// Package search provides a simple query builder for Wodify search queries.
package search

import (
	"strings"
)

// Filterable represents a field that can be filtered by a search query.
type Filterable interface {
	~string
}

// Builder constructs a Wodify search query using the pipe-delimited syntax.
// Multiple clauses are AND'd together. Call Encode() to get the URL-safe
// value for the 'q' query parameter.
type Builder[T Filterable] struct {
	clauses []clause
}

// New returns a new search builder
func New[T Filterable]() *Builder[T] {
	return &Builder[T]{}
}

// Lt adds a less than clause to the search query
func (b *Builder[T]) Lt(field T, value any) *Builder[T] {
	return b.add(lessThan, field, value)
}

// Lte adds a less than or equal to clause to the search query
func (b *Builder[T]) Lte(field T, value any) *Builder[T] {
	return b.add(lessThanEqualTo, field, value)
}

// Gt adds a greater than clause to the search query
func (b *Builder[T]) Gt(field T, value any) *Builder[T] {
	return b.add(greaterThan, field, value)
}

// Gte adds a greater than or equal to clause to the search query
func (b *Builder[T]) Gte(field T, value any) *Builder[T] {
	return b.add(greaterThanEqualTo, field, value)
}

// Eq adds an equality clause to the search query
func (b *Builder[T]) Eq(field T, value any) *Builder[T] {
	return b.add(equal, field, value)
}

// Neq adds an inequality clause to the search query
func (b *Builder[T]) Neq(field T, value any) *Builder[T] {
	return b.add(notEqual, field, value)
}

// Between adds a between clause to the search query
func (b *Builder[T]) Between(field T, min, max any) *Builder[T] {
	return b.add(between, field, min, max)
}

// In adds an in clause to the search query
func (b *Builder[T]) In(field T, values ...any) *Builder[T] {
	return b.add(in, field, values...)
}

// NotIn adds a not in clause to the search query
func (b *Builder[T]) NotIn(field T, values ...any) *Builder[T] {
	return b.add(notIn, field, values...)
}

// IsNull adds a null clause to the search query
func (b *Builder[T]) IsNull(field T) *Builder[T] {
	return b.add(null, field)
}

// IsNotNull adds a not null clause to the search query
func (b *Builder[T]) IsNotNull(field T) *Builder[T] {
	return b.add(notNull, field)
}

// Contains adds a contains clause to the search query
func (b *Builder[T]) Contains(field T, value string) *Builder[T] {
	return b.add(contains, field, value)
}

// StartsWith adds a starts with clause to the search query
func (b *Builder[T]) StartsWith(field T, value string) *Builder[T] {
	return b.add(startsWith, field, value)
}

// EndsWith adds an ends with clause to the search query
func (b *Builder[T]) EndsWith(field T, value string) *Builder[T] {
	return b.add(endsWith, field, value)
}

// Encode returns the URL-safe value for the 'q' query parameter.
func (b *Builder[T]) Encode() string {
	return b.String()
}

// String returns the unencoded query string
func (b *Builder[T]) String() string {
	parts := make([]string, len(b.clauses))
	for i, c := range b.clauses {
		parts[i] = c.encode()
	}
	return strings.Join(parts, ";")
}

// add inspects the value audit.go and applies quoting for strings.
func (b *Builder[T]) add(operator operator, field T, values ...any) *Builder[T] {
	b.clauses = append(b.clauses, newClause(string(field), operator, values...))
	return b
}
