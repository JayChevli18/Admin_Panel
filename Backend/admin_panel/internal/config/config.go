package config

import "os"

type Config struct {
	AppEnv   string
	AppPort  string
	MongoURI string
	MongoDB  string
}

func getenv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:   getenv("APP_ENV", "development"),
		AppPort:  getenv("APP_PORT", "8080"),
		MongoURI: getenv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getenv("MONGO_DB", "admin_panel"),
	}
}
