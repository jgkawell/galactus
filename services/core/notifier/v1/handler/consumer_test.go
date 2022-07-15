package handler

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"notifier/service"

	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	mb "github.com/circadence-official/galactus/pkg/chassis/messagebus"
	ut "github.com/circadence-official/galactus/pkg/chassis/util"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	spy "bou.ke/monkey"
)

const (
	Success = "Handle Success, NOTIFICATION_DELIVERED"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		testName string
		testMsg  *mb.MockQueuedMessage
		event    *espb.Event
	}{
		{
			testName: Success,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 2,
					},
				},
				EventData:     "test",
				TransactionId: "123",
			},
		},
		{
			testName: ErrorFailedToParseMessageData,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 1,
					},
				},
				EventData:     "test",
				TransactionId: "123",
			},
		},
		{
			testName: ErrorFailedToDeliverNotification,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 1,
					},
				},
				EventData:     "test",
				TransactionId: "123",
			},
		},
		{
			testName: WarningMessageTypeNotMatched,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 99,
					},
				},
				EventData:     "test",
				TransactionId: "123",
			},
		},
	}

	defer spy.UnpatchAll()
	for _, tc := range testCases {
		// setup
		logger, _ := l.CreateNullLogger()
		s := service.MockService{}
		c := NewConsumer(logger, &s)

		spy.Patch(
			ut.GetEventAndMessageData,
			func(ctx context.Context, logger l.Logger, msg mb.QueuedMessage) (*espb.Event, []byte, error) {
				if tc.testName == ErrorFailedGetEventAndMessageData {
					assert.Equal(t, tc.testMsg, msg, tc.testName)
					assert.NotNil(t, ctx, tc.testName)
					assert.NotNil(t, logger, tc.testName)
					return nil, nil, errors.New(tc.testName)
				} else {
					return tc.event, []byte(tc.event.GetEventData()), nil
				}
			})

		spy.PatchInstanceMethod(
			reflect.TypeOf(c), "DeliverNotification",
			func(_ *consumer, ctx context.Context, evt *espb.Event) error {
				assert.NotNil(t, ctx, tc.testName)
				assert.NotNil(t, evt, tc.testName)

				if tc.testName == ErrorFailedToParseMessageData {
					return errors.New(tc.testName)
				}

				if tc.testName == ErrorFailedToDeliverNotification {
					return errors.New(tc.testName)
				}

				return nil
			})

		// test
		err := c.Handle(context.Background(), tc.testMsg)

		// verify
		if tc.testName != Success && tc.testName != WarningMessageTypeNotMatched {
			assert.Equal(t, errors.New(tc.testName), err, tc.testName)
		} else {
			evt, data, _ := ut.GetEventAndMessageData(context.Background(), logger, tc.testMsg)
			assert.Equal(t, string(data), tc.event.GetEventData(), tc.testName)
			assert.Equal(t, evt, tc.event, tc.testName)
			assert.Nil(t, err, tc.testName)
		}

		spy.Unpatch(ut.GetEventAndMessageData)
	}
}

func TestDeliverNotification(t *testing.T) {
	testCases := []struct {
		testName string
		testMsg  *mb.MockQueuedMessage
		event    *espb.Event
	}{
		{
			testName: ErrorFailedToDeliverNotification,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 1,
					},
				},
				EventData:     `"{"test": "test"}"`,
				TransactionId: "123",
			},
		},
		{
			testName: Success,
			event: &espb.Event{
				Timestamp: 0,
				EventType: &espb.EventType{
					Code: &espb.EventType_NotificationCode{
						NotificationCode: 1,
					},
				},
				EventData:     `"{"book":{"name":"Harry Potter and the Goblet of Fire","author":"J. K. Rowling","year":2000,"genre":"Fantasy Fiction","bestseller":true}}"`,
				TransactionId: "123",
			},
		},
	}

	defer spy.UnpatchAll()
	for _, tc := range testCases {
		// setup
		logger, _ := l.CreateNullLogger()
		s := service.MockService{}
		c := NewConsumer(logger, &s)

		// mocks
		mockDeliver := s.On("Deliver", mock.Anything, mock.Anything)
		if tc.testName == ErrorFailedToDeliverNotification {
			fmt.Println("test_name: ", tc.testName)
			mockDeliver.Return(errors.New(tc.testName))
		} else {
			mockDeliver.Return(nil)
		}

		err := c.DeliverNotification(context.Background(), tc.event)

		fmt.Println("what is the error: ", err, tc.testName, tc.event)

		if tc.testName == ErrorFailedToDeliverNotification {
			assert.NotNil(t, err, tc.testName)
		} else {
			assert.Nil(t, err, tc.testName)
		}
	}
}
