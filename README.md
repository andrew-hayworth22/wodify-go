# wodify-go

[![Go Reference](https://pkg.go.dev/badge/github.com/andrew-hayworth22/wodify-go.svg)](https://pkg.go.dev/github.com/andrew-hayworth22/wodify-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/andrew-hayworth22/wodify-go)](https://goreportcard.com/report/github.com/andrew-hayworth22/wodify-go)
[![Coverage](https://codecov.io/gh/andrew-hayworth22/wodify-go/branch/main/graph/badge.svg)](https://codecov.io/gh/andrew-hayworth22/wodify-go)


A Go SDK for the [Wodify](https://docs.wodify.com/reference) API.

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

## Utils

Reference data for use with other API operations, including genders, countries, states, days of week, object types, and object action types.

```go
// List genders
genders, err := client.Utils.ListGenders(ctx, utils.ListGendersRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewGenderSort(utils.GenderFieldID, false),
})

// List countries
countries, err := client.Utils.ListCountries(ctx, utils.ListCountriesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewCountrySort(utils.CountryFieldName, false),
})

// Search countries
countries, err := client.Utils.SearchCountries(ctx, utils.SearchCountriesRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: utils.NewCountryQuery().Eq(utils.CountryFieldName, "United States"),
})

// List states
states, err := client.Utils.ListStates(ctx, utils.ListStatesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewStateSort(utils.StateFieldName, false),
})

// Search states
states, err := client.Utils.SearchStates(ctx, utils.SearchStatesRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: utils.NewStateQuery().Eq(utils.StateFieldName, "California"),
})

// List days of week
days, err := client.Utils.ListDaysOfWeek(ctx, utils.ListDaysOfWeekRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewDayOfWeekSort(utils.DayOfWeekFieldID, false),
})

// List object types
objectTypes, err := client.Utils.ListObjectTypes(ctx, utils.ListObjectTypesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewObjectTypeSort(utils.ObjectTypeFieldID, false),
})

// Search object types
objectTypes, err := client.Utils.SearchObjectTypes(ctx, utils.SearchObjectTypesRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldName, "Lead"),
})

// List object action types
actionTypes, err := client.Utils.ListObjectActionTypes(ctx, utils.ListObjectActionTypesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: utils.NewObjectActionTypeSort(utils.ObjectActionTypeFieldID, false),
})

// Search object action types
actionTypes, err := client.Utils.SearchObjectActionTypes(ctx, utils.SearchObjectActionTypesRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldObjectTypeID, 1),
})
```

## Leads

Lead management, including CRUD operations, conversion to clients, statuses, sources, tags, appointment bookings, class sign-ins, class reservations, and performance results.

```go
// Get a lead by ID
lead, err := client.Leads.Get(ctx, id)

// Create a lead
lead, err := client.Leads.Create(ctx, leads.CreateLeadRequest{
    FirstName:   "Jane",
    LastName:    "Doe",
    Email:       "jane@example.com",
    LocationID:  11337,
    Gender:      models.Genders.Female,
    DateOfBirth: models.Date{Time: time.Now().AddDate(-30, 0, 0)},
})

// List leads
results, err := client.Leads.List(ctx, leads.ListRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewLeadSort(leads.LeadFieldFirstName, false),
})

// Search leads
results, err := client.Leads.Search(ctx, leads.SearchRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: leads.NewLeadQuery().Eq(leads.LeadFieldFirstName, "Jane"),
})

// Update a lead
req := lead.ToUpdateRequest()
req.FirstName = "Janet"
lead, err = client.Leads.Update(ctx, lead.ID, req)

// Delete a lead
res, err := client.Leads.Delete(ctx, id)

// Convert a lead to a client
res, err := client.Leads.Convert(ctx, id, req)

// List lead statuses
statuses, err := client.Leads.ListStatuses(ctx, leads.ListStatusesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewStatusSort(leads.StatusFieldID, false),
})

// List lead sources
sources, err := client.Leads.ListSources(ctx, leads.ListSourcesRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewSourceSort(leads.SourceFieldName, false),
})

// Add tags to a lead
res, err := client.Leads.AddTags(ctx, id, leads.UpdateTagsRequest{
    Tags: []string{"vip", "trial"},
})

// Delete tags from a lead
res, err := client.Leads.DeleteTags(ctx, id, leads.UpdateTagsRequest{
    Tags: []string{"trial"},
})

// List a lead's appointment bookings
bookings, err := client.Leads.ListBookings(ctx, id, leads.ListBookingsRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewBookingSort(leads.BookingFieldID, false),
})

// Search a lead's appointment bookings
bookings, err := client.Leads.SearchBookings(ctx, id, leads.SearchBookingsRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: leads.NewBookingQuery().Eq(leads.BookingFieldStatusID, 2),
})

// List a lead's class sign-ins
signIns, err := client.Leads.ListClassSignIns(ctx, id, leads.ListClassSignInsRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewClassSignInSort(leads.ClassSignInFieldID, false),
})

// Search a lead's class sign-ins
signIns, err := client.Leads.SearchClassSignIns(ctx, id, leads.SearchClassSignInsRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldProgramID, 5),
})

// List a lead's class reservations
reservations, err := client.Leads.ListReservations(ctx, id, leads.ListReservationsRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
    Sort: leads.NewReservationSort(leads.ReservationFieldID, false),
})

// Search a lead's class reservations
reservations, err := client.Leads.SearchReservations(ctx, id, leads.SearchReservationsRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: leads.NewReservationQuery().Eq(leads.ReservationFieldStatusID, 1),
})

// List a lead's performance results
results, err := client.Leads.ListPerformanceResults(ctx, id, leads.ListPerformanceResultsRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
})

// List a lead's performance results for a specific component
results, err := client.Leads.ListPerformanceResultsByComponent(ctx, id, componentID, leads.ListPerformanceResultsRequest{
    Page: models.PaginationRequest{Page: 1, PageSize: 10},
})
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
```


## Testing

```sh
# Run all tests
make test

# Run tests for a specific package or group of packages
make test-wodify
make test-internal
make test-utils
make test-leads
```

## Profiling

```sh
# Coverage profile (generates coverage.out and opens HTML report)
make profile-test
```

## License

MIT