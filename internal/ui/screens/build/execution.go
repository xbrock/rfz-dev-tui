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

	boxWidth := m.execBoxWidth()
	// Inner content width = boxWidth - 2 (padding left+right)
	innerWidth := boxWidth - 2

	// Fixed column widths; name gets the remaining space
	const prefixWidth = 3 // "├─ " or "> " padded to 3
	colStatus := 4
	colPhase := 14
	colProgress := 16
	colTime := 6
	colName := innerWidth - prefixWidth - colStatus - colPhase - colProgress - colTime
	if colName < 10 {
		colName = 10
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(components.ColorTextSecondary)

	// Header with prefix-width padding to align with tree content
	headerPad := strings.Repeat(" ", prefixWidth)
	header := headerPad + lipgloss.JoinHorizontal(lipgloss.Top,
		headerStyle.Width(colStatus).Render("St"),
		headerStyle.Width(colName).Render("Component"),
		headerStyle.Width(colPhase).Align(lipgloss.Right).Render("Phase"),
		headerStyle.Width(colProgress).Align(lipgloss.Right).Render("Progress"),
		headerStyle.Width(colTime).Align(lipgloss.Right).Render("Time"),
	)

	divider := components.TuiDivider(components.DividerSingle, innerWidth)

	var rows []string
	rows = append(rows, header)
	rows = append(rows, divider)

	lastIdx := len(m.buildStates) - 1
	for i, state := range m.buildStates {
		isLast := i == lastIdx
		row := m.viewComponentRow(i, state, colStatus, colName, colPhase, colProgress, colTime, isLast, innerWidth)
		rows = append(rows, row)
	}

	tableContent := strings.Join(rows, "\n")

	return lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorCyan).
		Padding(0, 1).
		Width(boxWidth).
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
	isLast bool,
	totalWidth int,
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

	// Phase (right-aligned)
	phaseCell := lipgloss.NewStyle().
		Width(colPhase).
		Align(lipgloss.Right).
		Foreground(phaseColor(state.Phase)).
		Render(state.Phase.String())

	// Progress bar (braille blocks with color)
	progressCell := m.renderBrailleProgress(state, colProgress)

	// Elapsed time (right-aligned)
	elapsed := formatDuration(state.Elapsed)
	timeCell := lipgloss.NewStyle().
		Width(colTime).
		Align(lipgloss.Right).
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
		return lipgloss.NewStyle().
			Background(components.ColorCyan).
			Foreground(components.ColorBackground).
			Width(totalWidth).
			Render("> " + " " + rowCells)
	}

	// Tree prefix: ├─ for non-last, └─ for last
	prefix := "├─ "
	if isLast {
		prefix = "└─ "
	}

	return prefix + rowCells
}

// renderBrailleProgress renders a braille-block progress bar for a component.
func (m Model) renderBrailleProgress(state componentBuildState, width int) string {
	barWidth := width - 2 // leave margin
	if barWidth < 4 {
		barWidth = 4
	}

	switch state.Phase {
	case domain.PhasePending:
		// Dots pattern for pending
		dots := strings.Repeat("·", barWidth)
		return lipgloss.NewStyle().
			Width(width).
			Align(lipgloss.Right).
			Foreground(components.ColorTextMuted).
			Render(dots)
	case domain.PhaseDone:
		// Full braille blocks in green
		blocks := strings.Repeat("⣿", barWidth)
		return lipgloss.NewStyle().
			Width(width).
			Align(lipgloss.Right).
			Foreground(components.ColorGreen).
			Render(blocks)
	case domain.PhaseFailed:
		// Partial braille blocks in red
		filled := int(float64(barWidth) * state.Progress)
		blocks := strings.Repeat("⣿", filled) + strings.Repeat("·", barWidth-filled)
		return lipgloss.NewStyle().
			Width(width).
			Align(lipgloss.Right).
			Foreground(components.ColorDestructive).
			Render(blocks)
	default:
		// Running: braille blocks in cyan
		filled := int(float64(barWidth) * state.Progress)
		if filled < 1 && state.Progress > 0 {
			filled = 1
		}
		empty := barWidth - filled
		filledStr := lipgloss.NewStyle().Foreground(components.ColorCyan).Render(strings.Repeat("⣿", filled))
		emptyStr := lipgloss.NewStyle().Foreground(components.ColorTextMuted).Render(strings.Repeat("·", empty))
		return lipgloss.NewStyle().
			Width(width).
			Align(lipgloss.Right).
			Render(filledStr + emptyStr)
	}
}

// viewProgressBox renders the progress section with overall bar and status counters in a bordered box.
func (m Model) viewProgressBox() string {
	prog := m.overallProgress()
	boxWidth := m.execBoxWidth()

	label := lipgloss.NewStyle().
		Foreground(components.ColorTextSecondary).
		Bold(true).
		Render("Overall:")

	// Block-style overall progress bar (█ filled, ░ empty)
	barWidth := boxWidth - 30 // account for label + percentage + padding
	if barWidth < 10 {
		barWidth = 10
	}
	filled := int(float64(barWidth) * prog)
	empty := barWidth - filled
	filledStr := lipgloss.NewStyle().Foreground(components.ColorCyan).Render(strings.Repeat("█", filled))
	emptyStr := lipgloss.NewStyle().Foreground(components.ColorTextMuted).Render(strings.Repeat("░", empty))
	percentStr := lipgloss.NewStyle().Foreground(components.ColorTextSecondary).Render(fmt.Sprintf(" %3.0f%%", prog*100))
	bar := filledStr + emptyStr + percentStr

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

// viewStatusCounters renders the success/failed/pending counters as colored pill badges.
// Running badge is omitted per design. Pending is hidden when count is 0.
func (m Model) viewStatusCounters() string {
	_, success, failed, pending := m.statusCounts()

	var badges []string

	successBadge := lipgloss.NewStyle().
		Background(components.ColorGreen).
		Foreground(components.ColorBackground).
		Bold(true).
		Padding(0, 1).
		Render(fmt.Sprintf("✓ Success: %d", success))
	badges = append(badges, successBadge)

	failedBadge := lipgloss.NewStyle().
		Background(components.ColorDestructive).
		Foreground(components.ColorTextPrimary).
		Bold(true).
		Padding(0, 1).
		Render(fmt.Sprintf("✗ Failed: %d", failed))
	badges = append(badges, failedBadge)

	if pending > 0 {
		pendingBadge := lipgloss.NewStyle().
			Background(components.ColorSecondary).
			Foreground(components.ColorTextPrimary).
			Padding(0, 1).
			Render(fmt.Sprintf("○ Pending: %d", pending))
		badges = append(badges, pendingBadge)
	}

	return strings.Join(badges, "  ")
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
