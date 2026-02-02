// Package welcome provides the Welcome screen model.
//
// The Welcome screen is the first screen users see when launching the CLI.
// It displays:
// - ASCII art logo (RFZ-CLI)
// - Version and build information
// - Quick start hints
// - Ready status indicator
package welcome

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// Model represents the Welcome screen state.
type Model struct {
	// Window dimensions
	width  int
	height int

	// Version information
	version   string
	buildDate string

	// Ready state
	ready bool
}

// New creates a new Welcome screen model.
func New() Model {
	return Model{
		version:   "1.0.0",
		buildDate: "2026-02-02",
		ready:     true,
	}
}

// SetSize updates the model dimensions.
// Called when the terminal is resized.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// Init initializes the Welcome screen.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages for the Welcome screen.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

// View renders the Welcome screen.
func (m Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// ASCII art logo
	// TODO: Implement actual ASCII art rendering
	// The "RFZ" should be in Brand color (red)
	// The "CLI" should be in Cyan color
	logo := m.renderLogo()

	// Version badge
	version := components.StyleBadgeVersion.Render("v" + m.version)

	// Tagline
	tagline := components.StyleTagline.Render("Terminal Orchestration Tool")

	// Ready status
	readyStatus := m.renderReadyStatus()

	// Quick start hints
	hints := m.renderQuickStart()

	// Combine all elements vertically
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		logo,
		version,
		tagline,
		"",
		readyStatus,
		"",
		hints,
	)

	// Center content in available space
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

// renderLogo renders the ASCII art logo.
func (m Model) renderLogo() string {
	// TODO: Implement actual ASCII art
	// For now, return a simple placeholder
	// The actual implementation should use:
	// - components.StyleASCIIArt (ColorBrand) for "RFZ"
	// - components.StyleASCIIArtCyan (ColorCyan) for "CLI"

	rfzStyle := components.StyleASCIIArt
	cliStyle := components.StyleASCIIArtCyan

	// Placeholder ASCII art
	rfz := rfzStyle.Render("RFZ")
	cli := cliStyle.Render("-CLI")

	return lipgloss.JoinHorizontal(lipgloss.Center, rfz, cli)
}

// renderReadyStatus renders the ready status indicator.
func (m Model) renderReadyStatus() string {
	if m.ready {
		icon := lipgloss.NewStyle().Foreground(components.ColorGreen).Render("●")
		text := components.StyleBody.Render(" Ready")
		return icon + text
	}

	icon := lipgloss.NewStyle().Foreground(components.ColorYellow).Render("○")
	text := components.StyleBodySecondary.Render(" Initializing...")
	return icon + text
}

// renderQuickStart renders the quick start hints.
func (m Model) renderQuickStart() string {
	hints := []string{
		"Press " + components.StyleKeyboard.Render("2") + " to start building",
		"Press " + components.StyleKeyboard.Render("?") + " for help",
	}

	return lipgloss.JoinVertical(lipgloss.Center, hints...)
}
