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
		{
			in:   "12am",
			want: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			in:   "12pm",
			want: time.Date(2023, time.July, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		got := Time{t: base}.At(test.in)
		if got.Time() != test.want {
			t.Errorf("Time.At(%q) = %v, want %v", test.in, got.Time().Format(time.DateTime), test.want)
		}
	}
}

func TestRepeatTime(t *testing.T) {
	base := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)
	tt := ParseDate(base, "next monday").At("10pm")

	repeated := tt.Repeat(time.Date(2023, time.July, 25, 0, 0, 0, 0, time.UTC))
	if len(repeated) != 2 {
		t.Errorf("ParseRange() = %v, want %v", len(repeated), 2)
	}

	expectedDate := time.Date(2023, time.July, 17, 22, 0, 0, 0, time.UTC)
	if repeated[0].t != expectedDate {
		t.Errorf("Repeat() = %v, want %v", repeated[0].t.Format(time.DateTime), expectedDate.Format(time.DateTime))
	}
	expectedDate = time.Date(2023, time.July, 24, 22, 0, 0, 0, time.UTC)
	if repeated[1].t != expectedDate {
		t.Errorf("Repeat() = %v, want %v", repeated[1].t.Format(time.DateTime), expectedDate.Format(time.DateTime))
	}
}
