package domain

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id           *uuid.UUID
	email        string
	passwordHash string
	status       bool
	createdAt    time.Time
	modifiedAt   time.Time
	createdBy    *uuid.UUID
	modifiedBy   *uuid.UUID
}

func NewUser(email string, password string) (*User, error) {
	emailErr := EmailVerify(email)
	if emailErr != nil {
		return nil, emailErr
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	return &User{
		email:        email,
		passwordHash: password,
		status:       true,
	}, nil
}

func InstanceUser(id *uuid.UUID, email string, passwordHash string, status bool, createdAt time.Time, modifiedAt time.Time, createdBy *uuid.UUID, modifiedBy *uuid.UUID) (*User, error) {
	emailErr := EmailVerify(email)
	if emailErr != nil {
		return nil, emailErr
	}

	return &User{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		status:       status,
		createdAt:    createdAt,
		modifiedAt:   modifiedAt,
		createdBy:    createdBy,
		modifiedBy:   modifiedBy,
	}, nil
}

func EmailVerify(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) GetID() *uuid.UUID {
	return u.id
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPasswordHash() string {
	return u.passwordHash
}

func (u *User) GetStatus() bool {
	return u.status
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetModifiedAt() time.Time {
	return u.modifiedAt
}

func (u *User) GetCreatedBy() *uuid.UUID {
	return u.createdBy
}

func (u *User) GetModifiedBy() *uuid.UUID {
	return u.modifiedBy
}

func (u *User) UpdateEmail(newEmail string) error {
	emailErr := EmailVerify(newEmail)
	if emailErr != nil {
		return emailErr
	}
	u.email = newEmail
	return nil
}

func (u *User) UpdatePassword(newPassword string) error {
	if newPassword == "" {
		return errors.New("password cannot be empty")
	}
	u.passwordHash = newPassword
	return nil
}
