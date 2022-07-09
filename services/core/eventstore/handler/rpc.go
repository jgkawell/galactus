package handler

import (
	"context"

	s "eventstore/service"

	pb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type eventStoreHandler struct {
	logger  l.Logger
	service s.EventStoreService
}

func NewEventStoreHandler(logger l.Logger, service s.EventStoreService) *eventStoreHandler {
	return &eventStoreHandler{
		logger, service,
	}
}

func (h *eventStoreHandler) Create(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	logger := h.logger.WithRPCContext(ctx)
	logger.Debug("create method was called")

	err := req.ValidateAll()
	if err != nil {
		logger.WithError(err).Error("request validation failed")
		return nil, err
	}

	id, isPublished, customErr := h.service.Create(ctx, logger, req.GetEvent())
	if customErr != nil {
		logger.WithError(customErr).WithFields(customErr.Fields()).Error("failed to create event")
		return nil, customErr.Unwrap()
	}

	return &pb.CreateEventResponse{
		Id:          id,
		IsPublished: isPublished,
	}, nil
}
