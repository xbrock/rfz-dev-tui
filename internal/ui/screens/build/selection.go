package build

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// viewSelection renders the component selection phase.
func (m Model) viewSelection() string {
	title := components.StyleH2.Render("Build")

	// Header section with title and key hints
	selectedCount := len(components.GetSelected(m.items))
	totalCount := len(m.items)
	subtitle := fmt.Sprintf("Select components to build (%d/%d selected)", selectedCount, totalCount)
	subtitleRendered := components.StyleBodySecondary.Render(subtitle)

	shortcutHints := components.TuiKeyHints([]components.KeyHint{
		{Key: "Space", Label: "Toggle"},
		{Key: "a", Label: "All"},
		{Key: "n", Label: "None"},
	}, m.width)

	headerLine := lipgloss.JoinHorizontal(lipgloss.Top,
		subtitleRendered,
		lipgloss.NewStyle().Width(m.width-lipgloss.Width(subtitleRendered)-lipgloss.Width(shortcutHints)).Render(""),
		shortcutHints,
	)

	// Component list
	listContent := components.TuiList(m.items, m.cursorIndex, components.ListMultiSelect, m.focused, false)

	listBox := lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorCyan).
		Padding(0, 1).
		Width(m.width).
		Render(
			components.StyleH3.Render("Build RFZ Components") + "\n" +
				headerLine + "\n\n" +
				listContent,
		)

	// Actions section
	actions := m.viewActions()

	// Legend
	legend := m.viewLegend()

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		listBox,
		"",
		actions,
		"",
		legend,
	)
}

// viewActions renders the action buttons section.
func (m Model) viewActions() string {
	hasSelection := len(components.GetSelected(m.items)) > 0

	var buildVariant components.ButtonVariant
	if hasSelection {
		buildVariant = components.ButtonPrimary
	} else {
		buildVariant = components.ButtonSecondary
	}

	buildBtn := components.TuiButton("Build Selected", buildVariant, "Enter", false)
	selectAllBtn := components.TuiButton("Select All", components.ButtonSecondary, "a", false)
	clearBtn := components.TuiButton("Clear Selection", components.ButtonSecondary, "n", false)

	tabHint := components.TuiKeyHints([]components.KeyHint{
		{Key: "Tab", Label: "Switch focus"},
	}, 0)

	buttons := lipgloss.JoinHorizontal(lipgloss.Top,
		buildBtn,
		"    ",
		selectAllBtn,
		"    ",
		clearBtn,
	)

	// Right-align the tab hint
	buttonsWidth := lipgloss.Width(buttons)
	tabHintWidth := lipgloss.Width(tabHint)
	gapWidth := m.width - buttonsWidth - tabHintWidth - 6 // account for box padding/border
	if gapWidth < 1 {
		gapWidth = 1
	}

	actionsContent := lipgloss.JoinHorizontal(lipgloss.Top,
		buttons,
		lipgloss.NewStyle().Width(gapWidth).Render(""),
		tabHint,
	)

	return lipgloss.NewStyle().
		Border(components.BorderRounded).
		BorderForeground(components.ColorBorder).
		BorderTop(true).
		BorderBottom(true).
		BorderLeft(true).
		BorderRight(true).
		Padding(0, 1).
		Width(m.width).
		Render(
			components.StyleH3.Render("Actions") + "\n" +
				actionsContent,
		)
}

// viewLegend renders the selection legend.
func (m Model) viewLegend() string {
	keyStyle := lipgloss.NewStyle().Foreground(components.ColorCyan).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(components.ColorTextSecondary)

	legend := lipgloss.JoinHorizontal(lipgloss.Top,
		keyStyle.Render("[x]"),
		" ",
		labelStyle.Render("Selected"),
		"    ",
		keyStyle.Render("[ ]"),
		" ",
		labelStyle.Render("Not selected"),
		"    ",
		keyStyle.Render(">"),
		" ",
		labelStyle.Render("Current"),
	)

	return legend
}
