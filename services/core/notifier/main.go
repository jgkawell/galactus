package main

import (
	"fmt"
	h "notifier/handler"
	s "notifier/service"

	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/notifier/v1"
	evpb "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"

	"github.com/circadence-official/galactus/pkg/chassis"
	cf "github.com/circadence-official/galactus/pkg/chassis/clientfactory"
	mb "github.com/circadence-official/galactus/pkg/chassis/messagebus"

	"google.golang.org/grpc"
)

func main() {
	var esc espb.EventStoreClient
	var svc s.Service

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "notifier",
		KeyVaultConfig: &chassis.KeyVaultConfig{
			RequireKeyVault:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		MessageBusConfig: &chassis.MessageBusConfig{},
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
				svc = s.NewService(b.GetConfig().GetBool("isHeartbeatEnabled"), b.GetConfig().GetInt("heartbeatTimer"), esc)
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) {
				pb.RegisterNotifierServer(b.GetRpcServer(), h.NewNotifierHandler(b.GetLogger(), svc))
			},
			CreateBrokerConfig: func(b chassis.MainBuilder) chassis.ConsumerConfig {
				return chassis.ConsumerConfig{
					Configs: []chassis.HandlerConfig{
						{
							AggregateType: fmt.Sprint(evpb.AggregateType_AGGREGATE_TYPE_NOTIFICATION),
							EventType:     fmt.Sprint(evpb.NotificationEventCode_NOTIFICATION_DELIVERY_REQUESTED),
							// we want to "multicast" each message to all replicas of this service
							ConsumerKind: mb.ExchangeKindTopic,
							Handler:      h.NewConsumer(b.GetLogger(), svc),
						},
					},
				}
			},
		},
	})

	defer b.Close()
	b.Run()
}
