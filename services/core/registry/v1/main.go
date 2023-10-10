package main

import (
	h "registry/handler"
	s "registry/service"

	"github.com/jgkawell/galactus/pkg/chassis"
	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/databases/postgres/gorm"
	"github.com/jgkawell/galactus/pkg/vault"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
)

func main() {
	var (
		svc s.Service
	)
	gormdb := gorm.New("")


	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName:      "registry",
		DoNotRegisterService: true,
		SecretsConfig: &chassis.SecretsConfig{
			Client:   vault.New(),
			Required: func(b chassis.MainBuilder) bool { return !b.IsDevMode() },
		},
		DatabaseConfig: &chassis.DatabaseConfig{
			Databases: []database.Client{
				gormdb,
			},
		},
		DaoLayerConfig: &chassis.DaoLayerConfig{
			CreateDaoLayer: func(b chassis.MainBuilder) {
				gormdb.Client().AutoMigrate(
					&agpb.RegistrationORM{},
					&agpb.ServerORM{},
					&agpb.ConsumerORM{},
				)
			},
		},
		ServiceLayerConfig: &chassis.ServiceLayerConfig{
			CreateServiceLayer: func(b chassis.MainBuilder) {
				svc = s.NewService(b.GetLogger(), gormdb.Client(), b.IsDevMode())
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			HttpPortConfigKey: "httpPort",
			GrpcPortConfigKey: "grpcPort",
			CreateRpcHandlers: func(b chassis.MainBuilder) []chassis.GrpcHandlers {
				return []chassis.GrpcHandlers{
					{
						Desc:    pb.Registry_ServiceDesc,
						Handler: h.NewRegistryHandler(b.GetLogger(), svc),
					},
				}
			},
		},
	})
	defer b.Close()
	b.Run()
}
