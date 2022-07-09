package service

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	ntpb "github.com/circadence-official/galactus/api/gen/go/core/notifier/v1"
	evpb "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"

	ev "github.com/circadence-official/galactus/pkg/chassis/events"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type service struct {
	isHeartbeatEnabled bool
	heartbeatTimeout   int
	esc                espb.EventStoreClient
	inputChannels      *inputChannels
}

// inputChannels - The main underlying data structure to hold all user connections
type inputChannels struct {
	sync.Mutex
	// `Channels` is a map of user id's as the key and another map of client id's a the key and a notification channel as the value.
	// This structure gives us the ability to allow a user to have more than one client connected at a time while being
	// able to deliver messages to all the users clients.
	Channels map[string]map[string]chan evpb.NotificationDeliveryRequested
}

// NewService - Create an instance of the `service` struct that is implementing the `Service` interface.
func NewService(isHeartbeatEnabled bool, heartbeatTimeout int, esc espb.EventStoreClient) Service {
	ic := &inputChannels{
		Channels: make(map[string]map[string]chan evpb.NotificationDeliveryRequested),
	}

	return &service{
		isHeartbeatEnabled: isHeartbeatEnabled,
		heartbeatTimeout:   heartbeatTimeout,
		esc:                esc,
		inputChannels:      ic,
	}
}

// `NOTIFICATION_BUFFER_RATE` is the value the is responsible for how many messages are in a channel at a time.
// If the channel contains more messages in a channel then this number the message will not be delivered, and an
// error will be returned to the producer on the channel. After performance review `8` seems to find a decent
// balance between memeory used, and message processing. It's worth noting if this value is changed more memory
// will be consumed per connected client.
const NOTIFICATION_BUFFER_RATE = 8

// Error messages
const (
	ErrorFailedToEmitEvent                       = "failed to emit event"
	ErrorFailedToFindChannels                    = "failed to locate client channels for user"
	ErrorFailedHeartbeatDelivery                 = "failed to send heartbeat to user"
	ErrorFailedUnmarshal                         = "failed to unmarshal request"
	ErrorFailedToSendNotification                = "failed to send notification"
	ErrorFailedToSendConfirmationEvent           = "failed to send confirmation event"
	ErrorFailedToRemoveUserConnection            = "failed to remove the users notificationChannel.ction"
	ErrorFailedToRemoveInputChannel              = "failed to remove the users input channel"
	ErrorFailedToBuildNotificationChannel        = "failed to build notification channel for a new grpc connection"
	ErrorFailedToAddNotificationChannelForClient = "failed to add notification channel for client"
	ErrorInvalidConnectionMetadata               = "failed to create connection user metadata is invalid"
	ErrorInvalidNotificationRequest              = "invalid notification sent by producer, actorID is not found"
)

const (
	ErrorDuplicateDevice          = "device already connected"
	ErrorActorConnectionsNotFound = "actor connections not found"
)

// Service - public interface for `Connecting`, `DisnotificationChannel.cting` a client with the server, and a `Deliver` method
// allowing a different component to send messages to a specified `User`
type Service interface {
	// Connect, and begin to process incoming messages for a client
	Connect(context.Context, l.Logger, *ntpb.ConnectionRequest, ntpb.Notifier_ConnectServer) (NotificationChannel, l.Error)

	// Deliver - receives and event from the `Notification` exchange and sends it to the correct client
	Deliver(context.Context, l.Logger, *espb.Event) l.Error

	// SpawnWorker - Spawn a worker routine to process each incoming message on it's own `NotificationChannel`
	SpawnWorker(logger l.Logger, notificationChannel NotificationChannel)

	// NewConnection - Starts, and stores a long term grpc server streaming connection with a specified client
	NewNotificationChannelForConnection(logger l.Logger, ck *ConnectionKey, stream ntpb.Notifier_ConnectServer) NotificationChannel

	// RemoveChannel - Cleans up the user channel when given a `ConnectionKey`
	RemoveChannel(ck *ConnectionKey)

	// Multicast - sends one notification to each one of the connected clients a user has in their current session.
	Multicast(logger l.Logger, actorID string, msg evpb.NotificationDeliveryRequested) l.Error

	// ReadInputChannels - read channel mapping- unit test helper func
	ReadInputChannels() map[string]map[string]chan evpb.NotificationDeliveryRequested

	// ReadConnectsion  - read connections mapping for given userId - unit test helper func
	ReadConnections(userId string) (map[string]chan evpb.NotificationDeliveryRequested, bool)
}

// Connect - when called will establish a long term connection with the client and provide a `NotificationChannel` that can be used to communicate
// with the running goroutine, and stop it's message processing.
func (s *service) Connect(ctx context.Context, logger l.Logger, req *ntpb.ConnectionRequest, stream ntpb.Notifier_ConnectServer) (NotificationChannel, l.Error) {
	// create configuration from request
	ck, err := s.NewConnectionKey(req.GetUserId(), req.GetClientId())
	if err != nil {
		return NotificationChannel{}, logger.WrapError(l.NewError(err, ErrorInvalidConnectionMetadata))
	}
	logger.WithFields(l.Fields{"client_id": ck.clientID, "actor_id": ck.actorID}).Debug("client connecting")

	// create the `NotificationChannel` for the worker to receive messages to work on
	nc := s.NewNotificationChannelForConnection(logger, ck, stream)
	logger.Debug("notification channel created")

	// begin the worker process
	s.SpawnWorker(logger, nc)

	return nc, logger.WrapError(<-nc.Error)
}

// NewNotificationChannelForConnection - Starts, and stores a long term grpc server streaming connection for a specified client
// once the connection is stored a channel to communicate to the goroutine is created and returned in the `NotificationChannel` struct
func (s *service) NewNotificationChannelForConnection(logger l.Logger, ck *ConnectionKey, stream ntpb.Notifier_ConnectServer) NotificationChannel {
	logger.WithField("new_client_connection", ck.clientID).Debug("client connection")

	s.inputChannels.Mutex.Lock()
	defer s.inputChannels.Mutex.Unlock()

	clients, uc := s.inputChannels.Channels[ck.actorID]
	if !uc {
		logger.WithField("actor_id", ck.actorID).Debug("adding user to channels map")
		clients = make(map[string]chan evpb.NotificationDeliveryRequested)
		s.inputChannels.Channels[ck.actorID] = clients
	}

	clients[ck.clientID] = make(chan evpb.NotificationDeliveryRequested, NOTIFICATION_BUFFER_RATE)

	return NotificationChannel{
		connectionKey: ck,
		stream:        stream,
		Error:         make(chan error),
		channel:       clients[ck.clientID],
	}
}

func (s *service) ReadInputChannels() map[string]map[string]chan evpb.NotificationDeliveryRequested {
	return s.inputChannels.Channels
}

func (s *service) ReadConnections(userId string) (map[string]chan evpb.NotificationDeliveryRequested, bool) {
	c, ok := s.inputChannels.Channels[userId]
	return c, ok
}

type ErrorChan chan<- error

// NotificationChannel - holds a `ConnectionKey`, the grpc stream `Notifier_ConnectServer`, an error channel, and a channel
// for communicating to the worker goroutine.
type NotificationChannel struct {
	connectionKey *ConnectionKey
	stream        ntpb.Notifier_ConnectServer
	Error         chan error
	channel       chan evpb.NotificationDeliveryRequested
}

// CreateHeartbeat - build a notification for testing
func (c *NotificationChannel) CreateHeartbeat(logger l.Logger) *ntpb.Notification {
	confirmation := &ntpb.Heartbeat{
		ClientId:           c.connectionKey.clientID,
		ExpirationDeadline: 0,
	}

	byt, err := json.Marshal(confirmation)
	if err != nil {
		logger.WithError(err).Error(ErrorFailedHeartbeatDelivery)
	}

	return &ntpb.Notification{
		NotificationType: ntpb.NotificationType_HEARTBEAT,
		Data:             string(byt),
	}
}

// RemoveChannel - Cleans up the user channel when given a `ConnectionKey`
func (s *service) RemoveChannel(ck *ConnectionKey) {
	s.inputChannels.Mutex.Lock()
	defer s.inputChannels.Mutex.Unlock()
	// The user can have more than one client, and thus we should check
	// for the specific client for the user before attempting to delete
	// the channel.
	_, ch := s.inputChannels.Channels[ck.actorID][ck.clientID]
	if ch {
		delete(s.inputChannels.Channels[ck.actorID], ck.clientID)
	}

	// Check if the map of user connections is equal to zero. That mean the previous
	// operation deleted the last connection for that user. If this is the case
	// lets clean up the user from the `Channels` map. The `Channels` map can also
	// be thought of as how many active users are connected to the service.
	if len(s.inputChannels.Channels[ck.actorID]) == 0 {
		_, u := s.inputChannels.Channels[ck.actorID]
		if u {
			delete(s.inputChannels.Channels, ck.actorID)
		}
	}
}

// Multicast - sends one notification to each one of the connected clients an actor has in their current session.
func (s *service) Multicast(logger l.Logger, actorID string, msg evpb.NotificationDeliveryRequested) l.Error {
	s.inputChannels.Mutex.Lock()
	defer s.inputChannels.Mutex.Unlock()

	connections, ok := s.ReadConnections(actorID)
	if !ok {
		return logger.WithField("actor_id", actorID).WrapError(errors.New(ErrorActorConnectionsNotFound))
	}

	for clientID, ch := range connections {
		logger.WithField("client_id", clientID).Debug("broadcast to channel")
		ch <- msg
	}

	return nil
}

// ConnectionKey - The configuration struct of a connection, meta data like `actorID`, `clientID`, and other options
// the actorID could be a user, or a minionD client
type ConnectionKey struct {
	actorID, clientID  string
	isHeartbeatEnabled bool
	heartbeatTimeout   int
}

// NewConnectionKey - validates the actorID, clientID, and returns a new `ConnectionKey`
func (s *service) NewConnectionKey(actorID, clientID string) (*ConnectionKey, error) {
	if actorID == "" || clientID == "" {
		return nil, errors.New(ErrorInvalidConnectionMetadata)
	}

	return &ConnectionKey{
		actorID:            actorID,
		clientID:           clientID,
		isHeartbeatEnabled: s.isHeartbeatEnabled,
		heartbeatTimeout:   s.heartbeatTimeout,
	}, nil
}

// Deliver - is a method called by the `Consumer` that will pipe the notification to the users `InputChannel`
// and send the message to the go routine maintaining the user's notificationChannel.action. When the message has been sent to
// the users notificationChannel.action a notification_delivered event is published.
func (s *service) Deliver(ctx context.Context, logger l.Logger, event *espb.Event) l.Error {
	data := []byte(event.GetEventData())
	var req evpb.NotificationDeliveryRequested
	if err := json.Unmarshal(data, &req); err != nil {
		return logger.WrapError(l.NewError(err, ErrorFailedUnmarshal))
	}

	//set event transaction id in the Notification request
	transactionID := event.GetTransactionId()
	req.GetNotification().TransactionId = transactionID

	actorID := req.GetActorId()
	if actorID == "" {
		return logger.WrapError(errors.New(ErrorInvalidNotificationRequest))
	}

	// TODO: implement unicast messaging
	// clientID := req.GetClientId()

	logger.WithField("actor_id", actorID).WithField("transaction_id", transactionID).Debug("notification delivery requested")

	/* // TODO: implement unicast messaging
	 * if clientID != "" {
	 *   // if client_id is provided `unicast`
	 *   if err := s.Inputs.Unicast(logger, actorID, clientID, req); err != nil {
	 *     logger.WithError(err).Error(FailedToFindChannels)
	 *     return err
	 *   }
	 * } else {
	 *   // else `multicast`
	 *   if err := s.Inputs.Multicast(logger, actorID, req); err != nil {
	 *     logger.WithError(err).Error(FailedToFindChannels)
	 *     return err
	 *   }
	 * }  */

	if err := s.Multicast(logger, actorID, req); err != nil {
		return logger.WrapError(err)
	}

	// emit notification delivered event
	r := &evpb.NotificationDelivered{}
	et := evpb.EventType{Code: &evpb.EventType_NotificationCode{NotificationCode: evpb.NotificationEventCode_NOTIFICATION_DELIVERED}}
	if err := ev.CreateAndSendEventWithTransactionID(ctx, logger, s.esc, r, actorID, event.GetTransactionId(), evpb.AggregateType_AGGREGATE_TYPE_NOTIFICATION, et); err != nil {
		return logger.WrapError(err)
	}

	logger.WithField("actor_id", actorID).WithField("transaction_id", transactionID).Debug("notification published to actor")

	return nil
}

// SpawnWorker - Run in a separate goroutine for each client notificationChannel.ction. When a notification is added to the
// client channel its sent to the user via their grpc notificationChannel.ction.
func (s *service) SpawnWorker(logger l.Logger, notificationChannel NotificationChannel) {
	var timer *time.Ticker
	if notificationChannel.connectionKey.isHeartbeatEnabled {
		timer = time.NewTicker(time.Duration(notificationChannel.connectionKey.heartbeatTimeout) * time.Second)
	}

	for {
		select {
		case <-timer.C:
			notification := notificationChannel.CreateHeartbeat(logger)

			if err := notificationChannel.stream.Send(notification); err != nil {
				logger.WithError(err).Error(ErrorFailedHeartbeatDelivery)
			}
		case <-notificationChannel.stream.Context().Done():
			s.RemoveChannel(notificationChannel.connectionKey)
			<-notificationChannel.Error
		case msg := <-notificationChannel.channel:
			logger.WithField("message", msg).Debug("notification delivery requested")

			if err := notificationChannel.stream.Send(msg.GetNotification()); err != nil {
				logger.WithError(err).Error("message delivery failed")
			}

			// fire off notification delivered message

			logger.WithField("message", msg).Debug("notification sent without error")
		}
	}
}
