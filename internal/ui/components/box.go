// Package components provides shared UI components and styles.
//
// This file contains the TuiBox bordered container component.
package components

import (
	"github.com/charmbracelet/lipgloss"
)

// BoxStyle defines the border style for TuiBox.
type BoxStyle string

const (
	// BoxSingle uses single-line borders (┌─┐).
	BoxSingle BoxStyle = "single"

	// BoxDouble uses double-line borders (╔═╗).
	BoxDouble BoxStyle = "double"

	// BoxRounded uses rounded corner borders (╭─╮).
	BoxRounded BoxStyle = "rounded"

	// BoxHeavy uses thick/heavy borders.
	BoxHeavy BoxStyle = "heavy"
)

// getBorder returns the lipgloss.Border for the given BoxStyle.
func getBorder(style BoxStyle) lipgloss.Border {
	switch style {
	case BoxDouble:
		return BorderDouble
	case BoxRounded:
		return BorderRounded
	case BoxHeavy:
		return BorderHeavy
	default:
		return BorderSingle
	}
}

// TuiBox renders content inside a bordered container.
// The focused parameter highlights the border in cyan when true.
func TuiBox(content string, style BoxStyle, focused bool) string {
	border := getBorder(style)

	borderColor := ColorBorder
	if focused {
		borderColor = ColorCyan
	}

	boxStyle := lipgloss.NewStyle().
		Border(border).
		BorderForeground(borderColor).
		Padding(0, 1)

	return boxStyle.Render(content)
}

// TuiBoxWithWidth renders content inside a bordered container with a fixed width.
// Content exceeding the width is truncated with an ellipsis.
func TuiBoxWithWidth(content string, style BoxStyle, focused bool, width int) string {
	border := getBorder(style)

	borderColor := ColorBorder
	if focused {
		borderColor = ColorCyan
	}

	// Account for border (2 chars) and padding (2 chars)
	innerWidth := width - 4
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Truncate content if needed
	truncated := Truncate(content, innerWidth)

	boxStyle := lipgloss.NewStyle().
		Border(border).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(width)

	return boxStyle.Render(truncated)
}
