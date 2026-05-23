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

func TestClient_ListStates(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/state_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/states",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListStatesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewStateSort(utils.StateFieldName, false),
	}
	resp, err := svc.ListStates(context.Background(), req)
	if err != nil {
		t.Fatalf("listing states: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "state" {
		t.Errorf("request sort: expected=%s; got=%s", "state", query.Get("sort"))
	}

	// Check response
	if len(resp.States) != 2 {
		t.Errorf("response states list length: expected=%d; got=%d", 2, len(resp.States))
	}
}

func TestClient_SearchStates(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/state_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/states/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.SearchStatesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  utils.NewStateSort(utils.StateFieldID, false),
		Query: utils.NewStateQuery().Eq(utils.StateFieldID, 123),
	}
	resp, err := svc.SearchStates(context.Background(), req)
	if err != nil {
		t.Fatalf("searching states: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "state_id" {
		t.Errorf("request sort: expected=%s; got=%s", "state_id", query.Get("sort"))
	}
	if query.Get("q") != "state_id|eq|123" {
		t.Errorf("request query: expected=%s; got=%s", "state_id|eq|123", query.Get("q"))
	}

	// Check response
	if len(resp.States) != 2 {
		t.Errorf("response states list length: expected=%d; got=%d", 2, len(resp.States))
	}
}
