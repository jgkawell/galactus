package broker

import (
	"context"
	"fmt"

	messagebus "github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type BrokerType int

type BrokerDefinition struct {
	Exchange         string `mapstructure:"exchange"`
	RoutingKey       string `mapstructure:"routingkey"`
	QueueName        string `mapstructure:"queuename"`
}

const (
	_ BrokerType = iota
	BrokerTypeQueue
	BrokerTypeTopic
)

type BrokerSender interface {
	Send(ctx context.Context, i interface{}) error
}

type brokerSender struct {
	bus        messagebus.MessageBus
	definition *BrokerDefinition
	brokerType BrokerType
}

func NewBrokerSender(logger l.Logger, bus messagebus.MessageBus, definition *BrokerDefinition, brokerType BrokerType) BrokerSender {
	sender := &brokerSender{
		bus:        bus,
		definition: definition,
		brokerType: brokerType,
	}

	if brokerType == BrokerTypeQueue {
		RegisterQueueSender(logger, bus, definition)
	} else if brokerType == BrokerTypeTopic {
		RegisterTopicSender(logger, bus, definition)
	} else {
		panic(fmt.Sprintf("unknown broker type %d", brokerType))
	}
	return sender
}

func (s *brokerSender) Send(ctx context.Context, i interface{}) error {
	return s.bus.SendMessage(ctx, s.definition.Exchange, s.definition.RoutingKey, i)
}
