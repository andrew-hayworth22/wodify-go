package wodify

import (
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
)

// NewPaginationRequest creates a new pagination request
func NewPaginationRequest(page, pageSize int) request.PaginationRequest {
	return request.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

// SortDescending creates a new descending sort request
func SortDescending[Field ~string](field Field) sort.Sort[Field] {
	return sort.NewSort(field, true)
}

// SortAscending creates a new ascending sort request
func SortAscending[Field ~string](field Field) sort.Sort[Field] {
	return sort.NewSort(field, false)
}
