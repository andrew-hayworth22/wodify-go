package testutil

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
)

// AssertPaginationParams asserts that the given query parameters contain the expected pagination values.
func AssertPaginationParams(t *testing.T, query url.Values, expected request.PaginationRequest) {
	actualPage := query.Get("page")
	if actualPage != strconv.Itoa(expected.Page) {
		t.Errorf("page: expected=%d; got=%s", expected.Page, actualPage)
	}

	actualPageSize := query.Get("page_size")
	if actualPageSize != strconv.Itoa(expected.PageSize) {
		t.Errorf("page_size: expected=%d; got=%s", expected.PageSize, actualPageSize)
	}
}

// AssertSortParams asserts that the given query parameters contain the expected sort values.
func AssertSortParam[T ~string](t *testing.T, query url.Values, expected sort.Sort[T]) {
	actualSort := query.Get("sort")
	expectedSort := expected.String()
	if actualSort != expectedSort {
		t.Errorf("sort: expected=%s; got=%s", expectedSort, actualSort)
	}
}

// AssertQueryParam asserts that the given query parameters contain the expected values.
func AssertQueryParam[T ~string](t *testing.T, query url.Values, expected *query.Builder[T]) {
	actualQuery := query.Get("q")
	expectedQuery := expected.String()
	if actualQuery != expectedQuery {
		t.Errorf("query: expected=%s; got=%s", expectedQuery, actualQuery)
	}
}

// AssertQueryEqual asserts that two url.Values are equal.
func AssertQueryEqual(t *testing.T, expected, actual url.Values) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Errorf("query param count: expected=%d got=%d (expected=%v actual=%v)", len(expected), len(actual), expected, actual)
		return
	}
	for k, v := range expected {
		if actual.Get(k) != v[0] {
			t.Errorf("param %q: expected=%s got=%s", k, v[0], actual.Get(k))
		}
	}
}
