package build

// View renders the Build screen based on current phase.
func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	switch m.phase {
	case phaseSelecting:
		return m.viewSelection()
	case phaseConfiguring:
		return m.viewConfig()
	case phaseExecuting, phaseCompleted:
		return m.viewExecution()
	default:
		return m.viewSelection()
	}
}
