package request

import (
	"net/url"
	"strconv"
)

// PaginationRequest represents a request for a specific page of data.
type PaginationRequest struct {
	// Specifies the page of records to be returned (1-indexed).
	Page int `json:"page"`
	// Specifies the number of records to be returned.
	PageSize int `json:"page_size"`
}

// ToQuery converts the PaginationRequest to URL query parameters.
func (p PaginationRequest) ToQuery() url.Values {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", strconv.Itoa(p.Page))
	}
	if p.PageSize > 0 {
		q.Set("page_size", strconv.Itoa(p.PageSize))
	}
	return q
}
