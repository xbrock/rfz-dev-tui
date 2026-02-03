// Package demo provides demo screens for showcasing UI components.
package demo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"rfz-cli/internal/ui/components"
)

// Gallery is a Bubble Tea model that displays all UI components in sections.
type Gallery struct {
	viewport viewport.Model
	width    int
	height   int
	ready    bool
}

// New creates a new Gallery model.
func New() Gallery {
	return Gallery{}
}

// Init implements tea.Model.
func (g Gallery) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (g Gallery) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		g.width = msg.Width
		g.height = msg.Height

		if !g.ready {
			g.viewport = viewport.New(msg.Width, msg.Height-2)
			g.viewport.SetContent(g.renderContent())
			g.ready = true
		} else {
			g.viewport.Width = msg.Width
			g.viewport.Height = msg.Height - 2
			g.viewport.SetContent(g.renderContent())
		}
		return g, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return g, tea.Quit
		case "j", "down":
			g.viewport.LineDown(1)
		case "k", "up":
			g.viewport.LineUp(1)
		}
	}

	g.viewport, cmd = g.viewport.Update(msg)
	return g, cmd
}

// View implements tea.Model.
func (g Gallery) View() string {
	if !g.ready {
		return "Initializing..."
	}

	header := lipgloss.NewStyle().
		Foreground(components.ColorCyan).
		Bold(true).
		Render("Component Gallery") +
		"  " +
		lipgloss.NewStyle().
			Foreground(components.ColorTextMuted).
			Render("(j/k to scroll, q to quit)")

	return header + "\n" + g.viewport.View()
}

// renderContent generates the full gallery content.
func (g Gallery) renderContent() string {
	var sections []string

	sections = append(sections, g.renderBoxSection())
	sections = append(sections, g.renderDividerSection())
	sections = append(sections, g.renderButtonSection())
	sections = append(sections, g.renderStatusSection())

	return strings.Join(sections, "\n\n")
}

// renderBoxSection renders the TuiBox demo section.
func (g Gallery) renderBoxSection() string {
	var sb strings.Builder

	sb.WriteString(components.StyleH2.Render("TuiBox"))
	sb.WriteString("\n\n")

	// Box variants
	sb.WriteString("Border Variants:\n")
	boxRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		components.TuiBox("Single Border", components.BoxSingle, false),
		"  ",
		components.TuiBox("Double Border", components.BoxDouble, false),
		"  ",
		components.TuiBox("Rounded Border", components.BoxRounded, false),
		"  ",
		components.TuiBox("Heavy Border", components.BoxHeavy, false),
	)
	sb.WriteString(boxRow)
	sb.WriteString("\n\n")

	sb.WriteString("Focus State:\n")
	focusRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		components.TuiBox("Normal", components.BoxSingle, false),
		"  ",
		components.TuiBox("Focused", components.BoxSingle, true),
	)
	sb.WriteString(focusRow)

	return sb.String()
}

// renderDividerSection renders the TuiDivider demo section.
func (g Gallery) renderDividerSection() string {
	var sb strings.Builder

	sb.WriteString(components.StyleH2.Render("TuiDivider"))
	sb.WriteString("\n\n")

	sb.WriteString("Single Line:\n")
	sb.WriteString(components.TuiDivider(components.DividerSingle, 40))
	sb.WriteString("\n\n")

	sb.WriteString("Double Line:\n")
	sb.WriteString(components.TuiDivider(components.DividerDouble, 40))

	return sb.String()
}

// renderButtonSection renders the TuiButton demo section.
func (g Gallery) renderButtonSection() string {
	var sb strings.Builder

	sb.WriteString(components.StyleH2.Render("TuiButton"))
	sb.WriteString("\n\n")

	sb.WriteString("Button Variants:\n")
	variantRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		components.TuiButton("Primary", components.ButtonPrimary, "", false),
		"  ",
		components.TuiButton("Secondary", components.ButtonSecondary, "", false),
		"  ",
		components.TuiButton("Destructive", components.ButtonDestructive, "", false),
	)
	sb.WriteString(variantRow)
	sb.WriteString("\n\n")

	sb.WriteString("With Shortcuts:\n")
	shortcutRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		components.TuiButton("Build", components.ButtonPrimary, "Enter", false),
		"  ",
		components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false),
		"  ",
		components.TuiButton("Delete", components.ButtonDestructive, "D", false),
	)
	sb.WriteString(shortcutRow)
	sb.WriteString("\n\n")

	sb.WriteString("Focus State:\n")
	focusBtnRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		components.TuiButton("Primary", components.ButtonPrimary, "", true),
		"  ",
		components.TuiButton("Secondary", components.ButtonSecondary, "", true),
		"  ",
		components.TuiButton("Destructive", components.ButtonDestructive, "", true),
	)
	sb.WriteString(focusBtnRow)

	return sb.String()
}

// renderStatusSection renders the TuiStatus demo section.
func (g Gallery) renderStatusSection() string {
	var sb strings.Builder

	sb.WriteString(components.StyleH2.Render("TuiStatus"))
	sb.WriteString("\n\n")

	sb.WriteString("Status Badges:\n")
	statuses := []components.Status{
		components.StatusPending,
		components.StatusRunning,
		components.StatusSuccess,
		components.StatusFailed,
		components.StatusError,
	}
	for _, s := range statuses {
		sb.WriteString(components.TuiStatus(s))
		sb.WriteString("  ")
	}
	sb.WriteString("\n\n")

	sb.WriteString("Compact Variants:\n")
	for _, s := range statuses {
		sb.WriteString(components.TuiStatusCompact(s))
		sb.WriteString(fmt.Sprintf(" %s  ", s.String()))
	}

	return sb.String()
}
