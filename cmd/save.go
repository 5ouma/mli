package cmd

import (
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
	if err := c.loginItems.Get(); err != nil {
		return err
	}
	if err := c.loginItems.Save("./login_item.json", false); err != nil {
		return err
	}
	return nil
}
