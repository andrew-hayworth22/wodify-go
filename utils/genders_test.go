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

func TestClient_ListGenders(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/gender_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/genders",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListGendersRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewGenderSort(utils.GenderFieldID, false),
	}
	resp, err := svc.ListGenders(context.Background(), req)
	if err != nil {
		t.Fatalf("listing genders: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "gender_id" {
		t.Errorf("request sort: expected=%s; got=%s", "gender_id", query.Get("sort"))
	}

	// Check response
	if len(resp.Genders) != 2 {
		t.Errorf("response genders list length: expected=%d; got=%d", 2, len(resp.Genders))
	}
}
