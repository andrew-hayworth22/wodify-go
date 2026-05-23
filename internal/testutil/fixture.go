package testutil

import (
	"encoding/json"
	"os"
	"testing"
)

// MustReadJSONFixture reads a fixture and returns the raw JSON bytes.
func MustReadJSONFixture(t *testing.T, path string) json.RawMessage {
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	return json.RawMessage(body)
}
