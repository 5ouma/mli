package utils

import (
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
	items := []string{"/System/Applications/Photos.app", "/System/Applications/Notes.app", "/System/Applications/App Store.app"}
	t.Logf("  items: %s", items)
	for _, path := range items {
		if err := AddLoginItem(path); err != nil {
			t.Fatalf("🚨 %v", err)
		}
	}

	t.Log("📘 Read Login Items")
	if paths, err := GetLoginItemPaths(); err != nil {
		t.Fatalf("🚨 %v", err)
	} else {
		if !reflect.DeepEqual(paths, items) {
			t.Error("🚨 Items are not the same")
			t.Fatalf("  paths: %s", paths)
		}
	}
}
