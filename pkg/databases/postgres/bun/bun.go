package bun

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/chassis/env"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	Wrapper interface {
		database.Client
		Client() *bun.DB
	}
	wrapper struct {
		client    *bun.DB
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

func (w *wrapper) Client() *bun.DB {
	return w.client
}

func (w *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	url := config.GetString(w.configKey)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	client := bun.NewDB(sqldb, pgdialect.New())
	w.client = client
	return nil
}

func (w *wrapper) Disconnect(ctx context.Context) error {
	err := w.client.Close()
	if err != nil {
		return fmt.Errorf("failed to close the db connection for disconnect")
	}
	return nil
}

func (w *wrapper) Ping(ctx context.Context) error {
	err := w.client.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping the db")
	}
	return nil
}
