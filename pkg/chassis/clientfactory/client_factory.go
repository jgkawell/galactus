package clientfactory

import (
	"context"
	"errors"
	"net/url"
	"time"

	espb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	l "github.com/jgkawell/galactus/pkg/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultTimeout = time.Second * 15

// ClientFactory is a quick way to create connections to internal RPC clients
type ClientFactory interface {
	// CloseConnection will close the specified gRPC connection
	CloseConnection(logger l.Logger, connection *grpc.ClientConn)
	// CreatEventerClient creates a client connection to the EventStore RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreatEventerClient(logger l.Logger, url string) (espb.EventerClient, *grpc.ClientConn, error)
	// CreateRegistryClient creates a client connection to the Registry RPC interface.
	// `defer CloseConnection(logger, conn)` will need to be called by the caller of this method.
	CreateRegistryClient(logger l.Logger, url string) (rgpb.RegistryClient, *grpc.ClientConn, error)
	// TODO
	Create(logger l.Logger, desc grpc.ServiceDesc) *grpc.ClientConn
}

type clientFactory struct{}

func NewClientFactory() ClientFactory {
	return &clientFactory{}
}

func (f *clientFactory) Create(logger l.Logger, desc grpc.ServiceDesc) *grpc.ClientConn {
	// TODO: look up connection in registry
	target := "todo"

	conn, _ := f.createRpcConnection(logger, target)
	registry := rgpb.NewRegistryClient(conn)

	request := &rgpb.ConnectionRequest{
		Route: desc.ServiceName,
	}
	_, _ = registry.Connection(context.Background(), request)

	conn, err := f.createRpcConnection(logger, target)
	if err != nil {
		logger.WithField("service_name", desc.ServiceName).WithError(err).Fatal("failed to create rpc connection")
	}

	return conn
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

func (f *clientFactory) CreatEventerClient(logger l.Logger, target string) (espb.EventerClient, *grpc.ClientConn, error) {
	connection, err := f.createRpcConnection(logger, target)
	if err != nil {
		return nil, nil, err
	}
	return espb.NewEventerClient(connection), connection, nil
}

// func (f *clientFactory) CreateCommandHandlerClient(logger l.Logger, target string) (chpb.CommandHandlerClient, *grpc.ClientConn, error) {
// 	connection, err := f.createRpcConnection(logger, target)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return chpb.NewCommandHandlerClient(connection), connection, nil
// }

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

	// Get hostname for tracing
	u, stdErr := url.Parse(target)
	if stdErr != nil {
		return nil, logger.WrapError(l.NewError(stdErr, "failed to parse target as url"))
	}

	// Create connection
	connection, stdErr := grpc.Dial(
		u.Host,
		// NOTE: We use insecure here since TLS is handled by the Istio sidecar
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: defaultTimeout}),
	)
	if stdErr != nil {
		return nil, logger.WrapError(l.NewError(stdErr, "failed while trying to dial target"))
	}

	return connection, nil
}
