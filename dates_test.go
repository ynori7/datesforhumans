package datesforhumans

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	base := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)

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
			want: time.Date(2023, time.August, 15, 17, 4, 0, 0, time.UTC),
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
			in:   "last saturday",
			want: time.Date(2023, time.July, 8, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "yesterday",
			want: time.Date(2023, time.July, 14, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last month",
			want: time.Date(2023, time.June, 15, 17, 4, 0, 0, time.UTC),
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
			want: time.Date(2023, time.January, 15, 17, 4, 0, 0, time.UTC),
		},
		{
			in:   "last july",
			want: time.Date(2022, time.July, 15, 17, 4, 0, 0, time.UTC),
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
	}

	for _, test := range tests {
		got := ParseDate(base, test.in)
		if got.Time() != test.want {
			t.Errorf("ParseDate(%q) = %v, want %v", test.in, got, test.want)
		}
	}
}
