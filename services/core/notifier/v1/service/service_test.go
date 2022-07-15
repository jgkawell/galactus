package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	ntpb "github.com/circadence-official/galactus/api/gen/go/core/notifier/v1"
	mm "github.com/circadence-official/galactus/pkg/chassis/clientfactory/mockclient"
	ut "github.com/circadence-official/galactus/pkg/chassis/util"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	spy "bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

const (
	InvalidConnectionValues      = "client did not connect with the correct values"
	FailedToAddConnection        = "Error - Failed to add connection"
	SuccessWorkerCreated         = "worker was created with a valid notification channel"
	SuccessNotificationDelivered = "successs notification was delivered"
)

func TestNewService(t *testing.T) {
	// Setup
	logger, _ := l.CreateNullLogger()
	//Test and Verify
	assert.NotNil(t, NewService(logger, true, 10, &mm.MockEventStoreClient{}))
}

func TestService_Multicast(t *testing.T) {
	type testCase int
	const (
		fail testCase = iota
		success
	)
	testCases := []struct {
		TestName string
		TestCase testCase
	}{
		{"fail multicast", fail},
		{"success", success},
	}

	for _, tc := range testCases {
		// Setup
		logger, hook := l.CreateNullLogger()
		expectedErr := ErrorActorConnectionsNotFound
		actorId := "123"
		clientId := "456"
		svc := NewService(logger, false, 0, nil)
		msg := espb.NotificationDeliveryRequested{
			ActorId:  actorId,
			ClientId: clientId,
			Notification: &ntpb.Notification{
				NotificationType: ntpb.NotificationType_HEARTBEAT,
				Data:             "data",
			},
		}
		// Spy service call - ReadConnections
		spy.PatchInstanceMethod(
			reflect.TypeOf(svc), "ReadConnections",
			func(_ *service, u string) (map[string]chan espb.NotificationDeliveryRequested, bool) {
				assert.Equal(t, actorId, u)
				if tc.TestCase == fail {
					return nil, false
				}
				return make(map[string]chan espb.NotificationDeliveryRequested), true
			})
		//test
		err := svc.Multicast(logger, actorId, msg)

		if tc.TestCase < success {
			assert.NotNil(t, err)
			assert.Equal(t, expectedErr, err)
			if assert.True(t, len(hook.Entries) > 0, tc.TestName) {
				assert.Equal(t, "error", hook.LastEntry().Level.String(), tc.TestName)
			}
		} else {
			assert.Nil(t, err)
		}
	}
}
func TestService_Connect(t *testing.T) {
	// test cases
	testCases := []struct {
		name             string
		expectedUserID   string
		expectedClientID string
		request          *ntpb.ConnectionRequest
		stream           *mm.MockNotifier_ConnectServer
		connectionKey    ConnectionKey
	}{
		{
			name:             InvalidConnectionValues,
			expectedUserID:   "",
			expectedClientID: "",
		},
		// NOTE: this makes the test run forever. Working through a way of testing this. The Error, or signal channel
		// that is given to the routine is not stoping the execution of the test, therefor it runs forever and blocks
		// the remaining tests. During runtime the close channel is signaled with the users disconnects and the routine
		// is closed.
		/* {
		 *   name:             FailedToAddConnection,
		 *   expectedUserID:   "userid",
		 *   expectedClientID: "clientd",
		 *   request: &ntpb.ConnectionRequest{
		 *     UserId:   uuid.New().String(),
		 *     ClientId: uuid.New().String(),
		 *   },
		 * }, */
	}

	defer spy.UnpatchAll()
	for _, tc := range testCases {
		// setup
		logger, _ := l.CreateNullLogger()
		esClient := &mm.MockEventStoreClient{}
		mockStream := &mm.MockNotifier_ConnectServer{}

		ms := mockStream.On("Send", mock.Anything)
		ms.Return(errors.New(""))

		s := NewService(logger, false, 0, esClient)

		spy.PatchInstanceMethod(
			reflect.TypeOf(s),
			"NewConnectionKey",
			func(_ *service, actorID, clientID string) (*ConnectionKey, error) {
				if tc.name == InvalidConnectionValues {
					assert.Equal(t, actorID, tc.expectedUserID, tc.name)
					assert.Equal(t, clientID, tc.expectedClientID, tc.name)
					return &ConnectionKey{}, errors.New(ErrorInvalidConnectionMetadata)
				}

				return &ConnectionKey{
					actorID:            tc.expectedUserID,
					clientID:           tc.expectedClientID,
					isHeartbeatEnabled: false,
					heartbeatTimeout:   0,
				}, nil
			})

		spy.PatchInstanceMethod(
			reflect.TypeOf(s),
			"NewNotificationChannelForConnection",
			func(_ *service, logger l.Logger, ck *ConnectionKey, stream ntpb.Notifier_ConnectServer) NotificationChannel {
				assert.NotNil(t, logger, tc.name)
				assert.NotNil(t, ck, tc.name)
				assert.NotNil(t, stream, tc.name)

				return NotificationChannel{}

			})

		spy.PatchInstanceMethod(
			reflect.TypeOf(s),
			"SpawnWorker",
			func(_ *service, nc NotificationChannel) {
				assert.NotNil(t, nc, tc.name)
				assert.NotNil(t, nc.Error, tc.name)
			})

		// test
		nc, err := s.Connect(context.Background(), tc.request, mockStream)

		// verify
		if tc.name == InvalidConnectionValues {
			assert.NotNil(t, err, tc.name)
		} else {
			assert.NotNil(t, nc, tc.name)
		}
	}
}

func TestService_Deliver(t *testing.T) {
	var notification = &ntpb.Notification{
		TransactionId: "transID",
	}
	var ed = espb.NotificationDeliveryRequested{
		Notification: notification,
		ActorId:      "actorID",
	}
	j, _ := json.Marshal(ed)
	// {"user_id":"id"}
	var eventValidUser = &espb.Event{
		EventData:     string(j),
		TransactionId: "12345",
	}
	testCases := []struct {
		name                      string
		testReq                   *espb.Event
		testNotificationDelivered *espb.NotificationDelivered
		testEventType             espb.EventType
	}{
		{
			name: "invalid actorID",
			testReq: &espb.Event{
				EventData:     "{\"user_id\":\"\"}",
				TransactionId: "123",
			},
		},
		{
			name:    ErrorFailedToFindChannels,
			testReq: eventValidUser,
		},
		{
			name:                      ErrorFailedToEmitEvent,
			testReq:                   eventValidUser,
			testNotificationDelivered: &espb.NotificationDelivered{},
			testEventType:             espb.EventType{Code: &espb.EventType_NotificationCode{NotificationCode: espb.NotificationEventCode_NOTIFICATION_DELIVERED}},
		},
		{
			name:                      ErrorFailedUnmarshal,
			testReq:                   eventValidUser,
			testNotificationDelivered: &espb.NotificationDelivered{},
			testEventType:             espb.EventType{Code: &espb.EventType_NotificationCode{NotificationCode: espb.NotificationEventCode_NOTIFICATION_DELIVERED}},
		},
		{
			name:                      SuccessNotificationDelivered,
			testReq:                   eventValidUser,
			testNotificationDelivered: &espb.NotificationDelivered{},
			testEventType:             espb.EventType{Code: &espb.EventType_NotificationCode{NotificationCode: espb.NotificationEventCode_NOTIFICATION_DELIVERED}},
		},
	}

	for _, tc := range testCases {
		//setup
		logger, _ := l.CreateNullLogger()
		esClient := &mm.MockEventStoreClient{}
		s := NewService(logger, false, 0, esClient)

		spy.PatchInstanceMethod(
			reflect.TypeOf(s),
			"Multicast",
			func(_ *service, logger l.Logger, actorID string, req espb.NotificationDeliveryRequested) error {
				assert.NotNil(t, req, tc.name)
				assert.NotNil(t, logger, tc.name)
				assert.NotEqual(t, actorID, "", tc.name)

				if tc.name == ErrorFailedToFindChannels {
					return errors.New(tc.name)
				}

				return nil
			})

		spy.Patch(ut.CreateAndSendEventWithTransactionID,
			func(ctx context.Context, logger l.Logger, esc espb.EventStoreClient, req interface{}, actorID string, transactionID string, aggregate espb.AggregateType, event espb.EventType) error {
				assert.NotNil(t, ctx, tc.name)
				assert.NotNil(t, logger, tc.name)
				assert.NotNil(t, req, tc.name)
				assert.NotNil(t, aggregate, tc.name)
				assert.NotNil(t, transactionID, tc.name)
				assert.NotNil(t, event, tc.name)
				assert.NotEqual(t, actorID, "", tc.name)

				if tc.name == ErrorFailedToEmitEvent {
					return errors.New(ErrorFailedToEmitEvent)
				}

				return nil
			})

		spy.Patch(
			json.Unmarshal,
			func(b []byte, i interface{}) error {
				if tc.name == ErrorFailedUnmarshal {
					return errors.New(tc.name)
				}
				i.(*espb.NotificationDeliveryRequested).ActorId = ""
				i.(*espb.NotificationDeliveryRequested).Notification = notification
				if tc.name == SuccessNotificationDelivered {
					i.(*espb.NotificationDeliveryRequested).ActorId = "id"
				}
				return nil
			})

		err := s.Deliver(context.Background(), tc.testReq)

		if tc.name == SuccessNotificationDelivered {
			assert.Nil(t, err, tc.name)
		} else {
			assert.NotNil(t, err, tc.name)
		}
	}
}

func TestMockService_NewNotificationChannelForConnection(t *testing.T) {
	//Setup
	logger, _ := l.CreateNullLogger()
	userId := "123"
	clientId := "456"
	ck := ConnectionKey{
		actorID:            userId,
		clientID:           clientId,
		isHeartbeatEnabled: false,
		heartbeatTimeout:   0,
	}

	svc := NewService(logger, false, 0, nil)
	//Test
	nc := svc.NewNotificationChannelForConnection(logger, &ck, nil)
	ic := svc.ReadInputChannels()

	assert.NotNil(t, nc)
	assert.Equal(t, ck.actorID, userId)
	assert.Equal(t, ck.clientID, clientId)
	assert.Equal(t, nc.connectionKey, &ck)
	assert.NotNil(t, nc.Error)
	assert.Equal(t, nc.channel, ic[userId][clientId])
}

func TestService_RemoveChannel(t *testing.T) {
	logger, _ := l.CreateNullLogger()
	userId := "123"
	clientId := "456"
	ck := ConnectionKey{
		actorID:            userId,
		clientID:           clientId,
		isHeartbeatEnabled: false,
		heartbeatTimeout:   0,
	}
	svc := NewService(logger, false, 0, nil)

	//Test - call disconnect with created test channel
	svc.RemoveChannel(&ck)

	//Verify
	ic := svc.ReadInputChannels()
	assert.Nil(t, ic[userId][clientId])
	assert.Nil(t, ic[userId])
}

func TestService_ReadInputChannels(t *testing.T) {
	//Setup
	logger, _ := l.CreateNullLogger()
	svc := NewService(logger, false, 0, nil)
	//Test
	ic := svc.ReadInputChannels()
	assert.NotNil(t, ic)
}

func TestService_SpawnWorker(t *testing.T) {
	t.Skip()
	testCases := []struct {
		name                    string
		testNotificationChannel NotificationChannel
	}{
		{
			name: "called",
			testNotificationChannel: NotificationChannel{
				connectionKey: &ConnectionKey{
					actorID:            "actorID",
					clientID:           "clientID",
					isHeartbeatEnabled: false,
					heartbeatTimeout:   5,
				},
				stream:  nil,
				Error:   make(chan error),
				channel: make(chan espb.NotificationDeliveryRequested),
			},
		},
	}

	spy.Patch(
		ntpb.Notifier_ConnectServer.Send,
		func(notifier ntpb.Notifier_ConnectServer, notification *ntpb.Notification) error {
			fmt.Println("did I get called?")
			return nil
		})

	for _, tc := range testCases {
		// setup
		logger, _ := l.CreateNullLogger()
		s := NewService(logger, false, 0, nil)

		s.SpawnWorker(tc.testNotificationChannel)

	}
}
