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

// ListDaysOfWeek fetches a list of days of week.
func (c *Client) ListDaysOfWeek(ctx context.Context, req ListDaysOfWeekRequest) (*ListDaysOfWeekResponse, error) {
	var out ListDaysOfWeekResponse
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

// DayOfWeekSort represents a day of week sort order.
type DayOfWeekSort = sort.Sort[DayOfWeekField]

// NewDayOfWeekSort creates a new day of week sort.
func NewDayOfWeekSort(field DayOfWeekField, isDescending bool) DayOfWeekSort {
	return sort.NewSort(field, isDescending)
}

// ListDaysOfWeekRequest represents a request to list days of week.
type ListDaysOfWeekRequest struct {
	Page models.PaginationRequest
	Sort DayOfWeekSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListDaysOfWeekRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListDaysOfWeekResponse represents a response to a list days of week request.
type ListDaysOfWeekResponse struct {
	DaysOfWeek []models.DayOfWeek `json:"days_of_week"`
	Pagination models.PaginationResponse
}
