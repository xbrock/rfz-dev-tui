package build

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/ui/components"
)

// viewExecution renders the build execution phase.
func (m Model) viewExecution() string {
	title := components.StyleH2.Render("Build")

	// Command preview
	cmdLine := components.StylePrompt.Render("$") + " " +
		lipgloss.NewStyle().Foreground(components.ColorGreen).Render(m.config.ToCommand())

	// Component table
	table := m.viewComponentTable()

	// Overall progress
	progressSection := m.viewOverallProgress()

	// Status counters
	counters := m.viewStatusCounters()

	// Actions
	actions := m.viewExecutionActions()

	content := lipgloss.JoinVertical(lipgloss.Left,
		cmdLine,
		"",
		table,
		"",
		progressSection,
		"",
		counters,
	)

	contentBox := lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorCyan).
		Padding(0, 1).
		Width(m.width).
		Render(
			components.StyleH3.Render("Build Execution") + "\n" +
				content,
		)

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		contentBox,
		"",
		actions,
	)
}

// viewComponentTable renders the component build table with per-row status.
func (m Model) viewComponentTable() string {
	if len(m.buildStates) == 0 {
		return ""
	}

	// Column widths
	colStatus := 11
	colName := 20
	colPhase := 12
	colProgress := 22
	colTime := 8

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(components.ColorTextSecondary)

	header := lipgloss.JoinHorizontal(lipgloss.Top,
		headerStyle.Width(colStatus).Render("Status"),
		headerStyle.Width(colName).Render("Component"),
		headerStyle.Width(colPhase).Render("Phase"),
		headerStyle.Width(colProgress).Render("Progress"),
		headerStyle.Width(colTime).Render("Time"),
	)

	divider := components.TuiDivider(components.DividerSingle, colStatus+colName+colPhase+colProgress+colTime)

	var rows []string
	rows = append(rows, header)
	rows = append(rows, divider)

	for i, state := range m.buildStates {
		row := m.viewComponentRow(i, state, colStatus, colName, colPhase, colProgress, colTime)
		rows = append(rows, row)
	}

	return strings.Join(rows, "\n")
}

// viewComponentRow renders a single component row in the build table.
func (m Model) viewComponentRow(
	idx int,
	state componentBuildState,
	colStatus, colName, colPhase, colProgress, colTime int,
) string {
	isFocused := idx == m.buildCursor

	// Status badge
	status := phaseToStatus(state.Phase)
	statusStr := components.TuiStatus(status)
	statusCell := lipgloss.NewStyle().Width(colStatus).Render(statusStr)

	// Component name
	nameStyle := lipgloss.NewStyle().
		Width(colName).
		Foreground(components.ColorTextPrimary)
	if isFocused {
		nameStyle = nameStyle.Bold(true).Foreground(components.ColorCyan)
	}
	nameCell := nameStyle.Render(state.Name)

	// Phase
	phaseStyle := lipgloss.NewStyle().
		Width(colPhase).
		Foreground(phaseColor(state.Phase))
	phaseCell := phaseStyle.Render(state.Phase.String())

	// Progress bar (inline)
	var progressCell string
	switch state.Phase {
	case domain.PhasePending:
		progressCell = lipgloss.NewStyle().
			Width(colProgress).
			Foreground(components.ColorTextMuted).
			Render("waiting...")
	case domain.PhaseDone:
		progressCell = lipgloss.NewStyle().
			Width(colProgress).
			Foreground(components.ColorGreen).
			Render("complete")
	case domain.PhaseFailed:
		progressCell = lipgloss.NewStyle().
			Width(colProgress).
			Foreground(components.ColorDestructive).
			Render("failed")
	default:
		progressCell = components.TuiProgress(state.Progress, colProgress-6, false)
	}

	// Elapsed time
	elapsed := formatDuration(state.Elapsed)
	timeCell := lipgloss.NewStyle().
		Width(colTime).
		Foreground(components.ColorTextSecondary).
		Render(elapsed)

	// Row cursor indicator
	cursor := "  "
	if isFocused {
		cursor = lipgloss.NewStyle().Foreground(components.ColorCyan).Render("> ")
	}

	return cursor + lipgloss.JoinHorizontal(lipgloss.Top,
		statusCell,
		nameCell,
		phaseCell,
		progressCell,
		timeCell,
	)
}

// viewOverallProgress renders the overall build progress section.
func (m Model) viewOverallProgress() string {
	progress := m.overallProgress()

	label := lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Bold(true).
		Render("Overall Progress")

	bar := components.TuiProgress(progress, m.width-30, true)

	return label + "  " + bar
}

// viewStatusCounters renders the running/success/failed/pending counters.
func (m Model) viewStatusCounters() string {
	running, success, failed, pending := m.statusCounts()

	runningStyle := lipgloss.NewStyle().Foreground(components.ColorCyan).Bold(true)
	successStyle := lipgloss.NewStyle().Foreground(components.ColorGreen).Bold(true)
	failedStyle := lipgloss.NewStyle().Foreground(components.ColorDestructive).Bold(true)
	pendingStyle := lipgloss.NewStyle().Foreground(components.ColorTextMuted).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(components.ColorTextSecondary)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		runningStyle.Render(fmt.Sprintf("%d", running)),
		labelStyle.Render(" Running"),
		lipgloss.NewStyle().Render("    "),
		successStyle.Render(fmt.Sprintf("%d", success)),
		labelStyle.Render(" Success"),
		lipgloss.NewStyle().Render("    "),
		failedStyle.Render(fmt.Sprintf("%d", failed)),
		labelStyle.Render(" Failed"),
		lipgloss.NewStyle().Render("    "),
		pendingStyle.Render(fmt.Sprintf("%d", pending)),
		labelStyle.Render(" Pending"),
	)
}

// viewExecutionActions renders the action buttons during/after build execution.
func (m Model) viewExecutionActions() string {
	var buttons string
	var hintsStr string

	if m.buildDone() || m.buildCanceled {
		// Completed/Canceled: View Logs + Rebuild Failed + Back, with Tab Switch Focus hint
		viewLogsBtn := components.TuiButton("View Logs", components.ButtonSecondary, "l", false)
		rebuildBtn := components.TuiButton("Rebuild Failed", components.ButtonPrimary, "r", false)
		backBtn := components.TuiButton("Back", components.ButtonSecondary, "Esc", false)
		buttons = viewLogsBtn + "  " + rebuildBtn + "  " + backBtn
		hintsStr = components.TuiKeyHints([]components.KeyHint{
			{Key: "Tab", Label: "Switch Focus"},
		}, 0)
	} else {
		// Running: View Logs + Cancel Build, no hints on right
		viewLogsBtn := components.TuiButton("View Logs", components.ButtonSecondary, "l", false)
		cancelBtn := components.TuiButton("Cancel Build", components.ButtonDestructive, "Esc", false)
		buttons = viewLogsBtn + "  " + cancelBtn
	}

	var actionsContent string
	if hintsStr != "" {
		buttonsWidth := lipgloss.Width(buttons)
		hintsWidth := lipgloss.Width(hintsStr)
		gapWidth := m.width - buttonsWidth - hintsWidth - 6
		if gapWidth < 1 {
			gapWidth = 1
		}
		actionsContent = lipgloss.JoinHorizontal(lipgloss.Top,
			buttons,
			lipgloss.NewStyle().Width(gapWidth).Render(""),
			hintsStr,
		)
	} else {
		actionsContent = buttons
	}

	borderColor := components.ColorBorder
	if m.focused {
		borderColor = components.ColorCyan
	}

	return lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(m.width).
		Render(
			components.StyleH3.Render("Actions") + "\n" +
				actionsContent,
		)
}

// phaseToStatus maps a BuildPhase to a component Status for badge rendering.
func phaseToStatus(phase domain.BuildPhase) components.Status {
	switch phase {
	case domain.PhasePending:
		return components.StatusPending
	case domain.PhaseDone:
		return components.StatusSuccess
	case domain.PhaseFailed:
		return components.StatusFailed
	default:
		return components.StatusRunning
	}
}

// phaseColor returns the color for a build phase.
func phaseColor(phase domain.BuildPhase) lipgloss.Color {
	switch phase {
	case domain.PhasePending:
		return components.ColorTextMuted
	case domain.PhaseDone:
		return components.ColorGreen
	case domain.PhaseFailed:
		return components.ColorDestructive
	default:
		return components.ColorCyan
	}
}

// formatDuration formats a duration as MM:SS.
func formatDuration(d time.Duration) string {
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}
