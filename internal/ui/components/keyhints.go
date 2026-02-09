// Package components provides shared UI components and styles.
//
// This file contains the TuiKeyHints keyboard shortcut display component.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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

	separator := " | " // pipe separator matching design prototype
	separatorWidth := 3

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

// Tree connector symbols for tree-style hint rendering.
const (
	SymbolTreeBranch = "├──" // Branch connector (non-last item)
	SymbolTreeLast   = "└──" // Last item connector
)

// TuiKeyHintsTree renders a vertical tree-style list of keyboard hints.
// hints: slice of KeyHint items to render
func TuiKeyHintsTree(hints []KeyHint) string {
	if len(hints) == 0 {
		return ""
	}

	keyStyle := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
	treeStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)

	var lines []string
	for i, hint := range hints {
		connector := SymbolTreeBranch
		if i == len(hints)-1 {
			connector = SymbolTreeLast
		}
		line := treeStyle.Render(connector) + " " + keyStyle.Render(hint.Key) + " " + labelStyle.Render(hint.Label)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
