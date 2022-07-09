package handler

import (
	"context"

	s "commandhandler/service"

	pb "github.com/circadence-official/galactus/api/gen/go/core/commandhandler/v1"

	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

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

const (
	ErrorApplyFailed                    = "apply failed"
	ErrorFailedToCreateExecutionContext = "failed to create execution context"
)

func (h *handler) Apply(ctx context.Context, req *pb.ApplyCommandRequest) (*pb.ApplyCommandResponse, error) {
	logger := h.logger.WithRPCContext(ctx)
	// Creates transactionId and adds to execution context
	executionCtx, err := ct.NewExecutionContext(ctx, logger, uuid.New().String())
	if err != nil {
		executionCtx.GetLogger().WithFields(err.Fields()).WithError(err).Error(ErrorFailedToCreateExecutionContext)
		return nil, err.Unwrap()
	}

	res, err := h.service.Apply(executionCtx, req)
	if err != nil {
		executionCtx.GetLogger().WithFields(err.Fields()).WithError(err).Error(ErrorApplyFailed)
		return nil, err.Unwrap()
	}

	return res, nil
}
