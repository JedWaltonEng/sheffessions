package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sheffessions/api/store"
)

type Confession struct {
	Content            string `json:"content"`
	SourceOfConfession string `json:"source_of_confession"`
}

func HandleConfessions(w http.ResponseWriter, r *http.Request) {
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
		// Save to PostgreSQL
		result, err := store.DB.Exec("INSERT INTO confessions (confession_text, source_of_confession) VALUES ($1, $2)", confession.Content, confession.SourceOfConfession)
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

	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
