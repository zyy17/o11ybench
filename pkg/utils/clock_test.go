package utils

import (
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	testDuration := 3 * time.Second
	start, _ := time.Parse(time.RFC3339, "2025-03-10T00:00:00Z")
	end := start.Add(testDuration)

	clock, err := NewClock(start, end)
	if err != nil {
		t.Fatal(err)
	}

	var (
		lastNow  time.Time
		interval = 100 * time.Millisecond
	)

	for i := 0; i < 10; i++ {
		now := clock.Now()
		checkRange(t, start, end, now)

		if !lastNow.IsZero() {
			checkInterval(t, interval, now.Sub(lastNow), 10*time.Millisecond)
		}

		lastNow = now
		time.Sleep(interval)
	}

	clock.Forward(1 * time.Second)
	now := clock.Now()
	checkRange(t, start, end, now)
	checkInterval(t, 1*time.Second, now.Sub(lastNow), 110*time.Millisecond)

	clock.Forward(3 * time.Second)
	// Reach the end and it will always return the end time.
	for i := 0; i < 3; i++ {
		if !clock.Now().Equal(end) {
			t.Fatalf("now is '%s', expected '%s'", clock.Now(), end)
		}
	}
}

func checkRange(t *testing.T, start, end, now time.Time) {
	if now.Before(start) || now.After(end) {
		t.Fatalf("now is '%s', expected between '%s' and '%s'", now, start, end)
	}
}

func checkInterval(t *testing.T, expected, actual, tolerance time.Duration) {
	if actual < expected-tolerance || actual > expected+tolerance {
		t.Fatalf("actual '%s' is not within tolerance '%s' of expected '%s', delta: '%s'", actual, tolerance, expected, actual-expected)
	}
}
