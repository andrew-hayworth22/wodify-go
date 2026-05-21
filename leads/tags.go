package leads

import (
	"context"
	"fmt"
	"net/http"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// AddTags adds a list of tags to a Lead
func (c *Client) AddTags(ctx context.Context, id int64, req UpdateTagsRequest) (*UpdateTagsResponse, error) {
	var out UpdateTagsResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}

// DeleteTags deletes a list of tags to a Lead
func (c *Client) DeleteTags(ctx context.Context, id int64, req UpdateTagsRequest) (*UpdateTagsResponse, error) {
	var out UpdateTagsResponse
	err := c.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// UpdateTagsRequest represents a request to add tags to a Lead or remove tags from a Lead
type UpdateTagsRequest struct {
	Tags []string `json:"tags"`
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// UpdateTagsResponse represents a response to an add lead tags request
type UpdateTagsResponse struct {
	Tags      []string `json:"active_tags"`
	IsSuccess bool     `json:"is_success"`
}
