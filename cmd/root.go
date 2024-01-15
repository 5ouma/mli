package cmd

import (
	"github.com/5ouma/mli/utils/api"
	"github.com/spf13/cobra"
)

type Cmd struct {
	root       *cobra.Command
	loginItems *api.LoginItems
}

func New() *Cmd {
	cmd := &Cmd{
		root: &cobra.Command{
			Use:          "mli",
			Short:        "short desc",
			Long:         "long desc",
			SilenceUsage: true,
		},
		loginItems: new(api.LoginItems),
	}
	cmd.root.AddCommand(
		cmd.newSaveCmd(),
		cmd.newLoadCmd(),
	)
	return cmd
}

func (c *Cmd) Execute() error {
	return c.root.Execute()
}
