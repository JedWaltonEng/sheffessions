package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sheffessions/api/handlers"
	"sheffessions/api/store"
	"testing"
)

type FakePostStore struct {
	Confessions []*store.Confession
}

func (fps *FakePostStore) SaveConfession(content, source string) (int64, error) {
	return 0, nil // Not necessary for this test
}

func (fps *FakePostStore) RandomConfession() (*store.Confession, error) {
	// For simplicity, let's always return the first confession
	return fps.Confessions[0], nil
}

func (fps *FakePostStore) DeleteConfessionByID(id int64) error {
	return nil // Not necessary for this test
}

func TestPostGenerationHandler(t *testing.T) {
	// Create a fake confession for testing
	confession := &store.Confession{
		ID:                 1,
		ConfessionText:     "Test Confession",
		SourceOfConfession: "Test Source",
	}
	fakeStore := &FakePostStore{Confessions: []*store.Confession{confession}}

	handler := handlers.NewPostGenerationHandler(fakeStore)

	req, err := http.NewRequest("GET", "/generate-post", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content-type: got %v, want %v", contentType, expectedContentType)
	}

	// Deserialize the returned confession for easier inspection
	var returnedConfession store.Confession
	err = json.Unmarshal(rr.Body.Bytes(), &returnedConfession)
	if err != nil {
		t.Fatal("Failed to unmarshal returned confession:", err)
	}

	if returnedConfession.ConfessionText != confession.ConfessionText {
		t.Errorf("handler returned unexpected confession text: got %v, want %v", returnedConfession.ConfessionText, confession.ConfessionText)
	}

	// If there's a function to check post duplicity or any other function, invoke and check them
	// For now, we'll use the provided `isPostDuplicate` function as an example
	// if isPostDuplicate() {
	// 	t.Error("Post is detected as a duplicate when it shouldn't be.")
	// }
}
