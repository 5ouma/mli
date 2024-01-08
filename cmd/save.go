package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSaveCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "save short desc",
		Long:  "save long desc",
		Args:  cobra.NoArgs,
		RunE:  execSaveCmd,
	}
	return saveCmd
}

func execSaveCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("a message from save")
	return nil
}
