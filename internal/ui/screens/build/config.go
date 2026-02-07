package build

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// viewConfig renders the build configuration form as a TuiModal overlay.
func (m Model) viewConfig() string {
	// Header: "Building N components: name1, name2, ..."
	header := fmt.Sprintf("Building %d components: %s", len(m.selectedComponents), strings.Join(m.selectedComponents, ", "))

	termW := m.termW
	termH := m.termH
	if termW == 0 {
		termW = m.width
	}
	if termH == 0 {
		termH = m.height
	}

	modalWidth := termW - 16 // leave margin on each side
	if modalWidth < 60 {
		modalWidth = 60
	}
	sectionWidth := modalWidth - 6 // modal border (2) + modal padding (2) + margin (2)

	goalSection := m.viewGoalSection(sectionWidth)
	profilesSection := m.viewProfilesSection(sectionWidth)
	portSection := m.viewPortSection(sectionWidth)
	optionsSection := m.viewOptionsSection(sectionWidth)
	previewSection := m.viewPreviewSection(sectionWidth)
	hintsRow := m.viewConfigHints()

	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		goalSection,
		profilesSection,
		portSection,
		optionsSection,
		"",
		previewSection,
		"",
		components.TuiDivider(components.DividerSingle, sectionWidth),
		"",
		m.viewConfigButtons(sectionWidth),
		"",
		hintsRow,
	)

	modalConfig := components.TuiModalConfig{
		Title:   "Build Configuration",
		Content: content,
		Width:   modalWidth,
	}

	return components.TuiModal(modalConfig, termW, termH)
}

// sectionBox renders a bordered section with a title label in the top border.
func sectionBox(title string, content string, width int, focused bool) string {
	borderColor := components.ColorBorder
	titleColor := components.ColorTextSecondary
	if focused {
		borderColor = components.ColorCyan
		titleColor = components.ColorCyan
	}

	borderStyle := lipgloss.NewStyle().Foreground(borderColor)

	titleRendered := lipgloss.NewStyle().
		Bold(true).
		Foreground(titleColor).
		Render(title)

	innerWidth := width - 4 // border (2) + padding (2)
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Build top border with embedded title: ╭─ Title ──────╮
	titleVisualWidth := lipgloss.Width(titleRendered)
	dashesAfter := innerWidth - titleVisualWidth - 1 // -1 for space before title
	if dashesAfter < 0 {
		dashesAfter = 0
	}

	topLine := borderStyle.Render("╭─") + " " + titleRendered + " " +
		borderStyle.Render(strings.Repeat("─", dashesAfter)+"╮")

	// Build bottom border: ╰──────────────╯
	bottomLine := borderStyle.Render("╰" + strings.Repeat("─", innerWidth+2) + "╯")

	// Render content with side borders
	contentStyle := lipgloss.NewStyle().Width(innerWidth)
	renderedContent := contentStyle.Render(content)

	var lines []string
	lines = append(lines, topLine)
	for _, line := range strings.Split(renderedContent, "\n") {
		lineWidth := lipgloss.Width(line)
		pad := innerWidth - lineWidth
		if pad < 0 {
			pad = 0
		}
		lines = append(lines, borderStyle.Render("│")+" "+line+strings.Repeat(" ", pad)+" "+borderStyle.Render("│"))
	}
	lines = append(lines, bottomLine)

	return strings.Join(lines, "\n")
}

// viewGoalSection renders the Maven Goal radio group section.
func (m Model) viewGoalSection(width int) string {
	focused := m.section == sectionGoal

	labels := make([]string, len(mavenGoals))
	for i, g := range mavenGoals {
		labels[i] = mavenGoalLabels[g]
	}

	selectedIdx := m.goalIndex
	focusedIdx := -1
	if focused {
		focusedIdx = m.goalIndex
	}

	radios := components.TuiRadioGroup(labels, selectedIdx, focusedIdx, true)

	hint := components.StyleBodyMuted.Render("\u2190\u2192 or h/l to select")

	content := radios + "\n" + hint

	return sectionBox("Maven Goal", content, width, focused)
}

// viewProfilesSection renders the Maven Profiles checkbox section.
func (m Model) viewProfilesSection(width int) string {
	focused := m.section == sectionProfiles

	activeProfiles := make(map[string]bool)
	for _, p := range m.config.Profiles {
		activeProfiles[p] = true
	}

	var lines []string
	for i, opt := range profileOptions {
		checked := activeProfiles[opt]
		isFocused := focused && i == m.profileCursor
		cb := components.TuiCheckbox(opt, checked, isFocused, false)
		lines = append(lines, "  "+cb)
	}

	hint := components.StyleBodyMuted.Render("\u2191\u2193 navigate | Space toggle")

	content := strings.Join(lines, "\n") + "\n" + hint

	return sectionBox("Maven Profiles (multi-select)", content, width, focused)
}

// viewPortSection renders the Traktion Port radio group section.
func (m Model) viewPortSection(width int) string {
	focused := m.section == sectionPort

	labels := make([]string, len(portOptions))
	for i, p := range portOptions {
		labels[i] = fmt.Sprintf("Port %d", p)
	}

	selectedIdx := m.portIndex
	focusedIdx := -1
	if focused {
		focusedIdx = m.portIndex
	}

	radios := components.TuiRadioGroup(labels, selectedIdx, focusedIdx, true)

	hint := components.StyleBodyMuted.Render("\u2190\u2192 or h/l to select | Appends use_traktion_* profile")

	content := radios + "\n" + hint

	return sectionBox("Traktion Port", content, width, focused)
}

// viewOptionsSection renders the Build Options (Skip Tests) section.
func (m Model) viewOptionsSection(width int) string {
	focused := m.section == sectionOptions

	cb := components.TuiCheckbox("Skip Tests", m.config.SkipTests, focused, false)

	desc := components.StyleBodyMuted.Render("(adds -DskipTests)")
	line := "  " + cb + "  " + desc

	hint := components.StyleBodyMuted.Render("Space or Enter to toggle")

	content := line + "\n" + hint

	return sectionBox("Build Options", content, width, focused)
}

// viewPreviewSection renders the command preview section.
func (m Model) viewPreviewSection(width int) string {
	cmd := m.config.ToCommand()
	preview := components.StylePrompt.Render("$") + " " +
		lipgloss.NewStyle().Foreground(components.ColorGreen).Render(cmd)

	return sectionBox("Command Preview", preview, width, false)
}

// viewConfigButtons renders the Cancel and Start Build buttons.
func (m Model) viewConfigButtons(width int) string {
	focused := m.section == sectionButtons

	cancelFocused := focused && m.buttonIndex == 0
	startFocused := focused && m.buttonIndex == 1

	cancelBtn := components.TuiButton("Cancel", components.ButtonSecondary, "Esc", cancelFocused)
	startBtn := components.TuiButton("Start Build", components.ButtonPrimary, "Enter", startFocused)

	buttons := lipgloss.JoinHorizontal(lipgloss.Top,
		cancelBtn,
		"     ",
		startBtn,
	)

	return lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Render(buttons)
}

// viewConfigHints renders the keyboard hints at the bottom.
func (m Model) viewConfigHints() string {
	hints := components.TuiKeyHints([]components.KeyHint{
		{Key: "Tab", Label: "Switch sections"},
		{Key: "Enter", Label: "Confirm"},
		{Key: "Esc", Label: "Cancel"},
	}, m.width)

	return lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(hints)
}
