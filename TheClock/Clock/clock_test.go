package theclock_test

import (
	"math"
	"testing"
	"time"

	theclock "BadassStuff.com/TheClock/Clock"
)

func SimpleTime(hh, mm, ss int) time.Time {
	return time.Date(137, time.February, 1, hh, mm, ss, 0, time.UTC)
}

func TestSecondsHand(t *testing.T) {
	t.Run("Test Convert to radians", func(t *testing.T) {
		cases := []struct {
			currVal  float64
			divsions int
			want     float64
		}{
			{0, 60, 0},
			{30, 60, math.Pi},
			{45, 60, 3 * math.Pi / 2},
			{0, 12, 0},
			{6, 12, math.Pi},
			{9, 12, 3 * math.Pi / 2},
		}
		for _, test := range cases {
			got := theclock.InRadians(test.currVal, test.divsions)
			if got != test.want {
				t.Errorf("got %f want %f", got, test.want)
			}
		}
	})

	t.Run("Test Position of hand in Unit Circle", func(t *testing.T) {

		cases := []struct {
			tm   time.Time
			want theclock.Clock
		}{
			{SimpleTime(0, 0, 0), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(0, 0, 30), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 0, 1}}},
			{SimpleTime(0, 0, 15), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 1, 0}}},
			{SimpleTime(0, 0, 45), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, -1, 0}}},
			{SimpleTime(0, 30, 0), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 0, 1}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(0, 15, 0), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, 1, 0}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(0, 45, 0), theclock.Clock{theclock.HoursHand{0, 0, 0, -1}, theclock.MinutesHand{0, 0, -1, 0}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(6, 0, 0), theclock.Clock{theclock.HoursHand{0, 0, 0, 1}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(3, 0, 0), theclock.Clock{theclock.HoursHand{0, 0, 1, 0}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 0, -1}}},
			{SimpleTime(9, 0, 0), theclock.Clock{theclock.HoursHand{0, 0, -1, 0}, theclock.MinutesHand{0, 0, 0, -1}, theclock.SecondsHand{0, 0, 0, -1}}},
		}
		for _, test := range cases {
			got := theclock.GetUnitPoint(test.tm)
			AssertClock(t, got, test.want, test.tm)
		}
	})
	t.Run("Test Hands Position", func(t *testing.T) {
		cases := []struct {
			tm   time.Time
			want theclock.Clock
		}{
			{SimpleTime(0, 0, 0), theclock.Clock{theclock.HoursHand{150, 150, 150, 100}, theclock.MinutesHand{150, 150, 150, 60}, theclock.SecondsHand{150, 150, 150, 60}}},
			{SimpleTime(0, 0, 30), theclock.Clock{theclock.HoursHand{150, 150, 150, 100}, theclock.MinutesHand{150, 150, 150, 60}, theclock.SecondsHand{150, 150, 150, 240}}},
			{SimpleTime(0, 0, 15), theclock.Clock{theclock.HoursHand{150, 150, 150, 100}, theclock.MinutesHand{150, 150, 150, 60}, theclock.SecondsHand{150, 150, 240, 150}}},
			{SimpleTime(0, 0, 45), theclock.Clock{theclock.HoursHand{150, 150, 150, 100}, theclock.MinutesHand{150, 150, 150, 60}, theclock.SecondsHand{150, 150, 60, 150}}},
		}
		for _, test := range cases {
			got := theclock.GetPoint(test.tm)
			AssertClock(t, got, test.want, test.tm)
		}
	})
}

func AssertClock(t testing.TB, got, want theclock.Clock, tm time.Time) {
	t.Helper()
	if !almostEqualPoints(theclock.Point(got.HoursHand), theclock.Point(want.HoursHand), 5) {
		t.Errorf("%+v : HourHand : got %f want %f", tm, got.HoursHand, want.HoursHand)
	}
	if !almostEqualPoints(theclock.Point(got.MinutesHand), theclock.Point(want.MinutesHand), 8) {
		t.Errorf("%+v : MinutesHand : got %f want %f", tm, got.MinutesHand, want.MinutesHand)
	}
	if !almostEqualPoints(theclock.Point(got.SecondsHand), theclock.Point(want.SecondsHand), 1) {
		t.Errorf("%+v : SecondsHand : got %f want %f", tm, got.SecondsHand, want.SecondsHand)
	}
}

func almostEqualPoints(got, want theclock.Point, thr float64) bool {
	return almostEqualFloats(got.X1, want.X1, thr) && almostEqualFloats(got.X2, want.X2, thr) && almostEqualFloats(got.Y1, want.Y1, thr) && almostEqualFloats(got.Y2, want.Y2, thr)
}

func almostEqualFloats(got, want float64, thr float64) bool {
	return math.Abs(got-want) < thr
}
