package ports

import (
	"context"

	"github.com/RodrigoMarcelin/Authlume/src/internal/domain"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}
