package utils

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/andybrewer/mack"
)

func Test_OAScript(t *testing.T) {
	t.Log("🗑️ Delete Login Items")
	if _, err := mack.Tell("System Events", "delete login items"); err != nil {
		t.Fatalf("🚨 %v", err)
	}

	t.Log("🖋️ Add Login Items")
	previousLoginItems := LoginItems{
		{"App Store", "/System/Applications/App Store.app", false},
		{"Calculator", "/System/Applications/Calculator.app", false},
		{"Notes", "/System/Applications/Notes.app", false},
		{"Photos", "/System/Applications/Photos.app", false},
	}
	data, _ := json.Marshal(previousLoginItems)
	t.Logf("  previousLoginItems: %v", string(data))
	if err := previousLoginItems.Add(); err != nil {
		t.Fatalf("🚨 %v", err)
	}

	t.Log("📘 Read Login Items")
	currentLoginItems := LoginItems{}
	if err := currentLoginItems.Get(); err != nil {
		t.Fatalf("🚨 %v", err)
	}
	if !reflect.DeepEqual(previousLoginItems, currentLoginItems) {
		t.Error("🚨 Items are not the same")
		data, _ := json.Marshal(currentLoginItems)
		t.Fatalf("  currentLoginItems: %v", string(data))
	}
}

func ExampleLoginItems() {
	loginItems := LoginItems{{"Not found", "/Applications/Not found.app", false}}
	loginItems.Add()
	// Output:
	// ⚠️ app not found: "/Applications/Not found.app"
}
