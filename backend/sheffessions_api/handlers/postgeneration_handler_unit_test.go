package handlers

import (
	"net/http"
	"net/http/httptest"
	"sheffessions/api/store"
	"testing"
)

type StubPostGenerationService struct {
	Called bool
}

func (s *StubPostGenerationService) GeneratePost() (*store.Confession, error) {
	s.Called = true
	return nil, nil
}

func (s *StubPostGenerationService) IsPostDuplicate() bool {
	return false // or true, whatever makes sense for your test
}

func TestHandlePostGeneration(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/cron-job", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	stubService := &StubPostGenerationService{}
	handler := http.HandlerFunc(HandlePostGeneration(stubService))
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
