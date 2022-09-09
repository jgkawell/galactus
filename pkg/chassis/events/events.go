package events

import (
	"context"
	"encoding/json"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	es "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	ev "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type EventManager interface {
	CreateAndSendEvent(ctx context.Context, logger l.Logger, ed interface{}, id string, at ev.AggregateType, et ev.EventType, ec string) l.Error
	GetEventAndMessageData(ctx context.Context, logger l.Logger, msg messagebus.Message) (*agpb.Event, []byte, l.Error)
	ThrowSystemError(ctx context.Context, logger l.Logger, ed interface{}, id string, systemError *ev.SystemError)
}

type Manager struct {
	Client es.EventStoreClient
}

// CreateAndSendEvent - Create a sourced event and store it with the `EventStore`
//
//   ed - interface of data to be marshalled to json for the `EventData` payload.
//   id - `aggregate_id` value used to create a correlation between models, user actions, and state changes
//   at - `aggregate_type` union of aggregates of the system
//   et - `event_type` an event message able to be marshaled/unmarshaled to and from JSON
func (m *Manager) CreateAndSendEvent(ctx context.Context, logger l.Logger, ed interface{}, id string, at ev.AggregateType, et ev.EventType, ec string) l.Error {
	j, err := json.Marshal(ed)
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to marshall event data into json"))
	}
	req := es.CreateRequest{
		AggregateType: AggregateTypeAsString(at),
		EventType:     EventTypeAsString(&et),
		EventCode:     EventCodeAsString(ec),
		AggregateId:   id,
		EventData:     string(j),
	}
	logger = logger.WithField("request", req)
	resp, err := m.Client.Create(ctx, &req)
	if err != nil {
		return logger.WithError(err).WrapError(l.NewError(err, "failed to create event source"))
	}
	logger.WithField("event_id", resp.GetId()).Debug("created event")
	return nil
}

// GetEventAndMessageData - A utility to unpacked event type, and event data from queued message. Event data is nested
// within event, returned here in event and as []byte for ease of parsing and passing by callers
func (m *Manager) GetEventAndMessageData(ctx context.Context, logger l.Logger, msg messagebus.Message) (*agpb.Event, []byte, l.Error) {
	var evt agpb.Event
	if err := msg.GetMessage(&evt); err != nil {
		e := ev.SystemError{Code: ev.SystemErrorCode_SYSTEM_ERROR_CODE_MALFORMED_EVENT_DATA}
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
func (m *Manager) ThrowSystemError(ctx context.Context, logger l.Logger, ed interface{}, aggregateId string, systemError *ev.SystemError) {
	aggregateType := ev.AggregateType_AGGREGATE_TYPE_SYSTEM
	eventType := ev.EventType{ Code: &ev.EventType_SystemCode{} }
	eventCode := EventCodeAsString(ev.SystemEventCode_SYSTEM_EVENT_CODE_ERROR)
	err := m.CreateAndSendEvent(ctx, logger, ed, aggregateId, aggregateType, eventType, eventCode)
	if err != nil {
		// Only log the error. If the `CreateAndSendEvent` method errors out then we are unable to produce system errors and something awful has happened.
		// Logging is all we will be able to do, and should trigger an alert in whatever monitoring tool is used (e.g. Datadog).
		logger.WithError(err).Error(ev.SystemErrorCode_SYSTEM_ERROR_CODE_FAILED_EVENT_PUBLISH.String())
	}
}
