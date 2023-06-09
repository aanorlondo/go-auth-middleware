package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseHostname  string
	DatabaseName      string
	DatabasePort      string
	DatabaseTableName string
	DatabaseUsername  string
	DatabasePassword  string
	RedisHostname     string
	RedisPort         string
	RedisPassword     string
	AppSecretKey      string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	config := &Config{
		DatabaseHostname:  getEnv("DATABASE_HOSTNAME", "undefined"),
		DatabaseName:      getEnv("DATABASE_NAME", "undefined"),
		DatabasePort:      getEnv("DATABASE_PORT", "undefined"),
		DatabaseTableName: getEnv("DATABASE_TABLENAME", "undefined"),
		DatabaseUsername:  getEnv("DATABASE_USERNAME", "undefined"),
		DatabasePassword:  getEnv("DATABASE_PASSWORD", "undefined"),
		RedisHostname:     getEnv("REDIS_HOSTNAME", "undefined"),
		RedisPort:         getEnv("REDIS_PORT", "undefined"),
		RedisPassword:     getEnv("REDIS_PASSWORD", "undefined"),
		AppSecretKey:      getEnv("APP_SECRET_KEY", "undefined"),
	}

	return config, nil
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHostname, c.DatabasePort, c.DatabaseName)
}

func (c *Config) GetDatabaseTableName() string {
	return c.DatabaseTableName
}

func (c *Config) GetAppSecretKey() string {
	return c.AppSecretKey
}

func (c *Config) GetRedisURL() string {
	return fmt.Sprintf("%s:%s", c.RedisHostname, c.RedisPort)
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}
