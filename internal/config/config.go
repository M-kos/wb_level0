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

type KafkaConfig struct {
	Host     string `envconfig:"KAFKA_HOST" default:"localhost"`
	Port     string `envconfig:"KAFKA_PORT" default:"9092"`
	Topic    string `envconfig:"KAFKA_TOPIC" default:"orders"`
	GroupID  string `envconfig:"KAFKA_GROUP_ID" default:"orderRepository-consumer"`
	MinBytes int    `envconfig:"KAFKA_MIN_BYTES" default:"1"`
	MaxBytes int    `envconfig:"KAFKA_MAX_BYTES" default:"10000000"`
}

type Config struct {
	Port      int            `envconfig:"SERVICE_PORT" default:"8080"`
	LogLevel  string         `envconfig:"LOG_LEVEL" default:"DEBUG"`
	CacheSize int            `envconfig:"CACHE_SIZE" default:"100"`
	Postgres  PostgresConfig `envconfig:"POSTGRES"`
	Kafka     KafkaConfig    `envconfig:"KAFKA"`
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
