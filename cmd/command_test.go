package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

func TestCommandExecution(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "version", args: []string{"--version"}, want: "📑 mli "},
		{name: "help", args: []string{"--help"}, want: "Manage macOS Login Items"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			command := New()
			var output bytes.Buffer
			command.command.SetOut(&output)
			command.command.SetArgs(test.args)
			if err := command.Execute(); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), test.want) {
				t.Fatalf("output %q does not contain %q", output.String(), test.want)
			}
		})
	}
}

func TestCommandFlagErrors(t *testing.T) {
	command := New()
	for _, test := range []struct {
		name string
		run  func() error
	}{
		{name: "check file flag", run: func() error { return command.execCheckCmd(&cobra.Command{}, nil) }},
		{name: "load file flag", run: func() error { return command.execLoadCmd(&cobra.Command{}, nil) }},
		{
			name: "save force flag",
			run: func() error {
				cobraCommand := &cobra.Command{}
				cobraCommand.Flags().String("file", "items.json", "")
				return command.execSaveCmd(cobraCommand, nil)
			},
		},
		{
			name: "save file flag",
			run: func() error {
				return command.execSaveCmd(&cobra.Command{}, nil)
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(); err == nil {
				t.Fatal("command succeeded without a required flag")
			}
		})
	}
}

func TestCommandsRejectUnexpectedArgs(t *testing.T) {
	for _, name := range []string{"check", "load", "save"} {
		t.Run(name, func(t *testing.T) {
			command := New()
			command.command.SetArgs([]string{name, "unexpected"})
			if err := command.Execute(); err == nil {
				t.Fatal("Execute() succeeded with an unexpected argument")
			}
		})
	}
}

func TestCommandHeaderOutputErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "check", args: []string{"check"}},
		{name: "load", args: []string{"load"}},
		{name: "save", args: []string{"save"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := os.Open(os.DevNull)
			if err != nil {
				t.Fatal(err)
			}
			originalOutput := lipgloss.Writer.Forward
			lipgloss.Writer.Forward = output
			t.Cleanup(func() {
				lipgloss.Writer.Forward = originalOutput
				if err := output.Close(); err != nil {
					t.Error(err)
				}
			})
			command := New()
			command.command.SetArgs(test.args)
			runErr := command.Execute()
			if runErr == nil {
				t.Fatal("command succeeded when stdout was not writable")
			}
		})
	}
}
