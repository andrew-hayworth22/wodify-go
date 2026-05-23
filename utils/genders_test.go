package utils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/testutil"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func TestClient_ListGenders(t *testing.T) {
	// Load response fixture
	body := testutil.MustReadJSONFixture(t, "testdata/gender_list.json")

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/genders",
		StatusCode: http.StatusOK,
		Body:       body,
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	p := wodify.NewPaginationRequest(1, 10)
	s := wodify.SortAscending(utils.GenderFieldID)
	req := utils.NewGenderListRequest(p, s)
	resp, err := svc.ListGenders(context.Background(), req)
	if err != nil {
		t.Fatalf("listing genders: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	testutil.AssertPaginationParams(t, query, p)
	testutil.AssertSortParam(t, query, s)

	// Check response
	if len(resp.Genders) != 2 {
		t.Errorf("response genders list length: expected=%d; got=%d", 2, len(resp.Genders))
	}
}
