package config

import (
	"fmt"
	"os"
)

type Config struct {
	// Configuration variables
	DatabaseHostname string
	DatabaseUsername string
	DatabasePassword string
	AppSecretKey     string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	// Load and return the configuration
	config := &Config{
		DatabaseHostname: getEnv("DATABASE_HOSTNAME", "localhost"),
		DatabaseUsername: getEnv("DATABASE_USERNAME", "user"),
		DatabasePassword: getEnv("DATABASE_PASSWORD", "password"),
		AppSecretKey:     getEnv("APP_SECRET_KEY", "your-secret-key"),
	}

	return config, nil
}

func (c *Config) GetDatabaseURL() string {
	// Get the database URL
	return fmt.Sprintf("%s:%s@tcp(%s)/database", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHostname)
}

func (c *Config) GetAppSecretKey() string {
	// Get the app secret key
	return c.AppSecretKey
}
