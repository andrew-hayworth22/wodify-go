package leads

import (
	"context"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListStatuses fetches a list of lead statuses
func (c *Client) ListStatuses(ctx context.Context, req StatusListRequest) (*StatusListResponse, error) {
	var out StatusListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/statuses", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// StatusField represents a field that lead statuses can be sorted/filtered on
type StatusField string

const (
	StatusFieldID   StatusField = "id"
	StatusFieldName StatusField = "status"
)

// StatusListRequest represents a request to list lead statuses
type StatusListRequest = request.ListRequest[StatusField]

// NewStatusListRequest creates a new StatusListRequest with the given pagination and sort.
func NewStatusListRequest(pagination request.PaginationRequest, sort sort.Sort[StatusField]) StatusListRequest {
	return StatusListRequest{
		Page: pagination,
		Sort: sort,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// StatusListResponse represents a response to a lead status fetch
type StatusListResponse struct {
	Statuses   []models.LeadStatus       `json:"statuses"`
	Pagination models.PaginationResponse `json:"pagination"`
}
