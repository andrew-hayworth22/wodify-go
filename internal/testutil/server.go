// Package testutil provides test utilities
package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
)

// Handler is a test HTTP endpoint call.
type Handler struct {
	// Method is the expected HTTP method that will be called.
	Method string
	// Path is the expected path that will be called.
	Path string
	// StatusCode is the HTTP status code to return.
	StatusCode int
	// Body is the response body to return.
	Body any
	// Request is the request that was received.
	Request *http.Request
	// RequestBody is the body of the request that was received
	RequestBody []byte
}

// NewServer spins up a new test HTTP server that mocks the Wodify API and returns a client configured to use it.
func NewServer(t *testing.T, h *Handler) *httpclient.Client {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store request in handler
		h.Request = r
		var err error
		h.RequestBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("reading request body: %v", err)
		}

		// Check path and method
		if h.Method != r.Method {
			t.Errorf("expected method %q, got %q", h.Method, r.Method)
		}
		if h.Path != r.URL.Path {
			t.Errorf("expected path %q, got %q", h.Path, r.URL.Path)
		}

		// Write test response body
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(h.StatusCode)
		if err := json.NewEncoder(w).Encode(h.Body); err != nil {
			t.Fatalf("failed to encode test body: %v", err)
		}
	}))

	t.Cleanup(srv.Close)
	return httpclient.New(&http.Client{}, srv.URL, "test-key", 0)
}
