// Package placeholder provides a generic placeholder screen for unimplemented features.
package placeholder

import (
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// Model represents a placeholder screen state.
type Model struct {
	Title  string
	width  int
	height int
}

// New creates a new placeholder screen with the given title.
func New(title string, width, height int) Model {
	return Model{
		Title:  title,
		width:  width,
		height: height,
	}
}

// SetSize updates the screen dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// View renders the placeholder screen.
func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	title := components.StyleH2.Render(m.Title + " - Coming Soon")
	hint := components.StyleBodyMuted.Render("Press Esc to return to Welcome")

	content := lipgloss.JoinVertical(lipgloss.Center, title, "", hint)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
