package start

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/zyy17/o11ybench/pkg/generator"
	"github.com/zyy17/o11ybench/pkg/loader"
)

type StartOptions struct {
	Config    string
	BatchSize int
	Rate      int
	Endpoint  string
	Database  string
	Duration  string
	Pipeline  string
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
	flags.StringVarP(&opts.Config, "config", "c", "", "Config file")
	flags.IntVarP(&opts.BatchSize, "batch-size", "b", 2, "Batch size")
	flags.StringVarP(&opts.Endpoint, "endpoint", "e", "http://localhost:4000", "Endpoint")
	flags.StringVarP(&opts.Database, "database", "d", "public", "Database")
	flags.StringVarP(&opts.Duration, "duration", "D", "infinity", "Duration")
	flags.StringVarP(&opts.Pipeline, "pipeline", "p", "greptime_identity", "Pipeline name")
	flags.IntVarP(&opts.Rate, "rate", "r", 10, "Rate")
	return cmd
}

func start(opts *StartOptions) error {
	loaderOpts := &loader.LoaderOptions{
		Rate:         opts.Rate,
		Endpoint:     opts.Endpoint,
		DB:           opts.Database,
		PipelineName: opts.Pipeline,
		WorkerNum:    opts.Rate / opts.BatchSize,
		EnableGzip:   true,
		Table:        "logs",
	}

	generator, err := generator.NewGeneratorFromFile(opts.Config)
	if err != nil {
		return err
	}
	loaderOpts.Generator = generator

	if opts.Duration == "" || opts.Duration == "infinity" {
		loaderOpts.IsInfinite = true
	} else {
		parsedDuration, err := time.ParseDuration(opts.Duration)
		if err != nil {
			return err
		}
		loaderOpts.Duration = parsedDuration
	}

	loader, err := loader.NewLoader(loaderOpts)
	if err != nil {
		return err
	}

	if err := loader.Start(); err != nil {
		return err
	}

	return nil
}
