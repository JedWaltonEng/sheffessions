package integration

import (
	"io"
	"net/http"
	"testing"
)

func TestCronJobEndpoint(t *testing.T) {
	// Define the target endpoint.
	url := "http://localhost:8080/cron-job"

	// Create a new request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Add headers.
	req.Header.Set("Authorization", "securetoken")

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make a request: %v", err)
	}
	defer resp.Body.Close()

	// Check for expected response status.
	// if resp.StatusCode != http.StatusOK {
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status OK but got %v. Response body: %s", resp.Status, string(bodyBytes))
	}

	// Additional tests can be added here:
	// - Check the body of the response.
	// - Check against expected mock data.
	// - Other business logic validations.
}
