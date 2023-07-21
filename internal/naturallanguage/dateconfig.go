package naturallanguage

import (
	"strconv"
	"strings"
)

// DateConfig represents a date configuration parsed from a natural language string
type DateConfig struct {
	Direction int
	WeekDay   int
	Month     int
	Period    int
	Amount    int //will be some number of days, weeks, months, years, hours, minutes, or seconds

	IsValid bool
}

var defaultDate = DateConfig{
	Direction: UnknownDirection,
	WeekDay:   Unknown,
	Month:     Unknown,
	Period:    Unknown,
	Amount:    Unknown,
	IsValid:   false,
}

func (d DateConfig) IsEmpty() bool {
	return d.Direction == UnknownDirection && d.WeekDay == Unknown && d.Month == Unknown && d.Period == Unknown
}

// ParseNaturalDateString parses a natural language date string and returns the date configuration
// Date must be in the format [next|last|this|in] (amount) [weekday|month|period] [ago] (at) (time of day) where the amount and time of day are optional (depending on the situation)
func ParseNaturalDateString(s string) (DateConfig, TimeOfDay) {
	s = strings.ToLower(strings.TrimSpace(s))
	timeOfDay := defaultTimeOfDay

	date := defaultDate

	//If there is a time of day, parse it out
	parts := strings.Split(s, " "+timeOfDayIndicator+" ")
	if len(parts) == 2 {
		timeOfDay = ParseTime(parts[1])
		s = strings.TrimSpace(parts[0])
		if timeOfDay.IsEmpty() {
			return defaultDate, defaultTimeOfDay
		}
	}
	if len(parts) > 2 { //invalid string
		return defaultDate, defaultTimeOfDay
	}

	//handle the case where the time of day is at the beginning of the string ("at 10am" is equivalent of "today at 10am")
	if len(parts) == 1 && strings.HasPrefix(s, timeOfDayIndicator+" ") {
		timeOfDay = ParseTime(strings.TrimPrefix(s, timeOfDayIndicator+" "))
		if timeOfDay.IsEmpty() {
			return defaultDate, defaultTimeOfDay
		}
		date.Period = Day
		date.Amount = 1
		date.Direction = Current
		date.IsValid = true
		return date, timeOfDay
	}

	//Parse day shortcuts like "today" or "tomorrow"
	if d := parseDayShortcuts(s); d != Unknown {
		date.Period = Day
		date.Amount = 1
		switch d {
		case tomorrow:
			date.Direction = Next
		case yesterday:
			date.Direction = Last
		case today:
			date.Direction = Current
		}
		date.IsValid = true
		return date, timeOfDay
	}

	//Parse the direction
	date.Direction = parseDirection(s)
	if date.Direction == UnknownDirection { //invalid string
		return defaultDate, defaultTimeOfDay
	}

	//Parse the date
	found := 0
	parts = strings.Split(s, " ")
	for _, p := range parts {
		p = strings.TrimSpace(p)

		if d := parseWeekDay(p); d != Unknown {
			date.WeekDay = d
			found++
			continue
		}

		if d := parseMonth(p); d != Unknown {
			date.Month = d
			found++
			continue
		}

		if d := parseTimePeriod(p); d != Unknown {
			date.Period = d
			found++
			continue
		}

		if i, err := strconv.Atoi(p); err == nil {
			date.Amount = i
			continue
		}
	}
	if found != 1 { //invalid string
		return defaultDate, defaultTimeOfDay
	}

	date.IsValid = true

	return date, timeOfDay
}

