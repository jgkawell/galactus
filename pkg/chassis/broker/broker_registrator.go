package broker

import (
	"context"

	"github.com/jgkawell/galactus/pkg/chassis/messagebus"

	l "github.com/jgkawell/galactus/pkg/logging/v2"
)

type BrokerDefinition struct {
	Exchange   string
	RoutingKey string
	QueueName  string
	Handler    messagebus.ClientHandler
}

func RegisterConsumer(logger l.Logger, bus messagebus.MessageBus, definition BrokerDefinition) {
	logger = logger.WithField("queue_definition", definition)
	logger.Info("establishing listener to queue")

	// make the response channel
	done := make(chan messagebus.DoneChannelResponse, 1)

	// establish consumer on queue
	err := bus.Consume(
		context.Background(),
		definition.QueueName,
		definition.RoutingKey,
		definition.Handler,
		done)
	if err != nil {
		logger.WithError(err).WithField("queue", definition).Fatal("failed to register queue listener")
	}

	// wait for the consumer to exit
	go func() {
		select {
		case response := <-done:
			if response.Error != nil {
				logger.WithError(response.Error).Fatal("queue consumer exited with error")
			}
			if !response.Done {
				logger.Fatal("queue consumer exited without done=true")
			}
			if response.Message != "" {
				logger.Info(response.Message)
			}
			logger.WithField("queue", definition).Debug("queue consumer exited")
		}
	}()
}
