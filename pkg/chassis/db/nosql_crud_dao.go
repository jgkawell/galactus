package db

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CrudDaoConfig struct {
	SoftDelete           bool
	Timeout              time.Duration
	AllowUpsert          bool
	UniqueKeyColumnNames []string
}

type CrudDao interface {
	// InitializeCompositeUniqueIndex will set a unique key tupling constraint to the collection.
	InitializeCompositeUniqueIndex(logger l.Logger, keys bson.D) l.Error

	// Create inserts the specified record into the database.
	Create(ctx context.Context, logger l.Logger, model interface{}) (modelId string, err l.Error)
	// CreateWithId inserts the specified record into the database with the given id.
	CreateWithId(ctx context.Context, logger l.Logger, id string, model interface{}) (modelId string, err l.Error)

	// Read retrieves the records specified by the parameters.
	Read(ctx context.Context, logger l.Logger, params bson.M, decodeModelCallback func(bsonBytes []byte, modelId string) (interface{}, error), opts ...*options.FindOptions) ([]CrudDaoModel, l.Error)
	// ReadById retrieves the record specified by the id.
	ReadById(ctx context.Context, logger l.Logger, id string, decodeModelCallback func(bsonBytes []byte, modelId string) (interface{}, error), opts ...*options.FindOptions) (model CrudDaoModel, found bool, err l.Error)

	// Update changes records based off the parameters.
	Update(ctx context.Context, logger l.Logger, params bson.M, model interface{}) l.Error
	// Update changes the record specified by the id.
	UpdateById(ctx context.Context, logger l.Logger, id string, model interface{}) l.Error

	// Delete removes one or more records from the database.
	Delete(ctx context.Context, logger l.Logger, params bson.M) l.Error
	// DeleteById removes the record matching the id.
	DeleteById(ctx context.Context, logger l.Logger, id string) l.Error

	// GetCollection returns the mongo collection being used.
	// Useful for database transactions not covered in the other actions of this struct.
	GetCollection() *mongo.Collection

	// GenerateBase64URLEncodedId returns a base64url encoded random identifier or an error if it cannot generate one.
	GenerateBase64URLEncodedId(logger l.Logger) (string, l.Error)

	// GetConfiguration returns the configuration used by the dao.
	GetConfiguration() *CrudDaoConfig

	// SetRequiredParams is used to add required params for database calls.
	SetRequiredParams(params bson.M)
}

type crudDao struct {
	client         *mongo.Client
	dbName         string
	collectionName string
	config         *CrudDaoConfig
}

type CrudDaoModel interface {
	GetModel() interface{}
	GetHistory() *ModelHistory
}

type crudDaoModel struct {
	Id         string        `bson:"_id"`
	Model      interface{}   `bson:"model"`
	History    *ModelHistory `bson:"history"`
	CustomerId string        `bson:"customer_id,omitempty"` // deprecated
}

func (m *crudDaoModel) GetModel() interface{}     { return m.Model }
func (m *crudDaoModel) GetHistory() *ModelHistory { return m.History }

func NewCrudDao(logger l.Logger, client *mongo.Client, namespace string, dbName string, collectionName string, config *CrudDaoConfig) (CrudDao, l.Error) {
	if config == nil {
		return nil, logger.WrapError(errors.New("CrudDaoConfig must be set"))
	}
	if config.Timeout == 0 {
		logger.Warn("CrudDaoConfig Timeout not set. Defaulting to 10 seconds")
		config.Timeout = 10 * time.Second
	}

	if namespace != "" {
		collectionName = fmt.Sprintf("%s-%s", namespace, collectionName)
	}

	d := &crudDao{
		client:         client,
		dbName:         dbName,
		collectionName: collectionName,
		config:         config,
	}
	if len(config.UniqueKeyColumnNames) > 0 {
		if err := initializeNoSqlUniqueIndexes(logger, d.GetCollection(), d.dbName, d.collectionName, config.UniqueKeyColumnNames...); err != nil {
			return nil, logger.WrapError(l.NewError(err, "failed to apply unique key constraint to specified columns"))
		}
	}
	return d, nil
}

// NewCrudDaoAndClient creates a new NoSQL client and CRUD DAO
func NewCrudDaoAndClient(logger l.Logger, dbAddress string, namespace string, dbName string, collectionName string, config *CrudDaoConfig) (CrudDao, *mongo.Client, l.Error) {
	mongoClient, err := CreateNoSqlClient(logger, dbAddress)
	if err != nil {
		return nil, nil, logger.WrapError(l.NewError(err, "failed to create database client"))
	}
	logger.Info("Connected to " + dbAddress)
	dao, err := NewCrudDao(logger, mongoClient, namespace, dbName, collectionName, config)
	if err != nil {
		return nil, nil, logger.WrapError(l.NewError(err, "failed to create crud dao"))
	}
	return dao, mongoClient, nil
}

// InitializeCompositeUniqueIndex will set a unique key constraint to the collection
func (d *crudDao) InitializeCompositeUniqueIndex(logger l.Logger, keys bson.D) l.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := d.GetCollection().Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(true).SetName("compositeUniqueIndex"),
		}); err != nil {
		return logger.WrapError(l.NewError(err, "failed to create composite unique index"))
	}
	return nil
}

func (d *crudDao) Create(ctx context.Context, logger l.Logger, model interface{}) (string, l.Error) {
	return d.CreateWithId(ctx, logger, cuid.New(), model)
}

func (d *crudDao) CreateWithId(ctx context.Context, logger l.Logger, id string, model interface{}) (string, l.Error) {

	logger = logger.WithFields(l.Fields{
		"model": model,
	})
	logger.Debug("Create")

	crudModel := &crudDaoModel{
		Id:      id,
		Model:   model,
		History: &ModelHistory{Created: NewActionHistory(time.Now())},
	}

	// Insert the model
	ctx, cancel := context.WithTimeout(ctx, d.config.Timeout)
	defer cancel()
	_, err := d.GetCollection().InsertOne(ctx, crudModel)
	if err != nil {
		return "", logger.WrapError(l.NewError(err, "failed to insert model"))
	}

	return crudModel.Id, nil
}

// Query method for querying for Environments
func (d *crudDao) Read(ctx context.Context, logger l.Logger, params bson.M, decodeModelCallback func(bsonBytes []byte, id string) (interface{}, error), opts ...*options.FindOptions) ([]CrudDaoModel, l.Error) {

	logger = logger.WithFields(l.Fields{
		"params": params,
	})
	logger.Debug("Query")

	d.SetRequiredParams(params)

	ctx, cancel := context.WithTimeout(ctx, d.config.Timeout)
	defer cancel()
	cursor, err := d.GetCollection().Find(ctx, params, opts...)
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed query"))
	}
	defer func() { _ = cursor.Close(ctx) }()

	// iterate through results and decode them from found cursor
	var found []CrudDaoModel
	for cursor.Next(ctx) {

		var m crudDaoModel
		if err := cursor.Decode(&m); err != nil {
			return nil, logger.WrapError(l.NewError(err, "failed to decode model"))
		}

		bsonBytes, err := bson.Marshal(m.Model)
		if err != nil {
			return nil, logger.WithField("model", m.Model).WrapError(l.NewError(err, "failed to marshal the model"))
		}

		cm, err := decodeModelCallback(bsonBytes, m.Id)
		if err != nil {
			return nil, logger.WithField("model", m.Model).WrapError(l.NewError(err, "failed to decode model"))
		}

		m.Model = cm
		if cm != nil {
			found = append(found, &m)
		}
	}

	return found, nil
}

func (d *crudDao) ReadById(ctx context.Context, logger l.Logger, id string, decodeModelCallback func(bsonBytes []byte, id string) (interface{}, error), opts ...*options.FindOptions) (CrudDaoModel, bool, l.Error) {
	result, err := d.Read(ctx, logger, bson.M{"_id": id}, decodeModelCallback, opts...)
	if err != nil {
		return nil, false, err
	}
	if len(result) == 0 {
		return nil, false, nil
	}
	return result[0], true, nil
}

func (d *crudDao) Update(ctx context.Context, logger l.Logger, params bson.M, model interface{}) l.Error {

	logger = logger.WithFields(l.Fields{
		"model": model,
	})
	logger.Debug("Update")

	if len(params) == 0 {
		return logger.WrapError(errors.New("no parameters specified for update"))
	}
	d.SetRequiredParams(params)

	ctx, cancel := context.WithTimeout(context.Background(), d.config.Timeout)
	defer cancel()

	var result *mongo.UpdateResult
	var err error
	result, err = d.GetCollection().UpdateOne(
		ctx, params,
		bson.M{"$set": bson.M{
			"model":    model,
			UpdatedKey: NewActionHistory(time.Now()),
		}},
		&options.UpdateOptions{
			Upsert: func(b bool) *bool { return &b }(d.config.AllowUpsert),
		})
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to update model"))
	}
	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		return logger.WrapError(l.NewError(err, "failed to update model"))
	}

	return nil
}

func (d *crudDao) UpdateById(ctx context.Context, logger l.Logger, id string, model interface{}) l.Error {
	return d.Update(ctx, logger, bson.M{"_id": id}, model)
}

func (d *crudDao) Delete(ctx context.Context, logger l.Logger, params bson.M) l.Error {

	logger = logger.WithFields(l.Fields{
		"params": params,
	})
	logger.Debug("Delete")

	if len(params) == 0 {
		return logger.WrapError(errors.New("no parameters specified for delete"))
	}
	d.SetRequiredParams(params)

	ctx, cancel := context.WithTimeout(context.Background(), d.config.Timeout)
	defer cancel()
	var err error
	if d.config.SoftDelete {
		_, err = d.GetCollection().UpdateMany(ctx, params, bson.M{
			"$set": bson.M{
				DeletedKey: NewActionHistory(time.Now()),
			}})
	} else {
		_, err = d.GetCollection().DeleteMany(ctx, params, nil)
	}
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to delete model"))
	}

	return nil
}

func (d *crudDao) DeleteById(ctx context.Context, logger l.Logger, id string) l.Error {
	return d.Delete(ctx, logger, bson.M{"_id": id})
}

func (d *crudDao) GetCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection(d.collectionName)
}

func (d *crudDao) SetRequiredParams(params bson.M) {
	// Set essential params here to prevent outside users from setting them
	params[DeletedKey] = nil
	params["model"] = bson.M{"$ne": nil}
	// Can add role checks here (currently not applicable)
	// Probably should check roles through a policy engine instead (OPA?)
}

func (d *crudDao) GenerateBase64URLEncodedId(logger l.Logger) (string, l.Error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", logger.WrapError(l.NewError(err, "failed to generate random uuid"))
	}
	b, err := u.MarshalBinary()
	if err != nil {
		return "", logger.WrapError(l.NewError(err, "failed to marshal uuid to bytes"))
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (d *crudDao) GetConfiguration() *CrudDaoConfig {
	return d.config
}
