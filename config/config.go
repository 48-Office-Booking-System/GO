package config

import (
	"os"
)

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
}

func InitConfiguration() Config {
	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "kelompok48"),
		DB_NAME:        GetOrDefault("DB_NAME", "koba-test"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "localhost"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "punten"),
	}
}

func GetOrDefault(envKey, defaultValue string) string {
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}

	return defaultValue
}
