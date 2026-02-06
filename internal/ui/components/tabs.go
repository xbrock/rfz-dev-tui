// Package components provides shared UI components and styles.
//
// This file contains the TuiTabs horizontal tab navigation component.
package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// maxTabShortcuts is the maximum number of tabs that get numeric shortcuts (1-9).
const maxTabShortcuts = 9

// SymbolTabSeparator is the pipe separator between tabs.
const SymbolTabSeparator = "|"

// TuiTab represents a single tab item.
type TuiTab struct {
	Label  string // Display text for the tab
	Badge  int    // Optional count badge (0 = no badge)
	Active bool   // Whether this tab is the currently active view
}

// TuiTabs renders a horizontal tab bar with numeric shortcuts.
// tabs: slice of TuiTab items to render
// focusedIndex: index of the keyboard-focused tab (-1 for no focus)
// width: available width (0 = no width constraint)
func TuiTabs(tabs []TuiTab, focusedIndex int, width int) string {
	if len(tabs) == 0 {
		return ""
	}

	separatorStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
	separator := separatorStyle.Render(" " + SymbolTabSeparator + " ")
	separatorWidth := 3 // " | " = 3 chars

	var renderedTabs []string
	usedWidth := 0

	for i, tab := range tabs {
		rendered, tabWidth := renderTab(tab, i, focusedIndex)

		// Account for separator before this tab (not the first one)
		extraWidth := 0
		if i > 0 {
			extraWidth = separatorWidth
		}

		// Check if this tab fits within the width constraint
		if width > 0 && usedWidth+extraWidth+tabWidth > width {
			break
		}

		renderedTabs = append(renderedTabs, rendered)
		usedWidth += extraWidth + tabWidth
	}

	if len(renderedTabs) == 0 {
		return ""
	}

	return strings.Join(renderedTabs, separator)
}

// renderTab renders a single tab and returns the rendered string and its plain-text width.
func renderTab(tab TuiTab, index int, focusedIndex int) (string, int) {
	isFocused := index == focusedIndex

	// Build label with optional truncation
	label := tab.Label

	// Add badge if present
	badgeStr := ""
	if tab.Badge > 0 {
		badgeStr = fmt.Sprintf(" (%d)", tab.Badge)
	}

	// Build shortcut prefix (1-9 only)
	shortcutStr := ""
	if index < maxTabShortcuts {
		shortcutStr = fmt.Sprintf("%d:", index+1)
	}

	// Calculate plain-text width
	plainText := shortcutStr + label + badgeStr
	plainWidth := lipgloss.Width(plainText)

	// Choose style based on state
	switch {
	case tab.Active:
		return renderActiveTab(shortcutStr, label, badgeStr), plainWidth
	case isFocused:
		return renderFocusedTab(shortcutStr, label, badgeStr), plainWidth
	default:
		return renderNormalTab(shortcutStr, label, badgeStr), plainWidth
	}
}

func renderActiveTab(shortcut, label, badge string) string {
	style := StyleTabActive
	text := shortcut + label + badge
	return style.Render(text)
}

func renderFocusedTab(shortcut, label, badge string) string {
	style := StyleTabFocused
	text := shortcut + label + badge
	return style.Render(text)
}

func renderNormalTab(shortcut, label, badge string) string {
	var parts []string

	if shortcut != "" {
		shortcutStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		parts = append(parts, shortcutStyle.Render(shortcut))
	}

	labelStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
	parts = append(parts, labelStyle.Render(label))

	if badge != "" {
		badgeStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		parts = append(parts, badgeStyle.Render(badge))
	}

	return strings.Join(parts, "")
}
