package api

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func Test_Save(t *testing.T) {
	tests := []struct {
		name          string
		force         bool
		expectedError error
	}{
		{
			name:          "✅ Save correct with force",
			force:         true,
			expectedError: nil,
		},
		{
			name:          "🔴 Save failed with an already exists error",
			force:         false,
			expectedError: os.ErrExist,
		},
	}

	t.Log("🏗️ Prepare LoginItems Structure")
	loginItems := LoginItems{
		{"App Store", "/System/Applications/App Store.app", false},
		{"Calculator", "/System/Applications/Calculator.app", false},
		{"Notes", "/System/Applications/Notes.app", false},
		{"Photos", "/System/Applications/Photos.app", false},
	}
	data, err := json.Marshal(loginItems)
	if err != nil {
		t.Fatalf("🚨 %v", err)
	}
	t.Log(string(data))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("🖋️ Save Login Items")
			path := filepath.Join(t.TempDir(), "login_items.json")
			os.Create(path)
			err := loginItems.Save(path, test.force)
			if test.expectedError == nil && err != nil {
				t.Errorf("🚨 %v", err)
			} else if test.expectedError != nil && err == nil {
				t.Error("🚨 expected and error")
			} else if !errors.Is(test.expectedError, err) {
				t.Errorf("🚨 expected: %v", test.expectedError)
				t.Errorf("  actual: %v", err)
			}
		})
	}
}

func Test_Load(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		createFile    bool
		expectedError error
	}{
		{
			name: "✅ Load correctly",
			content: `[
  {
    "Name": "App Store",
    "Path": "/System/Applications/App Store.app",
    "Hidden": false
  },
  {
    "Name": "Calculator",
    "Path": "/System/Applications/Calculator.app",
    "Hidden": false
  },
  {
    "Name": "Notes",
    "Path": "/System/Applications/Notes.app",
    "Hidden": false
  },
  {
    "Name": "Photos",
    "Path": "/System/Applications/Photos.app",
    "Hidden": false
  }
]
`,
			createFile:    true,
			expectedError: nil,
		},
		{
			name: "🔴 Load failed with a not exist error",
			content: `[
  {
    "Name": "App Store",
    "Path": "/System/Applications/App Store.app",
    "Hidden": false
  },
  {
    "Name": "Calculator",
    "Path": "/System/Applications/Calculator.app",
    "Hidden": false
  },
  {
    "Name": "Notes",
    "Path": "/System/Applications/Notes.app",
    "Hidden": false
  },
  {
    "Name": "Photos",
    "Path": "/System/Applications/Photos.app",
    "Hidden": false
  }
]
`,
			createFile:    false,
			expectedError: os.ErrNotExist,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("🏗️ Prepare JSON Structure")
			t.Log(test.content)

			path := filepath.Join(t.TempDir(), "login_items.json")
			if test.createFile {
				t.Log("🖋️ Save Login Items")
				if err := os.WriteFile(path, []byte(test.content), os.ModePerm); err != nil {
					t.Fatalf("🚨 %v", err)
				}
			}

			t.Log("📘 Read the Saved File")
			var loginItems LoginItems
			err := loginItems.Load(path)
			if test.expectedError == nil && err != nil {
				t.Errorf("🚨 %v", err)
			} else if test.expectedError != nil && err == nil {
				t.Error("🚨 expected and error")
			} else if !errors.Is(test.expectedError, err) {
				t.Errorf("🚨 expected: %v", test.expectedError)
				t.Errorf("  actual: %v", err)
			}
		})
	}
}
