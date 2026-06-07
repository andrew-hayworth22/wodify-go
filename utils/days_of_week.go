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

// ListDaysOfWeek fetches a list of days of week.
func (c *Client) ListDaysOfWeek(ctx context.Context, req DayOfWeekListRequest) (*DayOfWeekListResponse, error) {
	var out DayOfWeekListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/days-of-week", req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// DayOfWeekField represents a field that day of week lists can be sorted/filtered on.
type DayOfWeekField string

const (
	DayOfWeekFieldID   DayOfWeekField = "day_of_week_id"
	DayOfWeekFieldName DayOfWeekField = "day_of_week"
)

// DayOfWeekListRequest represents a request to list days of week.
type DayOfWeekListRequest = request.ListRequest[DayOfWeekField]

// NewDayOfWeekListRequest creates a new DayOfWeekListRequest with the given pagination and sort.
func NewDayOfWeekListRequest(pagination request.PaginationRequest, sort sort.Sort[DayOfWeekField]) DayOfWeekListRequest {
	return DayOfWeekListRequest{
		Page: pagination,
		Sort: sort,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// DayOfWeekListResponse represents a response to a list days of week request.
type DayOfWeekListResponse struct {
	DaysOfWeek []models.DayOfWeek        `json:"days_of_week"`
	Pagination models.PaginationResponse `json:"pagination"`
}
