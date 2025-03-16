package collector

import (
	"math"
	"sync"
	"testing"
	"time"
)

func TestCollector(t *testing.T) {
	collector := New()
	testWorkerNum := 10
	expectedRate := 100
	duration := 2 * time.Second

	var wg sync.WaitGroup

	collector.Start()
	for i := 0; i < testWorkerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			collector.IncSuccessCount(int64(expectedRate / testWorkerNum))
			elapsed := time.Since(start)
			time.Sleep(time.Second - elapsed)
		}()
	}
	wg.Wait()
	collector.Stop()
	collector.Print()

	actualRate := collector.Rate()
	delta := 1.0
	if math.Abs(actualRate-float64(expectedRate)) > delta {
		t.Fatalf("actual rate: '%f', expected rate: '%d', delta: '%f'", actualRate, expectedRate, delta)
	}

	if math.Abs(float64(collector.Duration().Seconds())-duration.Seconds()) > delta {
		t.Fatalf("actual duration: '%s', expected duration: '%s', delta: '%f'", collector.Duration(), duration, delta)
	}
}
