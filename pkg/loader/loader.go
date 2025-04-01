package loader

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/zyy17/o11ybench/pkg/collector"
	"github.com/zyy17/o11ybench/pkg/generator"
	logstypes "github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

type Loader struct {
	cfg       *Config
	generator generator.Generator
	collector *collector.Collector
}

func New(cfg *Config, generator generator.Generator, collector *collector.Collector) (*Loader, error) {
	return &Loader{cfg: cfg, generator: generator, collector: collector}, nil
}

func (l *Loader) Start() error {
	var (
		// infinite is true if the loader is running infinitely.
		infinite bool

		// stopTime is the time when the loader will stop if it's not infinite.
		stopTime time.Time

		// requestsNum is the number of requests for each worker(goroutine).
		requestsNum = l.cfg.Rate / l.cfg.Workers

		// wg is the wait group for the workers.
		wg sync.WaitGroup
	)

	if l.cfg.Duration == 0 {
		infinite = true
	} else {
		stopTime = time.Now().Add(l.cfg.Duration)
	}

	// Start the collector.
	l.collector.Start()

	// Setup the signal handling to ensure we print stats even if terminated.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		fmt.Println("Received interrupt or termination signal, printing stats before exit...")
		l.collector.Stop()
		l.collector.Print()
		os.Exit(0)
	}()

	for i := 0; i < l.cfg.Workers; i++ {
		wg.Add(1)
		w := &worker{id: i, requestsNum: requestsNum, infinite: infinite, stopTime: stopTime}
		go func() {
			defer wg.Done()
			l.workerLoop(w)
		}()
	}

	wg.Wait()

	// Stop the collector.
	l.collector.Stop()

	return nil
}

type worker struct {
	id          int
	requestsNum int
	infinite    bool
	stopTime    time.Time
}

func (l *Loader) workerLoop(w *worker) {
	for {
		// Check if the worker should stop.
		if !w.infinite && time.Now().After(w.stopTime) {
			return
		}

		hc, err := l.httpClient()
		if err != nil {
			panic(err)
		}

		start := time.Now()
		for i := 0; i < w.requestsNum; i++ {
			if err := l.doRequest(hc); err != nil {
				l.collector.IncFailureCount(1)
				fmt.Printf("worker [%d] failed to make request: %v\n", w.id, err)
			} else {
				l.collector.IncSuccessCount(1)
				l.collector.IncRecordsCount(int64(l.cfg.Logs.RecordsPerRequest))
			}

			// Check if the worker should stop.
			if !w.infinite && time.Now().After(w.stopTime) {
				return
			}
		}
		elapsed := time.Since(start)

		// sleep if the requests are not enough to reach the rate.
		if elapsed < time.Second {
			time.Sleep(time.Second - elapsed)
		}
	}
}

func (l *Loader) doRequest(hc *http.Client) error {
	req, err := l.makeHTTPRequest()
	if err != nil {
		return err
	}

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request '%s' failed with status code '%d' and body '%s'", req.URL, resp.StatusCode, string(body))
	}

	return nil
}

func (l *Loader) makeHTTPRequest() (*http.Request, error) {
	// Generates the payload for the request.
	output, err := l.generator.Generate(l.generatorOptions())
	if err != nil {
		return nil, err
	}

	requestURL, err := l.constructURL()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if l.cfg.HTTP.Compression == "gzip" {
		// Compress the payload using gzip.
		writer := gzip.NewWriter(&buf)
		if _, err := writer.Write(output.Data); err != nil {
			return nil, err
		}
		writer.Close()
	} else {
		buf.Write(output.Data)
	}

	req, err := http.NewRequest(strings.ToUpper(l.cfg.HTTP.Method), requestURL, &buf)
	if err != nil {
		return nil, err
	}

	for k, v := range l.cfg.HTTP.Headers {
		req.Header.Set(k, v)
	}

	if l.cfg.HTTP.Compression == "gzip" {
		req.Header.Set("Content-Encoding", "gzip")
	}

	return req, nil
}

func (l *Loader) httpClient() (*http.Client, error) {
	client := http.DefaultClient

	client.Transport = &http.Transport{
		ResponseHeaderTimeout: l.cfg.HTTP.ResponseHeaderTimeout,
	}

	return client, nil
}

func (l *Loader) constructURL() (string, error) {
	return fmt.Sprintf("http://%s:%d%s", l.cfg.HTTP.Host, l.cfg.HTTP.Port, l.cfg.HTTP.URI), nil
}

func (l *Loader) generatorOptions() *generator.GeneratorOptions {
	return &generator.GeneratorOptions{
		Logs: &logstypes.GeneratorOptions{
			LogsCount: l.cfg.Logs.RecordsPerRequest,
			Timestamp: time.Now(),
		},
	}
}
