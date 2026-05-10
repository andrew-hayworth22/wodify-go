package sort_test

import (
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/sort"
)

type testSortable string

const (
	testSortableString testSortable = "string"
)

func TestSort_String(t *testing.T) {
	ascending := sort.NewSort(testSortableString, false)
	ascendingStr := ascending.String()
	if ascendingStr != "string" {
		t.Errorf("expected=%s; got=%s", testSortableString, ascendingStr)
	}

	descending := sort.NewSort(testSortableString, true)
	descendingStr := descending.String()
	if descendingStr != "desc_string" {
		t.Errorf("expected=%s; got=%s", "desc_string", descendingStr)
	}
}
