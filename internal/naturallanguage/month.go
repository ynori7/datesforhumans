package naturallanguage

import (
	"strings"
	"time"
)

const (
	January   = "january"
	February  = "february"
	March     = "march"
	April     = "april"
	May       = "may"
	June      = "june"
	July      = "july"
	August    = "august"
	September = "september"
	October   = "october"
	November  = "november"
	December  = "december"
)

//returns true if the input string begins with a month name
func startsWithMonth(s string) bool {
	return strings.HasPrefix(s, January) || strings.HasPrefix(s, February) || strings.HasPrefix(s, March) ||
		strings.HasPrefix(s, April) || strings.HasPrefix(s, May) || strings.HasPrefix(s, June) || strings.HasPrefix(s, July) ||
		strings.HasPrefix(s, August) || strings.HasPrefix(s, September) || strings.HasPrefix(s, October) ||
		strings.HasPrefix(s, November) || strings.HasPrefix(s, December)
}

//returns the month number (1-12) if the input string is a month name, otherwise returns Unknown
func parseMonth(s string) int {
	switch s {
	case January:
		return int(time.January)
	case February:
		return int(time.February)
	case March:
		return int(time.March)
	case April:
		return int(time.April)
	case May:
		return int(time.May)
	case June:
		return int(time.June)
	case July:
		return int(time.July)
	case August:
		return int(time.August)
	case September:
		return int(time.September)
	case October:
		return int(time.October)
	case November:
		return int(time.November)
	case December:
		return int(time.December)
	}
	return Unknown
}
