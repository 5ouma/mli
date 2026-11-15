package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/5ouma/mli/internal/loginitem"
	"go.yaml.in/yaml/v4"
)

func TestSaveCommand(t *testing.T) {
	tests := []struct {
		name      string
		get       func(*loginitem.LoginItems) error
		existing  bool
		force     bool
		wantErr   bool
		wantSaved bool
	}{
		{
			name: "save current items",
			get: func(items *loginitem.LoginItems) error {
				return yaml.Unmarshal([]byte("- name: App\n  path: /App\n  hidden: true\n"), items)
			},
			wantSaved: true,
		},
		{
			name: "overwrite existing file",
			get: func(items *loginitem.LoginItems) error {
				return yaml.Unmarshal([]byte("- name: App\n  path: /App\n  hidden: true\n"), items)
			},
			existing:  true,
			force:     true,
			wantSaved: true,
		},
		{
			name: "refuse existing file",
			get: func(*loginitem.LoginItems) error {
				return nil
			},
			existing: true,
			wantErr:  true,
		},
		{
			name: "System Events error",
			get: func(*loginitem.LoginItems) error {
				return errors.New("System Events failed")
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "login_items.yml")
			if test.existing {
				if err := os.WriteFile(path, []byte("old"), 0600); err != nil {
					t.Fatal(err)
				}
			}
			command := New()
			command.getLoginItems = test.get
			command.command.SetArgs([]string{"save", "--file", path})
			if test.force {
				command.command.SetArgs([]string{"save", "--file", path, "--force"})
			}
			err := command.Execute()
			if (err != nil) != test.wantErr {
				t.Fatalf("Execute() error = %v, wantErr %v", err, test.wantErr)
			}
			if !test.wantSaved {
				return
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(data), "name: App") {
				t.Fatalf("saved file = %s, want App entry", data)
			}
		})
	}
}
