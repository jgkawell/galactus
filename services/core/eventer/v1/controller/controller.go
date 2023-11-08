package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	pb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/events"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	l "github.com/jgkawell/galactus/pkg/logging"

	eventspb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Eventer interface {
	PublishAndSave(ctx ct.Context, event *eventspb.CloudEvent) (string, l.Error)
}

type eventer struct {
	db  *gorm.DB
	bus messagebus.Client
	env string
}

func New(db *gorm.DB, mb messagebus.Client, env string) Eventer {
	return &eventer{
		db, mb, env,
	}
}

func (e *eventer) PublishAndSave(ctx ct.Context, event *eventspb.CloudEvent) (string, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	receivedTime := time.Now()
	err := e.publish(ctx, event)
	publishedTime := time.Now()
	if err != nil {
		return "", ctx.Logger().WrapError(err)
	}
	go e.save(ctx, event, receivedTime, publishedTime)
	return event.Id, nil
}

// publish attempts to publish the event to the messagebus with an exponential backoff and returns an error on terminal failure
func (e *eventer) publish(ctx ct.Context, event *eventspb.CloudEvent) l.Error {
	ctx, span := ctx.Span()
	defer span.End()

	var err error
	success := false
	for i := 0; i < 5; i++ {
		err = e.bus.Publish(ctx, messagebus.PublishParams{
			Event: event,
			// Tags: TODO,
		})
		if err == nil {
			success = true
			break
		}
		ctx.Logger().WithError(err).Error("failed to publish event to messagebus. retrying...")
		wait := time.Duration(math.Pow(2, float64(i+1)))
		time.Sleep(wait * time.Second)
	}
	if !success {
		return ctx.Logger().WrapError(l.NewError(err, "failed to publish event to messagebus after exhausting all retries"))
	}
	return nil
}

// save attempts to save the event to the database with an exponential backoff and handles terminal failures
func (e *eventer) save(ctx ct.Context, event *eventspb.CloudEvent, received, published time.Time) {
	ctx, span := ctx.Span()
	defer span.End()

	data, err := json.Marshal(event.Data)
	if err != nil {
		ctx.Logger().WithError(err).Error("failed to marshal event data to json")
		e.failSave(ctx, event, received, published, err)
	}

	success := false
	for i := 0; i < 5; i++ {
		err = e.db.Create(&pb.EventORM{
			Id:            event.Id,
			ReceivedTime:  &received,
			PublishedTime: &published,
			EventSource:   event.Source,
			EventType:     event.Type,
			EventData:     string(data),
			// TODO: how to handle this?
			TransactionId: uuid.New().String(),
		}).Error
		if err == nil {
			success = true
			break
		}
		ctx.Logger().WithError(err).Error("failed to save event to database. retrying...")
		wait := time.Duration(math.Pow(2, float64(i+1)))
		time.Sleep(wait * time.Second)
	}

	if !success {
		ctx.Logger().WithError(err).Error("failed to save event to database after exhausting all retries")
		e.failSave(ctx, event, received, published, err)
	}
}

// failSave attempts to emit an event when the event terminally fails to save to the database
func (e *eventer) failSave(ctx ct.Context, event *eventspb.CloudEvent, received, published time.Time, originalErr error) {
	ctx, span := ctx.Span()
	defer span.End()

	newEvent, newErr := events.New(&pb.SaveFailed{
		Event:         event,
		Error:         originalErr.Error(),
		ReceivedTime:  timestamppb.New(received),
		PublishedTime: timestamppb.New(published),
	})
	if newErr != nil {
		err := fmt.Errorf("original error: %s - events.New error: %s", originalErr.Error(), newErr.Error())
		ctx.Logger().WithError(err).Error("failed to create SaveFailed event while processing error from failed json marshal of event data during save")
		return
	}
	newErr = e.bus.Publish(ctx, messagebus.PublishParams{
		Event: newEvent,
		// Tags: TODO,
	})
	if newErr != nil {
		err := fmt.Errorf("json.Marshal error: %s - events.New error: %s", originalErr.Error(), newErr.Error())
		ctx.Logger().WithError(err).Error("failed to publish SaveFailed event while processing error from failed json marshal of event data during save")
	}
}
