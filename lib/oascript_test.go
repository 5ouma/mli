package lib

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		invalidPath bool
		empty       bool
		wantLoaded  bool
		wantErr     bool
	}{
		{name: "skip missing app"},
		{name: "empty list", empty: true, wantLoaded: true},
		{name: "return invalid path error", invalidPath: true, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "Missing.app")
			if test.invalidPath {
				path = "\x00"
			}
			var items LoginItems
			if !test.empty {
				items = LoginItems{{"App", path, false}}
			}
			loaded, err := items.Add()
			if (err != nil) != test.wantErr {
				t.Fatalf("Add() error = %v, wantErr %v", err, test.wantErr)
			}
			if loaded != test.wantLoaded {
				t.Fatalf("Add() loaded = %v, want %v", loaded, test.wantLoaded)
			}
		})
	}
}

func TestAddOutputError(t *testing.T) {
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
	items := LoginItems{{"Missing", filepath.Join(t.TempDir(), "Missing.app"), false}}
	loaded, addErr := items.Add()
	if addErr == nil {
		t.Fatal("Add() succeeded when its output destination was not writable")
	}
	if loaded {
		t.Fatal("Add() reported a missing app as loaded")
	}
}

func TestGetWithScriptedResults(t *testing.T) {
	dir := t.TempDir()
	script := `#!/bin/sh
case "$2" in
  *"get name of login items"*) printf '%s' "$LOGIN_ITEM_NAMES" ;;
  *"get path of login items"*) printf '%s' "$LOGIN_ITEM_PATHS" ;;
  *"get hidden of login items"*) printf '%s' "$LOGIN_ITEM_HIDDEN" ;;
  *) exit 1 ;;
esac
`
	if err := os.WriteFile(filepath.Join(dir, "osascript"), []byte(script), 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))

	tests := []struct {
		name    string
		names   string
		paths   string
		hidden  string
		wantErr bool
	}{
		{name: "empty results"},
		{name: "mismatched result lengths", names: "App1, App2", paths: "/App1", hidden: "false, false", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("LOGIN_ITEM_NAMES", test.names)
			t.Setenv("LOGIN_ITEM_PATHS", test.paths)
			t.Setenv("LOGIN_ITEM_HIDDEN", test.hidden)
			var items LoginItems
			err := items.Get()
			if (err != nil) != test.wantErr {
				t.Fatalf("Get() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestGetReadsCurrentLoginItems(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("System Events is only available on macOS")
	}

	output, err := os.CreateTemp(t.TempDir(), "mli-output-")
	if err != nil {
		t.Fatal(err)
	}
	originalOutput := lipgloss.Writer.Forward
	lipgloss.Writer.Forward = output
	defer func() {
		lipgloss.Writer.Forward = originalOutput
		if err := output.Close(); err != nil {
			t.Error(err)
		}
	}()

	var items LoginItems
	if err := items.Get(); err != nil {
		t.Fatal(err)
	}
	for i := 1; i < len(items); i++ {
		if items[i-1].Name > items[i].Name {
			t.Fatalf("Get() returned unsorted items at indexes %d and %d", i-1, i)
		}
	}
}
