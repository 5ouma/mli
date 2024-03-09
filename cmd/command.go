package cmd

import (
	"github.com/5ouma/mli/lib"
	"github.com/spf13/cobra"
)

type cmd struct {
	command    *cobra.Command
	loginItems *lib.LoginItems
}

func New() *cmd {
	cmd := &cmd{
		command: &cobra.Command{
			Use:          "mli",
			Short:        "📑 Manage macOS Login Items",
			Long:         "📑 Manage macOS Login Items with JSON",
			Version:      lib.Version(),
			SilenceUsage: true,
		},
		loginItems: new(lib.LoginItems),
	}
	cmd.command.CompletionOptions.HiddenDefaultCmd = true
	cmd.command.SetVersionTemplate("📑 {{.Use}} {{.Version}}\n")
	cmd.command.SetErrPrefix("🚨")
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
