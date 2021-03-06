package types

import (
	"encoding/json"
	"strings"
	"time"
)

// Time is the time responsed by Plurk API 2.0.
// Its format is `Mon, 02 Jan 2006 15:04:05 MST`
type Time time.Time

// UnmarshalJSON implement Unmarshaler for time from Plurk API 2.0.
func (t *Time) UnmarshalJSON(b []byte) error {
	if tp, err := time.Parse(time.RFC1123, strings.Trim(string(b), "\"")); err == nil {
		*t = Time(tp)
	} else {
		return err
	}
	return nil
}

// MarshalJSON implement Marshaler for time from Plurk API 2.0.
func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

// Format time from Plurk API 2.0
func (t *Time) Format(s string) string {
	return time.Time(*t).Format(s)
}
