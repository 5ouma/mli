package loginitem

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/5ouma/mli/internal/utils"
	"github.com/andybrewer/mack"
)

func getLoginItem(entry string) ([]string, error) {
	data, err := mack.Tell("System Events", fmt.Sprintf("get %s of login items", entry))
	if data == "" {
		return nil, err
	}
	return strings.Split(data, ", "), err
}

func (loginItems *LoginItems) Get() error {
	names, err := getLoginItem("name")
	if err != nil {
		return err
	}
	path, err := getLoginItem("path")
	if err != nil {
		return err
	}
	hides, err := getLoginItem("hidden")
	if err != nil {
		return err
	}
	if len(names) != len(path) || len(names) != len(hides) {
		return fmt.Errorf("login item data length mismatch: names=%d paths=%d hidden=%d", len(names), len(path), len(hides))
	}

	for i := range names {
		hidden, err := strconv.ParseBool(hides[i])
		if err != nil {
			return err
		}
		*loginItems = append(*loginItems, &loginItem{names[i], path[i], hidden})
		if _, err := lipgloss.Println(utils.CheckedItem.Render(), names[i]); err != nil {
			return err
		}
	}
	sort.Slice(*loginItems, func(before, after int) bool {
		return (*loginItems)[before].Name < (*loginItems)[after].Name
	})

	return nil
}

func (loginItems *LoginItems) Add() (bool, error) {
	loadedCorrectly := true
	for _, loginItem := range *loginItems {
		exists, err := isExist(loginItem.Path)
		if err != nil {
			return false, err
		}
		if !exists {
			if _, err := lipgloss.Println(utils.WarnedItem.Render(), loginItem.Path); err != nil {
				return false, err
			}
			loadedCorrectly = false
			continue
		}
		if _, err := mack.Tell("System Events", fmt.Sprintf(`make login item at end with properties { name: "%s", path: "%s", hidden: %v }`, loginItem.Name, loginItem.Path, loginItem.Hidden)); err != nil {
			return false, err
		}
		if _, err := lipgloss.Println(utils.CheckedItem.Render(), loginItem.Name); err != nil {
			return false, err
		}
	}

	return loadedCorrectly, nil
}
