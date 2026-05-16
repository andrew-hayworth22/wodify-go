package httpclient_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

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
