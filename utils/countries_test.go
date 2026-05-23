package utils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListCountries(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/country_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/countries",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.CountryFieldName)
	req := utils.NewCountryListRequest(p, s)
	resp, err := svc.ListCountries(context.Background(), req)
	if err != nil {
		t.Fatalf("listing countries: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.Countries) != 2 {
		t.Errorf("response countries list length: expected=%d; got=%d", 2, len(resp.Countries))
	}
}

func TestClient_SearchCountries(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/country_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/countries/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.CountryFieldID)
	q := utils.NewCountryQuery().Eq(utils.CountryFieldID, 123)
	req := utils.NewCountrySearchRequest(p, s, q)
	resp, err := svc.SearchCountries(context.Background(), req)
	if err != nil {
		t.Fatalf("searching countries: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.Countries) != 2 {
		t.Errorf("response countries list length: expected=%d; got=%d", 2, len(resp.Countries))
	}
}
