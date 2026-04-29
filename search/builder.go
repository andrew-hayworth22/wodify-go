// Package search provides a simple query builder for Wodify search queries.
package search

import (
	"net/url"
	"strings"
)

// Builder constructs a Wodify search query using the pipe-delimited syntax.
// Multiple clauses are AND'd together. Call Encode() to get the URL-safe
// value for the 'q' query parameter.
type Builder struct {
	clauses []clause
}

// New returns a new search builder
func New() *Builder {
	return &Builder{}
}

// Lt adds a less than clause to the search query
func (b *Builder) Lt(field string, value any) *Builder {
	return b.add(lessThan, field, value)
}

// Lte adds a less than or equal to clause to the search query
func (b *Builder) Lte(field string, value any) *Builder {
	return b.add(lessThanEqualTo, field, value)
}

// Gt adds a greater than clause to the search query
func (b *Builder) Gt(field string, value any) *Builder {
	return b.add(greaterThan, field, value)
}

// Gte adds a greater than or equal to clause to the search query
func (b *Builder) Gte(field string, value any) *Builder {
	return b.add(greaterThanEqualTo, field, value)
}

// Eq adds an equality clause to the search query
func (b *Builder) Eq(field string, value any) *Builder {
	return b.add(equal, field, value)
}

// Neq adds an inequality clause to the search query
func (b *Builder) Neq(field string, value any) *Builder {
	return b.add(notEqual, field, value)
}

// Between adds a between clause to the search query
func (b *Builder) Between(field string, min, max any) *Builder {
	return b.add(between, field, min, max)
}

// In adds an in clause to the search query
func (b *Builder) In(field string, values ...any) *Builder {
	return b.add(in, field, values...)
}

// NotIn adds a not in clause to the search query
func (b *Builder) NotIn(field string, values ...any) *Builder {
	return b.add(notIn, field, values...)
}

// IsNull adds a null clause to the search query
func (b *Builder) IsNull(field string) *Builder {
	return b.add(null, field)
}

// IsNotNull adds a not null clause to the search query
func (b *Builder) IsNotNull(field string) *Builder {
	return b.add(notNull, field)
}

// Contains adds a contains clause to the search query
func (b *Builder) Contains(field, value string) *Builder {
	return b.add(contains, field, value)
}

// StartsWith adds a starts with clause to the search query
func (b *Builder) StartsWith(field, value string) *Builder {
	return b.add(startsWith, field, value)
}

// EndsWith adds an ends with clause to the search query
func (b *Builder) EndsWith(field, value string) *Builder {
	return b.add(endsWith, field, value)
}

// Encode returns the URL-safe value for the 'q' query parameter.
func (b *Builder) Encode() string {
	return url.QueryEscape(b.String())
}

// String returns the unencoded query string
func (b *Builder) String() string {
	parts := make([]string, len(b.clauses))
	for i, c := range b.clauses {
		parts[i] = c.encode()
	}
	return strings.Join(parts, ";")
}

// add inspects the value types.go and applies quoting for strings.
func (b *Builder) add(operator operator, field string, values ...any) *Builder {
	b.clauses = append(b.clauses, newClause(field, operator, values...))
	return b
}
