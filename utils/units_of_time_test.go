package utils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListUnitsOfTime(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/unit_of_time_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/units-of-time",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.UnitOfTimeFieldNameSingular)
	req := utils.NewUnitOfTimeListRequest(p, s)
	resp, err := svc.ListUnitsOfTime(context.Background(), req)
	if err != nil {
		t.Fatalf("listing units of time: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.UnitsOfTime) != 2 {
		t.Errorf("response units of time list length: expected=%d; got=%d", 2, len(resp.UnitsOfTime))
	}
}

func TestClient_SearchUnitsOfTime(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/unit_of_time_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/units-of-time/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.UnitOfTimeFieldNameSingular)
	q := utils.NewUnitOfTimeQuery().Eq(utils.UnitOfTimeFieldID, 123)
	req := utils.NewUnitOfTimeSearchRequest(p, s, q)
	resp, err := svc.SearchUnitsOfTime(context.Background(), req)
	if err != nil {
		t.Fatalf("searching units of time: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.UnitsOfTime) != 2 {
		t.Errorf("response units of time list length: expected=%d; got=%d", 2, len(resp.UnitsOfTime))
	}
}
