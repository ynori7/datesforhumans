package naturallanguage

import (
	"strings"
	"time"
)

const (
	Monday    = "monday"
	Tuesday   = "tuesday"
	Wednesday = "wednesday"
	Thursday  = "thursday"
	Friday    = "friday"
	Saturday  = "saturday"
	Sunday    = "sunday"
)

//startsWithWeekday returns true if the string starts with a weekday
func startsWithWeekday(s string) bool {
	return strings.HasPrefix(s, Monday) || strings.HasPrefix(s, Tuesday) || strings.HasPrefix(s, Wednesday) ||
		strings.HasPrefix(s, Thursday) || strings.HasPrefix(s, Friday) || strings.HasPrefix(s, Saturday) || strings.HasPrefix(s, Sunday)
}

//parseWeekDay returns the weekday number (0-6) if the input string is a weekday, otherwise returns Unknown
func parseWeekDay(s string) int {
	switch s {
	case Monday:
		return int(time.Monday)
	case Tuesday:
		return int(time.Tuesday)
	case Wednesday:
		return int(time.Wednesday)
	case Thursday:
		return int(time.Thursday)
	case Friday:
		return int(time.Friday)
	case Saturday:
		return int(time.Saturday)
	case Sunday:
		return int(time.Sunday)
	}

	return Unknown
}
