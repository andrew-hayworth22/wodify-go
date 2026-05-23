package utils

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

func (c *Client) ListGenders(ctx context.Context, req GenderListRequest) (*GenderListResponse, error) {
	var out GenderListResponse
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

// GenderListRequest represents a request to list genders.
type GenderListRequest = request.ListRequest[GenderField]

// NewGenderListRequest creates a new GenderListRequest with the given pagination and sort.
func NewGenderListRequest(pagination request.PaginationRequest, sort sort.Sort[GenderField]) GenderListRequest {
	return GenderListRequest{
		Page: pagination,
		Sort: sort,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// GenderListResponse represents a response to a list genders request.
type GenderListResponse struct {
	Genders    []models.Gender           `json:"genders"`
	Pagination models.PaginationResponse `json:"pagination"`
}
