package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sheffessions/api/handlers"
	"sheffessions/api/services"
	"sheffessions/api/store"
	"testing"

	// Other imports...
	_ "github.com/lib/pq"
)

func TestHandleConfessionsIntegration(t *testing.T) {
	// Setup
	// Assuming you have some utility functions to start your test database
	confessionStore := store.InitDB()
	// defer teardownTestDatabase(db)

	// Create instances of store, service, and handler
	confessionService := services.NewConfessionService(confessionStore)
	handler := handlers.HandleConfessions(confessionService)

	// Test
	confessionData := `{"content": "Integration Test Confession", "source_of_confession": "Integration Test Source"}`
	req, err := http.NewRequest("POST", "/confessions", bytes.NewBufferString(confessionData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code and response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Confession received"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

	// Teardown (handled by defer above)
}
