package lib

import (
	"errors"
	"os"
)

type loginItem struct {
	Name   string
	Path   string
	Hidden bool
}
type LoginItems []*loginItem

func isExist(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
