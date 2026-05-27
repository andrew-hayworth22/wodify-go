package leads_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/models"
)

func TestClient_Get(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead.json")

	// Create mock server and client
	svr := testutil.NewServer(t, &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/12345",
		StatusCode: http.StatusOK,
		Body:       body,
	})
	svc := leads.New(svr)

	// Make request
	resp, err := svc.Get(context.Background(), 12345)
	if err != nil {
		t.Fatalf("getting resp: %v", err)
	}

	// Check response
	if resp.ID != 12345 {
		t.Errorf("resp ID: expected=%d; got=%d", 12345, resp.ID)
	}

	if resp.FirstName != "John" {
		t.Errorf("resp first name: expected=%s; got=%s", "John", resp.FirstName)
	}

	if resp.LastName != "Doe" {
		t.Errorf("resp last name: expected=%s; got=%s", "Doe", resp.LastName)
	}

	expectedDateOfBirth := models.NewDate(time.Date(2001, time.December, 31, 0, 0, 0, 0, time.UTC))
	if resp.DateOfBirth != expectedDateOfBirth {
		t.Errorf("resp date of birth: expected=%s; got=%s", expectedDateOfBirth, resp.DateOfBirth)
	}

	expectednextClassReservation := models.NewDateTime(time.Date(2014, time.December, 31, 23, 59, 59, 938_000_000, time.UTC))
	if resp.NextClassReservation != expectednextClassReservation {
		t.Errorf("resp next class reservation: expected=%s; got=%s", expectednextClassReservation, resp.NextClassReservation)
	}
}

func TestClient_List(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(leads.LeadFieldFirstName)
	req := leads.NewLeadListRequest(p, s)
	resp, err := svc.List(context.Background(), req)
	if err != nil {
		t.Fatalf("listing leads: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if resp.Pagination.Page != req.Page.Page {
		t.Errorf("response page: expected=%d; got=%d", req.Page.Page, resp.Pagination.Page)
	}
	if resp.Pagination.PageSize != req.Page.PageSize {
		t.Errorf("response page size: expected=%d; got=%d", req.Page.PageSize, resp.Pagination.PageSize)
	}
	if len(resp.Leads) != 2 {
		t.Errorf("response leads list length: expected=%d; got=%d", 2, len(resp.Leads))
	}
}

func TestClient_Search(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request

	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortDescending(leads.LeadFieldFirstName)
	q := leads.NewLeadQuery().Eq(leads.LeadFieldFirstName, "john")
	req := leads.NewLeadSearchRequest(p, s, q)
	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		t.Fatalf("searching leads: %v", err)
	}

	// Check query parameters
	testutil.AssertPaginationParams(t, hdl.Request.URL.Query(), p)
	testutil.AssertSortParam(t, hdl.Request.URL.Query(), s)
	testutil.AssertQueryParam(t, hdl.Request.URL.Query(), q)

	// Check response
	if resp.Pagination.Page != req.Page.Page {
		t.Errorf("response page: expected=%d; got=%d", req.Page.Page, resp.Pagination.Page)
	}
	if resp.Pagination.PageSize != req.Page.PageSize {
		t.Errorf("response page size: expected=%d; got=%d", req.Page.PageSize, resp.Pagination.PageSize)
	}
	if len(resp.Leads) != 2 {
		t.Errorf("response leads list length: expected=%d; got=%d", 2, len(resp.Leads))
	}
}

func TestClient_Create(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead.json")

	// Create mock server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPost,
		Path:       "/leads",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.LeadCreateRequest{
		FirstName:  "John",
		LastName:   "Doe",
		LocationID: 2998,
	}
	resp, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("creating resp: %v", err)
	}

	// Check sent request
	var sentRequest leads.LeadCreateRequest
	if err := json.Unmarshal(hdl.RequestBody, &sentRequest); err != nil {
		t.Fatalf("decoding request: %v", err)
	}
	if sentRequest.FirstName != req.FirstName {
		t.Errorf("request first name: expected=%s; got=%s", req.FirstName, sentRequest.FirstName)
	}
	if sentRequest.LastName != req.LastName {
		t.Errorf("request last name: expected=%s; got=%s", req.LastName, sentRequest.LastName)
	}
	if sentRequest.LocationID != req.LocationID {
		t.Errorf("request location ID: expected=%d; got=%d", req.LocationID, sentRequest.LocationID)
	}

	// Check response
	if resp.ID != 12345 {
		t.Errorf("resp ID: expected=%d; got=%d", 12345, resp.ID)
	}
}

func TestClient_Delete(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead_delete.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodDelete,
		Path:       "/leads/123",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	resp, err := svc.Delete(context.Background(), 123)
	if err != nil {
		t.Fatalf("deleting leads: %v", err)
	}

	// Check response
	if resp.LeadID != 123 {
		t.Errorf("lead ID: expected=%d; got=%d", 123, resp.LeadID)
	}
	if !resp.IsSuccess {
		t.Error("lead response: expected=true; got=false")
	}
}

func TestClient_Update(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead.json")

	// Create test server and handler
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/leads/123",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.LeadUpdateRequest{
		FirstName: "Update",
		LastName:  "Lead",
		Email:     "updated@example.com",
	}
	resp, err := svc.Update(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("updating lead: %v", err)
	}

	// Check sent request
	var sentRequest leads.LeadUpdateRequest
	if err := json.Unmarshal(hdl.RequestBody, &sentRequest); err != nil {
		t.Fatalf("unmarshaling request: %v", err)
	}
	if sentRequest.FirstName != req.FirstName {
		t.Errorf("request first name: expected=%s; got=%s", req.FirstName, sentRequest.FirstName)
	}
	if sentRequest.LastName != req.LastName {
		t.Errorf("request last name: expected=%s; got=%s", req.LastName, sentRequest.LastName)
	}
	if sentRequest.Email != req.Email {
		t.Errorf("request email: expected=%s; got=%s", req.Email, sentRequest.Email)
	}

	// Check response
	if resp.ID != 12345 {
		t.Errorf("response ID: expected=%d; got=%d", 12345, resp.ID)
	}
}

func TestClient_Convert(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/lead_convert.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPost,
		Path:       "/leads/123/convert",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.LeadConvertRequest{
		LocationID: 13,
	}
	resp, err := svc.Convert(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("converting lead: %v", err)
	}

	// Check request
	var sentRequest leads.LeadConvertRequest
	if err := json.Unmarshal(hdl.RequestBody, &sentRequest); err != nil {
		t.Fatalf("decoding request: %v", err)
	}
	if sentRequest.LocationID != req.LocationID {
		t.Errorf("request location ID: expected=%d; got=%d", req.LocationID, sentRequest.LocationID)
	}

	// Check response
	if !resp.IsSuccess {
		t.Error("response is_success: expected=true; got=false")
	}
	if resp.ConvertedLeadID != 123 {
		t.Errorf("response converted lead ID: expected=%d; got=%d", 123, resp.ConvertedLeadID)
	}
	if resp.ClientData.ID != 1 {
		t.Errorf("response client ID: expected=%d; got=%d", 1, resp.ClientData.ID)
	}
}

func TestLeadUpdateRequestFrom(t *testing.T) {
	lead := &models.Lead{
		ID:         123,
		LocationID: 456,
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
	}

	updateReq := leads.LeadUpdateRequestFrom(lead)

	if lead.LocationID != updateReq.LocationID {
		t.Errorf("location ID: expected=%d; got=%d", updateReq.LocationID, lead.LocationID)
	}
	if lead.FirstName != updateReq.FirstName {
		t.Errorf("first name: expected=%s; got=%s", updateReq.FirstName, lead.FirstName)
	}
	if lead.LastName != updateReq.LastName {
		t.Errorf("last name: expected=%s; got=%s", updateReq.LastName, lead.LastName)
	}
	if lead.Email != updateReq.Email {
		t.Errorf("email: expected=%s; got=%s", updateReq.Email, lead.Email)
	}
}

func TestLeadConversionRequestFrom(t *testing.T) {
	lead := &models.Lead{
		ID:         123,
		LocationID: 456,
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
	}
	conversionReq := leads.LeadConvertRequestFrom(lead)

	if lead.LocationID != conversionReq.LocationID {
		t.Errorf("location ID: expected=%d; got=%d", 123, conversionReq.LocationID)
	}
	if lead.FirstName != conversionReq.FirstName {
		t.Errorf("first name: expected=%s; got=%s", "John", conversionReq.FirstName)
	}
	if lead.LastName != conversionReq.LastName {
		t.Errorf("last name: expected=%s; got=%s", "Doe", conversionReq.LastName)
	}
	if lead.Email != conversionReq.Email {
		t.Errorf("email: expected=%s; got=%s", lead.Email, conversionReq.Email)
	}
	if conversionReq.ClientStatusID != 0 {
		t.Errorf("client_status_id: expected=%d; got=%d", 0, conversionReq.ClientStatusID)
	}
}
