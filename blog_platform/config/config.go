package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASSWORD", "1234"),
		DBName:    getEnv("DB_NAME", "blog_db"),
		DBPort:    getEnv("DB_PORT", "5432"),
		JWTSecret: getEnv("JWT_SECRET", "ndjoqwhruew409i29eio"),
	}
}
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
