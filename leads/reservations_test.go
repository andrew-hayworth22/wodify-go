package leads_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_ListReservations(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/reservation_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/reservations",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.ReservationFieldClassID)
	req := leads.NewReservationListRequest(p, s)
	resp, err := svc.ListReservations(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing reservations: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.Reservations) != 2 {
		t.Errorf("response reservations list length: expected=%d; got=%d", 2, len(resp.Reservations))
	}
}

func TestClient_SearchReservations(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/reservation_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/reservations/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.ReservationFieldID)
	q := leads.NewReservationQuery().Eq(leads.ReservationFieldIsLateCancellation, true)
	req := leads.NewReservationSearchRequest(p, s, q)
	resp, err := svc.SearchReservations(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching reservations: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.Reservations) != 2 {
		t.Errorf("response reservation list length: expected=%d; got=%d", 2, len(resp.Reservations))
	}
}
