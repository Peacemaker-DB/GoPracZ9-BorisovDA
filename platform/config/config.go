package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DSN     string
	BcryptCost int
	Addr       string
}

func Load() Config {
	_ = godotenv.Load()

	cost := 12
	if v := os.Getenv("BCRYPT_COST"); v != "" {
	}
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return Config{
		DB_DSN:     os.Getenv("DATABASE_URL"),
		BcryptCost: cost,
		Addr:       addr,
	}
}
