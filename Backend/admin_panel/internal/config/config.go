package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string
	AppPort  string
	MongoURI string
	MongoDB  string
}

func getenv(key string) string {
	return os.Getenv(key)
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables: %v", err)
	}

	return &Config{
		AppEnv:   getenv("APP_ENV"),
		AppPort:  getenv("APP_PORT"),
		MongoURI: getenv("MONGO_URI"),
		MongoDB:  getenv("MONGO_DB"),
	}
}
