// Package wodify provides a client for the Wodify API.
package wodify

import (
	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
)

type Client struct {
	httpClient *httpclient.Client
}

func New(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	hc := httpclient.New(cfg.httpClient, cfg.baseURL, cfg.apiKey, cfg.maxRetries)
	return &Client{
		httpClient: hc,
	}, nil
}
