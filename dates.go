package datesforhumans

import "time"

// ParseDate parses a natural language date string and returns the date time object
func ParseDate(from time.Time, s string) Time {
	dateConfig, timeConfig := parseNaturalDateString(s)

	if dateConfig.isEmpty() {
		return Time{t: from}
	}

	//Use the current time if no time was specified
	if timeConfig.isEmpty() {
		timeConfig.hour = from.Hour()
		timeConfig.minute = from.Minute()
		timeConfig.second = from.Second()
	}

	//This is a "period" like "next week" or "last month"
	if dateConfig.period != unknown {
		switch dateConfig.period {
		case day:
			return Time{t: from.AddDate(0, 0, dateConfig.direction*getAmountOrOne(dateConfig)), dateString: s}.at(timeConfig)
		case week:
			return Time{t: from.AddDate(0, 0, dateConfig.direction*7*getAmountOrOne(dateConfig)), dateString: s}.at(timeConfig)
		case month:
			if dateConfig.amount == unknown {
				return Time{t: time.Date(from.Year(), from.Month()+time.Month(dateConfig.direction), 1, 0, 0, 0, 0, from.Location()), dateString: s}
			}
			return Time{t: from.AddDate(0, dateConfig.direction*dateConfig.amount, 0), dateString: s}.at(timeConfig)
		case year:
			return Time{t: from.AddDate(dateConfig.direction*getAmountOrOne(dateConfig), 0, 0), dateString: s}.at(timeConfig)
		case hour:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Hour), dateString: s}
		case minute:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Minute), dateString: s}
		case second:
			return Time{t: from.Add(time.Duration(dateConfig.direction*getAmountOrOne(dateConfig)) * time.Second), dateString: s}
		}
	}

	//This is a "weekday" like "next monday" or "last tuesday"
	if dateConfig.weekDay != unknown {
		return Time{t: getNextWeekday(from, time.Weekday(dateConfig.weekDay), dateConfig.direction), dateString: s}.at(timeConfig)
	}

	//This is a "month" like "next january" or "last february"
	if dateConfig.month != unknown {
		return Time{t: getNextMonth(from, time.Month(dateConfig.month), dateConfig.direction), dateString: s}
	}

	return Time{t: from, dateString: s}.at(timeConfig)
}

func getNextWeekday(date time.Time, weekday time.Weekday, direction int) time.Time {
	currentWeekday := date.Weekday()
	daysDiff := int(weekday - currentWeekday)
	if daysDiff <= 0 {
		daysDiff += 7
	}

	if direction == last {
		daysDiff -= 7
		if daysDiff == 0 { //if it's the same day of the week, go back a week
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
		if monthsDiff == 0 { //if it's the same month, go back a year
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
