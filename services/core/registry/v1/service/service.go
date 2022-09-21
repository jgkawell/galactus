package service

import (
	"context"
	"errors"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	// check for existing registration before creating new one
	registrationORM := &agpb.RegistrationORM{}
	r := s.db.Where("name = ? AND version = ?", req.GetName(), req.GetVersion()).
		Preload("Protocols").
		Preload("Consumers").
		First(registrationORM)
	// failed to check for existing registration
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {
		return nil, ctx.Logger.WrapError(l.NewError(r.Error, "failed to check for existing registration"))
	}
	// no previous registration found, create new one
	if r.Error == gorm.ErrRecordNotFound {
		ctx.Logger.Info("no previous registration found, creating new one")
		var err l.Error
		registrationORM, err = s.createDatabaseEntry(ctx, req)
		if err != nil {
			return nil, ctx.Logger.WrapError(err)
		}
	}

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

func (s *service) createDatabaseEntry(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*agpb.RegistrationORM, l.Error) {

	registrationId := uuid.NewString()

	// if remote: address is service name (proxied through Istio)
	// if local (devMode): address is localhost
	serviceAddress := req.GetName()
	if s.isDevMode {
		serviceAddress = "localhost"
	}

	// generate ORM protocols from PB protocols
	protocolsORM := make([]*agpb.ProtocolORM, len(req.GetProtocols()))
	for i, protocolPB := range req.GetProtocols() {
		protocolORM, err := s.convertProtocolRequestToORM(ctx.Logger, protocolPB, req.GetVersion())
		if err != nil {
			return nil, ctx.Logger.WrapError(err)
		}
		protocolsORM[i] = protocolORM
	}

	// generate ORM consumers from PB consumers
	consumersORM := make([]*agpb.ConsumerORM, len(req.GetConsumers()))
	for i, consumerPB := range req.GetConsumers() {
		consumerORM := agpb.ConsumerORM{
			Id:             uuid.NewString(),
			Kind:           int32(consumerPB.GetKind()),
			RegistrationId: &registrationId,
			RoutingKey:     generateRoutingKey(consumerPB.GetAggregateType(), consumerPB.GetEventType(), consumerPB.GetEventCode()),
		}
		consumersORM[i] = &consumerORM
	}

	// create new entry
	registrationORM := &agpb.RegistrationORM{
		Id:          registrationId,
		Name:        req.GetName(),
		Version:     req.GetVersion(),
		Description: req.GetDescription(),
		Address:     serviceAddress,
		Status:      int32(agpb.ServiceStatus_SERVICE_STATUS_REGISTERED),
		Protocols:   protocolsORM,
		Consumers:   consumersORM,
	}
	err := s.db.Create(&registrationORM).Error
	if err != nil {
		return nil, ctx.Logger.WrapError(l.NewError(err, "failed to create registration in db"))
	}

	return registrationORM, nil
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
