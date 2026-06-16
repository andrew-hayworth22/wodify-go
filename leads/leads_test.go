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

func TestClient(t *testing.T) {
	t.Parallel()

	pagination := wodify.NewPaginationRequest(1, 10)

	leadFixture := testutil.MustReadJSONFixture(t, "testdata/lead.json")
	leadListFixture := testutil.MustReadJSONFixture(t, "testdata/lead_list.json")
	leadDeleteFixture := testutil.MustReadJSONFixture(t, "testdata/lead_delete.json")
	leadConvertFixture := testutil.MustReadJSONFixture(t, "testdata/lead_convert.json")
	leadSort := wodify.SortDescending(leads.LeadFieldLastName)
	leadQuery := leads.NewLeadQuery().Eq(leads.LeadFieldCity, "Canal Fulton")
	leadCreateReq := leads.LeadCreateRequest{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		LeadStatusID:          5,
		LocationID:            20,
		GenderID:              1,
		PhoneNumber:           "555-867-5309",
		DateOfBirth:           models.NewDate(time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)),
		StreetAddress1:        "123 Main St",
		StreetAddress2:        "Apt 4B",
		City:                  "Springfield",
		StateID:               40,
		Province:              "",
		ZipCode:               "62701",
		CountryID:             50,
		Tags:                  []string{"vip", "new"},
		Notes:                 "VIP client",
		EmergencyContactName:  "Mary Doe",
		EmergencyContactPhone: "555-111-2222",
		LeadSourceID:          90,
		ReferredByFromWeb:     "Google",
		ReferredByUserID:      100,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       false,
		LeadOwnerID:           130,
	}
	leadUpdateReq := leads.LeadUpdateRequest{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		LeadStatusID:          5,
		LocationID:            20,
		GenderID:              1,
		PhoneNumber:           "555-867-5309",
		DateOfBirth:           models.NewDate(time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)),
		StreetAddress1:        "123 Main St",
		StreetAddress2:        "Apt 4B",
		City:                  "Springfield",
		StateID:               40,
		Province:              "",
		ZipCode:               "62701",
		CountryID:             50,
		Notes:                 "VIP client",
		EmergencyContactName:  "Mary Doe",
		EmergencyContactPhone: "555-111-2222",
		LeadSourceID:          90,
		ReferredByFromWeb:     "Google",
		ReferredByUserID:      100,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       false,
		LeadOwnerID:           130,
	}
	leadConvertReq := leads.LeadConvertRequest{
		LocationID:     12,
		Email:          "john.doe@example.com",
		FirstName:      "John",
		LastName:       "Doe",
		ClientStatusID: 2,
		GenderID:       1,
		BillingCCEmail: "johnnydoe@example.com",
		MobileNumber:   "123-123-1234",
		DateOfBirth:    models.Date{Time: time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)},
		StreetAddress1: "123 Main St",
		StreetAddress2: "Apt 4B",
		City:           "Canal Fulton",
		StateID:        5,
		Province:       "Province",
		CountryID:      10,
		ZipCode:        "12345",
		ClientOwnerID:  89,
	}

	bookingListFixture := testutil.MustReadJSONFixture(t, "testdata/booking_list.json")
	bookingSort := wodify.SortDescending(leads.BookingFieldAppointmentID)
	bookingQuery := leads.NewBookingQuery().Eq(leads.BookingFieldLocationID, 12)

	classSignInListFixture := testutil.MustReadJSONFixture(t, "testdata/class_sign_in_list.json")
	classSignInSort := wodify.SortDescending(leads.ClassSignInFieldID)
	classSignInQuery := leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldID, 12)

	reservationListFixture := testutil.MustReadJSONFixture(t, "testdata/reservation_list.json")
	reservationSort := wodify.SortDescending(leads.ReservationFieldID)
	reservationQuery := leads.NewReservationQuery().Eq(leads.ReservationFieldID, 12)

	performanceResultListFixture := testutil.MustReadJSONFixture(t, "testdata/performance_result_list.json")

	sourceListFixture := testutil.MustReadJSONFixture(t, "testdata/source_list.json")
	sourceSort := wodify.SortDescending(leads.SourceFieldID)

	statusListFixture := testutil.MustReadJSONFixture(t, "testdata/status_list.json")
	statusSort := wodify.SortDescending(leads.StatusFieldID)

	tagAddFixture := testutil.MustReadJSONFixture(t, "testdata/tag_add.json")
	tagDeleteFixture := testutil.MustReadJSONFixture(t, "testdata/tag_remove.json")
	tagUpdateReq := leads.TagsUpdateRequest{Tags: []string{"vip", "new"}}

	groupRoleListFixture := testutil.MustReadJSONFixture(t, "testdata/group_role_list.json")
	groupRoleSort := wodify.SortDescending(leads.GroupRoleFieldID)
	groupRoleQuery := leads.NewGroupRoleQuery().Eq(leads.GroupRoleFieldID, 1)

	groupFixture := testutil.MustReadJSONFixture(t, "testdata/group.json")
	convertDependentFixture := testutil.MustReadJSONFixture(t, "testdata/lead_convert_dependent.json")
	groupCreateReq := leads.GroupCreateRequest{
		GroupParticipants: []leads.GroupParticipantInput{
			{
				GroupParticipantLeadID: 1,
				GroupRoleID:            2,
			},
			{
				GroupParticipantLeadID: 3,
				GroupRoleID:            4,
			},
		},
	}
	participantsReq := leads.GroupParticipantsRequest{
		LeadIDs: []int64{123, 345, 678},
	}
	convertDependentReq := leads.ConvertFromDependentRequest{Email: "john.doe@example.com"}

	testCases := []struct {
		name     string
		endpoint *testutil.Endpoint
		run      func(*testing.T, *leads.Client)
	}{
		{
			name: "get",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123", http.StatusOK,
				testutil.WithResponseBody(leadFixture),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Get(context.Background(), 123)
				if err != nil {
					t.Fatalf("getting lead: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadFixture, respJSON)
			},
		},
		{
			name:     "get - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Get(context.Background(), 123)
				if err == nil {
					t.Fatalf("getting lead: expected error")
				}
				if resp != nil {
					t.Fatalf("getting lead: expected nil response")
				}
			},
		},
		{
			name: "list",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads", http.StatusOK,
				testutil.WithResponseBody(leadListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(leadSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.List(context.Background(), leads.NewLeadListRequest(pagination, leadSort))
				if err != nil {
					t.Fatalf("listing leads: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadListFixture, respJSON)
			},
		},
		{
			name:     "list - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.List(context.Background(), leads.NewLeadListRequest(pagination, leadSort))
				if err == nil {
					t.Fatalf("listing leads: expected error")
				}
				if resp != nil {
					t.Fatalf("listing leads: expected nil response")
				}
			},
		},
		{
			name: "search",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/search", http.StatusOK,
				testutil.WithResponseBody(leadListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(leadSort),
				testutil.WithExpectedRequestQuery(leadQuery),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Search(context.Background(), leads.NewLeadSearchRequest(pagination, leadSort, leadQuery))
				if err != nil {
					t.Fatalf("searching leads: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadListFixture, respJSON)
			},
		},
		{
			name:     "search - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Search(context.Background(), leads.NewLeadSearchRequest(pagination, leadSort, leadQuery))
				if err == nil {
					t.Fatalf("searching leads: expected error")
				}
				if resp != nil {
					t.Fatalf("searching leads: expected nil response")
				}
			},
		},
		{
			name: "create",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads", http.StatusOK,
				testutil.WithResponseBody(leadFixture),
				testutil.WithExpectedRequestBody(leadCreateReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Create(context.Background(), leadCreateReq)
				if err != nil {
					t.Fatalf("creating lead: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadFixture, respJSON)
			},
		},
		{
			name:     "create - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Create(context.Background(), leadCreateReq)
				if err == nil {
					t.Fatalf("creating lead: expected error")
				}
				if resp != nil {
					t.Fatalf("creating lead: expected nil response")
				}
			},
		},
		{
			name: "delete",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/123", http.StatusOK,
				testutil.WithResponseBody(leadDeleteFixture),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Delete(context.Background(), 123)
				if err != nil {
					t.Fatalf("deleting lead: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadDeleteFixture, respJSON)
			},
		},
		{
			name:     "delete - error",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/123", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Delete(context.Background(), 123)
				if err == nil {
					t.Fatalf("deleting lead: expected error")
				}
				if resp != nil {
					t.Fatalf("deleting lead: expected nil response")
				}
			},
		},
		{
			name: "update",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123", http.StatusOK,
				testutil.WithResponseBody(leadFixture),
				testutil.WithExpectedRequestBody(leadUpdateReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Update(context.Background(), 123, leadUpdateReq)
				if err != nil {
					t.Fatalf("updating lead: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadFixture, respJSON)
			},
		},
		{
			name:     "update - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Update(context.Background(), 123, leadUpdateReq)
				if err == nil {
					t.Fatalf("updating lead: expected error")
				}
				if resp != nil {
					t.Fatalf("updating lead: expected nil response")
				}
			},
		},
		{
			name: "convert",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads/123/convert", http.StatusOK,
				testutil.WithResponseBody(leadConvertFixture),
				testutil.WithExpectedRequestBody(leadConvertReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Convert(context.Background(), 123, leadConvertReq)
				if err != nil {
					t.Fatalf("converting lead: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, leadConvertFixture, respJSON)
			},
		},
		{
			name:     "convert - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads/123/convert", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.Convert(context.Background(), 123, leadConvertReq)
				if err == nil {
					t.Fatalf("converting lead: expected error")
				}
				if resp != nil {
					t.Fatalf("converting lead: expected nil response")
				}
			},
		},
		{
			name: "list bookings",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/appointments/bookings", http.StatusOK,
				testutil.WithResponseBody(bookingListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(bookingSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListBookings(context.Background(), 123, leads.NewBookingListRequest(pagination, bookingSort))
				if err != nil {
					t.Fatalf("listing bookings: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, bookingListFixture, respJSON)
			},
		},
		{
			name:     "list bookings - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/appointments/bookings", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListBookings(context.Background(), 123, leads.NewBookingListRequest(pagination, bookingSort))
				if err == nil {
					t.Fatalf("listing bookings: expected error")
				}
				if resp != nil {
					t.Fatalf("listing bookings: expected nil response")
				}
			},
		},
		{
			name: "search bookings",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/appointments/bookings/search", http.StatusOK,
				testutil.WithResponseBody(bookingListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(bookingSort),
				testutil.WithExpectedRequestQuery(bookingQuery),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchBookings(context.Background(), 123, leads.NewBookingSearchRequest(pagination, bookingSort, bookingQuery))
				if err != nil {
					t.Fatalf("searching bookings: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, bookingListFixture, respJSON)
			},
		},
		{
			name:     "search bookings - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/appointments/bookings/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchBookings(context.Background(), 123, leads.NewBookingSearchRequest(pagination, bookingSort, bookingQuery))
				if err == nil {
					t.Fatalf("searching bookings: expected error")
				}
				if resp != nil {
					t.Fatalf("searching bookings: expected nil response")
				}
			},
		},
		{
			name: "list class sign-ins",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/sign-ins", http.StatusOK,
				testutil.WithResponseBody(classSignInListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(classSignInSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListClassSignIns(context.Background(), 123, leads.NewClassSignInListRequest(pagination, classSignInSort))
				if err != nil {
					t.Fatalf("listing class sign-ins: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, classSignInListFixture, respJSON)
			},
		},
		{
			name:     "list class sign-ins - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/sign-ins", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListClassSignIns(context.Background(), 123, leads.NewClassSignInListRequest(pagination, classSignInSort))
				if err == nil {
					t.Fatalf("listing class sign-ins: expected error")
				}
				if resp != nil {
					t.Fatalf("listing class sign-ins: expected nil response")
				}
			},
		},
		{
			name: "search class sign-ins",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/sign-ins/search", http.StatusOK,
				testutil.WithResponseBody(classSignInListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(classSignInSort),
				testutil.WithExpectedRequestQuery(classSignInQuery),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchClassSignIns(context.Background(), 123, leads.NewClassSignInSearchRequest(pagination, classSignInSort, classSignInQuery))
				if err != nil {
					t.Fatalf("searching class sign-ins: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, classSignInListFixture, respJSON)
			},
		},
		{
			name:     "search class sign-ins - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/sign-ins/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchClassSignIns(context.Background(), 123, leads.NewClassSignInSearchRequest(pagination, classSignInSort, classSignInQuery))
				if err == nil {
					t.Fatalf("searching class sign-ins: expected error")
				}
				if resp != nil {
					t.Fatalf("searching class sign-ins: expected nil response")
				}
			},
		},
		{
			name: "list performance results",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/performance-results", http.StatusOK,
				testutil.WithResponseBody(performanceResultListFixture),
				testutil.WithExpectedRequestPagination(pagination),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListPerformanceResults(context.Background(), 123, leads.NewPerformanceResultListRequest(pagination))
				if err != nil {
					t.Fatalf("listing performance results: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, performanceResultListFixture, respJSON)
			},
		},
		{
			name:     "list performance results - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/performance-results", http.StatusBadRequest), run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListPerformanceResults(context.Background(), 123, leads.NewPerformanceResultListRequest(pagination))
				if err == nil {
					t.Fatalf("listing performance results: expected error")
				}
				if resp != nil {
					t.Fatalf("listing performance results: expected nil response")
				}
			},
		},
		{
			name: "list performance results by component",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/performance-results/components/456", http.StatusOK,
				testutil.WithResponseBody(performanceResultListFixture),
				testutil.WithExpectedRequestPagination(pagination),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListPerformanceResultsByComponent(context.Background(), 123, 456, leads.NewPerformanceResultListRequest(pagination))
				if err != nil {
					t.Fatalf("listing performance results by component: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, performanceResultListFixture, respJSON)
			},
		},
		{
			name:     "list performance results by component - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/performance-results/components/456", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListPerformanceResultsByComponent(context.Background(), 123, 456, leads.NewPerformanceResultListRequest(pagination))
				if err == nil {
					t.Fatalf("listing performance results by component: expected error")
				}
				if resp != nil {
					t.Fatalf("listing performance results by component: expected nil response")
				}
			},
		},
		{
			name: "list reservations",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/reservations", http.StatusOK,
				testutil.WithResponseBody(reservationListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(reservationSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListReservations(context.Background(), 123, leads.NewReservationListRequest(pagination, reservationSort))
				if err != nil {
					t.Fatalf("listing reservations: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, reservationListFixture, respJSON)
			},
		},
		{
			name:     "list reservations - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/reservations", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListReservations(context.Background(), 123, leads.NewReservationListRequest(pagination, reservationSort))
				if err == nil {
					t.Fatalf("listing reservations: expected error")
				}
				if resp != nil {
					t.Fatalf("listing reservations: expected nil response")
				}
			},
		},
		{
			name: "search reservations",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/reservations/search", http.StatusOK,
				testutil.WithResponseBody(reservationListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(reservationSort),
				testutil.WithExpectedRequestQuery(reservationQuery),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchReservations(context.Background(), 123, leads.NewReservationSearchRequest(pagination, reservationSort, reservationQuery))
				if err != nil {
					t.Fatalf("searching reservations: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, reservationListFixture, respJSON)
			},
		},
		{
			name:     "search reservations - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/123/classes/reservations/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchReservations(context.Background(), 123, leads.NewReservationSearchRequest(pagination, reservationSort, reservationQuery))
				if err == nil {
					t.Fatalf("searching reservations: expected error")
				}
				if resp != nil {
					t.Fatalf("searching reservations: expected nil response")
				}
			},
		},
		{
			name: "list sources",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/sources", http.StatusOK,
				testutil.WithResponseBody(sourceListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(sourceSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListSources(context.Background(), leads.NewSourceListRequest(pagination, sourceSort))
				if err != nil {
					t.Fatalf("listing sources: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, sourceListFixture, respJSON)
			},
		},
		{
			name:     "list sources - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/sources", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListSources(context.Background(), leads.NewSourceListRequest(pagination, sourceSort))
				if err == nil {
					t.Fatalf("listing sources: expected error")
				}
				if resp != nil {
					t.Fatalf("listing sources: expected nil response")
				}
			},
		},
		{
			name: "list statuses",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/statuses", http.StatusOK,
				testutil.WithResponseBody(statusListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(statusSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListStatuses(context.Background(), leads.NewStatusListRequest(pagination, statusSort))
				if err != nil {
					t.Fatalf("listing statuses: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, statusListFixture, respJSON)
			},
		},
		{
			name:     "list statuses - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/statuses", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListStatuses(context.Background(), leads.NewStatusListRequest(pagination, statusSort))
				if err == nil {
					t.Fatalf("listing statuses: expected error")
				}
				if resp != nil {
					t.Fatalf("listing statuses: expected nil response")
				}
			},
		},
		{
			name: "add tags",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123/tags", http.StatusOK,
				testutil.WithResponseBody(tagAddFixture),
				testutil.WithExpectedRequestBody(tagUpdateReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.AddTags(context.Background(), 123, tagUpdateReq)
				if err != nil {
					t.Fatalf("adding tags: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, tagAddFixture, respJSON)
			},
		},
		{
			name:     "add tags - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123/tags", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.AddTags(context.Background(), 123, tagUpdateReq)
				if err == nil {
					t.Fatalf("adding tags: expected error")
				}
				if resp != nil {
					t.Fatalf("adding tags: expected nil response")
				}
			},
		},
		{
			name: "delete tags",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/123/tags", http.StatusOK,
				testutil.WithResponseBody(tagDeleteFixture),
				testutil.WithExpectedRequestBody(tagUpdateReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.DeleteTags(context.Background(), 123, tagUpdateReq)
				if err != nil {
					t.Fatalf("deleting tags: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, tagDeleteFixture, respJSON)
			},
		},
		{
			name:     "delete tags - error",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/123/tags", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.DeleteTags(context.Background(), 123, tagUpdateReq)
				if err == nil {
					t.Fatalf("deleting tags: expected error")
				}
				if resp != nil {
					t.Fatalf("deleting tags: expected nil response")
				}
			},
		},
		{
			name: "list group roles",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/group/roles", http.StatusOK,
				testutil.WithResponseBody(groupRoleListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(groupRoleSort),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListGroupRoles(context.Background(), leads.NewGroupRoleListRequest(pagination, groupRoleSort))
				if err != nil {
					t.Fatalf("listing group roles: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupRoleListFixture, respJSON)
			},
		},
		{
			name:     "list group roles - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/group/roles", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ListGroupRoles(context.Background(), leads.NewGroupRoleListRequest(pagination, groupRoleSort))
				if err == nil {
					t.Fatalf("listing group roles: expected error")
				}
				if resp != nil {
					t.Fatalf("listing group roles: expected nil response")
				}
			},
		},
		{
			name: "search group roles",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/group/roles/search", http.StatusOK,
				testutil.WithResponseBody(groupRoleListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(groupRoleSort),
				testutil.WithExpectedRequestQuery(groupRoleQuery),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchGroupRoles(context.Background(), leads.NewGroupRoleSearchRequest(pagination, groupRoleSort, groupRoleQuery))
				if err != nil {
					t.Fatalf("searching group roles: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupRoleListFixture, respJSON)
			},
		},
		{
			name:     "search group roles - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/leads/group/roles/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.SearchGroupRoles(context.Background(), leads.NewGroupRoleSearchRequest(pagination, groupRoleSort, groupRoleQuery))
				if err == nil {
					t.Fatalf("searching group roles: expected error")
				}
				if resp != nil {
					t.Fatalf("searching group roles: expected nil response")
				}
			},
		},
		{
			name: "create group",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads/group/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(groupCreateReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.CreateGroup(context.Background(), groupCreateReq)
				if err != nil {
					t.Fatalf("creating group: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupFixture, respJSON)
			},
		},
		{
			name:     "create group - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/leads/group/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.CreateGroup(context.Background(), groupCreateReq)
				if err == nil {
					t.Fatalf("creating group: expected error")
				}
				if resp != nil {
					t.Fatalf("creating group: expected nil response")
				}
			},
		},
		{
			name: "add participants from group",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/group/123/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(participantsReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.AddGroupParticipants(context.Background(), 123, participantsReq)
				if err != nil {
					t.Fatalf("adding participants from group: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupFixture, respJSON)
			},
		},
		{
			name:     "add participants from group - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/group/123/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.AddGroupParticipants(context.Background(), 123, participantsReq)
				if err == nil {
					t.Fatalf("adding participants from group: expected error")
				}
				if resp != nil {
					t.Fatalf("adding participants from group: expected nil response")
				}
			},
		},
		{
			name: "remove participants from group",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/group/123/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(participantsReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.RemoveGroupParticipants(context.Background(), 123, participantsReq)
				if err != nil {
					t.Fatalf("removing participants from group: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupFixture, respJSON)
			},
		},
		{
			name:     "remove participants from group - error",
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/leads/group/123/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.RemoveGroupParticipants(context.Background(), 123, participantsReq)
				if err == nil {
					t.Fatalf("removing participants from group: expected error")
				}
				if resp != nil {
					t.Fatalf("removing participants from group: expected nil response")
				}
			},
		},
		{
			name: "convert lead from dependent",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123/convert-from-dependent", http.StatusOK,
				testutil.WithResponseBody(convertDependentFixture),
				testutil.WithExpectedRequestBody(convertDependentReq),
			),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ConvertFromDependent(context.Background(), 123, convertDependentReq)
				if err != nil {
					t.Fatalf("converting lead from dependent: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, convertDependentFixture, respJSON)
			},
		},
		{
			name:     "convert lead from dependent - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/leads/123/convert-from-dependent", http.StatusBadRequest),
			run: func(t *testing.T, svc *leads.Client) {
				resp, err := svc.ConvertFromDependent(context.Background(), 123, convertDependentReq)
				if err == nil {
					t.Fatalf("converting lead from dependent: expected error")
				}
				if resp != nil {
					t.Fatalf("converting lead from dependent: expected nil response")
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			svr := testutil.NewWodifyClient(t, "test-key", 0, testCase.endpoint)
			svc := leads.New(svr)
			testCase.run(t, svc)
		})
	}
}

func TestLeadUpdateRequestFrom(t *testing.T) {
	t.Parallel()
	dob := models.NewDate(time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC))
	l := &models.Lead{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		LeadStatusID:          5,
		LocationID:            20,
		GenderID:              1,
		PhoneNumber:           "555-867-5309",
		DateOfBirth:           dob,
		StreetAddress1:        "123 Main St",
		StreetAddress2:        "Apt 4B",
		City:                  "Springfield",
		StateID:               40,
		Province:              "",
		ZipCode:               "62701",
		CountryID:             50,
		Notes:                 "VIP client",
		EmergencyContactName:  "Mary Doe",
		EmergencyContactPhone: "555-111-2222",
		LeadSourceID:          90,
		ReferredByFromWeb:     "Google",
		ReferredByUserID:      100,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       false,
		LeadOwnerID:           130,
	}
	expected := leads.LeadUpdateRequest{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		LeadStatusID:          5,
		LocationID:            20,
		GenderID:              1,
		PhoneNumber:           "555-867-5309",
		DateOfBirth:           dob,
		StreetAddress1:        "123 Main St",
		StreetAddress2:        "Apt 4B",
		City:                  "Springfield",
		StateID:               40,
		Province:              "",
		ZipCode:               "62701",
		CountryID:             50,
		Notes:                 "VIP client",
		EmergencyContactName:  "Mary Doe",
		EmergencyContactPhone: "555-111-2222",
		LeadSourceID:          90,
		ReferredByFromWeb:     "Google",
		ReferredByUserID:      100,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       false,
		LeadOwnerID:           130,
	}

	actual := leads.LeadUpdateRequestFrom(l)
	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("marshaling actual: %v", err)
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("marshaling expected: %v", err)
	}
	testutil.AssertJSONEqual(t, expectedJSON, actualJSON)
}

func TestLeadConvertRequestFrom(t *testing.T) {
	t.Parallel()
	dob := models.NewDate(time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC))
	l := &models.Lead{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		PhoneNumber:    "555-867-5309",
		LocationID:     20,
		GenderID:       1,
		DateOfBirth:    dob,
		StreetAddress1: "123 Main St",
		StreetAddress2: "Apt 4B",
		City:           "Springfield",
		StateID:        40,
		Province:       "",
		CountryID:      50,
		ZipCode:        "62701",
		LeadOwnerID:    130,
	}
	expected := leads.LeadConvertRequest{
		LocationID:     20,
		Email:          "john.doe@example.com",
		FirstName:      "John",
		LastName:       "Doe",
		ClientStatusID: 0,
		GenderID:       1,
		BillingCCEmail: "",
		MobileNumber:   "555-867-5309",
		DateOfBirth:    dob,
		StreetAddress1: "123 Main St",
		StreetAddress2: "Apt 4B",
		City:           "Springfield",
		StateID:        40,
		Province:       "",
		CountryID:      50,
		ZipCode:        "62701",
		ClientOwnerID:  130,
	}

	actual := leads.LeadConvertRequestFrom(l)
	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("marshaling actual: %v", err)
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("marshaling expected: %v", err)
	}
	testutil.AssertJSONEqual(t, expectedJSON, actualJSON)
}
