package cmd

import (
	"fmt"
	"path/filepath"
	"reflect"

	"charm.land/lipgloss/v2"
	"github.com/5ouma/mli/lib"
	"github.com/spf13/cobra"
)

func (cmd *cmd) newCheckCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check Login Items",
		Long:  "🔍 Check the Login Items are up-to-date",
		Args:  cobra.NoArgs,
		RunE:  cmd.execCheckCmd,
	}
	checkCmd.Flags().String("file", "./login_items.json", "Check Login Items from this JSON file")
	return checkCmd
}

func (cmd *cmd) execCheckCmd(command *cobra.Command, args []string) error {
	file, err := command.Flags().GetString("file")
	if err != nil {
		return err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	if _, err := lipgloss.Println(lib.H2.Render("🔍 Check Login Items")); err != nil {
		return err
	}
	if err := cmd.loginItems.Load(path); err != nil {
		return err
	}
	loginItems := new(lib.LoginItems)
	if err := loginItems.Get(); err != nil {
		return err
	}
	if !reflect.DeepEqual(cmd.loginItems, loginItems) {
		fmt.Println()
		return fmt.Errorf("login items are out-of-date")
	}
	if _, err := lipgloss.Println(lib.H1.Render("✅ Login Items are up-to-date!")); err != nil {
		return err
	}

	return nil
}
