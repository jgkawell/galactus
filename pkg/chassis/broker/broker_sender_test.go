package broker

import (
	"context"
	"testing"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	"github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/undefinedlabs/go-mpatch"
)

func TestNewBrokerSender(t *testing.T) {
	testCases := []struct {
		testName   string
		brokerType BrokerType
	}{
		{testName: "brokerType = BrokerTypeQueue", brokerType: BrokerTypeQueue},
		{testName: "brokerType = BrokerTypeTopic", brokerType: BrokerTypeTopic},
		{testName: "with errors"},
	}
	rqsPatch, err := mpatch.PatchMethod(RegisterQueueSender, func(logger logging.Logger, bus messagebus.MessageBus, definition *BrokerDefinition) {
	})
	require.NoError(t, err)
	defer rqsPatch.Unpatch()
	rtsPatch, err := mpatch.PatchMethod(RegisterTopicSender, func(logger logging.Logger, bus messagebus.MessageBus, definition *BrokerDefinition) {
	})
	require.NoError(t, err)
	defer rtsPatch.Unpatch()
	for _, tt := range testCases {
		log, _ := logging.CreateNullLogger()
		bus := &messagebus.MockMessageBus{}
		definition := &BrokerDefinition{}
		t.Run(tt.testName, func(t *testing.T) {
			var sender BrokerSender
			if tt.brokerType == BrokerTypeQueue || tt.brokerType == BrokerTypeTopic {
				require.NotPanics(t, func() {
					sender = NewBrokerSender(log, bus, definition, tt.brokerType)
				})
				require.NotNil(t, sender)
			} else {
				require.Panics(t, func() {
					sender = NewBrokerSender(log, bus, definition, tt.brokerType)
				})
				require.Nil(t, sender)
			}
		})
	}
}

func TestBrokerSender_Send(t *testing.T) {
	testCases := []struct {
		testName   string
		err        error
		brokerType BrokerType
	}{
		{testName: "broker type BrokerTypeQueue", brokerType: BrokerTypeQueue},
		{testName: "broker type BrokerTypeTopic", brokerType: BrokerTypeTopic},
	}
	for _, tt := range testCases {
		ctx := context.TODO()
		bus := &messagebus.MockMessageBus{}
		definition := &BrokerDefinition{}
		sender := &brokerSender{
			bus:        bus,
			definition: definition,
			brokerType: tt.brokerType,
		}
		t.Run(tt.testName, func(t *testing.T) {
			if tt.brokerType == BrokerTypeQueue {
				bus.On("SendToQueue", ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			}
			bus.On("SendToTopic", ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			sendErr := sender.Send(ctx, nil)
			require.Equal(t, tt.err, sendErr)
		})
	}
}
