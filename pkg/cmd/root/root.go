package root

import (
	"github.com/spf13/cobra"

	"github.com/zyy17/o11ybench/pkg/cmd/logs"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "o11ybench <command> <subcommand>",
		Short: "o11ybench is a tool for benchmarking observability ecosystems",
	}

	cmd.AddCommand(logs.NewLogsCmd())

	return cmd
}
