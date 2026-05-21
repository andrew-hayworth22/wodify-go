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

func TestClient_ListClassSignIns(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/class_sign_in_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/sign-ins",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.ListClassSignInsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: leads.NewClassSignInSort(leads.ClassSignInFieldClassID, false),
	}
	resp, err := svc.ListClassSignIns(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing sign-ins: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "class_id" {
		t.Errorf("request sort: expected=%s; got=%s", "class_id", query.Get("sort"))
	}

	// Check response
	if len(resp.SignIns) != 2 {
		t.Errorf("response sign-ins list length: expected=%d; got=%d", 2, len(resp.SignIns))
	}
}

func TestClient_SearchClassSignIns(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/class_sign_in_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/sign-ins/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.SearchClassSignInsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  leads.NewClassSignInSort(leads.ClassSignInFieldID, false),
		Query: leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldIsDropIn, true),
	}
	resp, err := svc.SearchClassSignIns(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching sign-ins: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "id" {
		t.Errorf("request sort: expected=%s; got=%s", "id", query.Get("sort"))
	}
	if query.Get("q") != "is_drop_in|eq|true" {
		t.Errorf("request query: expected=%s; got=%s", "is_drop_in|eq|true", query.Get("q"))
	}

	// Check response
	if len(resp.SignIns) != 2 {
		t.Errorf("response sign-ins list length: expected=%d; got=%d", 2, len(resp.SignIns))
	}
}
