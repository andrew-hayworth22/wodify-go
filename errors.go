package wodify

import "github.com/andrew-hayworth22/wodify-go/internal/httpclient"

type APIError = httpclient.APIError

var (
	ErrNotFound     = httpclient.ErrNotFound
	ErrUnauthorized = httpclient.ErrUnauthorized
	ErrRateLimited  = httpclient.ErrRateLimited
	ErrBadRequest   = httpclient.ErrBadRequest
)
