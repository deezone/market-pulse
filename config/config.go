// config package handles setting up, parsing, and storing all configuration settings for the application.
// The main package sets up an instance of this package and all other packages can use the "GetInstance"
// method to get a reference to the config struct and all of it's values
package config

import (
	// Standard lib
	"encoding/json"
	"time"

	// Third-party
	"github.com/marksost/configurator"
	log "github.com/sirupsen/logrus"
)

const (
	// Environment names
	EnvDevelopment = "dev"
	EnvProduction  = "prod"
	EnvTesting     = "test"

	// Environment variable prefix
	EnvPrefix = "ForexClock_"

	// Environment variable where the path to an outside
	// configuration file is located
	ConfigLocation = EnvPrefix + "CONFIG"

	// API versions
	Version10 = "1.0"
)

var (
	config *Config

	// Supported API versions
	SupportedVersions = []string{Version10}
)

type (
	// Component-specific configuration

	// Struct containing configuration settings for a database
	DB struct {
		// The username to use when connecting to a DB server
		Username string `json:"username" env:"DB_USERNAME" default:""`
		// The password to use when connecting to a DB server
		Password string `json:"password" env:"DB_PASSWORD" default:""`
		// The name of the database to use
		DatabaseName string `json:"database-name" env:"DB_DB_NAME" default:""`
		// Struct containing information for TCP connections
		TCP DBTCP `json:"tcp"`
		// The max time (in seconds) to wait for operations to complete
		Timeout int `json:"timeout" env:"DB_TIMEOUT" default:"5"`
		// The database dialect to use
		Dialect string `json:"dialect" env:"DB_DIALECT" default:"mysql"`
	}

	// Struct containing configuration settings for a database TCP connection
	DBTCP struct {
		// The host of a DB server
		Host string `json:"host" env:"DB_TCP_HOST" default:"localhost"`
		// The port of a DB server
		Port int `json:"port" env:"DB_TCP_PORT" default:"3306"`
	}

	// Struct containing configuration settings for application logging
	Log struct {
		// The formatter to use
		Formatter string `json:"formatter" env:"LOG_FORMATTER" default:"text"`
		// The log level to use
		Level string `json:"level" env:"LOG_LEVEL" default:"info"`
	}

	// Struct containing configuration settings for the application server
	Server struct {
		// Port the server should listen on
		Port int `json:"port" env:"SERVER_PORT" default:"6010"`
		// Various timeouts for the server
		Timeouts struct {
			// Timeout (in seconds) allowed for server read operations
			Read int `json:"read" env:"SERVER_READ_TIMEOUT" default:"30"`
			// Timeout (in seconds) allowed for server to shutdown
			ShutDown int `json:"shutdown" env:"SERVER_SHUTDOWN_TIMEOUT" default:"5"`
			// Timeout (in seconds) allowed for server write operations
			Write int `json:"write" env:"SERVER_WRITE_TIMEOUT" default:"30"`
		} `json:"timeouts"`
	}

	// Config is a struct containing all configuration settings for the application.
	// NOTE: Only a single instance of this struct should be used throughout the application
	// so as to reference the same configuration state.
	Config struct {
		/* Top-level configuration */

		// The environment the application is running in
		Environment string `json:"environment" env:"ENVIRONMENT" default:"dev"`

		// Name of the application
		Name string `json:"name" env:"NAME" default:"fed"`

		// The current release version of the application
		ReleaseVersion string `json:"release-version" env:"RELEASE_VERSION" default:""`

		// Start time of the application
		StartTime time.Time

		/* Component-specific configuration */

		// Settings for the database
		DB DB `json:"db"`

		// Settings for the logger
		Log Log `json:"log"`

		// Settings for the server
		Server Server `json:"server"`
	}
)

// setLoggerSettings sets the application logger's various properties
func (c *Config) setLoggerSettings() {
	// Set logging level based on config value
	switch c.Log.Level {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	// Set logging formatter based on config value
	// NOTE: Add more cases here for custom formatters
	switch c.Log.Formatter {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}

// Init creates a new config instance and initializes it
func Init() {
	// Create new config instance
	config = &Config{}

	// Used configurator to set up config
	configurator.InitializeConfig(config)

	// Set logging settings
	config.setLoggerSettings()

	// Set start time
	config.StartTime = time.Now()
}

// GetInstance returns the initialized config instance
func GetInstance() *Config {
	// Check if config has been initialized yet
	if config == nil {
		Init()
	}

	return config
}

func init() {
	// Set up configurator settings
	configurator.EnvPrefix = EnvPrefix
	configurator.ConfigLocation = ConfigLocation
}
