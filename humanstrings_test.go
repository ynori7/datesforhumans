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
				amount:    1,
			},
		},
		{
			in: "next year",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    year,
				amount:    1,
			},
		},
		{
			in: "next week",
			want: dateConfig{
				direction: next,
				weekDay:   unknown,
				month:     unknown,
				period:    week,
				amount:    1,
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
				amount:    1,
			},
		},
		{
			in: "last year",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    year,
				amount:    1,
			},
		},
		{
			in: "last week",
			want: dateConfig{
				direction: last,
				weekDay:   unknown,
				month:     unknown,
				period:    week,
				amount:    1,
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
		got := parseNaturalDateString(test.in)
		if got != test.want {
			t.Errorf("ParseNaturalDateString(%q) = %v, want %v", test.in, got, test.want)
		}
	}
}
