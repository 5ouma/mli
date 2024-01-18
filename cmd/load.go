package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func (cmd *cmd) newLoadCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "Load Login Items",
		Long:  "📂 Load Login Items from JSON file",
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
	path, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	if err := cmd.loginItems.Load(path); err != nil {
		return err
	}
	if err := cmd.loginItems.Add(); err != nil {
		return err
	}
	fmt.Printf("✅ Successfully loaded from \"%s\"!\n", path)

	return nil
}
