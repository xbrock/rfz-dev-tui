// Package components provides shared UI components and styles.
//
// This file contains the TuiSpinner loading indicator component.
package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SpinnerVariant defines the animation style of a spinner.
type SpinnerVariant string

const (
	// SpinnerBraille uses braille dots pattern (default, most visually appealing).
	SpinnerBraille SpinnerVariant = "braille"

	// SpinnerLine uses |/-\ ASCII characters (fallback for simple terminals).
	SpinnerLine SpinnerVariant = "line"

	// SpinnerCircle uses quarter circle characters for rotation effect.
	SpinnerCircle SpinnerVariant = "circle"

	// SpinnerBounce uses vertical bounce animation.
	SpinnerBounce SpinnerVariant = "bounce"
)

// SpinnerColor defines the color of a spinner.
type SpinnerColor string

const (
	// SpinnerColorCyan is the default spinner color.
	SpinnerColorCyan SpinnerColor = "cyan"

	// SpinnerColorGreen for success-related spinners.
	SpinnerColorGreen SpinnerColor = "green"

	// SpinnerColorYellow for warning-related spinners.
	SpinnerColorYellow SpinnerColor = "yellow"
)

// TuiSpinnerModel is a styled spinner wrapping bubbles/spinner.
type TuiSpinnerModel struct {
	spinner spinner.Model
	label   string
	color   SpinnerColor
}

// NewTuiSpinner creates a new spinner with the specified variant and label.
func NewTuiSpinner(variant SpinnerVariant, label string) TuiSpinnerModel {
	s := spinner.New()
	s.Spinner = mapVariantToSpinner(variant)
	s.Style = lipgloss.NewStyle().Foreground(ColorCyan)

	return TuiSpinnerModel{
		spinner: s,
		label:   label,
		color:   SpinnerColorCyan,
	}
}

// mapVariantToSpinner converts our variant to bubbles spinner type.
func mapVariantToSpinner(variant SpinnerVariant) spinner.Spinner {
	switch variant {
	case SpinnerLine:
		return spinner.Line
	case SpinnerCircle:
		return spinner.MiniDot
	case SpinnerBounce:
		return spinner.Jump
	case SpinnerBraille:
		fallthrough
	default:
		return spinner.Dot
	}
}

// Init implements tea.Model.
func (m TuiSpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update implements tea.Model.
func (m TuiSpinnerModel) Update(msg tea.Msg) (TuiSpinnerModel, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m TuiSpinnerModel) View() string {
	spinnerView := m.spinner.View()

	if m.label != "" {
		labelStyle := lipgloss.NewStyle().Foreground(ColorTextPrimary)
		return spinnerView + " " + labelStyle.Render(m.label)
	}

	return spinnerView
}

// SetColor sets the spinner color.
func (m *TuiSpinnerModel) SetColor(color SpinnerColor) {
	m.color = color

	var c lipgloss.Color
	switch color {
	case SpinnerColorGreen:
		c = ColorGreen
	case SpinnerColorYellow:
		c = ColorYellow
	case SpinnerColorCyan:
		fallthrough
	default:
		c = ColorCyan
	}

	m.spinner.Style = lipgloss.NewStyle().Foreground(c)
}

// SetLabel sets the spinner label.
func (m *TuiSpinnerModel) SetLabel(label string) {
	m.label = label
}

// Label returns the current label.
func (m TuiSpinnerModel) Label() string {
	return m.label
}

// TuiSpinnerStatic renders a static spinner frame for gallery/documentation display.
// This shows the first frame of the animation without requiring a running Bubble Tea model.
func TuiSpinnerStatic(variant SpinnerVariant, label string, color SpinnerColor) string {
	// Get the first frame of the spinner
	s := mapVariantToSpinner(variant)
	frame := s.Frames[0]

	// Apply color
	var c lipgloss.Color
	switch color {
	case SpinnerColorGreen:
		c = ColorGreen
	case SpinnerColorYellow:
		c = ColorYellow
	case SpinnerColorCyan:
		fallthrough
	default:
		c = ColorCyan
	}

	spinnerStyle := lipgloss.NewStyle().Foreground(c)
	spinnerView := spinnerStyle.Render(frame)

	if label != "" {
		labelStyle := lipgloss.NewStyle().Foreground(ColorTextPrimary)
		return spinnerView + " " + labelStyle.Render(label)
	}

	return spinnerView
}
