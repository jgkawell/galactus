package clientfactory

import (
	"errors"
	"strings"
	"time"

	chpb "github.com/circadence-official/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	rgpb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"

	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

const timeout = time.Second * 15

// ClientFactory is a quick way to create connections to internal RPC clients
type ClientFactory interface {
	// CloseConnection will close the specified gRPC connection
	CloseConnection(connection *grpc.ClientConn)
	// CreateEventStoreClient creates a client connection to the EventStore RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateEventStoreClient(url string) (espb.EventStoreClient, *grpc.ClientConn, error)
	// CreateCommandHandlerClient creates a client connection to the CommandHandler RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateCommandHandlerClient(url string) (chpb.CommandHandlerClient, *grpc.ClientConn, error)
	// CreateRegistryClient creates a client connection to the Registry RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateRegistryClient(url string) (rgpb.RegistryClient, *grpc.ClientConn, error)
}

type clientFactory struct {
	logger l.Logger
}

func NewClientFactory(logger l.Logger) ClientFactory {
	return &clientFactory{
		logger: logger.WithField("struct", "ClientFactory"),
	}
}

// CloseConnection will close the specified gRPC connection
func CloseConnection(logger l.Logger, connection *grpc.ClientConn) {
	defer func() {
		if r := recover(); r != nil {
			logger.Warn("recovered from grpc connection failed to close")
		}
	}()
	if err := connection.Close(); err != nil {
		logger.WithError(err).Error("failed to close grpc connection")
	}
}

func (f *clientFactory) CloseConnection(connection *grpc.ClientConn) {
	CloseConnection(f.logger, connection)
}

func (f *clientFactory) CreateEventStoreClient(target string) (espb.EventStoreClient, *grpc.ClientConn, error) {
	if connection, err := f.createRpcConnection("EventStoreServiceClient", target); err != nil {
		return nil, nil, err
	} else {
		return espb.NewEventStoreClient(connection), connection, nil
	}
}

func (f *clientFactory) CreateCommandHandlerClient(target string) (chpb.CommandHandlerClient, *grpc.ClientConn, error) {
	if connection, err := f.createRpcConnection("CommandHandlerServiceClient", target); err != nil {
		return nil, nil, err
	} else {
		return chpb.NewCommandHandlerClient(connection), connection, nil
	}
}

func (f *clientFactory) CreateRegistryClient(target string) (rgpb.RegistryClient, *grpc.ClientConn, error) {
	if connection, err := f.createRpcConnection("RegistryServiceClient", target); err != nil {
		return nil, nil, err
	} else {
		return rgpb.NewRegistryClient(connection), connection, nil
	}
}

func (f *clientFactory) createRpcConnection(clientName string, target string) (*grpc.ClientConn, l.Error) {
	logger := f.logger.WithField("target", target)
	logger.Debug("Creating RPC Client for " + clientName)

	// Check that target isn't empty. Dial doesn't error out but tries to connect (to empty!) indefinitely
	if target == "" {
		err := errors.New("gRPC dial target cannot be empty")
		return nil, logger.WrapError(err)
	}

	// Strip off port from target for trace service name
	s := strings.Split(target, ":")
	si := grpctrace.StreamClientInterceptor(grpctrace.WithServiceName(s[0]))
	ui := grpctrace.UnaryClientInterceptor(grpctrace.WithServiceName(s[0]))

	if connection, err := grpc.Dial(
		target,
		// we use insecure here since TLS is handled by the Istio sidecar
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: timeout}),
		grpc.WithStreamInterceptor(si),
		grpc.WithUnaryInterceptor(ui)); err != nil {
		return nil, logger.WrapError(err)
	} else {
		return connection, nil
	}
}
