package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log/slog"
)

const (
	EnvFile      = ".env"
	EnvLocalFile = ".env.local"
)

type PostgresConfig struct {
	User     string `envconfig:"POSTGRES_USER" default:"root"`
	Password string `envconfig:"POSTGRES_PASSWORD" default:"root"`
	Name     string `envconfig:"POSTGRES_NAME" default:"wb"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
}

type KafkaConfig struct {
	Host          string `envconfig:"KAFKA_HOST" default:"localhost"`
	Port          string `envconfig:"KAFKA_PORT" default:"9092"`
	ExternalPort  string `envconfig:"KAFKA_EXTERNAL_PORT" default:"29092"`
	Topic         string `envconfig:"KAFKA_TOPIC" default:"orders"`
	GroupID       string `envconfig:"KAFKA_GROUP_ID" default:"order-consumer"`
	MinBytes      int    `envconfig:"KAFKA_MIN_BYTES" default:"1"`
	MaxBytes      int    `envconfig:"KAFKA_MAX_BYTES" default:"10000000"`
	MaxRetries    int    `envconfig:"KAFKA_MAX_RETRIES" default:"5"`
	DlqTopic      string `envconfig:"DLQ_TOPIC" default:"order-dlq"`
	MaxDlqRetries int    `envconfig:"KAFKA_MAX_RETRIES" default:"3"`
}

type Config struct {
	Port       int            `envconfig:"SERVICE_PORT" default:"8083"`
	LogLevel   string         `envconfig:"LOG_LEVEL" default:"DEBUG"`
	CacheSize  int            `envconfig:"CACHE_SIZE" default:"100"`
	Postgres   PostgresConfig `envconfig:"POSTGRES"`
	Kafka      KafkaConfig    `envconfig:"KAFKA"`
	DbHost     string         `envconfig:"DB_HOST" default:"postgres"`
	BrokerHost string         `envconfig:"BROKER_HOST" default:"kafka"`
}

func New() *Config {
	err := godotenv.Load(EnvLocalFile)
	if err != nil {
		slog.Error(err.Error())

		err = godotenv.Load(EnvFile)
		if err != nil {
			slog.Error(err.Error())
		}
	}

	var c Config

	err = envconfig.Process("", &c)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	slog.Info("config is loaded", c)

	return &c
}
