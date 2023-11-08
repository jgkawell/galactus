package handler

import (
	"context"

	s "registry/service"

	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type registryHandler struct {
	telemetry ct.Telemetry
	service   s.Service
}

func NewRegistryHandler(telemetry ct.Telemetry, service s.Service) pb.RegistryServer {
	return &registryHandler{telemetry, service}
}

func (h *registryHandler) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ctx, span := ct.New(c, h.telemetry).Span()
	defer span.End()

	stdErr := req.ValidateAll()
	if stdErr != nil {
		ctx.Logger().WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	id, err := h.service.Register(ctx, req)
	if err != nil {
		ctx.Logger().WithFields(err.Fields()).WithError(err).Error("failed to register")
		return nil, err.Unwrap()
	}

	return &pb.RegisterResponse{
		Id: id,
	}, err
}

func (h *registryHandler) RegisterGrpcServer(c context.Context, req *pb.RegisterGrpcServerRequest) (*pb.RegisterGrpcServerResponse, error) {
	ctx, span := ct.New(c, h.telemetry).Span()
	defer span.End()

	stdErr := req.ValidateAll()
	if stdErr != nil {
		ctx.Logger().WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	port, err := h.service.RegisterGrpcServer(ctx, req)
	if err != nil {
		ctx.Logger().WithFields(err.Fields()).WithError(err).Error("failed to register grpc server")
		return nil, err.Unwrap()
	}

	return &pb.RegisterGrpcServerResponse{
		Port: port,
	}, err
}

func (h *registryHandler) RegisterHttpServer(c context.Context, req *pb.RegisterHttpServerRequest) (*pb.RegisterHttpServerResponse, error) {
	_, span := ct.New(c, h.telemetry).Span()
	defer span.End()
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (h *registryHandler) RegisterConsumers(c context.Context, req *pb.RegisterConsumersRequest) (*pb.RegisterConsumersResponse, error) {
	_, span := ct.New(c, h.telemetry).Span()
	defer span.End()
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

// Connection
func (h *registryHandler) Connection(c context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	ctx, span := ct.New(c, h.telemetry).Span()
	defer span.End()

	stdErr := req.ValidateAll()
	if stdErr != nil {
		ctx.Logger().WithError(stdErr).Warn("invalid request received")
		return nil, stdErr
	}

	response, err := h.service.Connection(ctx, req)
	if err != nil {
		ctx.Logger().WithFields(err.Fields()).WithError(err).Error("failed to get connection information")
		return nil, err.Unwrap()
	}

	return response, err
}
