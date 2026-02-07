// Package app provides the main RFZ CLI application model.
package app

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
	"rfz-cli/internal/ui/screens/build"
	"rfz-cli/internal/ui/screens/placeholder"
	"rfz-cli/internal/ui/screens/welcome"
)

// Focus area within the application.
type focusArea int

const (
	focusNav     focusArea = iota // Navigation sidebar is focused
	focusContent                  // Content area is focused
)

// activeScreen identifies which screen is displayed in the content area.
type activeScreen int

const (
	screenWelcome  activeScreen = iota // Welcome/Home screen
	screenBuild                        // Build Components
	screenLogs                         // View Logs
	screenDiscover                     // Discover
	screenConfig                       // Configuration
)

// Navigation panel width (fixed).
const navWidth = 30

// Minimum terminal dimensions.
const (
	minWidth  = 80
	minHeight = 24
)

// Navigation item indices (0-based for cursor positioning).
const (
	navBuild   = 0
	navLogs    = 1
	navDiscover = 2
	navConfig  = 3
	navExit    = 4
)

// navItems defines the navigation menu entries.
var navItems = []components.TuiNavItem{
	{Label: "Build Components", Number: 1, Shortcut: "1"},
	{Label: "View Logs", Number: 2, Shortcut: "2"},
	{Label: "Discover", Number: 3, Shortcut: "3"},
	{Label: "Configuration", Number: 4, Shortcut: "4"},
	{Label: "Exit", Number: 5, Shortcut: "q"},
}

// screenNames maps activeScreen to display name for the status bar.
var screenNames = map[activeScreen]string{
	screenBuild:    "Build Components",
	screenLogs:     "View Logs",
	screenDiscover: "Discover",
	screenConfig:   "Configuration",
}

// Model is the top-level Bubble Tea model for the RFZ CLI application.
type Model struct {
	width       int
	height      int
	focus       focusArea
	cursorIndex int          // Navigation cursor position
	activeIndex int          // Active nav highlight (-1 = none/welcome)
	screen      activeScreen // Currently displayed screen
	currentTime time.Time

	showModal       bool // Whether the quit confirmation modal is visible
	modalFocusIndex int  // Focused button in modal (0=Yes, 1=No)

	welcome welcome.Model
	build   build.Model
	phLogs  placeholder.Model
	phDisc   placeholder.Model
	phConfig placeholder.Model
}

// New creates a new application model.
func New() Model {
	return Model{
		cursorIndex: 0,
		activeIndex: -1, // No active nav item on welcome
		screen:      screenWelcome,
		currentTime: time.Now(),
		welcome: welcome.New(0, 0),
		build:   build.New(0, 0),
		phLogs:      placeholder.New("View Logs", 0, 0),
		phDisc:      placeholder.New("Discover", 0, 0),
		phConfig:    placeholder.New("Configuration", 0, 0),
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
		cw, ch := m.contentWidth(), m.contentHeight()
		m.welcome = m.welcome.SetSize(cw, ch)
		m.build = m.build.SetSize(cw, ch)
		m.phLogs = m.phLogs.SetSize(cw, ch)
		m.phDisc = m.phDisc.SetSize(cw, ch)
		m.phConfig = m.phConfig.SetSize(cw, ch)
		return m, nil
	case TickMsg:
		m.currentTime = time.Time(msg)
		return m, tickCmd()
	case build.OpenConfigMsg:
		// Build screen wants to open configuration (handled in BUILD-003)
		return m, nil
	}
	return m, nil
}

// navigateTo switches to the given screen and updates cursor/active state.
func (m *Model) navigateTo(s activeScreen) {
	m.screen = s
	if s == screenWelcome {
		m.activeIndex = -1
	} else {
		idx := int(s) - 1 // screenBuild=1 → navBuild=0, etc.
		m.cursorIndex = idx
		m.activeIndex = idx
	}
}

// handleKey processes keyboard input.
func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Modal captures all input when visible
	if m.showModal {
		return m.handleModalKey(msg)
	}

	// Global keys that always work
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit

	case "tab":
		if m.focus == focusNav {
			m.focus = focusContent
			m.build = m.build.SetFocused(m.screen == screenBuild)
		} else {
			m.focus = focusNav
			m.build = m.build.SetFocused(false)
		}
		return m, nil

	case "esc":
		if m.screen != screenWelcome {
			m.navigateTo(screenWelcome)
			m.build = m.build.SetFocused(false)
		}
		return m, nil
	}

	// Delegate to content screen when focused
	if m.focus == focusContent && m.screen == screenBuild {
		// q in content focus should not trigger quit modal for build screen
		if msg.String() == "q" {
			m.showModal = true
			m.modalFocusIndex = 1
			return m, nil
		}
		var cmd tea.Cmd
		m.build, cmd = m.build.Update(msg)
		return m, cmd
	}

	// Navigation-focused keys
	switch msg.String() {
	case "q":
		m.showModal = true
		m.modalFocusIndex = 1 // Default focus on "No" for safety
		return m, nil

	case "up", "k":
		if m.focus == focusNav {
			m.cursorIndex = (m.cursorIndex - 1 + len(navItems)) % len(navItems)
		}

	case "down", "j":
		if m.focus == focusNav {
			m.cursorIndex = (m.cursorIndex + 1) % len(navItems)
		}

	case "1":
		m.navigateTo(screenBuild)
	case "2":
		m.navigateTo(screenLogs)
	case "3":
		m.navigateTo(screenDiscover)
	case "4":
		m.navigateTo(screenConfig)

	case "enter":
		if m.focus == focusNav {
			if m.cursorIndex == navExit {
				m.showModal = true
				m.modalFocusIndex = 1 // Default focus on "No" for safety
				return m, nil
			}
			// Map nav index to screen: navBuild(0)→screenBuild(1), etc.
			m.navigateTo(activeScreen(m.cursorIndex + 1))
		}
	}
	return m, nil
}

// handleModalKey processes keyboard input when the quit confirmation modal is visible.
func (m Model) handleModalKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "y":
		return m, tea.Quit
	case "n", "esc":
		m.showModal = false
		return m, nil
	case "left", "right", "tab":
		m.modalFocusIndex = 1 - m.modalFocusIndex // Toggle between 0 and 1
	case "enter":
		if m.modalFocusIndex == 0 {
			return m, tea.Quit
		}
		m.showModal = false
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

	// Modal overlays entire screen
	if m.showModal {
		return m.viewQuitModal()
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

// viewQuitModal renders the quit confirmation modal overlay.
func (m Model) viewQuitModal() string {
	config := components.TuiModalConfig{
		Title:   "Quit RFZ-CLI?",
		Content: "Are you sure you want to quit?",
		Buttons: []components.TuiModalButton{
			{Label: "Yes", Variant: components.ButtonPrimary, Shortcut: "y"},
			{Label: "No", Variant: components.ButtonSecondary, Shortcut: "n"},
		},
		FocusedIndex: m.modalFocusIndex,
	}
	return components.TuiModal(config, m.width, m.height)
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
		{Key: "1-4", Label: "Quick nav"},
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
	switch m.screen {
	case screenBuild:
		contentBody = m.build.View()
	case screenLogs:
		contentBody = m.phLogs.View()
	case screenDiscover:
		contentBody = m.phDisc.View()
	case screenConfig:
		contentBody = m.phConfig.View()
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

// viewStatusBar renders the bottom status bar.
func (m Model) viewStatusBar() string {
	// Determine context badge based on active screen
	var contextBadge string
	if name, ok := screenNames[m.screen]; ok {
		contextBadge = name
	} else {
		contextBadge = navItems[m.cursorIndex].Label
	}

	modeBadge := "HOME"
	var hints []components.KeyHint

	if m.screen == screenBuild && m.focus == focusContent {
		modeBadge = "SELECT"
		contextBadge = m.build.CurrentItemLabel()
		hints = []components.KeyHint{
			{Key: "Tab", Label: "Focus"},
			{Key: "\u2191\u2193", Label: "Nav"},
			{Key: "Enter", Label: "Select"},
			{Key: "Esc", Label: "Back"},
		}
	} else {
		hints = []components.KeyHint{
			{Key: "Tab", Label: "Focus"},
			{Key: "\u2191\u2193", Label: "Nav"},
			{Key: "Enter", Label: "Select"},
			{Key: "Esc", Label: "Back"},
		}
	}

	return components.TuiStatusBar(components.TuiStatusBarConfig{
		ModeBadge:      modeBadge,
		ModeBadgeColor: components.ColorCyan,
		ContextBadge:   contextBadge,
		Hints:          hints,
		QuitHint:       &components.KeyHint{Key: "q", Label: "Quit"},
		Width:          m.width,
	})
}
