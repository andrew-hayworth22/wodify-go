package clients_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/models"
)

func TestClient_Get(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client.json")

	// Create mock server and client
	svr := testutil.NewServer(t, &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/clients/12345",
		StatusCode: http.StatusOK,
		Body:       body,
	})
	svc := clients.New(svr)

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
	body := testutil.MustReadJSONFixture(t, "testdata/client_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/clients",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(clients.ClientFieldLastName)
	req := clients.NewClientListRequest(p, s)
	resp, err := svc.List(context.Background(), req)
	if err != nil {
		t.Fatalf("listing clients: %v", err)
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
	if len(resp.Clients) != 2 {
		t.Errorf("response client list length: expected=%d; got=%d", 2, len(resp.Clients))
	}
}

func TestClient_Search(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/clients/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortDescending(clients.ClientFieldFirstName)
	q := clients.NewClientQuery().Eq(clients.ClientFieldFirstName, "john")
	req := clients.NewClientSearchRequest(p, s, q)
	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		t.Fatalf("searching clients: %v", err)
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
	if len(resp.Clients) != 2 {
		t.Errorf("response clients list length: expected=%d; got=%d", 2, len(resp.Clients))
	}
}

func TestClient_Create(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client.json")

	// Create mock server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPost,
		Path:       "/clients",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	req := clients.ClientCreateRequest{
		FirstName:  "John",
		LastName:   "Doe",
		LocationID: 2998,
	}
	resp, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("creating client: %v", err)
	}

	// Check sent request
	var sentRequest clients.ClientCreateRequest
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

func TestClient_Deactivate(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client_action.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/clients/123/deactivate",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	resp, err := svc.Deactivate(context.Background(), 123)
	if err != nil {
		t.Fatalf("deactivating client: %v", err)
	}

	// Check response
	if resp.Client.ID != 123 {
		t.Errorf("client ID: expected=%d; got=%d", 123, resp.Client.ID)
	}
	if !resp.IsSuccess {
		t.Error("client response: expected=true; got=false")
	}
}

func TestClient_Reactivate(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client_action.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/clients/123/reactivate",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	resp, err := svc.Reactivate(context.Background(), 123)
	if err != nil {
		t.Fatalf("reactivating client: %v", err)
	}

	// Check response
	if resp.Client.ID != 123 {
		t.Errorf("client ID: expected=%d; got=%d", 123, resp.Client.ID)
	}
	if !resp.IsSuccess {
		t.Error("client response: expected=true; got=false")
	}
}

func TestClient_Suspend(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client_action.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/clients/123/suspend",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	resp, err := svc.Suspend(context.Background(), 123)
	if err != nil {
		t.Fatalf("suspending client: %v", err)
	}

	// Check response
	if resp.Client.ID != 123 {
		t.Errorf("client ID: expected=%d; got=%d", 123, resp.Client.ID)
	}
	if !resp.IsSuccess {
		t.Error("client response: expected=true; got=false")
	}
}

func TestClient_Reinstate(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client_action.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/clients/123/reinstate",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	resp, err := svc.Reinstate(context.Background(), 123)
	if err != nil {
		t.Fatalf("reinstating client: %v", err)
	}

	// Check response
	if resp.Client.ID != 123 {
		t.Errorf("client ID: expected=%d; got=%d", 123, resp.Client.ID)
	}
	if !resp.IsSuccess {
		t.Error("client response: expected=true; got=false")
	}
}

func TestClient_Update(t *testing.T) {
	// Load test fixture
	body := testutil.MustReadJSONFixture(t, "testdata/client.json")

	// Create test server and handler
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/clients/123",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := clients.New(svr)

	// Make request
	req := clients.ClientUpdateRequest{
		FirstName: "Update",
		LastName:  "Client",
		Email:     "updated@example.com",
	}
	resp, err := svc.Update(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("updating client: %v", err)
	}

	// Check sent request
	var sentRequest clients.ClientUpdateRequest
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

func TestClientUpdateRequestFrom(t *testing.T) {
	client := &models.Client{
		ID:         123,
		LocationID: 456,
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
	}

	updateReq := clients.ClientUpdateRequestFrom(client)

	if client.LocationID != updateReq.LocationID {
		t.Errorf("location ID: expected=%d; got=%d", updateReq.LocationID, client.LocationID)
	}
	if client.FirstName != updateReq.FirstName {
		t.Errorf("first name: expected=%s; got=%s", updateReq.FirstName, client.FirstName)
	}
	if client.LastName != updateReq.LastName {
		t.Errorf("last name: expected=%s; got=%s", updateReq.LastName, client.LastName)
	}
	if client.Email != updateReq.Email {
		t.Errorf("email: expected=%s; got=%s", updateReq.Email, client.Email)
	}
}
