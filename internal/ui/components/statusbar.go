// Package components provides shared UI components and styles.
//
// This file contains the TuiStatusBar bottom bar component.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// TuiStatusBarConfig holds the configuration for rendering a status bar.
type TuiStatusBarConfig struct {
	Status     string    // Left section: status text (e.g., "Build: Running")
	StatusColor lipgloss.Color // Color for the status text (e.g., ColorGreen, ColorYellow, ColorDestructive)
	Info       string    // Center section: screen name or info text
	Hints      []KeyHint // Right section: keyboard hints (rendered via TuiKeyHints)
	Width      int       // Total available width
}

// TuiStatusBar renders a full-width bottom status bar with three sections.
// Left: status text with configurable color
// Center: info/screen name
// Right: keyboard hints via TuiKeyHints
func TuiStatusBar(cfg TuiStatusBarConfig) string {
	if cfg.Width <= 0 {
		return ""
	}

	barStyle := StyleFooter.Width(cfg.Width)

	// StyleFooter has Padding(0, 1) → 2 chars of horizontal padding
	innerWidth := cfg.Width - 2
	if innerWidth <= 0 {
		return barStyle.Render("")
	}

	// Render right section first (priority: hints should not be truncated)
	rightContent := ""
	rightWidth := 0
	if len(cfg.Hints) > 0 {
		maxHintsWidth := innerWidth / 2
		rightContent = TuiKeyHints(cfg.Hints, maxHintsWidth)
		rightWidth = lipgloss.Width(rightContent)
	}

	// Calculate remaining width for left and center
	remaining := innerWidth - rightWidth
	if remaining < 0 {
		remaining = 0
	}

	// Split remaining between left and center
	leftMaxWidth := remaining / 2
	centerMaxWidth := remaining - leftMaxWidth

	// Render left section (status)
	leftContent := ""
	leftWidth := 0
	if cfg.Status != "" {
		statusText := Truncate(cfg.Status, leftMaxWidth)
		statusStyle := lipgloss.NewStyle().Bold(true)
		if cfg.StatusColor != "" {
			statusStyle = statusStyle.Foreground(cfg.StatusColor)
		} else {
			statusStyle = statusStyle.Foreground(ColorTextSecondary)
		}
		leftContent = statusStyle.Render(statusText)
		leftWidth = lipgloss.Width(leftContent)
	}

	// Render center section (info)
	centerContent := ""
	if cfg.Info != "" {
		infoText := Truncate(cfg.Info, centerMaxWidth)
		infoStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
		centerContent = infoStyle.Render(infoText)
	}

	// Build the three-column layout
	// Left-aligned status | centered info | right-aligned hints
	centerWidth := innerWidth - leftWidth - rightWidth
	if centerWidth < 0 {
		centerWidth = 0
	}

	centerSection := lipgloss.NewStyle().
		Width(centerWidth).
		Align(lipgloss.Center).
		Render(centerContent)

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		leftContent,
		centerSection,
		rightContent,
	)

	return barStyle.Render(content)
}
