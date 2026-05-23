package leads_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_ListBookings(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/booking_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/appointments/bookings",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.BookingFieldID)
	req := leads.NewBookingListRequest(p, s)
	resp, err := svc.ListBookings(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing bookings: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.Bookings) != 2 {
		t.Errorf("response bookings list length: expected=%d; got=%d", 2, len(resp.Bookings))
	}
}

func TestClient_SearchBookings(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/booking_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/appointments/bookings/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.BookingFieldID)
	q := leads.NewBookingQuery().Eq(leads.BookingFieldID, 123)
	req := leads.NewBookingSearchRequest(p, s, q)
	resp, err := svc.SearchBookings(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching bookings: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.Bookings) != 2 {
		t.Errorf("response bookings list length: expected=%d; got=%d", 2, len(resp.Bookings))
	}
}
