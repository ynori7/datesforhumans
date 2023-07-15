package datesforhumans

import (
	"testing"
	"time"
)

func TestAt(t *testing.T) {
	base := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)

	tests := []struct {
		in   string
		want time.Time
	}{
		{
			in:   "5:02",
			want: time.Date(2023, time.July, 15, 5, 2, 0, 0, time.UTC),
		},
		{
			in:   "5:02 am",
			want: time.Date(2023, time.July, 15, 5, 2, 0, 0, time.UTC),
		},
		{
			in:   "5:02 pm",
			want: time.Date(2023, time.July, 15, 17, 2, 0, 0, time.UTC),
		},
		{
			in:   "5:02:20",
			want: time.Date(2023, time.July, 15, 5, 2, 20, 0, time.UTC),
		},
		{
			in:   "5:02:20 am",
			want: time.Date(2023, time.July, 15, 5, 2, 20, 0, time.UTC),
		},
		{
			in:   "5:02:20 pm",
			want: time.Date(2023, time.July, 15, 17, 2, 20, 0, time.UTC),
		},
		{
			in:   "5am",
			want: time.Date(2023, time.July, 15, 5, 0, 0, 0, time.UTC),
		},
		{
			in:   "5pm",
			want: time.Date(2023, time.July, 15, 17, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		got := Time{t: base}.At(test.in)
		if got.Time() != test.want {
			t.Errorf("Time.At(%q) = %v, want %v", test.in, got, test.want)
		}
	}

}
