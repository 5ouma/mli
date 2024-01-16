package lib

import (
	"encoding/json"
	"os"
)

func (loginItems *LoginItems) Save(path string, force bool) error {
	if isExist, err := isExist(path); err != nil {
		return err
	} else if !force && isExist && err == nil {
		return os.ErrExist
	}
	data, err := json.MarshalIndent(loginItems, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (loginItems *LoginItems) Load(path string) error {
	if isExist, err := isExist(path); err != nil {
		return err
	} else if !isExist {
		return os.ErrNotExist
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, loginItems); err != nil {
		return err
	}
	return nil
}
