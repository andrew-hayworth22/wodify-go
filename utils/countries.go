package utils

import (
	"context"
	"net/http"
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListCountries fetches a list of countries.
func (c *Client) ListCountries(ctx context.Context, req ListCountriesRequest) (*ListCountriesResponse, error) {
	var out ListCountriesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/countries", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchCountries fetches a list of countries matching a search criteria.
func (c *Client) SearchCountries(ctx context.Context, req SearchCountriesRequest) (*ListCountriesResponse, error) {
	var out ListCountriesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/countries/search", req.ToQuery(), nil, &out)
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

// CountrySort represents a country sort order.
type CountrySort = sort.Sort[CountryField]

// NewCountrySort creates a new country sort.
func NewCountrySort(field CountryField, isDescending bool) CountrySort {
	return sort.NewSort(field, isDescending)
}

// CountryQuery represents a country search query.
type CountryQuery = search.Builder[CountryField]

// NewCountryQuery creates a new country search query builder.
func NewCountryQuery() *CountryQuery {
	return search.New[CountryField]()
}

// ListCountriesRequest represents a request to list countries.
type ListCountriesRequest struct {
	Page models.PaginationRequest
	Sort CountrySort
}

// ToQuery converts the request to URL query string parameters.
func (r ListCountriesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchCountriesRequest represents a request to query for countries.
type SearchCountriesRequest struct {
	Page  models.PaginationRequest
	Sort  CountrySort
	Query *CountryQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchCountriesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	if r.Query != nil {
		q.Set("q", r.Query.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListCountriesResponse represents a response to a list countries request.
type ListCountriesResponse struct {
	Countries  []models.Country          `json:"countries"`
	Pagination models.PaginationResponse `json:"pagination"`
}
