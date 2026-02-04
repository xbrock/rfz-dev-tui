// Package components provides shared UI components and styles.
//
// This file contains the TuiProgress progress bar component.
package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Default progress bar width
const defaultProgressWidth = 40

// TuiProgressModel is a styled progress bar wrapping bubbles/progress.
type TuiProgressModel struct {
	progress    progress.Model
	percent     float64
	showPercent bool
}

// NewTuiProgress creates a new progress bar with RFZ styling.
// width: the width of the progress bar in characters
// showPercent: whether to display the percentage value next to the bar
func NewTuiProgress(width int, showPercent bool) TuiProgressModel {
	if width <= 0 {
		width = defaultProgressWidth
	}

	// Create progress bar with yellow-to-green gradient
	p := progress.New(
		progress.WithGradient(string(ColorYellow), string(ColorGreen)),
		progress.WithWidth(width),
		progress.WithoutPercentage(), // We handle percentage display ourselves
	)

	return TuiProgressModel{
		progress:    p,
		percent:     0,
		showPercent: showPercent,
	}
}

// Init implements tea.Model.
func (m TuiProgressModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m TuiProgressModel) Update(msg tea.Msg) (TuiProgressModel, tea.Cmd) {
	var cmd tea.Cmd
	progressModel, cmd := m.progress.Update(msg)
	m.progress = progressModel.(progress.Model)
	return m, cmd
}

// View implements tea.Model.
func (m TuiProgressModel) View() string {
	bar := m.progress.ViewAs(m.percent)

	if m.showPercent {
		percentStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
		percentText := percentStyle.Render(fmt.Sprintf(" %3.0f%%", m.percent*100))
		return bar + percentText
	}

	return bar
}

// SetPercent sets the current progress (0.0 to 1.0).
func (m *TuiProgressModel) SetPercent(percent float64) {
	if percent < 0 {
		percent = 0
	}
	if percent > 1 {
		percent = 1
	}
	m.percent = percent
}

// Percent returns the current progress (0.0 to 1.0).
func (m TuiProgressModel) Percent() float64 {
	return m.percent
}

// SetWidth sets the width of the progress bar.
func (m *TuiProgressModel) SetWidth(w int) {
	m.progress.Width = w
}

// SetShowPercent enables or disables percentage display.
func (m *TuiProgressModel) SetShowPercent(show bool) {
	m.showPercent = show
}

// TuiProgress renders a static progress bar for gallery/documentation display.
// percent: progress value from 0.0 to 1.0
// width: the width of the progress bar in characters
// showPercent: whether to display the percentage value
func TuiProgress(percent float64, width int, showPercent bool) string {
	if width <= 0 {
		width = defaultProgressWidth
	}
	if percent < 0 {
		percent = 0
	}
	if percent > 1 {
		percent = 1
	}

	// Create progress bar with yellow-to-green gradient
	p := progress.New(
		progress.WithGradient(string(ColorYellow), string(ColorGreen)),
		progress.WithWidth(width),
		progress.WithoutPercentage(),
	)

	bar := p.ViewAs(percent)

	if showPercent {
		percentStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary)
		percentText := percentStyle.Render(fmt.Sprintf(" %3.0f%%", percent*100))
		return bar + percentText
	}

	return bar
}
