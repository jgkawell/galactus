package service

import (
	chpb "github.com/circadence-official/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"

	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type service struct {
	eventStoreClient espb.EventStoreClient
}

type Service interface {
	Apply(ct.ExecutionContext, *chpb.ApplyCommandRequest) (*chpb.ApplyCommandResponse, l.Error)
}

func NewService(eventStoreClient espb.EventStoreClient) Service {
	return &service{
		eventStoreClient: eventStoreClient,
	}
}

const (
	ErrorEventStoreCreate = "failed to create event"
)

// Processing a command in this respect means to convert the command request to an event, and create
// a new event in the eventstore. The eventstore will then broker the message to the correct data stream
func (s *service) Apply(ctx ct.ExecutionContext, cmdReq *chpb.ApplyCommandRequest) (*chpb.ApplyCommandResponse, l.Error) {
	ctx.GetLogger().Debug("applying command")

	eventRequest := &espb.CreateEventRequest{
		Event: &espb.Event{
			TransactionId: ctx.GetTransactionID(),
			Publish:       true, // all command events must be published for a consumer to act on the event
			EventType:     cmdReq.GetEventType(),
			AggregateType: cmdReq.GetAggregateType(),
			AggregateId:   cmdReq.GetAggregateId(),
			EventData:     cmdReq.GetCommandData(),
		},
	}

	eventResponse, err := s.eventStoreClient.Create(ctx.GetContext(), eventRequest)
	if err != nil {
		return nil, ctx.GetLogger().WrapError(l.NewError(err, ErrorEventStoreCreate))
	}
	ctx.GetLogger().WithField("event_id", eventResponse.GetId()).Debug("event created")

	res := &chpb.ApplyCommandResponse{
		Id:            cmdReq.GetAggregateId(),
		TransactionId: ctx.GetTransactionID(),
	}
	return res, nil
}
