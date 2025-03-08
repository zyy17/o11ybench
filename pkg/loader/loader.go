package loader

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/zyy17/o11ybench/pkg/generator"
)

type Loader struct {
	opts *LoaderOptions
}

type LoaderOptions struct {
	Rate         int
	WorkerNum    int
	Duration     time.Duration
	Endpoint     string
	DB           string
	Table        string
	PipelineName string
	EnableGzip   bool
	IsInfinite   bool
	Generator    generator.Generator
}

func NewLoader(opts *LoaderOptions) (*Loader, error) {
	return &Loader{
		opts: opts,
	}, nil
}

func (l *Loader) Start() error {
	var (
		// endTime is the time when the loader will stop.
		endTime = time.Now().Add(l.opts.Duration)

		// FIXME: make it configurable.
		requestNum = 1

		wg sync.WaitGroup
	)

	for i := 0; i < l.opts.WorkerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// check if we should stop.
				if !l.opts.IsInfinite && time.Now().After(endTime) {
					return
				}

				start := time.Now()
				for i := 0; i < requestNum; i++ {
					if err := l.doHTTPRequest(); err != nil {
						fmt.Println(err)
					}

					// check if we should stop.
					if !l.opts.IsInfinite && time.Now().After(endTime) {
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

func (l *Loader) doHTTPRequest() error {
	// FIXME: make it configurable.
	url := fmt.Sprintf("%s/v1/events/logs?db=%s&table=%s&pipeline_name=%s", l.opts.Endpoint, l.opts.DB, l.opts.Table, l.opts.PipelineName)

	client, err := NewClient(&ClientOptions{
		TimeoutInMillisecond: 10000,
	})
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	var headers = make(map[string]string)
	payload, err := l.opts.Generator.Generate()
	if err != nil {
		return err
	}
	headers["Content-Type"] = "application/json"

	if l.opts.EnableGzip {
		writer := gzip.NewWriter(&buffer)
		if _, err := writer.Write(payload); err != nil {
			return err
		}
		writer.Close()

		// Set the content encoding to gzip.
		headers["Content-Encoding"] = "gzip"
	} else {
		buffer = *bytes.NewBuffer(payload)
	}

	rawReq, err := http.NewRequest(http.MethodPost, url, &buffer)
	if err != nil {
		return err
	}

	for k, v := range headers {
		rawReq.Header.Set(k, v)
	}

	resp, err := client.Do(rawReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
