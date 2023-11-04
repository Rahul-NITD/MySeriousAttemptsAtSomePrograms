package theclock_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	theclock "BadassStuff.com/TheClock/Clock"
)

func TestClockAcceptanceTest(t *testing.T) {
	t.Run("Test Clock At 00:00:00", func(t *testing.T) {
		tm := time.Date(137, time.January, 1, 0, 0, 0, 0, time.UTC)
		b := bytes.Buffer{}
		got := theclock.BuildClock(&b, tm)
		want := theclock.Clock{
			HoursHand:   theclock.HoursHand{X1: 150, Y1: 150, X2: 150, Y2: 100},
			MinutesHand: theclock.MinutesHand{X1: 150, Y1: 150, X2: 150, Y2: 60},
			SecondsHand: theclock.SecondsHand{X1: 150, Y1: 150, X2: 150, Y2: 60},
		}
		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
	t.Run("Test Clock At 6:00:00", func(t *testing.T) {
		tm := time.Date(137, time.January, 1, 6, 0, 0, 0, time.UTC)
		b := bytes.Buffer{}
		got := theclock.BuildClock(&b, tm)
		want := theclock.Clock{
			HoursHand:   theclock.HoursHand{X1: 150, Y1: 150, X2: 150, Y2: 200},
			MinutesHand: theclock.MinutesHand{X1: 150, Y1: 150, X2: 150, Y2: 60},
			SecondsHand: theclock.SecondsHand{X1: 150, Y1: 150, X2: 150, Y2: 60},
		}
		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
	t.Run("Test if Clock returns correct SVG", func(t *testing.T) {
		tm := time.Date(137, time.January, 1, 6, 0, 0, 0, time.UTC)
		b := bytes.Buffer{}
		theclock.BuildClock(&b, tm)
		svg := theclock.SVG{}
		xml.Unmarshal(b.Bytes(), &svg)
		want := theclock.Line(theclock.Point{X1: 150, Y1: 150, X2: 150, Y2: 60})
		if !LineExists(want, svg.Line) {
			t.Errorf("Does not contain required line")
		}
	})
}

func LineExists(want theclock.Line, got []theclock.Line) bool {
	for _, line := range got {
		if line == want {
			return true
		}
	}
	return false
}
