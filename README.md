# wodify-go

[![Go Reference](https://pkg.go.dev/badge/github.com/andrew-hayworth22/wodify-go.svg)](https://pkg.go.dev/github.com/andrew-hayworth22/wodify-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/andrew-hayworth22/wodify-go)](https://goreportcard.com/report/github.com/andrew-hayworth22/wodify-go)
[![Coverage](https://codecov.io/gh/andrew-hayworth22/wodify-go/branch/main/graph/badge.svg)](https://codecov.io/gh/andrew-hayworth22/wodify-go)


A Go SDK for the [Wodify](https://docs.wodify.com/reference) API.

> **Note:** wodify-go is a community-maintained SDK and is not officially supported by Wodify.

## Installation

```sh
go get github.com/andrew-hayworth22/wodify-go
```

## Authentication

An API key is required. Provide it via environment variable or the `WithAPIKey` option:

```sh
export WODIFY_API_KEY=your_api_key
```

## Quick Start

```go
client, err := wodify.New()
if err != nil {
    log.Fatal(err)
}

lead, err := client.Leads.Get(ctx, 12345)
```

## Domain Coverage

| Domain | Status |
|---|---|
| Utils | ✅ complete |
| Leads | ✅ complete |
| Clients | 🚧 in-development |
| Customers | ⏳ coming soon |
| Documents | ⏳ coming soon |
| Classes | ⏳ coming soon |
| Appointments | ⏳ coming soon |
| Memberships | ⏳ coming soon |
| Financials | ⏳ coming soon |
| Communications | ⏳ coming soon |
| Workouts | ⏳ coming soon |

## Configuration

`wodify.New()` accepts functional options to override defaults:

| Option | Default | Description |
|---|---|---|
| `WithAPIKey(key)` | `$WODIFY_API_KEY` | API key |
| `WithBaseURL(url)` | `$WODIFY_BASE_URL` | API base URL |
| `WithTimeout(d)` | `10s` | Request timeout |
| `WithMaxRetries(n)` | `$WODIFY_MAX_RETRIES` / `3` | Max retries on rate limit |
| `WithHTTPClient(c)` | built-in | Custom `*http.Client` |

```go
client, err := wodify.New(
    wodify.WithAPIKey("your_api_key"),
    wodify.WithTimeout(30 * time.Second),
    wodify.WithMaxRetries(5),
)
```

## Error Handling

Errors can be inspected using the exported sentinel values:

```go
_, err := client.Leads.Get(ctx, id)
if errors.Is(err, wodify.ErrNotFound) {
    // lead does not exist
}
```

| Sentinel | HTTP Status |
|---|---|
| `wodify.ErrNotFound` | 404 |
| `wodify.ErrUnauthorized` | 401, 403 |
| `wodify.ErrRateLimited` | 429 |
| `wodify.ErrBadRequest` | 400, 422 |

For full error details, use `errors.As`:

```go
var apiErr *wodify.APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.HTTPCode, apiErr.MoreInfo, apiErr.DeveloperMessage)
}
```

## Pagination and Sorting

Pagination and sorting helpers are available on the top-level `wodify` package:

```go
p := wodify.NewPaginationRequest(1, 10)
s := wodify.SortAscending(leads.LeadFieldFirstName)
```

## Utils

Reference data for use with other API operations, including genders, countries, states, days of week, object types, and object action types.

```go
// List genders
p := wodify.NewPaginationRequest(1, 10)
genders, err := client.Utils.ListGenders(ctx, utils.NewGenderListRequest(p, wodify.SortAscending(utils.GenderFieldID)))

// List countries
countries, err := client.Utils.ListCountries(ctx, utils.NewCountryListRequest(p, wodify.SortAscending(utils.CountryFieldName)))

// Search countries
q := utils.NewCountryQuery().Eq(utils.CountryFieldName, "United States")
countries, err := client.Utils.SearchCountries(ctx, utils.NewCountrySearchRequest(p, wodify.SortAscending(utils.CountryFieldName), q))

// List states
states, err := client.Utils.ListStates(ctx, utils.NewStateListRequest(p, wodify.SortAscending(utils.StateFieldName)))

// Search states
q := utils.NewStateQuery().Eq(utils.StateFieldName, "California")
states, err := client.Utils.SearchStates(ctx, utils.NewStateSearchRequest(p, wodify.SortAscending(utils.StateFieldName), q))

// List days of week
days, err := client.Utils.ListDaysOfWeek(ctx, utils.NewDayOfWeekListRequest(p, wodify.SortAscending(utils.DayOfWeekFieldID)))

// List object types
objectTypes, err := client.Utils.ListObjectTypes(ctx, utils.NewObjectTypeListRequest(p, wodify.SortAscending(utils.ObjectTypeFieldID)))

// Search object types
q := utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldName, "Lead")
objectTypes, err := client.Utils.SearchObjectTypes(ctx, utils.NewObjectTypeSearchRequest(p, wodify.SortAscending(utils.ObjectTypeFieldID), q))

// List object action types
actionTypes, err := client.Utils.ListObjectActionTypes(ctx, utils.NewObjectActionTypeListRequest(p, wodify.SortAscending(utils.ObjectActionTypeFieldID)))

// Search object action types
q := utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldObjectTypeID, "1")
actionTypes, err := client.Utils.SearchObjectActionTypes(ctx, utils.NewObjectActionTypeSearchRequest(p, wodify.SortAscending(utils.ObjectActionTypeFieldID), q))

// List units of time
unitsOfTime, err := client.Utils.ListUnitsOfTime(ctx, utils.NewUnitOfTimeListRequest(p, wodify.SortAscending(utils.UnitOfTimeFieldID)))

// Search units of time
q := utils.NewUnitOfTimeQuery().Eq(utils.UnitOfTimeFieldID, "1")
unitsOfTime, err := client.Utils.SearchUnitsOfTime(ctx, utils.NewUnitOfTimeSearchRequest(p, wodify.SortAscending(utils.UnitOfTimeFieldID), q))
```

## Leads

Lead management, including CRUD operations, conversion to clients, statuses, sources, tags, appointment bookings, class sign-ins, class reservations, and performance results.

```go
p := wodify.NewPaginationRequest(1, 10)

// Get a lead by ID
lead, err := client.Leads.Get(ctx, id)

// Create a lead
lead, err := client.Leads.Create(ctx, leads.LeadCreateRequest{
    FirstName:  "Jane",
    LastName:   "Doe",
    Email:      "jane@example.com",
    LocationID: 11337,
    GenderID:   2,
})

// List leads
results, err := client.Leads.List(ctx, leads.NewLeadListRequest(p, wodify.SortAscending(leads.LeadFieldFirstName)))

// Search leads
q := leads.NewLeadQuery().Eq(leads.LeadFieldFirstName, "Jane")
results, err := client.Leads.Search(ctx, leads.NewLeadSearchRequest(p, wodify.SortAscending(leads.LeadFieldFirstName), q))

// Update a lead
req := leads.LeadUpdateRequestFrom(lead)
req.FirstName = "Janet"
lead, err = client.Leads.Update(ctx, lead.ID, req)

// Delete a lead
res, err := client.Leads.Delete(ctx, id)

// Convert a lead to a client
req := leads.LeadConvertRequestFrom(lead)
req.ClientStatusID = 1
res, err := client.Leads.Convert(ctx, id, req)

// List lead statuses
statuses, err := client.Leads.ListStatuses(ctx, leads.NewStatusListRequest(p, wodify.SortAscending(leads.StatusFieldID)))

// List lead sources
sources, err := client.Leads.ListSources(ctx, leads.NewSourceListRequest(p, wodify.SortAscending(leads.SourceFieldName)))

// Add tags to a lead
res, err := client.Leads.AddTags(ctx, id, leads.TagsUpdateRequest{
    Tags: []string{"vip", "trial"},
})

// Delete tags from a lead
res, err := client.Leads.DeleteTags(ctx, id, leads.TagsUpdateRequest{
    Tags: []string{"trial"},
})

// List a lead's appointment bookings
bookings, err := client.Leads.ListBookings(ctx, id, leads.NewBookingListRequest(p, wodify.SortAscending(leads.BookingFieldID)))

// Search a lead's appointment bookings
q := leads.NewBookingQuery().Eq(leads.BookingFieldStatusID, "2")
bookings, err := client.Leads.SearchBookings(ctx, id, leads.NewBookingSearchRequest(p, wodify.SortAscending(leads.BookingFieldID), q))

// List a lead's class sign-ins
signIns, err := client.Leads.ListClassSignIns(ctx, id, leads.NewClassSignInListRequest(p, wodify.SortAscending(leads.ClassSignInFieldID)))

// Search a lead's class sign-ins
q := leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldProgramID, "5")
signIns, err := client.Leads.SearchClassSignIns(ctx, id, leads.NewClassSignInSearchRequest(p, wodify.SortAscending(leads.ClassSignInFieldID), q))

// List a lead's class reservations
reservations, err := client.Leads.ListReservations(ctx, id, leads.NewReservationListRequest(p, wodify.SortAscending(leads.ReservationFieldID)))

// Search a lead's class reservations
q := leads.NewReservationQuery().Eq(leads.ReservationFieldStatusID, "1")
reservations, err := client.Leads.SearchReservations(ctx, id, leads.NewReservationSearchRequest(p, wodify.SortAscending(leads.ReservationFieldID), q))

// List a lead's performance results
results, err := client.Leads.ListPerformanceResults(ctx, id, leads.NewPerformanceResultListRequest(p))

// List a lead's performance results for a specific component
results, err := client.Leads.ListPerformanceResultsByComponent(ctx, id, componentID, leads.NewPerformanceResultListRequest(p))

// List lead group roles
roles, err := client.Leads.ListGroupRoles(ctx, leads.NewGroupRoleListRequest(p, wodify.SortAscending(leads.GroupRoleFieldName)))

// Search lead group roles
q := leads.NewGroupRoleQuery().Eq(leads.GroupRoleFieldName, "Guardian")
roles, err := client.Leads.SearchGroupRoles(ctx, leads.NewGroupRoleSearchRequest(p, wodify.SortAscending(leads.GroupRoleFieldID), q))
```

## Clients

Client management, including CRUD operations and status changes.

```go
p := wodify.NewPaginationRequest(1, 10)

// Get a client by ID
client, err := client.Clients.Get(ctx, id)

// Create a client
c, err := client.Clients.Create(ctx, clients.ClientCreateRequest{
    FirstName:      "Jane",
    LastName:       "Doe",
    Email:          "jane@example.com",
    LocationID:     11337,
    ClientStatusId: 1,
    GenderID:       2,
})

// List clients
results, err := client.Clients.List(ctx, clients.NewClientListRequest(p, wodify.SortAscending(clients.ClientFieldFirstName)))

// Search clients
q := clients.NewClientQuery().Eq(clients.ClientFieldFirstName, "Jane")
results, err := client.Clients.Search(ctx, clients.NewClientSearchRequest(p, wodify.SortAscending(clients.ClientFieldFirstName), q))

// Update a client
req := clients.ClientUpdateRequestFrom(c)
req.FirstName = "Janet"
c, err = client.Clients.Update(ctx, c.ID, req)

// Deactivate a client
res, err := client.Clients.Deactivate(ctx, id)

// Reactivate a client
res, err := client.Clients.Reactivate(ctx, id)

// Suspend a client
res, err := client.Clients.Suspend(ctx, id)

// Reinstate a client
res, err := client.Clients.Reinstate(ctx, id)

// List client statuses
statuses, err := client.Clients.ListStatuses(ctx, clients.NewStatusListRequest(p, wodify.SortAscending(clients.StatusFieldName)))

// Search client statuses
q := clients.NewStatusQuery().Eq(clients.StatusFieldName, "Active")
statuses, err := client.Clients.SearchStatuses(ctx, clients.NewStatusSearchRequest(p, wodify.SortAscending(clients.StatusFieldName), q))
```

## Examples

```sh
# Listing genders
make utils-genders

# Listing and searching countries
make utils-countries

# Listing and searching states
make utils-states

# Listing and searching object types and action types
make utils-object-types

# Listing and searching units of time
make utils-units-of-time

# Listing days of week
make utils-days-of-week

# CRUD operations on leads
make leads-crud

# Listing and searching leads
make leads-search

# Converting a lead to a client
make leads-convert

# Listing lead statuses
make leads-statuses

# Listing lead sources
make leads-sources

# Adding and deleting lead tags
make leads-tags

# Listing and searching lead appointment bookings
make leads-bookings

# Listing and searching lead class sign-ins
make leads-class-sign-ins

# Listing and searching lead class reservations
make leads-reservations

# Listing lead performance results
make leads-performance-results

# Listing and searching lead group roles
make leads-group-roles

# CRUD operations on clients
make clients-crud

# Searching clients
make clients-search

# Deactivating, reactivating, suspending, and reinstating clients
make clients-actions

# Listing and searching client statuses
make clients-statuses
```


## Testing

```sh
# Run all tests
make test

# Run tests for a specific package or group of packages
make test-wodify
make test-models
make test-internal
make test-utils
make test-leads
make test-clients
```

## Profiling

```sh
# Coverage profile (generates coverage.out and opens HTML report)
make profile-test
```

## License

MIT