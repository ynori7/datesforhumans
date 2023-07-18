package datesforhumans

import (
	"testing"
	"time"
)

func TestParseNaturalDateString(t *testing.T) {
	tests := []struct {
		in   string
		want dateConfig
	}{
		{
			in: "tomorrow",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
		},
		{
			in: "yesterday",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
		},
		{
			in: "today",
			want: dateConfig{
				direction: current,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
		},
		{
			in: "next monday",
			want: dateConfig{
				direction: next,
				weekDay:   int(time.Monday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "next month",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    month,
				amount:    unknown,
			},
		},
		{
			in: "next year",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    year,
				amount:    unknown,
			},
		},
		{
			in: "next week",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    week,
				amount:    unknown,
			},
		},
		{
			in: "next tuesday",
			want: dateConfig{
				direction: next,
				weekDay:   int(time.Tuesday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "next sunday",
			want: dateConfig{
				direction: next,
				weekDay:   int(time.Sunday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "last monday",
			want: dateConfig{
				direction: last,
				weekDay:   int(time.Monday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "yesterday",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
		},
		{
			in: "last month",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    month,
				amount:    unknown,
			},
		},
		{
			in: "last year",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    year,
				amount:    unknown,
			},
		},
		{
			in: "last week",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    week,
				amount:    unknown,
			},
		},
		{
			in: "last tuesday",
			want: dateConfig{
				direction: last,
				weekDay:   int(time.Tuesday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "last january",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     int(time.January),
				period:    unknown,
				amount:    unknown,
			},
		},
		{
			in: "3 months ago",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    month,
				amount:    3,
			},
		},
		{
			in: "1 year ago",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    year,
				amount:    1,
			},
		},
		{
			in: "2 days from now",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    2,
			},
		},
		{
			in: "in 6 hours",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    hour,
				amount:    6,
			},
		},
		{
			in: "in 22 minutes",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    minute,
				amount:    22,
			},
		},
	}

	for _, test := range tests {
		got, timeOfDay := parseNaturalDateString(test.in)
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
		want     dateConfig
		wantTime timeOfDay
	}{
		{
			in: "tomorrow at 10pm",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
			wantTime: timeOfDay{hour: 22, minute: 0, second: 0},
		},
		{
			in: "yesterday at 1am",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
			wantTime: timeOfDay{hour: 1, minute: 0, second: 0},
		},
		{
			in: "next monday at 10:00",
			want: dateConfig{
				direction: next,
				weekDay:   int(time.Monday),
				month:     unknown,
				period:    unknown,
				amount:    unknown,
			},
			wantTime: timeOfDay{hour: 10, minute: 0, second: 0},
		},
		{
			in: "next month at 11:30:00",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    month,
				amount:    unknown,
			},
			wantTime: timeOfDay{hour: 11, minute: 30, second: 0},
		},
		{
			in: "tomorrow at midnight",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
			wantTime: timeOfDay{hour: 0, minute: 0, second: 0},
		},
		{
			in: "yesterday at noon",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    day,
				amount:    1,
			},
			wantTime: timeOfDay{hour: 12, minute: 0, second: 0},
		},
	}

	for _, test := range tests {
		got, timeOfDay := parseNaturalDateString(test.in)
		if got != test.want {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, got, test.want)
		}
		if timeOfDay != test.wantTime {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, timeOfDay, test.wantTime)
		}
	}
}
