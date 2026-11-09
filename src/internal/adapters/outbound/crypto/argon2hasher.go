package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Hasher struct {
	MemoryMib   int
	Time        int
	Parallelism int
	SaltLength  int
	KeyLength   int
}

func NewArgon2Hasher(memoryMib, time, parallelism, saltLength, keyLength int) *Argon2Hasher {
	return &Argon2Hasher{
		MemoryMib:   memoryMib,
		Time:        time,
		Parallelism: parallelism,
		SaltLength:  saltLength,
		KeyLength:   keyLength,
	}
}

func (h *Argon2Hasher) CreateSaltPassword() ([]byte, error) {
	newByteSalt := make([]byte, h.SaltLength)
	_, err := rand.Read(newByteSalt)
	if err != nil {
		return nil, err
	}
	return newByteSalt, nil
}

func (h *Argon2Hasher) CreateHashPassword(password string, salt []byte) []byte {
	passwordByte := []byte(password)
	hashByte := argon2.IDKey(
		passwordByte,
		salt,
		uint32(h.Time),
		uint32(h.MemoryMib)*1024,
		uint8(h.Parallelism),
		uint32(h.KeyLength),
	)
	return hashByte
}

func (h *Argon2Hasher) Argon2IdHash(hash []byte, salt []byte) string {
	hashKey := base64.RawStdEncoding.EncodeToString(hash)
	saltKey := base64.RawStdEncoding.EncodeToString(salt)
	argon2Hash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.MemoryMib*1024,
		h.Time,
		h.Parallelism,
		saltKey,
		hashKey,
	)
	return argon2Hash
}
