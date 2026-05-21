package leads

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListPerformanceResults fetches a list of a Lead's performance results
func (c *Client) ListPerformanceResults(ctx context.Context, id int64, req ListPerformanceResultsRequest) (*ListPerformanceResultsResponse, error) {
	var out ListPerformanceResultsResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/performance-results", id), req.ToQuery(), nil, &out)
	return &out, err
}

// ListPerformanceResultsByComponent fetches a list of a Lead's performance results for a specific component'
func (c *Client) ListPerformanceResultsByComponent(ctx context.Context, leadID int64, componentID int64, req ListPerformanceResultsRequest) (*ListPerformanceResultsResponse, error) {
	var out ListPerformanceResultsResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/performance-results/components/%d", leadID, componentID), req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

type ListPerformanceResultsRequest struct {
	Page models.PaginationRequest
}

// ToQuery converts the request to URL query string parameters.
func (r ListPerformanceResultsRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

type ListPerformanceResultsResponse struct {
	PerformanceResults []models.PerformanceResult `json:"performance_results"`
	Pagination         models.PaginationResponse  `json:"pagination"`
}
