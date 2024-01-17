package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (cmd *cmd) newLoadCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "short load desc",
		Long:  "long load desc",
		Args:  cobra.NoArgs,
		RunE:  cmd.execLoadCmd,
	}
	loadCmd.Flags().String("file", "./login_items.json", "Load from this JSON file")
	return loadCmd
}

func (cmd *cmd) execLoadCmd(command *cobra.Command, args []string) error {
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
	fmt.Printf("✅ Successfully loaded from \"%s\"!\n", file)

	return nil
}
