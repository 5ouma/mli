package lib

import (
	"errors"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type loginItem struct {
	Name   string
	Path   string
	Hidden bool
}
type LoginItems []*loginItem

var (
	heading = lipgloss.NewStyle().
		Foreground(lipgloss.CompleteColor{TrueColor: "#007aff", ANSI256: "27"}).
		Bold(true).
		Padding(1)
	H1 = heading
	H2 = heading.SetString("▌")

	item = lipgloss.NewStyle().
		PaddingLeft(2)
	CheckedItem = item.
			Foreground(lipgloss.CompleteColor{TrueColor: "#63b946", ANSI256: "41"}).
			SetString("✔︎")
	WarnedItem = item.
			Foreground(lipgloss.CompleteColor{TrueColor: "#ffc627", ANSI256: "226"}).
			SetString("⚠︎")
)

func isExist(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
