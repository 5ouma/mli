package loginitem

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

//go:embed testdata/osascript
var fakeOSAScriptContents string

type fakeOSAScriptResults struct {
	names, paths, hidden string
	failOn               string
}

type errorWriter struct{}

func (errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func setFakeOSAScript(t *testing.T, results fakeOSAScriptResults) string {
	t.Helper()
	dir := t.TempDir()
	logPath := filepath.Join(dir, "commands.log")
	scriptPath := filepath.Join(dir, "osascript")
	if err := os.WriteFile(scriptPath, []byte(fakeOSAScriptContents), 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("LOGIN_ITEM_COMMAND_LOG", logPath)
	t.Setenv("LOGIN_ITEM_FAIL_ON", results.failOn)
	t.Setenv("LOGIN_ITEM_NAMES", results.names)
	t.Setenv("LOGIN_ITEM_PATHS", results.paths)
	t.Setenv("LOGIN_ITEM_HIDDEN", results.hidden)
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	return logPath
}

func setLipglossOutput(t *testing.T) {
	t.Helper()
	originalOutput := lipgloss.Writer.Forward
	lipgloss.Writer.Forward = errorWriter{}
	t.Cleanup(func() {
		lipgloss.Writer.Forward = originalOutput
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		itemPath    string
		empty       bool
		pathExists  bool
		results     fakeOSAScriptResults
		outputError bool
		wantLoaded  bool
		wantErr     string
		wantCommand string
	}{
		{name: "skip missing app"},
		{name: "empty list", empty: true, wantLoaded: true},
		{name: "invalid path", itemPath: "\x00", wantErr: "invalid argument"},
		{
			name:        "existing item",
			pathExists:  true,
			wantLoaded:  true,
			wantCommand: `make login item at end with properties { name: "App", path: "`,
		},
		{
			name:       "System Events error",
			pathExists: true,
			results:    fakeOSAScriptResults{failOn: "add"},
			wantErr:    "add failed",
		},
		{name: "missing item output error", outputError: true, wantErr: "write failed"},
		{name: "existing item output error", pathExists: true, outputError: true, wantErr: "write failed"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			commandLog := setFakeOSAScript(t, test.results)
			var items LoginItems
			var path string
			if !test.empty {
				path = test.itemPath
				if path == "" {
					path = filepath.Join(t.TempDir(), "Missing.app")
				}
				if test.pathExists {
					path = filepath.Join(t.TempDir(), "App.app")
					if err := os.WriteFile(path, nil, 0600); err != nil {
						t.Fatal(err)
					}
				}
				items = LoginItems{{"App", path, false}}
			}
			if test.outputError {
				setLipglossOutput(t)
			}
			loaded, err := items.Add()
			if test.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantErr) {
					t.Fatalf("Add() error = %v, want error containing %q", err, test.wantErr)
				}
			} else if err != nil {
				t.Fatalf("Add() error = %v", err)
			}
			if loaded != test.wantLoaded {
				t.Fatalf("Add() loaded = %v, want %v", loaded, test.wantLoaded)
			}
			if test.wantCommand != "" {
				gotCommand, err := os.ReadFile(commandLog)
				if err != nil {
					t.Fatal(err)
				}
				if !strings.Contains(string(gotCommand), test.wantCommand+path+`", hidden: false }`) ||
					!strings.Contains(string(gotCommand), `tell application "System Events"`) {
					t.Fatalf("osascript command = %q, want command for %q", gotCommand, test.wantCommand+path)
				}
			}
		})
	}
}

func TestGetWithSystemEventsResults(t *testing.T) {
	tests := []struct {
		name        string
		results     fakeOSAScriptResults
		want        LoginItems
		wantErr     string
		outputError bool
	}{
		{name: "empty results"},
		{
			name: "successful results are sorted",
			results: fakeOSAScriptResults{
				names:  "Zed, Alpha",
				paths:  "/Zed, /Alpha",
				hidden: "false, true",
			},
			want: LoginItems{{"Alpha", "/Alpha", true}, {"Zed", "/Zed", false}},
		},
		{
			name: "mismatched result lengths",
			results: fakeOSAScriptResults{
				names:  "App1, App2",
				paths:  "/App1",
				hidden: "false, false",
			},
			wantErr: "length mismatch",
		},
		{
			name: "invalid hidden value",
			results: fakeOSAScriptResults{
				names:  "App",
				paths:  "/App",
				hidden: "sometimes",
			},
			wantErr: "strconv.ParseBool",
		},
		{
			name:    "name query error",
			results: fakeOSAScriptResults{failOn: "name"},
			wantErr: "name failed",
		},
		{
			name: "path query error",
			results: fakeOSAScriptResults{
				names:  "App",
				failOn: "path",
			},
			wantErr: "path failed",
		},
		{
			name: "hidden query error",
			results: fakeOSAScriptResults{
				names:  "App",
				paths:  "/App",
				failOn: "hidden",
			},
			wantErr: "hidden failed",
		},
		{
			name: "output error",
			results: fakeOSAScriptResults{
				names:  "App",
				paths:  "/App",
				hidden: "false",
			},
			wantErr:     "write failed",
			outputError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setFakeOSAScript(t, test.results)
			if test.outputError {
				setLipglossOutput(t)
			}
			var items LoginItems
			err := items.Get()
			if test.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantErr) {
					t.Fatalf("Get() error = %v, want error containing %q", err, test.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Get() error = %v", err)
			}
			if !reflect.DeepEqual(items, test.want) {
				t.Fatalf("Get() = %#v, want %#v", items, test.want)
			}
		})
	}
}
