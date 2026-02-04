// Package components provides shared UI components and styles.
//
// This file contains the TuiCheckbox toggle component.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// maxCheckboxLabelWidth is the maximum width for checkbox labels before truncation.
const maxCheckboxLabelWidth = 40

// TuiCheckbox renders a checkbox toggle with charm-style symbols.
// Format: "☐ Label" (unchecked) or "☑ Label" (checked)
// The focused parameter highlights the checkbox in cyan.
// The disabled parameter renders in muted colors.
func TuiCheckbox(label string, checked bool, focused bool, disabled bool) string {
	// Select the symbol based on checked state
	symbol := SymbolCheckboxUnchecked
	if checked {
		symbol = SymbolCheckboxChecked
	}

	// Truncate the label if needed
	truncatedLabel := Truncate(label, maxCheckboxLabelWidth)

	// Determine the style based on state
	var style lipgloss.Style

	switch {
	case disabled:
		// Disabled state: muted colors for both symbol and text
		style = lipgloss.NewStyle().
			Foreground(ColorTextDisabled)
	case focused:
		// Focused state: cyan highlight
		style = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Bold(true)
	default:
		// Normal state: primary text color
		style = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)
	}

	// Render symbol and label with consistent styling
	return style.Render(symbol + " " + truncatedLabel)
}
