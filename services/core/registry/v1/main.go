package main

import (
	h "registry/handler"
	s "registry/service"

	"github.com/circadence-official/galactus/pkg/chassis"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
)

func main() {
	var svc s.Service

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "registry",
		DoNotRegisterService: true,
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
					&agpb.RegistrationORM{},
					&agpb.ProtocolORM{},
					&agpb.ConsumerORM{},
				)
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(b.GetLogger(), b.GetConfig().GetString("env"), b.GetSqlClient(), b.GetBroker(), b.IsDevMode())
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			HttpPortVariable: "httpPort",
			RpcPortVariable:  "grpcPort",
			CreateRpcHandlers: func(b chassis.MainBuilder) {
				pb.RegisterRegistryServer(b.GetRpcServer(), h.NewRegistryHandler(b.GetLogger(), svc))
			},
		},
	})

	defer b.Close()

	b.Run()
}