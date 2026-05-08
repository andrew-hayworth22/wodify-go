package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/andrew-hayworth22/wodify-go/internal/version"
)

// Client is the internal HTTP client for the library.
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
	maxRetries int
}

// New creates a new HTTP client.
func New(httpClient *http.Client, baseURL, apiKey string, maxRetries int) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
		maxRetries: maxRetries,
	}
}

// Do makes an HTTP call to the API.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body, out any) error {
	// Add query parameters to the path
	if len(query) != 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	var attempt int
	for {
		// Build the request
		req, err := c.buildRequest(ctx, method, path, body)
		if err != nil {
			return err
		}

		// Make the HTTP request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}

		// If too many requests, retry with exponential backoff
		if resp.StatusCode == http.StatusTooManyRequests && attempt < c.maxRetries {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff(attempt)):
			}
			attempt++
			continue
		}

		// Handle errors
		if resp.StatusCode >= 400 {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
			return newAPIError(resp.StatusCode, resp.Status)
		}

		return c.decode(resp, out)
	}
}

// buildRequest builds an HTTP request for the API.
func (c *Client) buildRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	// Build the request body
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, buf)
	if err != nil {
		return nil, err
	}

	// Inject headers
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", version.UserAgent)
	return req, nil
}

// backoff calculates the backoff duration for a retry attempt.
func backoff(attempt int) time.Duration {
	// Formula: 2^attempt * 1 second capped at 30 seconds
	return min(time.Duration(1<<attempt)*time.Second, 30*time.Second)
}

// decode decodes the response body into the provided struct.
func (c *Client) decode(resp *http.Response, out any) error {
	defer resp.Body.Close()

	// If ignoring the response or no content, return
	if out == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Read the response body
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for API error
	// Wodify does not return HTTP status codes for some errors, so we need to check the body.
	var apiErr APIError
	if err := json.Unmarshal(b, &apiErr); err == nil && apiErr.HTTPCode >= 400 {
		return &apiErr
	}

	return json.Unmarshal(b, out)
}
