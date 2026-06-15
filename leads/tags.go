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
func (c *Client) AddTags(ctx context.Context, id int64, req TagsUpdateRequest) (*AddTagsResponse, error) {
	var out AddTagsResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// DeleteTags deletes a list of tags to a Lead
func (c *Client) DeleteTags(ctx context.Context, id int64, req TagsUpdateRequest) (*DeleteTagsResponse, error) {
	var out DeleteTagsResponse
	err := c.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// TagsUpdateRequest represents a request to add or remove tags from a lead
type TagsUpdateRequest struct {
	Tags []string `json:"tags"`
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// DeleteTagsResponse represents a response to a delete lead tags request
type DeleteTagsResponse struct {
	Tags      []string `json:"active_tags"`
	IsSuccess bool     `json:"is_success"`
}

// AddTagsResponse represents a response to an add lead tags request
type AddTagsResponse struct {
	Tags []string `json:"active_tags"`
}
