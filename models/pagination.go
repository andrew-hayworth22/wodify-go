package models

import (
	"net/url"
	"strconv"
)

// PaginationResponse represents a Wodify pagination object.
type PaginationResponse struct {
	// Specifies the page of records that was returned (1-indexed).
	Page int `json:"page"`
	// Specifies the number of records that was returned.
	PageSize int `json:"page_size"`
	// Indicates whether there are more records to be returned.
	HasMore bool `json:"has_more"`
}

// PaginationRequest represents a request for a specific page of data.
type PaginationRequest struct {
	// Specifies the page of records to be returned (1-indexed).
	Page int `json:"page"`
	// Specifies the number of records to be returned.
	PageSize int `json:"page_size"`
}

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
