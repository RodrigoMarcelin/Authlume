package ports

import "context"

type ConnectionConfig interface {
	NewConnectionConfig(databaseURL string, poolMaxConns, poolMinConns, poolMaxConnLifetime, poolMaxConnIdle int) *ConnectionConfig
	NewConnection(ctx context.Context) (interface{}, error)
}
