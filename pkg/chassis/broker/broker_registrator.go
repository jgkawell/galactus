package broker

import (
	"context"
	"time"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"

	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

func RegisterQueueSender(logger l.Logger, bus messagebus.MessageBus, definition *BrokerDefinition) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := bus.RegisterQueue(ctx, definition.QueueName, definition.Exchange, definition.RoutingKey); err != nil {
		logger.WithError(err).Panic("failed to register queue sender")
	}
}

func RegisterTopicSender(logger l.Logger, bus messagebus.MessageBus, definition *BrokerDefinition) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := bus.RegisterTopic(ctx, definition.QueueName, definition.Exchange, definition.RoutingKey); err != nil {
		logger.WithError(err).Panic("failed to register topic sender")
	}
}

func RegisterConsumer(logger l.Logger, bus messagebus.MessageBus, listener messagebus.ClientHandler, definition *BrokerDefinition) {
	logger = logger.WithField("queue_definition", definition)
	logger.Info("establishing listener to queue")

	// make the response channel
	done := make(chan messagebus.DoneChannelResponse, 1)

	// establish consumer on queue
	err := bus.Consume(
		context.Background(),
		definition.QueueName,
		definition.RoutingKey,
		listener,
		done)
	if err != nil {
		logger.WithError(err).WithField("queue", definition).Panic("failed to register queue listener")
	}

	// wait for the consumer to exit
	go func() {
		select {
		case response := <-done:
			if response.Error != nil {
				logger.WithError(response.Error).Panic("queue consumer exited with error")
			}
			if !response.Done {
				logger.Panic("queue consumer exited without done=true")
			}
			if response.Message != "" {
				logger.Info(response.Message)
			}
			logger.WithField("queue", definition).Debug("queue consumer exited")
		}
	}()
}
