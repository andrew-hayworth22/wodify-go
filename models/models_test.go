package models_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/andrew-hayworth22/wodify-go/models"
)

var testDate time.Time
var testDateTime time.Time

const dateString = "2014-04-15"
const dateTimeString = "2014-04-15T13:34:23"

func init() {
	testDate = time.Date(2014, 4, 15, 0, 0, 0, 0, time.UTC)
	testDateTime = time.Date(2014, 4, 15, 13, 34, 23, 0, time.UTC)
}

func TestNewDate(t *testing.T) {
	date := models.NewDate(testDate)
	if date.Compare(testDate) != 0 {
		t.Errorf("expected=%s; got=%s", testDate, date)
	}
}

func TestDate_String(t *testing.T) {
	date := models.NewDate(testDate)
	if date.String() != dateString {
		t.Errorf("expected=%s; got=%s", dateString, date.String())
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	var date models.Date
	err := json.Unmarshal([]byte(`"2014-04-15"`), &date)
	if err != nil {
		t.Fatalf("error unmarshalling: %v", err)
	}
	if date.Compare(testDate) != 0 {
		t.Errorf("expected=%s; got=%s", testDate, date)
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	date := models.NewDate(testDate)
	marshalled, err := json.Marshal(date)
	if err != nil {
		t.Fatalf("error marshalling: %v", err)
	}

	expected := fmt.Sprintf(`"%s"`, dateString)
	if string(marshalled) != expected {
		t.Errorf("expected=%s; got=%s", expected, string(marshalled))
	}
}

func TestNewDateTime(t *testing.T) {
	dateTime := models.NewDateTime(testDateTime)
	if dateTime.Compare(testDateTime) != 0 {
		t.Errorf("expected=%s; got=%s", testDateTime, dateTime)
	}
}

func TestDateTime_String(t *testing.T) {
	dateTime := models.NewDateTime(testDateTime)
	if dateTime.String() != dateTimeString {
		t.Errorf("expected=%s; got=%s", dateTimeString, dateTime.String())
	}
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	var dateTime models.DateTime
	err := json.Unmarshal([]byte(`"2014-04-15T13:34:23"`), &dateTime)
	if err != nil {
		t.Fatalf("error unmarshalling: %v", err)
	}
	if dateTime.Compare(testDateTime) != 0 {
		t.Errorf("expected=%s; got=%s", testDateTime, dateTime)
	}
}

func TestDateTime_MarshalJSON(t *testing.T) {
	dateTime := models.NewDateTime(testDateTime)
	marshalled, err := json.Marshal(dateTime)
	if err != nil {
		t.Fatalf("error marshalling: %v", err)
	}

	expected := fmt.Sprintf(`"%s"`, dateTimeString)
	if string(marshalled) != expected {
		t.Errorf("expected=%s; got=%s", expected, string(marshalled))
	}
}

func TestPaginationRequest_ToQuery(t *testing.T) {
	paginationRequest := models.PaginationRequest{
		Page:     1,
		PageSize: 10,
	}
	query := paginationRequest.ToQuery()
	page := query.Get("page")
	pageSize := query.Get("page_size")
	if page != "1" {
		t.Errorf("page expected=%s; got=%s", "1", page)
	}
	if pageSize != "10" {
		t.Errorf("page_size expected=%s; got=%s", "10", pageSize)
	}
}
