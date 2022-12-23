package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	l "github.com/jgkawell/galactus/pkg/logging/v2"

	spy "bou.ke/monkey"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(m *testing.M) {
	fmt.Println("This test must be run with the \"-gcflags=-l\" flag")
	os.Exit(m.Run())
}

func TestNewCrudDao(t *testing.T) {
	type testCase int
	const (
		failNilCrudDaoConfig testCase = iota
		failInitializeUniqueIndexes
		noCrudDaoConfigTimeoutSpecified
		noUniqueKeys
		hasUniqueKeys
	)
	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"nil crud dao config", failNilCrudDaoConfig},
		{"initializeUniqueIndexes fails", failInitializeUniqueIndexes},
		{"CrudDaoConfig timeout not set", noCrudDaoConfigTimeoutSpecified},
		{"no unique keys", noUniqueKeys},
		{"has unique keys", hasUniqueKeys},
	}

	expectedUniqueKeys := []string{"uk1", "uk2"}
	defer spy.UnpatchAll()

	for _, tc := range testCases {
		// Setup
		var config *CrudDaoConfig
		if tc.testCase != failNilCrudDaoConfig {
			config = &CrudDaoConfig{}
			if tc.testCase == failInitializeUniqueIndexes || tc.testCase == hasUniqueKeys {
				config.UniqueKeyColumnNames = expectedUniqueKeys
			}
			if tc.testCase != noCrudDaoConfigTimeoutSpecified {
				config.Timeout = 10
			}
		}

		// Patch initializeUniqueIndexes
		spy.Patch(initializeNoSqlUniqueIndexes,
			func(logger l.Logger, collection *mongo.Collection, dbName string, collectionName string, uniqueKeys ...string) l.Error {
				switch tc.testCase {
				case failInitializeUniqueIndexes:
					return logger.WrapError(errors.New(tc.testName))
				case hasUniqueKeys:
					assert.Equal(t, expectedUniqueKeys, uniqueKeys, tc.testName)
				default:
					if len(uniqueKeys) == 0 {
						assert.Fail(t, "initializeUniqueIndexes should not be called when no unique keys are specified")
					}
				}
				return nil
			})

		// Create a mock logger
		logger, hook := l.CreateNullLogger()

		// Create a mock mongo client
		mockMongoClient := createMockMongoClient(&mongo.Collection{})

		// Test
		result, err := NewCrudDao(logger, mockMongoClient, config)

		// Verify
		if tc.testCase < noCrudDaoConfigTimeoutSpecified {
			assert.Nil(t, result, tc.testName)
			assert.NotNil(t, err, tc.testName)

			// Check that error logs were created
			expectedMessage := ""
			if tc.testCase == failNilCrudDaoConfig {
				expectedMessage = "CrudDaoConfig must be set"
			} else if tc.testCase == failInitializeUniqueIndexes {
				expectedMessage = "failed to apply unique key constraint to specified columns"
			}
			AssertLastLogMessage(t, hook, tc.testName, "error", expectedMessage)
		} else {
			assert.NotNil(t, result, tc.testName)
			assert.Nil(t, err, tc.testName)

			// Check log messages
			if tc.testCase == noCrudDaoConfigTimeoutSpecified {
				AssertLastLogMessage(t, hook, tc.testName, "warning", "CrudDaoConfig Timeout not set. Defaulting to 10 seconds")
			} else {
				AssertNoWarnOrErrorLogMessages(t, hook, tc.testName)
			}
		}
	}
}

// TestCreate unit test for EnvironmentDao.Create
func TestCrudDao_Create(t *testing.T) {
	type testCase int
	const (
		failInsertOne testCase = iota
		success
	)
	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"fail InsertOne", failInsertOne},
		{"success", success},
	}

	// Test setup
	defer spy.UnpatchAll()
	expectedRecord := struct{}{}

	for _, tc := range testCases {
		// Setup
		var insertedRecord *crudDaoModel
		collection := &mongo.Collection{}

		// Mock collection.InsertOne
		spy.PatchInstanceMethod(
			reflect.TypeOf(collection), "InsertOne",
			func(c *mongo.Collection, ctx context.Context, record interface{}, options ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				// Verify model
				insertedRecord, _ = record.(*crudDaoModel)

				if tc.testCase == failInsertOne {
					return nil, errors.New("")
				}
				return nil, nil
			})

		// Mock the logger
		logger, hook := l.CreateNullLogger()

		// Create the CrudDao
		d := crudDao{
			client: createMockMongoClient(collection),
			config: &CrudDaoConfig{Timeout: 10},
		}

		// Test
		result, err := d.Create(context.Background(), logger, expectedRecord)

		// Verify
		if tc.testCase < success {
			assert.Equal(t, "", result, tc.testName)
			assert.NotNil(t, err, tc.testName)

			expectedMessage := ""
			if tc.testCase == failInsertOne {
				expectedMessage = "failed to insert model"
			}
			AssertLastLogMessage(t, hook, tc.testName, "error", expectedMessage)
		} else {
			assert.NotEqual(t, "", result, tc.testName)
			assert.Nil(t, err, tc.testName)

			AssertNoWarnOrErrorLogMessages(t, hook, tc.testName)

			if assert.NotNil(t, insertedRecord, tc.testName, "record is crudDaoModel") {
				assert.NotEqual(t, "", insertedRecord.Id, tc.testName, "record id should be set")
				assert.Equal(t, expectedRecord, insertedRecord.Model, tc.testName)
				if assert.NotNil(t, insertedRecord.History, tc.testName, "record's history should be set") {
					h := insertedRecord.History
					if assert.NotNil(t, h.Created, tc.testName, "record should have create history") {
						assert.False(t, h.Created.At.IsZero(), "record history should not be zero time", tc.testName)
					}
					assert.Nil(t, h.Updated, "update should not be set.", tc.testName)
					assert.Nil(t, h.Deleted, "delete should not be set.", tc.testName)
				}
			}
		}
	}
}

func TestCrudDao_Read(t *testing.T) {
	type testCase int
	const (
		failFind testCase = iota
		failCursorDecode
		failBsonMarshal
		failCallback
		noRecordsFound
		success
	)
	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"fail Find", failFind},
		{"fail Cursor.Decode", failCursorDecode},
		{"fail bson.Marshal", failBsonMarshal},
		{"fail Callback", failCallback},
		{"no records found", noRecordsFound},
		{"success", success},
	}

	// Test setup
	defer spy.UnpatchAll()

	for _, tc := range testCases {
		// Setup
		var findParams bson.M
		collection := &mongo.Collection{}
		cursor := &mongo.Cursor{}
		var expectedModels []testModel
		var expectedCrudRecords []CrudDaoModel
		index := -1

		if tc.testCase == failCursorDecode ||
			tc.testCase == failBsonMarshal ||
			tc.testCase == failCallback ||
			tc.testCase == success {
			expectedModels = []testModel{{s: "s1"}, {s: "s2"}}
			expectedCrudRecords = []CrudDaoModel{
				&crudDaoModel{Id: "1", Model: expectedModels[0]},
				&crudDaoModel{Id: "2", Model: expectedModels[1]},
			}
		}

		// Mock collection.InsertOne
		spy.PatchInstanceMethod(
			reflect.TypeOf(collection), "Find",
			func(c *mongo.Collection, ctx context.Context, filter interface{}, options ...*options.FindOptions) (*mongo.Cursor, error) {
				// Verify model
				findParams, _ = filter.(bson.M)
				if tc.testCase == failFind {
					return nil, errors.New("")
				}
				return cursor, nil
			})

		// Mock cursor.Next
		spy.PatchInstanceMethod(
			reflect.TypeOf(cursor), "Next",
			func(c *mongo.Cursor, ctx context.Context) bool {
				index++
				return index < len(expectedCrudRecords)
			})

		// Mock cursor.Decode
		spy.PatchInstanceMethod(
			reflect.TypeOf(cursor), "Decode",
			func(c *mongo.Cursor, val interface{}) error {
				if tc.testCase == failCursorDecode {
					return errors.New("")
				}
				m := val.(*crudDaoModel)
				expected := expectedCrudRecords[index].(*crudDaoModel)
				m.Id = expected.Id
				m.Model = expected.Model
				return nil
			})

		// Mock cursor.Close
		spy.PatchInstanceMethod(
			reflect.TypeOf(cursor), "Close",
			func(c *mongo.Cursor, ctx context.Context) error {
				return nil
			})

		// Mock bson.Marshal
		spy.Patch(
			bson.Marshal, func(val interface{}) ([]byte, error) {
				if tc.testCase == failBsonMarshal {
					return nil, errors.New("")
				}
				return []byte{}, nil
			})

		// Mock the logger
		logger, hook := l.CreateNullLogger()

		// Create the CrudDao
		d := crudDao{
			client: createMockMongoClient(collection),
			config: &CrudDaoConfig{Timeout: 10},
		}

		// Test
		result, err := d.Read(context.Background(), logger, bson.M{"p": "1"}, func(b []byte, id string) (interface{}, error) {
			if tc.testCase == failCallback {
				return nil, errors.New("")
			}
			return expectedModels[index], nil
		})

		// Verify
		if tc.testCase < noRecordsFound {
			assert.Nil(t, result, tc.testName)
			assert.NotNil(t, err, tc.testName)
			AssertLastLogMessage(t, hook, tc.testName, "error", "")
		} else if tc.testCase == noRecordsFound {
			assert.Equal(t, 0, len(result), tc.testName)
			assert.Nil(t, err, tc.testName)
		} else {
			if assert.NotNil(t, 0, len(result), tc.testName) {
				for i := range expectedCrudRecords {
					assert.Equal(t, expectedCrudRecords[i], result[i], tc.testName)
				}
			}
			assert.Nil(t, err, tc.testName)
			param, found := findParams["p"]
			if assert.True(t, found, tc.testName, "params had a value for \"p\"") {
				assert.Equal(t, "1", param, tc.testName)
			}
		}
	}
}

func TestCrudDao_ReadById(t *testing.T) {
	type testCase int
	const (
		failedRead testCase = iota
		noRecordFound
		success
	)

	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"fail CrudDao.Read", failedRead},
		{"no records found", noRecordFound},
		{"success", success},
	}

	// Test setup
	defer spy.UnpatchAll()

	for _, tc := range testCases {
		// Setup
		d := &crudDao{}

		expectedResult := &crudDaoModel{Id: "id"}

		// Setup the the mock for "Read"
		spy.PatchInstanceMethod(
			reflect.TypeOf(d), "Read",
			// callback function matching the signature of the "Read" method that is being mocked including the function response parameters.
			func(d *crudDao, ctx context.Context, logger l.Logger, params bson.M, decodeModelCallback func(bsonBytes []byte, id string) (interface{}, error), opts ...*options.FindOptions) ([]CrudDaoModel, error) {
				if tc.testCase == failedRead {
					return nil, errors.New("")
				} else if tc.testCase == noRecordFound {
					return []CrudDaoModel{}, nil
				}

				return []CrudDaoModel{
					expectedResult,
				}, nil
			})

		// Mock the logger
		logger, _ := l.CreateNullLogger()

		// Test
		result, found, err := d.ReadById(context.Background(), logger, "id", nil)

		// Verify
		if tc.testCase == failedRead {
			assert.Nil(t, result, tc.testName)
			assert.False(t, found, tc.testName)
			assert.NotNil(t, err, tc.testName)
		}
		if tc.testCase == noRecordFound {
			assert.Nil(t, result, tc.testName)
			assert.False(t, found, tc.testName)
			assert.Nil(t, err, tc.testName)
		}
		if tc.testCase == success {
			assert.Equal(t, expectedResult, result, tc.testName)
			assert.True(t, found, tc.testName)
			assert.Nil(t, err, tc.testName)
		}

	}
}

func TestCrudDao_Update(t *testing.T) {
	// define cases
	type testCase int
	const (
		failNoParams = iota
		failUpdateOne
		failNoModelsUpdated
		success
	)

	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"fail no query parameters", failNoParams},
		{"fail UpdateOne", failUpdateOne},
		{"fail No Models Updated", failNoModelsUpdated},
		{"success", success},
	}

	defer spy.UnpatchAll()

	for _, tc := range testCases {
		// Setup
		collection := &mongo.Collection{}
		updatedRecord := &testModel{
			s: "1",
		}
		params := bson.M{}

		if tc.testCase != failNoParams {
			params["p"] = "i"
		}

		// Mock the logger
		logger, hook := l.CreateNullLogger()

		d := crudDao{
			config: &CrudDaoConfig{},
		}

		// mock UpdateOne Called
		spy.PatchInstanceMethod(
			reflect.TypeOf(collection), "UpdateOne",
			func(c *mongo.Collection, ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
				m := update.(bson.M)

				if tc.testCase == failUpdateOne {
					return nil, errors.New("")
				}

				if set, ok := m["$set"]; !ok {
					assert.Fail(t, "$set not found", tc.testName)
				} else {
					mset := set.(bson.M)
					if model, ok := mset["model"]; !ok {
						assert.Fail(t, "model not found", tc.testName)
					} else {
						assert.Equal(t, updatedRecord, model, tc.testName)
					}
				}

				params := filter.(bson.M)
				if tc.testCase == success {
					v, _ := params["p"]
					assert.Equal(t, "i", v)
				}

				result := &mongo.UpdateResult{}
				if tc.testCase == failNoModelsUpdated {
					result.MatchedCount = 0
					result.ModifiedCount = 0
				} else {
					result.MatchedCount = 1
					result.ModifiedCount = 1
				}

				return result, nil
			})

		// test
		err := d.Update(context.Background(), logger, params, updatedRecord)

		// verify
		if tc.testCase < success {
			assert.NotNil(t, err, tc.testName)
			AssertLastLogMessage(t, hook, tc.testName, "error", "")
		} else {
			assert.Nil(t, err, tc.testName)
			AssertNoWarnOrErrorLogMessages(t, hook, tc.testName)
		}
	}
}

func TestCrudDao_UpdateById(t *testing.T) {
	// Test setup
	defer spy.UnpatchAll()

	// Setup
	d := &crudDao{}

	expectedId := "abc123"
	expectedResult := &crudDaoModel{Id: "id"}
	expectedError := errors.New("")

	// Setup the the mock for "Update"
	spy.PatchInstanceMethod(
		reflect.TypeOf(d), "Update",
		// callback function matching the signature of the "Update" method that is being mocked including the function response parameters.
		func(d *crudDao, ctx context.Context, logger l.Logger, params bson.M, model interface{}) error {
			id, _ := params["_id"]
			assert.Equal(t, expectedId, id)
			assert.Equal(t, expectedResult, model)
			return expectedError
		})

	// Mock the logger
	logger, _ := l.CreateNullLogger()

	// Test
	err := d.UpdateById(context.Background(), logger, expectedId, expectedResult)

	// Verify
	assert.Equal(t, expectedError, err)
}

func TestCrudDao_Delete(t *testing.T) {
	type testCase int
	const (
		failNoParams testCase = iota
		failUpdateMany
		failDeleteMany
		successUpdateMany
		successDeleteMany
	)

	testCases := []struct {
		testName string
		testCase testCase
	}{
		{"fail NoParams", failNoParams},
		{"fail UpdateMany", failUpdateMany},
		{"fail DeleteMany", failDeleteMany},
		{"success UpdateMany", successUpdateMany},
		{"success UpdateDelete", successDeleteMany},
	}

	defer spy.UnpatchAll()

	for _, tc := range testCases {
		// Setup
		collection := &mongo.Collection{}

		var params bson.M
		if tc.testCase != failNoParams {
			params = bson.M{"k": "v"}
		}

		config := &CrudDaoConfig{
			SoftDelete: false,
		}
		if tc.testCase == failUpdateMany || tc.testCase == successUpdateMany {
			config.SoftDelete = true
		}

		// Mock collection.UpdateMany
		spy.PatchInstanceMethod(
			reflect.TypeOf(collection), "UpdateMany",
			func(collection *mongo.Collection, ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
				if tc.testCase == failUpdateMany {
					return nil, errors.New("")
				}

				// Get the passed in params
				b, _ := filter.(bson.M)
				kValue, _ := b["k"]
				assert.Equal(t, "v", kValue, tc.testName)

				b, _ = update.(bson.M)
				if b, ok := b["$set"]; !ok {
					assert.Fail(t, "$set not set in update", tc.testName)
				} else {
					if b, ok := b.(bson.M); !ok {
						assert.Fail(t, "$set's value is not a bson.M", tc.testName)
					} else {
						deleteValue, _ := b[DeletedKey]
						if assert.NotNil(t, deleteValue, tc.testCase) {
							if history, ok := deleteValue.(*ActionHistory); !ok {
								assert.Fail(t, "delete history not set", tc.testName)
							} else {
								assert.NotNil(t, history.At, tc.testName)
							}
						}
					}
				}

				return nil, nil
			})

		// Mock collection.DeleteMany
		spy.PatchInstanceMethod(
			reflect.TypeOf(collection), "DeleteMany",
			func(collection *mongo.Collection, ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
				if tc.testCase == failDeleteMany {
					return nil, errors.New("")
				}
				b, _ := filter.(bson.M)
				kValue, _ := b["k"]
				assert.Equal(t, "v", kValue, tc.testName)

				return nil, nil
			})

		// Create the mock logger
		logger, hook := l.CreateNullLogger()

		// Create CrudDao
		d := &crudDao{
			client: createMockMongoClient(collection),
			config: config,
		}

		// Test
		err := d.Delete(context.Background(), logger, params)

		// Verify
		if tc.testCase < successUpdateMany {
			assert.NotNil(t, err, tc.testName)
			AssertLastLogMessage(t, hook, tc.testName, "error", "")
		} else {
			assert.Nil(t, err, tc.testName)
			AssertNoWarnOrErrorLogMessages(t, hook, tc.testName)
		}
	}
}

func TestCrudDao_DeleteById(t *testing.T) {
	// Test setup
	defer spy.UnpatchAll()

	// Setup
	d := &crudDao{}

	expectedId := "abc123"
	expectedError := errors.New("")

	// Setup the the mock for "Update"
	spy.PatchInstanceMethod(
		reflect.TypeOf(d), "Delete",
		// callback function matching the signature of the "Update" method that is being mocked including the function response parameters.
		func(d *crudDao, ctx context.Context, logger l.Logger, params bson.M) error {
			id, _ := params["_id"]
			assert.Equal(t, expectedId, id)
			return expectedError
		})

	// Mock the logger
	logger, _ := l.CreateNullLogger()

	// Test
	err := d.DeleteById(context.Background(), logger, expectedId)

	// Verify
	assert.Equal(t, expectedError, err)
}

func createMockMongoClient(mockCollection *mongo.Collection) *mongo.Client {
	client := &mongo.Client{}
	database := &mongo.Database{}
	spy.PatchInstanceMethod(
		reflect.TypeOf(client), "Database",
		func(client *mongo.Client, name string, opts ...*options.DatabaseOptions) *mongo.Database {
			return database
		})
	spy.PatchInstanceMethod(
		reflect.TypeOf(database), "Collection",
		func(database *mongo.Database, name string, opts ...*options.CollectionOptions) *mongo.Collection {
			return mockCollection
		})
	return client
}

func AssertLastLogMessage(t *testing.T, hook *test.Hook, testName string, expectedLevel string, expectedMessage string) {
	if assert.NotNil(t, hook.LastEntry(), testName, "should have made a log message") {
		assert.Equal(t, expectedLevel, hook.LastEntry().Level.String(), testName, "log message at wrong level")
		if expectedMessage != "" {
			assert.Equal(t, expectedMessage, hook.LastEntry().Message, testName)
		}
	}
}

func AssertNoWarnOrErrorLogMessages(t *testing.T, hook *test.Hook, testName string) {
	for _, entry := range hook.Entries {
		if entry.Level.String() == "error" || entry.Level.String() == "warning" {
			assert.Fail(t, "no warning or error log messages should have been created", testName, entry.Message)
		}
	}
}

type testModel struct {
	s string `bson:"s"`
}
