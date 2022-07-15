package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"

	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/google/uuid"
)

const localPortConflictRetryLimit = 1000

// registry reserves 3500 (http) and 3501 (grpc) for itself when running locally
const localMinPort = 3502
const localMaxPort = 4500

func (s *service) convertProtocolRequestToORM(logger l.Logger, protocolPB *pb.Protocol, serviceVersion string) (*agpb.ProtocolORM, l.Error) {
	// check the protocol kind is valid
	if protocolPB.GetKind() == agpb.ProtocolKind_PROTOCOL_KIND_INVALID {
		return nil, logger.WithField("protocol_kind", protocolPB.GetKind()).WrapError(errors.New("invalid protocol kind"))
	}

	// generate the protocol port
	var port int32
	var err l.Error
	if s.isDevMode {
		port, err = s.generateLocalPort(logger)
		if err != nil {
			return nil, err
		}
	} else {
		port, err = s.generateRemotePort(logger, protocolPB.GetKind())
		if err != nil {
			return nil, err
		}
	}

	// get the protocol version (ex. if serviceVersion is "v2.3.5", then protocolVersion is "v2")
	version := strings.Split(serviceVersion, ".")[0]
	if version == "" {
		return nil, logger.WithField("service_version", serviceVersion).WrapError(errors.New("invalid service version"))
	}

	return &agpb.ProtocolORM{
		Id:      uuid.NewString(),
		Kind:    int32(protocolPB.GetKind()),
		Port:    port,
		Version: version,
	}, nil
}

/*
generateLocalPort will generate a random port between 3500 and 4500 making sure it is not already in use.

If an unused port is not found after 1000 attempts, an error is returned.
*/
func (s *service) generateLocalPort(logger l.Logger) (int32, l.Error) {
	for i := 0; i < localPortConflictRetryLimit; i++ {
		randomPort := rand.Intn(localMaxPort-localMinPort) + localMinPort
		var count int64
		err := s.db.Model(&agpb.ProtocolORM{}).Where("port = ?", randomPort).Count(&count).Error
		if err != nil {
			return 0, logger.WrapError(l.NewError(err, "failed to query for port usage while generating random local port"))
		}
		if count == 0 {
			return int32(randomPort), nil
		}
	}
	return 0, logger.WrapError(errors.New("failed to generate local port after maximum conflict retries"))
}

/*
generateRemotePort will generate the remote port based on the protocol kind:
	http = 8080
	grpc = 8090
If an invalid protocol kind is given, an error is returned.
*/
func (s *service) generateRemotePort(logger l.Logger, kind agpb.ProtocolKind) (int32, l.Error) {
	switch kind {
	case agpb.ProtocolKind_PROTOCOL_KIND_HTTP:
		return 8080, nil
	case agpb.ProtocolKind_PROTOCOL_KIND_GRPC:
		return 8090, nil
	default:
		return 0, logger.WrapError(errors.New("unsupported protocol kind"))
	}
}

/*
generateExchangeName will generate an exchange name based on the given base name and the service environment (k8s namespace or local)

The result will have the following form:
	exchangeName = "ENV.EXCHANGE_NAME"
*/
func (s *service) generateExchangeName(exchangeName string) string {
	return fmt.Sprintf("%s.%s", s.env, exchangeName)
}

/*
generateExchangeAndQueueNames will generate the exchange and queue names based on the given service name, service environment, and the
provided ConsumerORM values.

The result will have the following form:
    exchangeName = "ENV.EXCHANGE_NAME"
    // if consumer.Identifier is not empty
    queueName = "ENV.EXCHANGE_NAME.ROUTING_KEY.SERVICE_NAME.IDENTIFIER"
	// if consumer.Identifier is empty
	queueName = "ENV.EXCHANGE_NAME.ROUTING_KEY.SERVICE_NAME"
*/
func (s *service) generateExchangeAndQueueNames(serviceName string, consumer *agpb.ConsumerORM) (exchangeName, queueName string) {
	exchangeName = s.generateExchangeName(consumer.Exchange)
	queueName = fmt.Sprintf("%s.%s.%s.%s", exchangeName, consumer.RoutingKey, serviceName, consumer.Queue)
	// if the queue name is empty, remove trailing period
	if consumer.Queue == "" {
		queueName = strings.TrimSuffix(queueName, ".")
	}
	return exchangeName, queueName
}
