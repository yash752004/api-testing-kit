package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             string
	DatabaseURL      string
	DatabaseMaxConns int32
}

func Load() Config {
	return Config{
		Port:             getEnv("API_PORT", "8080"),
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		DatabaseMaxConns: getEnvInt32("DATABASE_MAX_CONNS", 4),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getEnvInt32(key string, fallback int32) int32 {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil || parsed <= 0 {
		return fallback
	}

	return int32(parsed)
}
