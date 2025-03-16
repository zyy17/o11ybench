package config

import (
	"bytes"
	"fmt"
	"os"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"

	"github.com/zyy17/o11ybench/pkg/generator"
	"github.com/zyy17/o11ybench/pkg/loader"
)

// Config is the top level configuration for the application.
type Config struct {
	// GeneratorConfig is the configuration for the generator.
	GeneratorConfig *generator.Config `yaml:"generator,omitempty"`

	// LoaderConfig is the configuration for the loader.
	LoaderConfig *loader.Config `yaml:"loader,omitempty"`
}

// New creates a new Config from a file.
func New(configFile string) (*Config, error) {
	if configFile == "" {
		return nil, fmt.Errorf("config file path is required")
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if err := setDefaults(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// PrintConfig prints the config in colorized JSON format.
func (c *Config) Print() error {
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(c); err != nil {
		return err
	}

	fmt.Printf("--- config ---\n")
	fmt.Printf("%s", buf.String())
	fmt.Printf("--------------\n")

	return nil
}

func (c *Config) validate() error {
	if c.GeneratorConfig != nil {
		if err := c.GeneratorConfig.Validate(); err != nil {
			return err
		}
	}

	if c.LoaderConfig != nil {
		if err := c.LoaderConfig.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func setDefaults(cfg *Config) error {
	if cfg.GeneratorConfig != nil {
		if err := mergo.Merge(cfg.GeneratorConfig, cfg.GeneratorConfig.Defaults()); err != nil {
			return err
		}
	}

	if cfg.LoaderConfig != nil {
		if err := mergo.Merge(cfg.LoaderConfig, cfg.LoaderConfig.Defaults()); err != nil {
			return err
		}
	}

	return nil
}
