package cmd

import (
	"github.com/spf13/cobra"
)

func (cmd *Cmd) newSaveCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "save short desc",
		Long:  "save long desc",
		Args:  cobra.NoArgs,
		RunE:  cmd.execSaveCmd,
	}
	return saveCmd
}

func (cmd *Cmd) execSaveCmd(command *cobra.Command, args []string) error {
	if err := cmd.loginItems.Get(); err != nil {
		return err
	}
	if err := cmd.loginItems.Save("./login_items.json", false); err != nil {
		return err
	}
	return nil
}
