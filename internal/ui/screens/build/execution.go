package build

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/ui/components"
)

// execBoxWidth returns the Width value for bordered boxes so visual width = m.width.
func (m Model) execBoxWidth() int {
	w := m.width - 2 // subtract border (left+right)
	if w < 1 {
		w = 1
	}
	return w
}

// viewExecution renders the build execution phase.
func (m Model) viewExecution() string {
	title := components.StyleH2.Render("Build")
	boxWidth := m.execBoxWidth()

	// Build Execution box: command preview
	cmdLine := components.StylePrompt.Render("$") + " " +
		lipgloss.NewStyle().Foreground(components.ColorGreen).Render(m.config.ToCommand())

	buildExecBox := lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorBorder).
		Padding(0, 1).
		Width(boxWidth).
		Render(
			components.StyleH3.Render("Build Execution") + "\n" +
				cmdLine,
		)

	// Components box
	componentsBox := m.viewComponentTable()

	// Progress box (includes status counters)
	progressBox := m.viewProgressBox()

	// Actions box
	actions := m.viewExecutionActions()

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		buildExecBox,
		"",
		componentsBox,
		"",
		progressBox,
		"",
		actions,
	)
}

// viewComponentTable renders the component build table wrapped in a bordered box.
func (m Model) viewComponentTable() string {
	if len(m.buildStates) == 0 {
		return ""
	}

	// Column widths
	colStatus := 4
	colName := 20
	colPhase := 12
	colProgress := 22
	colTime := 8

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(components.ColorTextSecondary)

	// Use padding to account for the tree-branch prefix width (4 chars: "├── " or ">   ")
	headerPad := "    "
	header := headerPad + lipgloss.JoinHorizontal(lipgloss.Top,
		headerStyle.Width(colStatus).Render("St"),
		headerStyle.Width(colName).Render("Component"),
		headerStyle.Width(colPhase).Render("Phase"),
		headerStyle.Width(colProgress).Render("Progress"),
		headerStyle.Width(colTime).Render("Time"),
	)

	divider := components.TuiDivider(components.DividerSingle, colStatus+colName+colPhase+colProgress+colTime+4)

	var rows []string
	rows = append(rows, header)
	rows = append(rows, divider)

	for i, state := range m.buildStates {
		row := m.viewComponentRow(i, state, colStatus, colName, colPhase, colProgress, colTime)
		rows = append(rows, row)
	}

	tableContent := strings.Join(rows, "\n")

	return lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorCyan).
		Padding(0, 1).
		Width(m.execBoxWidth()).
		Render(
			components.StyleH3.Render("Components") + "\n" +
				tableContent,
		)
}

// viewComponentRow renders a single component row in the build table.
func (m Model) viewComponentRow(
	idx int,
	state componentBuildState,
	colStatus, colName, colPhase, colProgress, colTime int,
) string {
	isFocused := idx == m.buildCursor

	// Status icon (compact single character)
	status := phaseToStatus(state.Phase)
	statusStr := components.TuiStatusCompact(status)
	statusCell := lipgloss.NewStyle().Width(colStatus).Render(statusStr)

	// Component name
	nameCell := lipgloss.NewStyle().
		Width(colName).
		Foreground(components.ColorTextPrimary).
		Render(state.Name)

	// Phase
	phaseCell := lipgloss.NewStyle().
		Width(colPhase).
		Foreground(phaseColor(state.Phase)).
		Render(state.Phase.String())

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

	rowCells := lipgloss.JoinHorizontal(lipgloss.Top,
		statusCell,
		nameCell,
		phaseCell,
		progressCell,
		timeCell,
	)

	// Tree-branch prefix and highlight
	if isFocused {
		totalWidth := colStatus + colName + colPhase + colProgress + colTime + 4
		return lipgloss.NewStyle().
			Background(components.ColorCyan).
			Foreground(components.ColorBackground).
			Width(totalWidth).
			Render("> " + "  " + rowCells)
	}

	return "├── " + rowCells
}

// viewProgressBox renders the progress section with overall bar and status counters in a bordered box.
func (m Model) viewProgressBox() string {
	progress := m.overallProgress()
	boxWidth := m.execBoxWidth()

	label := lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Bold(true).
		Render("Overall:")

	// Inner width = boxWidth - 2 (padding); subtract label+gap for bar
	bar := components.TuiProgress(progress, boxWidth-30, true)

	progressLine := label + "  " + bar
	counters := m.viewStatusCounters()
	content := progressLine + "\n" + counters

	return lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorBorder).
		Padding(0, 1).
		Width(boxWidth).
		Render(
			components.StyleH3.Render("Progress") + "\n" +
				content,
		)
}

// viewStatusCounters renders the running/success/failed/pending counters as colored pill badges.
func (m Model) viewStatusCounters() string {
	running, success, failed, pending := m.statusCounts()

	runningBadge := lipgloss.NewStyle().
		Background(components.ColorCyan).
		Foreground(components.ColorBackground).
		Bold(true).
		Padding(0, 1).
		Render(fmt.Sprintf("● Running: %d", running))

	successBadge := lipgloss.NewStyle().
		Background(components.ColorGreen).
		Foreground(components.ColorBackground).
		Bold(true).
		Padding(0, 1).
		Render(fmt.Sprintf("✓ Success: %d", success))

	failedBadge := lipgloss.NewStyle().
		Background(components.ColorDestructive).
		Foreground(components.ColorTextPrimary).
		Bold(true).
		Padding(0, 1).
		Render(fmt.Sprintf("✗ Failed: %d", failed))

	pendingBadge := lipgloss.NewStyle().
		Background(components.ColorSecondary).
		Foreground(components.ColorTextPrimary).
		Padding(0, 1).
		Render(fmt.Sprintf("○ Pending: %d", pending))

	return lipgloss.JoinHorizontal(lipgloss.Top,
		runningBadge, "  ", successBadge, "  ", failedBadge, "  ", pendingBadge)
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

	boxWidth := m.execBoxWidth()
	actInnerWidth := boxWidth - 2 // subtract padding

	var actionsContent string
	if hintsStr != "" {
		buttonsWidth := lipgloss.Width(buttons)
		hintsWidth := lipgloss.Width(hintsStr)
		gapWidth := actInnerWidth - buttonsWidth - hintsWidth
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
		Width(boxWidth).
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
