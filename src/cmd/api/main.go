package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/outbound/database"
	"github.com/RodrigoMarcelin/Authlume/src/internal/bootstrap"
	"github.com/RodrigoMarcelin/Authlume/src/internal/config"
	"github.com/go-chi/chi/v5"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	deConfig := database.NewConnectionConfig(
		cfg.DatabaseURL,
		cfg.DBPoolMaxConns,
		cfg.DBPoolMinConns,
		cfg.DBPoolMaxConnLifetime,
		cfg.DBPoolMaxConnIdle,
	)

	pool, err := deConfig.NewConnection(ctx)
	if err != nil {
		panic(err)
	}

	var databaseName string
	var serverAddress string
	var serverPort int

	err = pool.QueryRow(ctx, `
		SELECT
			current_database(),
			inet_server_addr()::text,
			inet_server_port()
	`).Scan(
		&databaseName,
		&serverAddress,
		&serverPort,
	)

	if err != nil {
		panic(err)
	}

	slog.Info(
		"PostgreSQL connection established",
		"database", databaseName,
		"server", serverAddress,
		"port", serverPort,
	)

	defer pool.Close()

	app := chi.NewRouter()

	err = bootstrap.Run(ctx, cfg, pool, app)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", app)
}
