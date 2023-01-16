package db

import (
	"fmt"

	log "github.com/techotron/dog-ceo-api/backend/log"
	"github.com/techotron/dog-ceo-api/backend/config"

	"github.com/jmoiron/sqlx"

	// create package-level variables and execute the init function of that package.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

const (
    // ConnTypeMigrate for migrate operations
    ConnTypeMigrate = "migrate"
    // ConnTypeApi for api operations
    ConnTypeApi = "api"
)

// Conn to database, can be used for all database queries
var Conn *sqlx.DB = nil

// MigrateConn is the connection to use just for migration operations
var MigrateConn *sqlx.DB = nil

// SetupConn to database, use db.Conn to execute any queries, this func is only used to setup the initial connection pool.
func SetupConn(connType string) error {
	appConfig, err := config.GetAppConfigFromEnv()
	if err != nil {
		return err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName, appConfig.DBSSLMode)

	// wrapper of https://golang.org/pkg/database/sql/#Open
	// creates the initial connection pool to the database
	if connType == ConnTypeMigrate {
        MigrateConn, err = sqlx.Open("postgres", psqlInfo)
		if err != nil {
			return err
		}
	} else {
		Conn, err = sqlx.Open("postgres", psqlInfo)
		if err != nil {
			return err
		}		
    }

	err = Conn.Ping()
	if err != nil {
		log.Error(err)
		log.Error("Database connection could not be set, ensure that database credentials are correctly")
		return err
	}
	log.Info("Database connection is set")
	return nil
}

// CloseConn to the database (all connections)
func CloseConn() error {
	return Conn.Close()
}

// MigrateUp the database to current version
func MigrateUp() error {
	m, err := initMigrate()
	if err != nil {
		return err
	}

	currentVersion, _, _ := m.Version()
	log.Info("Migrating Database from version ", currentVersion)
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Info(err)
			return nil
		}
		return err
	}
	currentVersion, _, _ = m.Version()
	log.Info("Migrated Database to version ", currentVersion)
	return nil
}

// Drop deletes everything in the database
func Drop() error {
	m, err := initMigrate()
	if err != nil {
		return err
	}

	err = m.Drop()
	if err != nil {
		return err
	}
	return nil
}

// Version of database using go.migrate.Version()
func Version() (version uint, dirty bool, err error) {
	m, err := initMigrate()
	if err != nil {
		return version, dirty, err
	}

	return m.Version()
}

// StepMigrate migrates database based on input steps. E.g. if you are on v39 and steps param is `-1`, db will be migrated down to v38
func StepMigrate(steps int) error {
	m, err := initMigrate()
	if err != nil {
		return err
	}

	return m.Steps(steps)
}

// initMigrate returns a migrate instance based on current db setup
func initMigrate() (*migrate.Migrate, error) {
	err := SetupConn(ConnTypeMigrate)
	if err != nil {
		log.Errorf("An error occured while setting up the migration db connection, %s", err)
	}
	appConfig, _ := config.GetAppConfigFromEnv()

	// https://github.com/golang-migrate/migrate#use-in-your-go-project
	// go migrate requires the sql.DB not sqlx.DB, first is available via sqlx.DB.DB
	driver, err := postgres.WithInstance(MigrateConn.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+appConfig.DBMigrationsPath,
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	MigrateConn.Close()
	return m, nil
}
