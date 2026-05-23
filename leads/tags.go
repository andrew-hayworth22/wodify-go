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
func (c *Client) AddTags(ctx context.Context, id int64, req TagsUpdateRequest) (*TagsUpdateResponse, error) {
	var out TagsUpdateResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}

// DeleteTags deletes a list of tags to a Lead
func (c *Client) DeleteTags(ctx context.Context, id int64, req TagsUpdateRequest) (*TagsUpdateResponse, error) {
	var out TagsUpdateResponse
	err := c.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// TagsUpdateRequest represents a request to add tags to a Lead or remove tags from a Lead
type TagsUpdateRequest struct {
	Tags []string `json:"tags"`
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// TagsUpdateResponse represents a response to an add lead tags request
type TagsUpdateResponse struct {
	Tags      []string `json:"active_tags"`
	IsSuccess bool     `json:"is_success"`
}
