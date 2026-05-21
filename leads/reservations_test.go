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

func TestClient_ListReservations(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/reservation_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/reservations",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.ListReservationsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: leads.NewReservationSort(leads.ReservationFieldClassID, false),
	}
	resp, err := svc.ListReservations(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("listing reservations: %v", err)
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
	if len(resp.Reservations) != 2 {
		t.Errorf("response reservations list length: expected=%d; got=%d", 2, len(resp.Reservations))
	}
}

func TestClient_SearchReservations(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/reservation_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/123/classes/reservations/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.SearchReservationsRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  leads.NewReservationSort(leads.ReservationFieldID, false),
		Query: leads.NewReservationQuery().Eq(leads.ReservationFieldIsLateCancellation, true),
	}
	resp, err := svc.SearchReservations(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("searching reservations: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "reservation_id" {
		t.Errorf("request sort: expected=%s; got=%s", "reservation_id", query.Get("sort"))
	}
	if query.Get("q") != "is_late_cancellation|eq|true" {
		t.Errorf("request query: expected=%s; got=%s", "is_late_cancellation|eq|true", query.Get("q"))
	}

	// Check response
	if len(resp.Reservations) != 2 {
		t.Errorf("response reservation list length: expected=%d; got=%d", 2, len(resp.Reservations))
	}
}
