package service

import (
	"fmt"

	"github.com/google/uuid"
	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"

	"gorm.io/gorm"
)

type Service interface {
	Register(ctx ct.Context, req *pb.RegisterRequest) (string, l.Error)
	RegisterGrpcServer(ctx ct.Context, req *pb.RegisterGrpcServerRequest) (string, l.Error)
	Connection(ctx ct.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error)
}

type service struct {
	db        *gorm.DB
	isDevMode bool
}

func NewService(logger l.Logger, db *gorm.DB, isDevMode bool) Service {
	return &service{
		db:        db,
		isDevMode: isDevMode,
	}
}

func (s *service) Register(ctx ct.Context, req *pb.RegisterRequest) (string, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	// check for existing registration before creating new one
	registrationORM := &pb.RegistrationORM{}
	r := s.db.
		Where("domain = ? AND name = ? AND version = ?", req.Domain, req.Name, req.Version).
		Find(registrationORM)

	// failed to check for existing registration
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {
		return "", ctx.Logger().WrapError(l.NewError(r.Error, "failed to check for existing registration"))
	}

	if registrationORM.Id != "" {
		return registrationORM.Id, nil
	}

	// create new registration
	registrationORM = &pb.RegistrationORM{
		Id:      uuid.New().String(),
		Domain:  req.Domain,
		Name:    req.Name,
		Version: req.Version,
	}

	// insert into database
	stdErr := s.db.Create(registrationORM).Error
	if stdErr != nil {
		return "", ctx.Logger().WrapError(l.NewError(stdErr, "failed to create registration in db"))
	}

	ctx.Logger().Info("registered service")

	return registrationORM.Id, nil
}

func (s *service) RegisterGrpcServer(ctx ct.Context, req *pb.RegisterGrpcServerRequest) (string, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	// check for existing server before creating new one
	server := &pb.ServerORM{}
	err := s.db.
		Where("route = ?", req.Route).
		First(server).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", ctx.Logger().WrapError(l.NewError(err, "failed to check for existing server"))
	}

	if server.Id != "" {
		return server.Port, nil
	}

	// create new server

	// get registration
	registration := &pb.RegistrationORM{}
	err = s.db.
		Where("id = ?", req.Id).
		First(registration).Error
	if err != nil {
		return "", ctx.Logger().WrapError(l.NewError(err, "failed to find registration with matching id"))
	}

	// get port
	port, err := s.generatePort(ctx, pb.ServerKind_SERVER_KIND_GRPC)
	if err != nil {
		return "", ctx.Logger().WrapError(err)
	}

	// define scheme and host
	scheme := "http"
	host := "localhost"
	if !s.isDevMode {
		scheme = "https"
		host = fmt.Sprintf("%s.%s.svc.cluster.local", registration.Name, registration.Domain)
	}

	// save to database
	server = &pb.ServerORM{
		Id:     uuid.New().String(),
		Route:  req.Route,
		Scheme: scheme,
		Host:   host,
		Port:   port,
		Kind:   int32(pb.ServerKind_SERVER_KIND_GRPC.Number()),
	}
	err = s.db.Create(server).Error
	if err != nil {
		return "", ctx.Logger().WrapError(l.NewError(err, "failed to create server in db"))
	}

	ctx.Logger().Info("registered grpc server")

	return server.Port, nil
}

func (s *service) Connection(ctx ct.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	result := &pb.ServerORM{}
	err := s.db.
		Model(&pb.ServerORM{}).
		Where("route = ?", req.Route).
		First(result).Error
	if err != nil {
		return nil, ctx.Logger().WrapError(l.NewError(err, "failed to find route with matching path"))
	}

	address := fmt.Sprintf("%s://%s:%s", result.Scheme, result.Host, result.Port)
	return &pb.ConnectionResponse{
		Address: address,
		// TODO: check the health of the service
		Status: pb.Status_STATUS_HEALTHY,
	}, nil
}
