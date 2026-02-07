package build

import (
	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/ui/components"
)

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.phase {
		case phaseSelecting:
			return m.handleSelectionKey(msg)
		case phaseConfiguring:
			return m.handleConfigKey(msg)
		}
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

// handleConfigKey processes keyboard input during build configuration.
func (m Model) handleConfigKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	key := msg.String()

	// Global config keys
	switch key {
	case "esc":
		// Cancel: return to selection phase
		m.phase = phaseSelecting
		return m, nil
	case "tab":
		// Cycle to next section
		m.section = (m.section + 1) % sectionCount
		return m, nil
	case "shift+tab":
		// Cycle to previous section
		m.section = (m.section - 1 + sectionCount) % sectionCount
		return m, nil
	}

	// Section-specific keys
	switch m.section {
	case sectionGoal:
		return m.handleGoalKey(key)
	case sectionProfiles:
		return m.handleProfilesKey(key)
	case sectionPort:
		return m.handlePortKey(key)
	case sectionOptions:
		return m.handleOptionsKey(key)
	case sectionButtons:
		return m.handleButtonsKey(key)
	}

	return m, nil
}

// handleGoalKey handles keys in the Maven Goal section.
func (m Model) handleGoalKey(key string) (Model, tea.Cmd) {
	switch key {
	case "left", "h":
		if m.goalIndex > 0 {
			m.goalIndex--
		}
		m.config.Goal = mavenGoals[m.goalIndex]
	case "right", "l":
		if m.goalIndex < len(mavenGoals)-1 {
			m.goalIndex++
		}
		m.config.Goal = mavenGoals[m.goalIndex]
	case "enter":
		// Confirm and start build
		return m.startBuild()
	}
	return m, nil
}

// handleProfilesKey handles keys in the Maven Profiles section.
func (m Model) handleProfilesKey(key string) (Model, tea.Cmd) {
	switch key {
	case "up", "k":
		if m.profileCursor > 0 {
			m.profileCursor--
		}
	case "down", "j":
		if m.profileCursor < len(profileOptions)-1 {
			m.profileCursor++
		}
	case " ", "enter":
		// Toggle the profile at cursor
		profile := profileOptions[m.profileCursor]
		m.config.Profiles = toggleProfile(m.config.Profiles, profile)
	}
	return m, nil
}

// handlePortKey handles keys in the Traktion Port section.
func (m Model) handlePortKey(key string) (Model, tea.Cmd) {
	switch key {
	case "left", "h":
		if m.portIndex > 0 {
			m.portIndex--
		}
		m.config.Port = portOptions[m.portIndex]
	case "right", "l":
		if m.portIndex < len(portOptions)-1 {
			m.portIndex++
		}
		m.config.Port = portOptions[m.portIndex]
	case "enter":
		return m.startBuild()
	}
	return m, nil
}

// handleOptionsKey handles keys in the Build Options section.
func (m Model) handleOptionsKey(key string) (Model, tea.Cmd) {
	switch key {
	case " ", "enter":
		m.config.SkipTests = !m.config.SkipTests
	}
	return m, nil
}

// handleButtonsKey handles keys in the action buttons section.
func (m Model) handleButtonsKey(key string) (Model, tea.Cmd) {
	switch key {
	case "left", "h":
		m.buttonIndex = 0
	case "right", "l":
		m.buttonIndex = 1
	case "enter":
		if m.buttonIndex == 0 {
			// Cancel
			m.phase = phaseSelecting
			return m, nil
		}
		// Start Build
		return m.startBuild()
	}
	return m, nil
}

// startBuild sends a StartBuildMsg and transitions to executing phase.
func (m Model) startBuild() (Model, tea.Cmd) {
	m.phase = phaseExecuting
	cfg := m.config
	selected := m.selectedComponents
	return m, func() tea.Msg {
		return StartBuildMsg{Config: cfg, Selected: selected}
	}
}

// toggleProfile adds or removes a profile from the list.
// Returns a new slice to avoid mutating the original.
func toggleProfile(profiles []string, profile string) []string {
	var result []string
	found := false
	for _, p := range profiles {
		if p == profile {
			found = true
			continue
		}
		result = append(result, p)
	}
	if !found {
		result = append(result, profile)
	}
	return result
}
