package bootstrap

import (
	"context"

	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/inbound/rest"
	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/outbound/crypto"
	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/outbound/database"
	"github.com/RodrigoMarcelin/Authlume/src/internal/application/usecase"
	"github.com/RodrigoMarcelin/Authlume/src/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(ctx context.Context, cfg *config.Config, pool *pgxpool.Pool, router *chi.Mux) error {

	userRepository := database.NewUserRepository(pool)

	criptoHasher := crypto.NewArgon2Hasher(cfg.Argon2MemoryMib, cfg.Argon2Time, cfg.Argon2Parallelism, cfg.Argon2SaltLength, cfg.Argon2KeyLength)

	createUserUseCase := usecase.NewCreateUserUseCase(userRepository, criptoHasher)

	rest.NewUserController(createUserUseCase).RoutesUser(router)

	return nil
}
