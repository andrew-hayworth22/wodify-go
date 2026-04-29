package leads

import (
	"context"
	"fmt"

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
func (s *Service) Get(ctx context.Context, id int64) (Lead, error) {
	var out Lead
	err := s.hc.Do(ctx, "GET", fmt.Sprintf("/leads/%d", id), nil, &out)
	return out, err
}
