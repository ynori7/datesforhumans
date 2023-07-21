package datesforhumans

import (
	"time"

	"github.com/ynori7/datesforhumans/internal/naturallanguage"
)

// Time is a wrapper around the stdlib time object. In case parsing was not possible, IsValid will be false
type Time struct {
	t       time.Time
	IsValid bool

	dateString string
	timeString string
}

// Time returns the underlying stdlib time object
func (t Time) Time() time.Time {
	return t.t
}

// At parses a natural language time string and returns a new time object at the specified time
func (t Time) At(s string) Time {
	timeConfig := naturallanguage.ParseTime(s)
	t = t.at(timeConfig)
	t.timeString = s

	return t
}

// at sets the time of day on the time object
func (t Time) at(timeConfig naturallanguage.TimeOfDay) Time {
	if !timeConfig.IsEmpty() {
		t.t = time.Date(t.t.Year(), t.t.Month(), t.t.Day(), timeConfig.Hour, timeConfig.Minute, timeConfig.Second, 0, t.t.Location())
	}

	return t
}

// Repeat returns a slice of times that repeat until the specified time
func (t Time) Repeat(until time.Time) []Time {
	var times []Time
	times = append(times, t)

	current := t.t
	for current.Before(until) {
		t2 := ParseDate(current, t.dateString).At(t.timeString)
		if t2.t.After(until) {
			break
		}
		times = append(times, t2)
		current = t2.t
	}

	return times
}
