package leads

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
	"github.com/andrew-hayworth22/wodify-go/models"
)

// Client provides access to the Wodify Leads API.
type Client struct {
	hc *httpclient.Client
}

// New returns a new Leads client.
func New(hc *httpclient.Client) *Client {
	return &Client{hc: hc}
}

// Get fetches a single lead by ID.
func (s *Client) Get(ctx context.Context, id int64) (*models.Lead, error) {
	var out models.Lead
	err := s.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	return &out, err
}

// Create creates a new lead.
func (s *Client) Create(ctx context.Context, req CreateLeadRequest) (*models.Lead, error) {
	var out models.Lead
	err := s.hc.Do(ctx, http.MethodPost, "/leads", nil, req, &out)
	return &out, err
}

// List fetches a list of leads.
func (s *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var out ListResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads", req.ToQuery(), nil, &out)
	return &out, err
}

// Search fetches a list of leads matching the search criteria.
func (s *Client) Search(ctx context.Context, req SearchRequest) (*ListResponse, error) {
	var out ListResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads/search", req.ToQuery(), nil, &out)
	return &out, err
}

// Delete deletes a lead by ID.
func (s *Client) Delete(ctx context.Context, id int64) (*DeleteLeadResponse, error) {
	var out DeleteLeadResponse
	err := s.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	return &out, err
}

// Update updates a lead by ID.
func (s *Client) Update(ctx context.Context, id int64, req UpdateLeadRequest) (*models.Lead, error) {
	var out models.Lead
	err := s.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d", id), nil, req, &out)
	return &out, err
}

// Convert converts a lead to a Client.
func (s *Client) Convert(ctx context.Context, id int64, req ConvertLeadRequest) (*ConvertLeadResponse, error) {
	var out ConvertLeadResponse
	err := s.hc.Do(ctx, http.MethodPost, fmt.Sprintf("/leads/%d/convert", id), nil, req, &out)
	return &out, err
}

// ListStatuses fetches a list of lead statuses
func (s *Client) ListStatuses(ctx context.Context, req ListStatusesRequest) (*ListStatusesResponse, error) {
	var out ListStatusesResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads/statuses", req.ToQuery(), nil, &out)
	return &out, err
}

// ListSources fetches a list of lead sources
func (s *Client) ListSources(ctx context.Context, req ListSourcesRequest) (*ListSourcesResponse, error) {
	var out ListSourcesResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads/sources", req.ToQuery(), nil, &out)
	return &out, err
}

// AddTags adds a list of tags to a Lead
func (s *Client) AddTags(ctx context.Context, id int64, req UpdateTagsRequest) (*UpdateTagsResponse, error) {
	var out UpdateTagsResponse
	err := s.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}

// DeleteTags deletes a list of tags to a Lead
func (s *Client) DeleteTags(ctx context.Context, id int64, req UpdateTagsRequest) (*UpdateTagsResponse, error) {
	var out UpdateTagsResponse
	err := s.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d/tags", id), nil, req, &out)
	return &out, err
}
