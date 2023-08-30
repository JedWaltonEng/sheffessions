package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sheffessions/api/services"
)

type Confession struct {
	Content            string `json:"content"`
	SourceOfConfession string `json:"source_of_confession"`
}

func HandleConfessions(s services.ConfessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

			affected, err := s.SaveConfession(confession.Content, confession.SourceOfConfession)
			if err != nil {
				log.Println("Failed to insert confession to database:", err)
				http.Error(w, "Failed to save confession", http.StatusInternalServerError)
				return
			}

			log.Printf("Inserted confession into database. Rows affected: %d", affected)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Confession received"))
			log.Println("Confession received: ", confession)

		} else {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	}
}
