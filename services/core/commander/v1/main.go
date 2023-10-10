package main

import (
	h "commander/handler"
	s "commander/service"

	pb "github.com/jgkawell/galactus/api/gen/go/core/commander/v1"
	espb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	"github.com/jgkawell/galactus/pkg/chassis"
	cf "github.com/jgkawell/galactus/pkg/chassis/clientfactory"

	"google.golang.org/grpc"
)

func main() {
	var svc s.Service
	var esc espb.EventStoreClient

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "commander",
		VaultConfig: &chassis.VaultConfig{
			Required:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		GatewayLayerConfig: &chassis.GatewayLayerConfig{
			// NOTE: most services do not need to create a dedicated eventer client like below, instead they should use the EventManager in the chassis
			//       commander is an exception though as it doesn't create events by types (Aggregate/Event) but creates them by strings so it can't use EventManager.
			CreateInternalClients: func(b chassis.MainBuilder, clientFactory cf.ClientFactory) []*grpc.ClientConn {
				var (
					conn        *grpc.ClientConn
					connections []*grpc.ClientConn
				)
				conn = clientFactory.Create(b.GetLogger(), rgpb.Registry_ServiceDesc)
				esc = espb.NewEventStoreClient(conn)
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
			CreateRpcHandlers: func(b chassis.MainBuilder) []chassis.GrpcHandlers {
				return []chassis.GrpcHandlers{
					{
						Desc:    pb.CommandHandler_ServiceDesc,
						Handler: h.NewHandler(b.GetLogger(), svc),
					},
				}
			},
		},
	})

	defer b.Close()
	b.Run()
}
