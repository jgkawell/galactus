package events

import (
	"context"
	"encoding/json"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"

	es "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	ev "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

// CreateAndSendEvent - Create a sourced event and store it with the `EventStore`
//
//   ed - interface of data to be marshalled to json for the `EventData` payload.
//   id - `aggregate_id` value used to create a correlation between models, user actions, and state changes
//   at - `aggregate_type` union of aggregates of the system
//   et - `event_type` an event message able to be marshaled/unmarshaled to and from JSON/BSON
func CreateAndSendEvent(ctx context.Context, logger l.Logger, client es.EventStoreClient, ed interface{}, id string, at ev.AggregateType, et ev.EventType) l.Error {
	logger = logger.WithField("event_data", ed)
	j, err := json.Marshal(ed)
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to marshall into json"))
	}
	req := es.CreateEventRequest{
		Event: &es.Event{
			EventType:     1, // TODO: &et,
			AggregateId:   id,
			AggregateType: int64(at.Number()),
			EventData:     string(j),
		},
	}
	logger = logger.WithField("request", req)
	_, err = client.Create(ctx, &req)
	if err != nil {
		return logger.WithError(err).WrapError(l.NewError(err, "failed to create event source"))
	}
	return nil
}

// CreateAndSendEventWithTransactionID - Create a sourced event and store it with the `EventStore`
//
//   ed - interface of data to be marshalled to json for the `EventData` payload.
//   aggregateID - root `aggregate_id` value used to create a correlation between models, user actions, and state changes
//   transactionID - transaction ID for traceability from client all through the system
//   at - `aggregate_type` union of aggregates of the system
//   et - `event_type` an event message able to be marshaled/unmarshaled to and from JSON/BSON
func CreateAndSendEventWithTransactionID(ctx context.Context, logger l.Logger, client es.EventStoreClient, ed interface{}, aggregateID, transactionID string, at ev.AggregateType, et ev.EventType) l.Error {
	logger = logger.WithField("event_data", ed)
	j, err := json.Marshal(ed)
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to marshall into json"))
	}
	req := es.CreateEventRequest{
		Event: &es.Event{
			EventType:     1, // TODO: &et,
			AggregateId:   aggregateID,
			TransactionId: transactionID,
			AggregateType: int64(at.Number()),
			EventData:     string(j),
		},
	}
	logger = logger.WithField("request", req)
	_, err = client.Create(ctx, &req)
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to create event source"))
	}
	return nil
}

// GetEventAndMessageData - A utility to unpacked event type, and event data from queued message. Event data is nested
// within event, returned here in event and as []byte for ease of parsing and passing by callers
func GetEventAndMessageData(ctx context.Context, logger l.Logger, msg messagebus.Message) (*es.Event, []byte, l.Error) {
	var evt es.Event
	if err := msg.GetMessage(&evt); err != nil {
		e := ev.SystemError{Code: ev.SystemErrorCode_INVALID_SYSTEM_MESSAGE_DATA}
		return nil, nil, logger.WrapError(l.NewError(err, e.Code.String()))
	}
	logger.Debug(evt.GetEventData())

	return &evt, []byte(evt.GetEventData()), nil
}

// ThrowSystemError - Create a system level error that is related to core components of the system, and not application level
//
// @ed - interface of data to provided error level context
// @id - aggregate_id value used to create a correlation between models, user actions, and state changes
// @systemError - a systemError code that builds a relationship between core system components and errors that might have happened.
func ThrowSystemError(ctx context.Context, logger l.Logger, client es.EventStoreClient, ed interface{}, id string, systemError *ev.SystemError) {
	eventType := ev.EventType{
		Code: &ev.EventType_SystemCode{
			SystemCode: ev.SystemEventCode_SYSTEM_ERROR,
		},
	}
	if err := CreateAndSendEvent(ctx, logger, client, ed, id, ev.AggregateType_AGGREGATE_TYPE_SYSTEM, eventType); err != nil {
		// Only log the error. If the `CreateAndSendEvent` method errors out then we are unable to produce system errors and something awful has happened.
		// Logging is all we will be able to do, and should trigger an alert.
		logger.WithError(err).Error(ev.SystemErrorCode_FAILED_EVENT_PUBLISH.String())
	}
}
