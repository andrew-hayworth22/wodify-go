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

func TestClient_ListBookings(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/booking_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/appointments/bookings",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.ListBookingsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: leads.NewBookingSort(leads.BookingFieldID, false),
	}
	resp, err := svc.ListBookings(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing bookings: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "booking_id" {
		t.Errorf("request sort: expected=%s; got=%s", "booking_id", query.Get("sort"))
	}

	// Check response
	if len(resp.Bookings) != 2 {
		t.Errorf("response bookings list length: expected=%d; got=%d", 2, len(resp.Bookings))
	}
}

func TestClient_SearchBookings(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/booking_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/appointments/bookings/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.SearchBookingsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  leads.NewBookingSort(leads.BookingFieldID, false),
		Query: leads.NewBookingQuery().Eq(leads.BookingFieldID, 123),
	}
	resp, err := svc.SearchBookings(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching bookings: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "booking_id" {
		t.Errorf("request sort: expected=%s; got=%s", "booking_id", query.Get("sort"))
	}
	if query.Get("q") != "booking_id|eq|123" {
		t.Errorf("request query: expected=%s; got=%s", "booking_id|eq|123", query.Get("q"))
	}

	// Check response
	if len(resp.Bookings) != 2 {
		t.Errorf("response bookings list length: expected=%d; got=%d", 2, len(resp.Bookings))
	}
}
