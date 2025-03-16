package logs

import (
	"github.com/spf13/cobra"

	generatecmd "github.com/zyy17/o11ybench/pkg/cmd/logs/generate"
	startcmd "github.com/zyy17/o11ybench/pkg/cmd/logs/start"
)

func NewLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs <command>",
		Short: "Run the logs benchmark suites",
	}

	cmd.AddCommand(startcmd.NewStartCmd())
	cmd.AddCommand(generatecmd.NewGenerateCmd())

	return cmd
}
