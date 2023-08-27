package main

import (
	"log"
	"net/http"
	"os"
	"sheffessions/api/handlers"
	"sheffessions/api/middleware"
	"sheffessions/api/store"

	_ "github.com/lib/pq"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	store.InitDB()
	defer store.DB.Close()

	http.HandleFunc("/confessions", middleware.Logging(handlers.HandleConfessions))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Server starting on :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
