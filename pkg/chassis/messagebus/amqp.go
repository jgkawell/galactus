package messagebus

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// MessageBus provides a connection to the message bus service.
type MessageBus interface {
	// Connect attempts to connect to to the AMQP server
	Connect(ctx context.Context, logger l.Logger) l.Error

	// RegisterExchange creates and configures an exchange on the AMQP server
	RegisterExchange(ctx context.Context, exchange string, kind ExchangeKind) error

	// RegisterQueue creates and configures a queue on the AMQP server
	RegisterQueue(ctx context.Context, queueName, exchangeName, routingKey string) error

	// RegisterTopic creates and configures a topic on the AMQP server
	RegisterTopic(ctx context.Context, identifier, exchangeName, routingKey string) error

	// SendMessage puts a message on the given exchange with a given routing key
	SendMessage(ctx context.Context, exchange, routingKey string, message interface{}) error

	// Consume... TODO
	Consume(ctx context.Context, queueName, routingKey string, h ClientHandler, done chan<- DoneChannelResponse) error

	// CancelConsumerChannels cancels all active listening consumer channels.
	CancelConsumerChannels(noWait bool) error

	// CancelConsumerChannelsBySuffix cancels all active listening consumer channels that have the given suffix substring in there consumer ID.
	// If suffix is blank, then all channels are cancelled.
	CancelConsumerChannelsBySuffix(suffix string, noWait bool) error

	// Shutdown cancels all consumer channels and closes the connection to the messagebus
	Shutdown(noWait bool) error
}

type messageBus struct {
	username          string            // username for auth
	password          string            // password for auth
	serverIPs         string            // this is a comma separated list of IP addresses
	connection        *amqp.Connection  // the persistent server connection
	consumers         map[string]string // map of consumer IDs to channel names
	persistentChannel *amqp.Channel     // the channel used throughout the lifetime of the client (management/publish). consuming goroutines should create their own channels: https://www.rabbitmq.com/tutorials/amqp-concepts.html#amqp-channels
}

// NewMessageBus creates...
func NewMessageBus(logger l.Logger, u, p, s string) MessageBus {
	return &messageBus{
		username:  u,
		password:  p,
		serverIPs: s,
		consumers: make(map[string]string),
	}
}

func (mb *messageBus) Connect(ctx context.Context, logger l.Logger) (l.Error) {

	// convert the comma separated list of IP addresses to a slice and randomize it to "load balance" the connections
	connectionStrings := GenerateConnectionStrings(mb.username, mb.password, mb.serverIPs)

	for _, connection := range connectionStrings {
		logger := logger.WithField("connection", connection)

		// attempt to open connection to the AMQP server
		var err error
		mb.connection, err = amqp.Dial(connection)
		if err != nil {
			logger.WithError(err).Error("failed to connect to broker on a certain ip address. trying next ip address...")
			continue // try the next connection
		}

		// attempt to open a channel on the connection
		ch, err := mb.connection.Channel()
		if err != nil {
			logger.WithError(err).Error("failed to open the management channel after connecting to the broker. trying next ip address...")
			continue // try the next connection
		}
		mb.persistentChannel = ch

		// we only want to connect to one IP, so this return exits the loop early when successfully connected
		logger.Info("connected to broker")
		return nil
	}
	return logger.WrapError(errors.New("failed to connect to broker on any connections"))
}

func (mb *messageBus) RegisterExchange(ctx context.Context, exchange string, kind ExchangeKind) error {

	// this configures the "original" kind of exchange running underneath the ExchangeKindDelayed
	// since we're using the `rabbitmq-delayed-message-exchange` plugin
	args := amqp.Table{
		rabbitMqDelayedExchangeArg: kind.String(), // the original type of exchange
	}

	// create (no-op if already exists) the exchange on the AMQP server
	err := mb.persistentChannel.ExchangeDeclare(
		exchange,                     // name: the name of the exchange
		ExchangeKindDelayed.String(), // kind: ref - https://www.rabbitmq.com/tutorials/amqp-concepts.html#exchange-direct
		true,                         // durable: we want to keep the exchange around if the server restarts
		false,                        // autoDelete: we DO NOT want the exchange to be deleted when there are no queues bound to it (is this true???)
		false,                        // internal: we want to allow publishing to this exchange
		false,                        // noWait: we want to wait for the server to respond to this declaration request before continuing
		args,                         // arguments: the arguments to include with this declaration request
	)
	if err != nil {
		return l.NewError(err, "failed to declare exchange")
	}
	return nil
}

// RegisterQueue registers a queue on the AMQP server. In this context, Queue refers to a load balanced unicast approach to consumers of the queue.
//
// In other words, when a message is routed from an exchange to a queue, the queue will round-robin select an attached consumer to handle the message.
func (mb *messageBus) RegisterQueue(ctx context.Context, queueName, exchangeName, routingKey string) error {
	// declare the queue on the AMQP server
	_, err := mb.persistentChannel.QueueDeclare(
		queueName, // name: the name of the queue
		true,      // durable: we want to keep the queue around if the server restarts
		false,     // autoDelete: we don't want to delete queues after all consumers drop. instead this will be handled on the unregister call
		false,     // exclusive: we want multiple consumers to attach (load balancing is round robin and handled by the server)
		false,     // noWait: we want to wait for the server to respond to this declaration request before continuing
		nil,       // arguments: no arguments needed
	)
	if err != nil {
		return l.NewError(err, "failed to declare queue")
	}

	// bind the queue to the exchange on the AMQP server
	err = mb.persistentChannel.QueueBind(
		queueName,    // name: the name of the queue
		routingKey,   // key: the routing key for the queue which decides how messages are pulled off the exchange and into this queue
		exchangeName, // exchange: the name of the exchange
		false,        // noWait: we want to wait for the server to respond to this bind request before continuing
		nil,          // arguments: no arguments needed
	)
	if err != nil {
		return l.NewError(err, "failed to bind queue to exchange")
	}

	return nil
}

// RegisterTopic registers a topic on the AMQP server. In this context, Topic refers to a load balanced multicast approach to consumers of the topic.
//
// Functionally this is almost the same as RegisterQueue(), HOWEVER the queueName passed should be UNIVERSALLY UNIQUE. This insures that you don't
// accidentally try and have multiple consumers on the same queue. Note that this function will set the queue to autoDelete=true and exclusive=true so that
// these unique queues are cleaned up as consumers disconnect and that we don't have multiple consumers attached to the same queue which could lead to
// some nasty runtime bugs.
func (mb *messageBus) RegisterTopic(ctx context.Context, queueName, exchangeName, routingKey string) error {

	args := amqp.Table{
		"x-expires": 1000 * 60 * 60 * 24, // 1 day
	}

	// declare the queue on the AMQP server
	_, err := mb.persistentChannel.QueueDeclare(
		queueName, // name: the name of the queue
		true,      // durable: we want to keep the queue around if the server restarts
		true,      // autoDelete: we DO want to delete the queue after the ONLY consumer disconnects. since topics are a single consumer per queue, this makes sure not to end up with a flood of queues on the AMQP server.
		false,     // exclusive: we want to only allow one consumer to attach to this queue since there is a 1:1 relationship between the topic and a consumer. HOWEVER, since the regristry will create this queue we have to make it non-exclusive so the consumer can connection.
		false,     // noWait: we want to wait for the server to respond to this declaration request before continuing
		args,      // arguments: extra arguments on declaration
	)
	if err != nil {
		return l.NewError(err, "failed to declare queue")
	}

	// bind the queue to the exchange on the AMQP server
	err = mb.persistentChannel.QueueBind(
		queueName,    // name: the name of the queue
		routingKey,   // key: the routing key for the queue which decides how messages are pulled off the exchange and into this queue
		exchangeName, // exchange: the name of the exchange
		false,        // noWait: we want to wait for the server to respond to this bind request before continuing
		nil,          // arguments: no arguments needed
	)
	if err != nil {
		return l.NewError(err, "failed to bind queue to exchange")
	}

	return nil
}

func (mb *messageBus) Consume(ctx context.Context, queueName, routingKey string, h ClientHandler, done chan<- DoneChannelResponse) error {
	// the consumer id consists of a generated UUID suffixed with the given routingKey
	// this allows for us to selective cancel consumers based on the routingKey to facilitate blue/green deployments
	consumerID := fmt.Sprintf("%s-%s", uuid.New().String(), routingKey)

	// open a new channel on the existing connection
	channel, err := mb.connection.Channel()
	if err != nil {
		return l.NewError(err, "failed to open channel on connection")
	}

	// establish a consume connection to receive messages off of the queue into this channel
	msgs, err := channel.Consume(
		queueName,  // queue: the name of the queue
		consumerID, // consumer: the id of the consumer
		false,      // autoAck: we want to manually ack messages after we've processed them
		false,      // exclusive: allow sharing of this queue by multiple consumers (topics are not shared but that's handled in the QueueDeclare() call)
		false,      // noLocal: this is not supported by RabbitMQ
		false,      // noWait: we want to wait for the server to respond to this consume request before continuing
		nil,        // args: no arguments needed
	)
	if err != nil {
		return l.NewError(err, "failed to establish consume on channel")
	}

	// now that we've successfully connected, save the consumer id for later use in consumer.Cancel() call
	mb.consumers[consumerID] = queueName

	// watch for a closed or cancelled signal on this channel
	go func() {
		nclose := make(chan *amqp.Error)
		channel.NotifyClose(nclose)

		ncancel := make(chan string)
		channel.NotifyCancel(ncancel)

		response := DoneChannelResponse{}
		select {
		case err := <-nclose:
			if err != nil {
				response.Error = err
			} else {
				response.Message = "channel closed with no error"
			}
		case tag := <-ncancel:
			response.Error = fmt.Errorf("channel with tag (%s) cancelled: ", tag)
		}
		done <- response
	}()

	// watch for messages on this channel
	go func() {
		for msg := range msgs {
			message := &message{
				acked: false,
				msg:   &msg,
			}
			if message.msg.Body != nil {
				err = h.Handle(ctx, message)
				if err != nil {
					// TODO: do we need logic to make sure we don't hit an infinite retry loop?
					message.Retry()
				} else {
					// TODO: implicit ack?
					message.Complete()
				}
			}
		}
		done <- DoneChannelResponse{
			Done:    true,
			Message: "why did this happen?",
		}
	}()

	return nil
}

func (mb *messageBus) SendMessage(ctx context.Context, exchange, routingKey string, message interface{}) error {

	// generate a unique message id
	messageID := uuid.New().String()

	// serialize the message to JSON
	body, err := serializeMessage(message)
	if err != nil {
		return l.NewError(err, "failed to serialize message")
	}

	// publish the message to the exchange
	err = mb.persistentChannel.Publish(
		exchange,   // exchange: the name of the exchange
		routingKey, // routingKey: the routing key
		false,      // mandatory: we don't require delivery of the message to a queue
		false,      // immediate: what is this?
		amqp.Publishing{
			Type:        getMessageType(message),
			Timestamp:   time.Now(),
			MessageId:   messageID,
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return l.NewError(err, "failed to publish message to broker")
	}
	return nil
}

func (mb *messageBus) CancelConsumerChannels(noWait bool) error {
	return mb.CancelConsumerChannelsBySuffix("", noWait)
}

func (mb *messageBus) CancelConsumerChannelsBySuffix(suffix string, noWait bool) error {
	for consumerID, _ := range mb.consumers {
		if suffix == "" || strings.HasSuffix(consumerID, suffix) {
			// TODO: do we have a single channel or multiple?
			if err := mb.persistentChannel.Cancel(consumerID, noWait); err != nil {
				return l.NewError(err, "failed to cancel consumer channel")
			}
			delete(mb.consumers, consumerID)
		}
	}
	return nil
}

func (mb *messageBus) Shutdown(noWait bool) error {
	// will close() the deliveries channel
	if err := mb.CancelConsumerChannels(noWait); err != nil {
		return l.NewError(err, "failed to cancel consumer channels")
	}
	// close the connection
	if err := mb.connection.Close(); err != nil {
		return l.NewError(err, "failed to close connection")
	}
	return nil
}
