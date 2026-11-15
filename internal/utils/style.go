package utils

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

var (
	heading = lipgloss.NewStyle().
		Foreground(compat.CompleteColor{TrueColor: lipgloss.Color("#007aff"), ANSI256: lipgloss.Color("27")}).
		Bold(true).
		Padding(1)
	// Renders the primary heading for CLI output.
	H1 = heading
	// Renders the secondary heading for CLI output.
	H2 = heading.SetString("▌")

	item = lipgloss.NewStyle().
		PaddingLeft(2)
	// Renders the success marker for a login item.
	CheckedItem = item.
			Foreground(compat.CompleteColor{TrueColor: lipgloss.Color("#63b946"), ANSI256: lipgloss.Color("41")}).
			SetString("✔︎")
	// Renders the warning marker for a login item.
	WarnedItem = item.
			Foreground(compat.CompleteColor{TrueColor: lipgloss.Color("#ffc627"), ANSI256: lipgloss.Color("226")}).
			SetString("⚠︎")
)
