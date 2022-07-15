package service

import (
	"errors"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
	ct "github.com/circadence-official/galactus/pkg/chassis/context"
	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	Register(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*pb.RegisterResponse, l.Error)
	Connection(ctx ct.ExecutionContext, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error)
}

type service struct {
	env       string
	db        *gorm.DB
	mb        messagebus.MessageBus
	isDevMode bool
}

func NewService(env string, db *gorm.DB, mb messagebus.MessageBus, isDevMode bool) Service {
	return &service{
		env:       env,
		db:        db,
		mb:        mb,
		isDevMode: isDevMode,
	}
}

func (s *service) Register(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*pb.RegisterResponse, l.Error) {
	logger := ctx.GetLogger()

	// check for existing registration before creating new one
	registrationORM := &agpb.RegistrationORM{}
	r := s.db.Where("name = ? AND version = ?", req.GetName(), req.GetVersion()).
		Preload("Protocols").
		Preload("Producers").
		Preload("Consumers").
		First(registrationORM)
	// failed to check for existing registration
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {
		return nil, logger.WrapError(l.NewError(r.Error, "failed to check for existing registration"))
	}
	// no previous registration found, create new one
	if r.Error == gorm.ErrRecordNotFound {
		logger.Info("no previous registration found, creating new one")
		var err l.Error
		registrationORM, err = s.createDatabaseEntry(ctx, req)
		if err != nil {
			return nil, logger.WrapError(err)
		}
	}

	// process producers to register exchanges with the messagebus
	for _, p := range registrationORM.Producers {
		// TODO: just using topic for now. may need to change to direct?
		s.mb.RegisterExchange(ctx.GetContext(), s.generateExchangeName(p.Exchange), messagebus.ExchangeKindTopic)
	}

	// process consumers to register queues (and their associated exchanges) with the messagebus and bind them to their associated exchanges
	for i, c := range registrationORM.Consumers {
		var exchangeName, queueName string
		switch c.Kind {
		case int32(agpb.ConsumerKind_CONSUMER_KIND_QUEUE):
			exchangeName, queueName = s.generateExchangeAndQueueNames(registrationORM.Name, c)
			s.mb.RegisterExchange(ctx.GetContext(), exchangeName, messagebus.ExchangeKindTopic)
			s.mb.RegisterQueue(ctx.GetContext(), queueName, exchangeName, c.RoutingKey)
			registrationORM.Consumers[i].Queue = queueName
		case int32(agpb.ConsumerKind_CONSUMER_KIND_TOPIC):
			// topics need unique queue names, so generate a uuid to append
			c.Queue = uuid.NewString()
			exchangeName, queueName = s.generateExchangeAndQueueNames(registrationORM.Name, c)
			s.mb.RegisterExchange(ctx.GetContext(), exchangeName, messagebus.ExchangeKindTopic)
			s.mb.RegisterTopic(ctx.GetContext(), queueName, exchangeName, c.RoutingKey)
		default:
			return nil, logger.WithField("kind", c.Kind).WrapError(errors.New("unsupported consumer kind"))
		}

		// save exchange and queue names so they can be returned to caller
		registrationORM.Consumers[i].Queue = queueName
		registrationORM.Consumers[i].Exchange = exchangeName
	}

	// convert entry to proto
	registrationPB, err := registrationORM.ToPB(ctx.GetContext())
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to convert registration orm to proto"))
	}

	// TODO: generate events
	// evt := espb.EventType{Code: &espb.EventType_LabEventCode{LabEventCode: espb.LabEventCode_LAB_EVENT_CODE_LAB_CREATED}}
	// if err := et.CreateAndSendEventWithTransactionID(ctx.GetContext(), logger, s.eventStoreClient, lab, lab.GetId(), transactionId, espb.AggregateType_LAB, evt); err != nil {
	// 	logger.WithError(err).Error(ErrorFailedToEmitEvent)
	// 	return nil, err
	// }

	logger.Info("registered service")

	return &pb.RegisterResponse{
		Registration: &registrationPB,
	}, nil
}

func (s *service) createDatabaseEntry(ctx ct.ExecutionContext, req *pb.RegisterRequest) (*agpb.RegistrationORM, l.Error) {
	// if remote: address is service name (proxied through Istio)
	// if local (devMode): address is localhost
	serviceAddress := req.GetName()
	if s.isDevMode {
		serviceAddress = "localhost"
	}

	// generate ORM protocols from PB protocols
	protocolsORM := make([]*agpb.ProtocolORM, len(req.GetProtocols()))
	for i, protocolPB := range req.GetProtocols() {
		protocolORM, err := s.convertProtocolRequestToORM(ctx.GetLogger(), protocolPB, req.GetVersion())
		if err != nil {
			return nil, ctx.GetLogger().WrapError(err)
		}
		protocolsORM[i] = protocolORM
	}

	// generate ORM producers from PB producers
	producersORM := make([]*agpb.ProducerORM, len(req.GetProducers()))
	for i, producerPB := range req.GetProducers() {
		producerORM, err := producerPB.ToORM(ctx.GetContext())
		if err != nil {
			return nil, ctx.GetLogger().WrapError(l.NewError(err, "failed to convert producer pb to orm"))
		}
		producerORM.Id = uuid.NewString()
		producersORM[i] = &producerORM
	}

	// generate ORM consumers from PB consumers
	consumersORM := make([]*agpb.ConsumerORM, len(req.GetConsumers()))
	for i, consumerPB := range req.GetConsumers() {
		consumerORM, err := consumerPB.ToORM(ctx.GetContext())
		if err != nil {
			return nil, ctx.GetLogger().WrapError(l.NewError(err, "failed to convert consumer pb to orm"))
		}
		consumerORM.Id = uuid.NewString()
		consumersORM[i] = &consumerORM
	}

	// create new entry
	registrationORM := &agpb.RegistrationORM{
		Id:          uuid.NewString(),
		Name:        req.GetName(),
		Version:     req.GetVersion(),
		Description: req.GetDescription(),
		Address:     serviceAddress,
		Status:      int32(agpb.ServiceStatus_SERVICE_STATUS_REGISTERED),
		Protocols:   protocolsORM,
		Producers:   producersORM,
		Consumers:   consumersORM,
	}
	err := s.db.Create(&registrationORM).Error
	if err != nil {
		return nil, ctx.GetLogger().WrapError(l.NewError(err, "failed to create registration in db"))
	}

	return registrationORM, nil
}

func (s *service) Connection(ctx ct.ExecutionContext, req *pb.ConnectionRequest) (*pb.ConnectionResponse, l.Error) {
	logger := ctx.GetLogger().WithFields(l.Fields{
		"service_name": req.GetName(),
		"service_version": req.GetVersion(),
		"protocol_kind": req.GetType(),
	})

	result := &agpb.RegistrationORM{}
	err := s.db.Model(&agpb.RegistrationORM{}).Where("name = ? AND version = ?", req.GetName(), req.GetVersion()).Preload("Protocols").Find(result).Error
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to find registration"))
	}

	var port int32
	for _, protocolORM := range result.Protocols {
		if protocolORM.Kind == int32(req.GetType()) {
			port = protocolORM.Port
		}
	}
	if port == 0 {
		return nil, logger.WrapError(errors.New("failed to find port for given protocol kind"))
	}

	// TODO: check the health of the service

	return &pb.ConnectionResponse{
		Address: result.Address,
		Port:    port,
	}, nil
}
