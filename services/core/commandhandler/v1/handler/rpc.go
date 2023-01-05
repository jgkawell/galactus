package handler

import (
	"context"

	s "commandhandler/service"

	pb "github.com/jgkawell/galactus/api/gen/go/core/commandhandler/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"

	"github.com/google/uuid"
)

type handler struct {
	logger  l.Logger
	service s.Service
}

func NewHandler(logger l.Logger, service s.Service) pb.CommandHandlerServer {
	return &handler{
		logger, service,
	}
}

// TODO: there needs to be a policy where we can check commands and data against the user calling this method
func (h *handler) Apply(c context.Context, req *pb.ApplyCommandRequest) (*pb.ApplyCommandResponse, error) {
	logger := h.logger.WithRPCContext(c)

	ctx, err := ct.NewExecutionContext(c, logger, uuid.New().String())
	if err != nil {
		logger.WrappedError(err, "failed to create execution context. this should never happen here as we are creating a new transaction id ourselves")
		return nil, err.Unwrap()
	}

	// TODO: validate the request parameters

	err = h.service.Apply(*ctx, req)
	if err != nil {
		ctx.Logger.WrappedError(err, "failed to apply command")
		return nil, err.Unwrap()
	}

	return &pb.ApplyCommandResponse{
		Id:            req.GetAggregateId(),
		TransactionId: ctx.GetTransactionID(),
	}, nil
}
