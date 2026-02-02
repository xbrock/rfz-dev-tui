// Package components provides the TuiNavItem component.
//
// TuiNavItem renders navigation menu items with proper focus and active states.
// Format: "> 1. Build Components  1" (focused) or "  2. View Logs  2" (normal)
package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// NavItemState represents the state of a navigation item.
type NavItemState int

const (
	NavItemNormal NavItemState = iota
	NavItemFocused
	NavItemActive
	NavItemFocusedActive // Both focused and active
)

// TuiNavItem renders a navigation menu item.
//
// Parameters:
// - index: The menu item number (1-5)
// - label: The display label (e.g., "Build Components")
// - shortcut: The keyboard shortcut hint (e.g., "1")
// - state: The current state of the item
func TuiNavItem(index int, label string, shortcut string, state NavItemState) string {
	// Cursor prefix
	prefix := "  "
	if state == NavItemFocused || state == NavItemFocusedActive {
		prefix = "> "
	}

	// Build the main text
	text := fmt.Sprintf("%s%d. %s", prefix, index, label)

	// Select style based on state
	var style lipgloss.Style
	switch state {
	case NavItemFocused:
		style = StyleNavItemFocused
	case NavItemActive:
		style = StyleNavItemActive
	case NavItemFocusedActive:
		// Combine focused and active styling
		style = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Background(ColorSecondary).
			Bold(true).
			PaddingLeft(0)
	default:
		style = StyleNavItem
	}

	// Render the main text
	rendered := style.Render(text)

	// Render the shortcut hint on the right
	shortcutStyle := lipgloss.NewStyle().
		Foreground(ColorTextMuted)
	shortcutRendered := shortcutStyle.Render(shortcut)

	// Join with spacing
	return lipgloss.JoinHorizontal(lipgloss.Top, rendered, "  ", shortcutRendered)
}

// TuiNavigation renders a complete navigation panel.
//
// Parameters:
// - items: List of navigation item labels
// - activeIndex: The currently active screen (0-based)
// - focusedIndex: The currently focused item (0-based, -1 if none)
// - width: The desired width of the navigation panel
func TuiNavigation(items []string, activeIndex, focusedIndex, width int) string {
	var renderedItems []string

	for i, label := range items {
		shortcut := fmt.Sprintf("%d", i+1)

		// Determine state
		var state NavItemState
		isFocused := i == focusedIndex
		isActive := i == activeIndex

		switch {
		case isFocused && isActive:
			state = NavItemFocusedActive
		case isFocused:
			state = NavItemFocused
		case isActive:
			state = NavItemActive
		default:
			state = NavItemNormal
		}

		renderedItems = append(renderedItems, TuiNavItem(i+1, label, shortcut, state))
	}

	// Join all items vertically
	content := lipgloss.JoinVertical(lipgloss.Left, renderedItems...)

	// Apply navigation panel styling
	panelStyle := lipgloss.NewStyle().
		Border(BorderSingle).
		BorderForeground(ColorBorder).
		Width(width).
		Padding(1)

	return panelStyle.Render(content)
}

// TuiNavHelp renders the help footer within navigation.
func TuiNavHelp(shortcuts []string) string {
	helpStyle := lipgloss.NewStyle().
		Foreground(ColorTextMuted).
		MarginTop(1).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(ColorBorder).
		PaddingTop(1)

	content := lipgloss.JoinVertical(lipgloss.Left, shortcuts...)
	return helpStyle.Render(content)
}
