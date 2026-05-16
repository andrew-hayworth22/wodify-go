package httpclient_test

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

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
	hdl := &testutil.Handler{
		Method:     http.MethodPost,
		Path:       "/test",
		StatusCode: http.StatusOK,
		Body:       expectedResp,
	}
	_ = testutil.NewServer(t, hdl)
	client := httpclient.New(&http.Client{}, hdl.BaseURL, "test-key", 0)

	var actualResp Response
	err := client.Do(context.Background(), http.MethodPost, "/test", queryParams, req, &actualResp)
	if err != nil {
		t.Fatalf("do: %s", err)
	}

	if actualResp.Status != expectedResp.Status {
		t.Errorf("expect status %s, but got %s", expectedResp.Status, actualResp.Status)
	}
}

func TestClient_Do_Retry(t *testing.T) {
	t.Parallel()
	hdl := &testutil.Handler{
		Method:              http.MethodPost,
		Path:                "/test",
		StatusCode:          http.StatusOK,
		Body:                nil,
		TooManyRequestCount: 2,
	}
	_ = testutil.NewServer(t, hdl)
	client := httpclient.New(&http.Client{}, hdl.BaseURL, "test-key", 3)

	err := client.Do(context.Background(), http.MethodPost, "/test", nil, nil, nil)
	if err != nil {
		t.Fatalf("do: %s", err)
	}

	if hdl.CallCount != 3 {
		t.Errorf("expected call count %d; got %d", 3, hdl.CallCount)
	}
}

func TestClient_Do_Errors(t *testing.T) {
	tc := []struct {
		name             string
		hdl              *testutil.Handler
		expectedErr      wodify.APIError
		expectedSentinel error
	}{
		{
			name: "validation error (400 code)",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusBadRequest,
			},
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusBadRequest,
				MoreInfo: "400 Bad Request",
			},
			expectedSentinel: httpclient.ErrBadRequest,
		},
		{
			name: "validation error (200 code)",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusOK,
				Body: httpclient.APIError{
					HTTPCode:         http.StatusUnprocessableEntity,
					DeveloperMessage: "Developer Message",
					UserMessage:      "User Message",
					MoreInfo:         "More Info",
				},
			},
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusUnprocessableEntity,
				MoreInfo: "More Info",
			},
			expectedSentinel: httpclient.ErrBadRequest,
		},
		{
			name: "not found error",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusNotFound,
			},
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusNotFound,
				MoreInfo: "404 Not Found",
			},
			expectedSentinel: httpclient.ErrNotFound,
		},
		{
			name: "unauthorized error (forbidden)",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusForbidden,
			},
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusForbidden,
				MoreInfo: "403 Forbidden",
			},
			expectedSentinel: httpclient.ErrUnauthorized,
		},
		{
			name: "unauthorized error (unauthorized)",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusUnauthorized,
			},
			expectedErr: httpclient.APIError{
				HTTPCode: http.StatusUnauthorized,
				MoreInfo: "401 Unauthorized",
			},
			expectedSentinel: httpclient.ErrUnauthorized,
		},
		{
			name: "error rate limited",
			hdl: &testutil.Handler{
				Method:     http.MethodPost,
				Path:       "/test",
				StatusCode: http.StatusTooManyRequests,
			},
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

			// Create test server and client
			_ = testutil.NewServer(t, c.hdl)
			client := httpclient.New(&http.Client{}, c.hdl.BaseURL, "test-key", 0)

			// Make request and check response
			err := client.Do(context.Background(), http.MethodPost, "/test", nil, nil, c.expectedErr)
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
