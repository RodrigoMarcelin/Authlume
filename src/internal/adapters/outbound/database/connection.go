package database

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

type ConnectionConfig struct {
	DatabaseURL         string
	PoolMaxConns        int
	PoolMinConns        int
	PoolMaxConnLifetime int
	PoolMaxConnIdle     int
}

func NewConnectionConfig(databaseURL string, poolMaxConns int, poolMinConns int, poolMaxConnLifetime int, poolMaxConnIdle int) *ConnectionConfig {
	return &ConnectionConfig{
		DatabaseURL:         databaseURL,
		PoolMaxConns:        poolMaxConns,
		PoolMinConns:        poolMinConns,
		PoolMaxConnLifetime: poolMaxConnLifetime,
		PoolMaxConnIdle:     poolMaxConnIdle,
	}
}

func (config *ConnectionConfig) NewConnection(ctx context.Context) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: tracelog.LoggerFunc(
			func(
				ctx context.Context,
				level tracelog.LogLevel,
				msg string,
				data map[string]any,
			) {
				slog.Log(ctx, slog.LevelDebug, msg, "details", data)
			},
		),
		LogLevel: tracelog.LogLevelDebug,
	}

	poolConfig.MaxConns = int32(config.PoolMaxConns)
	poolConfig.MinConns = int32(config.PoolMinConns)
	poolConfig.MaxConnLifetime = time.Duration(config.PoolMaxConnLifetime) * time.Minute
	poolConfig.MaxConnIdleTime = time.Duration(config.PoolMaxConnIdle) * time.Minute
	return pgxpool.NewWithConfig(ctx, poolConfig)
}
