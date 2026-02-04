// Package components provides shared UI components and styles.
//
// This file contains the TuiTextInput text entry component.
package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TuiTextInputModel is a styled text input field wrapping bubbles/textinput.
type TuiTextInputModel struct {
	input    textinput.Model
	prompt   string
	disabled bool
}

// NewTuiTextInput creates a new text input with RFZ styling.
// placeholder: text shown when input is empty
// prompt: optional prefix symbol (e.g., "$", ">") shown before input area
func NewTuiTextInput(placeholder string, prompt string) TuiTextInputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(ColorTextMuted)
	ti.TextStyle = lipgloss.NewStyle().Foreground(ColorTextPrimary)
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(ColorCyan)
	// Disable the default prompt from bubbles/textinput (we use our own)
	ti.Prompt = ""

	return TuiTextInputModel{
		input:    ti,
		prompt:   prompt,
		disabled: false,
	}
}

// Init implements tea.Model.
func (m TuiTextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model.
func (m TuiTextInputModel) Update(msg tea.Msg) (TuiTextInputModel, tea.Cmd) {
	if m.disabled {
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m TuiTextInputModel) View() string {
	// Build the prompt prefix
	var promptStr string
	if m.prompt != "" {
		promptStyle := lipgloss.NewStyle().Foreground(ColorYellow).Bold(true)
		promptStr = promptStyle.Render(m.prompt) + " "
	}

	// Choose border style based on state
	var borderStyle lipgloss.Style
	switch {
	case m.disabled:
		borderStyle = lipgloss.NewStyle().
			Border(BorderSingle).
			BorderForeground(ColorTextDisabled).
			Foreground(ColorTextDisabled).
			Padding(0, 1)
	case m.input.Focused():
		borderStyle = StyleInputFocused
	default:
		borderStyle = StyleInputNormal
	}

	// Render the input field
	inputView := m.input.View()

	// Combine prompt and input, then apply border
	content := promptStr + inputView
	return borderStyle.Render(content)
}

// Focus sets focus on the input field.
func (m *TuiTextInputModel) Focus() tea.Cmd {
	return m.input.Focus()
}

// Blur removes focus from the input field.
func (m *TuiTextInputModel) Blur() {
	m.input.Blur()
}

// Focused returns whether the input is focused.
func (m TuiTextInputModel) Focused() bool {
	return m.input.Focused()
}

// SetValue sets the input value.
func (m *TuiTextInputModel) SetValue(s string) {
	m.input.SetValue(s)
}

// Value returns the current input value.
func (m TuiTextInputModel) Value() string {
	return m.input.Value()
}

// SetCharLimit sets the maximum number of characters allowed.
func (m *TuiTextInputModel) SetCharLimit(n int) {
	m.input.CharLimit = n
}

// CharLimit returns the current character limit.
func (m TuiTextInputModel) CharLimit() int {
	return m.input.CharLimit
}

// SetDisabled enables or disables the input field.
func (m *TuiTextInputModel) SetDisabled(disabled bool) {
	m.disabled = disabled
	if disabled {
		m.input.Blur()
	}
}

// Disabled returns whether the input is disabled.
func (m TuiTextInputModel) Disabled() bool {
	return m.disabled
}

// SetWidth sets the width of the input field.
func (m *TuiTextInputModel) SetWidth(w int) {
	m.input.Width = w
}
