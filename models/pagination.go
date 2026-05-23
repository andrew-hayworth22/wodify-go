package models

// PaginationResponse represents a Wodify pagination object.
type PaginationResponse struct {
	// Specifies the page of records that was returned (1-indexed).
	Page int `json:"page"`
	// Specifies the number of records that was returned.
	PageSize int `json:"page_size"`
	// Indicates whether there are more records to be returned.
	HasMore bool `json:"has_more"`
}
