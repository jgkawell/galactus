package handler

import (
	"context"

	s "registry/service"

	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type registryHandler struct {
	logger  l.Logger
	service s.Service
}

func NewRegistryHandler(logger l.Logger, service s.Service) pb.RegistryServer {
	return &registryHandler{logger, service}
}

func (h *registryHandler) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, h.logger)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to create application context from current context")
		return nil, err.Unwrap()
	}

	stdErr := req.ValidateAll()
	if stdErr != nil {
		h.logger.WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	id, err := h.service.Register(*ctx, req)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to register")
		return nil, err.Unwrap()
	}

	return &pb.RegisterResponse{
		Id: id,
	}, err
}

func (h *registryHandler) RegisterGrpcServer(c context.Context, req *pb.RegisterGrpcServerRequest) (*pb.RegisterGrpcServerResponse, error) {
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, h.logger)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to create application context from current context")
		return nil, err.Unwrap()
	}

	stdErr := req.ValidateAll()
	if stdErr != nil {
		h.logger.WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	port, err := h.service.RegisterGrpcServer(*ctx, req)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to register grpc server")
		return nil, err.Unwrap()
	}

	return &pb.RegisterGrpcServerResponse{
		Port: port,
	}, err
}

func (h *registryHandler) RegisterHttpServer(c context.Context, req *pb.RegisterHttpServerRequest) (*pb.RegisterHttpServerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (h *registryHandler) RegisterConsumers(c context.Context, req *pb.RegisterConsumersRequest) (*pb.RegisterConsumersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

// Connection
func (h *registryHandler) Connection(c context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	ctx, err := ct.NewExecutionContextFromContextWithMetadata(c, h.logger)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to create application context from current context")
		return nil, err
	}

	stdErr := req.ValidateAll()
	if stdErr != nil {
		h.logger.WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	response, err := h.service.Connection(*ctx, req)
	if err != nil {
		h.logger.WithFields(err.Fields()).WithError(err).Error("failed to get connection information")
		return nil, err.Unwrap()
	}

	return response, err
}
