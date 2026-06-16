// Package sort provides a generic sort builder for Wodify's sort URL parameter
package sort

import (
	"fmt"
)

// Sortable represents a column that can be sorted on.
type Sortable interface {
	~string
}

// Sort represents a sort passed to the API.
// Used to order results from a list request.
type Sort[T Sortable] struct {
	Field        T
	IsDescending bool
}

// NewSort creates a new sort with the given field and direction.
func NewSort[T Sortable](field T, isDescending bool) Sort[T] {
	return Sort[T]{Field: field, IsDescending: isDescending}
}

// String returns a string representation of the sort.
// Should be passed to the API as the 'sort' query parameter.
func (s Sort[T]) String() string {
	var prefix string
	if s.IsDescending {
		prefix = "desc_"
	}

	return fmt.Sprintf("%s%s", prefix, s.Field)
}
