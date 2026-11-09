package ports

type Argon2Hasher interface {
	CreateSaltPassword() ([]byte, error)
	CreateHashPassword(password string, salt []byte) []byte
	Argon2IdHash(hash []byte, salt []byte) string
}
