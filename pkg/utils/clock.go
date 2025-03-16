package utils

import (
	"fmt"
	"time"
)

// Clock is a utility for generating time values within a given time range.
type Clock struct {
	start time.Time
	end   time.Time

	// record the last real time Now() was called.
	lastCall time.Time

	// record the last time Now() was returned.
	lastNow time.Time
}

// NewClock creates a new clock with the given start and end times.
// If end is zero, the clock will not have an end time.
func NewClock(start, end time.Time) (*Clock, error) {
	if start.IsZero() {
		return nil, fmt.Errorf("start time is zero")
	}

	if !end.IsZero() && start.After(end) {
		return nil, fmt.Errorf("start time %s is after end time %s", start, end)
	}

	return &Clock{
		start:   start,
		end:     end,
		lastNow: start,
	}, nil
}

// Now returns the current time based on the clock's start time and end time.
func (c *Clock) Now() time.Time {
	// First call.
	if c.lastCall.IsZero() {
		c.lastCall = time.Now()
		return c.lastNow
	}

	delta := time.Since(c.lastCall)
	now := c.lastNow.Add(delta)

	// If the end time is set, and the delta is greater than the end time minus the start time, always return the end time.
	if !c.end.IsZero() && now.After(c.end) {
		return c.end
	}

	c.lastCall = time.Now()
	c.lastNow = now

	return now
}

// Reset resets the clock to the start time.
func (c *Clock) Reset() {
	c.lastCall = time.Time{}
	c.lastNow = c.start
}

// Forward moves the clock forward by the given duration.
func (c *Clock) Forward(d time.Duration) {
	if !c.lastCall.IsZero() {
		now := c.lastNow.Add(d)
		if !c.end.IsZero() && now.After(c.end) {
			c.lastNow = c.end
			return
		}

		c.lastNow = now
	}
}
