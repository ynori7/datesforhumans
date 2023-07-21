package naturallanguage

import (
	"strconv"
	"strings"
)

const (
	timeOfDayIndicator = "at"
)

const (
	am = "am"
	pm = "pm"
)

var defaultTimeOfDay = TimeOfDay{Hour: -1, IsValid: true}

type TimeOfDay struct {
	Hour   int
	Minute int
	Second int

	IsValid bool
}

func (t TimeOfDay) IsEmpty() bool {
	return t.Hour < 0
}

func ParseTime(s string) TimeOfDay {
	s = strings.ToLower(strings.TrimSpace(s))

	// Parse time shortcuts like "midnight" or "noon"
	if d := parseTimeShortcuts(s); d != Unknown {
		switch d {
		case midnight:
			return TimeOfDay{Hour: 0, IsValid: true}
		case noon:
			return TimeOfDay{Hour: 12, IsValid: true}
		}
	}

	// Parse am/pm
	hasAmPm := strings.HasSuffix(s, am) || strings.HasSuffix(s, pm)
	amPmModifier := 0
	s = strings.TrimSuffix(s, am)
	if strings.HasSuffix(s, pm) {
		s = strings.TrimSuffix(s, pm)
		amPmModifier = 12
	}
	s = strings.TrimSpace(s) //in case we trimmed off the am/pm

	// Parse the time
	parts := strings.Split(s, ":")
	if len(parts) > 3 || len(parts) == 0 {
		return defaultTimeOfDay
	}

	// Parse the hour
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return defaultTimeOfDay
	}
	if hour < 0 || (!hasAmPm && hour > 24) || (hasAmPm && hour > 12) { //validate hour
		return defaultTimeOfDay
	}
	// Handle am/pm when hour = 12
	if hour == 12 && amPmModifier == 12 {
		amPmModifier = 0
	} else if hour == 12 && amPmModifier == 0 {
		amPmModifier = -12
	}

	if len(parts) == 1 {
		return TimeOfDay{
			Hour:    hour + amPmModifier,
			Minute:  0,
			Second:  0,
			IsValid: true,
		}
	}

	// Parse the minute
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return defaultTimeOfDay
	}
	if minute < 0 || minute > 59 { //validate minute
		return defaultTimeOfDay
	}

	// Parse the second
	second := 0
	if len(parts) == 3 {
		second, err = strconv.Atoi(parts[2])
		if err != nil {
			return defaultTimeOfDay
		}
		if second < 0 || second > 59 { //validate second
			return defaultTimeOfDay
		}
	}

	return TimeOfDay{
		Hour:    hour + amPmModifier,
		Minute:  minute,
		Second:  second,
		IsValid: true,
	}
}

const (
	midnight = iota
	noon
)

func parseTimeShortcuts(s string) int {
	switch s {
	case "midnight":
		return midnight
	case "noon":
		return noon
	}
	return Unknown
}
