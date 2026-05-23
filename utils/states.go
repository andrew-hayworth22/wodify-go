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

// ListStates fetches a list of US states.
func (c *Client) ListStates(ctx context.Context, req ListStatesRequest) (*ListStatesResponse, error) {
	var out ListStatesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/states", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchStates fetches a list of US states matching a search criteria.
func (c *Client) SearchStates(ctx context.Context, req SearchStatesRequest) (*ListStatesResponse, error) {
	var out ListStatesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/states/search", req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// StateField represents a field that state lists can be sorted/filtered by.
type StateField string

const (
	StateFieldID   StateField = "state_id"
	StateFieldName StateField = "state"
)

// StateSort represents a state sort order.
type StateSort = sort.Sort[StateField]

// NewStateSort creates a new state sort.
func NewStateSort(field StateField, isDescending bool) StateSort {
	return sort.NewSort(field, isDescending)
}

// StateQuery represents a state search query.
type StateQuery = search.Builder[StateField]

// NewStateQuery creates a new state search query builder.
func NewStateQuery() *StateQuery {
	return search.New[StateField]()
}

// ListStatesRequest represents a request to list states.
type ListStatesRequest struct {
	Page models.PaginationRequest
	Sort StateSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListStatesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchStatesRequest represents a request to search states.
type SearchStatesRequest struct {
	Page  models.PaginationRequest
	Sort  StateSort
	Query *StateQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchStatesRequest) ToQuery() url.Values {
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

// ListStatesResponse represents a response to a list states request.
type ListStatesResponse struct {
	States     []models.State            `json:"states"`
	Pagination models.PaginationResponse `json:"pagination"`
}
