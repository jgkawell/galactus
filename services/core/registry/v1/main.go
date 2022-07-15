package main

import (
	h "registry/handler"
	s "registry/service"

	mb "github.com/circadence-official/galactus/pkg/chassis"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	pb "github.com/circadence-official/galactus/api/gen/go/core/registry/v1"
)

func main() {
	var svc s.Service

	b := mb.NewMainBuilder(&mb.MainBuilderConfig{
		ApplicationName: "registry",
		DoNotRegisterService: true,
		KeyVaultConfig: &mb.KeyVaultConfig{
			RequireKeyVault:               func(b mb.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		MessageBusConfig: &mb.MessageBusConfig{},
		SqlConfig: &mb.SqlConfig{
			SqlDbHost:   "sqlDbHost",
			SqlDbPort:   "sqlDbPort",
			SqlDbName:   "sqlDbName",
			SqlDbUser:   "sqlDbUser",
			SqlDbSecret: "sqlDbSecret",
			SqlDbSchema: "namespace",
		},
		DaoLayerConfig: &mb.DaoLayerConfig{
			CreateDaoLayer: func(b mb.MainBuilder) {
				db := b.GetSqlClient()
				db.AutoMigrate(
					&agpb.RegistrationORM{},
					&agpb.ProtocolORM{},
					&agpb.ProducerORM{},
					&agpb.ConsumerORM{},
				)
			},
		},
		ServiceLayerConfig: &mb.ServiceLayerConfig{
			CreateServiceLayer: func(b mb.MainBuilder) {
				svc = s.NewService(b.GetConfig().GetString("env"), b.GetSqlClient(), b.GetBroker(), b.IsDevMode())
			},
		},
		HandlerLayerConfig: &mb.HandlerLayerConfig{
			HttpPortVariable: "httpPort",
			RpcPortVariable:  "grpcPort",
			CreateRpcHandlers: func(b mb.MainBuilder) {
				pb.RegisterRegistryServer(b.GetRpcServer(), h.NewRegistryHandler(b.GetLogger(), svc))
			},
		},
	})

	defer b.Close()

	b.Run()
}
