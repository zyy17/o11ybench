package collector

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Collector is used to collect the metrics during the load test.
type Collector struct {
	start    time.Time
	stop     time.Time
	success  atomic.Int64
	failure  atomic.Int64
	duration time.Duration
}

// New creates a new Collector.
func New() *Collector {
	return &Collector{}
}

// Start starts the Collector.
func (c *Collector) Start() {
	c.start = time.Now()
}

// Stop stops the Collector.
func (c *Collector) Stop() {
	c.stop = time.Now()
	c.duration = c.stop.Sub(c.start)
}

// IncSuccessCount increments the success counter.
func (c *Collector) IncSuccessCount(inc int64) {
	c.success.Add(inc)
}

// IncFailureCount increments the failure counter.
func (c *Collector) IncFailureCount(inc int64) {
	c.failure.Add(inc)
}

// Rate returns the actual rate of the load test.
func (c *Collector) Rate() float64 {
	return float64(c.success.Load()) / float64(c.duration.Seconds())
}

// Duration returns the duration of the load test.
func (c *Collector) Duration() time.Duration {
	return c.duration
}

func (c *Collector) Print() {
	fmt.Printf("Success: %d, Failure: %d, Duration: %s, Rate: %f\n", c.success.Load(), c.failure.Load(), c.duration, c.Rate())
}
