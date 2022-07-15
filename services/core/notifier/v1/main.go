package main

import (
	"context"
	"fmt"
	h "notifier/handler"
	s "notifier/service"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	espb "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/notifier/v1"
	rgpb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
	evpb "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"
	"github.com/google/uuid"

	"github.com/circadence-official/galactus/pkg/chassis"
	cf "github.com/circadence-official/galactus/pkg/chassis/clientfactory"
	ct "github.com/circadence-official/galactus/pkg/chassis/context"
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
							AggregateType: fmt.Sprint(int64(evpb.AggregateType_AGGREGATE_TYPE_NOTIFICATION)),
							EventType:     fmt.Sprint(int64(evpb.NotificationEventCode_NOTIFICATION_DELIVERY_REQUESTED)),
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
