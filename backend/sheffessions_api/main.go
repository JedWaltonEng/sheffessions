package main

import (
	"log"
	"net/http"
	"os"
	"sheffessions/api/handlers"
	"sheffessions/api/middleware"
	"sheffessions/api/services"
	"sheffessions/api/store"

	_ "github.com/lib/pq"
)

type ServerConfig struct {
	SecretToken    string
	AllowedOrigins []string
	Port           string
	DBStore        *store.DBStore
}

func NewServerConfig() *ServerConfig {
	secretToken := os.Getenv("SECRET_TOKEN")
	if secretToken == "" {
		log.Fatal("SECRET_TOKEN environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	allowedOrigins := []string{
		"http://localhost:3000",
		"https://sheffessions-fe-s3vpbhlkuq-ew.a.run.app",
		"https://sheffessions-fe-staging-5pu6yezepq-ew.a.run.app",
	}

	dbStore := store.InitDB()

	return &ServerConfig{
		SecretToken:    secretToken,
		AllowedOrigins: allowedOrigins,
		Port:           port,
		DBStore:        dbStore,
	}
}

func main() {
	config := NewServerConfig()
	defer config.DBStore.Close()

	mux := http.NewServeMux()
	mux = SetupRoutes(config, mux)

	log.Printf("Server starting on :%s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func SetupRoutes(config *ServerConfig, mux *http.ServeMux) *http.ServeMux {
	confessionsService := services.NewConfessionService(config.DBStore)
	handleConfessionsFunc := handlers.HandleConfessions(confessionsService)
	confessionsHandler := middleware.Chain(handleConfessionsFunc, middleware.CORSMiddleware(config.AllowedOrigins), middleware.Logging)

	postGenHandler := handlers.NewPostGenerationHandler(config.DBStore)
	postGenChainedHandler := middleware.Chain(
		postGenHandler.ServeHTTP,
		middleware.TokenAuthMiddleware(config.SecretToken),
		middleware.Logging,
	)

	// Register the endpoint with the chained handler
	mux.HandleFunc("/confessions", confessionsHandler)
	mux.HandleFunc("/cron-job", postGenChainedHandler)

	return mux
}
