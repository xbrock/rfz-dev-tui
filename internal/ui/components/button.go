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
// Format: "[Label] shortcut" where shortcut is in a small bordered box.
// The focused parameter adds solid background fill.
func TuiButton(label string, variant ButtonVariant, shortcut string, focused bool) string {
	// Format the label with brackets
	labelText := fmt.Sprintf("[%s]", Truncate(label, maxLabelWidth))

	// Get label style based on focus state and variant
	var labelStyle lipgloss.Style
	if focused {
		// Focused state: solid background with variant color
		switch variant {
		case ButtonPrimary:
			labelStyle = lipgloss.NewStyle().
				Background(ColorGreen).
				Foreground(ColorBackground).
				Bold(true).
				Padding(0, 1)
		case ButtonSecondary:
			labelStyle = lipgloss.NewStyle().
				Background(ColorCyan).
				Foreground(ColorBackground).
				Bold(true).
				Padding(0, 1)
		case ButtonDestructive:
			labelStyle = lipgloss.NewStyle().
				Background(ColorDestructive).
				Foreground(ColorBackground).
				Bold(true).
				Padding(0, 1)
		default:
			labelStyle = lipgloss.NewStyle().
				Background(ColorGreen).
				Foreground(ColorBackground).
				Bold(true).
				Padding(0, 1)
		}
	} else {
		// Unfocused state: text only, variant determines color
		switch variant {
		case ButtonPrimary:
			labelStyle = lipgloss.NewStyle().
				Foreground(ColorGreen)
		case ButtonSecondary:
			labelStyle = lipgloss.NewStyle().
				Foreground(ColorTextPrimary)
		case ButtonDestructive:
			labelStyle = lipgloss.NewStyle().
				Foreground(ColorDestructive)
		default:
			labelStyle = lipgloss.NewStyle().
				Foreground(ColorGreen)
		}
	}

	// Render the label
	result := labelStyle.Render(labelText)

	// Add shortcut with subtle grey background (keyboard hint style)
	if shortcut != "" {
		shortcutStyle := lipgloss.NewStyle().
			Background(ColorSecondary).
			Foreground(ColorTextMuted).
			Padding(0, 1)
		result += " " + shortcutStyle.Render(shortcut)
	}

	return result
}
