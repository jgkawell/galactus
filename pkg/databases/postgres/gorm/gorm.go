package gorm

import (
	"context"
	"fmt"

	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/chassis/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Wrapper interface {
		database.Client
		Client() *gorm.DB
	}
	wrapper struct {
		client    *gorm.DB
		configKey string
	}
)

// New instantiates a new client wrapper. A call to Initialize is required before use.
// The configKey parameter dictates which key in the configuration file will be read during
// initialization. If configKey is empty, the default value of "postgres.url" will be used.
// The configuration can be in various formats, but the following is an example of a yaml file:
//
//	postgres:
//	  url: postgres://user:password@localhost:5432/dbname
func New(configKey string) Wrapper {
	if configKey == "" {
		configKey = "postgres.url"
	}
	return &wrapper{
		configKey: configKey,
	}
}

func (w *wrapper) Client() *gorm.DB {
	return w.client
}

func (w *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	url := config.GetString(w.configKey)
	client, err := gorm.Open(postgres.Open(url), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to sql db")
	}
	w.client = client
	return nil
}

func (w *wrapper) Disconnect(ctx context.Context) error {
	db, err := w.client.DB()
	if err != nil {
		return fmt.Errorf("failed to get the db connection for disconnect")
	}
	err = db.Close()
	if err != nil {
		return fmt.Errorf("failed to close the db connection for disconnect")
	}
	return nil
}

func (w *wrapper) Ping(ctx context.Context) error {
	db, err := w.client.DB()
	if err != nil {
		return fmt.Errorf("failed to get the db connection for ping")
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping sql db")
	}
	return nil
}
