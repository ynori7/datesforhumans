package datesforhumans

import (
	"time"

	"github.com/ynori7/datesforhumans/internal/naturallanguage"
)

// FrowNow transforms the current time based on the input string. This is a shorthand for ParseDate(time.Now(), s)
func FromNow(s string) Time {
	return ParseDate(time.Now(), s)
}

// ParseDate parses a natural language date string and returns the date time object
func ParseDate(from time.Time, s string) Time {
	dateConfig, timeConfig := naturallanguage.ParseNaturalDateString(s)

	if dateConfig.IsEmpty() {
		return Time{t: from}
	}

	//Use the current time if no time was specified
	if timeConfig.IsEmpty() {
		timeConfig.Hour = from.Hour()
		timeConfig.Minute = from.Minute()
		timeConfig.Second = from.Second()
	}

	t := Time{dateString: s, IsValid: dateConfig.IsValid}

	//This is a "period" like "next week" or "last month"
	if dateConfig.Period != naturallanguage.Unknown {
		switch dateConfig.Period {
		case naturallanguage.Day:
			t.t = from.AddDate(0, 0, dateConfig.Direction*getAmountOrOne(dateConfig))
			return t.at(timeConfig)
		case naturallanguage.Week:
			t.t = from.AddDate(0, 0, dateConfig.Direction*7*getAmountOrOne(dateConfig))
			return t.at(timeConfig)
		case naturallanguage.Month:
			if dateConfig.Amount == naturallanguage.Unknown {
				t.t = time.Date(from.Year(), from.Month()+time.Month(dateConfig.Direction), 1, 0, 0, 0, 0, from.Location())
				return t
			}
			t.t = from.AddDate(0, dateConfig.Direction*dateConfig.Amount, 0)
			return t.at(timeConfig)
		case naturallanguage.Year:
			t.t = from.AddDate(dateConfig.Direction*getAmountOrOne(dateConfig), 0, 0)
			return t.at(timeConfig)
		case naturallanguage.Hour:
			t.t = from.Add(time.Duration(dateConfig.Direction*getAmountOrOne(dateConfig)) * time.Hour)
			return t
		case naturallanguage.Minute:
			t.t = from.Add(time.Duration(dateConfig.Direction*getAmountOrOne(dateConfig)) * time.Minute)
			return t
		case naturallanguage.Second:
			t.t = from.Add(time.Duration(dateConfig.Direction*getAmountOrOne(dateConfig)) * time.Second)
			return t
		}
	}

	//This is a "weekday" like "next monday" or "last tuesday"
	if dateConfig.WeekDay != naturallanguage.Unknown {
		t.t = getNextWeekday(from, time.Weekday(dateConfig.WeekDay), dateConfig.Direction)
		return t.at(timeConfig)
	}

	//This is a "month" like "next january" or "last february"
	if dateConfig.Month != naturallanguage.Unknown {
		t.t = getNextMonth(from, time.Month(dateConfig.Month), dateConfig.Direction)
		return t
	}

	return Time{t: from, dateString: s}.at(timeConfig)
}

//getNextWeekday returns the next weekday in the specified direction. If the direction is 0 (current), then decides based on whether that day comes before or after the current day
func getNextWeekday(date time.Time, weekday time.Weekday, direction int) time.Time {
	currentWeekday := date.Weekday()
	if direction == 0 {
		if currentWeekday == weekday {
			return date
		}
		if currentWeekday > weekday {
			direction = naturallanguage.Last
		} else {
			direction = naturallanguage.Next
		}
	}

	daysDiff := int(weekday - currentWeekday)
	if daysDiff <= 0 {
		daysDiff += 7
	}

	if direction == naturallanguage.Last {
		daysDiff -= 7
		if daysDiff == 0 { //if it's the same day of the week, go back a week
			daysDiff = -7
		}
	}

	return date.AddDate(0, 0, daysDiff)
}

//getNextMonth returns the next month in the specified direction. If the direction is 0 (current), then decides based on whether that month comes before or after the current month
func getNextMonth(date time.Time, month time.Month, direction int) time.Time {
	currentMonth := date.Month()
	if direction == 0 {
		if currentMonth == month {
			return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
		}
		if currentMonth > month {
			direction = naturallanguage.Last
		} else {
			direction = naturallanguage.Next
		}
	}

	monthsDiff := int(month - currentMonth)
	if monthsDiff <= 0 {
		monthsDiff += 12
	}

	if direction == naturallanguage.Last {
		monthsDiff -= 12
		if monthsDiff == 0 { //if it's the same month, go back a year
			monthsDiff = -12
		}
	}

	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location()).AddDate(0, monthsDiff, 0)
}

//getAmountOrOne returns the amount if it is known, otherwise returns 1
func getAmountOrOne(dateConfig naturallanguage.DateConfig) int {
	if dateConfig.Amount == naturallanguage.Unknown {
		return 1
	}
	return dateConfig.Amount
}
