// Package components provides the TuiStatus component.
//
// TuiStatus renders build status badges with appropriate colors.
// Variants: pending, running, success, failed, error
package components

import "github.com/charmbracelet/lipgloss"

// Status represents the build status type
type Status int

const (
	StatusPending Status = iota
	StatusRunning
	StatusSuccess
	StatusFailed
	StatusError
	StatusSkipped
)

// String returns the display string for a status.
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusRunning:
		return "RUNNING"
	case StatusSuccess:
		return "SUCCESS"
	case StatusFailed:
		return "FAILED"
	case StatusError:
		return "ERROR"
	case StatusSkipped:
		return "SKIPPED"
	default:
		return "UNKNOWN"
	}
}

// TuiStatus renders a status badge with appropriate styling.
// Uses Lip Gloss for all styling - no manual ANSI codes.
func TuiStatus(status Status) string {
	var style lipgloss.Style

	switch status {
	case StatusPending:
		style = lipgloss.NewStyle().
			Background(ColorSecondary).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 1)

	case StatusRunning:
		style = lipgloss.NewStyle().
			Background(ColorCyan).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 1)

	case StatusSuccess:
		style = lipgloss.NewStyle().
			Background(ColorGreen).
			Foreground(ColorBackground).
			Bold(true).
			Padding(0, 1)

	case StatusFailed:
		style = lipgloss.NewStyle().
			Background(ColorDestructive).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 1)

	case StatusError:
		style = lipgloss.NewStyle().
			Background(ColorDestructive).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 1)

	case StatusSkipped:
		style = lipgloss.NewStyle().
			Background(ColorSecondary).
			Foreground(ColorTextMuted).
			Padding(0, 1)

	default:
		style = lipgloss.NewStyle().
			Background(ColorSecondary).
			Foreground(ColorTextSecondary).
			Padding(0, 1)
	}

	return style.Render(status.String())
}

// TuiStatusCompact renders a compact status indicator (single character).
// Useful for list views where space is limited.
func TuiStatusCompact(status Status) string {
	var icon string
	var color lipgloss.Color

	switch status {
	case StatusPending:
		icon = "○"
		color = ColorTextMuted
	case StatusRunning:
		icon = "●"
		color = ColorCyan
	case StatusSuccess:
		icon = "✓"
		color = ColorGreen
	case StatusFailed:
		icon = "✗"
		color = ColorDestructive
	case StatusError:
		icon = "!"
		color = ColorDestructive
	case StatusSkipped:
		icon = "-"
		color = ColorTextMuted
	default:
		icon = "?"
		color = ColorTextSecondary
	}

	return lipgloss.NewStyle().Foreground(color).Render(icon)
}
