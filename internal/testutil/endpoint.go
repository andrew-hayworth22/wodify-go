package testutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
)

// Endpoint represents an endpoint that will be configured on the test server
type Endpoint struct {
	// Testing context
	t *testing.T

	// Method of the endpoint
	method string
	// Path of the endpoint
	path string
	// StatusCode of the endpoint's response
	statusCode int
	// Body of the endpoint's response
	responseBody any
	// Raw bytes of endpoint's response used to test malformed JSON responses
	rawResponseBody []byte

	// Expected request body
	expectedRequestBody any
	// Expected request URL parameters
	expectedRequestURLParams url.Values
	// Expected request header
	expectedRequestHeader http.Header
	// Optional number of times the endpoint is expected to be called
	expectedCallCount int
	// Optional rate limiting callCount of the number of times the endpoint will return http.StatusTooManyRequests
	rateLimitingCount int
	// Optional duration to wait in the endpoint's handler
	wait time.Duration

	// Mutex for concurrent access
	mutex sync.RWMutex

	// Number of times the handler has been invoked
	callCount int
}

// EndpointOption represents an option for an endpoint
type EndpointOption func(*Endpoint)

// NewEndpoint creates a new endpoint
func NewEndpoint(t *testing.T, method, path string, statusCode int, opts ...EndpointOption) *Endpoint {
	t.Helper()

	e := &Endpoint{
		t:          t,
		method:     method,
		path:       path,
		statusCode: statusCode,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithResponseBody defines the body of the response from the endpoint
func WithResponseBody(body any) EndpointOption {
	return func(e *Endpoint) {
		e.responseBody = body
	}
}

// WithRawResponseBody defines a raw byte slice response from the endpoint
// Used to test malformed JSON responses
func WithRawResponseBody(rawBody []byte) EndpointOption {
	return func(e *Endpoint) {
		e.rawResponseBody = rawBody
	}
}

// WithExpectedCount defines the number of times an endpoint is expected to be called
func WithExpectedCount(count int) EndpointOption {
	return func(e *Endpoint) {
		e.expectedCallCount = count
	}
}

// WithExpectedRequestBody defines an expected request structure in the handler
func WithExpectedRequestBody(body any) EndpointOption {
	return func(e *Endpoint) {
		e.expectedRequestBody = body
	}
}

// WithExpectedRequestHeader defines expected headers to the request
func WithExpectedRequestHeader(header http.Header) EndpointOption {
	return func(e *Endpoint) {
		e.expectedRequestHeader = header
	}
}

// WithExpectedRequestURLValues defines expected URL parameters passed with the request
func WithExpectedRequestURLValues(params url.Values) EndpointOption {
	return func(e *Endpoint) {
		e.expectedRequestURLParams = params
	}
}

// WithExpectedRequestPagination adds expected URL values related to a pagination request
func WithExpectedRequestPagination(expected request.PaginationRequest) EndpointOption {
	return func(e *Endpoint) {
		if e.expectedRequestURLParams == nil {
			e.expectedRequestURLParams = url.Values{}
		}
		for key, val := range expected.ToQuery() {
			e.expectedRequestURLParams[key] = val
		}
	}
}

// WithExpectedRequestSort adds expected headers related to a sort request
func WithExpectedRequestSort[T ~string](expected sort.Sort[T]) EndpointOption {
	return func(e *Endpoint) {
		if e.expectedRequestURLParams == nil {
			e.expectedRequestURLParams = url.Values{}
		}
		e.expectedRequestURLParams["sort"] = []string{expected.String()}
	}
}

// WithExpectedRequestQuery adds expected headers related to a query request
func WithExpectedRequestQuery[T ~string](expected *query.Builder[T]) EndpointOption {
	return func(e *Endpoint) {
		if e.expectedRequestURLParams == nil {
			e.expectedRequestURLParams = url.Values{}
		}
		e.expectedRequestURLParams["q"] = []string{expected.String()}
	}
}

// WithRateLimitingCount defines the number of times an endpoint will return http.StatusTooManyRequests
// before returning the configured StatusCode and Body
func WithRateLimitingCount(count int) EndpointOption {
	return func(e *Endpoint) {
		e.rateLimitingCount = count
	}
}

// WithWait defines a duration that each call should take
func WithWait(wait time.Duration) EndpointOption {
	return func(e *Endpoint) {
		e.wait = wait
	}
}

// Pattern returns the HTTP pattern of the endpoint
func (e *Endpoint) Pattern() string {
	return fmt.Sprintf("%s %s", e.method, e.path)
}

// Handler creates the HTTP handler for the endpoint
func (e *Endpoint) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e.mutex.Lock()
		defer e.mutex.Unlock()

		e.callCount += 1

		if e.expectedCallCount != 0 && e.callCount > e.expectedCallCount {
			e.t.Fatalf("expected %d calls to %s, got %d", e.expectedCallCount, e.Pattern(), e.callCount)
		}

		if e.rateLimitingCount > 0 && e.callCount <= e.rateLimitingCount {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		if e.expectedRequestBody != nil {
			expected, err := json.Marshal(e.expectedRequestBody)
			if err != nil {
				e.t.Fatal(err)
			}
			actual, err := io.ReadAll(r.Body)
			if err != nil {
				e.t.Fatalf("failed to read request body: %v", err)
			}
			AssertJSONEqual(e.t, expected, actual)
		}

		if e.expectedRequestURLParams != nil {
			AssertURLValuesEqual(e.t, e.expectedRequestURLParams, r.URL.Query())
		}

		if e.expectedRequestHeader != nil {
			for key, values := range e.expectedRequestHeader {
				if !reflect.DeepEqual(r.Header[key], values) {
					e.t.Errorf("expected request header: %v; got %v", values, r.Header[key])
				}
			}
		}

		if e.wait > 0 {
			time.Sleep(e.wait)
		}

		// Write test response body
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(e.statusCode)
		if e.rawResponseBody != nil {
			_, err := w.Write(e.rawResponseBody)
			if err != nil {
				e.t.Fatalf("failed to write test body: %v", err)
			}
			return
		}
		if err := json.NewEncoder(w).Encode(e.responseBody); err != nil {
			e.t.Fatalf("failed to encode test body: %v", err)
		}
	}
}
