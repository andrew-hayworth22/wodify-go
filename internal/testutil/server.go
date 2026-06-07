package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
)

// NewClient creates a new test server and returns a generic HTTP client
func NewClient(t *testing.T, endpoints ...*Endpoint) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()

	for _, endpoint := range endpoints {
		mux.HandleFunc(endpoint.Pattern(), endpoint.Handler())
	}

	mux.HandleFunc("/", UnexpectedCall(t))
	return httptest.NewServer(mux)
}

// NewWodifyClient creates a new test server and returns an internal Wodify API client
func NewWodifyClient(t *testing.T, apiKey string, maxRetries int, endpoints ...*Endpoint) *httpclient.Client {
	t.Helper()
	svr := NewClient(t, endpoints...)
	return httpclient.New(&http.Client{}, svr.URL, apiKey, maxRetries)
}

// UnexpectedCall is the default handler for the test server
func UnexpectedCall(t *testing.T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected call to %s %s", r.Method, r.URL)
	}
}
