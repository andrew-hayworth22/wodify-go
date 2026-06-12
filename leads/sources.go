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

// ListSources fetches a list of lead sources
func (c *Client) ListSources(ctx context.Context, req SourceListRequest) (*SourceListResponse, error) {
	var out SourceListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/sources", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// SourceField represents a field that lead sources can be sorted/filtered on
type SourceField string

const (
	SourceFieldID   SourceField = "id"
	SourceFieldName SourceField = "source"
)

// SourceListRequest represents a request to list lead sources
type SourceListRequest = request.ListRequest[SourceField]

// NewSourceListRequest creates a new SourceListRequest with the given pagination and sort.
func NewSourceListRequest(pagination request.PaginationRequest, sort sort.Sort[SourceField]) SourceListRequest {
	return SourceListRequest{
		Page: pagination,
		Sort: sort,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// SourceListResponse represents a response to a lead source fetch
type SourceListResponse struct {
	Sources    []models.LeadSource       `json:"sources"`
	Pagination models.PaginationResponse `json:"pagination"`
}
