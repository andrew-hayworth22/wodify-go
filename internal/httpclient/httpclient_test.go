package httpclient_test

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
)

type Request struct {
	Data string `json:"data"`
}

type Response struct {
	Status string `json:"status"`
}

func TestClient_Do_Success(t *testing.T) {
	t.Parallel()
	queryParams := url.Values{}
	queryParams.Add("data", "test")
	req := Request{
		Data: "hello world",
	}
	expectedResp := Response{
		Status: "OK",
	}
	svr := testutil.NewWodifyClient(t, "test-key", 0,
		testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusOK,
			testutil.WithResponseBody(expectedResp),
		),
	)

	var actualResp Response
	err := svr.Do(context.Background(), http.MethodPost, "/test", queryParams, req, &actualResp)
	if err != nil {
		t.Fatalf("do: %s", err)
	}

	if actualResp.Status != expectedResp.Status {
		t.Errorf("expect status %s, but got %s", expectedResp.Status, actualResp.Status)
	}
}

func TestClient_Do_Retry(t *testing.T) {
	t.Parallel()
	svr := testutil.NewWodifyClient(t, "test-key", 2,
		testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusOK,
			testutil.WithRateLimitingCount(2),
			testutil.WithExpectedCount(3),
		),
	)

	err := svr.Do(context.Background(), http.MethodPost, "/test", nil, nil, nil)
	if err != nil {
		t.Fatalf("do: %s", err)
	}
}

func TestClient_Do_Errors(t *testing.T) {
	t.Parallel()
	tc := []struct {
		name             string
		endpoint         *testutil.Endpoint
		expectedErr      wodify.APIError
		expectedSentinel error
	}{
		{
			name:     "validation error (400 code)",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusBadRequest),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusBadRequest,
				MoreInfo: "400 Bad Request",
			},
			expectedSentinel: httpclient.ErrBadRequest,
		},
		{
			name: "validation error (200 code)",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusOK,
				testutil.WithResponseBody(httpclient.APIError{
					HTTPCode:         http.StatusUnprocessableEntity,
					DeveloperMessage: "Developer Message",
					UserMessage:      "User Message",
					MoreInfo:         "More Info",
				}),
			),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusUnprocessableEntity,
				MoreInfo: "More Info",
			},
			expectedSentinel: httpclient.ErrBadRequest,
		},
		{
			name:     "not found error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusNotFound),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusNotFound,
				MoreInfo: "404 Not Found",
			},
			expectedSentinel: httpclient.ErrNotFound,
		},
		{
			name:     "unauthorized error (forbidden)",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusForbidden),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusForbidden,
				MoreInfo: "403 Forbidden",
			},
			expectedSentinel: httpclient.ErrUnauthorized,
		},
		{
			name:     "unauthorized error (unauthorized)",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusUnauthorized),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusUnauthorized,
				MoreInfo: "401 Unauthorized",
			},
			expectedSentinel: httpclient.ErrUnauthorized,
		},
		{
			name:     "error rate limited",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusTooManyRequests),
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusTooManyRequests,
				MoreInfo: "429 Too Many Requests",
			},
			expectedSentinel: httpclient.ErrRateLimited,
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			svr := testutil.NewWodifyClient(t, "test-key", 0, c.endpoint)

			err := svr.Do(context.Background(), http.MethodPost, "/test", nil, nil, c.expectedErr)
			if err == nil {
				t.Fatal("expected error")
			}

			var apiErr *httpclient.APIError
			if !errors.As(err, &apiErr) {
				t.Errorf("expected APIError, but got %v", err)
			}
			if apiErr.HTTPCode != c.expectedErr.HTTPCode {
				t.Errorf("expected HTTP code %d, but got %d", c.expectedErr.HTTPCode, apiErr.HTTPCode)
			}
			if apiErr.MoreInfo != c.expectedErr.MoreInfo {
				t.Errorf("expected moreInfo %s, but got %s", c.expectedErr.MoreInfo, apiErr.MoreInfo)
			}
			if !errors.Is(err, c.expectedSentinel) {
				t.Errorf("expected %v, but got %v", c.expectedSentinel, err)
			}
			if apiErr.Error() != c.expectedErr.Error() {
				t.Errorf("expected error message %s, but got %s", c.expectedErr.Error(), apiErr.Error())
			}
		})
	}
}

func TestClient_Do_InvalidRequestJSON(t *testing.T) {
	t.Parallel()
	svr := testutil.NewWodifyClient(t, "test-key", 0,
		testutil.NewEndpoint(t, http.MethodPost, "/test", http.StatusOK),
	)

	err := svr.Do(context.Background(), http.MethodPost, "/test", nil, testutil.MarshalError{}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClient_Do_InvalidResponseJSON(t *testing.T) {
	t.Parallel()
	svr := testutil.NewWodifyClient(t, "test-key", 0,
		testutil.NewEndpoint(t, http.MethodGet, "/test", http.StatusOK,
			testutil.WithRawResponseBody([]byte(`"not an object"`)),
		),
	)

	var out Response
	err := svr.Do(context.Background(), http.MethodGet, "/test", nil, nil, &out)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClient_Do_InvalidMethod(t *testing.T) {
	t.Parallel()
	svr := testutil.NewWodifyClient(t, "test-key", 0,
		testutil.NewEndpoint(t, http.MethodGet, "/test", http.StatusOK),
	)
	err := svr.Do(context.Background(), "INVALID METHOD", "/test", nil, nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClient_Do_TransportError(t *testing.T) {
	t.Parallel()
	client := httpclient.New(&http.Client{Transport: testutil.TransportError{}}, "http://example.com", "test-key", 0)
	err := client.Do(context.Background(), http.MethodGet, "/test", nil, nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClient_Do_TransportReadError(t *testing.T) {
	t.Parallel()
	client := httpclient.New(&http.Client{Transport: testutil.TransportReadError{}}, "http://example.com", "test-key", 0)
	var out Response
	err := client.Do(context.Background(), http.MethodGet, "/test", nil, nil, &out)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClient_Do_ContextCancelled(t *testing.T) {
	t.Parallel()
	svr := testutil.NewWodifyClient(t, "test-key", 10,
		testutil.NewEndpoint(t, http.MethodGet, "/test", http.StatusOK,
			testutil.WithRateLimitingCount(999),
			testutil.WithExpectedCount(10),
		),
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	err := svr.Do(ctx, http.MethodGet, "/test", nil, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}
