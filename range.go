package datesforhumans

import "time"

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
