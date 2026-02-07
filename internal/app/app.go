// Package app provides the main RFZ CLI application model.
package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the top-level Bubble Tea model for the RFZ CLI application.
type Model struct {
	width  int
	height int
}

// New creates a new application model.
func New() Model {
	return Model{}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

// View implements tea.Model.
func (m Model) View() string {
	return "RFZ Developer CLI"
}
