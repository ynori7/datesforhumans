package datesforhumans

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	base := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC) //a saturday

	tests := []struct {
		in   string
		want time.Time
	}{
		{
			in:   "tomorrow",
			want: time.Date(2023, time.July, 16, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "yesterday",
			want: time.Date(2023, time.July, 14, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "today",
			want: base,
		},
		{
			in:   "next monday",
			want: time.Date(2023, time.July, 17, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "next tuesday",
			want: time.Date(2023, time.July, 18, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "next saturday",
			want: time.Date(2023, time.July, 22, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "next sunday",
			want: time.Date(2023, time.July, 16, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "next month",
			want: time.Date(2023, time.August, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "next year",
			want: time.Date(2024, time.July, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "next week",
			want: time.Date(2023, time.July, 22, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last monday",
			want: time.Date(2023, time.July, 10, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last tuesday",
			want: time.Date(2023, time.July, 11, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "this tuesday",
			want: time.Date(2023, time.July, 11, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last saturday",
			want: time.Date(2023, time.July, 8, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "this saturday",
			want: time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "yesterday",
			want: time.Date(2023, time.July, 14, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last month",
			want: time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "last year",
			want: time.Date(2022, time.July, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last week",
			want: time.Date(2023, time.July, 8, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last january",
			want: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "this january",
			want: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "last july",
			want: time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "next july",
			want: time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "this july",
			want: time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC), // should not preserve the day and time
		},
		{
			in:   "3 months ago",
			want: time.Date(2023, time.April, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "1 year ago",
			want: time.Date(2022, time.July, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "2 days from now",
			want: time.Date(2023, time.July, 17, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "in 6 hours",
			want: time.Date(2023, time.July, 15, 23, 4, 0, 0, time.UTC),
		},
		{
			in:   "in 22 minutes",
			want: time.Date(2023, time.July, 15, 17, 26, 0, 0, time.UTC),
		},
		{
			in:   "next monday at 10pm",
			want: time.Date(2023, time.July, 17, 22, 0, 0, 0, time.UTC),
		},
		{
			in:   "last monday at 10pm",
			want: time.Date(2023, time.July, 10, 22, 0, 0, 0, time.UTC),
		},
		{
			in:   "today at 10pm",
			want: time.Date(2023, time.July, 15, 22, 0, 0, 0, time.UTC),
		},
		{
			in:   "today at noon",
			want: time.Date(2023, time.July, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			in:   "tomorrow at midnight",
			want: time.Date(2023, time.July, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			in:   "this sunday",
			want: time.Date(2023, time.July, 9, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "this friday",
			want: time.Date(2023, time.July, 14, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "this july",
			want: time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			in:   "this august",
			want: time.Date(2023, time.August, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			in:   "this june",
			want: time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		got := ParseDate(base, test.in)
		if !got.IsValid {
			t.Errorf("ParseDate(%q) = %v, want %v", test.in, got.IsValid, true)
		}
		if got.Time() != test.want {
			t.Errorf("ParseDate(%q) = %v, want %v", test.in, got.t.Format(time.DateTime), test.want.Format(time.DateTime))
		}
	}
}

func TestParseInvalidDate(t *testing.T) {
	var tests = []string{
		"in 2 hours ago",
		"tomorrow at 10:00:00:00",
		"asdf",
		"nextblah",
		"tomrrow at 25",
		"tomorrow at 13pm",
		"tomorrow at 13am",
	}

	for _, test := range tests {
		now := time.Now()
		got := ParseDate(now, test)
		if got.IsValid {
			t.Errorf("ParseDate(%q) = %v, want %v", test, got.IsValid, false)
		}
		if got.t != now {
			t.Errorf("ParseDate(%q) = %v, want %v", test, got.t.Format(time.DateTime), now.Format(time.DateTime))
		}
	}
}
