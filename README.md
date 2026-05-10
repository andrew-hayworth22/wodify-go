# wodify-go

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

## Leads

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
    Sort: leads.NewSort(leads.SortByFirstName, false),
})

// Search leads
results, err := client.Leads.Search(ctx, leads.SearchRequest{
    Page:  models.PaginationRequest{Page: 1, PageSize: 10},
    Query: leads.NewQuery().Eq(leads.FilterByFirstName, "Jane"),
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
    Sort: leads.NewStatusSort(leads.FieldStatusID, false),
})
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

## Examples

```sh
# CRUD operations on leads
make leads-crud

# Listing and searching leads
make leads-search

# Listing lead statuses
make leads-statuses
```

## Testing

```sh
# Run all tests
make test

# Run tests for a specific package
make test-leads
make test-search
```

## Profiling

```sh
# Coverage profile (generates coverage.out and opens HTML report)
make profile-test
```

## License

MIT