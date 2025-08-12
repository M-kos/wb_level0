package db

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/logger"
	"net"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DBScheme = "postgres"

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, config *config.Config, log logger.Logger) (*PostgresDB, error) {
	pool, err := MakePostgresPool(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %w", err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	log.Info("connected to database")

	return &PostgresDB{
		Pool: pool,
	}, nil
}

func MakePostgresPool(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	connString := ConnectionString(config)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func ConnectionString(config *config.Config) string {
	query := url.Values{}
	query.Set("sslmode", "disable")

	value := url.URL{
		Scheme:   DBScheme,
		User:     url.UserPassword(config.Postgres.User, config.Postgres.Password),
		Host:     net.JoinHostPort(config.Postgres.Host, config.Postgres.Port),
		Path:     config.Postgres.Name,
		RawQuery: query.Encode(),
	}

	return value.String()
}
