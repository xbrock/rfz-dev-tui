// Package config provides the Configuration screen model.
//
// The Config screen allows users to:
// - Set Maven options (profiles, skip tests, etc.)
// - Configure workspace paths
// - Set default build options
// - Manage component registry
package config

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/service"
	"rfz-cli/internal/ui/components"
)

// FocusedField tracks which input field is focused.
type FocusedField int

const (
	FieldMavenPath FocusedField = iota
	FieldWorkspace
	FieldProfiles
	FieldCount // Used to cycle through fields
)

// Model represents the Config screen state.
type Model struct {
	// Text inputs for configuration values
	// CRITICAL: Use Bubbles textinput, not custom input handling
	mavenPathInput textinput.Model
	workspaceInput textinput.Model
	profilesInput  textinput.Model

	// Currently focused field
	focused FocusedField

	// Configuration state
	skipTests    bool
	parallelism  int
	cleanBuild   bool

	// Window dimensions
	width  int
	height int

	// Config service (injected)
	configSvc *service.ConfigService
}

// New creates a new Config screen model.
func New(configSvc *service.ConfigService) Model {
	// Create text inputs with RFZ styling
	mavenPath := textinput.New()
	mavenPath.Placeholder = "/usr/bin/mvn"
	mavenPath.CharLimit = 256
	mavenPath.Width = 40

	workspace := textinput.New()
	workspace.Placeholder = "~/workspace"
	workspace.CharLimit = 256
	workspace.Width = 40

	profiles := textinput.New()
	profiles.Placeholder = "base,local"
	profiles.CharLimit = 256
	profiles.Width = 40

	// Focus first input
	mavenPath.Focus()

	return Model{
		mavenPathInput: mavenPath,
		workspaceInput: workspace,
		profilesInput:  profiles,
		focused:        FieldMavenPath,
		parallelism:    4,
		configSvc:      configSvc,
	}
}

// SetSize updates the model dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// Init initializes the Config screen.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages for the Config screen.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down":
			// Move to next field
			m.focusNext()
			return m, nil

		case "shift+tab", "up":
			// Move to previous field
			m.focusPrev()
			return m, nil

		case "t":
			// Toggle skip tests
			m.skipTests = !m.skipTests
			return m, nil

		case "c":
			// Toggle clean build
			m.cleanBuild = !m.cleanBuild
			return m, nil

		case "ctrl+s":
			// Save configuration
			return m, m.saveConfigCmd()
		}
	}

	// Update the focused input
	switch m.focused {
	case FieldMavenPath:
		m.mavenPathInput, cmd = m.mavenPathInput.Update(msg)
	case FieldWorkspace:
		m.workspaceInput, cmd = m.workspaceInput.Update(msg)
	case FieldProfiles:
		m.profilesInput, cmd = m.profilesInput.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the Config screen.
func (m Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Title
	title := components.StyleH2.Render("Configuration")

	// Input fields
	mavenField := m.renderField("Maven Path", m.mavenPathInput.View(), m.focused == FieldMavenPath)
	workspaceField := m.renderField("Workspace", m.workspaceInput.View(), m.focused == FieldWorkspace)
	profilesField := m.renderField("Profiles", m.profilesInput.View(), m.focused == FieldProfiles)

	// Toggle options
	toggles := m.renderToggles()

	// Footer hints
	footer := m.renderFooter()

	// Combine all sections
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		mavenField,
		workspaceField,
		profilesField,
		"",
		toggles,
		"",
		footer,
	)

	return components.StyleContent.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// renderField renders a labeled input field.
func (m Model) renderField(label string, input string, focused bool) string {
	labelStyle := components.StyleBodySecondary
	if focused {
		labelStyle = components.StyleH3
	}

	labelText := labelStyle.Render(label + ":")

	var inputStyle lipgloss.Style
	if focused {
		inputStyle = components.StyleInputFocused
	} else {
		inputStyle = components.StyleInputNormal
	}

	inputText := inputStyle.Render(input)

	return lipgloss.JoinVertical(lipgloss.Left, labelText, inputText)
}

// renderToggles renders the toggle options.
func (m Model) renderToggles() string {
	skipTestsToggle := m.renderToggle("Skip Tests", m.skipTests, "t")
	cleanBuildToggle := m.renderToggle("Clean Build", m.cleanBuild, "c")

	return lipgloss.JoinVertical(lipgloss.Left, skipTestsToggle, cleanBuildToggle)
}

// renderToggle renders a single toggle option.
func (m Model) renderToggle(label string, enabled bool, key string) string {
	icon := "[ ]"
	if enabled {
		icon = "[x]"
	}

	iconStyle := lipgloss.NewStyle().Foreground(components.ColorCyan)
	if enabled {
		iconStyle = iconStyle.Bold(true)
	}

	shortcut := components.StyleKeyboard.Render(key)

	return iconStyle.Render(icon) + " " +
		components.StyleBody.Render(label) + "  " +
		shortcut
}

// renderFooter renders the keyboard hints.
func (m Model) renderFooter() string {
	return components.StyleBodyMuted.Render(
		"Tab: next field | Ctrl+S: save",
	)
}

// focusNext moves focus to the next field.
func (m *Model) focusNext() {
	m.blurAll()
	m.focused = (m.focused + 1) % FieldCount
	m.focusCurrent()
}

// focusPrev moves focus to the previous field.
func (m *Model) focusPrev() {
	m.blurAll()
	m.focused = (m.focused - 1 + FieldCount) % FieldCount
	m.focusCurrent()
}

// blurAll removes focus from all inputs.
func (m *Model) blurAll() {
	m.mavenPathInput.Blur()
	m.workspaceInput.Blur()
	m.profilesInput.Blur()
}

// focusCurrent focuses the current field.
func (m *Model) focusCurrent() {
	switch m.focused {
	case FieldMavenPath:
		m.mavenPathInput.Focus()
	case FieldWorkspace:
		m.workspaceInput.Focus()
	case FieldProfiles:
		m.profilesInput.Focus()
	}
}

// saveConfigCmd returns a command to save the configuration.
func (m Model) saveConfigCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Use configSvc to save configuration
		return configSavedMsg{}
	}
}

// Messages

type configSavedMsg struct{}

type configLoadedMsg struct {
	MavenPath string
	Workspace string
	Profiles  string
	SkipTests bool
	CleanBuild bool
}
