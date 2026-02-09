// Package welcome provides the Welcome Screen for the RFZ CLI application.
package welcome

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// ASCII art logo lines for "RFZ" and "CLI" portions.
// Each line pair is rendered side by side: RFZ (brand) + CLI (cyan).
var (
	logoRFZ = []string{
		" ██████╗  ███████╗ ███████╗",
		" ██╔══██╗ ██╔════╝ ╚══███╔╝",
		" ██████╔╝ █████╗     ███╔╝ ",
		" ██╔══██╗ ██╔══╝    ███╔╝  ",
		" ██║  ██║ ██║      ███████╗",
		" ╚═╝  ╚═╝ ╚═╝      ╚══════╝",
	}
	logoCLI = []string{
		"  ██████╗ ██╗     ██╗",
		" ██╔════╝ ██║     ██║",
		" ██║      ██║     ██║",
		" ██║      ██║     ██║",
		" ╚██████╗ ███████╗██║",
		"  ╚═════╝ ╚══════╝╚═╝",
	}
)

// Model represents the Welcome Screen state.
type Model struct {
	width  int
	height int
}

// New creates a new Welcome Screen model.
func New(width, height int) Model {
	return Model{
		width:  width,
		height: height,
	}
}

// SetSize updates the screen dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height
	return m
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

// View renders the Welcome Screen.
func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	logo := m.renderLogo()
	subtitle := lipgloss.NewStyle().Foreground(components.ColorTextPrimary).Render("Terminal Orchestration Tool")
	tagline := components.StyleTagline.Render(`"First, solve the problem. Then, write the code."`)
	divider := lipgloss.NewStyle().Foreground(components.ColorTextMuted).Render(strings.Repeat("⣿", 30))
	badges := m.renderBadges()
	status := m.renderStatus()
	hints := m.renderHints()

	content := lipgloss.JoinVertical(lipgloss.Center,
		logo,
		"",
		subtitle,
		tagline,
		"",
		divider,
		"",
		badges,
		"",
		status,
		"",
		hints,
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

// renderLogo builds the ASCII art logo with RFZ in brand color and CLI in cyan.
func (m Model) renderLogo() string {
	rfzStyle := components.StyleASCIIArt
	cliStyle := components.StyleASCIIArtCyan

	var lines []string
	for i := range logoRFZ {
		line := rfzStyle.Render(logoRFZ[i]) + cliStyle.Render(logoCLI[i])
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

// renderBadges builds the version, org, and info badge row.
func (m Model) renderBadges() string {
	version := components.StyleBadgeVersion.Render("v1.0.0")

	org := lipgloss.NewStyle().
		Background(components.ColorSecondary).
		Foreground(components.ColorTextPrimary).
		Padding(0, 1).
		Render("Deutsche Bahn")

	info := lipgloss.NewStyle().
		Background(lipgloss.Color("#164e63")).
		Foreground(components.ColorTextPrimary).
		Padding(0, 1).
		Render("Internal Tool")

	return lipgloss.JoinHorizontal(lipgloss.Center, version, " ", org, " ", info)
}

// renderStatus builds the readiness status line.
func (m Model) renderStatus() string {
	prompt := components.StylePrompt.Render("$")
	text := components.StyleBody.Render(" rfz-cli ready")

	statusLine := lipgloss.JoinHorizontal(lipgloss.Center, prompt, text)

	hint := lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Render("Use navigation panel to get started")

	return lipgloss.JoinVertical(lipgloss.Center, statusLine, "", hint)
}

// renderHints builds the keyboard hints block.
func (m Model) renderHints() string {
	return components.TuiKeyHintsTree([]components.KeyHint{
		{Key: "↑↓/jk", Label: "navigate"},
		{Key: "Enter", Label: "select"},
		{Key: "q", Label: "quit"},
	})
}
