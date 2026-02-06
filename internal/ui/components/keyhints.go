// Package components provides shared UI components and styles.
//
// This file contains the TuiKeyHints keyboard shortcut display component.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// SymbolKeySeparator is the middle dot separator between key hints.
const SymbolKeySeparator = "·"

// KeyHint represents a single keyboard shortcut hint.
type KeyHint struct {
	Key   string // The key(s) to press (e.g., "Enter", "q", "Esc")
	Label string // The action description (e.g., "Select", "Quit")
}

// TuiKeyHints renders a horizontal list of keyboard hints with separators.
// hints: slice of KeyHint items to render
// width: available width (0 = no width constraint, hints that don't fit are omitted)
func TuiKeyHints(hints []KeyHint, width int) string {
	if len(hints) == 0 {
		return ""
	}

	keyStyle := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
	separatorStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)

	separator := separatorStyle.Render(" " + SymbolKeySeparator + " ")
	separatorWidth := 3 // " · " = 3 chars

	var renderedHints []string
	usedWidth := 0

	for i, hint := range hints {
		rendered := keyStyle.Render(hint.Key) + " " + labelStyle.Render(hint.Label)
		hintWidth := lipgloss.Width(hint.Key) + 1 + lipgloss.Width(hint.Label)

		// Account for separator before this hint (not the first one)
		extraWidth := 0
		if i > 0 {
			extraWidth = separatorWidth
		}

		// Check if this hint fits within the width constraint
		if width > 0 && usedWidth+extraWidth+hintWidth > width {
			break
		}

		renderedHints = append(renderedHints, rendered)
		usedWidth += extraWidth + hintWidth
	}

	if len(renderedHints) == 0 {
		return ""
	}

	return strings.Join(renderedHints, separator)
}
