package start

import (
	"github.com/spf13/cobra"
)

type StartOptions struct {
}

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the logs benchmark",
	}

	return cmd
}
