package leads

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListPerformanceResults fetches a list of a Lead's performance results
func (c *Client) ListPerformanceResults(ctx context.Context, id int64, req PerformanceResultListRequest) (*PerformanceResultListResponse, error) {
	var out PerformanceResultListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/performance-results", id), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// ListPerformanceResultsByComponent fetches a list of a Lead's performance results for a specific component'
func (c *Client) ListPerformanceResultsByComponent(ctx context.Context, leadID int64, componentID int64, req PerformanceResultListRequest) (*PerformanceResultListResponse, error) {
	var out PerformanceResultListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/performance-results/components/%d", leadID, componentID), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// PerformanceResultListRequest represents a request to list a Lead's performance results
type PerformanceResultListRequest = request.ListRequest[string]

// NewPerformanceResultListRequest creates a new PerformanceResultListRequest with the given pagination.
func NewPerformanceResultListRequest(pagination request.PaginationRequest) PerformanceResultListRequest {
	return PerformanceResultListRequest{
		Page: pagination,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

type PerformanceResultListResponse struct {
	PerformanceResults []models.PerformanceResult `json:"performance_results"`
	Pagination         models.PaginationResponse  `json:"pagination"`
}
