package start

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zyy17/o11ybench/pkg/collector"
	"github.com/zyy17/o11ybench/pkg/config"
	"github.com/zyy17/o11ybench/pkg/generator"
	"github.com/zyy17/o11ybench/pkg/loader"
)

// StartOptions is the command options for `start` subcommand.
type StartOptions struct {
	// ConfigFile is the configuration file path.
	ConfigFile string
}

func NewStartCmd() *cobra.Command {
	opts := &StartOptions{}

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the logs benchmark",
		RunE: func(cmd *cobra.Command, args []string) error {
			return start(opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.ConfigFile, "config", "c", "", "The path to the config file")
	return cmd
}

func start(opts *StartOptions) error {
	cfg, err := config.New(opts.ConfigFile)
	if err != nil {
		return err
	}

	if cfg.GeneratorConfig == nil {
		return fmt.Errorf("generator config is required")
	}

	if cfg.GeneratorConfig.Logs == nil {
		return fmt.Errorf("logs generator config is required")
	}

	if cfg.LoaderConfig == nil {
		return fmt.Errorf("loader config is required")
	}

	if err := cfg.Print(); err != nil {
		return err
	}

	// Setup the generator.
	generator, err := generator.New(cfg.GeneratorConfig)
	if err != nil {
		return err
	}

	// Setup the collector.
	collector := collector.New()

	// Setup the loader.
	loader, err := loader.New(cfg.LoaderConfig, generator, collector)
	if err != nil {
		return err
	}

	// Start the loader.
	if err := loader.Start(); err != nil {
		return err
	}

	// Print the stats.
	collector.Print()

	return nil
}
