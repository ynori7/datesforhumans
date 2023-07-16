package datesforhumans

import "time"

// Time is a wrapper around the stdlib time object
type Time struct {
	t time.Time

	dateString string
	timeString string
}

// Time returns the underlying stdlib time object
func (t Time) Time() time.Time {
	return t.t
}

// At parses a natural language time string and returns a new time object at the specified time
func (t Time) At(s string) Time {
	timeConfig := parseTime(s)
	t = t.at(timeConfig)
	t.timeString = s

	return t
}

func (t Time) at(timeConfig timeOfDay) Time {
	if !timeConfig.isEmpty() {
		t.t = time.Date(t.t.Year(), t.t.Month(), t.t.Day(), timeConfig.hour, timeConfig.minute, timeConfig.second, 0, t.t.Location())
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

// Range is a tuple of a starting and ending time
type Range struct {
	Start Time
	End   Time

	startString string
	endString   string
}

// ParseRange parses a natural language date string and returns the date time object
func ParseRange(from time.Time, start string, end string) Range {
	r := Range{
		startString: start,
		endString:   end,
	}
	r.Start = ParseDate(from, start)
	r.End = ParseDate(r.Start.t, end)

	return r
}

// Repeat returns a slice of ranges that repeat until the specified time
func (r Range) Repeat(until time.Time) []Range {
	var ranges []Range
	ranges = append(ranges, r)

	current := r.Start.t
	for current.Before(until) {
		r2 := ParseRange(current, r.startString, r.endString)
		if r2.Start.t.After(until) {
			break
		}
		ranges = append(ranges, r2)
		current = r2.Start.t
	}

	return ranges
}
