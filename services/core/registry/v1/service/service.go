package service

import (
	"context"
	"errors"
	"strings"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	defaultExchangeName = "default"
)

type Service interface {
	Register(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*pb.RegisterResponse, l.Error)
	Connection(ctx ct.ExecutionContext, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error)
}

type service struct {
	defaultExchange string
	db              *gorm.DB
	mb              messagebus.MessageBus
	isDevMode       bool
}

func NewService(logger l.Logger, env string, db *gorm.DB, mb messagebus.MessageBus, isDevMode bool) Service {
	// attempt to register the default exchange
	exchange := generateExchangeName(env, defaultExchangeName)
	err := mb.RegisterExchange(context.Background(), exchange, messagebus.ExchangeKindTopic)
	if err != nil {
		logger.WithField("exchange_name", exchange).WithError(err).Fatal("failed to register default exchange with messagebus")
	}

	return &service{
		defaultExchange: exchange,
		db:              db,
		mb:              mb,
		isDevMode:       isDevMode,
	}
}

func (s *service) Register(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*pb.RegisterResponse, l.Error) {

	// get the version (e.g. if requested version is "v2.3.5", then queried version is "v2")
	version := strings.Split(req.Version, ".")[0]
	if version == "" {
		return nil, ctx.Logger.WithField("requested_version", req.Version).WrapError(errors.New("invalid requested service version. must have the form vX.Y.Z"))
	}

	// check for existing registration before creating new one
	registrationORM := &agpb.RegistrationORM{}
	r := s.db.Where("domain = ? AND name = ? AND version = ?", req.Domain, req.Name, req.Version).
		Preload("Protocols").
		Preload("Consumers").
		Find(registrationORM)

	// failed to check for existing registration
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {
		return nil, ctx.Logger.WrapError(l.NewError(r.Error, "failed to check for existing registration"))
	}

	// merge requested registration with existing data
	registrationORM, err := s.mergeRegistrations(ctx, req, registrationORM)
	if err != nil {
		return nil, ctx.Logger.WrapError(err)
	}

	// upsert into database
	stdErr := s.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(registrationORM).Error
	if stdErr != nil {
		return nil, ctx.Logger.WrapError(l.NewError(stdErr, "failed to create registration in db"))
	}

	// initialize response
	response := &pb.RegisterResponse{
		Protocols: make([]*pb.ProtocolResponse, len(req.GetProtocols())),
		Consumers: make([]*pb.ConsumerResponse, len(req.GetConsumers())),
	}

	// pull out protocol values from ORM
	for i, p := range registrationORM.Protocols {
		response.Protocols[i] = &pb.ProtocolResponse{
			Kind: agpb.ProtocolKind(p.Kind),
			Port: p.Port,
		}
	}

	// process consumers to register queues (and their associated exchanges) with the messagebus and bind them to their associated exchanges
	for i, c := range registrationORM.Consumers {
		var queueName string
		switch c.Kind {
		case int32(agpb.ConsumerKind_CONSUMER_KIND_QUEUE):
			queueName = s.generateQueueName(s.defaultExchange, c.RoutingKey, registrationORM.Name, "")
			err := s.mb.RegisterQueue(ctx.GetContext(), queueName, s.defaultExchange, c.RoutingKey)
			if err != nil {
				return nil, ctx.Logger.WrapError(err)
			}
		case int32(agpb.ConsumerKind_CONSUMER_KIND_TOPIC):
			// topics need unique queue names, so generate a uuid to append
			queueName = s.generateQueueName(s.defaultExchange, c.RoutingKey, registrationORM.Name, uuid.NewString())
			err := s.mb.RegisterTopic(ctx.GetContext(), queueName, s.defaultExchange, c.RoutingKey)
			if err != nil {
				return nil, ctx.Logger.WrapError(err)
			}
		default:
			return nil, ctx.Logger.WithField("kind", c.Kind).WrapError(errors.New("unsupported consumer kind"))
		}

		// save values for return to caller
		response.Consumers[i] = &pb.ConsumerResponse{
			Kind:       agpb.ConsumerKind(c.Kind),
			RoutingKey: c.RoutingKey,
			Exchange:   s.defaultExchange,
			QueueName:  queueName,
		}
	}

	// TODO: generate events
	// evt := espb.EventType{Code: &espb.EventType_LabEventCode{LabEventCode: espb.LabEventCode_LAB_EVENT_CODE_LAB_CREATED}}
	// if err := et.CreateAndSendEventWithTransactionID(ctx.GetContext(), logger, s.eventStoreClient, lab, lab.GetId(), transactionId, espb.AggregateType_LAB, evt); err != nil {
	// 	ctx.Logger.WithError(err).Error(ErrorFailedToEmitEvent)
	// 	return nil, err
	// }

	ctx.Logger.Info("registered service")

	return response, nil
}

func (s *service) Connection(ctx ct.ExecutionContext, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error) {
	ctx.Logger = ctx.Logger.WithFields(l.Fields{
		"service_name":    req.GetName(),
		"service_version": req.GetVersion(),
		"protocol_kind":   req.GetType(),
	})

	result := &agpb.RegistrationORM{}
	err := s.db.Model(&agpb.RegistrationORM{}).Where("name = ? AND version = ?", req.GetName(), req.GetVersion()).Preload("Protocols").Find(result).Error
	if err != nil {
		return nil, ctx.Logger.WrapError(l.NewError(err, "failed to find registration"))
	}

	var port int32
	for _, protocolORM := range result.Protocols {
		if protocolORM.Kind == int32(req.GetType()) {
			port = protocolORM.Port
		}
	}
	if port == 0 {
		return nil, ctx.Logger.WrapError(errors.New("failed to find port for given protocol kind"))
	}

	// TODO: check the health of the service

	return &pb.ConnectionResponse{
		Address: result.Address,
		Port:    port,
	}, nil
}
