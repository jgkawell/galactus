package main

import (
	"github.com/jgkawell/galactus/pkg/chassis"
	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/databases/postgres/gorm"
	"github.com/jgkawell/galactus/pkg/secrets/vault"
)

func main() {
	db := gorm.New("")
	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "secrets-example",
		SecretsConfig: &chassis.SecretsConfig{
			Client: vault.New(),
			Required: func(b chassis.MainBuilder) bool { return true },
		},
		DatabaseConfig: &chassis.DatabaseConfig{
			Databases: []database.Client{
				db,
			},
		},
	})
	defer b.Close()
	b.Run()
}
