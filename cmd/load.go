package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newLoadCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "short load desc",
		Long:  "long load desc",
		Args:  cobra.NoArgs,
		RunE:  c.execLoadCmd,
	}
	return loadCmd
}

func (c *Cmd) execLoadCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("a message from load")
	return nil
}
