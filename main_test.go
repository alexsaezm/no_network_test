package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServer is a function to test if the server is correctly serving the given string.
func TestServer(t *testing.T) {
	// Start a test server
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	// Perform a GET request to the server
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to perform GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(body))
	}
}
