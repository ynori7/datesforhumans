package datesforhumans

import "time"

// ParseDate parses a natural language date string and returns the date time object
func ParseDate(from time.Time, s string) Time {
	dateConfig := parseNaturalDateString(s)

	if dateConfig.isEmpty() {
		return Time{t: from}
	}

	if dateConfig.period != unknown {
		switch dateConfig.period {
		case day:
			return Time{t: from.AddDate(0, 0, dateConfig.direction*dateConfig.amount), dateString: s}
		case week:
			return Time{t: from.AddDate(0, 0, dateConfig.direction*7*dateConfig.amount), dateString: s}
		case month:
			return Time{t: from.AddDate(0, dateConfig.direction*dateConfig.amount, 0), dateString: s}
		case year:
			return Time{t: from.AddDate(dateConfig.direction*dateConfig.amount, 0, 0), dateString: s}
		case hour:
			return Time{t: from.Add(time.Duration(dateConfig.direction*dateConfig.amount) * time.Hour), dateString: s}
		case minute:
			return Time{t: from.Add(time.Duration(dateConfig.direction*dateConfig.amount) * time.Minute), dateString: s}
		case second:
			return Time{t: from.Add(time.Duration(dateConfig.direction*dateConfig.amount) * time.Second), dateString: s}
		}
	}

	if dateConfig.weekDay != unknown {
		return Time{t: getNextWeekday(from, time.Weekday(dateConfig.weekDay), dateConfig.direction), dateString: s}
	}

	if dateConfig.month != unknown {
		return Time{t: getNextMonth(from, time.Month(dateConfig.month), dateConfig.direction), dateString: s}
	}

	return Time{t: from, dateString: s}
}

func getNextWeekday(date time.Time, weekday time.Weekday, direction int) time.Time {
	currentWeekday := date.Weekday()
	daysDiff := int(weekday - currentWeekday)
	if daysDiff <= 0 {
		daysDiff += 7
	}

	if direction == last {
		daysDiff -= 7
		if daysDiff == 0 {
			daysDiff = -7
		}
	}

	return date.AddDate(0, 0, daysDiff)
}

func getNextMonth(date time.Time, month time.Month, direction int) time.Time {
	currentMonth := date.Month()
	monthsDiff := int(month - currentMonth)
	if monthsDiff <= 0 {
		monthsDiff += 12
	}

	if direction == last {
		monthsDiff -= 12
		if monthsDiff == 0 {
			monthsDiff = -12
		}
	}

	return date.AddDate(0, monthsDiff, 0)
}
