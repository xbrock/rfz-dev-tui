// Package buildconfig provides the Build Configuration modal.
//
// This modal appears before starting a build and allows users to:
// - Select Maven profiles
// - Toggle skip tests option
// - Set parallelism level
// - Choose clean vs incremental build
package buildconfig

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/ui/components"
)

// Model represents the Build Config modal state.
type Model struct {
	// Configuration being edited
	config domain.BuildConfig

	// Focused option
	focused int

	// Window dimensions
	width  int
	height int

	// Visible state
	visible bool
}

// New creates a new Build Config modal.
func New() Model {
	return Model{
		config: domain.BuildConfig{
			Parallelism: 4,
			SkipTests:   false,
			CleanBuild:  true,
			Profiles:    []string{"base"},
		},
	}
}

// Show makes the modal visible and returns it.
func (m Model) Show() Model {
	m.visible = true
	return m
}

// Hide makes the modal invisible and returns it.
func (m Model) Hide() Model {
	m.visible = false
	return m
}

// IsVisible returns whether the modal is visible.
func (m Model) IsVisible() bool {
	return m.visible
}

// Config returns the current configuration.
func (m Model) Config() domain.BuildConfig {
	return m.config
}

// SetSize updates the modal dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// Init initializes the modal.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages for the modal.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.visible {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Cancel and hide
			return m.Hide(), nil

		case "enter":
			// Confirm and hide
			return m.Hide(), func() tea.Msg {
				return BuildConfigConfirmedMsg{Config: m.config}
			}

		case "up", "k":
			m.focused--
			if m.focused < 0 {
				m.focused = 3 // Wrap to last option
			}

		case "down", "j":
			m.focused++
			if m.focused > 3 {
				m.focused = 0 // Wrap to first option
			}

		case " ", "x":
			// Toggle current option
			m.toggleCurrent()

		case "+", "=":
			if m.focused == 2 {
				m.config.Parallelism++
				if m.config.Parallelism > 8 {
					m.config.Parallelism = 8
				}
			}

		case "-":
			if m.focused == 2 {
				m.config.Parallelism--
				if m.config.Parallelism < 1 {
					m.config.Parallelism = 1
				}
			}
		}
	}

	return m, nil
}

// View renders the modal.
func (m Model) View() string {
	if !m.visible {
		return ""
	}

	// Modal title
	title := components.StyleH2.Render("Build Configuration")

	// Options
	options := []string{
		m.renderOption(0, "Skip Tests", m.config.SkipTests),
		m.renderOption(1, "Clean Build", m.config.CleanBuild),
		m.renderParallelism(2),
		m.renderProfiles(3),
	}

	optionsView := lipgloss.JoinVertical(lipgloss.Left, options...)

	// Footer
	footer := components.StyleBodyMuted.Render(
		"Enter: confirm | Esc: cancel | Space: toggle",
	)

	// Combine content
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		optionsView,
		"",
		footer,
	)

	// Modal styling
	modalWidth := 50
	modalHeight := 15

	modalStyle := lipgloss.NewStyle().
		Border(components.BorderDouble).
		BorderForeground(components.ColorCyan).
		Padding(1, 2).
		Width(modalWidth).
		Height(modalHeight)

	modal := modalStyle.Render(content)

	// Center the modal
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		modal,
	)
}

// renderOption renders a toggle option.
func (m Model) renderOption(index int, label string, enabled bool) string {
	icon := "[ ]"
	if enabled {
		icon = "[x]"
	}

	var style lipgloss.Style
	if index == m.focused {
		style = lipgloss.NewStyle().
			Foreground(components.ColorCyan).
			Bold(true)
	} else {
		style = components.StyleBody
	}

	cursor := "  "
	if index == m.focused {
		cursor = "> "
	}

	return cursor + style.Render(icon+" "+label)
}

// renderParallelism renders the parallelism option.
func (m Model) renderParallelism(index int) string {
	var style lipgloss.Style
	if index == m.focused {
		style = lipgloss.NewStyle().
			Foreground(components.ColorCyan).
			Bold(true)
	} else {
		style = components.StyleBody
	}

	cursor := "  "
	if index == m.focused {
		cursor = "> "
	}

	value := components.StyleKeyboard.Render("%d", m.config.Parallelism)
	return cursor + style.Render("Parallelism: ") + value
}

// renderProfiles renders the profiles option.
func (m Model) renderProfiles(index int) string {
	var style lipgloss.Style
	if index == m.focused {
		style = lipgloss.NewStyle().
			Foreground(components.ColorCyan).
			Bold(true)
	} else {
		style = components.StyleBody
	}

	cursor := "  "
	if index == m.focused {
		cursor = "> "
	}

	profiles := ""
	for i, p := range m.config.Profiles {
		if i > 0 {
			profiles += ", "
		}
		profiles += p
	}

	return cursor + style.Render("Profiles: "+profiles)
}

// toggleCurrent toggles the current option.
func (m *Model) toggleCurrent() {
	switch m.focused {
	case 0:
		m.config.SkipTests = !m.config.SkipTests
	case 1:
		m.config.CleanBuild = !m.config.CleanBuild
	}
}

// Messages

// BuildConfigConfirmedMsg is sent when the user confirms the configuration.
type BuildConfigConfirmedMsg struct {
	Config domain.BuildConfig
}

// BuildConfigCancelledMsg is sent when the user cancels the modal.
type BuildConfigCancelledMsg struct{}
