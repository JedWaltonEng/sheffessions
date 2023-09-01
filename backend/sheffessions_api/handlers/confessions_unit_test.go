package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sheffessions/api/services"
	"sheffessions/api/store"
	"testing"
)

type FakeStore struct {
	Confessions []*store.Confession
}

func (fs *FakeStore) SaveConfession(content, source string) (int64, error) {
	confession := &store.Confession{
		ID:                 int64(len(fs.Confessions) + 1),
		ConfessionText:     content,
		SourceOfConfession: source,
	}
	fs.Confessions = append(fs.Confessions, confession)
	return confession.ID, nil
}

func (fs *FakeStore) RandomConfession() (*store.Confession, error) {
	// Just return the first confession for simplicity
	return fs.Confessions[0], nil
}

func (fs *FakeStore) DeleteConfessionByID(id int64) error {
	// Omitted for brevity
	return nil
}

func TestHandleConfessions(t *testing.T) {
	fakeStore := &FakeStore{}
	confessionService := services.NewConfessionService(fakeStore)
	handler := HandleConfessions(confessionService)

	confessionData := `{"content": "Test Confession", "source_of_confession": "Test Source"}`
	req, err := http.NewRequest("POST", "/confessions", bytes.NewBufferString(confessionData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Confession received"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

	// Verify that the confession was stored in our fake store
	if len(fakeStore.Confessions) != 1 {
		t.Fatalf("expected 1 confession in the store, got %d", len(fakeStore.Confessions))
	}
}
