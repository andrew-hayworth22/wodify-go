package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Sentinel error values
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrRateLimited  = errors.New("rate limited")
	ErrBadRequest   = errors.New("bad request")
)

// APIError is an error returned by the API.
type APIError struct {
	HTTPCode         int    `json:"HTTPCode"`
	DeveloperMessage string `json:"DeveloperMessage"`
	UserMessage      string `json:"UserMessage"`
	MoreInfo         string `json:"MoreInfo"`
	err              error
}

// Error returns the error message.
func (e *APIError) Error() string {
	return fmt.Sprintf("wodify api error %d: %s", e.HTTPCode, e.MoreInfo)
}

// Unwrap returns the underlying error value
func (e *APIError) Unwrap() error {
	return e.err
}

// UnmarshalJSON implements the JSON unmarshaler
func (e *APIError) UnmarshalJSON(data []byte) error {
	// Alias the error type to get default behavior with no infinite looping
	type alias APIError
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*e = APIError(a)
	e.err = sentinelFor(e.HTTPCode)
	return nil
}

// newAPIError constructs a new APIError value
func newAPIError(code int, status string) *APIError {
	return &APIError{
		HTTPCode: code,
		MoreInfo: status,
		err:      sentinelFor(code),
	}
}

// sentinelFor returns the sentinel error for the provided HTTP status code
func sentinelFor(code int) error {
	switch code {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusUnauthorized, http.StatusForbidden:
		return ErrUnauthorized
	case http.StatusTooManyRequests:
		return ErrRateLimited
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return ErrBadRequest
	}
	return nil
}
