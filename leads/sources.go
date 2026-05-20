package leads

import (
	"context"
	"net/http"
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListSources fetches a list of lead sources
func (s *Client) ListSources(ctx context.Context, req ListSourcesRequest) (*ListSourcesResponse, error) {
	var out ListSourcesResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads/sources", req.ToQuery(), nil, &out)
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

// SourceSort represents a lead source sort order
type SourceSort = sort.Sort[SourceField]

// NewSourceSort creates a new lead source sort
func NewSourceSort(field SourceField, isDescending bool) SourceSort {
	return sort.NewSort(field, isDescending)
}

// ListSourcesRequest represents a request to list lead sources
type ListSourcesRequest struct {
	Page models.PaginationRequest
	Sort SourceSort
}

// ToQuery converts the request to URL query string parameters
func (r ListSourcesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListSourcesResponse represents a response to a lead source fetch
type ListSourcesResponse struct {
	Sources    []models.LeadSource       `json:"sources"`
	Pagination models.PaginationResponse `json:"pagination"`
}
