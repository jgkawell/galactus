package broker

import (
	"context"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"

	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type BrokerDefinition struct {
	Exchange         string `mapstructure:"exchange"`
	RoutingKey       string `mapstructure:"routingkey"`
	QueueName        string `mapstructure:"queuename"`
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
