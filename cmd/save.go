package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newSaveCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "save short desc",
		Long:  "save long desc",
		Args:  cobra.NoArgs,
		RunE:  c.execSaveCmd,
	}
	return saveCmd
}

func (c *Cmd) execSaveCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("a message from save")
	return nil
}
