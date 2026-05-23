package request_test

import (
	"net/url"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
)

type testField string

const (
	testFieldName testField = "name"
)

func TestPaginationRequest_ToQuery(t *testing.T) {
	tc := []struct {
		name     string
		req      request.PaginationRequest
		expected url.Values
	}{
		{
			name: "pagination",
			req: request.PaginationRequest{
				Page:     1,
				PageSize: 10,
			},
			expected: url.Values{
				"page":      {"1"},
				"page_size": {"10"},
			},
		},
		{
			name:     "no pagination",
			req:      request.PaginationRequest{},
			expected: url.Values{},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			testutil.AssertQueryEqual(t, c.expected, c.req.ToQuery())
		})
	}
}

func TestListRequest_ToQuery(t *testing.T) {
	tc := []struct {
		name     string
		req      request.ListRequest[testField]
		expected url.Values
	}{
		{
			name: "pagination & sorting",
			req: request.ListRequest[testField]{
				Page: request.PaginationRequest{Page: 2, PageSize: 5},
				Sort: sort.NewSort(testFieldName, false),
			},
			expected: url.Values{
				"page":      {"2"},
				"page_size": {"5"},
				"sort":      {string(testFieldName)},
			},
		},
		{
			name:     "empty request",
			req:      request.ListRequest[testField]{},
			expected: url.Values{},
		},
		{
			name: "pagination only",
			req: request.ListRequest[testField]{
				Page: request.PaginationRequest{Page: 1, PageSize: 10},
			},
			expected: url.Values{
				"page":      {"1"},
				"page_size": {"10"},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			testutil.AssertQueryEqual(t, c.expected, c.req.ToQuery())
		})
	}
}

func TestSearchRequest_ToQuery(t *testing.T) {
	tc := []struct {
		name     string
		req      request.SearchRequest[testField]
		expected url.Values
	}{
		{
			name: "pagination, sorting, and query",
			req: request.SearchRequest[testField]{
				Page:  request.PaginationRequest{Page: 2, PageSize: 5},
				Sort:  sort.NewSort(testFieldName, false),
				Query: query.New[testField]().Eq(testFieldName, "andy"),
			},
			expected: url.Values{
				"page":      {"2"},
				"page_size": {"5"},
				"sort":      {string(testFieldName)},
				"q":         {"name|eq|'andy'"},
			},
		},
		{
			name:     "empty request",
			req:      request.SearchRequest[testField]{},
			expected: url.Values{},
		},
		{
			name: "pagination only",
			req: request.SearchRequest[testField]{
				Page: request.PaginationRequest{Page: 1, PageSize: 10},
			},
			expected: url.Values{
				"page":      {"1"},
				"page_size": {"10"},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			testutil.AssertQueryEqual(t, c.expected, c.req.ToQuery())
		})
	}
}
