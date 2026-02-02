// Package build provides the Build screen model.
//
// The Build screen allows users to:
// - View discovered components
// - Select components for building
// - Configure build options
// - Start and monitor builds
// - View build status for each component
package build

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/service"
	"rfz-cli/internal/ui/components"
)

// Model represents the Build screen state.
type Model struct {
	// Bubbles list component for component selection
	// CRITICAL: Use Bubbles list, not custom implementation
	list list.Model

	// Selected components (component ID -> selected)
	selected map[string]bool

	// Build state
	building    bool
	buildStatus map[string]components.Status

	// Components data
	components []domain.Component

	// Window dimensions
	width  int
	height int

	// Build service (injected)
	buildSvc *service.BuildService
}

// componentItem implements list.Item for the Bubbles list.
type componentItem struct {
	component domain.Component
	status    components.Status
}

func (i componentItem) Title() string       { return i.component.Name }
func (i componentItem) Description() string { return i.component.Path }
func (i componentItem) FilterValue() string { return i.component.Name }

// New creates a new Build screen model.
func New(buildSvc *service.BuildService) Model {
	// Create a default list with empty items
	// Items will be populated when components are loaded
	delegate := list.NewDefaultDelegate()
	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = "Build Components"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

	// Apply RFZ styles to the list
	// CRITICAL: Use Lip Gloss styles, not custom rendering
	l.Styles.Title = components.StyleH2

	return Model{
		list:        l,
		selected:    make(map[string]bool),
		buildStatus: make(map[string]components.Status),
		buildSvc:    buildSvc,
	}
}

// SetSize updates the model dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height

	// Update list dimensions
	// Account for borders and padding
	listWidth := width - 4
	listHeight := height - 4
	m.list.SetSize(listWidth, listHeight)

	return m
}

// Init initializes the Build screen.
func (m Model) Init() tea.Cmd {
	// Return a command to load components
	return m.loadComponentsCmd()
}

// Update handles messages for the Build screen.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width-4, msg.Height-4)

	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			// Toggle selection of current item
			if item, ok := m.list.SelectedItem().(componentItem); ok {
				m.selected[item.component.ID] = !m.selected[item.component.ID]
			}
			return m, nil

		case "b", "enter":
			// Start build if not already building
			if !m.building && len(m.selected) > 0 {
				m.building = true
				return m, m.startBuildCmd()
			}
			return m, nil
		}

	case componentsLoadedMsg:
		// Update components and refresh list
		m.components = msg.Components
		m.updateListItems()

	case buildStartedMsg:
		// Update status for building components
		for id := range m.selected {
			m.buildStatus[id] = components.StatusRunning
		}
		m.updateListItems()

	case buildProgressMsg:
		// Update individual component status
		m.buildStatus[msg.ComponentID] = msg.Status
		m.updateListItems()

	case buildCompleteMsg:
		// Build finished
		m.building = false
		for id := range m.selected {
			if msg.Success {
				m.buildStatus[id] = components.StatusSuccess
			} else {
				m.buildStatus[id] = components.StatusFailed
			}
		}
		m.updateListItems()
	}

	// Update the list component
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the Build screen.
func (m Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// The list component handles most of the rendering
	// CRITICAL: Use Bubbles list view, not custom rendering
	listView := m.list.View()

	// Add selection summary
	summary := m.renderSelectionSummary()

	// Combine list and summary
	content := lipgloss.JoinVertical(lipgloss.Left, listView, summary)

	// Apply container styling
	return components.StyleContent.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// updateListItems refreshes the list items with current status.
func (m *Model) updateListItems() {
	items := make([]list.Item, len(m.components))
	for i, comp := range m.components {
		status := m.buildStatus[comp.ID]
		if status == 0 {
			status = components.StatusPending
		}
		items[i] = componentItem{
			component: comp,
			status:    status,
		}
	}
	m.list.SetItems(items)
}

// renderSelectionSummary renders the selection count and build button.
func (m Model) renderSelectionSummary() string {
	count := 0
	for _, selected := range m.selected {
		if selected {
			count++
		}
	}

	if count == 0 {
		return components.StyleBodyMuted.Render("No components selected")
	}

	text := components.StyleBody.Render(
		lipgloss.NewStyle().Render("Selected: ") +
			components.StyleKeyboard.Render("%d component(s)"),
	)

	if m.building {
		return text + "  " + components.TuiStatus(components.StatusRunning)
	}

	return text + "  " + components.StyleButtonPrimary.Render("Press Enter to build")
}

// Commands

func (m Model) loadComponentsCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Use buildSvc to load components
		// For now, return empty list
		return componentsLoadedMsg{Components: []domain.Component{}}
	}
}

func (m Model) startBuildCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Use buildSvc to start build
		return buildStartedMsg{}
	}
}

// Messages

type componentsLoadedMsg struct {
	Components []domain.Component
}

type buildStartedMsg struct{}

type buildProgressMsg struct {
	ComponentID string
	Status      components.Status
}

type buildCompleteMsg struct {
	Success bool
}
