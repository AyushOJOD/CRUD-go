package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() *Config {

	_ = godotenv.Load()

	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBName: getEnv("DB_NAME", "taskmanager"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return fallback
}