package generator

import (
	"fmt"

	"github.com/zyy17/o11ybench/pkg/generator/logs"
	logstypes "github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

// Generator is the interface for the data generator.
type Generator interface {
	// Generates the data for the stress test by the given options.
	Generate(opts *GeneratorOptions) (*GeneratorOutput, error)
}

// GeneratorOptions is the options for configuring the data generation.
type GeneratorOptions struct {
	Logs *logstypes.GeneratorOptions
}

// GeneratorOutput is the output of the data generation.
type GeneratorOutput struct {
	// Data is the generated data.
	Data []byte
}

// GeneratorType is the type of the generator.
type GeneratorType string

const (
	// GeneratorTypeLogs is the type of the generator for the logs.
	GeneratorTypeLogs GeneratorType = "logs"

	// GeneratorTypeTraces is the type of the generator for the traces.
	GeneratorTypeTraces GeneratorType = "traces"

	// GeneratorTypeMetrics is the type of the generator for the metrics.
	GeneratorTypeMetrics GeneratorType = "metrics"
)

type generator struct {
	typ  GeneratorType
	logs *logs.LogsGenerator
}

var _ Generator = &generator{}

// New creates a new generator based on the given configuration.
func New(cfg *Config) (Generator, error) {
	if cfg.Logs != nil {
		logsGenerator, err := logs.NewLogsGenerator(cfg.Logs, cfg.Time)
		if err != nil {
			return nil, err
		}

		return &generator{typ: GeneratorTypeLogs, logs: logsGenerator}, nil
	}

	return nil, fmt.Errorf("invalid generator config")
}

func (g *generator) Generate(opts *GeneratorOptions) (*GeneratorOutput, error) {
	if g.typ == GeneratorTypeLogs {
		var options *logstypes.GeneratorOptions
		if opts != nil && opts.Logs != nil {
			options = opts.Logs
		}

		data, err := g.logs.Generate(options)
		if err != nil {
			return nil, err
		}

		return &GeneratorOutput{Data: data}, nil
	}

	return nil, fmt.Errorf("no generator found")
}
