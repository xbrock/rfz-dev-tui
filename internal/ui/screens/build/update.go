package build

import (
	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/ui/components"
)

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.phase != phaseSelecting {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleSelectionKey(msg)
	}

	return m, nil
}

// handleSelectionKey processes keyboard input during component selection.
func (m Model) handleSelectionKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if len(m.items) > 0 {
			m.cursorIndex = (m.cursorIndex - 1 + len(m.items)) % len(m.items)
		}

	case "down", "j":
		if len(m.items) > 0 {
			m.cursorIndex = (m.cursorIndex + 1) % len(m.items)
		}

	case " ":
		m.items = components.ToggleSelection(m.items, m.cursorIndex, components.ListMultiSelect)

	case "a":
		m.items = components.SelectAll(m.items)

	case "n":
		m.items = components.DeselectAll(m.items)

	case "enter":
		selected := components.GetSelectedLabels(m.items)
		if len(selected) > 0 {
			return m, func() tea.Msg {
				return OpenConfigMsg{Selected: selected}
			}
		}
	}

	return m, nil
}
