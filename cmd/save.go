package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func (cmd *cmd) newSaveCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "Save Login Items",
		Long:  "💾 Save Login Items to JSON file",
		Args:  cobra.NoArgs,
		RunE:  cmd.execSaveCmd,
	}
	saveCmd.Flags().String("file", "./login_items.json", "Save to this JSON file")
	saveCmd.Flags().BoolP("force", "f", false, "Overwrite existing file")
	return saveCmd
}

func (cmd *cmd) execSaveCmd(command *cobra.Command, args []string) error {
	file, err := command.Flags().GetString("file")
	if err != nil {
		return err
	}
	path, err := filepath.Abs(file)
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
	if err := cmd.loginItems.Save(path, force); err != nil {
		return err
	}
	fmt.Printf("✅ Successfully saved to \"%s\"!\n", path)

	return nil
}
