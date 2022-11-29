package main

import (
	h "eventstore/handler"
	s "eventstore/service"

	"github.com/jgkawell/galactus/pkg/chassis"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/eventstore/v1"
)

func main() {
	var svc s.Service

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "eventstore",
		KeyVaultConfig: &chassis.KeyVaultConfig{
			RequireKeyVault:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		MessageBusConfig: &chassis.MessageBusConfig{},
		SqlConfig: &chassis.SqlConfig{
			SqlDbHost:   "sqlDbHost",
			SqlDbPort:   "sqlDbPort",
			SqlDbName:   "sqlDbName",
			SqlDbUser:   "sqlDbUser",
			SqlDbSecret: "sqlDbSecret",
			SqlDbSchema: "namespace",
		},
		DaoLayerConfig: &chassis.DaoLayerConfig{
			CreateDaoLayer: func(b chassis.MainBuilder) {
				db := b.GetSqlClient()
				db.AutoMigrate(
					&agpb.EventORM{},
				)
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(b.GetSqlClient(), b.GetBroker(), b.GetConfig().GetString("env"))
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) {
				pb.RegisterEventStoreServer(b.GetRpcServer(), h.NewEventStoreHandler(b.GetLogger(), svc))
			},
		},
	})

	defer b.Close()

	b.Run()
}
