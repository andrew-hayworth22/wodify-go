package utils_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/models"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListObjectTypes(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/object_type_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-types",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListObjectTypesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewObjectTypeSort(utils.ObjectTypeFieldID, false),
	}
	resp, err := svc.ListObjectTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "object_type_id" {
		t.Errorf("request sort: expected=%s; got=%s", "object_type_id", query.Get("sort"))
	}

	// Check response
	if len(resp.ObjectTypes) != 2 {
		t.Errorf("response object types list length: expected=%d; got=%d", 2, len(resp.ObjectTypes))
	}
}

func TestClient_SearchObjectTypes(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/object_type_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-types/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.SearchObjectTypesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  utils.NewObjectTypeSort(utils.ObjectTypeFieldName, false),
		Query: utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldID, 123),
	}
	resp, err := svc.SearchObjectTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("searching object types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "object_type" {
		t.Errorf("request sort: expected=%s; got=%s", "object_type", query.Get("sort"))
	}
	if query.Get("q") != "object_type_id|eq|123" {
		t.Errorf("request query: expected=%s; got=%s", "object_type_id|eq|123", query.Get("q"))
	}

	// Check response
	if len(resp.ObjectTypes) != 2 {
		t.Errorf("response object type list length: expected=%d; got=%d", 2, len(resp.ObjectTypes))
	}
}

func TestClient_ListObjectActionTypes(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/object_action_type_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-type-action-types",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListObjectActionTypesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewObjectActionTypeSort(utils.ObjectActionTypeFieldName, false),
	}
	resp, err := svc.ListObjectActionTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object action types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "action_type" {
		t.Errorf("request sort: expected=%s; got=%s", "action_type", query.Get("sort"))
	}

	// Check response
	if len(resp.ObjectActionTypes) != 2 {
		t.Errorf("response object action types list length: expected=%d; got=%d", 2, len(resp.ObjectActionTypes))
	}
}

func TestClient_SearchObjectActionTypes(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/object_action_type_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-type-action-types/search",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.SearchObjectActionTypesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  utils.NewObjectActionTypeSort(utils.ObjectActionTypeFieldName, false),
		Query: utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldID, 123),
	}
	resp, err := svc.SearchObjectActionTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object action types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "action_type" {
		t.Errorf("request sort: expected=%s; got=%s", "action_type", query.Get("sort"))
	}
	if query.Get("q") != "action_type_id|eq|123" {
		t.Errorf("request query: expected=%s; got=%s", "action_type_id|eq|123", query.Get("q"))
	}

	// Check response
	if len(resp.ObjectActionTypes) != 2 {
		t.Errorf("response object action types list length: expected=%d; got=%d", 2, len(resp.ObjectActionTypes))
	}
}
