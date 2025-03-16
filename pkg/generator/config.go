package generator

import (
	"fmt"

	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
	"github.com/zyy17/o11ybench/pkg/generator/metrics"
	"github.com/zyy17/o11ybench/pkg/generator/traces"
)

// Config is the configuration for the generator.
type Config struct {
	// Logs is the configuration for the logs generator.
	Logs *types.LogsGeneratorConfig `yaml:"logs,omitempty"`

	// Traces is the configuration for the traces generator.
	Traces *traces.TracesGeneratorConfig `yaml:"traces,omitempty"`

	// Metrics is the configuration for the metrics generator.
	Metrics *metrics.MetricsGeneratorConfig `yaml:"metrics,omitempty"`

	// Time is the configuration for the time of the data to be generated.
	Time *common.TimeConfig `yaml:"time,omitempty"`
}

// Validate validates the configuration for the generator.
func (c *Config) Validate() error {
	var typ GeneratorType
	if c.Logs != nil {
		typ = GeneratorTypeLogs
		if err := c.Logs.Validate(); err != nil {
			return err
		}
	}

	if c.Traces != nil {
		if typ != "" {
			return fmt.Errorf("only one generator can be set")
		}
		typ = GeneratorTypeTraces
	}

	if c.Metrics != nil {
		if typ != "" {
			return fmt.Errorf("only one generator can be set")
		}
		typ = GeneratorTypeMetrics
	}

	if typ == "" {
		return fmt.Errorf("no generator is configured")
	}

	if c.Time != nil {
		if err := c.Time.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Defaults returns the default configuration for the generator.
func (c Config) Defaults() *Config {
	if c.Logs != nil {
		return &Config{
			Logs: types.LogsGeneratorConfig{}.Defaults(),
			Time: common.TimeConfig{}.Defaults(),
		}
	}

	return &c
}
