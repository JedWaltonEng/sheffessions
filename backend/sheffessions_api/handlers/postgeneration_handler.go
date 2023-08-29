package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sheffessions/api/services"
	"sheffessions/api/store"
)

type PostGenerationHandler struct {
	service services.PostGenerationService
}

func NewPostGenerationHandler(db store.StorerConfessions) *PostGenerationHandler {
	service := services.NewPostGenerationService(db)
	return &PostGenerationHandler{service: service}
}

func (h *PostGenerationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	confession, err := h.service.GeneratePost()
	if err != nil {
		log.Printf("%v", err)
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
}
