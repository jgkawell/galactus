package main

import (
	h "eventstore/handler"
	s "eventstore/service"
	"time"

	"github.com/circadence-official/galactus/pkg/chassis"
	"github.com/circadence-official/galactus/pkg/chassis/db"

	es "github.com/circadence-official/galactus/api/gen/go/core/eventstore/v1"
)

func main() {
	var dao db.CrudDao
	var svc s.EventStoreService

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "eventstore",
		KeyVaultConfig: &chassis.KeyVaultConfig{
			RequireKeyVault:               func(b chassis.MainBuilder) bool { return !b.GetConfig().GetBool("isDevMode") },
			KeyVaultResourceGroupVariable: "resourceGroup",
			KeyVaultNameVariable:          "keyVault",
			KeyVaultOverridesVariable:     "keyVaultOverrides",
		},
		MessageBusConfig: &chassis.MessageBusConfig{},
		NoSqlConfig: &chassis.NoSqlConfig{
			DbAddressVariable: "dbAddress",
		},
		DaoLayerConfig: &chassis.DaoLayerConfig{
			CreateDaoLayer: func(b chassis.MainBuilder) {
				config := b.GetConfig()
				var err error
				dao, err = db.NewCrudDao(
					b.GetLogger(),
					b.GetMongoClient(),
					config.GetString("namespace"),
					config.GetString("dbName"),
					config.GetString("dbCollection"),
					&db.CrudDaoConfig{
						SoftDelete:           true,
						Timeout:              config.GetDuration("dbTimeout") * time.Second,
						AllowUpsert:          true,
						UniqueKeyColumnNames: []string{},
					},
				)
				if err != nil {
					b.GetLogger().WithError(err).Fatal("failed to create dao")
				}
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(dao, b.GetBroker())
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) {
				es.RegisterEventStoreServer(b.GetRpcServer(), h.NewEventStoreHandler(b.GetLogger(), svc))
			},
		},
	})

	defer b.Close()

	b.Run()
}
