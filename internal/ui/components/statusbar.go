// Package components provides shared UI components and styles.
//
// This file contains the TuiStatusBar bottom bar component.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// TuiStatusBarConfig holds the configuration for rendering a badge-based status bar.
type TuiStatusBarConfig struct {
	// Badge section (left)
	ModeBadge         string         // Mode text (e.g., "LOGS", "SELECT", "BUILD")
	ModeBadgeColor    lipgloss.Color // Background color for mode badge
	ContextBadge      string         // Context text (e.g., "rfz-dispatcher", "Boss")
	ContextBadgeColor lipgloss.Color // Background color for context badge (optional, defaults to ColorSecondary)

	// Key hints section (right)
	Hints    []KeyHint // Main key hints (rendered via TuiKeyHints)
	QuitHint *KeyHint  // Separate quit hint (far right, visually separated)

	// Layout
	Width int // Total available width
}

// TuiStatusBar renders a full-width bottom status bar with badge-based layout.
// Layout: [ModeBadge] [ContextBadge] ...gap... [Hints] [q Quit]
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

	// Build left section: badges
	leftParts := []string{}
	leftWidth := 0

	if cfg.ModeBadge != "" {
		badgeColor := cfg.ModeBadgeColor
		if badgeColor == "" {
			badgeColor = ColorCyan
		}
		badge := lipgloss.NewStyle().
			Background(badgeColor).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 1).
			Render(cfg.ModeBadge)
		leftParts = append(leftParts, badge)
		leftWidth += lipgloss.Width(badge)
	}

	if cfg.ContextBadge != "" {
		badgeColor := cfg.ContextBadgeColor
		if badgeColor == "" {
			badgeColor = ColorSecondary
		}
		badge := lipgloss.NewStyle().
			Background(badgeColor).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 1).
			Render(cfg.ContextBadge)
		if len(leftParts) > 0 {
			leftWidth++ // space between badges
		}
		leftParts = append(leftParts, badge)
		leftWidth += lipgloss.Width(badge)
	}

	leftContent := ""
	if len(leftParts) > 0 {
		leftContent = lipgloss.JoinHorizontal(lipgloss.Top, interleave(leftParts, " ")...)
	}

	// Build right section: hints + quit
	rightContent := ""
	rightWidth := 0

	if len(cfg.Hints) > 0 || cfg.QuitHint != nil {
		maxHintsWidth := innerWidth - leftWidth - 2 // leave gap
		if maxHintsWidth < 0 {
			maxHintsWidth = 0
		}

		hintsStr := ""
		if len(cfg.Hints) > 0 {
			hintsStr = TuiKeyHints(cfg.Hints, maxHintsWidth)
		}

		if cfg.QuitHint != nil {
			quitStr := FooterItem(cfg.QuitHint.Key, cfg.QuitHint.Label)
			if hintsStr != "" {
				rightContent = hintsStr + "  " + quitStr
			} else {
				rightContent = quitStr
			}
		} else {
			rightContent = hintsStr
		}
		rightWidth = lipgloss.Width(rightContent)
	}

	// Calculate gap between left and right
	gapWidth := innerWidth - leftWidth - rightWidth
	if gapWidth < 1 {
		gapWidth = 1
	}

	gap := lipgloss.NewStyle().Width(gapWidth).Render("")

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		leftContent,
		gap,
		rightContent,
	)

	return barStyle.Render(content)
}

// interleave inserts sep between each element of parts.
func interleave(parts []string, sep string) []string {
	if len(parts) <= 1 {
		return parts
	}
	result := make([]string, 0, len(parts)*2-1)
	for i, p := range parts {
		if i > 0 {
			result = append(result, sep)
		}
		result = append(result, p)
	}
	return result
}
