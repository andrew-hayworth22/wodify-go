package leads_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_ListStatuses(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/status_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/statuses",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.StatusFieldName)
	req := leads.NewStatusListRequest(p, s)
	resp, err := svc.ListStatuses(context.Background(), req)
	if err != nil {
		t.Fatalf("listing lead statuses: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if resp.Pagination.Page != req.Page.Page {
		t.Errorf("response page: expected=%d; got=%d", req.Page.Page, resp.Pagination.Page)
	}
	if resp.Pagination.PageSize != req.Page.PageSize {
		t.Errorf("response page size: expected=%d; got=%d", req.Page.PageSize, resp.Pagination.PageSize)
	}
	if len(resp.Statuses) != 3 {
		t.Errorf("response lead statuses list length: expected=%d; got=%d", 3, len(resp.Statuses))
	}
}
