package types

// Pagination represents a Wodify pagination object.
type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasMore  bool `json:"has_more"`
}

// PageParameters represents a request for a specific page of data.
type PageParameters struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
