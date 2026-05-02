package leads_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/models"
)

func TestGet(t *testing.T) {
	body, err := os.ReadFile("testdata/lead.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	svr := testutil.NewServer(t, &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads/12345",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	})
	svc := leads.New(svr)

	lead, err := svc.Get(context.Background(), 12345)
	if err != nil {
		t.Fatalf("getting lead: %v", err)
	}

	const expectedID = 12345
	if lead.ID != expectedID {
		t.Errorf("lead ID: expected=%d; got=%d", expectedID, lead.ID)
	}

	const expectedFirstName = "John"
	if lead.FirstName != expectedFirstName {
		t.Errorf("lead first name: expected=%s; got=%s", expectedFirstName, lead.FirstName)
	}

	const expectedLastName = "Doe"
	if lead.LastName != expectedLastName {
		t.Errorf("lead last name: expected=%s; got=%s", expectedLastName, lead.LastName)
	}

	expectedDateOfBirth := models.NewDate(time.Date(2001, time.December, 31, 0, 0, 0, 0, time.UTC))
	if lead.DateOfBirth != expectedDateOfBirth {
		t.Errorf("lead date of birth: expected=%s; got=%s", expectedDateOfBirth, lead.DateOfBirth)
	}

	const expectedGender = models.GenderFemale
	if lead.Gender.Name != expectedGender {
		t.Errorf("lead gender: expected=%s; got=%s", expectedGender, lead.Gender.Name)
	}

	expectednextClassReservation := models.NewDateTime(time.Date(2014, time.December, 31, 23, 59, 59, 938_000_000, time.UTC))
	if lead.NextClassReservation != expectednextClassReservation {
		t.Errorf("lead next class reservation: expected=%s; got=%s", expectednextClassReservation, lead.NextClassReservation)
	}
}

func TestCreate(t *testing.T) {
	body, err := os.ReadFile("testdata/lead.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	hdl := &testutil.Handler{
		Method:     http.MethodPost,
		Path:       "/leads",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	req := leads.CreateLeadRequest{
		FirstName:  "John",
		LastName:   "Doe",
		LocationID: 2998,
	}

	lead, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("creating lead: %v", err)
	}

	var sentRequest leads.CreateLeadRequest
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

	if lead.ID != 12345 {
		t.Errorf("lead ID: expected=%d; got=%d", 12345, lead.ID)
	}
}
func TestList(t *testing.T) {
	body, err := os.ReadFile("testdata/lead_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/leads",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	req := leads.ListRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: leads.NewSort(leads.SortByFirstName, false),
	}

	leadList, err := svc.List(context.Background(), req)
	if err != nil {
		t.Fatalf("creating lead: %v", err)
	}

	if leadList.Pagination.Page != req.Page.Page {
		t.Errorf("pagination page: expected=%d; got=%d", req.Page.Page, leadList.Pagination.Page)
	}
}
