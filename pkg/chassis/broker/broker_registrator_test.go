package broker

import (
	"errors"
	"testing"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	"github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterQueueListener(t *testing.T) {
	testCases := []struct {
		testName string
		err      error
	}{
		{testName: "with errors", err: errors.New("err")},
		{testName: "without errors"},
	}
	for _, tt := range testCases {
		log, _ := logging.CreateNullLogger()
		bus := &messagebus.MockMessageBus{}
		listener := &messagebus.MockClientHandler{}
		definition := &BrokerDefinition{}
		t.Run(tt.testName, func(t *testing.T) {
			bus.On("ListenToQueue", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, tt.err).Once()
			bus.On("RegisterQueue", mock.Anything, mock.Anything).Return(nil).Once()
			if tt.err != nil {
				require.Panics(t, func() {
					RegisterConsumer(log, bus, listener, definition)
				})
			} else {
				require.NotPanics(t, func() {
					RegisterConsumer(log, bus, listener, definition)
				})
			}

		})
	}
}

func TestRegisterTopicListener(t *testing.T) {
	testCases := []struct {
		testName string
		err      error
	}{
		{testName: "with errors", err: errors.New("err")},
		{testName: "without errors"},
	}
	for _, tt := range testCases {
		log, _ := logging.CreateNullLogger()
		bus := &messagebus.MockMessageBus{}
		listener := &messagebus.MockClientHandler{}
		definition := &BrokerDefinition{}
		t.Run(tt.testName, func(t *testing.T) {
			bus.On("ListenToTopic", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, tt.err).Once()
			bus.On("RegisterTopic", mock.Anything, mock.Anything).Return(nil).Once()
			if tt.err != nil {
				require.Panics(t, func() {
					RegisterConsumer(log, bus, listener, definition)
				})
			} else {
				require.NotPanics(t, func() {
					RegisterConsumer(log, bus, listener, definition)
				})
			}

		})
	}
}
