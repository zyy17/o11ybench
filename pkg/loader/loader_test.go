package loader

import (
	"fmt"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/zyy17/o11ybench/pkg/collector"
	"github.com/zyy17/o11ybench/pkg/generator"
	"github.com/zyy17/o11ybench/pkg/utils"
)

type mockGenerator struct{}

var _ generator.Generator = &mockGenerator{}

func (g *mockGenerator) Generate(options *generator.GeneratorOptions) (*generator.GeneratorOutput, error) {
	return &generator.GeneratorOutput{
		Data: []byte("test"),
	}, nil
}

type mockTargetService struct {
	port     int
	endpoint string
	rate     int
}

func (s *mockTargetService) Start() error {
	// Create a new http server and always return 200 OK for the endpoint.
	mux := http.NewServeMux()
	mux.HandleFunc(s.endpoint, func(w http.ResponseWriter, r *http.Request) {
		if s.rate > 0 {
			processTime := time.Second / time.Duration(s.rate)
			time.Sleep(processTime)
		}

		w.WriteHeader(http.StatusOK)
	})

	// Choose a random port.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}

	return server.ListenAndServe()
}

func TestLoader(t *testing.T) {
	cfg := &Config{
		Rate:     100,
		Workers:  10,
		Duration: 2 * time.Second,
		Logs: &LogsGeneratorConfig{
			RecordsPerRequest: 10,
		},
		HTTP: HTTPConfig{
			Host:        "localhost",
			Port:        int(utils.RandomNumber(20000, 40000)),
			URI:         "/api/load",
			Method:      "POST",
			Compression: "gzip",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	collector := collector.New()
	loader, err := New(cfg, &mockGenerator{}, collector)
	if err != nil {
		t.Fatalf("failed to create loader: %v", err)
	}

	// Start the mock target service.
	mockTargetService := &mockTargetService{
		port:     cfg.HTTP.Port,
		endpoint: cfg.HTTP.URI,
		rate:     cfg.Rate,
	}
	go mockTargetService.Start()

	// Start the loader.
	loader.Start()

	delta := 1.0
	if math.Abs(collector.Rate()-float64(cfg.Rate)) > delta {
		t.Fatalf("actual rate: '%f', expected rate: '%d', delta: '%f'", collector.Rate(), cfg.Rate, delta)
	}

	if math.Abs(float64(collector.Duration().Seconds())-cfg.Duration.Seconds()) > delta {
		t.Fatalf("actual duration: '%s', expected duration: '%s', delta: '%f'", collector.Duration(), cfg.Duration, delta)
	}
}
