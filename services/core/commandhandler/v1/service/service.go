package service

import (
	chpb "github.com/jgkawell/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"
)

type service struct {
	eventStoreClient espb.EventStoreClient
}

type Service interface {
	Apply(ct.ExecutionContext, *chpb.ApplyCommandRequest) l.Error
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
func (s *service) Apply(ctx ct.ExecutionContext, cmdReq *chpb.ApplyCommandRequest) l.Error {
	ctx.Logger.Debug("applying command")

	// build out event request
	eventRequest := &espb.CreateRequest{
		AggregateType: cmdReq.GetAggregateType(),
		EventType:     cmdReq.GetEventType(),
		EventCode:     cmdReq.GetEventCode(),
		AggregateId:   cmdReq.GetAggregateId(),
		EventData:     cmdReq.GetCommandData(),
	}

	// send create event request
	eventResponse, err := s.eventStoreClient.Create(ctx.GetContext(), eventRequest)
	if err != nil {
		return ctx.Logger.WrapError(l.NewError(err, ErrorEventStoreCreate))
	}
	ctx.Logger.WithField("event_id", eventResponse.GetId()).Debug("event created")

	return nil
}
