package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	AppPort       string
	AppEnv        string
	AppURL        string
	SessionSecret string

	DBDriver     string
	DBSQLitePath string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string

	AdminEmail    string
	AdminPassword string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using OS environment")
	}
	return &Config{
		AppName:       env("APP_NAME", "LagariGo"),
		AppPort:       env("APP_PORT", "3000"),
		AppEnv:        env("APP_ENV", "development"),
		AppURL:        env("APP_URL", "http://localhost:3000"),
		SessionSecret: env("SESSION_SECRET", "change-me"),

		DBDriver:     env("DB_DRIVER", "sqlite"),
		DBSQLitePath: env("DB_SQLITE_PATH", "./lagarigo.db"),
		DBHost:       env("DB_HOST", "127.0.0.1"),
		DBPort:       env("DB_PORT", "3306"),
		DBUser:       env("DB_USER", "root"),
		DBPassword:   env("DB_PASSWORD", ""),
		DBName:       env("DB_NAME", "lagarigo"),

		AdminEmail:    env("ADMIN_EMAIL", "admin@lagarigo.local"),
		AdminPassword: env("ADMIN_PASSWORD", "admin123"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
