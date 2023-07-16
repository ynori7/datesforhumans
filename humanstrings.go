package datesforhumans

import (
	"strconv"
	"strings"
	"time"
)

// dateConfig represents a date configuration parsed from a natural language string
type dateConfig struct {
	direction int
	weekDay   int
	month     int
	period    int
	amount    int //will be some number of days, weeks, months, years, hours, minutes, or seconds
}

func (d dateConfig) isEmpty() bool {
	return d.direction == 0 && d.weekDay == unknown && d.month == unknown && d.period == unknown
}

// ParseNaturalDateString parses a natural language date string and returns the date configuration
func parseNaturalDateString(s string) (dateConfig, timeOfDay) {
	s = strings.ToLower(strings.TrimSpace(s))
	timeOfDay := defaultTimeOfDay

	date := dateConfig{
		direction: parseDirection(s),
		weekDay:   unknown,
		month:     unknown,
		period:    unknown,
		amount:    unknown,
	}

	parts := strings.Split(s, " "+timeOfDayIndicator+" ")
	if len(parts) == 2 {
		timeOfDay = parseTime(parts[1])
		s = strings.TrimSpace(parts[0])
	}

	if d := parseDayShortcuts(s); d != unknown {
		date.period = day
		date.amount = 1
		switch d {
		case tomorrow:
			date.direction = next
		case yesterday:
			date.direction = last
		}
		return date, timeOfDay
	}

	parts = strings.Split(s, " ")
	for _, p := range parts {
		p = strings.TrimSpace(p)

		if d := parseWeekDay(p); d != unknown {
			date.weekDay = d
			continue
		}

		if d := parseMonth(p); d != unknown {
			date.month = d
			continue
		}

		if d := parseTimePeriod(p); d != unknown {
			date.period = d
			continue
		}

		if i, err := strconv.Atoi(p); err == nil {
			date.amount = i
			continue
		}
	}

	if date.direction == 0 {
		date.direction = next
	}

	return date, timeOfDay
}

const unknown = -1

const (
	next = 1
	last = -1
)

var directionIndicators = map[string]int{
	"next":     next,
	"last":     last,
	"this":     next,
	"ago":      last,
	"from now": next,
	"in":       next,
}

func parseDirection(s string) int {
	for k, v := range directionIndicators {
		if strings.HasPrefix(s, k) || strings.HasSuffix(s, k) {
			return v
		}
	}

	return 0
}

func parseWeekDay(s string) int {
	switch s {
	case "monday":
		return int(time.Monday)
	case "tuesday":
		return int(time.Tuesday)
	case "wednesday":
		return int(time.Wednesday)
	case "thursday":
		return int(time.Thursday)
	case "friday":
		return int(time.Friday)
	case "saturday":
		return int(time.Saturday)
	case "sunday":
		return int(time.Sunday)
	}

	return unknown
}

func parseMonth(s string) int {
	switch s {
	case "january":
		return int(time.January)
	case "february":
		return int(time.February)
	case "march":
		return int(time.March)
	case "april":
		return int(time.April)
	case "may":
		return int(time.May)
	case "june":
		return int(time.June)
	case "july":
		return int(time.July)
	case "august":
		return int(time.August)
	case "september":
		return int(time.September)
	case "october":
		return int(time.October)
	case "november":
		return int(time.November)
	case "december":
		return int(time.December)
	}
	return unknown
}

// Human strings indicating the time period
const (
	day = iota
	week
	month
	year
	hour
	minute
	second
)

func parseTimePeriod(s string) int {
	switch s {
	case "day", "days":
		return day
	case "week", "weeks":
		return week
	case "month", "months":
		return month
	case "year", "years":
		return year
	case "hour", "hours":
		return hour
	case "minute", "minutes":
		return minute
	case "second", "seconds":
		return second
	}
	return unknown
}

const (
	tomorrow = iota
	yesterday
)

func parseDayShortcuts(s string) int {
	switch s {
	case "tomorrow":
		return tomorrow
	case "yesterday":
		return yesterday
	}
	return unknown
}

const (
	timeOfDayIndicator = "at"
)

const (
	am = "am"
	pm = "pm"
)

var defaultTimeOfDay = timeOfDay{hour: -1}

type timeOfDay struct {
	hour   int
	minute int
	second int
}

func (t timeOfDay) isEmpty() bool {
	return t.hour < 0
}

func parseTime(s string) timeOfDay {
	s = strings.ToLower(strings.TrimSpace(s))

	amPmModifier := 0
	s = strings.TrimSuffix(s, am)
	if strings.HasSuffix(s, pm) {
		s = strings.TrimSuffix(s, pm)
		amPmModifier = 12
	}
	s = strings.TrimSpace(s) //in case we trimmed off the am/pm

	parts := strings.Split(s, ":")
	if len(parts) > 3 || len(parts) == 0 {
		return defaultTimeOfDay
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return defaultTimeOfDay
	}
	if hour == 12 && amPmModifier == 12 {
		amPmModifier = 0
	} else if hour == 12 && amPmModifier == 0 {
		amPmModifier = -12
	}

	if len(parts) == 1 {
		return timeOfDay{
			hour:   hour + amPmModifier,
			minute: 0,
			second: 0,
		}
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return defaultTimeOfDay
	}

	second := 0
	if len(parts) == 3 {
		second, err = strconv.Atoi(parts[2])
		if err != nil {
			return defaultTimeOfDay
		}
	}

	return timeOfDay{
		hour:   hour + amPmModifier,
		minute: minute,
		second: second,
	}
}
