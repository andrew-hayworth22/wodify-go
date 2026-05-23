package leads_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_ListPerformanceResults(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/performance_result_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/performance-results",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	req := leads.NewPerformanceResultListRequest(p)
	resp, err := svc.ListPerformanceResults(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing performance results: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)

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
	body := testutil.MustReadJSONFixture(t, "testdata/performance_result_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/performance-results/components/123",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	req := leads.NewPerformanceResultListRequest(p)
	resp, err := svc.ListPerformanceResultsByComponent(context.Background(), 123, 123, req)
	if err != nil {
		t.Fatalf("listing performance results by component: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)

	// Check response
	if len(resp.PerformanceResults) != 2 {
		t.Errorf("response performance result list length: expected=%d; got=%d", 2, len(resp.PerformanceResults))
	}
	if len(resp.PerformanceResults[0].ChildResults) != 2 {
		t.Errorf("response child results list length: expected=%d; got=%d", 2, len(resp.PerformanceResults[0].ChildResults))
	}
}
