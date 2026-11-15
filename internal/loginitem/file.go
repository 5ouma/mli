package loginitem

import (
	"errors"
	"os"

	"go.yaml.in/yaml/v4"
)

func (loginItems *LoginItems) Save(path string, force bool) error {
	if isExist, err := isExist(path); err != nil {
		return err
	} else if !force && isExist {
		return os.ErrExist
	}

	data, err := yaml.Marshal(loginItems)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
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
	if err := yaml.Unmarshal(data, loginItems); err != nil {
		return err
	}

	return nil
}

func isExist(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
