package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all environment-driven configuration for the application.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	CORSOrigin string
}

// Load reads the .env file (if present) and populates a Config from environment variables.
func Load() *Config {
	// godotenv will not override variables that are already set in the environment.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading config from environment")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "apm"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		CORSOrigin: getEnv("CORS_ORIGIN", "http://localhost:4200"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
