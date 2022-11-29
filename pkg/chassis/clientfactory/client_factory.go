package clientfactory

import (
	"errors"
	"strings"
	"time"

	chpb "github.com/jgkawell/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

const defaultTimeout = time.Second * 15

// ClientFactory is a quick way to create connections to internal RPC clients
type ClientFactory interface {
	// CloseConnection will close the specified gRPC connection
	CloseConnection(logger l.Logger, connection *grpc.ClientConn)
	// CreateEventStoreClient creates a client connection to the EventStore RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateEventStoreClient(logger l.Logger, url string) (espb.EventStoreClient, *grpc.ClientConn, error)
	// CreateCommandHandlerClient creates a client connection to the CommandHandler RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateCommandHandlerClient(logger l.Logger, url string) (chpb.CommandHandlerClient, *grpc.ClientConn, error)
	// CreateRegistryClient creates a client connection to the Registry RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateRegistryClient(logger l.Logger, url string) (rgpb.RegistryClient, *grpc.ClientConn, error)
}

type clientFactory struct{}

func NewClientFactory(logger l.Logger) ClientFactory {
	return &clientFactory{}
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

func (f *clientFactory) CloseConnection(logger l.Logger, connection *grpc.ClientConn) {
	CloseConnection(logger, connection)
}

func (f *clientFactory) CreateEventStoreClient(logger l.Logger, target string) (espb.EventStoreClient, *grpc.ClientConn, error) {
	connection, err := f.createRpcConnection(logger, target)
	if err != nil {
		return nil, nil, err
	}
	return espb.NewEventStoreClient(connection), connection, nil
}

func (f *clientFactory) CreateCommandHandlerClient(logger l.Logger, target string) (chpb.CommandHandlerClient, *grpc.ClientConn, error) {
	connection, err := f.createRpcConnection(logger, target)
	if err != nil {
		return nil, nil, err
	}
	return chpb.NewCommandHandlerClient(connection), connection, nil
}

func (f *clientFactory) CreateRegistryClient(logger l.Logger, target string) (rgpb.RegistryClient, *grpc.ClientConn, error) {
	connection, err := f.createRpcConnection(logger, target)
	if err != nil {
		return nil, nil, err
	}
	return rgpb.NewRegistryClient(connection), connection, nil
}

func (f *clientFactory) createRpcConnection(logger l.Logger, target string) (*grpc.ClientConn, l.Error) {
	logger = logger.WithField("target", target)
	logger.Debug("creating a connection for a grpc client")

	// Check that target isn't empty. Dial doesn't error out but tries to connect indefinitely if target is empty
	if target == "" {
		return nil, logger.WrapError(errors.New("gRPC dial target cannot be empty"))
	}

	// Strip off port from target for trace service name
	s := strings.Split(target, ":")
	si := grpctrace.StreamClientInterceptor(grpctrace.WithServiceName(s[0]))
	ui := grpctrace.UnaryClientInterceptor(grpctrace.WithServiceName(s[0]))

	connection, stdErr := grpc.Dial(
		target,
		// NOTE: We use insecure here since TLS is handled by the Istio sidecar
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: defaultTimeout}),
		grpc.WithStreamInterceptor(si),
		grpc.WithUnaryInterceptor(ui))
	if stdErr != nil {
		return nil, logger.WrapError(l.NewError(stdErr, "failed while trying to dial target"))
	}

	return connection, nil
}
