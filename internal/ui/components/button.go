// Package components provides shared UI components and styles.
//
// This file contains the TuiButton interactive button component.
package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// ButtonVariant defines the visual style of a button.
type ButtonVariant string

const (
	// ButtonPrimary is a filled button with cyan background for primary actions.
	ButtonPrimary ButtonVariant = "primary"

	// ButtonSecondary is an outlined button for secondary actions.
	ButtonSecondary ButtonVariant = "secondary"

	// ButtonDestructive is a red button for dangerous/destructive actions.
	ButtonDestructive ButtonVariant = "destructive"
)

// maxLabelWidth is the maximum width for button labels before truncation.
const maxLabelWidth = 30

// TuiButton renders an interactive button with optional keyboard shortcut.
// The shortcut is displayed as "[shortcut] label" if provided.
// The focused parameter adds visual emphasis (bold + underline).
func TuiButton(label string, variant ButtonVariant, shortcut string, focused bool) string {
	// Format the button text
	text := label
	if shortcut != "" {
		text = fmt.Sprintf("[%s] %s", shortcut, label)
	}

	// Truncate long labels (preserve shortcut visibility)
	maxWidth := maxLabelWidth
	if shortcut != "" {
		maxWidth += len(shortcut) + 3 // account for "[shortcut] "
	}
	text = Truncate(text, maxWidth)

	// Get base style for variant
	var style lipgloss.Style
	switch variant {
	case ButtonPrimary:
		style = lipgloss.NewStyle().
			Background(ColorCyan).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 2)
	case ButtonSecondary:
		style = lipgloss.NewStyle().
			Border(BorderSingle).
			BorderForeground(ColorBorder).
			Foreground(ColorTextPrimary).
			Padding(0, 2)
	case ButtonDestructive:
		style = lipgloss.NewStyle().
			Background(ColorDestructive).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 2)
	default:
		// Default to primary
		style = lipgloss.NewStyle().
			Background(ColorCyan).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 2)
	}

	// Apply focus state
	if focused {
		style = style.Bold(true).Underline(true)
	}

	return style.Render(text)
}
