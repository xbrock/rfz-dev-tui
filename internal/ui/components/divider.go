// Package components provides shared UI components and styles.
//
// This file contains the TuiDivider horizontal separator component.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// DividerStyle defines the line style for TuiDivider.
type DividerStyle string

const (
	// DividerSingle uses single horizontal lines (─).
	DividerSingle DividerStyle = "single"

	// DividerDouble uses double horizontal lines (═).
	DividerDouble DividerStyle = "double"
)

// dividerChars maps DividerStyle to the character used for the line.
var dividerChars = map[DividerStyle]string{
	DividerSingle: "─",
	DividerDouble: "═",
}

// TuiDivider renders a horizontal divider line.
// Returns an empty string if width <= 0.
func TuiDivider(style DividerStyle, width int) string {
	if width <= 0 {
		return ""
	}

	char, ok := dividerChars[style]
	if !ok {
		char = dividerChars[DividerSingle]
	}

	line := strings.Repeat(char, width)

	dividerStyle := lipgloss.NewStyle().
		Foreground(ColorBorder)

	return dividerStyle.Render(line)
}
