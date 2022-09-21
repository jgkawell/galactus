package handler

import (
	"context"

	s "notifier/service"

	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	ev "github.com/circadence-official/galactus/pkg/chassis/events"
	mb "github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type consumer struct {
	logger  l.Logger
	service s.Service
	em      ev.EventManager
}

type Consumer interface {
	mb.ClientHandler
}

func NewConsumer(logger l.Logger, service s.Service, em ev.EventManager) Consumer {
	return &consumer{
		logger, service, em,
	}
}

var (
	ErrorFailedGetEventAndMessageData = "failed to get message data from queued message"
	ErrorFailedToDeliverNotification  = "failed to deliver notification"
)

// Handle - Subscribed to the `Notifier` exchange and processing messages of type `NOTIFICATION_DELIVERY_REQUESTED`
func (h *consumer) Handle(c context.Context, msg mb.Message) error {
	h.logger.Debug("handling message")

	evt, _, err := h.em.GetEventAndMessageData(h.logger, msg)
	if err != nil {
		h.logger.WrappedError(err, "failed to get event and message data from queued message")
		// reject since notifier listens to a topic and so is the only consumer on it's unique queue
		msg.Reject()
		return err
	}
	ctx, err := ct.NewExecutionContextFromEvent(c, h.logger, evt)
	if err != nil {
		h.logger.WrappedError(err, "failed to create execution context from event")
		msg.Reject()
		return err
	}

	err = h.service.Deliver(*ctx, evt)
	if err != nil {
		ctx.Logger.WrappedError(err, "failed to deliver message")
		// reject since notifier listens to a topic and so is the only consumer on it's unique queue
		msg.Reject()
		return err
	}

	msg.Complete()
	return nil
}
