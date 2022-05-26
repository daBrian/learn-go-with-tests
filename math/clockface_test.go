package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{150, 150 + 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

//func TestSecondsInRadians(t *testing.T) {
//	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
//	want := math.Pi
//	got := secondsInRadians(thirtySeconds)
//
//	if got != want {
//		t.Errorf("got %v want %v", got, want)
//	}
//}

func TestSecondsInRadians(t *testing.T) {
	tests := []struct {
		time time.Time
		want float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, tt := range tests {
		t.Run(simpleName(tt.time), func(t *testing.T) {
			if got := secondsInRadians(tt.time); got != tt.want {
				t.Errorf("secondsInRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}
func simpleTime(h, m, s int) time.Time {
	return time.Date(312, time.October, 28, h, m, s, 0, time.UTC)
}

func simpleName(t time.Time) string {
	return t.Format("15:04:05")
}
