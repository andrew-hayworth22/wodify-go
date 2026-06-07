package utils_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient(t *testing.T) {
	t.Parallel()

	pagination := wodify.NewPaginationRequest(1, 10)

	countryListFixture := testutil.MustReadJSONFixture(t, "testdata/country_list.json")
	countrySort := wodify.SortAscending(utils.CountryFieldName)
	countryQuery := utils.NewCountryQuery().Eq(utils.CountryFieldID, 123)

	dayOfWeekListFixture := testutil.MustReadJSONFixture(t, "testdata/day_of_week_list.json")
	dayOfWeekSort := wodify.SortAscending(utils.DayOfWeekFieldName)

	genderListFixture := testutil.MustReadJSONFixture(t, "testdata/gender_list.json")
	genderSort := wodify.SortAscending(utils.GenderFieldName)

	objectTypeListFixture := testutil.MustReadJSONFixture(t, "testdata/object_type_list.json")
	objectTypeSort := wodify.SortAscending(utils.ObjectTypeFieldName)
	objectTypeQuery := utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldID, 123)

	objectActionTypeListFixture := testutil.MustReadJSONFixture(t, "testdata/object_action_type_list.json")
	objectActionTypeSort := wodify.SortAscending(utils.ObjectActionTypeFieldName)
	objectActionTypeQuery := utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldID, 123)

	stateListFixture := testutil.MustReadJSONFixture(t, "testdata/state_list.json")
	stateSort := wodify.SortAscending(utils.StateFieldName)
	stateQuery := utils.NewStateQuery().Eq(utils.StateFieldID, 123)

	unitOfTimeListFixture := testutil.MustReadJSONFixture(t, "testdata/unit_of_time_list.json")
	unitOfTimeSort := wodify.SortAscending(utils.UnitOfTimeFieldNameSingular)
	unitOfTimeQuery := utils.NewUnitOfTimeQuery().Eq(utils.UnitOfTimeFieldNameSingular, "Hour")

	testCases := []struct {
		name     string
		endpoint *testutil.Endpoint
		run      func(*testing.T, *utils.Client)
	}{
		{
			name: "list countries",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/countries", http.StatusOK,
				testutil.WithResponseBody(countryListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(countrySort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListCountries(context.Background(), utils.NewCountryListRequest(pagination, countrySort))
				if err != nil {
					t.Fatalf("listing countries: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, countryListFixture, respJSON)
			},
		},
		{
			name: "search countries",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/countries/search", http.StatusOK,
				testutil.WithResponseBody(countryListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(countrySort),
				testutil.WithExpectedRequestQuery(countryQuery),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.SearchCountries(context.Background(), utils.NewCountrySearchRequest(pagination, countrySort, countryQuery))
				if err != nil {
					t.Fatalf("searching countries: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, countryListFixture, respJSON)
			},
		},
		{
			name: "list days of week",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/days-of-week", http.StatusOK,
				testutil.WithResponseBody(dayOfWeekListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(dayOfWeekSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListDaysOfWeek(context.Background(), utils.NewDayOfWeekListRequest(pagination, dayOfWeekSort))
				if err != nil {
					t.Fatalf("listing days of week: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, dayOfWeekListFixture, respJSON)
			},
		},
		{
			name: "list genders",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/genders", http.StatusOK,
				testutil.WithResponseBody(genderListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(genderSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListGenders(context.Background(), utils.NewGenderListRequest(pagination, genderSort))
				if err != nil {
					t.Fatalf("listing genders: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, genderListFixture, respJSON)
			},
		},
		{
			name: "list object types",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/object-types", http.StatusOK,
				testutil.WithResponseBody(objectTypeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(objectTypeSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListObjectTypes(context.Background(), utils.NewObjectTypeListRequest(pagination, objectTypeSort))
				if err != nil {
					t.Fatalf("listing object types: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, objectTypeListFixture, respJSON)
			},
		},
		{
			name: "search object types",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/object-types/search", http.StatusOK,
				testutil.WithResponseBody(objectTypeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(objectTypeSort),
				testutil.WithExpectedRequestQuery(objectTypeQuery),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.SearchObjectTypes(context.Background(), utils.NewObjectTypeSearchRequest(pagination, objectTypeSort, objectTypeQuery))
				if err != nil {
					t.Fatalf("searching object types: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, objectTypeListFixture, respJSON)
			},
		},
		{
			name: "list object action types",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/object-type-action-types", http.StatusOK,
				testutil.WithResponseBody(objectActionTypeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(objectActionTypeSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListObjectActionTypes(context.Background(), utils.NewObjectActionTypeListRequest(pagination, objectActionTypeSort))
				if err != nil {
					t.Fatalf("listing object action types: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, objectActionTypeListFixture, respJSON)
			},
		},
		{
			name: "search object action types",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/object-type-action-types/search", http.StatusOK,
				testutil.WithResponseBody(objectActionTypeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(objectActionTypeSort),
				testutil.WithExpectedRequestQuery(objectActionTypeQuery),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.SearchObjectActionTypes(context.Background(), utils.NewObjectActionTypeSearchRequest(pagination, objectActionTypeSort, objectActionTypeQuery))
				if err != nil {
					t.Fatalf("searching object action types: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, objectActionTypeListFixture, respJSON)
			},
		},
		{
			name: "list states",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/states", http.StatusOK,
				testutil.WithResponseBody(stateListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(stateSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListStates(context.Background(), utils.NewStateListRequest(pagination, stateSort))
				if err != nil {
					t.Fatalf("listing states: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, stateListFixture, respJSON)
			},
		},
		{
			name: "search states",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/states/search", http.StatusOK,
				testutil.WithResponseBody(stateListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(stateSort),
				testutil.WithExpectedRequestQuery(stateQuery),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.SearchStates(context.Background(), utils.NewStateSearchRequest(pagination, stateSort, stateQuery))
				if err != nil {
					t.Fatalf("searching states: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, stateListFixture, respJSON)
			},
		},
		{
			name: "list units of time",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/units-of-time", http.StatusOK,
				testutil.WithResponseBody(unitOfTimeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(unitOfTimeSort),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.ListUnitsOfTime(context.Background(), utils.NewUnitOfTimeListRequest(pagination, unitOfTimeSort))
				if err != nil {
					t.Fatalf("listing units of time: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, unitOfTimeListFixture, respJSON)
			},
		},
		{
			name: "search units of time",
			endpoint: testutil.NewEndpoint(t, http.MethodGet, "/utilities/units-of-time/search", http.StatusOK,
				testutil.WithResponseBody(unitOfTimeListFixture),
				testutil.WithExpectedRequestPagination(pagination),
				testutil.WithExpectedRequestSort(unitOfTimeSort),
				testutil.WithExpectedRequestQuery(unitOfTimeQuery),
			),
			run: func(t *testing.T, svc *utils.Client) {
				resp, err := svc.SearchUnitsOfTime(context.Background(), utils.NewUnitOfTimeSearchRequest(pagination, unitOfTimeSort, unitOfTimeQuery))
				if err != nil {
					t.Fatalf("searching units of time: %v", err)
				}
				respJSON, _ := json.Marshal(resp)
				testutil.AssertJSONEqual(t, unitOfTimeListFixture, respJSON)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			svr := testutil.NewWodifyClient(t, "test-key", 0, testCase.endpoint)
			svc := utils.New(svr)
			testCase.run(t, svc)
		})
	}
}
