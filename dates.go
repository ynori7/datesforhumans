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
			return Time{t: from.AddDate(0, 0, dateConfig.direction*getAmountOrOne(dateConfig)), dateString: s}
		case week:
			return Time{t: from.AddDate(0, 0, dateConfig.direction*7*getAmountOrOne(dateConfig)), dateString: s}
		case month:
			if dateConfig.amount == unknown {
				return Time{t: time.Date(from.Year(), from.Month()+time.Month(dateConfig.direction), 1, 0, 0, 0, 0, from.Location()), dateString: s}
			}
			return Time{t: from.AddDate(0, dateConfig.direction*dateConfig.amount, 0), dateString: s}
		case year:
			return Time{t: from.AddDate(dateConfig.direction*getAmountOrOne(dateConfig), 0, 0), dateString: s}
		case hour:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Hour), dateString: s}
		case minute:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Minute), dateString: s}
		case second:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Second), dateString: s}
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

	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location()).AddDate(0, monthsDiff, 0)
}

func getAmountOrOne(dateConfig dateConfig) int {
	if dateConfig.amount == unknown {
		return 1
	}
	return dateConfig.amount
}
