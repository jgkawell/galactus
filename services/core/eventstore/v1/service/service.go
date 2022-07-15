package service

import (
	"context"
	"errors"
	"fmt"

	es "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	db "github.com/circadence-official/galactus/pkg/chassis/db"
	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
)

var (
	errInvalidAggregateType = errors.New("invalid aggregate type in event")
	errInvalidEventType     = errors.New("invalid event type in event")
)

type EventStoreService interface {
	Create(ctx context.Context, log l.Logger, event *es.Event) (string, bool, l.Error)
}

type service struct {
	dao db.CrudDao
	mb  messagebus.MessageBus
	env string
}

func NewService(dao db.CrudDao, mb messagebus.MessageBus, env string) EventStoreService {
	return &service{
		dao, mb, env,
	}
}

// Create - Create a new event that will be stored, published to an exchange (if requested), and updated with its published timestamp
func (s *service) Create(ctx context.Context, logger l.Logger, event *es.Event) (string, bool, l.Error) {
	logger = logger.WithFields(l.Fields{
		"event_type":     event.GetEventType(),
		"aggregate_id":   event.GetAggregateId(),
		"aggregate_type": event.GetAggregateType(),
		"transaction_id": event.GetTransactionId(),
	})

	logger.Info("create event")
	if event == nil {
		return "", false, logger.WrapError(errors.New("event is nil"))
	}

	// create and set id
	event.EventId = uuid.New().String()
	logger = logger.WithField("event_id", event.GetEventId())

	// set timestamp
	event.ReceivedDate = timestamppb.Now()

	// get and check the validity of the provided aggregate type
	logger = logger.WithField("aggregate_type", event.GetAggregateType())
	// all we check for here is if the aggregate_type is 0 (invalid) since the client is
	// responsible for providing a valid aggregate_type that matches a proto definition
	if event.GetAggregateType() == 0 {
		return "", false, logger.WrapError(errInvalidAggregateType)
	}
	logger.Debug("aggregate type is valid")

	// get and check the validity of the provided event type
	logger = logger.WithField("event_type", event.GetEventType())
	// all we check for here is if the event_type is 0 (invalid) since the client is
	// responsible for providing a valid event_type that matches a proto definition
	if event.GetEventType() == 0 {
		return "", false, logger.WrapError(errInvalidEventType)
	}
	logger.Debug("event type is valid")

	// if the caller requests the event to be published, publish it
	if event.GetPublish() {
		// put the message on the messagebus with the mapping:
		// 	 - aggregate_type = exchange
		// 	 - event_type     = routing_key
		logger.WithFields(l.Fields{
			"exchange":    fmt.Sprint(event.GetAggregateType()),
			"routing_key": fmt.Sprint(event.GetEventType()),
			"event":       event,
		}).Debug("publishing event")
		err := s.mb.SendMessage(ctx, fmt.Sprintf("%s.%d", s.env, event.GetAggregateType()), fmt.Sprint(event.GetEventType()), event)
		if err != nil {
			return "", false, logger.WrapError(err)
		}
		// save the published timestamp
		event.PublishedDate = timestamppb.Now()
		logger.WithField("published_date", event.PublishedDate).Debug("generated published date")
	}

	// save the event to the db asynchronously since it's not a critical operation and if it fails the event creation is still successful
	go func() {
		id, err := s.dao.Create(ctx, logger, event)
		if err != nil {
			logger.WithError(err).Error("failed to store the event after it's been published")
		}
		logger.WithField("model_id", id).Debug("event stored in database")
	}()

	return event.GetEventId(), true, nil
}
