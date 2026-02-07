package build

// View renders the Build screen based on current phase.
func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	switch m.phase {
	case phaseSelecting:
		return m.viewSelection()
	default:
		return m.viewSelection()
	}
}
