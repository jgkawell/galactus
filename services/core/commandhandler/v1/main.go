package main

import (
	"context"
	"fmt"

	h "commandhandler/handler"
	s "commandhandler/service"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	rgpb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"

	"github.com/circadence-official/galactus/pkg/chassis"
	cf "github.com/circadence-official/galactus/pkg/chassis/clientfactory"
	ct "github.com/circadence-official/galactus/pkg/chassis/context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	var svc s.Service
	var esc espb.EventStoreClient

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "commandhandler",
		KeyVaultConfig: &chassis.KeyVaultConfig{
			RequireKeyVault:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		GatewayLayerConfig: &chassis.GatewayLayerConfig{
			CreateInternalClients: func(b chassis.MainBuilder, clientFactory cf.ClientFactory) []*grpc.ClientConn {
				var (
					conn        *grpc.ClientConn
					err         error
					connections []*grpc.ClientConn
				)

				ctx, _ := ct.NewExecutionContext(context.Background(), b.GetLogger(), uuid.NewString())
				connectionResponse, err := b.GetRegistryClient().Connection(ctx.GetContextWithTransactionID(), &rgpb.ConnectionRequest{
					Name:    "eventstore",
					Version: "0.0.0",
					Type:    agpb.ProtocolKind_PROTOCOL_KIND_GRPC,
				})

				if err != nil {
					b.GetLogger().WithError(err).Fatal("failed to get eventstore connection")
				}

				fullAddress := fmt.Sprintf("%s:%d", connectionResponse.GetAddress(), connectionResponse.GetPort())
				esc, conn, err = clientFactory.CreateEventStoreClient(fullAddress)
				if err != nil {
					b.GetLogger().WithError(err).Panic("failed to connect to eventstore service")
				}
				connections = append(connections, conn)

				return connections
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(esc)
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) {
				pb.RegisterCommandHandlerServer(b.GetRpcServer(), h.NewHandler(b.GetLogger(), svc))
			},
		},
	})

	defer b.Close()
	b.Run()
}
