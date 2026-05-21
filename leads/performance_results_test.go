package leads_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/models"
)

func TestClient_ListPerformanceResults(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/performance_result_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/performance-results",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.ListPerformanceResultsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}
	resp, err := svc.ListPerformanceResults(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing performance results: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}

	// Check response
	if len(resp.PerformanceResults) != 2 {
		t.Errorf("response performance result list length: expected=%d; got=%d", 2, len(resp.PerformanceResults))
	}
	if len(resp.PerformanceResults[0].ChildResults) != 2 {
		t.Errorf("response child results list length: expected=%d; got=%d", 2, len(resp.PerformanceResults[0].ChildResults))
	}
}

func TestClient_ListPerformanceResultsByComponent(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/performance_result_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/performance-results/components/123",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.ListPerformanceResultsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}
	resp, err := svc.ListPerformanceResultsByComponent(context.Background(), 123, 123, req)
	if err != nil {
		t.Fatalf("listing performance results by component: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}

	// Check response
	if len(resp.PerformanceResults) != 2 {
		t.Errorf("response performance result list length: expected=%d; got=%d", 2, len(resp.PerformanceResults))
	}
	if len(resp.PerformanceResults[0].ChildResults) != 2 {
		t.Errorf("response child results list length: expected=%d; got=%d", 2, len(resp.PerformanceResults[0].ChildResults))
	}
}
