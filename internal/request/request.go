// Package request provides common utilities related to building requests to the Wodify API
package request

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
)

// ListRequest represents a generic request to list a resource
// The Field type represents the fields that will be sortable/queriable
type ListRequest[Field ~string] struct {
	Page PaginationRequest
	Sort sort.Sort[Field]
}

// ToQuery converts the request to URL query string parameters
func (r ListRequest[Field]) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchRequest represents a generic request to query a resource
// The Field type represents the fields that will be sortable/queriable
type SearchRequest[Field ~string] struct {
	Page  PaginationRequest
	Sort  sort.Sort[Field]
	Query *query.Builder[Field]
}

// ToQuery converts the request to URL query string parameters
func (r SearchRequest[Field]) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	if r.Query != nil {
		q.Set("q", r.Query.String())
	}
	return q
}
