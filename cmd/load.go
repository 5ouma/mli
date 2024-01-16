package cmd

import (
	"github.com/spf13/cobra"
)

func (cmd *Cmd) newLoadCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "short load desc",
		Long:  "long load desc",
		Args:  cobra.NoArgs,
		RunE:  cmd.execLoadCmd,
	}
	loadCmd.Flags().String("file", "./login_items.json", "load from this JSON file")
	return loadCmd
}

func (cmd *Cmd) execLoadCmd(command *cobra.Command, args []string) error {
	file, err := command.Flags().GetString("file")
	if err != nil {
		return err
	}

	if err := cmd.loginItems.Load(file); err != nil {
		return err
	}
	if err := cmd.loginItems.Add(); err != nil {
		return err
	}
	return nil
}
