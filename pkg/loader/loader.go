package loader

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Loader struct {
	rate          int
	workerNum     int
	duration      time.Duration
	stats         *Stats
	clientOptions *ClientOptions
}

type LoaderOptions struct {
	Rate          int
	WorkerNum     int
	Duration      time.Duration
	ClientOptions *ClientOptions
}

type PayloadGenerator interface {
	Generate() ([]byte, error)
}

type Stats struct {
	TotalRequests atomic.Int64
	TotalErrors   atomic.Int64
}

type HTTPRequest struct {
	Method           string
	URL              string
	Headers          map[string]string
	PayloadGenerator PayloadGenerator
	EnableGzip       bool
}

func NewLoader(opts *LoaderOptions) (*Loader, error) {
	return &Loader{
		rate:          opts.Rate,
		workerNum:     opts.WorkerNum,
		duration:      opts.Duration,
		clientOptions: opts.ClientOptions,
		stats: &Stats{
			TotalRequests: atomic.Int64{},
			TotalErrors:   atomic.Int64{},
		},
	}, nil
}

func (l *Loader) Start(req *HTTPRequest) error {
	var (
		// endTime is the time when the loader will stop.
		endTime = time.Now().Add(l.duration)

		// calculate the number of requests per worker.
		requestNum = l.rate / l.workerNum

		wg sync.WaitGroup
	)

	for i := 0; i < l.workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// check if we should stop.
				if time.Now().After(endTime) {
					return
				}

				start := time.Now()
				for i := 0; i < requestNum; i++ {
					_, err := l.doHTTPRequest(req)
					if err != nil {
						l.stats.TotalErrors.Add(1)
					}
					l.stats.TotalRequests.Add(1)

					// check if we should stop.
					if time.Now().After(endTime) {
						return
					}
				}
				elapsed := time.Since(start)

				// sleep if the requests are not enough to reach the rate.
				if elapsed < time.Second {
					time.Sleep(time.Second - elapsed)
				}
			}
		}()
	}

	wg.Wait()

	return nil
}

func (l *Loader) doHTTPRequest(req *HTTPRequest) (int, error) {
	client, err := NewClient(l.clientOptions)
	if err != nil {
		return 0, err
	}

	var buffer bytes.Buffer
	payload, err := req.PayloadGenerator.Generate()
	if err != nil {
		return 0, err
	}

	if req.EnableGzip {
		if _, err := gzip.NewWriter(&buffer).Write(payload); err != nil {
			return 0, err
		}
		// Set the content encoding to gzip.
		req.Headers["Content-Encoding"] = "gzip"
	} else {
		buffer = *bytes.NewBuffer(payload)
	}

	rawReq, err := http.NewRequest(req.Method, req.URL, &buffer)
	if err != nil {
		return 0, err
	}

	for k, v := range req.Headers {
		rawReq.Header.Set(k, v)
	}

	resp, err := client.Do(rawReq)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
