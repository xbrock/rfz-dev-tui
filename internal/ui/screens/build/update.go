package build

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/ui/components"
)

// maxConcurrentBuilds limits how many components build simultaneously.
const maxConcurrentBuilds = 3

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.phase {
		case phaseSelecting:
			return m.handleSelectionKey(msg)
		case phaseConfiguring:
			return m.handleConfigKey(msg)
		case phaseExecuting:
			return m.handleExecutionKey(msg)
		case phaseCompleted:
			return m.handleCompletedKey(msg)
		}

	case StartBuildMsg:
		return m.handleStartBuild(msg)

	case BuildTickMsg:
		return m.handleBuildTick(time.Time(msg))

	case BuildPhaseMsg:
		return m.handleBuildPhase(msg)

	case BuildCompleteMsg:
		m.phase = phaseCompleted
		return m, nil
	}

	return m, nil
}

// handleStartBuild initializes the build execution state and starts the simulator.
func (m Model) handleStartBuild(msg StartBuildMsg) (Model, tea.Cmd) {
	m.phase = phaseExecuting
	m.config = msg.Config
	m.selectedComponents = msg.Selected
	m = m.initBuildStates()

	simStates, tickCmd := startSimulator(len(m.buildStates))
	m.simStates = simStates

	// Start initial components
	startMsgs := startPendingComponents(m.buildStates, maxConcurrentBuilds)
	for _, pm := range startMsgs {
		m.buildStates[pm.ComponentIndex].Phase = pm.Phase
		m.buildStates[pm.ComponentIndex].StartTime = time.Now()
	}

	return m, tickCmd
}

// handleBuildTick processes a simulator tick: updates elapsed times and advances phases.
func (m Model) handleBuildTick(_ time.Time) (Model, tea.Cmd) {
	if m.phase != phaseExecuting {
		return m, nil
	}

	now := time.Now()

	// Update elapsed time for active components
	for i := range m.buildStates {
		phase := m.buildStates[i].Phase
		if phase != domain.PhasePending && phase != domain.PhaseDone && phase != domain.PhaseFailed {
			m.buildStates[i].Elapsed = now.Sub(m.buildStates[i].StartTime)
		}
	}

	// Advance simulation
	phaseMsgs, allDone := advanceSimulation(m.buildStates, m.simStates)

	// Apply phase transitions
	for _, pm := range phaseMsgs {
		m.buildStates[pm.ComponentIndex].Phase = pm.Phase
		if pm.Phase == domain.PhaseDone || pm.Phase == domain.PhaseFailed {
			m.buildStates[pm.ComponentIndex].Progress = 1
		} else {
			m.buildStates[pm.ComponentIndex].Progress = 0
		}
	}

	// Start pending components if slots are available
	startMsgs := startPendingComponents(m.buildStates, maxConcurrentBuilds)
	for _, pm := range startMsgs {
		m.buildStates[pm.ComponentIndex].Phase = pm.Phase
		m.buildStates[pm.ComponentIndex].StartTime = now
	}

	if allDone && m.buildDone() {
		m.phase = phaseCompleted
		return m, nil
	}

	return m, simulatorTick()
}

// handleBuildPhase processes a single phase transition message.
func (m Model) handleBuildPhase(msg BuildPhaseMsg) (Model, tea.Cmd) {
	if msg.ComponentIndex >= 0 && msg.ComponentIndex < len(m.buildStates) {
		m.buildStates[msg.ComponentIndex].Phase = msg.Phase
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

// startBuild sends a StartBuildMsg to trigger the execution phase.
func (m Model) startBuild() (Model, tea.Cmd) {
	cfg := m.config
	selected := m.selectedComponents
	return m, func() tea.Msg {
		return StartBuildMsg{Config: cfg, Selected: selected}
	}
}

// handleExecutionKey processes keyboard input during build execution.
func (m Model) handleExecutionKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.buildCursor > 0 {
			m.buildCursor--
		}
	case "down", "j":
		if m.buildCursor < len(m.buildStates)-1 {
			m.buildCursor++
		}
	case "esc":
		// Cancel build: mark all running/pending as failed
		m.buildCanceled = true
		for i := range m.buildStates {
			phase := m.buildStates[i].Phase
			if phase != domain.PhaseDone && phase != domain.PhaseFailed {
				m.buildStates[i].Phase = domain.PhaseFailed
				m.buildStates[i].Progress = 0
			}
		}
		m.phase = phaseCompleted
		return m, nil
	}
	return m, nil
}

// handleCompletedKey processes keyboard input after build completion.
func (m Model) handleCompletedKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.buildCursor > 0 {
			m.buildCursor--
		}
	case "down", "j":
		if m.buildCursor < len(m.buildStates)-1 {
			m.buildCursor++
		}
	case "n":
		// New build: reset to selection phase
		m.phase = phaseSelecting
		m.buildStates = nil
		m.simStates = nil
		m.buildCursor = 0
		m.buildCanceled = false
		// Deselect all items
		m.items = components.DeselectAll(m.items)
	}
	return m, nil
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
