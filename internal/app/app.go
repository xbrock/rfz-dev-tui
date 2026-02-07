// Package app provides the main RFZ CLI application model.
package app

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
	"rfz-cli/internal/ui/screens/welcome"
)

// Focus area within the application.
type focusArea int

const (
	focusNav     focusArea = iota // Navigation sidebar is focused
	focusContent                  // Content area is focused
)

// Navigation panel width (fixed).
const navWidth = 30

// Minimum terminal dimensions.
const (
	minWidth  = 80
	minHeight = 24
)

// Screen indices for navigation items.
const (
	screenBuild  = 0
	screenLogs   = 1
	screenDiscover = 2
	screenConfig = 3
	screenExit   = 4
)

// navItems defines the navigation menu entries.
var navItems = []components.TuiNavItem{
	{Label: "Build Components", Number: 1, Shortcut: "1"},
	{Label: "View Logs", Number: 2, Shortcut: "2"},
	{Label: "Discover", Number: 3, Shortcut: "3"},
	{Label: "Configuration", Number: 4, Shortcut: "4"},
	{Label: "Exit", Number: 5, Shortcut: "q"},
}

// Model is the top-level Bubble Tea model for the RFZ CLI application.
type Model struct {
	width       int
	height      int
	focus       focusArea
	cursorIndex int // Navigation cursor position
	activeIndex int // Currently active screen (-1 = welcome/home)
	currentTime time.Time

	welcome welcome.Model
}

// New creates a new application model.
func New() Model {
	return Model{
		cursorIndex: 0,
		activeIndex: -1, // Welcome screen (no active nav item)
		currentTime: time.Now(),
		welcome:     welcome.New(0, 0),
	}
}

// tickCmd returns a command that sends a TickMsg every second.
func tickCmd() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return tickCmd()
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.welcome = m.welcome.SetSize(m.contentWidth(), m.contentHeight())
		return m, nil
	case TickMsg:
		m.currentTime = time.Time(msg)
		return m, tickCmd()
	}
	return m, nil
}

// handleKey processes keyboard input.
func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		if m.focus == focusNav {
			return m, tea.Quit
		}
	case "up", "k":
		if m.focus == focusNav && m.cursorIndex > 0 {
			m.cursorIndex--
		}
	case "down", "j":
		if m.focus == focusNav && m.cursorIndex < len(navItems)-1 {
			m.cursorIndex++
		}
	case "1":
		m.cursorIndex = screenBuild
	case "2":
		m.cursorIndex = screenLogs
	case "3":
		m.cursorIndex = screenDiscover
	case "4":
		m.cursorIndex = screenConfig
	case "5":
		m.cursorIndex = screenExit
	case "enter":
		if m.focus == focusNav {
			if m.cursorIndex == screenExit {
				return m, tea.Quit
			}
			m.activeIndex = m.cursorIndex
		}
	case "tab":
		if m.focus == focusNav {
			m.focus = focusContent
		} else {
			m.focus = focusNav
		}
	}
	return m, nil
}

// View implements tea.Model.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	// Check minimum terminal size
	if m.width < minWidth || m.height < minHeight {
		return m.viewTooSmall()
	}

	header := m.viewHeader()
	statusBar := m.viewStatusBar()

	// Calculate body height: total - header - statusbar
	headerHeight := lipgloss.Height(header)
	statusBarHeight := lipgloss.Height(statusBar)
	bodyHeight := m.height - headerHeight - statusBarHeight
	if bodyHeight < 1 {
		bodyHeight = 1
	}

	body := m.viewBody(bodyHeight)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		body,
		statusBar,
	)
}

// viewTooSmall renders a message when the terminal is too small.
func (m Model) viewTooSmall() string {
	msg := fmt.Sprintf("Terminal too small. Please resize to at least %dx%d.", minWidth, minHeight)
	style := lipgloss.NewStyle().
		Foreground(components.ColorYellow).
		Bold(true).
		Width(m.width).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Height(m.height)
	return style.Render(msg)
}

// viewHeader renders the top header bar.
func (m Model) viewHeader() string {
	title := components.StyleHeaderTitle.Render("RFZ-CLI v1.0.0")
	subtitle := components.StyleHeaderSubtitle.Render("Terminal Orchestration Tool")

	leftContent := lipgloss.JoinVertical(lipgloss.Left, title, subtitle)

	timeStr := m.currentTime.Format("3:04:05 PM")
	rightText := timeStr + " | Deutsche Bahn Internal"
	rightContent := lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Render(rightText)

	// Calculate gap
	leftWidth := lipgloss.Width(leftContent)
	rightWidth := lipgloss.Width(rightContent)
	headerInnerWidth := m.width - 2 // account for StyleHeader padding
	gapWidth := headerInnerWidth - leftWidth - rightWidth
	if gapWidth < 1 {
		gapWidth = 1
	}
	gap := lipgloss.NewStyle().Width(gapWidth).Render("")

	headerContent := lipgloss.JoinHorizontal(lipgloss.Top,
		leftContent,
		gap,
		rightContent,
	)

	return components.StyleHeader.Width(m.width).Render(headerContent)
}

// viewBody renders the main body with navigation sidebar and content area.
func (m Model) viewBody(height int) string {
	nav := m.viewNavigation(height)
	content := m.viewContent(height)

	return lipgloss.JoinHorizontal(lipgloss.Top, nav, content)
}

// viewNavigation renders the left navigation sidebar.
func (m Model) viewNavigation(height int) string {
	focused := m.focus == focusNav

	// Build footer with key hints
	footer := components.TuiKeyHints([]components.KeyHint{
		{Key: "\u2191/k", Label: "Up"},
		{Key: "\u2193/j", Label: "Down"},
		{Key: "Enter", Label: "Select"},
		{Key: "1-5", Label: "Quick nav"},
	}, navWidth-4) // account for box border + padding

	navContent := components.TuiNavigation(
		navItems,
		m.cursorIndex,
		m.activeIndex,
		focused,
		"Navigation",
		footer,
		navWidth-4, // inner width (box border 2 + padding 2)
	)

	borderColor := components.ColorBorder
	if focused {
		borderColor = components.ColorCyan
	}

	boxStyle := lipgloss.NewStyle().
		Border(components.BorderSingle).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(navWidth).
		Height(height - 2) // account for border top/bottom

	return boxStyle.Render(navContent)
}

// contentWidth returns the inner width available for content (excluding border/padding).
func (m Model) contentWidth() int {
	// Total width - nav - content box border (2) - content box padding (2)
	w := m.width - navWidth - 4
	if w < 1 {
		w = 1
	}
	return w
}

// contentHeight returns the inner height available for content.
func (m Model) contentHeight() int {
	header := m.viewHeader()
	statusBar := m.viewStatusBar()
	headerHeight := lipgloss.Height(header)
	statusBarHeight := lipgloss.Height(statusBar)
	// Total height - header - statusbar - content box border (2)
	h := m.height - headerHeight - statusBarHeight - 2
	if h < 1 {
		h = 1
	}
	return h
}

// viewContent renders the main content area.
func (m Model) viewContent(height int) string {
	contentWidth := m.width - navWidth
	if contentWidth < 1 {
		contentWidth = 1
	}

	var contentBody string
	switch m.activeIndex {
	case screenBuild:
		contentBody = placeholderScreen("Build Components")
	case screenLogs:
		contentBody = placeholderScreen("View Logs")
	case screenDiscover:
		contentBody = placeholderScreen("Discover")
	case screenConfig:
		contentBody = placeholderScreen("Configuration")
	default:
		contentBody = m.welcome.View()
	}

	borderColor := components.ColorBorder
	if m.focus == focusContent {
		borderColor = components.ColorCyan
	}

	boxStyle := lipgloss.NewStyle().
		Border(components.BorderSingle).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(contentWidth).
		Height(height - 2) // account for border top/bottom

	return boxStyle.Render(contentBody)
}

// placeholderScreen renders a placeholder title for screens not yet implemented.
func placeholderScreen(title string) string {
	return lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Bold(true).
		Render(title)
}

// viewStatusBar renders the bottom status bar.
func (m Model) viewStatusBar() string {
	// Determine context badge based on active screen
	var contextBadge string
	switch m.activeIndex {
	case screenBuild:
		contextBadge = "Build Components"
	case screenLogs:
		contextBadge = "View Logs"
	case screenDiscover:
		contextBadge = "Discover"
	case screenConfig:
		contextBadge = "Configuration"
	default:
		contextBadge = navItems[m.cursorIndex].Label
	}

	return components.TuiStatusBar(components.TuiStatusBarConfig{
		ModeBadge:      "HOME",
		ModeBadgeColor: components.ColorCyan,
		ContextBadge:   contextBadge,
		Hints: []components.KeyHint{
			{Key: "Tab", Label: "Focus"},
			{Key: "\u2191\u2193", Label: "Nav"},
			{Key: "Enter", Label: "Select"},
			{Key: "Esc", Label: "Back"},
		},
		QuitHint: &components.KeyHint{Key: "q", Label: "Quit"},
		Width:    m.width,
	})
}
