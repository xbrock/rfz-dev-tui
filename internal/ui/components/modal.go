// Package components provides shared UI components and styles.
//
// This file contains the TuiModal overlay dialog component.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TuiModalButton represents a button in the modal footer.
type TuiModalButton struct {
	Label    string        // Button label text
	Variant  ButtonVariant // Button style variant
	Shortcut string        // Keyboard shortcut hint
}

// TuiModalConfig holds the configuration for rendering a modal dialog.
type TuiModalConfig struct {
	Title        string           // Modal title displayed in the header
	Content      string           // Main content of the modal
	Buttons      []TuiModalButton // Footer action buttons
	Width        int              // Modal width (0 = auto)
	Height       int              // Modal height (0 = auto)
	FocusedIndex int              // Index of the focused button (-1 for none)
}

// TuiModal renders a centered overlay dialog with double border.
// termWidth and termHeight define the terminal dimensions for centering and backdrop.
func TuiModal(config TuiModalConfig, termWidth, termHeight int) string {
	modalWidth := config.Width
	if modalWidth <= 0 {
		modalWidth = 60
	}
	modalHeight := config.Height

	// Build modal sections
	var sections []string

	// Title section
	if config.Title != "" {
		titleStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan).
			Width(modalWidth - 4). // Account for border + padding
			Align(lipgloss.Center)
		sections = append(sections, titleStyle.Render(config.Title))
		sections = append(sections, TuiDivider(DividerSingle, modalWidth-4))
	}

	// Content section
	if config.Content != "" {
		contentStyle := lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Width(modalWidth - 4)
		sections = append(sections, contentStyle.Render(config.Content))
	}

	// Footer with buttons
	if len(config.Buttons) > 0 {
		sections = append(sections, TuiDivider(DividerSingle, modalWidth-4))

		var buttonParts []string
		for i, btn := range config.Buttons {
			focused := i == config.FocusedIndex
			rendered := TuiButton(btn.Label, btn.Variant, btn.Shortcut, focused)
			buttonParts = append(buttonParts, rendered)
		}

		footerContent := strings.Join(buttonParts, "    ")
		footerStyle := lipgloss.NewStyle().
			Width(modalWidth - 4).
			Align(lipgloss.Center)
		sections = append(sections, footerStyle.Render(footerContent))
	}

	// Join all sections vertically
	innerContent := lipgloss.JoinVertical(lipgloss.Left, sections...)

	// Apply double border modal style
	modalStyle := lipgloss.NewStyle().
		Border(BorderDouble).
		BorderForeground(ColorCyan).
		Padding(1, 1).
		Width(modalWidth)

	if modalHeight > 0 {
		modalStyle = modalStyle.Height(modalHeight)
	}

	modal := modalStyle.Render(innerContent)

	// If terminal dimensions provided, center with backdrop
	if termWidth > 0 && termHeight > 0 {
		return renderWithBackdrop(modal, termWidth, termHeight)
	}

	return modal
}

// renderWithBackdrop places the modal centered on a dimmed backdrop.
func renderWithBackdrop(modal string, termWidth, termHeight int) string {
	return lipgloss.Place(
		termWidth,
		termHeight,
		lipgloss.Center,
		lipgloss.Center,
		modal,
		lipgloss.WithWhitespaceBackground(ColorBackground),
		lipgloss.WithWhitespaceForeground(ColorTextDisabled),
	)
}
