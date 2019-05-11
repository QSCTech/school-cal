package schoolcal

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime is a Time for unmarshaling from json in layout of ISO8601
type CustomTime struct {
	time.Time
}

// ISO8601Layout is a time layout as ISO8601
const ISO8601Layout = "2006-01-02T15:04-07:00"

var nilTime = (time.Time{}).UnixNano()

// UnmarshalJSON is method to get CustomTime for JSON
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ISO8601Layout, s)
	return
}

// MarshalJSON is method to get JSON for CustomTime
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ISO8601Layout))), nil
}

// IsSet is method to check if CustomTime is set
func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
