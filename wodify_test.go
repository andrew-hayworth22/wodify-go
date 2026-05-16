package wodify

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
)

func TestNew(t *testing.T) {
	// Set up test server
	hdl := testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/test",
		StatusCode: http.StatusOK,
	}
	_ = testutil.NewServer(t, &hdl)

	// Create client
	key := "test-key"
	wc, err := New(
		WithAPIKey(key),
		WithBaseURL(hdl.BaseURL),
		WithHTTPClient(&http.Client{}),
		WithTimeout(time.Minute),
		WithMaxRetries(2),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Send request
	err = wc.httpClient.Do(context.Background(), hdl.Method, hdl.Path, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Check API key
	actualKey := hdl.Request.Header.Get("X-Api-Key")
	if actualKey != key {
		t.Errorf("got X-Api-Key %q; want %q\n", actualKey, key)
	}
}

func TestNew_NoAPIKey(t *testing.T) {
	_, err := New()
	if err == nil {
		t.Fatal("expected a missing API key error")
	}
}

func TestNew_EnvironmentDefault(t *testing.T) {
	// Set environment variables
	_ = os.Setenv("WODIFY_API_KEY", "test-key")
	_ = os.Setenv("WODIFY_BASE_URL", "http://localhost")
	_ = os.Setenv("WODIFY_MAX_RETRIES", "0")

	// Create client and assert no errors occur (should default API key to environment variable)
	_, err := New()
	if err != nil {
		t.Fatal("expected environment variable WODIFY_API_KEY to be set")
	}
}
