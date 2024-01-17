package cmd

import (
	"github.com/spf13/cobra"
)

func (cmd *cmd) newSaveCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "save short desc",
		Long:  "save long desc",
		Args:  cobra.NoArgs,
		RunE:  cmd.execSaveCmd,
	}
	saveCmd.Flags().String("file", "./login_items.json", "save to this JSON file")
	saveCmd.Flags().BoolP("force", "f", false, "overwrite existing file")
	return saveCmd
}

func (cmd *cmd) execSaveCmd(command *cobra.Command, args []string) error {
	file, err := command.Flags().GetString("file")
	if err != nil {
		return err
	}
	force, err := command.Flags().GetBool("force")
	if err != nil {
		return err
	}

	if err := cmd.loginItems.Get(); err != nil {
		return err
	}
	if err := cmd.loginItems.Save(file, force); err != nil {
		return err
	}
	return nil
}
