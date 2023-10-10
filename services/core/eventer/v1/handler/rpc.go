package handler

import (
	"context"
	"errors"

	"eventer/controller"

	pb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	logger  l.Logger
	service controller.Eventer
}

func New(logger l.Logger, service controller.Eventer) *handler {
	return &handler{
		logger, service,
	}
}

func (h *handler) Emit(c context.Context, req *pb.EmitRequest) (*pb.EmitResponse, error) {
	logger := h.logger.WithRPCContext(c)
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, logger)
	if err != nil {
		ctx.Logger.WrappedError(err, "failed to create execution context")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// validate the event
	err = validateRequest(*ctx, req)
	if err != nil {
		ctx.Logger.WithFields(err.Fields()).Warn("rejecting emit request due to invalid request values")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// populate logger with event fields
	logger = ctx.Logger.WithFields(l.Fields{
		"event_id":     req.Event.Id,
		"event_source": req.Event.Source,
		"event_type":   req.Event.Type,
	})

	// call service layer to create and publish event
	id, err := h.service.PublishAndSave(*ctx, req.Event)
	if err != nil {
		ctx.Logger.WithFields(err.Fields()).WithError(err).Error("failed to create event")
		// TODO: should we wrap one more time and return the error to the client? What shape will it take if we do that?
		return nil, status.Error(codes.Internal, "failed to either publish or save event during creation, this is most likely due to a network or infrastructure issue like RabbitMQ or MongoDB being down so a retry is suggested")
	}

	// return the event id and whether or not the event was published
	// if the calling client required the event to be published, then this allows it to know whether that was successful
	return &pb.EmitResponse{
		Id: id,
	}, nil
}

// HELPER FUNCTIONS

// validateRequest makes sure the event passed into Create() is valid and returns an error if it is not
func validateRequest(ctx ct.ExecutionContext, req *pb.EmitRequest) l.Error {

	// validate event is not nil
	if req.Event == nil {
		return ctx.Logger.WrapError(errors.New("event is a required field but was not set"))
	}

	// validate event source
	if req.Event.Source == "" {
		return ctx.Logger.WrapError(errors.New("event source is a required field but was not set"))
	}

	// validate event type
	if req.Event.Type == "" {
		return ctx.Logger.WrapError(errors.New("event type is a required field but was not set"))
	}

	// everything in the request is valid
	return nil
}
