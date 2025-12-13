package constants

import "time"

var EnvKeys = envKeys{
	// 1. App Environment (Dev/Prod)
	AppEnv: "APP_ENV",

	// 2. Server & Security
	ServerAddress:     "SERVER_ADDRESS",
	CorsAllowedOrigin: "CORS_ALLOWED_ORIGIN",

	// 3. Database (Postgres)
	DBDriver: "DB_DRIVER",
	DBSource: "DB_SOURCE",

	// 4. Infrastructure (The new additions)
	NatsUrl:       "NATS_URL",
	RedisHost:     "REDIS_HOST",
	MinioEndpoint: "MINIO_ENDPOINT",
	MinioBucket:   "MINIO_BUCKET",
}

var Headers = headers{
	Origin:        "Origin",
	ContentLength: "Content-Length",
	// [NEW] You need these for JSON APIs and Video Uploads
	ContentType:   "Content-Type",
	Authorization: "Authorization",
}

var MaxAge = 12 * time.Hour

// --- Struct Definitions ---

type envKeys struct {
	AppEnv            string
	ServerAddress     string
	CorsAllowedOrigin string
	DBDriver          string
	DBSource          string
	// Infrastructure
	NatsUrl       string
	RedisHost     string
	MinioEndpoint string
	MinioBucket   string
}

type headers struct {
	Origin        string
	ContentLength string
	ContentType   string
	Authorization string
}
