package amqp

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	events "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

const (
	exchange = "primary"
)

type (
	wrapper struct {
		connection        *amqp.Connection // the persistent server connection
		configKey         string
		consumers         map[string]string // map of consumer IDs to channel names
		persistentChannel *amqp.Channel     // the channel used throughout the lifetime of the client (management/publish). consuming goroutines should create their own channels: https://www.rabbitmq.com/tutorials/amqp-concepts.html#amqp-channels
	}
)

// New instantiates a new client wrapper. A call to Initialize is required before use.
// The configKey parameter dictates which key in the configuration file will be read during
// initialization. If configKey is empty, the default value of "amqp.url" will be used.
// The configuration can be in various formats, but the following is an example of a yaml file:
//
//	amqp:
//	  url: 'amqp://localhost:5672'
func New(configKey string) messagebus.Client {
	if configKey == "" {
		configKey = "amqp.url"
	}
	return &wrapper{
		configKey: configKey,
		consumers: make(map[string]string),
	}
}

func (w *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	url := config.GetString(w.configKey)
	// attempt to open connection to the AMQP server
	var err error
	w.connection, err = amqp.Dial(url)
	if err != nil {
		return err
	}

	// attempt to open a channel on the connection
	ch, err := w.connection.Channel()
	if err != nil {
		return err
	}
	w.persistentChannel = ch

	// create (no-op if already exists) the exchange on the AMQP server
	err = w.persistentChannel.ExchangeDeclare(
		exchange, // name: the name of the exchange
		"direct", // kind: ref - https://www.rabbitmq.com/tutorials/amqp-concepts.html#exchange-direct
		true,     // durable: we want to keep the exchange around if the server restarts
		false,    // autoDelete: we DO NOT want the exchange to be deleted when there are no queues bound to it (is this true???)
		false,    // internal: we want to allow publishing to this exchange
		false,    // noWait: we want to wait for the server to respond to this declaration request before continuing
		nil,      // arguments: the arguments to include with this declaration request
	)
	if err != nil {
		return err
	}

	return nil
}

func (w *wrapper) Publish(ctx context.Context, params messagebus.PublishParams) error {
	// serialize the event
	data, err := proto.Marshal(params.Event)
	if err != nil {
		return err
	}

	// publish the message to the exchange
	err = w.persistentChannel.PublishWithContext(
		ctx,      // ctx: the context for the publish request
		exchange, // exchange: the name of the exchange
		routingKey(params.Event, false, params.Tags...), // routingKey: the routing key
		false, // mandatory: we don't require delivery of the message to a queue
		false, // immediate: what is this?
		amqp.Publishing{
			Timestamp: time.Now(),
			MessageId: params.Event.Id,
			Body:      data,
		})
	if err != nil {
		return err
	}
	return nil
}

func (w *wrapper) Subscribe(ctx context.Context, params messagebus.SubscribeParams) error {
	identifier := uuid.New().String()

	// declare the queue on the AMQP server
	queueName := identifier
	exclusive := true
	if params.Group != "" {
		queueName = fmt.Sprintf("%s.%s", routingKey(params.Event, params.IgnoreType, params.Tags...), params.Group)
		exclusive = false
	}
	err := w.registerQueue(ctx, queueName, routingKey(params.Event, params.IgnoreType, params.Tags...), exclusive)
	if err != nil {
		return err
	}

	// open a new channel on the existing connection
	channel, err := w.connection.Channel()
	if err != nil {
		return err
	}

	// establish a consume connection to receive messages off of the queue into this channel
	msgs, err := channel.Consume(
		queueName,  // queue: the name of the queue
		identifier, // consumer: the id of the consumer
		false,      // autoAck: we want to manually ack messages after we've processed them
		false,      // exclusive: allow sharing of this queue by multiple consumers
		false,      // noLocal: this is not supported by RabbitMQ
		false,      // noWait: we want to wait for the server to respond to this consume request before continuing
		nil,        // args: no arguments needed
	)
	if err != nil {
		return err
	}

	// now that we've successfully connected, save the consumer id for later use in consumer.Cancel() call
	w.consumers[identifier] = queueName

	// watch for a closed or cancelled signal on this channel
	go func() {
		nclose := make(chan *amqp.Error)
		channel.NotifyClose(nclose)

		ncancel := make(chan string)
		channel.NotifyCancel(ncancel)

		var msg string
		select {
		case err := <-nclose:
			if err != nil {
				msg = err.Error()
			} else {
				msg = "channel closed with no error"
			}
		case tag := <-ncancel:
			msg = fmt.Sprintf("channel with tag (%s) cancelled: ", tag)
		}
		// TODO: return an error (from the message above) to a channel given back to the caller
		fmt.Println(msg)
	}()

	// watch for messages on this channel
	go func() {
		for msg := range msgs {
			// spin off the message handling in a separate goroutine to parallelize message consumption
			go process(ctx, params.Consumer, msg)
		}
	}()

	return nil
}

// registerQueue registers a queue on the AMQP server
func (w *wrapper) registerQueue(ctx context.Context, queueName, routingKey string, exclusive bool) error {

	args := amqp.Table{
		"x-expires": 1000 * 60 * 60 * 24, // 1 day
	}

	// declare the queue on the AMQP server
	_, err := w.persistentChannel.QueueDeclare(
		queueName,
		true, // durable: we want to keep the queue around if the server restarts
		true, // autoDelete: we want to delete the queue after all consumers have disconnected
		exclusive,
		false, // noWait: we want to wait for the server to respond to this declaration request before continuing
		args,
	)
	if err != nil {
		return err
	}

	// bind the queue to the exchange on the AMQP server
	err = w.persistentChannel.QueueBind(
		queueName,
		routingKey,
		exchange,
		false, // noWait: we want to wait for the server to respond to this bind request before continuing
		nil,   // arguments: no arguments needed
	)
	if err != nil {
		return err
	}

	return nil
}

// process sends the queue message to the ClientHandler interface for handling
func process(ctx context.Context, consumer messagebus.Consumer, msg amqp.Delivery) {
	if msg.Body != nil {
		// deserialize the event
		event := &events.CloudEvent{}
		err := proto.Unmarshal(msg.Body, event)
		if err != nil {
			msg.Nack(false, false)
		}
		// consume the event
		err = consumer.Consume(ctx, event)
		if err != nil {
			msg.Nack(false, true)
		}
		msg.Ack(false)
	}
}

func (w *wrapper) Unsubscribe(ctx context.Context, params messagebus.UnsubscribeParams) error {
	err := w.persistentChannel.Cancel(routingKey(params.Event, params.IgnoreType, params.Tags...), false)
	if err != nil {
		return err
	}
	delete(w.consumers, routingKey(params.Event, params.IgnoreType, params.Tags...))
	return nil
}

func (w *wrapper) Shutdown(force bool) error {
	// will close() the deliveries channel
	for _, consumer := range w.consumers {
		err := w.persistentChannel.Cancel(consumer, force)
		if err != nil {
			return err
		}
	}
	// close the connection
	// ignore the error as the connection is treated as closed regardless
	_ = w.connection.Close()
	return nil
}

// HELPERS

// routingKey returns the routing key for the given event and tags. The format will always follow the pattern:
//
// If no tags are provided and ignoreType is true: {event.Source}.#
// If no tags are provided and ignoreType is false: {event.Source}.{event.Type}.#
// If tags are provided and ignoreType is true: {event.Source}.*.{tags}
// If tags are provided and ignoreType is false: {event.Source}.{event.Type}.{tags}
//
// And the tags will be sorted alphabetically.
func routingKey(event *events.CloudEvent, ignoreType bool, tags ...string) string {
	if len(tags) == 0 {
		if ignoreType {
			return fmt.Sprintf("%s.#", event.Source)
		}
		return fmt.Sprintf("%s.%s.#", event.Source, event.Type)
	}
	sort.Strings(tags)
	if ignoreType {
		return fmt.Sprintf("%s.*.%s", event.Source, strings.Join(tags, "."))
	}
	return fmt.Sprintf("%s.%s.%s", event.Source, event.Type, strings.Join(tags, "."))
}
