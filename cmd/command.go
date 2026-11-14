package cmd

import (
	"github.com/5ouma/mli/lib"
	"github.com/spf13/cobra"
)

type cmd struct {
	command       *cobra.Command
	loginItems    *lib.LoginItems
	getLoginItems func(*lib.LoginItems) error
	addLoginItems func() (bool, error)
}

func New() *cmd {
	loginItems := new(lib.LoginItems)
	cmd := &cmd{
		command: &cobra.Command{
			Use:          "mli",
			Short:        "📑 Manage macOS Login Items",
			Long:         "📑 Manage macOS Login Items with JSON",
			Version:      lib.Version(),
			SilenceUsage: true,
		},
		loginItems:    loginItems,
		getLoginItems: (*lib.LoginItems).Get,
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
