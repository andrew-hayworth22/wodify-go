package wodify

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go/internal/sort"
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

func TestNewPaginationRequest(t *testing.T) {
	req := NewPaginationRequest(1, 10)
	if req.Page != 1 {
		t.Errorf("expected page=1; got=%d", req.Page)
	}
	if req.PageSize != 10 {
		t.Errorf("expected page_size=10; got=%d", req.PageSize)
	}
}

func TestSort(t *testing.T) {
	tc := []struct {
		name     string
		actual   sort.Sort[string]
		expected sort.Sort[string]
	}{
		{
			name:     "ascending",
			actual:   SortAscending("test"),
			expected: sort.NewSort("test", false),
		},
		{
			name:     "descending",
			actual:   SortDescending("test"),
			expected: sort.NewSort("test", true),
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			if c.actual != c.expected {
				t.Errorf("expected=%v; got=%v", c.expected, c.actual)
			}
		})
	}
}
