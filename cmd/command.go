package cmd

import (
	"runtime/debug"

	"github.com/5ouma/mli/internal/loginitem"
	"github.com/5ouma/mli/internal/utils"
	"github.com/spf13/cobra"
)

type cmd struct {
	command       *cobra.Command
	loginItems    *loginitem.LoginItems
	getLoginItems func(*loginitem.LoginItems) error
	addLoginItems func() (bool, error)
}

func New() *cmd {
	loginItems := new(loginitem.LoginItems)
	info, ok := debug.ReadBuildInfo()
	cmd := &cmd{
		command: &cobra.Command{
			Use:          "mli",
			Short:        "📑 Manage macOS Login Items",
			Long:         "📑 Manage macOS Login Items with JSON",
			Version:      utils.Version(info, ok),
			SilenceUsage: true,
		},
		loginItems:    loginItems,
		getLoginItems: (*loginitem.LoginItems).Get,
		addLoginItems: loginItems.Add,
	}
	cmd.command.CompletionOptions.HiddenDefaultCmd = true
	cmd.command.SetVersionTemplate("📑 {{.Use}} {{.Version}}\n")
	cmd.command.SetErrPrefix(" 🚨")
	cmd.command.AddCommand(
		cmd.newCheckCmd(),
		cmd.newLoadCmd(),
		cmd.newSaveCmd(),
	)
	return cmd
}

func (cmd *cmd) Execute() error {
	return cmd.command.Execute()
}
