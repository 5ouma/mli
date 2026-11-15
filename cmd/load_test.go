package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCommand(t *testing.T) {
	const validFile = "- name: App\n  path: /App\n  hidden: false\n"
	tests := []struct {
		name       string
		content    string
		createFile bool
		add        func() (bool, error)
		wantErr    bool
	}{
		{
			name:       "all apps loaded",
			content:    validFile,
			createFile: true,
			add:        func() (bool, error) { return true, nil },
		},
		{
			name:       "some apps skipped",
			content:    validFile,
			createFile: true,
			add:        func() (bool, error) { return false, nil },
		},
		{
			name:       "System Events error",
			content:    validFile,
			createFile: true,
			add:        func() (bool, error) { return false, errors.New("System Events failed") },
			wantErr:    true,
		},
		{name: "missing file", wantErr: true},
		{name: "invalid YAML", content: "name: [", createFile: true, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "login_items.yml")
			if test.createFile {
				if err := os.WriteFile(path, []byte(test.content), 0600); err != nil {
					t.Fatal(err)
				}
			}
			command := New()
			if test.add != nil {
				command.addLoginItems = test.add
			}
			command.command.SetArgs([]string{"load", "--file", path})
			err := command.Execute()
			if (err != nil) != test.wantErr {
				t.Fatalf("Execute() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
