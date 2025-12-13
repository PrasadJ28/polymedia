package configs

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string // [NEW] Stores "development" or "production"
	Server   serverConfig
	Database databaseConfig
    // Add Redis/NATS here later if needed
	Infrastructure infrastructureConfig
}

type serverConfig struct {
	Address string
}

type databaseConfig struct {
	DatabaseDriver string
	DatabaseSource string
}

type infrastructureConfig struct {
    NatsUrl     string
    MinioBucket string
    // ...
}

func NewConfig() *Config {
	// 1. Try to load .env, but DO NOT PANIC if it fails.
	// In Docker/Prod, this file won't exist, and that's okay because
	// variables are injected via docker-compose environment section.
	err := godotenv.Load("configs/dev.env")
	if err != nil {
		log.Println("Warning: .env file not found. Relying on System Env Variables.")
	}

	c := &Config{
		// [NEW] Load the Environment Type
		// We default to "development" if not set, for safety.
		AppEnv: getEnvOrDefault("APP_ENV", "development"),

		Server: serverConfig{
			Address: GetEnvOrPanic(constants.EnvKeys.ServerAddress),
		},
		Database: databaseConfig{
			DatabaseDriver: GetEnvOrPanic(constants.EnvKeys.DBDriver),
			DatabaseSource: GetEnvOrPanic(constants.EnvKeys.DBSource),
		},
	}

	return c
}

// Helper: Enforce required variables
func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("CRITICAL: environment variable %s not set", key))
	}
	return value
}

// Helper: Optional variables with fallback
func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func (conf *Config) CorsNew() gin.HandlerFunc {
	// Use getEnvOrDefault here just in case, or keep Panic if it's strict
	allowedOrigin := GetEnvOrPanic(constants.EnvKeys.CorsAllowedOrigin)

	return cors.New(cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{constants.Headers.Origin, "Content-Type", "Authorization"}, // Added common headers
		ExposeHeaders:    []string{constants.Headers.ContentLength},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowedOrigin
		},
		MaxAge: constants.MaxAge,
	})
}
