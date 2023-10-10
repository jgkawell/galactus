package main

import (
	h "notifier/handler"
	s "notifier/service"

	pb "github.com/jgkawell/galactus/api/gen/go/core/notifier/v1"
	evpb "github.com/jgkawell/galactus/api/gen/go/generic/events/v1"

	"github.com/jgkawell/galactus/pkg/chassis"
	mb "github.com/jgkawell/galactus/pkg/chassis/messagebus"
)

func main() {
	var svc s.Service

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName:        "notifier",
		CreateEventStoreClient: true,
		VaultConfig: &chassis.VaultConfig{
			Required:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		MessageBusConfig: &chassis.MessageBusConfig{},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(b.GetConfig().GetBool("isHeartbeatEnabled"), b.GetConfig().GetInt("heartbeatTimer"), b.GetEventManager())
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) []chassis.GrpcHandlers {
				return []chassis.GrpcHandlers{
					{
						Desc:    pb.Notifier_ServiceDesc,
						Handler: h.NewNotifierHandler(b.GetLogger(), svc),
					},
				}
			},
			CreateBrokerConfig: func(b chassis.MainBuilder) chassis.ConsumerConfig {
				return chassis.ConsumerConfig{
					Configs: []chassis.HandlerConfig{
						{
							AggregateType: evpb.AggregateType_AGGREGATE_TYPE_NOTIFICATION,
							EventType:     &evpb.EventType{Code: &evpb.EventType_NotificationCode{}},
							EventCode:     evpb.NotificationEventCode_NOTIFICATION_EVENT_CODE_DELIVERY_REQUESTED,
							// we want to "multicast" each message to all replicas of this service
							ConsumerKind: mb.ExchangeKindTopic,
							Handler:      h.NewConsumer(b.GetLogger(), svc, b.GetEventManager()),
						},
					},
				}
			},
		},
	})

	defer b.Close()
	b.Run()
}
