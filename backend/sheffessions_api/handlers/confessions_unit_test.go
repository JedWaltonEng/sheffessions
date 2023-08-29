package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// In handlers_test.go
type FakeStore struct {
	saved bool
}

func (fs *FakeStore) SaveConfession(content, source string) (int64, error) {
	fs.saved = true
	return 1, nil // always pretend to successfully save one row
}

func TestHandleConfessions(t *testing.T) {
	fakeStore := &FakeStore{}

	handler := HandleConfessions(fakeStore)

	// Example test: Handle a POST request with valid data
	confessionData := `{"content": "Test Confession", "source_of_confession": "Test Source"}`
	req, err := http.NewRequest("POST", "/confessions", bytes.NewBufferString(confessionData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check response body
	expected := "Confession received"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

	// Check that the fake store recorded a save
	if !fakeStore.saved {
		t.Errorf("handler did not save confession using the Storer interface")
	}
}
