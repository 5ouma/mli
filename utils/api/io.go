package api

import (
	"encoding/json"
	"os"

	"github.com/5ouma/mli/utils"
)

func (loginItems *LoginItems) Save(path string, force bool) error {
	if isExist, err := utils.IsExist(path); err != nil {
		return err
	} else if !force && isExist && err == nil {
		return os.ErrExist
	}
	data, _ := json.MarshalIndent(loginItems, "", "  ")
	if err := os.WriteFile(path, data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (loginItems *LoginItems) Load(path string) error {
	if isExist, err := utils.IsExist(path); err != nil {
		return err
	} else if !isExist && err == nil {
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
