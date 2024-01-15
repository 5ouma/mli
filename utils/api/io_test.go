package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_IO(t *testing.T) {
	tests := []struct {
		name      string
		force     bool
		wantError bool
	}{
		{
			name:      "✅ Force Save with no error",
			force:     true,
			wantError: false,
		},
		{
			name:      "🔴 Save with an exist error",
			force:     false,
			wantError: true,
		},
	}

	t.Log("🏗️ Prepare LoginItems Structure")
	previousLoginItems := LoginItems{
		{"App Store", "/System/Applications/App Store.app", false},
		{"Calculator", "/System/Applications/Calculator.app", false},
		{"Notes", "/System/Applications/Notes.app", false},
		{"Photos", "/System/Applications/Photos.app", false},
	}
	data, _ := json.Marshal(previousLoginItems)
	t.Logf("  loginItems: %v", string(data))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("🖋️ Save Login Items")
			path := filepath.Join(t.TempDir(), "login_items.json")
			os.Create(path)
			if err := previousLoginItems.Save(path, test.force); !test.wantError && err != nil {
				t.Fatalf("🚨 %v", err)
			} else if test.wantError && err == nil {
				t.Fatal("🚨 expected an error")
			} else if test.wantError {
				t.Logf(`  "%v"`, err)
				t.SkipNow()
			}

			t.Log("📘 Read the Saved File")
			var currentLoginItems LoginItems
			if err := currentLoginItems.Load(path); err != nil {
				t.Fatalf("🚨 %v", err)
			}
			if !reflect.DeepEqual(previousLoginItems, currentLoginItems) {
				t.Error("🚨 items are not the same")
				data, _ := json.Marshal(currentLoginItems)
				t.Fatalf("  currentLoginItems: %v", string(data))
			}
		})
	}
}
