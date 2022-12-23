package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging/v2"
)

const localPortConflictRetryLimit = 1000

// registry reserves 3500 (http) and 3501 (grpc) for itself when running locally
const localMinPort = 3502
const localMaxPort = 4500

func (s *service) mergeRegistrations(ctx ct.ExecutionContext, request *pb.RegisterRequest, existing *agpb.RegistrationORM) (*agpb.RegistrationORM, l.Error) {

	// generate new id if empty
	if existing.Id == "" {
		existing.Id = uuid.NewString()
	}
	// generate address if empty
	if existing.Address == "" {
		existing.Address = request.Name
		if s.isDevMode {
			existing.Address = "localhost"
		}
	}
	// copy over requested metadata
	existing.Name = request.Name
	existing.Domain = request.Domain
	existing.Version = request.Version
	existing.Description = request.Description

	// add any new protocols
	for _, newP := range request.Protocols {
		// if already exists, break
		found := false
		for _, existingP := range existing.Protocols {
			if newP.Route == existingP.Route {
				found = true
				break
			}
		}
		// create new protocol and add to existing
		if !found {
			port, err := s.generatePort(ctx, agpb.ProtocolKind(newP.Kind))
			if err != nil {
				return nil, ctx.Logger.WrapError(err)
			}
			new := &agpb.ProtocolORM{
				Id:             uuid.NewString(),
				Kind:           int32(newP.Kind),
				Port:           port,
				RegistrationId: &existing.Id,
				Route:          newP.Route,
				Version:        existing.Version,
			}
			existing.Protocols = append(existing.Protocols, new)
		}
	}

	// add any new consumers
	for _, newC := range request.Consumers {
		routingKey := generateRoutingKey(newC.AggregateType, newC.EventType, newC.EventCode)
		// if already exists, break
		// TODO: what do we do if there are existing consumers that are no longer being used?
		found := false
		for _, existingC := range existing.Consumers {
			if routingKey == existingC.RoutingKey {
				found = true
				break
			}
		}
		// create new consumer and add to existing
		if !found {
			new := &agpb.ConsumerORM{
				Id:             uuid.NewString(),
				Kind:           int32(newC.Kind),
				RegistrationId: &existing.Id,
				RoutingKey:     routingKey,
			}
			existing.Consumers = append(existing.Consumers, new)
		}
	}

	return existing, nil
}

func (s *service) generatePort(ctx ct.ExecutionContext, kind agpb.ProtocolKind) (int32, l.Error) {
	var port int32
	var err l.Error
	if s.isDevMode {
		port, err = s.generateLocalPort(ctx.Logger)
		if err != nil {
			return 0, ctx.Logger.WrapError(err)
		}
	} else {
		port, err = s.generateRemotePort(ctx.Logger, kind)
		if err != nil {
			return 0, ctx.Logger.WrapError(err)
		}
	}
	return port, nil
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
func generateExchangeName(env, exchangeName string) string {
	return fmt.Sprintf("%s.%s", env, exchangeName)
}

func generateRoutingKey(aggregateType, eventType, eventCode string) string {
	return fmt.Sprintf("%s.%s.%s", aggregateType, eventType, eventCode)
}

/*
generateQueueName will generate queue name based on the given exchange, routingKey, service and identifier.

The result will have the following form:

	    exchangeName = "EXCHANGE_NAME"
	    // if identifier is not empty
	    queueName = "EXCHANGE_NAME.ROUTING_KEY.SERVICE_NAME.IDENTIFIER"
		// if identifier is empty
		queueName = "EXCHANGE_NAME.ROUTING_KEY.SERVICE_NAME"
*/
func (s *service) generateQueueName(exchangeName, routingKey, serviceName, identifier string) (queueName string) {
	queueName = fmt.Sprintf("%s.%s.%s.%s", exchangeName, routingKey, serviceName, identifier)
	// if the queue name is empty, remove trailing period
	if identifier == "" {
		queueName = strings.TrimSuffix(queueName, ".")
	}
	return queueName
}

// reduceVersion simplifies the semver to just the major version (e.g. if requested version is "v2.3.5", then queried version is "v2")
func reduceVersion(version string) string {
	version = strings.Split(version, ".")[0]
	if !strings.HasPrefix(version, "v") {
		return ""
	}
	return version
}
