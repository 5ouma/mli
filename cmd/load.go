package cmd

import (
	"fmt"
	"path/filepath"

	"charm.land/lipgloss/v2"
	"github.com/5ouma/mli/lib"
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

	if _, err := lipgloss.Println(lib.H2.Render("📂 Load Login Items")); err != nil {
		return err
	}
	if err := cmd.loginItems.Load(path); err != nil {
		return err
	}
	if loadedCorrectly, err := cmd.loginItems.Add(); err != nil {
		fmt.Println()
		return err
	} else if loadedCorrectly {
		if _, err := lipgloss.Println(lib.H1.Render("✅ Successfully loaded")); err != nil {
			return err
		}
	} else {
		if _, err := lipgloss.Println(lib.H1.Render("⚠️ Some apps weren't loaded correctly")); err != nil {
			return err
		}
	}

	return nil
}
