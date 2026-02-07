// Package build provides the Build Component Selection screen for the RFZ CLI.
package build

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/ui/components"
)

// buildPhase tracks the current phase of the build workflow.
type buildPhase int

const (
	phaseSelecting   buildPhase = iota // Selecting components to build
	phaseConfiguring                   // Configuring build options (modal)
	phaseExecuting                     // Build in progress
	phaseCompleted                     // Build finished
)

// configSection identifies which section of the config form is focused.
type configSection int

const (
	sectionGoal     configSection = iota // Maven Goal radio group
	sectionProfiles                      // Maven Profiles checkboxes
	sectionPort                          // Traktion Port radio group
	sectionOptions                       // Build Options (skip tests)
	sectionButtons                       // Action buttons (Cancel / Start Build)
	sectionCount                         // Sentinel: total number of sections
)

// OpenConfigMsg is sent when the user confirms a component selection to open the config modal.
type OpenConfigMsg struct {
	Selected []string
}

// StartBuildMsg is sent when the user confirms the build configuration.
type StartBuildMsg struct {
	Config   domain.BuildConfig
	Selected []string
}

// BuildTickMsg is sent periodically to update elapsed time during execution.
type BuildTickMsg time.Time

// BuildPhaseMsg signals a phase transition for a specific component.
type BuildPhaseMsg struct {
	ComponentIndex int
	Phase          domain.BuildPhase
}

// BuildCompleteMsg signals that all components have finished building.
type BuildCompleteMsg struct{}

// componentBuildState tracks build progress for a single component.
type componentBuildState struct {
	Name      string
	Phase     domain.BuildPhase
	StartTime time.Time
	Elapsed   time.Duration
	Progress  float64 // 0.0 to 1.0 within current phase
}

// mavenGoals defines the available Maven goals in display order.
var mavenGoals = []domain.MavenGoal{
	domain.GoalClean,
	domain.GoalInstall,
	domain.GoalPackage,
	domain.GoalCleanInstall,
}

// mavenGoalLabels maps goals to display labels.
var mavenGoalLabels = map[domain.MavenGoal]string{
	domain.GoalClean:        "clean",
	domain.GoalInstall:      "install",
	domain.GoalPackage:      "package",
	domain.GoalCleanInstall: "clean install",
}

// profileOptions defines the available Maven profiles.
var profileOptions = []string{"target_env_dev", "generate_local_config_files"}

// portOptions defines the available Traktion ports.
var portOptions = []int{11090, 11091}

// Model represents the Build Component Selection screen state.
type Model struct {
	width    int
	height   int
	termW    int // Full terminal width (for modal overlay)
	termH    int // Full terminal height (for modal overlay)
	phase    buildPhase

	// Selection phase
	items       []components.TuiListItem
	cursorIndex int
	focused     bool

	// Config phase
	selectedComponents []string
	config             domain.BuildConfig
	section            configSection
	goalIndex          int // Cursor within goal radio group
	profileCursor      int // Cursor within profiles list
	portIndex          int // Cursor within port radio group
	buttonIndex        int // 0=Cancel, 1=Start Build

	// Execution phase
	buildStates   []componentBuildState
	simStates     []simulatorState
	buildCursor   int  // Cursor in execution table
	buildCanceled bool // True if user canceled the build

	provider domain.ComponentProvider
}

// New creates a new Build Screen model.
func New(width, height int) Model {
	provider := domain.MockComponentProvider{}
	items := componentsToListItems(provider.Components())

	return Model{
		width:       width,
		height:      height,
		phase:       phaseSelecting,
		items:       items,
		cursorIndex: 0,
		focused:     true,
		provider:    provider,
	}
}

// SetSize updates the screen dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// SetTermSize stores the full terminal dimensions for modal overlay rendering.
func (m Model) SetTermSize(w, h int) Model {
	m.termW = w
	m.termH = h
	return m
}

// SetFocused sets the focus state of the build screen.
func (m Model) SetFocused(focused bool) Model {
	m.focused = focused
	return m
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// IsConfiguring returns true when the build screen is in the configuration phase.
func (m Model) IsConfiguring() bool {
	return m.phase == phaseConfiguring
}

// CurrentItemLabel returns the label of the currently focused item.
func (m Model) CurrentItemLabel() string {
	if m.cursorIndex >= 0 && m.cursorIndex < len(m.items) {
		return m.items[m.cursorIndex].Label
	}
	return ""
}

// IsExecuting returns true when a build is running.
func (m Model) IsExecuting() bool {
	return m.phase == phaseExecuting
}

// IsCompleted returns true when the build has finished.
func (m Model) IsCompleted() bool {
	return m.phase == phaseCompleted
}

// buildDone returns true when all components have finished (Done or Failed).
func (m Model) buildDone() bool {
	for _, s := range m.buildStates {
		if s.Phase != domain.PhaseDone && s.Phase != domain.PhaseFailed {
			return false
		}
	}
	return len(m.buildStates) > 0
}

// statusCounts returns running, success, failed, pending counts.
func (m Model) statusCounts() (running, success, failed, pending int) {
	for _, s := range m.buildStates {
		switch s.Phase {
		case domain.PhasePending:
			pending++
		case domain.PhaseDone:
			success++
		case domain.PhaseFailed:
			failed++
		default:
			running++
		}
	}
	return
}

// overallProgress returns the fraction of completed components (0.0 to 1.0).
func (m Model) overallProgress() float64 {
	if len(m.buildStates) == 0 {
		return 0
	}
	done := 0
	for _, s := range m.buildStates {
		if s.Phase == domain.PhaseDone || s.Phase == domain.PhaseFailed {
			done++
		}
	}
	return float64(done) / float64(len(m.buildStates))
}

// initBuildStates initializes build state tracking for the selected components.
func (m Model) initBuildStates() Model {
	states := make([]componentBuildState, len(m.selectedComponents))
	now := time.Now()
	for i, name := range m.selectedComponents {
		states[i] = componentBuildState{
			Name:      name,
			Phase:     domain.PhasePending,
			StartTime: now,
		}
	}
	m.buildStates = states
	m.buildCursor = 0
	m.buildCanceled = false
	return m
}

// componentsToListItems converts domain components to TuiListItems with type badges.
func componentsToListItems(comps []domain.Component) []components.TuiListItem {
	items := make([]components.TuiListItem, len(comps))
	for i, c := range comps {
		items[i] = components.TuiListItem{
			Label:    c.Name,
			Badge:    c.Type.String(),
			Selected: false,
		}
	}
	return items
}

// defaultConfig returns the default build configuration.
func defaultConfig() domain.BuildConfig {
	return domain.BuildConfig{
		Goal:      domain.GoalCleanInstall,
		Profiles:  []string{"target_env_dev"},
		Port:      11090,
		SkipTests: true,
	}
}

// OpenConfig transitions from selection to configuration phase.
func (m Model) OpenConfig(selected []string) Model {
	m.phase = phaseConfiguring
	m.selectedComponents = selected
	m.config = defaultConfig()
	m.section = sectionGoal
	m.goalIndex = indexOf(mavenGoals, m.config.Goal)
	m.profileCursor = 0
	m.portIndex = indexOfInt(portOptions, m.config.Port)
	m.buttonIndex = 1 // Default focus on Start Build
	return m
}

// indexOf returns the index of val in slice, or 0 if not found.
func indexOf(slice []domain.MavenGoal, val domain.MavenGoal) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return 0
}

// indexOfInt returns the index of val in slice, or 0 if not found.
func indexOfInt(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return 0
}
