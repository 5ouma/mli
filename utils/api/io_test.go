package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/5ouma/mli/utils/ci"
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("🏗️ Prepare LoginItems Structure")
			previousLoginItems := LoginItems{
				{"App Store", "/System/Applications/App Store.app", false},
				{"Calculator", "/System/Applications/Calculator.app", false},
				{"Notes", "/System/Applications/Notes.app", false},
				{"Photos", "/System/Applications/Photos.app", false},
			}
			data, _ := json.Marshal(previousLoginItems)
			t.Logf("  loginItems: %v", string(data))

			t.Log("🖋️ Save Login Items")
			dir := t.TempDir()
			path := filepath.Join(dir, "login_items.json")
			os.Create(path)
			err := previousLoginItems.Save(path, test.force)
			if !test.wantError && err != nil {
				t.Fatal(ci.Err(err))
			} else if test.wantError && err == nil {
				t.Fatal(ci.Err(fmt.Errorf("Expected an error")))
			}
			if test.wantError {
				t.Logf(`  "%v"`, err)
				t.SkipNow()
			}

			t.Log("📘 Read the Saved File")
			var currentLoginItems LoginItems
			if err := currentLoginItems.Load(path); err != nil {
				t.Fatal(ci.Err(err))
			}
			if !reflect.DeepEqual(previousLoginItems, currentLoginItems) {
				t.Error(ci.Err(ci.ErrorItemsNotSame))
				data, _ := json.Marshal(currentLoginItems)
				t.Fatalf("  currentLoginItems: %v", string(data))
			}
		})
	}
}
