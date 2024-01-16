package cmd

import (
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
	if err := c.loginItems.Load("./login_items.json"); err != nil {
		return err
	}
	if err := c.loginItems.Add(); err != nil {
		return err
	}
	return nil
}
