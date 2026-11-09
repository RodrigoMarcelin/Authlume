package database

import (
	"context"
	"errors"
	"time"

	"github.com/RodrigoMarcelin/Authlume/src/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, status, created_at, modified_at, created_by, modified_by
		FROM public.users
		WHERE email = $1;
	`

	var (
		id           *uuid.UUID
		userEmail    string
		passwordHash string
		status       bool
		createdAt    time.Time
		modifiedAt   time.Time
		createdBy    *uuid.UUID
		modifiedBy   *uuid.UUID
	)

	err := r.pool.QueryRow(ctx, query, email).Scan(
		&id,
		&userEmail,
		&passwordHash,
		&status,
		&createdAt,
		&modifiedAt,
		&createdBy,
		&modifiedBy,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return domain.InstanceUser(
		id,
		userEmail,
		passwordHash,
		status,
		createdAt,
		modifiedAt,
		createdBy,
		modifiedBy,
	)
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO public.users (email, password_hash, status)
		VALUES ($1, $2, $3);
	`
	_, err := r.pool.Exec(ctx, query, user.GetEmail(), user.GetPasswordHash(), user.GetStatus())
	return err
}
