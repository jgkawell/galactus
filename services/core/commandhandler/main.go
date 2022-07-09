package main

import (
	h "commandhandler/handler"
	s "commandhandler/service"

	pb "github.com/circadence-official/galactus/api/gen/go/core/commandhandler/v1"
	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"

	"github.com/circadence-official/galactus/pkg/chassis"
	cf "github.com/circadence-official/galactus/pkg/chassis/clientfactory"

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

				esc, conn, err = clientFactory.CreateEventStoreClient(b.GetConfig().GetString("eventstoreaddress"))
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
