// Package components provides shared UI components and styles.
//
// This file contains the TuiRadio button group component.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// TuiRadio renders a single radio button with charm-style symbols.
// Format: "◯ Label" (unselected) or "◉ Label" (selected)
// The focused parameter highlights the radio button in cyan.
func TuiRadio(label string, selected bool, focused bool) string {
	// Select the symbol based on selected state
	symbol := SymbolRadioUnselected
	if selected {
		symbol = SymbolRadioSelected
	}

	// Determine the style based on state
	var style lipgloss.Style

	switch {
	case focused:
		// Focused state: cyan highlight
		style = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Bold(true)
	case selected:
		// Selected state: primary text color (slightly emphasized)
		style = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)
	default:
		// Normal state: secondary text color (unselected options are dimmer)
		style = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)
	}

	// Render symbol and label with consistent styling
	return style.Render(symbol + " " + label)
}

// TuiRadioGroup renders a group of radio buttons with one selected option.
// options: slice of option labels
// selectedIndex: index of the currently selected option (-1 for none)
// focusedIndex: index of the currently focused option (-1 for none)
// horizontal: if true, renders options horizontally; otherwise vertically
func TuiRadioGroup(options []string, selectedIndex int, focusedIndex int, horizontal bool) string {
	if len(options) == 0 {
		return ""
	}

	// Render each radio button
	rendered := make([]string, len(options))
	for i, opt := range options {
		isSelected := i == selectedIndex
		isFocused := i == focusedIndex
		rendered[i] = TuiRadio(opt, isSelected, isFocused)
	}

	// Join based on layout preference
	if horizontal {
		// Horizontal layout: options separated by spaces
		separator := lipgloss.NewStyle().SetString("  ").String()
		return lipgloss.JoinHorizontal(lipgloss.Top, intersperse(rendered, separator)...)
	}

	// Vertical layout: each option on its own line
	return lipgloss.JoinVertical(lipgloss.Left, rendered...)
}

// intersperse inserts a separator between each element of the slice.
func intersperse(items []string, separator string) []string {
	if len(items) <= 1 {
		return items
	}

	result := make([]string, 0, len(items)*2-1)
	for i, item := range items {
		result = append(result, item)
		if i < len(items)-1 {
			result = append(result, separator)
		}
	}
	return result
}
