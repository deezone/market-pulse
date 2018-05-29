// db package contains all database implementations this application will need
// sc-db defines interaction with the SoulCycle central database
package db

import (
	// Standard lib
	"errors"

	// Internal
	"github.com/deezone/forex-clock/config"

	// Third-party
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

const (
	// Select queries

	// Insert queries
)

type (
	// Struct representing a single SoulCycle DB instance
	fcDB struct {
		dbType   string
		instance *sqlx.DB
	}
)

// NewSCDB creates and returns a new instance of a SoulCycle DB
func NewFCDB() DB {
	// Form configs
	c := config.GetInstance()
	dsnConfig := DSNConfig{
		Username:     c.DB.Username,
		Password:     c.DB.Password,
		DatabaseName: c.DB.DatabaseName,
		TCP:          DSNConfigTCP{Host: c.DB.TCP.Host, Port: c.DB.TCP.Port},
		Timeout:      c.DB.Timeout,
		Dialect:      c.DB.Dialect,
	}

	// Attempt to open DB, checking for errors
	// NOTE: Ignoring error here. `sqlx.Connect` only errors if the driver isn't imported
	// It's always imported, so no way for that to error
	i, _ := sqlx.Connect(dsnConfig.Dialect, formDSN(dsnConfig))

	// Form and return new DB
	return &fcDB{
		dbType:   DBTypeFC,
		instance: i,
	}
}

// Close is used to call a `Close` method of the underlying database driver
func (db *fcDB) Close() error { return db.instance.Close() }

// GetInstance returns an internal sqlx.DB instance to allow
// for direct access to the database
func (db *fcDB) GetInstance() *sqlx.DB { return db.instance }

// SetInstance sets an internal sqlx.DB instance to allow
// for direct access to the database
// NOTE: Returns the DB struct to allow for chaining of methods
func (db *fcDB) SetInstance(i *sqlx.DB) DB {
	db.instance = i
	return db
}

// String returns the "type" of database as a string - used in health checks
func (db *fcDB) String() string { return db.dbType }

// Ready checks if the DB is "ready" - used in health checks
func (db *fcDB) Ready() error {
	// Ensure database instance is valid
	if db.GetInstance() == nil {
		return errors.New("Nil database instance detected")
	}

	// Ping database
	return db.GetInstance().Ping()
}
