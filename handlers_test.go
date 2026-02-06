package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test the homeHandler
func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

// Test asciiArtHandler with valid POST data
func TestAsciiArtHandler_Valid(t *testing.T) {
	formData := strings.NewReader("text=hello&banner=standard")
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", formData)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	asciiArtHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

// Test asciiArtHandler with missing POST fields
func TestAsciiArtHandler_MissingInput(t *testing.T) {
	formData := strings.NewReader("text=&banner=")
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", formData)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	asciiArtHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}
}

// Test asciiArtHandler with GET method (not allowed)
func TestAsciiArtHandler_GetMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	w := httptest.NewRecorder()

	asciiArtHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for GET, got %d", resp.StatusCode)
	}
}
