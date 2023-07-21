package naturallanguage

const Unknown = -1

// Human strings indicating the time period
const (
	Day = iota
	Week
	Month
	Year
	Hour
	Minute
	Second
)

func parseTimePeriod(s string) int {
	switch s {
	case "day", "days":
		return Day
	case "week", "weeks":
		return Week
	case "month", "months":
		return Month
	case "year", "years":
		return Year
	case "hour", "hours":
		return Hour
	case "minute", "minutes":
		return Minute
	case "second", "seconds":
		return Second
	}
	return Unknown
}

const (
	tomorrow = iota
	yesterday
	today
)

func parseDayShortcuts(s string) int {
	switch s {
	case "tomorrow":
		return tomorrow
	case "yesterday":
		return yesterday
	case "today":
		return today
	}
	return Unknown
}
