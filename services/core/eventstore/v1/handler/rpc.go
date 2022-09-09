package handler

import (
	"context"
	"encoding/json"
	"errors"

	s "eventstore/service"

	pb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eventStoreHandler struct {
	logger  l.Logger
	service s.Service
}

func NewEventStoreHandler(logger l.Logger, service s.Service) *eventStoreHandler {
	return &eventStoreHandler{
		logger, service,
	}
}

func (h *eventStoreHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	logger := h.logger.WithRPCContext(ctx)
	// populate logger with event fields
	// NOTE: the EventData is not included as a field here because it could contain sensitive information
	logger = logger.WithFields(l.Fields{
		"aggregate_type": req.GetAggregateType(),
		"event_type":     req.GetEventType(),
		"event_code":     req.GetEventCode(),
		"aggregate_id":   req.GetAggregateId(),
	})
	logger.Debug("handling create event request")

	// validate the event
	err := validateRequest(logger, req)
	if err != nil {
		logger.WithFields(err.Fields()).Warn("rejecting create event request due to invalid request values")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// call service layer to create and publish event
	id, err := h.service.Create(ctx, logger, req)
	if err != nil {
		logger.WithFields(err.Fields()).WithError(err).Error("failed to create event")
		// TODO: should we wrap one more time and return the error to the client? What shape will it take if we do that?
		return nil, status.Error(codes.Internal, "failed to either publish or save event during creation, this is most likely due to a network or infrastructure issue like RabbitMQ or MongoDB being down so a retry is suggested")
	}

	// return the event id and whether or not the event was published
	// if the calling client required the event to be published, then this allows it to know whether that was successful
	return &pb.CreateResponse{
		Id: id,
	}, nil
}

// HELPER FUNCTIONS

// validateRequest makes sure the event passed into Create() is valid and returns an error if it is not
func validateRequest(logger l.Logger, req *pb.CreateRequest) l.Error {

	// validate aggregate type
	if req.GetAggregateType() == "" {
		return logger.WrapError(errors.New("aggregate type is a required field but was not set"))
	}

	// validate event type
	if req.GetEventType() == "" {
		return logger.WrapError(errors.New("event type is a required field but was not set"))
	}

	// validate event code
	if req.GetEventCode() == "" {
		return logger.WrapError(errors.New("event code is a required field but was not set"))
	}

	// validate aggregate id
	_, err := uuid.Parse(req.GetAggregateId())
	if err != nil {
		return logger.WrapError(errors.New("invalid aggregate id, must be a valid UUID"))
	}

	// validate event data
	validJson := json.Valid([]byte(req.GetEventData()))
	if !validJson {
		return logger.WrapError(errors.New("invalid event data, must be valid JSON"))
	}

	// everything in the request is valid
	return nil
}
