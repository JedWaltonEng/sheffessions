package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sheffessions/api/services"
)

func HandlePostGeneration(s services.PostGenerationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodGet {
			confession, err := s.GeneratePost()
			if err != nil {
				log.Printf("Failed to generate post: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			respBytes, err := json.Marshal(confession)
			if err != nil {
				log.Printf("Error marshaling confession: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(respBytes)

		} else {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	}
}
