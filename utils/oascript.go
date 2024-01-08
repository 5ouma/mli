package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/andybrewer/mack"
)

func GetLoginItemPaths() ([]string, error) {
	data, err := mack.Tell("System Events", "get path of login items")
	if err != nil {
		return nil, err
	}
	return strings.Split(data, ", "), nil
}

func AddLoginItem(path string) error {
	if _, err := os.Stat(path); err != nil {
		return errAppNotFound(path)
	}
	name := getFileName(path)
	msg, err := mack.Tell("System Events", fmt.Sprintf(`make login item at end with properties { name: "%s", path: "%s", hidden: true }`, name, path))
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}
