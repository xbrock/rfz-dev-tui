// Package components provides the TuiStatusBar component.
//
// TuiStatusBar renders the footer/status bar at the bottom of the screen.
// Shows keyboard shortcuts and contextual information.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// StatusBarItem represents a single item in the status bar.
type StatusBarItem struct {
	Key    string // The keyboard shortcut (e.g., "q", "Enter")
	Label  string // The action label (e.g., "Quit", "Select")
	Active bool   // Whether this item is currently active/highlighted
}

// RenderStatusBar renders the status bar with the given items.
func RenderStatusBar(width int, items []StatusBarItem) string {
	var renderedItems []string

	for _, item := range items {
		renderedItems = append(renderedItems, renderStatusBarItem(item))
	}

	// Join items with separators
	content := lipgloss.JoinHorizontal(lipgloss.Top, renderedItems...)

	// Apply footer styling and ensure full width
	style := lipgloss.NewStyle().
		Background(ColorCard).
		Foreground(ColorTextSecondary).
		Width(width).
		Padding(0, 1)

	return style.Render(content)
}

// renderStatusBarItem renders a single status bar item.
func renderStatusBarItem(item StatusBarItem) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(ColorTextSecondary)

	if item.Active {
		// Active items have inverted colors
		return lipgloss.NewStyle().
			Background(ColorCyan).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 1).
			Render(item.Label)
	}

	// Normal item: "Key Label"
	rendered := keyStyle.Render(item.Key) + " " + labelStyle.Render(item.Label)

	// Add spacing between items
	return rendered + "  "
}

// RenderContextInfo renders contextual information for the status bar.
// For example: current build status, selected component count, etc.
func RenderContextInfo(info string) string {
	return lipgloss.NewStyle().
		Foreground(ColorTextMuted).
		Align(lipgloss.Right).
		Render(info)
}
