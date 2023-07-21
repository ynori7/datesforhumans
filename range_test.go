package datesforhumans

import (
	"testing"
	"time"
)

func TestParseRange(t *testing.T) {
	base := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)

	actual := ParseRange(base, "next monday at 5pm", "next tuesday at 6pm")

	expectedStart := time.Date(2023, time.July, 17, 17, 0, 0, 0, time.UTC)
	if actual.Start.t != expectedStart {
		t.Errorf("ParseRange() = %v, want %v", actual.Start.t.Format(time.DateTime), expectedStart.Format(time.DateTime))
	}

	expectedEnd := time.Date(2023, time.July, 18, 18, 0, 0, 0, time.UTC)
	if actual.End.t != expectedEnd {
		t.Errorf("ParseRange() = %v, want %v", actual.End.t.Format(time.DateTime), expectedEnd.Format(time.DateTime))
	}

	repeated := actual.Repeat(time.Date(2023, time.July, 25, 0, 0, 0, 0, time.UTC))
	if len(repeated) != 2 {
		t.Errorf("ParseRange() = %v, want %v", len(repeated), 2)
	}

	if repeated[0].Start.t != expectedStart {
		t.Errorf("Repeat() = %v, want %v", repeated[0].Start.t.Format(time.DateTime), expectedStart.Format(time.DateTime))
	}
	if repeated[0].End.t != expectedEnd {
		t.Errorf("Repeat() = %v, want %v", repeated[0].End.t.Format(time.DateTime), expectedEnd.Format(time.DateTime))
	}
	expectedStart = time.Date(2023, time.July, 24, 17, 0, 0, 0, time.UTC)
	expectedEnd = time.Date(2023, time.July, 25, 18, 0, 0, 0, time.UTC)
	if repeated[1].Start.t != expectedStart {
		t.Errorf("Repeat() = %v, want %v", repeated[0].Start.t.Format(time.DateTime), expectedStart.Format(time.DateTime))
	}
	if repeated[1].End.t != expectedEnd {
		t.Errorf("Repeat() = %v, want %v", repeated[0].End.t.Format(time.DateTime), expectedEnd.Format(time.DateTime))
	}
}
