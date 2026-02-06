// Package components provides shared UI components and styles.
//
// This file contains the TuiTable wrapper around bubbles/table with RFZ styling.
package components

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// TuiTableConfig configures a TuiTable instance.
type TuiTableConfig struct {
	Columns     []table.Column // Column definitions
	Rows        []table.Row    // Data rows
	Width       int            // Total table width (0 = auto from columns)
	Height      int            // Visible row count (default 10)
	Focused     bool           // Whether the table is focused/interactive
	ZebraStripe bool           // Alternate row background colors
}

// TuiTableStyles returns pre-configured table.Styles with RFZ design tokens.
func TuiTableStyles() table.Styles {
	return table.Styles{
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorTextSecondary).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(ColorBorder).
			Padding(0, 1),
		Cell: lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Padding(0, 1),
		Selected: lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Background(ColorSecondary).
			Bold(true).
			Padding(0, 1),
	}
}

// NewTuiTable creates a new bubbles/table.Model with RFZ design system styling.
// Returns a configured table ready for use in a Bubble Tea model.
func NewTuiTable(cfg TuiTableConfig) table.Model {
	height := cfg.Height
	if height <= 0 {
		height = 10
	}

	opts := []table.Option{
		table.WithColumns(cfg.Columns),
		table.WithRows(cfg.Rows),
		table.WithHeight(height),
		table.WithStyles(TuiTableStyles()),
	}

	if cfg.Focused {
		opts = append(opts, table.WithFocused(true))
	}

	t := table.New(opts...)

	if cfg.Width > 0 {
		t.SetWidth(cfg.Width)
	}

	return t
}

// TuiTableEmpty renders an empty table placeholder with headers visible.
// columns: column definitions (rendered as header row)
// width: total width for the container
func TuiTableEmpty(columns []table.Column, width int) string {
	// Render header row
	headerParts := make([]string, 0, len(columns))
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorTextSecondary).
		Padding(0, 1)

	for _, col := range columns {
		cellStyle := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width).Inline(true)
		headerParts = append(headerParts, headerStyle.Render(cellStyle.Render(Truncate(col.Title, col.Width))))
	}

	header := lipgloss.JoinHorizontal(lipgloss.Left, headerParts...)

	// Separator line
	separatorStyle := lipgloss.NewStyle().Foreground(ColorBorder)
	totalWidth := 0
	for _, col := range columns {
		totalWidth += col.Width + 2 // +2 for padding
	}
	if width > 0 && width < totalWidth {
		totalWidth = width
	}
	separator := separatorStyle.Render(repeatChar("─", totalWidth))

	// "No data" message
	emptyStyle := lipgloss.NewStyle().
		Foreground(ColorTextMuted).
		Italic(true).
		Padding(1, 1)
	emptyMsg := emptyStyle.Render("No data")

	return header + "\n" + separator + "\n" + emptyMsg
}

// repeatChar repeats a character n times.
func repeatChar(ch string, n int) string {
	if n <= 0 {
		return ""
	}
	result := make([]byte, 0, n*len(ch))
	for i := 0; i < n; i++ {
		result = append(result, ch...)
	}
	return string(result)
}
