package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connectionString := os.Getenv("POSTGRESQL_URL")
	if connectionString == "" {
		log.Fatal("POSTGRESQL_URL environment variable is not set")
	}
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() // This will check if the connection is successful
	if err != nil {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		log.Println("Handled the request.")
	}
}

type Confession struct {
	Content            string `json:"content"`
	SourceOfConfession string `json:"source_of_confession"` // added this line
}

var confessions []Confession

func handleConfessions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Handle preflight request (for CORS)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var confession Confession
		if err := json.NewDecoder(r.Body).Decode(&confession); err != nil {
			http.Error(w, "Failed to decode confession", http.StatusBadRequest)
			return
		}
		confessions = append(confessions, confession)
		// Save to PostgreSQL
		result, err := db.Exec("INSERT INTO confessions (confession_text, source_of_confession) VALUES ($1, $2)", confession.Content, confession.SourceOfConfession)
		if err != nil {
			log.Println("Failed to insert confession to database:", err)
			http.Error(w, "Failed to save confession", http.StatusInternalServerError)
			return
		}
		affected, _ := result.RowsAffected()
		log.Printf("Inserted confession into database. Rows affected: %d", affected)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Confession received"))
		log.Println("Confession received: ", confession)

	} else if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(confessions)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/confessions", loggingMiddleware(handleConfessions))

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
