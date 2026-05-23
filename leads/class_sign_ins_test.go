package leads_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_ListClassSignIns(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/class_sign_in_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/sign-ins",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.ClassSignInFieldClassID)
	req := leads.NewClassSignInListRequest(p, s)
	resp, err := svc.ListClassSignIns(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing sign-ins: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.SignIns) != 2 {
		t.Errorf("response sign-ins list length: expected=%d; got=%d", 2, len(resp.SignIns))
	}
}

func TestClient_SearchClassSignIns(t *testing.T) {
	body := testutil.MustReadJSONFixture(t, "testdata/class_sign_in_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/sign-ins/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.ClassSignInFieldClassID)
	q := leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldIsDropIn, true)
	req := leads.NewClassSignInSearchRequest(p, s, q)
	resp, err := svc.SearchClassSignIns(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching sign-ins: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.SignIns) != 2 {
		t.Errorf("response sign-ins list length: expected=%d; got=%d", 2, len(resp.SignIns))
	}
}
