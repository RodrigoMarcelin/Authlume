package usecase

import (
	"context"
	"errors"

	cryptoports "github.com/RodrigoMarcelin/Authlume/src/internal/application/ports/crypto"
	ports "github.com/RodrigoMarcelin/Authlume/src/internal/application/ports/database"
	"github.com/RodrigoMarcelin/Authlume/src/internal/domain"
	"github.com/jackc/pgx/v5"
)

type CreateUserUseCase struct {
	userRepository ports.UserRepository
	cryptoHasher   cryptoports.Argon2Hasher
}

func NewCreateUserUseCase(userRepository ports.UserRepository, cryptoHasher cryptoports.Argon2Hasher) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
		cryptoHasher:   cryptoHasher,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, email, password string) error {
	newUser, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}

	userDatabase, err := uc.userRepository.GetUserByEmail(ctx, newUser.GetEmail())

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if userDatabase != nil {
		return errors.New(`user already exists with this email ${newUser.email}`)
	}

	salt, err := uc.cryptoHasher.CreateSaltPassword()
	if err != nil {
		return err
	}

	hashPassword := uc.cryptoHasher.CreateHashPassword(newUser.GetPasswordHash(), salt)

	newUser.UpdatePassword(uc.cryptoHasher.Argon2IdHash(hashPassword, salt))

	if err := uc.userRepository.CreateUser(ctx, newUser); err != nil {
		return err
	}

	return nil

}
