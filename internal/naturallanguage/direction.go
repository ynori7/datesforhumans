package naturallanguage

import "strings"

const (
	Next             = 1
	Last             = -1
	Current          = 0
	UnknownDirection = -2
)

var directionIndicators = map[string]int{
	"next":     Next,
	"last":     Last,
	"this":     Current,
	"ago":      Last,
	"from now": Next,
	"in":       Next,
}

func parseDirection(s string) int {
	prefixDirection := UnknownDirection
	suffixDirection := UnknownDirection

	for k, v := range directionIndicators {
		if strings.HasPrefix(s, k) {
			prefixDirection = v
		}
		if strings.HasSuffix(s, k) {
			suffixDirection = v
		}
	}

	if prefixDirection != UnknownDirection && suffixDirection != UnknownDirection {
		return UnknownDirection //can't have two direction indicators
	}

	if prefixDirection != UnknownDirection {
		return prefixDirection
	}

	if suffixDirection != UnknownDirection {
		return suffixDirection
	}

	if startsWithWeekday(s) || startsWithMonth(s){
		return Current
	}

	return UnknownDirection
}
