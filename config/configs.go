package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func LoadConfig() *Config {
	godotenv.Load()

	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("Can not get DSN")
	}

	return &Config{
		DSN: dsn,
	}
}
