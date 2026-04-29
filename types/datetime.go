package types

import (
	"strings"
	"time"
)

const dateTimeLayout = "2006-01-02T15:04:05"

// DateTime is a date-time value that marshals/unmarshals as "YYYY-MM-DD HH:MM:SS" to match Wodify's API.
type DateTime struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `""`)

	t, err := time.Parse(dateTimeLayout, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte{}, nil
	}
	return []byte(`"` + d.Format(dateTimeLayout) + `"`), nil
}

// String returns the date as a string in "YYYY-MM-DD" format.
func (d *DateTime) String() string {
	return d.Format(dateTimeLayout)
}
