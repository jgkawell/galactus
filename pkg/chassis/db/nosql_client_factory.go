package db

import (
	"context"
	"strings"
	"time"

	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// CreateNoSqlClient will create a NoSQL client to a Mongo database using the official mongo drivers
func CreateNoSqlClient(logger l.Logger, dbAddress string) (*mongo.Client, l.Error) {
	logger.Info("creating nosql client")

	// add mongodb prefix if not already there
	if !strings.Contains(dbAddress, "://") {
		dbAddress = "mongodb://" + dbAddress
	}

	// create new client from address
	client, err := mongo.NewClient(options.Client().ApplyURI(dbAddress))
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to create database client"))
	}

	// connect to database with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to connect to database"))
	}

	return client, nil
}

// DisconnectNoSqlClient will disconnect a NoSQL client from a Mongo database
func DisconnectNoSqlClient(logger l.Logger, client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		logger.WithError(err).Error("failed to disconnect mongo client")
	}
}

// PingNoSqlClient is called in a separate thread that is responding to health checks
func PingNoSqlClient(ctx context.Context, client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// initializeNoSqlUniqueIndexes will create unique indexes on the database given a collection, name, and keys
func initializeNoSqlUniqueIndexes(logger l.Logger, collection *mongo.Collection, dbName string, collectionName string, uniqueKeys ...string) l.Error {
	logger = logger.WithFields(l.Fields{
		"db_name":         dbName,
		"collection_name": collectionName,
		"unique_keys":     uniqueKeys,
	})
	logger.Info("initializing unique indexes")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var keys bsonx.Doc
	for _, k := range uniqueKeys {
		keys = append(keys, bsonx.Elem{Key: k, Value: bsonx.Int32(1)})
	}

	if _, err := collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(true),
		}); err != nil {
		return logger.WrapError(l.NewError(err, "failed to assign unique index for current db and collection"))
	}
	return nil
}
