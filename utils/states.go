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

// ListStates fetches a list of US states.
func (c *Client) ListStates(ctx context.Context, req StateListRequest) (*StateListResponse, error) {
	var out StateListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/states", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchStates fetches a list of US states matching a query criteria.
func (c *Client) SearchStates(ctx context.Context, req StateSearchRequest) (*StateListResponse, error) {
	var out StateListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/states/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
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

// StateListRequest represents a request to list states.
type StateListRequest = request.ListRequest[StateField]

// NewStateListRequest creates a new StateListRequest with the given pagination and sort.
func NewStateListRequest(pagination request.PaginationRequest, sort sort.Sort[StateField]) StateListRequest {
	return StateListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// StateSearchRequest represents a request to query states.
type StateSearchRequest = request.SearchRequest[StateField]

// NewStateSearchRequest creates a new StateSearchRequest with the given pagination, sort, and query.
func NewStateSearchRequest(pagination request.PaginationRequest, sort sort.Sort[StateField], query *query.Builder[StateField]) StateSearchRequest {
	return StateSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewStateQuery creates a new state query builder.
func NewStateQuery() *query.Builder[StateField] {
	return query.New[StateField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// StateListResponse represents a response to a list states request.
type StateListResponse struct {
	States     []models.State            `json:"states"`
	Pagination models.PaginationResponse `json:"pagination"`
}
