package utils

import (
	"context"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListCountries fetches a list of countries.
func (c *Client) ListCountries(ctx context.Context, req CountryListRequest) (*CountryListResponse, error) {
	var out CountryListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/countries", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchCountries fetches a list of countries matching a query criteria.
func (c *Client) SearchCountries(ctx context.Context, req CountrySearchRequest) (*CountryListResponse, error) {
	var out CountryListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/countries/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// CountryField represents a field that country lists can be sorted/filtered on.
type CountryField string

const (
	CountryFieldID   CountryField = "country_id"
	CountryFieldName CountryField = "country"
)

// CountryListRequest represents a request to list countries.
type CountryListRequest = request.ListRequest[CountryField]

// NewCountryListRequest creates a new CountryListRequest with the given pagination and sort.
func NewCountryListRequest(pagination request.PaginationRequest, sort sort.Sort[CountryField]) CountryListRequest {
	return CountryListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// CountrySearchRequest represents a request to query for countries.
type CountrySearchRequest = request.SearchRequest[CountryField]

// NewCountrySearchRequest creates a new CountrySearchRequest with the given pagination, sort, and query.
func NewCountrySearchRequest(pagination request.PaginationRequest, sort sort.Sort[CountryField], query *query.Builder[CountryField]) CountrySearchRequest {
	return CountrySearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewCountryQuery creates a new country query  builder.
func NewCountryQuery() *query.Builder[CountryField] {
	return query.New[CountryField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// CountryListResponse represents a response to a list countries request.
type CountryListResponse struct {
	Countries  []models.Country          `json:"countries"`
	Pagination models.PaginationResponse `json:"pagination"`
}
