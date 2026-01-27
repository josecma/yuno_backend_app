package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port         string
	Environment  string
	ReadTimeout  int
	WriteTimeout int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
	Expiry string
}

func Load() *Config {

	godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "3000"),
			Environment:  getEnv("ENV", "development"),
			ReadTimeout:  getIntEnv("READ_TIMEOUT", 15),
			WriteTimeout: getIntEnv("WRITE_TIMEOUT", 15),
		},
		Database: DatabaseConfig{
			Host:     getEnv("POSTGRES_DB_HOST", "localhost"),
			Port:     getEnv("POSTGRES_DB_PORT", "5434"),
			User:     getEnv("POSTGRES_DB_USER", "postgres"),
			Password: getEnv("POSTGRES_DB_PASSWORD", "postgres"),
			Name:     getEnv("POSTGRES_DB_NAME", "yuno_db_dev"),
			SSLMode:  getEnv("POSTGRES_DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "secret"),
			Expiry: getEnv("JWT_EXPIRY", "24h"),
		},
	}

}

func getEnv(key, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {

		return value

	}

	return defaultValue

}

func getIntEnv(key string, defaultValue int) int {

	strVal := getEnv(key, "")

	if strVal == "" {

		return defaultValue

	}

	val, err := strconv.Atoi(strVal)

	if err != nil {

		return defaultValue

	}

	return val

}
