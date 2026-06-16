package query_test

import (
	"errors"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
)

type filterField string

const (
	filterName   filterField = "name"
	filterWeight filterField = "weight"
	filterActive filterField = "is_active"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name        string
		builder     *query.Builder[filterField]
		expectedRaw string
		expectedErr error
	}{
		{
			name:        "eq text",
			builder:     query.New[filterField]().Eq(filterName, "andy"),
			expectedRaw: "name|eq|'andy'",
		},
		{
			name:        "eq number",
			builder:     query.New[filterField]().Eq(filterWeight, 180),
			expectedRaw: "weight|eq|180",
		},
		{
			name:        "eq bool",
			builder:     query.New[filterField]().Eq(filterActive, true),
			expectedRaw: "is_active|eq|true",
		},
		{
			name: "numeric comparison",
			builder: query.New[filterField]().
				Lt(filterWeight, 180).
				Lte(filterWeight, 180).
				Gt(filterWeight, 180).
				Gte(filterWeight, 180).
				Neq(filterWeight, 180),
			expectedRaw: "weight|lt|180;weight|lte|180;weight|gt|180;weight|gte|180;weight|neq|180",
		},
		{
			name:        "between",
			builder:     query.New[filterField]().Between(filterWeight, 180, 200),
			expectedRaw: "weight|between|180|200",
		},
		{
			name:        "in text",
			builder:     query.New[filterField]().In(filterName, "andy", "bob", "charlie"),
			expectedRaw: "name|in|{'andy','bob','charlie'}",
		},
		{
			name:        "in number",
			builder:     query.New[filterField]().In(filterWeight, 180, 190, 200),
			expectedRaw: "weight|in|{180,190,200}",
		},
		{
			name:        "not in",
			builder:     query.New[filterField]().NotIn(filterWeight, 180, 190, 200),
			expectedRaw: "weight|not_in|{180,190,200}",
		},
		{
			name:        "null/not null",
			builder:     query.New[filterField]().IsNull(filterActive).IsNotNull(filterWeight),
			expectedRaw: "is_active|is_null;weight|not_null",
		},
		{
			name:        "contains",
			builder:     query.New[filterField]().Contains(filterName, "andy"),
			expectedRaw: "name|contains|'andy'",
		},
		{
			name:        "starts/ends with",
			builder:     query.New[filterField]().StartsWith(filterName, "andy").EndsWith(filterName, "worth"),
			expectedRaw: "name|starts_with|'andy';name|ends_with|'worth'",
		},
		{
			name:        "reserved character error - base characters enforced",
			builder:     query.New[filterField]().StartsWith(filterName, "an'dy").EndsWith(filterName, "worth"),
			expectedErr: query.ErrReservedChar,
		},
		{
			name:        "reserved character error - reserved list characters not enforced",
			builder:     query.New[filterField]().StartsWith(filterName, "an{dy").EndsWith(filterName, "worth"),
			expectedRaw: "name|starts_with|'an{dy';name|ends_with|'worth'",
		},
		{
			name:        "reserved character error - list characters enforced",
			builder:     query.New[filterField]().In(filterName, "andy", "band,y").EndsWith(filterName, "worth"),
			expectedErr: query.ErrReservedChar,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedErr == nil && tt.builder.Err() != nil {
				t.Fatalf("expected no error, got '%v'", tt.builder.Err())
			}
			if tt.expectedErr != nil {
				if !errors.Is(tt.builder.Err(), tt.expectedErr) {
					t.Errorf("expected error '%v', got '%v'", tt.expectedErr, tt.builder.Err())
				}
				return
			}
			if got := tt.builder.String(); got != tt.expectedRaw {
				t.Errorf("expected='%s' got='%s'", tt.expectedRaw, got)
			}
		})
	}
}
