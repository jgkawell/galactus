package main

import (
	"eventer/controller"
	"eventer/handler"

	"github.com/jgkawell/galactus/pkg/chassis"
	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	"github.com/jgkawell/galactus/pkg/databases/postgres/gorm"
	"github.com/jgkawell/galactus/pkg/messaging/nats"
	"github.com/jgkawell/galactus/pkg/secrets/vault"

	pb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
)

func main() {
	var c controller.Eventer
	db := gorm.New("")
	bus := nats.New("")

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "eventer",
		SecretsConfig: &chassis.SecretsConfig{
			Client:   vault.New(),
			Required: func(b chassis.MainBuilder) bool { return !b.IsDevMode() },
		},
		MessageBusConfig: &chassis.MessageBusConfig{
			Buses: []messagebus.Client{
				bus,
			},
		},
		DatabaseConfig: &chassis.DatabaseConfig{
			Databases: []database.Client{
				db,
			},
		},
		DaoLayerConfig: &chassis.DaoLayerConfig{
			CreateDaoLayer: func(b chassis.MainBuilder) {
				db.Client().AutoMigrate(
					&pb.EventORM{},
				)
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				c = controller.New(db.Client(), bus, b.GetConfig().GetString("env"))
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateRpcHandlers: func(b chassis.MainBuilder) []chassis.GrpcHandlers {
				return []chassis.GrpcHandlers{
					{
						Desc:    pb.Eventer_ServiceDesc,
						Handler: handler.New(b.GetTelemetry(), c),
					},
				}
			},
		},
	})
	defer b.Close()
	b.Run()
}
