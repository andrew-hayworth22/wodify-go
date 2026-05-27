// Package wodify provides a Go SDK for the Wodify API.
//
// # Authentication
//
// A Wodify API key is required to use this SDK. It can be provided explicitly
// via [WithAPIKey] or through the environment.
//
// # Environment Variables
//
// The following environment variables are supported:
//
//	WODIFY_API_KEY      - Required. Your Wodify API key.
//	WODIFY_BASE_URL     - Optional. Override the default API base URL.
//	WODIFY_MAX_RETRIES  - Optional. Maximum number of retries for rate limited requests (default: 3).
//
// # Usage
//
//	client, err := wodify.New()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	lead, err := client.Leads.Get(ctx, "lead-id")
package wodify

import (
	"errors"
	"os"
	"strconv"

	"github.com/andrew-hayworth22/wodify-go/clients"
	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

// Client provides access to the Wodify API.
type Client struct {
	httpClient *httpclient.Client

	Utils   *utils.Client
	Leads   *leads.Client
	Clients *clients.Client
}

// New creates a new Wodify client.
func New(opts ...Option) (*Client, error) {
	// Start with the default configuration
	cfg := defaultConfig()

	// Apply environment variables
	applyEnv(&cfg)

	// Apply options
	for _, opt := range opts {
		opt(&cfg)
	}

	// Validate configuration
	if cfg.apiKey == "" {
		return nil, errors.New("wodify: no API key provided; set WODIFY_API_KEY or use WithAPIKey()")
	}

	internalClient := httpclient.New(cfg.httpClient, cfg.baseURL, cfg.apiKey, cfg.maxRetries)
	return &Client{
		httpClient: internalClient,
		Utils:      utils.New(internalClient),
		Leads:      leads.New(internalClient),
		Clients:    clients.New(internalClient),
	}, nil
}

// applyEnv applies environment variables to the configuration.
func applyEnv(cfg *config) {
	if v := os.Getenv("WODIFY_API_KEY"); v != "" {
		cfg.apiKey = v
	}
	if v := os.Getenv("WODIFY_BASE_URL"); v != "" {
		cfg.baseURL = v
	}
	if v := os.Getenv("WODIFY_MAX_RETRIES"); v != "" {
		// If the value is not a number, ignore it
		if n, err := strconv.Atoi(v); err == nil {
			cfg.maxRetries = n
		}
	}
}
