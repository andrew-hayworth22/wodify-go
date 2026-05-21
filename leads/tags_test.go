package leads_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func TestClient_AddTags(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/tag_update.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodPut,
		Path:       "/leads/123/tags",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.UpdateTagsRequest{
		Tags: []string{"tag 1", "tag 2"},
	}
	resp, err := svc.AddTags(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("adding lead tags: %v", err)
	}

	// Check request
	var sentRequest leads.UpdateTagsRequest
	if err := json.Unmarshal(hdl.RequestBody, &sentRequest); err != nil {
		t.Fatalf("decoding request: %v", err)
	}
	if len(sentRequest.Tags) != len(req.Tags) {
		t.Errorf("request tag length: expected=%d; got=%d", len(req.Tags), len(sentRequest.Tags))
	}

	// Check response
	if len(resp.Tags) != 3 {
		t.Errorf("expected 3 tags, got=%d", len(resp.Tags))
	}
}

func TestClient_DeleteTags(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/tag_update.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodDelete,
		Path:       "/leads/123/tags",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := leads.New(svr)

	// Make request
	req := leads.UpdateTagsRequest{
		Tags: []string{"tag 1", "tag 2"},
	}
	resp, err := svc.DeleteTags(context.Background(), 123, req)
	if err != nil {
		t.Fatalf("deleting lead tags: %v", err)
	}

	// Check request
	var sentRequest leads.UpdateTagsRequest
	if err := json.Unmarshal(hdl.RequestBody, &sentRequest); err != nil {
		t.Fatalf("decoding request: %v", err)
	}
	if len(sentRequest.Tags) != len(req.Tags) {
		t.Errorf("request tag length: expected=%d; got=%d", len(req.Tags), len(sentRequest.Tags))
	}

	// Check response
	if len(resp.Tags) != 3 {
		t.Errorf("expected 3 tags, got=%d", len(resp.Tags))
	}
}
