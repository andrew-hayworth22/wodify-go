package leads

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// StatusField represents a field that lead statuses can be sorted/filtered on
type StatusField string

const (
	StatusFieldID   StatusField = "id"
	StatusFieldName StatusField = "status"
)

// StatusSort represents a lead status sort order
type StatusSort = sort.Sort[StatusField]

// NewStatusSort creates a new lead status sort
func NewStatusSort(field StatusField, isDescending bool) StatusSort {
	return sort.NewSort(field, isDescending)
}

// ListStatusesRequest represents a request to list lead statuses
type ListStatusesRequest struct {
	Page models.PaginationRequest
	Sort StatusSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListStatusesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListStatusesResponse represents a response to a lead status fetch
type ListStatusesResponse struct {
	Statuses   []models.LeadStatus       `json:"statuses"`
	Pagination models.PaginationResponse `json:"pagination"`
}
