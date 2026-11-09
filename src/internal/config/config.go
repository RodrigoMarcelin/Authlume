package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL           string
	DBPoolMaxConns        int
	DBPoolMinConns        int
	DBPoolMaxConnLifetime int
	DBPoolMaxConnIdle     int
	Argon2MemoryMib       int
	Argon2Time            int
	Argon2Parallelism     int
	Argon2SaltLength      int
	Argon2KeyLength       int
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "postgres://authlume:authlume_dev@localhost:5432/authlume?sslmode=disable"
	}

	db_pool_max_conns, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_CONNS"))
	if err != nil {
		db_pool_max_conns = 10
	}

	db_pool_min_conns, err := strconv.Atoi(os.Getenv("DB_POOL_MIN_CONNS"))
	if err != nil {
		db_pool_min_conns = 2
	}

	db_pool_max_conn_lifetime, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_CONN_LIFETIME"))
	if err != nil {
		db_pool_max_conn_lifetime = 60
	}

	db_pool_max_conn_idle, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_CONN_IDLE"))
	if err != nil {
		db_pool_max_conn_idle = 15
	}

	argon2MemoryMib, err := strconv.Atoi(os.Getenv("ARGON2_MEMORY_MIB"))
	if err != nil {
		argon2MemoryMib = 64
	}

	argon2Time, err := strconv.Atoi(os.Getenv("ARGON2_TIME"))

	if err != nil {
		argon2Time = 1
	}

	argon2Parallelism, err := strconv.Atoi(os.Getenv("ARGON2_PARALLELISM"))

	if err != nil {
		argon2Parallelism = 4
	}

	argon2SaltLength, err := strconv.Atoi(os.Getenv("ARGON2_SALT_LENGTH"))

	if err != nil {
		argon2SaltLength = 16
	}

	argon2KeyLength, err := strconv.Atoi(os.Getenv("ARGON2_KEY_LENGTH"))

	if err != nil {
		argon2KeyLength = 32
	}

	return &Config{
		DatabaseURL:           databaseUrl,
		DBPoolMaxConns:        db_pool_max_conns,
		DBPoolMinConns:        db_pool_min_conns,
		DBPoolMaxConnLifetime: db_pool_max_conn_lifetime,
		DBPoolMaxConnIdle:     db_pool_max_conn_idle,
		Argon2MemoryMib:       argon2MemoryMib,
		Argon2Time:            argon2Time,
		Argon2Parallelism:     argon2Parallelism,
		Argon2SaltLength:      argon2SaltLength,
		Argon2KeyLength:       argon2KeyLength,
	}, nil
}
