package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/5ouma/mli/internal/loginitem"
	"github.com/spf13/cobra"
)

type commandErrorWriter struct {
	writesBeforeError int
}

func (writer *commandErrorWriter) Write(data []byte) (int, error) {
	if writer.writesBeforeError == 0 {
		return 0, errors.New("write failed")
	}
	writer.writesBeforeError--
	return len(data), nil
}

func setCommandLipglossOutput(t *testing.T, writesBeforeError int) {
	t.Helper()
	originalOutput := lipgloss.Writer.Forward
	lipgloss.Writer.Forward = &commandErrorWriter{writesBeforeError: writesBeforeError}
	t.Cleanup(func() {
		lipgloss.Writer.Forward = originalOutput
	})
}

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

func TestCommandOutputErrors(t *testing.T) {
	const validFile = `[{"Name":"App","Path":"/App","Hidden":false}]`
	tests := []struct {
		name              string
		args              []string
		setup             func(*testing.T, *cmd)
		writesBeforeError int
	}{
		{name: "check header", args: []string{"check"}},
		{name: "load header", args: []string{"load"}},
		{name: "save header", args: []string{"save"}},
		{
			name: "check final output",
			setup: func(t *testing.T, command *cmd) {
				path := filepath.Join(t.TempDir(), "items.json")
				if err := os.WriteFile(path, []byte(validFile), 0600); err != nil {
					t.Fatal(err)
				}
				command.getLoginItems = func(items *loginitem.LoginItems) error {
					return json.Unmarshal([]byte(validFile), items)
				}
				command.command.SetArgs([]string{"check", "--file", path})
			},
			writesBeforeError: 1,
		},
		{
			name: "load success final output",
			setup: func(t *testing.T, command *cmd) {
				path := filepath.Join(t.TempDir(), "items.json")
				if err := os.WriteFile(path, []byte(validFile), 0600); err != nil {
					t.Fatal(err)
				}
				command.addLoginItems = func() (bool, error) { return true, nil }
				command.command.SetArgs([]string{"load", "--file", path})
			},
			writesBeforeError: 1,
		},
		{
			name: "load partial final output",
			setup: func(t *testing.T, command *cmd) {
				path := filepath.Join(t.TempDir(), "items.json")
				if err := os.WriteFile(path, []byte(validFile), 0600); err != nil {
					t.Fatal(err)
				}
				command.addLoginItems = func() (bool, error) { return false, nil }
				command.command.SetArgs([]string{"load", "--file", path})
			},
			writesBeforeError: 1,
		},
		{
			name: "save final output",
			setup: func(t *testing.T, command *cmd) {
				path := filepath.Join(t.TempDir(), "items.json")
				command.getLoginItems = func(items *loginitem.LoginItems) error {
					return json.Unmarshal([]byte(validFile), items)
				}
				command.command.SetArgs([]string{"save", "--file", path})
			},
			writesBeforeError: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			command := New()
			if test.args != nil {
				command.command.SetArgs(test.args)
			}
			if test.setup != nil {
				test.setup(t, command)
			}
			setCommandLipglossOutput(t, test.writesBeforeError)
			runErr := command.Execute()
			if runErr == nil || !strings.Contains(runErr.Error(), "write failed") {
				t.Fatalf("Execute() error = %v, want output write error", runErr)
			}
		})
	}
}
