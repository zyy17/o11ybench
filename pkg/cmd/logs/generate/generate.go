package generate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/zyy17/o11ybench/pkg/config"
	"github.com/zyy17/o11ybench/pkg/generator"
)

// GenerateOptions is the command options for `generate` subcommand.
type GenerateOptions struct {
	// ConfigFile is the configuration file path.
	ConfigFile string

	// Output is the output file path. If not set, the logs will be printed to the console.
	Output string

	// PrintConfig is the flag to print the config.
	PrintConfig bool
}

func NewGenerateCmd() *cobra.Command {
	opts := &GenerateOptions{}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateLogs(opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.Output, "output", "o", "", "The path to the output file")
	flags.StringVarP(&opts.ConfigFile, "config", "c", "", "The path to the config file")
	flags.BoolVarP(&opts.PrintConfig, "print-config", "p", false, "Print the config")
	return cmd
}

func generateLogs(opts *GenerateOptions) error {
	cfg, err := config.New(opts.ConfigFile)
	if err != nil {
		return err
	}

	if cfg.GeneratorConfig == nil || cfg.GeneratorConfig.Logs == nil {
		return fmt.Errorf("logs generator config is required")
	}

	if opts.PrintConfig {
		if err := cfg.Print(); err != nil {
			return err
		}
	}

	// Setup the generator.
	generator, err := generator.New(cfg.GeneratorConfig)
	if err != nil {
		return err
	}

	logs, err := generator.Generate(nil)
	if err != nil {
		return err
	}

	if opts.Output != "" {
		return os.WriteFile(opts.Output, logs.Data, 0644)
	}

	// If the output is not set, print the logs to the stdout.
	fmt.Print(string(logs.Data))

	return nil
}
