package lib

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	tests := []struct {
		name        string
		fileExists  bool
		force       bool
		path        string
		wantErr     bool
		wantIs      error
		wantContent bool
	}{
		{name: "create", wantContent: true},
		{name: "overwrite", fileExists: true, force: true, wantContent: true},
		{name: "reject existing file", fileExists: true, wantErr: true, wantIs: os.ErrExist},
		{name: "missing parent", path: "missing/file.json", wantErr: true, wantIs: os.ErrNotExist},
		{name: "invalid path", path: "\x00", wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "login_items.json")
			if test.path != "" {
				path = filepath.Join(t.TempDir(), test.path)
			}
			if test.fileExists {
				if err := os.WriteFile(path, []byte("old"), 0600); err != nil {
					t.Fatal(err)
				}
			}

			items := LoginItems{{"Calculator", "/Applications/Calculator.app", true}}
			err := items.Save(path, test.force)
			if (err != nil) != test.wantErr || (test.wantIs != nil && !errors.Is(err, test.wantIs)) {
				t.Fatalf("Save() error = %v, wantErr %v (matching %v)", err, test.wantErr, test.wantIs)
			}
			if !test.wantContent {
				return
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var got LoginItems
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, items) {
				t.Fatalf("saved items = %#v, want %#v", got, items)
			}
			if data[len(data)-1] != '\n' {
				t.Fatal("saved JSON does not end in a newline")
			}
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		fileExists bool
		directory  bool
		want       LoginItems
		wantErr    bool
		wantIs     error
	}{
		{
			name:       "valid JSON",
			content:    `[{"Name":"Notes","Path":"/Applications/Notes.app","Hidden":false}]`,
			fileExists: true,
			want:       LoginItems{{"Notes", "/Applications/Notes.app", false}},
		},
		{name: "missing file", wantErr: true, wantIs: os.ErrNotExist},
		{name: "invalid JSON", content: "{", fileExists: true, wantErr: true},
		{name: "read directory", directory: true, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "login_items.json")
			if test.fileExists {
				if err := os.WriteFile(path, []byte(test.content), 0600); err != nil {
					t.Fatal(err)
				}
			}
			if test.directory {
				if err := os.Mkdir(path, 0700); err != nil {
					t.Fatal(err)
				}
			}

			var got LoginItems
			err := got.Load(path)
			if (err != nil) != test.wantErr || (test.wantIs != nil && !errors.Is(err, test.wantIs)) {
				t.Fatalf("Load() error = %v, wantErr %v (matching %v)", err, test.wantErr, test.wantIs)
			}
			if !test.wantErr && !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Load() = %#v, want %#v", got, test.want)
			}
		})
	}
}
