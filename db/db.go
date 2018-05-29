// database implementations
package db

import (
	// Standard lib
	"fmt"

	// Third-party
	"github.com/jmoiron/sqlx"
)

const (
	// Database "types"
	DBTypeFC = "fc-db"
)

type (
	// Interface that all database implementations must fulfill
	DB interface {
		// Close is used to call a `Close` method of the underlying database driver
		Close() error
		// GetInstance returns an internal sqlx.DB instance to allow
		// for direct access to the database
		GetInstance() *sqlx.DB
		// SetInstance sets an internal sqlx.DB instance to allow
		// for direct access to the database
		SetInstance(*sqlx.DB) DB
		// String returns the "type" of database as a string - used in health checks
		String() string
		// Ready checks if the DB is "ready" - used in health checks
		Ready() error
	}
	// Struct representing configuration settings that should be used to create
	// a database DSN string
	DSNConfig struct {
		Username     string       // The username to use when connecting to the DB server
		Password     string       // The password to use when connecting to the DB server
		DatabaseName string       // The name of the database to use
		TCP          DSNConfigTCP // Struct containing information for TCP connections
		Timeout      int          // The max time (in seconds) to wait for operations to complete
		Dialect      string       // The database dialect to use
	}
	// Struct containing information for TCP connections
	DSNConfigTCP struct {
		Host string
		Port int
	}
)

// formDSN takes a configuration struct and uses it's values to form a
// DSN (Data Source Name) that takes the form of:
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func formDSN(c DSNConfig) string {
	// Set protocol and address
	protocol := "tcp"
	address := fmt.Sprintf("%s:%d", c.TCP.Host, c.TCP.Port)

	// Set template based on dialect
	var template string
	switch c.Dialect {
	case "postgres":
		template = "postgres://%s:%s@%s%s/%s?connect_timeout=%d&connect_timeout=%d&sslmode=disable"
		protocol = ""
		break
	default: // NOTE: MySQL
		template = "%s:%s@%s(%s)/%s?timeout=%ds&readTimeout=%ds"
		break
	}

	// Return formatted DSN
	return fmt.Sprintf(template, c.Username, c.Password, protocol, address, c.DatabaseName, c.Timeout, c.Timeout)
}
