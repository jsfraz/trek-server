package utils

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm/logger"
)

type Config struct {

	// Gin mode
	GinMode string `envconfig:"GIN_MODE" required:"true"`

	// Postgres
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresServer   string `envconfig:"POSTGRES_SERVER" required:"true"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresDb       string `envconfig:"POSTGRES_DB" required:"true"`

	// Access token
	AccessTokenLifespan int    `envconfig:"ACCESS_TOKEN_LIFESPAN" required:"true"`
	AccessTokenSecret   string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`

	// Tracker token
	TrackerTokenSecret string `envconfig:"TRACKER_TOKEN_SECRET" required:"true"`

	// Superuser
	SuperuserUsername string `envconfig:"SUPERUSER_USERNAME" required:"true"`
	SuperuserPassword string `envconfig:"SUPERUSER_PASSWORD" required:"true"`
}

// Returns config from ENV values.
//
//	@return *Config
//	@return error
func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	return &config, nil
}

// Returns the Gorm log level according to the environment variable
//
//	@receiver c
//	@return logger.LogLevel
func (c *Config) GetGormLogLevel() logger.LogLevel {
	if c.GinMode == "release" {
		return logger.Error
	}
	return logger.Info
}
