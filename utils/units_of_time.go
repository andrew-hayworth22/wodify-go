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

// ListUnitsOfTime fetches a list of unit of time.
func (c *Client) ListUnitsOfTime(ctx context.Context, req UnitOfTimeListRequest) (*UnitOfTimeListResponse, error) {
	var out UnitOfTimeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/units-of-time", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchUnitsOfTime fetches a list of unit of time matching a query criteria.
func (c *Client) SearchUnitsOfTime(ctx context.Context, req UnitOfTimeSearchRequest) (*UnitOfTimeListResponse, error) {
	var out UnitOfTimeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/units-of-time/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// UnitOfTimeField represents a field that unit of time lists can be sorted/filtered by.
type UnitOfTimeField string

const (
	UnitOfTimeFieldID                 UnitOfTimeField = "unit_of_time_id"
	UnitOfTimeFieldNameSingular       UnitOfTimeField = "unit_of_time_singular"
	UnitOfTimeFieldNamePlural         UnitOfTimeField = "unit_of_time_plural"
	UnitOfTimeFieldNameSingularPlural UnitOfTimeField = "unit_of_time_singular_plural"
	UnitOfTimeFieldIsTimeUnit         UnitOfTimeField = "is_time_unit"
	UnitOfTimeFieldNumberOfDays       UnitOfTimeField = "number_of_days"
)

// UnitOfTimeListRequest represents a request to list unit of time.
type UnitOfTimeListRequest = request.ListRequest[UnitOfTimeField]

// NewUnitOfTimeListRequest creates a new UnitOfTimeListRequest with the given pagination and sort.
func NewUnitOfTimeListRequest(pagination request.PaginationRequest, sort sort.Sort[UnitOfTimeField]) UnitOfTimeListRequest {
	return UnitOfTimeListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// UnitOfTimeSearchRequest represents a request to query unit of time.
type UnitOfTimeSearchRequest = request.SearchRequest[UnitOfTimeField]

// NewUnitOfTimeSearchRequest creates a new UnitOfTimeSearchRequest with the given pagination, sort, and query.
func NewUnitOfTimeSearchRequest(pagination request.PaginationRequest, sort sort.Sort[UnitOfTimeField], query *query.Builder[UnitOfTimeField]) UnitOfTimeSearchRequest {
	return UnitOfTimeSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewUnitOfTimeQuery creates a new unit of time query builder.
func NewUnitOfTimeQuery() *query.Builder[UnitOfTimeField] {
	return query.New[UnitOfTimeField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// UnitOfTimeListResponse represents a response to a list unit of time request.
type UnitOfTimeListResponse struct {
	UnitsOfTime []models.UnitOfTime       `json:"units_of_time"`
	Pagination  models.PaginationResponse `json:"pagination"`
}
