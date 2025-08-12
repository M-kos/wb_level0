package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log/slog"
)

const (
	EnvFile = ".env"
)

type PostgresConfig struct {
	User     string `envconfig:"POSTGRES_USER" default:"root"`
	Password string `envconfig:"POSTGRES_PASSWORD" default:"root"`
	Name     string `envconfig:"POSTGRES_NAME" default:"wb"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
}

type Config struct {
	Port      int            `envconfig:"SERVICE_PORT" default:"8080"`
	LogLevel  string         `envconfig:"LOG_LEVEL" default:"DEBUG"`
	CacheSize int            `envconfig:"CACHE_SIZE" default:"100"`
	Postgres  PostgresConfig `envconfig:"POSTGRES"`
}

func New() *Config {
	err := godotenv.Load(EnvFile)
	if err != nil {
		slog.Error(err.Error())
	}

	var c Config

	err = envconfig.Process("", &c)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &c
}
