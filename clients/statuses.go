package clients

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

// ListStatuses fetches a list of client statuses
func (c *Client) ListStatuses(ctx context.Context, req StatusListRequest) (*StatusListResponse, error) {
	var out StatusListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/clients/statuses", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchStatuses fetches a list of client statuses matching the query criteria
func (c *Client) SearchStatuses(ctx context.Context, req StatusSearchRequest) (*StatusListResponse, error) {
	if req.Query != nil {
		if err := req.Query.Err(); err != nil {
			return nil, err
		}
	}
	var out StatusListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/clients/statuses/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// ClientStatusField represents a field that client statuses can be sorted/filtered on.
type ClientStatusField string

const (
	StatusFieldID   ClientStatusField = "id"
	StatusFieldName ClientStatusField = "status"
)

// StatusListRequest represents a request to list client statuses
type StatusListRequest = request.ListRequest[ClientStatusField]

// NewStatusListRequest creates a new request to list client statuses
func NewStatusListRequest(pagination request.PaginationRequest, sort sort.Sort[ClientStatusField]) StatusListRequest {
	return StatusListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// StatusSearchRequest represents a request to search client statuses
type StatusSearchRequest = request.SearchRequest[ClientStatusField]

// NewStatusSearchRequest creates a new request to search client statuses
func NewStatusSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ClientStatusField], query *query.Builder[ClientStatusField]) StatusSearchRequest {
	return StatusSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewStatusQuery creates a new query builder for client statuses
func NewStatusQuery() *query.Builder[ClientStatusField] {
	return query.New[ClientStatusField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// StatusListResponse represents a response to a request to list client statuses
type StatusListResponse struct {
	Statuses   []models.ClientStatus     `json:"statuses"`
	Pagination models.PaginationResponse `json:"pagination"`
}
