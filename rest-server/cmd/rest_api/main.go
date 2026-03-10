package main

import (
	"time"

	serve "github.com/PrasadJ28/gin-rest-server/api/server"
	routes "github.com/PrasadJ28/gin-rest-server/api/server/router"
	"github.com/PrasadJ28/gin-rest-server/configs"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/repository"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/database"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/handlers"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/services"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/storage"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config := configs.NewConfig()

	client, err := database.NewSQLClient(database.Config{
		DBDriver:          config.Database.DatabaseDriver,
		DBSource:          config.Database.DatabaseSource,
		MaxOpenConns:      25,
		MaxIdleConns:      25,
		ConnMaxIdleTime:   15 * time.Minute,
		ConnectionTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database client")
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Msgf("Failed to close database client: %v", err)
		}
	}()

	minioStorage, err := storage.NewMinioStorage(
		config.Infrastructure.MinioInternalEndpoint,
		config.Infrastructure.MinioExternalEndpoint,
		config.Infrastructure.MinioAccessKey,
		config.Infrastructure.MinioSecretKey,
		config.Infrastructure.MinioBucket,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize Minio storage")
		return
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(client.DB)

	//Initialize services
	userService := services.NewUserService(userRepo)

	// Pass services to handlers
	userHandler := handlers.NewUserHandler(userService)

	uploadService := services.NewUploadService(minioStorage)
	uploadHandler := handlers.NewUploadHandler(uploadService)

	cors := config.CorsNew()

	router := gin.Default()
	router.Use(cors)

	// Register routes
	routes.RegisterPublicEndpoints(router, userHandler)
	routes.RegisterUploadEndpoints(router, uploadHandler)

	server := serve.NewServer(log.Logger, router, config)
	server.Serve()
}
