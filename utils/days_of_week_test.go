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

func TestClient_ListDaysOfWeek(t *testing.T) {
	// Load response fixture
	body, err := os.ReadFile("testdata/day_of_week_list.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	// Create test server and client
	hdl := &testutil.Handler{
		Method:     http.MethodGet,
		Path:       "/utilities/days-of-week",
		StatusCode: http.StatusOK,
		Body:       json.RawMessage(body),
	}
	svr := testutil.NewServer(t, hdl)
	svc := utils.New(svr)

	// Make request
	req := utils.ListDaysOfWeekRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: utils.NewDayOfWeekSort(utils.DayOfWeekFieldID, false),
	}
	resp, err := svc.ListDaysOfWeek(context.Background(), req)
	if err != nil {
		t.Fatalf("listing days of week: %v", err)
	}

	// Check query parameters
	query := hdl.Request.URL.Query()
	if query.Get("page") != strconv.Itoa(req.Page.Page) {
		t.Errorf("request page: expected=%d; got=%s", req.Page.Page, query.Get("page"))
	}
	if query.Get("page_size") != strconv.Itoa(req.Page.PageSize) {
		t.Errorf("request page_size: expected=%d; got=%s", req.Page.PageSize, query.Get("page_size"))
	}
	if query.Get("sort") != "day_of_week_id" {
		t.Errorf("request sort: expected=%s; got=%s", "day_of_week_id", query.Get("sort"))
	}

	// Check response
	if len(resp.DaysOfWeek) != 2 {
		t.Errorf("response days of week list length: expected=%d; got=%d", 2, len(resp.DaysOfWeek))
	}
}
