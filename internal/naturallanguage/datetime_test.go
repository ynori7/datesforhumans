package naturallanguage

import (
	"testing"
	"time"
)

func TestParseNaturalDateString(t *testing.T) {
	tests := []struct {
		in   string
		want DateConfig
	}{
		{
			in: "tomorrow",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
		},
		{
			in: "yesterday",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
		},
		{
			in: "today",
			want: DateConfig{
				Direction: Current,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
		},
		{
			in: "next monday",
			want: DateConfig{
				Direction: Next,
				WeekDay:   int(time.Monday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "next month",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Month,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "next year",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Year,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "next week",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Week,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "next tuesday",
			want: DateConfig{
				Direction: Next,
				WeekDay:   int(time.Tuesday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "next sunday",
			want: DateConfig{
				Direction: Next,
				WeekDay:   int(time.Sunday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "last monday",
			want: DateConfig{
				Direction: Last,
				WeekDay:   int(time.Monday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "this monday",
			want: DateConfig{
				Direction: Current,
				WeekDay:   int(time.Monday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "yesterday",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
		},
		{
			in: "last month",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Month,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "this month",
			want: DateConfig{
				Direction: Current,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Month,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "last year",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Year,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "last week",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Week,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "last tuesday",
			want: DateConfig{
				Direction: Last,
				WeekDay:   int(time.Tuesday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "last january",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     int(time.January),
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
		},
		{
			in: "3 months ago",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Month,
				Amount:    3,
				IsValid:   true,
			},
		},
		{
			in: "1 year ago",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Year,
				Amount:    1,
				IsValid:   true,
			},
		},
		{
			in: "2 days from now",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    2,
				IsValid:   true,
			},
		},
		{
			in: "in 6 hours",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Hour,
				Amount:    6,
				IsValid:   true,
			},
		},
		{
			in: "in 22 minutes",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Minute,
				Amount:    22,
				IsValid:   true,
			},
		},
	}

	for _, test := range tests {
		got, timeOfDay := ParseNaturalDateString(test.in)
		if timeOfDay != defaultTimeOfDay {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, timeOfDay, defaultTimeOfDay)
		}
		if got != test.want {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, got, test.want)
		}
	}
}

func TestParseNaturalDateStringWithTime(t *testing.T) {
	tests := []struct {
		in       string
		want     DateConfig
		wantTime TimeOfDay
	}{
		{
			in: "tomorrow at 10pm",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 22, Minute: 0, Second: 0, IsValid: true},
		},
		{
			in: "yesterday at 1am",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 1, Minute: 0, Second: 0, IsValid: true},
		},
		{
			in: "next monday at 10:00",
			want: DateConfig{
				Direction: Next,
				WeekDay:   int(time.Monday),
				Month:     Unknown,
				Period:    Unknown,
				Amount:    Unknown,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 10, Minute: 0, Second: 0, IsValid: true},
		},
		{
			in: "next month at 11:30:00",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Month,
				Amount:    Unknown,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 11, Minute: 30, Second: 0, IsValid: true},
		},
		{
			in: "tomorrow at midnight",
			want: DateConfig{
				Direction: Next,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 0, Minute: 0, Second: 0, IsValid: true},
		},
		{
			in: "yesterday at noon",
			want: DateConfig{
				Direction: Last,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 12, Minute: 0, Second: 0, IsValid: true},
		},
		{
			in: "at noon",
			want: DateConfig{
				Direction: Current,
				WeekDay:   Unknown,
				Month:     Unknown,
				Period:    Day,
				Amount:    1,
				IsValid:   true,
			},
			wantTime: TimeOfDay{Hour: 12, Minute: 0, Second: 0, IsValid: true},
		},
	}

	for _, test := range tests {
		got, timeOfDay := ParseNaturalDateString(test.in)
		if got != test.want {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, got, test.want)
		}
		if timeOfDay != test.wantTime {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, timeOfDay, test.wantTime)
		}
	}
}

func TestParseInvalidDateString(t *testing.T) {
	tests := []string{
		"in 2 hours ago",
		"tomorrow at 10:00:00:00",
		"tomorrow at 25:00",
		"tomorrow at 25",
		"tomorrow at 13am",
		"tomorrow at 13pm",
		"tomorrow at asdf",
		"blahblah",
		"this is a test",
		"next",
		"today tomrrow",
		"next monday tuesday",
	}

	for _, test := range tests {
		got, timeOfDay := ParseNaturalDateString(test)
		if timeOfDay != defaultTimeOfDay {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test, timeOfDay, defaultTimeOfDay)
		}
		if got != defaultDate {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test, got, defaultDate)
		}
	}
}
