package cmd

import (
	"github.com/5ouma/mli/lib"
	"github.com/spf13/cobra"
)

type Cmd struct {
	command    *cobra.Command
	loginItems *lib.LoginItems
}

func New() *Cmd {
	cmd := &Cmd{
		command: &cobra.Command{
			Use:          "mli",
			Short:        "short desc",
			Long:         "long desc",
			SilenceUsage: true,
		},
		loginItems: new(lib.LoginItems),
	}
	cmd.command.AddCommand(
		cmd.newSaveCmd(),
		cmd.newLoadCmd(),
	)
	return cmd
}

func (cmd *Cmd) Execute() error {
	return cmd.command.Execute()
}
