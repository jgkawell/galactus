package chassis

import (
	"fmt"
	"net"
	_ "net/http/pprof"

	"github.com/jgkawell/galactus/pkg/chassis/terminator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// createRpcServer creates the rpc server
func (b *mainBuilder) createRpcServer(opts ...grpc.ServerOption) {
	if opts == nil {
		opts = make([]grpc.ServerOption, 0)
	}
	b.rpcServer = grpc.NewServer(opts...)
}

func (b *mainBuilder) startRpcServer() {
	logger := b.logger.WithField("port", b.rpcPort)
	logger.Info("starting grpc server")

	if b.rpcPort == "" {
		b.logger.Fatal("grpc server failed to start. grpc port not set")
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", b.rpcPort))
	if err != nil {
		b.logger.WithError(err).Fatal("failed to create grpc listener")
	}
	reflection.Register(b.rpcServer)
	if err := b.rpcServer.Serve(grpcListener); err != nil {
		b.logger.WithError(err).Fatal("grpc server failed to start. see error field for details")
	}
	terminator.TerminateApplication()
}

func (b *mainBuilder) stopRpcServer() {
	if b.rpcServer != nil {
		b.rpcServer.Stop()
		b.logger.Info("grpc server stopped")
	}
}
