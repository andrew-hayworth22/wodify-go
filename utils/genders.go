package utils

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

func (c *Client) ListGenders(ctx context.Context, req ListGendersRequest) (*ListGendersResponse, error) {
	var out ListGendersResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/genders", req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// GenderField represents a field that gender lists can be sorted/filtered by.
type GenderField string

const (
	GenderFieldID   GenderField = "gender_id"
	GenderFieldName GenderField = "gender"
)

// GenderSort represents a gender sort order.
type GenderSort = sort.Sort[GenderField]

// NewGenderSort creates a new gender sort.
func NewGenderSort(field GenderField, isDescending bool) GenderSort {
	return sort.NewSort(field, isDescending)
}

// ListGendersRequest represents a request to list genders.
type ListGendersRequest struct {
	Page models.PaginationRequest
	Sort GenderSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListGendersRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListGendersResponse represents a response to a list genders request.
type ListGendersResponse struct {
	Genders    []models.Gender           `json:"genders"`
	Pagination models.PaginationResponse `json:"pagination"`
}
