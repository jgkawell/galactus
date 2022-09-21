package handler

import (
	"context"

	s "registry/service"

	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
)

type registryHandler struct {
	logger  l.Logger
	service s.Service
}

// NewRegistryHandler
func NewRegistryHandler(logger l.Logger, service s.Service) pb.RegistryServer {
	return &registryHandler{logger, service}
}

// Register
func (h *registryHandler) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// build execution context, with request context that contains `transaction_id` in metadata
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, h.logger)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to create application context from current context")
		return nil, err.Unwrap()
	}

	response, err := h.service.Register(*ctx, req)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to register")
		return nil, err.Unwrap()
	}

	return response, err
}

// Connection
func (h *registryHandler) Connection(c context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	// build execution context, with request context that contains `transaction_id` in metadata
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, h.logger)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to create application context from current context")
		return nil, err
	}

	response, err := h.service.Connection(*ctx, req)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to get connection information")
		return nil, err.Unwrap()
	}

	return response, err
}

// HELPERS

func validateServiceVersion(version string) error {
	return nil
}
