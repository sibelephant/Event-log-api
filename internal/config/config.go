package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	DatabaseURL string
	JWTSecret   string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Parse port with default value
	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	// Get JWT secret with default
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-change-in-production"
		log.Println("Warning: JWT_SECRET not set, using default (not secure for production)")
	}

	return &Config{
		Port:        port,
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   jwtSecret,
	}
}

// GetEnv returns an environment variable with a fallback default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
