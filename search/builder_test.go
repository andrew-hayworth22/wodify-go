package search_test

import (
	"testing"

	"github.com/andrew-hayworth22/wodify-go/search"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name        string
		builder     *search.Builder
		expectedRaw string
	}{
		{
			name:        "eq text",
			builder:     search.New().Eq("name", "andy"),
			expectedRaw: "name|eq|'andy'",
		},
		{
			name:        "eq number",
			builder:     search.New().Eq("weight", 180),
			expectedRaw: "weight|eq|180",
		},
		{
			name:        "eq bool",
			builder:     search.New().Eq("is_active", true),
			expectedRaw: "is_active|eq|true",
		},
		{
			name:        "multiple clauses",
			builder:     search.New().Eq("name", "andy").Eq("weight", 180),
			expectedRaw: "name|eq|'andy';weight|eq|180",
		},
		{
			name:        "between",
			builder:     search.New().Between("weight", 180, 200),
			expectedRaw: "weight|between|180|200",
		},
		{
			name:        "in text",
			builder:     search.New().In("name", "andy", "bob", "charlie"),
			expectedRaw: "name|in|{'andy','bob','charlie'}",
		},
		{
			name:        "in number",
			builder:     search.New().In("weight", 180, 190, 200),
			expectedRaw: "weight|in|{180,190,200}",
		},
		{
			name:        "not in",
			builder:     search.New().NotIn("weight", 180, 190, 200),
			expectedRaw: "weight|not_in|{180,190,200}",
		},
		{
			name:        "null",
			builder:     search.New().IsNull("is_active"),
			expectedRaw: "is_active|is_null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.builder.String(); got != tt.expectedRaw {
				t.Errorf("expected='%s' got='%s'", tt.expectedRaw, got)
			}
		})
	}
}
