package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBurl string
	Port  string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		DBurl: os.Getenv("DB_URL"),
		Port: os.Getenv("PORT"),
	}
}