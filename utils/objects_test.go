package utils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListObjectTypes(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/object_type_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-types",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.ObjectTypeFieldID)
	req := utils.NewObjectTypeListRequest(p, s)
	resp, err := svc.ListObjectTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.ObjectTypes) != 2 {
		t.Errorf("response object types list length: expected=%d; got=%d", 2, len(resp.ObjectTypes))
	}
}

func TestClient_SearchObjectTypes(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/object_type_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-types/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.ObjectTypeFieldID)
	q := utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldID, 123)
	req := utils.NewObjectTypeSearchRequest(p, s, q)
	resp, err := svc.SearchObjectTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("searching object types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.ObjectTypes) != 2 {
		t.Errorf("response object type list length: expected=%d; got=%d", 2, len(resp.ObjectTypes))
	}
}

func TestClient_ListObjectActionTypes(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/object_action_type_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-type-action-types",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.ObjectActionTypeFieldName)
	req := utils.NewObjectActionTypeListRequest(p, s)
	resp, err := svc.ListObjectActionTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object action types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.ObjectActionTypes) != 2 {
		t.Errorf("response object action types list length: expected=%d; got=%d", 2, len(resp.ObjectActionTypes))
	}
}

func TestClient_SearchObjectActionTypes(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/object_action_type_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/object-type-action-types/search",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.ObjectActionTypeFieldName)
	q := utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldID, 123)
	req := utils.NewObjectActionTypeSearchRequest(p, s, q)
	resp, err := svc.SearchObjectActionTypes(context.Background(), req)
	if err != nil {
		t.Fatalf("listing object action types: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)
	testutil.AssertQueryParam(t, query, q)

	// Check response
	if len(resp.ObjectActionTypes) != 2 {
		t.Errorf("response object action types list length: expected=%d; got=%d", 2, len(resp.ObjectActionTypes))
	}
}
