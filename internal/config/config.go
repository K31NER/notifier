package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV             string
	PORT                string
	GMAIL_CLIENT_ID     string
	GMAIL_CLIENT_SECRET string
	GMAIL_REFRESH_TOKEN string
	GMAIL_SENDER        string
}

// Definimos el loader para cargar la estructura de varibales que usaremos
func Load() *Config {
	_ = godotenv.Load() // Cargamos las variables de entorno

	return &Config{
		APP_ENV:             getEnv("APP_ENV", "development"),
		PORT:                getEnv("PORT", "8080"),
		GMAIL_CLIENT_ID:     getEnv("GMAIL_CLIENT_ID", ""),
		GMAIL_CLIENT_SECRET: getEnv("GMAIL_CLIENT_SECRET", ""),
		GMAIL_REFRESH_TOKEN: getEnv("GMAIL_REFRESH_TOKEN", ""),
		GMAIL_SENDER:        getEnv("GMAIL_SENDER", ""),
	}
}

// Definimos la funcion para leer cada env
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}