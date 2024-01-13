package utils

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
)

func getLoginItem(entry string) ([]string, error) {
	data, err := mack.Tell("System Events", fmt.Sprintf("get %s of login items", entry))
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
	for i := range names {
		hidden, err := strconv.ParseBool(hides[i])
		if err != nil {
			return err
		}
		*loginItems = append(*loginItems, &loginItem{names[i], path[i], hidden})
	}
	sort.Slice(*loginItems, func(before, after int) bool {
		return (*loginItems)[before].Name < (*loginItems)[after].Name
	})
	return nil
}

func (loginItem *loginItem) Add() error {
	if _, err := os.Stat(loginItem.Path); err != nil {
		return errAppNotFound(loginItem.Path)
	}
	msg, err := mack.Tell("System Events", fmt.Sprintf(`make login item at end with properties { name: "%s", path: "%s", hidden: %v }`, loginItem.Name, loginItem.Path, loginItem.Hidden))
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}
