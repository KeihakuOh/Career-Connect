package config

import (
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    AppEnv     string
}

func Load() *Config {
    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "labcareer"),
        DBPassword: getEnv("DB_PASSWORD", "labcareer_pass"),
        DBName:     getEnv("DB_NAME", "labcareer_db"),
        JWTSecret:  getEnv("JWT_SECRET", "dev_secret"),
        AppEnv:     getEnv("APP_ENV", "development"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
