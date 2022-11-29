package service

import (
	"fmt"
	"time"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx ct.ExecutionContext, request *pb.CreateRequest) (string, l.Error)
}

type service struct {
	db  *gorm.DB
	mb  messagebus.MessageBus
	env string
}

func NewService(db *gorm.DB, mb messagebus.MessageBus, env string) Service {
	return &service{
		db, mb, env,
	}
}

// Create - Create a new event that will be stored, published to an exchange (if requested), and updated with its published timestamp
func (s *service) Create(ctx ct.ExecutionContext, req *pb.CreateRequest) (string, l.Error) {

	// create and set id and add it to the logger
	rTime := time.Now()
	event := &agpb.EventORM{
		Id: uuid.NewString(),
		AggregateId: req.GetAggregateId(),
		AggregateType: req.GetAggregateType(),
		EventData: req.GetEventData(),
		EventType: req.GetEventType(),
		EventCode: req.GetEventCode(),
		ReceivedTime: &rTime,
		TransactionId: ctx.GetTransactionID(),
	}
	ctx.Logger = ctx.Logger.WithFields(l.Fields{
		"event_id": event.Id,
	})
	ctx.Logger.Debug("creating event")

	// attempt to publish event
	eventPB, stdErr := event.ToPB(ctx.GetContext())
	if stdErr != nil {
		return event.Id, ctx.Logger.WrapError(l.NewError(stdErr, "failed to convert event from ORM to PB"))
	}
	err := s.publishEvent(ctx, eventPB)
	if err != nil {
		// if publishEvent returns an error, that means there was an unexpected system error.
		// we should thus return an error to the client as they may wish to create a new event to reattempt to publish
		// NOTE: this also means the event is NOT stored in the database
		return event.Id, ctx.Logger.WrapError(err)
	}
	pTime := time.Now()
	event.PublishedTime = &pTime

	// save event to db
	stdErr = s.db.Create(event).Error
	if stdErr != nil {
		// if Create returns an error, that means there was an unexpected system error.
		// the publish succeeded so simply log the error and continue.
		ctx.Logger.WrappedError(err, "failed to save event to database")
	}

	ctx.Logger.Debug("successfully created event")
	return event.Id, nil
}

func (s *service) publishEvent(ctx ct.ExecutionContext, event agpb.Event) l.Error {
	ctx.Logger.Debug("attempting to publish event")
	stdErr := s.mb.SendMessage(ctx.GetContext(), "local.default", generateRoutingKey(event.GetAggregateType(), event.GetEventType(), event.GetEventCode()), event)
	if stdErr != nil {
		return ctx.Logger.WrapError(l.NewError(stdErr, "failed to publish event as message on messagebus"))
	}
	ctx.Logger.Debug("event published")
	return nil
}

func generateRoutingKey(aggregateType, eventType, eventCode string) string {
	return fmt.Sprintf("%s.%s.%s", aggregateType, eventType, eventCode)
}
