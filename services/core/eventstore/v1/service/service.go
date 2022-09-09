package service

import (
	"context"
	"fmt"
	"time"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"

	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, logger l.Logger, request *pb.CreateRequest) (string, l.Error)
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
func (s *service) Create(ctx context.Context, logger l.Logger, req *pb.CreateRequest) (string, l.Error) {

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
	}
	logger = logger.WithFields(l.Fields{
		"event_id": event.Id,
	})
	logger.Debug("creating event")

	// attempt to publish event
	eventPB, stdErr := event.ToPB(ctx)
	if stdErr != nil {
		return event.Id, logger.WrapError(l.NewError(stdErr, "failed to convert event from ORM to PB"))
	}
	err := s.publishEvent(ctx, logger, eventPB)
	if err != nil {
		// if publishEvent returns an error, that means there was an unexpected system error.
		// we should thus return an error to the client as they may wish to create a new event to reattempt to publish
		// NOTE: this also means the event is NOT stored in the database
		return event.Id, logger.WrapError(err)
	}
	pTime := time.Now()
	event.PublishedTime = &pTime

	// save event to db
	stdErr = s.db.Create(event).Error
	if stdErr != nil {
		// if Create returns an error, that means there was an unexpected system error.
		// the publish succeeded so simply log the error and continue.
		logger.WrappedError(err, "failed to save event to database")
	}

	logger.Debug("successfully created event")
	return event.Id, nil
}

func (s *service) publishEvent(ctx context.Context, logger l.Logger, event agpb.Event) l.Error {
	logger.Debug("attempting to publish event")
	stdErr := s.mb.SendMessage(ctx, "local.default", generateRoutingKey(event.GetAggregateType(), event.GetEventType(), event.GetEventCode()), event)
	if stdErr != nil {
		return logger.WrapError(l.NewError(stdErr, "failed to publish event as message on messagebus"))
	}
	logger.Debug("event published")
	return nil
}

func generateRoutingKey(aggregateType, eventType, eventCode string) string {
	return fmt.Sprintf("%s.%s.%s", aggregateType, eventType, eventCode)
}
