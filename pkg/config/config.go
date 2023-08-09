package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	DBConf mysql.Config
}

func NewConfig() *Config {
	return &Config{DBConf: NewMySQLConfig()}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
