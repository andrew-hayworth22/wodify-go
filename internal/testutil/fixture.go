package testutil

import (
	"bytes"
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

	var minified bytes.Buffer
	err = json.Compact(&minified, body)
	if err != nil {
		t.Fatalf("compacting fixture: %v", err)
	}

	return json.RawMessage(minified.Bytes())
}
