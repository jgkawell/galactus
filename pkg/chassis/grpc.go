package chassis

import (
	"fmt"
	"net"
	_ "net/http/pprof"

	"github.com/circadence-official/galactus/pkg/chassis/terminator"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

// createRpcServer creates the rpc server
func (b *mainBuilder) createRpcServer(opts ...grpc.ServerOption) {
	if opts == nil {
		opts = make([]grpc.ServerOption, 0)
	}
	opts = append(
		opts,
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName(b.viper.GetString("traceName"))),
		)),
		grpc.StreamInterceptor(grpctrace.StreamServerInterceptor(grpctrace.WithServiceName(b.viper.GetString("traceName")))),
	)
	b.rpcServer = grpc.NewServer(opts...)
}

func (b *mainBuilder) StartRpcServer() {
	logger := b.logger.WithField("port", b.rpcPort)
	logger.Info("starting grpc server")

	if b.rpcPort == "" {
		b.logger.Panic("grpc server failed to start. grpc port not set")
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", b.rpcPort))
	if err != nil {
		b.logger.WithError(err).Panic("failed to create grpc listener")
	}
	if err := b.rpcServer.Serve(grpcListener); err != nil {
		b.logger.WithError(err).Panic("grpc server failed to start. see error field for details")
	}
	terminator.TerminateApplication()
}

func (b *mainBuilder) StopRpcServer() {
	if b.rpcServer != nil {
		b.rpcServer.Stop()
		b.logger.Info("grpc server stopped")
	}
}
