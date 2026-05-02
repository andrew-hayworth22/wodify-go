package leads

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
)

// Service provides access to the Wodify Leads API.
type Service struct {
	hc *httpclient.Client
}

// New returns a new Leads service.
func New(hc *httpclient.Client) *Service {
	return &Service{hc: hc}
}

// Get returns a single lead by ID.
func (s *Service) Get(ctx context.Context, id int64) (*Lead, error) {
	var out Lead
	err := s.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	return &out, err
}

// Create creates a new lead.
func (s *Service) Create(ctx context.Context, req CreateLeadRequest) (*Lead, error) {
	var out Lead
	err := s.hc.Do(ctx, http.MethodPost, "/leads", nil, req, &out)
	return &out, err
}

// List returns a list of leads.
func (s *Service) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var out ListResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads", req.ToQuery(), req, &out)
	return &out, err
}

// Search returns a list of leads matching the search criteria.
func (s *Service) Search(ctx context.Context, req SearchRequest) (*ListResponse, error) {
	var out ListResponse
	err := s.hc.Do(ctx, http.MethodGet, "/leads/search", req.ToQuery(), req, &out)
	return &out, err
}

// Delete deletes a lead by ID.
func (s *Service) Delete(ctx context.Context, id int64) (*DeleteLeadResponse, error) {
	var out DeleteLeadResponse
	err := s.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	return &out, err
}
