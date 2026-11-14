package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestMainEntryPoint(t *testing.T) {
	if mode := os.Getenv("MLI_MAIN_TEST"); mode != "" {
		if mode == "version" {
			os.Args = []string{"mli", "--version"}
		} else {
			os.Args = []string{"mli", "unknown-command"}
		}
		main()
		return
	}

	tests := []struct {
		name    string
		mode    string
		wantErr bool
	}{
		{name: "successful command", mode: "version"},
		{name: "command error exits nonzero", mode: "error", wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			command := exec.Command(os.Args[0], "-test.run=^TestMainEntryPoint$")
			command.Env = append(os.Environ(), "MLI_MAIN_TEST="+test.mode)
			err := command.Run()
			if (err != nil) != test.wantErr {
				t.Fatalf("child process error = %v, wantErr %v", err, test.wantErr)
			}
			if test.wantErr {
				var exitError *exec.ExitError
				if !errors.As(err, &exitError) || exitError.ExitCode() != 1 {
					t.Fatalf("child process error = %v, want exit code 1", err)
				}
			}
		})
	}
}
