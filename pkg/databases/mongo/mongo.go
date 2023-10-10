package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/jgkawell/galactus/pkg/chassis/database"
	"github.com/jgkawell/galactus/pkg/chassis/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	Wrapper interface {
		database.Client
		Client() *mongo.Client
	}
	wrapper struct {
		client    *mongo.Client
		configKey string
	}
	Config struct {
		Host     string
		Port     string
		User     string
		Password string
	}
)

// New instantiates a new client wrapper. A call to Initialize is required before use.
// The configKey parameter dictates which key in the configuration file will be read during
// initialization. If configKey is empty, the default value of "mongo.url" will be used.
// The configuration can be in various formats, but the following is an example of a yaml file:
//
//	mongo:
//	  url: mongodb://user:password@localhost:27017
func New(configKey string) Wrapper {
	if configKey == "" {
		configKey = "mongo.url"
	}
	return &wrapper{
		configKey: configKey,
	}
}

func (w *wrapper) Client() *mongo.Client {
	return w.client
}

func (w *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	// create new client from address
	url := config.GetString(w.configKey)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return fmt.Errorf("failed to create mongo database client with error: %s", err.Error())
	}
	w.client = client

	// connect to database with timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo database with error: %s", err.Error())
	}

	// since client.Connect() does not verify the connection, ping the database before returning
	err = w.Ping(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (w *wrapper) Disconnect(ctx context.Context) error {
	if err := w.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect mongo client with error: %s", err.Error())
	}
	return nil
}

func (w *wrapper) Ping(ctx context.Context) error {
	if err := w.client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to ping database with error: %s", err.Error())
	}
	return nil
}
