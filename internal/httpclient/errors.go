package httpclient

import "fmt"

// APIError is an error returned by the API.
type APIError struct {
	HTTPCode         int    `json:"HTTPCode"`
	DeveloperMessage string `json:"DeveloperMessage"`
	UserMessage      string `json:"UserMessage"`
	MoreInfo         string `json:"MoreInfo"`
}

// Error returns the error message.
func (e *APIError) Error() string {
	return fmt.Sprintf("wodify error %d: %s %s", e.HTTPCode, e.UserMessage, e.MoreInfo)
}
