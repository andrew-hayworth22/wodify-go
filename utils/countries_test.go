package utils_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/models"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListCountries(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/country_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/countries",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListCountriesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewCountrySort(utils.CountryFieldName, false),
	}
	resp, err := svc.ListCountries(context.Background(), req)
	if err != nil {
		t.Fatalf("listing countries: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "country" {
		t.Errorf("request sort: expected=%s; got=%s", "country", query.Get("sort"))
	}

	// Check response
	if len(resp.Countries) != 2 {
		t.Errorf("response countries list length: expected=%d; got=%d", 2, len(resp.Countries))
	}
}

func TestClient_SearchCountries(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/country_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/countries/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.SearchCountriesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  utils.NewCountrySort(utils.CountryFieldID, false),
		Query: utils.NewCountryQuery().Eq(utils.CountryFieldID, 123),
	}
	resp, err := svc.SearchCountries(context.Background(), req)
	if err != nil {
		t.Fatalf("searching countries: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "country_id" {
		t.Errorf("request sort: expected=%s; got=%s", "country_id", query.Get("sort"))
	}
	if query.Get("q") != "country_id|eq|123" {
		t.Errorf("request query: expected=%s; got=%s", "country_id|eq|123", query.Get("q"))
	}

	// Check response
	if len(resp.Countries) != 2 {
		t.Errorf("response countries list length: expected=%d; got=%d", 2, len(resp.Countries))
	}
}
