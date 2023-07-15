package datesforhumans

import "time"

// Time is a wrapper around the stdlib time object
type Time struct {
	t time.Time
}

// Time returns the underlying stdlib time object
func (t Time) Time() time.Time {
	return t.t
}

// At parses a natural language time string and returns a new time object at the specified time
func (t Time) At(s string) Time {
	timeConfig := parseTime(s)
	if !timeConfig.isEmpty() {
		t.t = time.Date(t.t.Year(), t.t.Month(), t.t.Day(), timeConfig.hour, timeConfig.minute, timeConfig.second, 0, t.t.Location())
	}

	return t
}
