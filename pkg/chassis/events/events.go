package events

import (
	"encoding/json"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	es "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"
	ev "github.com/jgkawell/galactus/api/gen/go/generic/events/v1"

	c "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	l "github.com/jgkawell/galactus/pkg/logging/v2"
)

type EventManager interface {
	CreateAndSendEvent(ctx c.ExecutionContext, event interface{}, aggregateId string, aggregateType ev.AggregateType, eventType ev.EventType, eventCode string) l.Error
	GetEventAndMessageData(logger l.Logger, msg messagebus.Message) (*agpb.Event, []byte, l.Error)
	ThrowSystemError(ctx c.ExecutionContext, event interface{}, aggregateId string, systemError *ev.SystemError)
}

type Manager struct {
	Client es.EventStoreClient
}

// CreateAndSendEvent sends creates an event on the eventstore service
func (m *Manager) CreateAndSendEvent(ctx c.ExecutionContext, event interface{}, aggregateId string, aggregateType ev.AggregateType, eventType ev.EventType, eventCode string) l.Error {
	j, err := json.Marshal(event)
	if err != nil {
		return ctx.Logger.WrapError(l.NewError(err, "failed to marshall event data into json"))
	}
	req := es.CreateRequest{
		AggregateType: AggregateTypeAsString(aggregateType),
		EventType:     EventTypeAsString(&eventType),
		EventCode:     EventCodeAsString(eventCode),
		AggregateId:   aggregateId,
		EventData:     string(j),
	}
	ctx.Logger = ctx.Logger.WithField("request", req)
	resp, err := m.Client.Create(ctx.GetContext(), &req)
	if err != nil {
		return ctx.Logger.WithError(err).WrapError(l.NewError(err, "failed to create event source"))
	}
	ctx.Logger.WithField("event_id", resp.GetId()).Debug("created event")
	return nil
}

// GetEventAndMessageData unpacks an event type and the corresponding event data from a queued message. Event data is nested
// within the event as well as returned as a []byte for ease of parsing and passing by callers
func (m *Manager) GetEventAndMessageData(logger l.Logger, msg messagebus.Message) (*agpb.Event, []byte, l.Error) {
	var evt agpb.Event
	if err := msg.GetMessage(&evt); err != nil {
		return nil, nil, logger.WrapError(l.NewError(err, "failed to parse queued message into event type"))
	}
	logger.Debug(evt.GetEventData())

	return &evt, []byte(evt.GetEventData()), nil
}

// ThrowSystemError - Create a system level error that is related to core components of the system, and not application level
//
// @ed - interface of data to provided error level context
// @id - aggregate_id value used to create a correlation between models, user actions, and state changes
// @systemError - a systemError code that builds a relationship between core system components and errors that might have happened.
func (m *Manager) ThrowSystemError(ctx c.ExecutionContext, ed interface{}, aggregateId string, systemError *ev.SystemError) {
	aggregateType := ev.AggregateType_AGGREGATE_TYPE_SYSTEM
	eventType := ev.EventType{Code: &ev.EventType_SystemCode{}}
	eventCode := EventCodeAsString(ev.SystemEventCode_SYSTEM_EVENT_CODE_ERROR)
	err := m.CreateAndSendEvent(ctx, ed, aggregateId, aggregateType, eventType, eventCode)
	if err != nil {
		// Only log the error. If the `CreateAndSendEvent` method errors out then we are unable to produce system errors and something awful has happened.
		// Logging is all we will be able to do, and should trigger an alert in whatever monitoring tool is used (e.g. Datadog).
		ctx.Logger.WithError(err).Error(ev.SystemErrorCode_SYSTEM_ERROR_CODE_FAILED_EVENT_PUBLISH.String())
	}
}
