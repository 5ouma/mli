package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/5ouma/mli/internal/loginitem"
)

func TestCheckCommand(t *testing.T) {
	const validFile = `[{"Name":"App","Path":"/App","Hidden":false}]`
	tests := []struct {
		name       string
		content    string
		createFile bool
		get        func(*loginitem.LoginItems) error
		wantErr    bool
	}{
		{
			name:       "up to date",
			content:    validFile,
			createFile: true,
			get: func(items *loginitem.LoginItems) error {
				return json.Unmarshal([]byte(validFile), items)
			},
		},
		{
			name:       "out of date",
			content:    validFile,
			createFile: true,
			get: func(items *loginitem.LoginItems) error {
				return json.Unmarshal([]byte(`[{"Name":"Other","Path":"/Other","Hidden":false}]`), items)
			},
			wantErr: true,
		},
		{
			name:       "System Events error",
			content:    validFile,
			createFile: true,
			get: func(*loginitem.LoginItems) error {
				return errors.New("System Events failed")
			},
			wantErr: true,
		},
		{name: "missing file", wantErr: true},
		{name: "invalid JSON", content: "{", createFile: true, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "login_items.json")
			if test.createFile {
				if err := os.WriteFile(path, []byte(test.content), 0600); err != nil {
					t.Fatal(err)
				}
			}
			command := New()
			if test.get != nil {
				command.getLoginItems = test.get
			}
			command.command.SetArgs([]string{"check", "--file", path})
			err := command.Execute()
			if (err != nil) != test.wantErr {
				t.Fatalf("Execute() error = %v, wantErr %v", err, test.wantErr)
			}
			if test.name == "out of date" && !strings.Contains(err.Error(), "out-of-date") {
				t.Fatalf("Execute() error = %v, want out-of-date error", err)
			}
		})
	}
}
