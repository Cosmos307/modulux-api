package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration values
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "mydb"),
	}

	return config
}

// getEnv reads an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
