package models

import (
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// Date is a date-only value with no time component that marshals/unmarshals as "YYYY-MM-DD" to match Wodify's API.
type Date struct {
	time.Time
}

// NewDate creates a new date that can be used with the Wodify API.
func NewDate(time time.Time) Date {
	return Date{Time: time}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `""`)

	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte{}, nil
	}
	return []byte(`"` + d.Format(dateLayout) + `"`), nil
}

// String returns the date as a string in "YYYY-MM-DD" format.
func (d *Date) String() string {
	return d.Format(dateLayout)
}
