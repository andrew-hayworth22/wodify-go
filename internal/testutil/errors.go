package testutil

import (
	"fmt"
	"io"
	"net/http"
)

// MarshalError is a type that will return an error when marshalled
type MarshalError struct{}

// MarshalJSON returns an error
func (MarshalError) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("intentional JSON marshal error")
}

// ReadError is a reader that will return an error when read from
type ReadError struct{}

// Read returns an error
func (ReadError) Read([]byte) (int, error) {
	return 0, fmt.Errorf("intentional read error")
}

// TransportError is an HTTP transport that returns an error on every request
type TransportError struct{}

// RoundTrip returns an error
func (TransportError) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("intentional HTTP transport error")
}

// TransportReadError is an HTTP transport that returns an error when reading the body
type TransportReadError struct{}

// RoundTrip returns a response that will error when reading the body
func (TransportReadError) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(ReadError{}),
	}, nil
}
