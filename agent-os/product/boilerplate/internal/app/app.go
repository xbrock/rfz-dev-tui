// Package app contains the main Bubble Tea application model.
//
// This is the root model that:
// - Contains all screen models as embedded fields
// - Routes messages to the active screen
// - Handles global key bindings (quit, navigation)
// - Manages modal state (overlay dialogs)
// - Coordinates screen transitions
package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/service"
	"rfz-cli/internal/ui/components"
	"rfz-cli/internal/ui/screens/build"
	"rfz-cli/internal/ui/screens/config"
	"rfz-cli/internal/ui/screens/discover"
	"rfz-cli/internal/ui/screens/logs"
	"rfz-cli/internal/ui/screens/welcome"
)

// Screen represents the currently active screen
type Screen int

const (
	ScreenWelcome Screen = iota
	ScreenBuild
	ScreenLogs
	ScreenDiscover
	ScreenConfig
)

// Model is the main application model containing all screen models.
// It implements tea.Model interface.
type Model struct {
	// Active screen tracking
	activeScreen Screen

	// Child screen models
	// All screens exist in memory; only the active one receives updates
	welcome  welcome.Model
	build    build.Model
	logs     logs.Model
	discover discover.Model
	config   config.Model

	// Modal state
	// Modals overlay on top of the active screen
	// TODO: Add modal models (buildconfig.Model, confirm.Model)
	showModal bool
	modalType ModalType

	// Window dimensions
	// Updated on tea.WindowSizeMsg
	width  int
	height int

	// Global key bindings
	keys keyMap

	// Application services (injected via constructor)
	// Services handle business logic; screens focus on UI state
	buildSvc  *service.BuildService
	scanSvc   *service.ScanService
	configSvc *service.ConfigService
}

// ModalType represents the type of modal currently displayed
type ModalType int

const (
	ModalNone ModalType = iota
	ModalBuildConfig
	ModalConfirm
)

// New creates a new application model with injected services.
// This is called from main.go with real or mock service implementations.
func New(buildSvc *service.BuildService, scanSvc *service.ScanService, configSvc *service.ConfigService) Model {
	return Model{
		activeScreen: ScreenWelcome,

		// Initialize child screen models
		// Each screen receives relevant services via its constructor
		welcome:  welcome.New(),
		build:    build.New(buildSvc),
		logs:     logs.New(),
		discover: discover.New(scanSvc),
		config:   config.New(configSvc),

		// Initialize key bindings
		keys: defaultKeyMap(),

		// Store service references for later use
		buildSvc:  buildSvc,
		scanSvc:   scanSvc,
		configSvc: configSvc,
	}
}

// Init initializes the application.
// Called once when the program starts.
func (m Model) Init() tea.Cmd {
	// Return any initialization commands
	// For example, could load initial configuration here
	return nil
}

// Update handles all messages for the application.
// It routes messages to the appropriate screen or handles global messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Handle window resize
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Propagate resize to all screens
		// Each screen adjusts its layout based on available space
		m.welcome = m.welcome.SetSize(msg.Width, msg.Height)
		m.build = m.build.SetSize(msg.Width, msg.Height)
		m.logs = m.logs.SetSize(msg.Width, msg.Height)
		m.discover = m.discover.SetSize(msg.Width, msg.Height)
		m.config = m.config.SetSize(msg.Width, msg.Height)

		return m, nil

	// Handle keyboard input
	case tea.KeyMsg:
		// Global key handling (quit, navigation)
		switch {
		case msg.String() == "ctrl+c" || msg.String() == "q":
			return m, tea.Quit

		// Navigation shortcuts (1-5 keys)
		case msg.String() == "1":
			m.activeScreen = ScreenWelcome
			return m, nil
		case msg.String() == "2":
			m.activeScreen = ScreenBuild
			return m, nil
		case msg.String() == "3":
			m.activeScreen = ScreenLogs
			return m, nil
		case msg.String() == "4":
			m.activeScreen = ScreenDiscover
			return m, nil
		case msg.String() == "5":
			m.activeScreen = ScreenConfig
			return m, nil
		}

	// Handle navigation messages from child screens
	case NavigateToScreenMsg:
		m.activeScreen = msg.Screen
		return m, nil

	// Handle modal messages
	case ShowModalMsg:
		m.showModal = true
		m.modalType = msg.Modal
		return m, nil

	case HideModalMsg:
		m.showModal = false
		m.modalType = ModalNone
		return m, nil
	}

	// Route message to active screen
	// Only the active screen processes non-global messages
	switch m.activeScreen {
	case ScreenWelcome:
		m.welcome, cmd = m.welcome.Update(msg)
	case ScreenBuild:
		m.build, cmd = m.build.Update(msg)
	case ScreenLogs:
		m.logs, cmd = m.logs.Update(msg)
	case ScreenDiscover:
		m.discover, cmd = m.discover.Update(msg)
	case ScreenConfig:
		m.config, cmd = m.config.Update(msg)
	}

	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View renders the application.
// Combines the active screen view with navigation, header, and footer.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	// Calculate layout dimensions
	// Header: 3 lines, Footer: 2 lines, Navigation: 30 chars wide
	navWidth := 30
	contentWidth := m.width - navWidth - 2 // 2 for borders
	contentHeight := m.height - 5          // header + footer

	// Render navigation sidebar
	nav := m.renderNavigation(navWidth, contentHeight)

	// Render active screen content
	var content string
	switch m.activeScreen {
	case ScreenWelcome:
		content = m.welcome.View()
	case ScreenBuild:
		content = m.build.View()
	case ScreenLogs:
		content = m.logs.View()
	case ScreenDiscover:
		content = m.discover.View()
	case ScreenConfig:
		content = m.config.View()
	}

	// Apply content styling
	contentStyle := lipgloss.NewStyle().
		Width(contentWidth).
		Height(contentHeight)
	content = contentStyle.Render(content)

	// Combine navigation and content
	main := lipgloss.JoinHorizontal(lipgloss.Top, nav, content)

	// Render header
	header := m.renderHeader()

	// Render footer/status bar
	footer := components.RenderStatusBar(m.width, m.getStatusBarItems())

	// Combine all sections
	return lipgloss.JoinVertical(lipgloss.Left, header, main, footer)
}

// renderNavigation renders the sidebar navigation panel.
func (m Model) renderNavigation(width, height int) string {
	// TODO: Implement navigation rendering using TuiNavItem component
	// Should show menu items with focus indicators and shortcuts
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.NormalBorder()).
		BorderForeground(components.ColorBorder).
		Render("Navigation")
}

// renderHeader renders the application header bar.
func (m Model) renderHeader() string {
	// TODO: Implement header with title, version, and context info
	return lipgloss.NewStyle().
		Width(m.width).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(components.ColorBrand).
		Render("RFZ-CLI v1.0.0")
}

// getStatusBarItems returns the current status bar items.
func (m Model) getStatusBarItems() []components.StatusBarItem {
	// Return context-appropriate status bar items
	return []components.StatusBarItem{
		{Key: "q", Label: "Quit"},
		{Key: "1-5", Label: "Navigate"},
		{Key: "?", Label: "Help"},
	}
}
