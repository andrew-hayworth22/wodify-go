package wodify

import (
	"net/http"
	"time"
)

// config is the configuration for the client.
type config struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
	maxRetries int
}

// defaultConfig returns the default configuration for the client.
func defaultConfig() config {
	return config{
		baseURL:    "https://api.wodify.com/v1",
		httpClient: &http.Client{},
		timeout:    10 * time.Second,
		maxRetries: 3,
	}
}

// Option is a function that configures the client.
type Option func(*config)

// WithAPIKey sets the API key for the client.
func WithAPIKey(key string) Option {
	return func(c *config) {
		c.apiKey = key
	}
}

// WithBaseURL sets the base URL for the client.
func WithBaseURL(url string) Option {
	return func(c *config) {
		c.baseURL = url
	}
}

// WithHTTPClient sets the HTTP client for the client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *config) {
		c.httpClient = client
	}
}

// WithTimeout sets the timeout for the client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *config) {
		c.timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries for the client.
func WithMaxRetries(maxRetries int) Option {
	return func(c *config) {
		c.maxRetries = maxRetries
	}
}
