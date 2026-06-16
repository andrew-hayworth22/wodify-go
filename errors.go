package wodify

import (
	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
	"github.com/andrew-hayworth22/wodify-go/internal/query"
)

type APIError = httpclient.APIError

var (
	ErrNotFound     = httpclient.ErrNotFound
	ErrUnauthorized = httpclient.ErrUnauthorized
	ErrRateLimited  = httpclient.ErrRateLimited
	ErrBadRequest   = httpclient.ErrBadRequest
	ErrInvalidQuery = query.ErrReservedChar
)
