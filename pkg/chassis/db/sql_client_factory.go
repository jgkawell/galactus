package db

import (
	"fmt"

	l "github.com/jgkawell/galactus/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// CreateSqlClient will create a SQL client to a Postgres database using gorm. If it fails it will
func CreateSqlClient(logger l.Logger, sqlDbUser, sqlDbSecret, sqlDbHost, sqlDbPort, sqlDbName, sqlDbSchema string, isDevMode bool) (*gorm.DB, l.Error) {
	logger = logger.WithFields(l.Fields{
		"user":   sqlDbUser,
		"host":   sqlDbHost,
		"port":   sqlDbPort,
		"name":   sqlDbName,
		"schema": sqlDbSchema,
	})
	logger.Info("creating sql client")

	// create the db connection string
	sqlDbString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?search_path=%s",
		sqlDbUser, sqlDbSecret, sqlDbHost, sqlDbPort, sqlDbName, sqlDbSchema)

	// disable ssl when running locally
	if isDevMode {
		sqlDbString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
			sqlDbUser, sqlDbSecret, sqlDbHost, sqlDbPort, sqlDbName, sqlDbSchema)
	}

	// open a connection to the target postgres datastore
	db, err := gorm.Open(postgres.Open(sqlDbString), &gorm.Config{
		FullSaveAssociations: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", sqlDbSchema),
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to connect to sql db"))
	}

	// check the schema, and table layout. If the db is missing the required schema, then it
	// will be created.
	if sqlDbSchema != "" {
		var schemas []string
		db.Raw("SELECT schema_name FROM information_schema.schemata").Scan(&schemas)
		// the target schema is not setup in the database
		if !doesStringSliceContainString(schemas, sqlDbSchema) {
			// schema does not exist so create it
			// NOTE: If the schema does not already exist and you're moving from a different datasource. Data may need to be
			//       migrated. This is not the responsiblity of the `chassis` module. It should be a different task
			//       that is automated.
			logger.WithField("schema", sqlDbSchema).Info("schema does not exist so it is being created")
			cs := fmt.Sprintf("CREATE SCHEMA %s", sqlDbSchema)
			tx := db.Exec(cs)
			if tx.Error != nil {
				return nil, logger.WithField("schema", sqlDbSchema).WrapError(l.NewError(err, "failed to create the schema"))
			}
		}
	}

	return db, nil
}

// DisconnectSqlClient will disconnect a gorm SQL client from a Postgres database
func DisconnectSqlClient(logger l.Logger, client *gorm.DB) l.Error {
	db, err := client.DB()
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to get the db connection for disconnect"))
	}
	db.Close()

	return nil
}

// PingSqlClient is called in a separate thread that is responding to health checks.
func PingSqlClient(logger l.Logger, client *gorm.DB) l.Error {
	db, err := client.DB()
	if err != nil {
		return logger.WrapError(l.NewError(err, "failed to get the db connection for ping"))
	}
	if err := db.Ping(); err != nil {
		return logger.WrapError(l.NewError(err, "failed to ping sql db"))
	}

	return nil
}
