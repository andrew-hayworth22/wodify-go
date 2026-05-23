package utils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListStates(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/state_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/states",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.StateFieldName)
	req := utils.NewStateListRequest(p, s)
	resp, err := svc.ListStates(context.Background(), req)
	if err != nil {
		t.Fatalf("listing states: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.States) != 2 {
		t.Errorf("response states list length: expected=%d; got=%d", 2, len(resp.States))
	}
}

func TestClient_SearchStates(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/state_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/states/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.StateFieldID)
	q := utils.NewStateQuery().Eq(utils.StateFieldID, 123)
	req := utils.NewStateSearchRequest(p, s, q)
	resp, err := svc.SearchStates(context.Background(), req)
	if err != nil {
		t.Fatalf("searching states: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.States) != 2 {
		t.Errorf("response states list length: expected=%d; got=%d", 2, len(resp.States))
	}
}
