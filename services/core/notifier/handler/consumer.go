package handler

import (
	"context"
	"errors"

	s "notifier/service"

	es "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	ev "github.com/circadence-official/galactus/pkg/chassis/events"
	mb "github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type consumer struct {
	logger  l.Logger
	service s.Service
}

type Consumer interface {
	mb.ClientHandler
	Deliverer
}

type Deliverer interface {
	DeliverNotification(context.Context, *es.Event) error
}

func NewConsumer(logger l.Logger, service s.Service) Consumer {
	return &consumer{
		logger, service,
	}
}

var (
	ErrorFailedGetEventAndMessageData = "failed to get message data from queued message"
	ErrorFailedUnmarshal              = "failed to unmarshal request"
	ErrorFailedToParseMessageData     = "failed to parse message data from event"
	ErrorFailedToDeliverNotification  = "failed to deliver notification"
	WarningMessageTypeNotMatched      = "warning type not matched"
)

// Handle - Subscribed to the `Notifier` exchange and processing messages of type `NOTIFICATION_DELIVERY_REQUESTED`
func (c *consumer) Handle(ctx context.Context, msg mb.Message) error {
	logger := c.logger.WithField("handle_incoming_message", msg)

	evt, _, err := ev.GetEventAndMessageData(ctx, logger, msg)
	if err != nil {
		logger.WithError(err).Error(ErrorFailedGetEventAndMessageData)
		return err
	}

	if err := c.DeliverNotification(ctx, evt); err != nil {
		logger.WithError(err).Error(ErrorFailedToParseMessageData)
		return err
	}

	return nil
}

func (c *consumer) DeliverNotification(ctx context.Context, event *es.Event) error {
	if err := c.service.Deliver(ctx, c.logger, event); err != nil {
		c.logger.WithError(err).Error(ErrorFailedToDeliverNotification)
		return errors.New(ErrorFailedToDeliverNotification)
	}

	return nil
}
