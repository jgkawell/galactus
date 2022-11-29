package main

import (
	"context"
	"fmt"

	h "commandhandler/handler"
	s "commandhandler/service"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	"github.com/jgkawell/galactus/pkg/chassis"
	cf "github.com/jgkawell/galactus/pkg/chassis/clientfactory"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"

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
			// NOTE: most services do not need to create a dedicated eventstore client like below, instead they should use the EventManager in the chassis
			//       commandhandler is an exception though as it doesn't create events by types (Aggregate/Event) but creates them by strings so it can't use EventManager.
			CreateInternalClients: func(b chassis.MainBuilder, clientFactory cf.ClientFactory) []*grpc.ClientConn {
				var (
					conn        *grpc.ClientConn
					err         error
					connections []*grpc.ClientConn
				)

				ctx, _ := ct.NewExecutionContext(context.Background(), b.GetLogger(), uuid.NewString())
				connectionResponse, err := b.GetRegistryClient().Connection(ctx.GetContext(), &rgpb.ConnectionRequest{
					Name:    "eventstore",
					Version: "0.0.0",
					Type:    agpb.ProtocolKind_PROTOCOL_KIND_GRPC,
				})

				if err != nil {
					b.GetLogger().WithError(err).Fatal("failed to get eventstore connection")
				}

				fullAddress := fmt.Sprintf("%s:%d", connectionResponse.GetAddress(), connectionResponse.GetPort())
				esc, conn, err = clientFactory.CreateEventStoreClient(b.GetLogger(), fullAddress)
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
