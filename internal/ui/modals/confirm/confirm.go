// Package confirm provides a generic confirmation modal.
//
// This modal is used for confirming destructive or important actions.
// It displays a message and Yes/No buttons.
package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// Model represents the Confirm modal state.
type Model struct {
	// Modal content
	title   string
	message string

	// Button state
	yesSelected bool

	// Window dimensions
	width  int
	height int

	// Visible state
	visible bool

	// Callback identifier
	actionID string
}

// New creates a new Confirm modal.
func New() Model {
	return Model{
		yesSelected: false,
	}
}

// Show displays the modal with the given content.
func (m Model) Show(title, message, actionID string) Model {
	m.title = title
	m.message = message
	m.actionID = actionID
	m.visible = true
	m.yesSelected = false // Default to No for safety
	return m
}

// Hide makes the modal invisible.
func (m Model) Hide() Model {
	m.visible = false
	return m
}

// IsVisible returns whether the modal is visible.
func (m Model) IsVisible() bool {
	return m.visible
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
		case "esc", "n":
			// Cancel / No
			return m.Hide(), func() tea.Msg {
				return ConfirmCancelledMsg{ActionID: m.actionID}
			}

		case "enter":
			if m.yesSelected {
				// Confirm
				return m.Hide(), func() tea.Msg {
					return ConfirmConfirmedMsg{ActionID: m.actionID}
				}
			}
			// No selected
			return m.Hide(), func() tea.Msg {
				return ConfirmCancelledMsg{ActionID: m.actionID}
			}

		case "y":
			// Quick confirm
			return m.Hide(), func() tea.Msg {
				return ConfirmConfirmedMsg{ActionID: m.actionID}
			}

		case "left", "h":
			m.yesSelected = true

		case "right", "l":
			m.yesSelected = false

		case "tab":
			m.yesSelected = !m.yesSelected
		}
	}

	return m, nil
}

// View renders the modal.
func (m Model) View() string {
	if !m.visible {
		return ""
	}

	// Title
	title := components.StyleH2.Render(m.title)

	// Message
	message := components.StyleBody.Render(m.message)

	// Buttons
	buttons := m.renderButtons()

	// Footer hint
	footer := components.StyleBodyMuted.Render(
		"y: Yes | n: No | Tab: switch | Enter: confirm",
	)

	// Combine content
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		message,
		"",
		buttons,
		"",
		footer,
	)

	// Modal styling
	modalWidth := 50
	modalHeight := 12

	modalStyle := lipgloss.NewStyle().
		Border(components.BorderDouble).
		BorderForeground(components.ColorYellow).
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

// renderButtons renders the Yes/No buttons.
func (m Model) renderButtons() string {
	var yesStyle, noStyle lipgloss.Style

	if m.yesSelected {
		yesStyle = components.StyleButtonDestructive
		noStyle = components.StyleButtonSecondary
	} else {
		yesStyle = components.StyleButtonSecondary
		noStyle = components.StyleButtonPrimary
	}

	yes := yesStyle.Render(" Yes ")
	no := noStyle.Render(" No ")

	return lipgloss.JoinHorizontal(lipgloss.Center, yes, "   ", no)
}

// Messages

// ConfirmConfirmedMsg is sent when the user confirms the action.
type ConfirmConfirmedMsg struct {
	ActionID string
}

// ConfirmCancelledMsg is sent when the user cancels the action.
type ConfirmCancelledMsg struct {
	ActionID string
}
