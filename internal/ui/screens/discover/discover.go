// Package discover provides the Discover screen model.
//
// The Discover screen allows users to:
// - Scan for RFZ components in the workspace
// - View component hierarchy and dependencies
// - Add/remove components from the registry
// - View component metadata (version, type, etc.)
package discover

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/service"
	"rfz-cli/internal/ui/components"
)

// Model represents the Discover screen state.
type Model struct {
	// Bubbles table for component display
	// CRITICAL: Use Bubbles table, not custom rendering
	table table.Model

	// Component data
	components []domain.Component
	scanning   bool

	// Window dimensions
	width  int
	height int

	// Scan service (injected)
	scanSvc *service.ScanService
}

// New creates a new Discover screen model.
func New(scanSvc *service.ScanService) Model {
	// Define table columns
	columns := []table.Column{
		{Title: "Name", Width: 25},
		{Title: "Type", Width: 15},
		{Title: "Path", Width: 40},
		{Title: "Status", Width: 10},
	}

	// Create table with empty rows
	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	// Apply RFZ styles to the table
	// CRITICAL: Use Lip Gloss styles
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(components.ColorBorder).
		BorderBottom(true).
		Bold(true).
		Foreground(components.ColorCyan)
	s.Selected = s.Selected.
		Foreground(components.ColorTextPrimary).
		Background(components.ColorSecondary).
		Bold(true)
	t.SetStyles(s)

	return Model{
		table:   t,
		scanSvc: scanSvc,
	}
}

// SetSize updates the model dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height

	// Update table height
	m.table.SetHeight(height - 8)

	return m
}

// Init initializes the Discover screen.
func (m Model) Init() tea.Cmd {
	// Start initial scan
	return m.scanCmd()
}

// Update handles messages for the Discover screen.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetHeight(msg.Height - 8)

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			// Refresh/rescan
			if !m.scanning {
				m.scanning = true
				return m, m.scanCmd()
			}
			return m, nil

		case "a":
			// Add selected to registry
			return m, m.addToRegistryCmd()

		case "d":
			// Remove from registry
			return m, m.removeFromRegistryCmd()
		}

	case scanStartedMsg:
		m.scanning = true

	case scanCompleteMsg:
		m.scanning = false
		m.components = msg.Components
		m.updateTableRows()

	case scanErrorMsg:
		m.scanning = false
		// TODO: Show error notification
	}

	// Update table
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

// View renders the Discover screen.
func (m Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Header
	header := m.renderHeader()

	// Table content
	tableView := m.table.View()

	// Footer with hints
	footer := m.renderFooter()

	// Combine sections
	content := lipgloss.JoinVertical(lipgloss.Left, header, tableView, footer)

	return components.StyleContent.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// renderHeader renders the screen header.
func (m Model) renderHeader() string {
	title := components.StyleH2.Render("Discover Components")

	status := ""
	if m.scanning {
		status = components.TuiStatus(components.StatusRunning)
	} else {
		count := len(m.components)
		status = components.StyleBodySecondary.Render(
			lipgloss.NewStyle().Render("Found: ") +
				components.StyleKeyboard.Render("%d component(s)"),
		)
		_ = count // Use count in actual implementation
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, title, "  ", status)
}

// renderFooter renders the footer with keyboard hints.
func (m Model) renderFooter() string {
	hints := []string{
		components.StyleKeyboard.Render("r") + " Rescan",
		components.StyleKeyboard.Render("a") + " Add to registry",
		components.StyleKeyboard.Render("d") + " Remove",
	}

	return components.StyleBodyMuted.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, hints...),
	)
}

// updateTableRows refreshes table rows with current components.
func (m *Model) updateTableRows() {
	rows := make([]table.Row, len(m.components))
	for i, comp := range m.components {
		rows[i] = table.Row{
			comp.Name,
			string(comp.Type),
			comp.Path,
			"Ready", // TODO: Get actual status
		}
	}
	m.table.SetRows(rows)
}

// Commands

func (m Model) scanCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Use scanSvc to scan for components
		return scanCompleteMsg{Components: []domain.Component{}}
	}
}

func (m Model) addToRegistryCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Add selected component to registry
		return nil
	}
}

func (m Model) removeFromRegistryCmd() tea.Cmd {
	return func() tea.Msg {
		// TODO: Remove selected component from registry
		return nil
	}
}

// Messages

type scanStartedMsg struct{}

type scanCompleteMsg struct {
	Components []domain.Component
}

type scanErrorMsg struct {
	Err error
}
