package testutil

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

// AssertURLValuesEqual asserts that two url.Values are equal.
func AssertURLValuesEqual(t *testing.T, expected, actual url.Values) {
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

// AssertJSONEqual asserts that two JSON payloads are equal ignoring key order
func AssertJSONEqual(t *testing.T, expected, actual []byte) {
	t.Helper()
	var expectedMap, actualMap map[string]any
	if err := json.Unmarshal(expected, &expectedMap); err != nil {
		t.Fatalf("expected json unmarshal: %v", err)
	}
	if err := json.Unmarshal(actual, &actualMap); err != nil {
		t.Fatalf("actual json unmarshal: %v", err)
	}

	if !reflect.DeepEqual(expectedMap, actualMap) {
		expectedPretty, _ := json.MarshalIndent(expectedMap, "", "  ")
		actualPretty, _ := json.MarshalIndent(actualMap, "", "  ")
		t.Errorf("JSON mismatch:\nexpected:\n%s\ngot:\n%s", string(expectedPretty), string(actualPretty))
	}
}
