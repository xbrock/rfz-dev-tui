// Package build provides the Build Component Selection screen for the RFZ CLI.
package build

import (
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

// OpenConfigMsg is sent when the user confirms a component selection to open the config modal.
type OpenConfigMsg struct {
	Selected []string
}

// Model represents the Build Component Selection screen state.
type Model struct {
	width  int
	height int
	phase  buildPhase

	items       []components.TuiListItem
	cursorIndex int
	focused     bool

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

// SetFocused sets the focus state of the build screen.
func (m Model) SetFocused(focused bool) Model {
	m.focused = focused
	return m
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// CurrentItemLabel returns the label of the currently focused item.
func (m Model) CurrentItemLabel() string {
	if m.cursorIndex >= 0 && m.cursorIndex < len(m.items) {
		return m.items[m.cursorIndex].Label
	}
	return ""
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
