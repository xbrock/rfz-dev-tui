// Package logs provides the Log Viewer screen model.
//
// The Logs screen allows users to:
// - View real-time build output
// - Filter logs by level (Info, Warn, Error)
// - Search within logs
// - Export logs to file
// - Scroll through log history
package logs

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// LogLevel represents log severity levels.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// LogEntry represents a single log line.
type LogEntry struct {
	Level     LogLevel
	Timestamp string
	Message   string
	Component string
}

// Model represents the Logs screen state.
type Model struct {
	// Bubbles viewport for scrollable content
	// CRITICAL: Use Bubbles viewport, not custom scrolling
	viewport viewport.Model

	// Log data
	logs        []LogEntry
	filteredLog []LogEntry

	// Filter state
	levelFilter LogLevel
	showAll     bool

	// Search state
	searchQuery  string
	searchActive bool

	// Window dimensions
	width  int
	height int

	// Auto-scroll to bottom
	autoScroll bool
}

// New creates a new Logs screen model.
func New() Model {
	vp := viewport.New(0, 0)
	vp.Style = components.StyleContent

	return Model{
		viewport:   vp,
		logs:       []LogEntry{},
		showAll:    true,
		autoScroll: true,
	}
}

// SetSize updates the model dimensions.
func (m Model) SetSize(width, height int) Model {
	m.width = width
	m.height = height

	// Update viewport dimensions
	// Account for header and footer
	m.viewport.Width = width - 4
	m.viewport.Height = height - 6

	return m
}

// Init initializes the Logs screen.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages for the Logs screen.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 6

	case tea.KeyMsg:
		switch msg.String() {
		case "f":
			// Toggle filter
			m.cycleFilter()
			m.applyFilter()
			m.updateViewportContent()
			return m, nil

		case "/":
			// Activate search
			m.searchActive = true
			return m, nil

		case "esc":
			// Cancel search
			m.searchActive = false
			m.searchQuery = ""
			m.applyFilter()
			m.updateViewportContent()
			return m, nil

		case "g":
			// Go to top
			m.viewport.GotoTop()
			m.autoScroll = false
			return m, nil

		case "G":
			// Go to bottom
			m.viewport.GotoBottom()
			m.autoScroll = true
			return m, nil
		}

	case logReceivedMsg:
		// Append new log entry
		m.logs = append(m.logs, msg.Entry)
		m.applyFilter()
		m.updateViewportContent()

		// Auto-scroll if enabled
		if m.autoScroll {
			m.viewport.GotoBottom()
		}
	}

	// Update viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the Logs screen.
func (m Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Header with filter status
	header := m.renderHeader()

	// Viewport content
	content := m.viewport.View()

	// Footer with scroll position
	footer := m.renderFooter()

	// Combine sections
	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

// renderHeader renders the log viewer header.
func (m Model) renderHeader() string {
	title := components.StyleH2.Render("Build Logs")

	filterLabel := m.getFilterLabel()
	filter := components.StyleBodySecondary.Render("Filter: ") +
		components.StyleKeyboard.Render(filterLabel)

	return lipgloss.JoinHorizontal(lipgloss.Top, title, "  ", filter)
}

// renderFooter renders the log viewer footer.
func (m Model) renderFooter() string {
	scrollInfo := components.StyleBodyMuted.Render(
		m.viewport.ScrollPercent(),
	)

	autoScrollIndicator := ""
	if m.autoScroll {
		autoScrollIndicator = components.StyleLogInfo.Render(" [AUTO]")
	}

	return scrollInfo + autoScrollIndicator
}

// cycleFilter cycles through log level filters.
func (m *Model) cycleFilter() {
	if m.showAll {
		m.showAll = false
		m.levelFilter = LogLevelInfo
	} else {
		switch m.levelFilter {
		case LogLevelInfo:
			m.levelFilter = LogLevelWarn
		case LogLevelWarn:
			m.levelFilter = LogLevelError
		case LogLevelError:
			m.showAll = true
		}
	}
}

// getFilterLabel returns the current filter label.
func (m Model) getFilterLabel() string {
	if m.showAll {
		return "All"
	}
	switch m.levelFilter {
	case LogLevelInfo:
		return "Info+"
	case LogLevelWarn:
		return "Warn+"
	case LogLevelError:
		return "Error"
	default:
		return "All"
	}
}

// applyFilter filters logs based on current filter settings.
func (m *Model) applyFilter() {
	if m.showAll && m.searchQuery == "" {
		m.filteredLog = m.logs
		return
	}

	m.filteredLog = make([]LogEntry, 0)
	for _, entry := range m.logs {
		// Level filter
		if !m.showAll && entry.Level < m.levelFilter {
			continue
		}

		// Search filter
		// TODO: Implement search matching

		m.filteredLog = append(m.filteredLog, entry)
	}
}

// updateViewportContent updates the viewport with filtered logs.
func (m *Model) updateViewportContent() {
	content := m.renderLogs()
	m.viewport.SetContent(content)
}

// renderLogs renders all filtered log entries.
func (m Model) renderLogs() string {
	if len(m.filteredLog) == 0 {
		return components.StyleBodyMuted.Render("No logs to display")
	}

	var lines []string
	for _, entry := range m.filteredLog {
		lines = append(lines, m.renderLogLine(entry))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// renderLogLine renders a single log entry.
func (m Model) renderLogLine(entry LogEntry) string {
	var levelStyle lipgloss.Style
	var levelText string

	switch entry.Level {
	case LogLevelDebug:
		levelStyle = components.StyleLogDebug
		levelText = "DEBUG"
	case LogLevelInfo:
		levelStyle = components.StyleLogInfo
		levelText = "INFO"
	case LogLevelWarn:
		levelStyle = components.StyleLogWarning
		levelText = "WARN"
	case LogLevelError:
		levelStyle = components.StyleLogError
		levelText = "ERROR"
	}

	timestamp := components.StyleLogTimestamp.Render(entry.Timestamp)
	level := levelStyle.Render("[" + levelText + "]")
	message := components.StyleLogMessage.Render(entry.Message)

	return lipgloss.JoinHorizontal(lipgloss.Top, timestamp, " ", level, " ", message)
}

// Messages

type logReceivedMsg struct {
	Entry LogEntry
}
