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

	dbStore := store.ConnectToPostgres()

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
	// Confessions Route
	mux.HandleFunc("/confessions", middleware.Chain(
		handlers.HandleConfessions(services.NewConfessionService(config.DBStore)),
		middleware.CORSMiddleware(config.AllowedOrigins),
		middleware.Logging,
	))

	// Post Generation Route
	mux.HandleFunc("/cron-job", middleware.Chain(
		handlers.HandlePostGeneration(services.NewPostGenerationService(config.DBStore, config.DBStore)),
		middleware.Logging,
		middleware.TokenAuthMiddleware(config.SecretToken),
	))

	return mux
}
