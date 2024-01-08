package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mli",
		Short:        "short desc",
		Long:         "long desc",
		SilenceUsage: true,
	}
	cmd.AddCommand(
		newSaveCmd(),
		newLoadCmd(),
	)
	return cmd
}
