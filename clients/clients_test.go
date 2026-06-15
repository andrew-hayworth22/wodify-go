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

func TestClient(t *testing.T) {
	t.Parallel()

	pagination := wodify.NewPaginationRequest(1, 10)

	clientFixture := testutil.MustReadJSONFixture(t, "testdata/client.json")
	clientListFixture := testutil.MustReadJSONFixture(t, "testdata/client_list.json")
	clientActionFixture := testutil.MustReadJSONFixture(t, "testdata/client_action.json")
	clientSort := wodify.SortDescending(clients.ClientFieldFirstName)
	clientQuery := clients.NewClientQuery().Eq(clients.ClientFieldFirstName, clientSort)
	clientCreateReq := clients.ClientCreateRequest{
		FirstName:             "john",
		LastName:              "doe",
		Email:                 "john.doe@example.com",
		PhoneNumber:           "1231231234",
		BillingCCEmail:        "johnnydoe@example.org",
		ClientStatusID:        1,
		LocationID:            123,
		DateOfBirth:           models.Date{Time: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		GenderID:              2,
		StreetAddress1:        "123 Example St",
		StreetAddress2:        "Apt. A",
		City:                  "Example",
		StateID:               3,
		Province:              "Province",
		ZipCode:               "12345",
		CountryID:             4,
		TimezoneID:            5,
		Height1Measurement:    6,
		Height2Measurement:    7,
		Weight:                128,
		IsEmailSubscribed:     true,
		Tags:                  []string{"tag 1", "tag 2"},
		Notes:                 "Notes",
		EmergencyContactName:  "Jane Doe",
		EmergencyContactPhone: "jane.doe@example.com",
		LeadSourceID:          8,
		ReferringUserID:       9,
		IsSMSSubscribed:       true,
		ClientOwnerId:         12,
	}
	clientUpdateReq := clients.ClientUpdateRequest{
		FirstName:             "john",
		LastName:              "doe",
		Email:                 "john.doe@example.com",
		PhoneNumber:           "1231231234",
		BillingCCEmail:        "johnnydoe@example.org",
		ClientStatusID:        1,
		LocationID:            123,
		DefaultProgramID:      2,
		DateOfBirth:           models.Date{Time: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		GenderID:              2,
		StreetAddress1:        "123 Example St",
		StreetAddress2:        "Apt. A",
		City:                  "Example",
		StateID:               3,
		Province:              "Province",
		ZipCode:               "12345",
		CountryID:             4,
		TimezoneID:            5,
		Height1Measurement:    6,
		Height2Measurement:    7,
		Weight:                128,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       true,
		Notes:                 "Notes",
		IsOverwriteNotes:      true,
		EmergencyContactName:  "Jane Doe",
		EmergencyContactPhone: "jane.doe@example.com",
		LeadSourceID:          8,
		ReferringUserID:       9,
		ClientOwnerID:         12,
		CoachTitle:            "Coach Title",
		CoachBio:              "Coach Bio",
		CoachLink1Icon:        "Link 1 Icon",
		CoachLink1URL:         "Link 1 URL",
		CoachLink2Icon:        "Link 2 Icon",
		CoachLink2URL:         "Link 2 URL",
		CoachLink3Icon:        "Link 3 Icon",
		CoachLink3URL:         "Link 3 URL",
		CoachLink4Icon:        "Link 4 Icon",
		CoachLink4URL:         "Link 4 URL",
		CoachLink5Icon:        "Link 5 Icon",
		CoachLink5URL:         "Link 5 URL",
	}
	registerLinkFixture := testutil.MustReadJSONFixture(t, "testdata/register_link.json")

	statusListFixture := testutil.MustReadJSONFixture(t, "testdata/status_list.json")
	statusSort := wodify.SortAscending(clients.StatusFieldName)
	statusQuery := clients.NewStatusQuery().Eq(clients.StatusFieldName, "Active")

	groupRoleListFixture := testutil.MustReadJSONFixture(t, "testdata/group_role_list.json")
	groupRoleSort := wodify.SortDescending(clients.GroupRoleFieldID)
	groupRoleQuery := clients.NewGroupRoleQuery().Eq(clients.GroupRoleFieldID, 1)

	groupFixture := testutil.MustReadJSONFixture(t, "testdata/group.json")
	convertDependentFixture := testutil.MustReadJSONFixture(t, "testdata/client_convert_dependent.json")
	groupCreateReq := clients.GroupCreateRequest{
		GroupParticipants: []clients.GroupParticipantInput{
			{
				GroupParticipantClientID: 1,
				GroupRoleID:              2,
			},
			{
				GroupParticipantClientID: 3,
				GroupRoleID:              4,
			},
		},
	}
	participantsReq := clients.GroupParticipantsRequest{
		ClientIDs: []int64{123, 345, 678},
	}
	convertDependentReq := clients.ConvertFromDependentRequest{Email: "john.doe@example.com"}

	testCases := []struct {
		name     string
		endpoint *testutil.Endpoint
		run      func(*testing.T, *clients.Client)
	}{
		{
			name: "get",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/123", http.StatusOK,
				testutil.WithResponseBody(clientFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Get(context.Background(), 123)
				if err != nil {
					t.Fatalf("getting client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientFixture, respJSON)
			},
		},
		{
			name:     "get - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/123", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Get(context.Background(), 123)
				if err == nil {
					t.Fatal("getting client: expected error")
				}
				if resp != nil {
					t.Fatal("getting client: expected nil response")
				}
			},
		},
		{
			name: "list",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients", http.StatusOK,
				testutil.WithResponseBody(clientListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(clientSort),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.List(context.Background(), clients.NewClientListRequest(pagination, clientSort))
				if err != nil {
					t.Fatalf("listing clients: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientListFixture, respJSON)
			},
		},
		{
			name:     "list - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.List(context.Background(), clients.NewClientListRequest(pagination, clientSort))
				if err == nil {
					t.Fatal("listing clients: expected error")
				}
				if resp != nil {
					t.Fatal("listing clients: expected nil response")
				}
			},
		},
		{
			name: "search",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/search", http.StatusOK,
				testutil.WithResponseBody(clientListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(clientSort),
				testutil.WithExpectedRequestQuery(clientQuery),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Search(context.Background(), clients.NewClientSearchRequest(pagination, clientSort, clientQuery))
				if err != nil {
					t.Fatalf("searching clients: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientListFixture, respJSON)
			},
		},
		{
			name:     "search - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Search(context.Background(), clients.NewClientSearchRequest(pagination, clientSort, clientQuery))
				if err == nil {
					t.Fatal("searching clients: expected error")
				}
				if resp != nil {
					t.Fatal("searching clients: expected nil response")
				}
			},
		},
		{
			name: "create",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients", http.StatusOK,
				testutil.WithResponseBody(clientFixture),
				testutil.WithExpectedRequestBody(clientCreateReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Create(context.Background(), clientCreateReq)
				if err != nil {
					t.Fatalf("creating client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientFixture, respJSON)
			},
		},
		{
			name:     "create - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Create(context.Background(), clientCreateReq)
				if err == nil {
					t.Fatal("creating client: expected error")
				}
				if resp != nil {
					t.Fatal("creating client: expected nil response")
				}
			},
		},
		{
			name: "deactivate",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/deactivate", http.StatusOK,
				testutil.WithResponseBody(clientActionFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Deactivate(context.Background(), 123)
				if err != nil {
					t.Fatalf("deactivating client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientActionFixture, respJSON)
			},
		},
		{
			name:     "deactivate - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/deactivate", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Deactivate(context.Background(), 123)
				if err == nil {
					t.Fatal("deactivating client: expected error")
				}
				if resp != nil {
					t.Fatal("deactivating client: expected nil response")
				}
			},
		},
		{
			name: "reactivate",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/reactivate", http.StatusOK,
				testutil.WithResponseBody(clientActionFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Reactivate(context.Background(), 123)
				if err != nil {
					t.Fatalf("reactivating client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientActionFixture, respJSON)
			},
		},
		{
			name:     "reactivate - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/reactivate", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Reactivate(context.Background(), 123)
				if err == nil {
					t.Fatal("reactivating client: expected error")
				}
				if resp != nil {
					t.Fatal("reactivating client: expected nil response")
				}
			},
		},
		{
			name: "suspend",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/suspend", http.StatusOK,
				testutil.WithResponseBody(clientActionFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Suspend(context.Background(), 123)
				if err != nil {
					t.Fatalf("suspending client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientActionFixture, respJSON)
			},
		},
		{
			name:     "suspend - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/suspend", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Suspend(context.Background(), 123)
				if err == nil {
					t.Fatal("suspend client: expected error")
				}
				if resp != nil {
					t.Fatal("suspend client: expected nil response")
				}
			},
		},
		{
			name: "reinstate",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/reinstate", http.StatusOK,
				testutil.WithResponseBody(clientActionFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Reinstate(context.Background(), 123)
				if err != nil {
					t.Fatalf("reinstating client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientActionFixture, respJSON)
			},
		},
		{
			name:     "reinstate - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/reinstate", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Reinstate(context.Background(), 123)
				if err == nil {
					t.Fatal("reinstating client: expected error")
				}
				if resp != nil {
					t.Fatal("reinstating client: expected nil response")
				}
			},
		},
		{
			name: "update",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123", http.StatusOK,
				testutil.WithResponseBody(clientFixture),
				testutil.WithExpectedRequestBody(clientUpdateReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Update(context.Background(), 123, clientUpdateReq)
				if err != nil {
					t.Fatalf("updating client: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, clientFixture, respJSON)
			},
		},
		{
			name:     "update - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.Update(context.Background(), 123, clientUpdateReq)
				if err == nil {
					t.Fatal("updating client: expected error")
				}
				if resp != nil {
					t.Fatal("updating client: expected nil response")
				}
			},
		},
		{
			name: "list statuses",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/statuses", http.StatusOK,
				testutil.WithResponseBody(statusListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(statusSort),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.ListStatuses(context.Background(), clients.NewStatusListRequest(pagination, statusSort))
				if err != nil {
					t.Fatalf("listing statuses: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, statusListFixture, respJSON)
			},
		},
		{
			name:     "list statuses - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/statuses", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.ListStatuses(context.Background(), clients.NewStatusListRequest(pagination, statusSort))
				if err == nil {
					t.Fatal("listing statuses: expected error")
				}
				if resp != nil {
					t.Fatal("listing statuses: expected nil response")
				}
			},
		},
		{
			name: "search statuses",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/statuses/search", http.StatusOK,
				testutil.WithResponseBody(statusListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(statusSort),
				testutil.WithExpectedRequestQuery(statusQuery),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.SearchStatuses(context.Background(), clients.NewStatusSearchRequest(pagination, statusSort, statusQuery))
				if err != nil {
					t.Fatalf("searching statuses: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, statusListFixture, respJSON)
			},
		},
		{
			name:     "search statuses - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/statuses/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.SearchStatuses(context.Background(), clients.NewStatusSearchRequest(pagination, statusSort, statusQuery))
				if err == nil {
					t.Fatal("searching statuses: expected error")
				}
				if resp != nil {
					t.Fatal("searching statuses: expected nil response")
				}
			},
		},
		{
			name: "list group roles",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/group/roles", http.StatusOK,
				testutil.WithResponseBody(groupRoleListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(groupRoleSort),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.ListGroupRoles(context.Background(), clients.NewGroupRoleListRequest(pagination, groupRoleSort))
				if err != nil {
					t.Fatalf("listing group roles: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupRoleListFixture, respJSON)
			},
		},
		{
			name:     "list group roles - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/group/roles", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.ListGroupRoles(context.Background(), clients.NewGroupRoleListRequest(pagination, groupRoleSort))
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
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/group/roles/search", http.StatusOK,
				testutil.WithResponseBody(groupRoleListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(groupRoleSort),
				testutil.WithExpectedRequestQuery(groupRoleQuery),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.SearchGroupRoles(context.Background(), clients.NewGroupRoleSearchRequest(pagination, groupRoleSort, groupRoleQuery))
				if err != nil {
					t.Fatalf("searching group roles: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, groupRoleListFixture, respJSON)
			},
		},
		{
			name:     "search group roles - error",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/clients/group/roles/search", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.SearchGroupRoles(context.Background(), clients.NewGroupRoleSearchRequest(pagination, groupRoleSort, groupRoleQuery))
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
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients/group/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(groupCreateReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients/group/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/group/123/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(participantsReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/group/123/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/clients/group/123/participants", http.StatusOK,
				testutil.WithResponseBody(groupFixture),
				testutil.WithExpectedRequestBody(participantsReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodDelete, "/clients/group/123/participants", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/convert-from-dependent", http.StatusOK,
				testutil.WithResponseBody(convertDependentFixture),
				testutil.WithExpectedRequestBody(convertDependentReq),
			),
			run: func(t *testing.T, svc *clients.Client) {
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
			endpoint: testutil.NewEndpoint(t, http.MethodPut, "/clients/123/convert-from-dependent", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.ConvertFromDependent(context.Background(), 123, convertDependentReq)
				if err == nil {
					t.Fatalf("converting lead from dependent: expected error")
				}
				if resp != nil {
					t.Fatalf("converting lead from dependent: expected nil response")
				}
			},
		},
		{
			name: "generate register link",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients/123/register-link", http.StatusOK,
				testutil.WithResponseBody(registerLinkFixture),
			),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.GenerateRegisterLink(context.Background(), 123)
				if err != nil {
					t.Fatalf("generating register link: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, registerLinkFixture, respJSON)
			},
		},
		{
			name:     "generate register link - error",
			endpoint: testutil.NewEndpoint(t, http.MethodPost, "/clients/123/register-link", http.StatusBadRequest),
			run: func(t *testing.T, svc *clients.Client) {
				resp, err := svc.GenerateRegisterLink(context.Background(), 123)
				if err == nil {
					t.Fatalf("generating register link: expected error")
				}
				if resp != nil {
					t.Fatalf("generating register link: expected nil response")
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			svr := testutil.NewWodifyClient(t, "test-key", 0, testCase.endpoint)
			svc := clients.New(svr)
			testCase.run(t, svc)
		})
	}
}

func TestClientUpdateRequestFrom(t *testing.T) {
	t.Parallel()
	c := &models.Client{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		PhoneNumber:           "555-867-5309",
		BillingCCEmail:        "billing@example.com",
		ClientStatusID:        10,
		LocationID:            20,
		DefaultProgramID:      30,
		DateOfBirth:           models.NewDate(time.Date(1990, time.June, 15, 0, 0, 0, 0, time.UTC)),
		GenderID:              2,
		StreetAddress1:        "123 Main St",
		StreetAddress2:        "Apt 4B",
		City:                  "Springfield",
		StateID:               40,
		Province:              "Province",
		ZipCode:               "62701",
		CountryID:             50,
		TimezoneID:            60,
		Height1Measurement:    5,
		Height2Measurement:    9,
		Weight:                160.5,
		IsEmailSubscribed:     true,
		IsSMSSubscribed:       false,
		Notes:                 "Some notes",
		EmergencyContactName:  "Mary Doe",
		EmergencyContactPhone: "555-111-2222",
		LeadSourceID:          90,
		ReferringUserID:       100,
		ClientOwnerID:         130,
		CoachTitle:            "Head Coach",
		CoachBio:              "Bio text",
		CoachLink1Icon:        "instagram",
		CoachLink1URL:         "https://instagram.com/johndoe",
		CoachLink2Icon:        "twitter",
		CoachLink2URL:         "https://twitter.com/johndoe",
		CoachLink3Icon:        "facebook",
		CoachLink3URL:         "https://facebook.com/johndoe",
		CoachLink4Icon:        "linkedin",
		CoachLink4URL:         "https://linkedin.com/in/johndoe",
		CoachLink5Icon:        "youtube",
		CoachLink5URL:         "https://youtube.com/johndoe",
	}

	expected := clients.ClientUpdateRequest{
		FirstName:             c.FirstName,
		LastName:              c.LastName,
		Email:                 c.Email,
		PhoneNumber:           c.PhoneNumber,
		BillingCCEmail:        c.BillingCCEmail,
		ClientStatusID:        c.ClientStatusID,
		LocationID:            c.LocationID,
		DefaultProgramID:      c.DefaultProgramID,
		DateOfBirth:           c.DateOfBirth,
		GenderID:              c.GenderID,
		StreetAddress1:        c.StreetAddress1,
		StreetAddress2:        c.StreetAddress2,
		City:                  c.City,
		StateID:               c.StateID,
		Province:              c.Province,
		ZipCode:               c.ZipCode,
		CountryID:             c.CountryID,
		TimezoneID:            c.TimezoneID,
		Height1Measurement:    c.Height1Measurement,
		Height2Measurement:    c.Height2Measurement,
		Weight:                c.Weight,
		IsEmailSubscribed:     c.IsEmailSubscribed,
		IsSMSSubscribed:       c.IsSMSSubscribed,
		Notes:                 c.Notes,
		IsOverwriteNotes:      true,
		EmergencyContactName:  c.EmergencyContactName,
		EmergencyContactPhone: c.EmergencyContactPhone,
		LeadSourceID:          c.LeadSourceID,
		ReferringUserID:       c.ReferringUserID,
		ClientOwnerID:         c.ClientOwnerID,
		CoachTitle:            c.CoachTitle,
		CoachBio:              c.CoachBio,
		CoachLink1Icon:        c.CoachLink1Icon,
		CoachLink1URL:         c.CoachLink1URL,
		CoachLink2Icon:        c.CoachLink2Icon,
		CoachLink2URL:         c.CoachLink2URL,
		CoachLink3Icon:        c.CoachLink3Icon,
		CoachLink3URL:         c.CoachLink3URL,
		CoachLink4Icon:        c.CoachLink4Icon,
		CoachLink4URL:         c.CoachLink4URL,
		CoachLink5Icon:        c.CoachLink5Icon,
		CoachLink5URL:         c.CoachLink5URL,
	}

	actual := clients.ClientUpdateRequestFrom(c)

	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("marshal actual: %v", err)
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("marshal expected: %v", err)
	}
	testutil.AssertJSONEqual(t, expectedJSON, actualJSON)
}
