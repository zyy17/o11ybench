package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/zyy17/o11ybench/pkg/generator"
)

type GenerateOptions struct {
	Config string
	Output string
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
	flags.StringVarP(&opts.Output, "output", "o", "", "Output file")
	flags.StringVarP(&opts.Config, "config", "c", "", "Config file")

	return cmd
}

func generateLogs(opts *GenerateOptions) error {
	generator, err := generator.NewGeneratorFromFile(opts.Config)
	if err != nil {
		return fmt.Errorf("create generator: %v", err)
	}

	logs, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("generate logs: %v", err)
	}

	// Add a newline to the end of the logs.
	output := strings.Join(logs, "\n") + "\n"

	if opts.Output != "" {
		return os.WriteFile(opts.Output, []byte(output), 0644)
	}

	// If the output is not set, print the logs to the stdout.
	fmt.Printf("%s", output)

	return nil
}
