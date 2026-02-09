// Package components provides shared UI components and styles.
//
// This file contains the TuiNavigation sidebar container and TuiNavItem menu item.
package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TuiNavItem represents a single navigation menu item.
type TuiNavItem struct {
	Label    string // Display text
	Number   int    // Item number (1-based, shown before label)
	Shortcut string // Keyboard shortcut (shown right-aligned, muted)
}

// TuiNavItemRender renders a single navigation item.
// item: the nav item to render
// cursor: whether the cursor is on this item
// active: whether this is the currently active/selected screen
// focused: whether the navigation container is focused
// width: available width for the item (0 = no width constraint)
func TuiNavItemRender(item TuiNavItem, cursor bool, active bool, focused bool, width int) string {
	numberStr := fmt.Sprintf("%d", item.Number)
	label := item.Label

	// Build the text content: "N. Label"
	text := fmt.Sprintf("%s. %s", numberStr, label)

	// Shortcut style
	shortcutStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)

	if active {
		// Active state: teal background, bold, shortcut right-aligned
		// If cursor is also on this item, show ">" prefix
		prefix := ""
		if cursor {
			prefix = SymbolListPointer + " "
		}

		fullText := prefix + text

		// Build right-aligned line with shortcut
		if item.Shortcut != "" && width > 0 {
			shortcutRendered := shortcutStyle.
				Background(ColorNavActiveBg).
				Render(item.Shortcut)
			textWidth := lipgloss.Width(fullText)
			shortcutWidth := lipgloss.Width(item.Shortcut)
			gap := width - textWidth - shortcutWidth
			if gap < 1 {
				gap = 1
			}
			fullText += strings.Repeat(" ", gap) + shortcutRendered
		}

		style := StyleNavItemActive.
			PaddingLeft(0). // we handle prefix manually
			PaddingRight(0).
			Width(width)
		return style.Render(fullText)
	}

	if cursor && focused {
		// Focused cursor (not active): cyan bold with ">" prefix
		cursorPrefix := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true).Render(SymbolListPointer)
		styledText := StyleNavItemFocused.Render(text)

		line := cursorPrefix + " " + styledText
		if item.Shortcut != "" && width > 0 {
			lineWidth := lipgloss.Width(line)
			shortcutWidth := lipgloss.Width(item.Shortcut)
			gap := width - lineWidth - shortcutWidth
			if gap < 1 {
				gap = 1
			}
			line += strings.Repeat(" ", gap) + shortcutStyle.Render(item.Shortcut)
		}
		return line
	}

	// Normal item: indented, shortcut right-aligned
	indent := "  " // 2 spaces to align with "› " prefix
	line := indent + StyleNavItem.PaddingLeft(0).Render(text)
	if item.Shortcut != "" && width > 0 {
		lineWidth := lipgloss.Width(line)
		shortcutWidth := lipgloss.Width(item.Shortcut)
		gap := width - lineWidth - shortcutWidth
		if gap < 1 {
			gap = 1
		}
		line += strings.Repeat(" ", gap) + shortcutStyle.Render(item.Shortcut)
	}
	return line
}

// TuiNavigation renders a complete sidebar navigation with optional header/footer.
// items: slice of nav items to render
// cursorIndex: index of the current cursor position (-1 for no cursor)
// activeIndex: index of the currently active screen (-1 for none)
// focused: whether the navigation is focused
// header: optional header text (empty string for no header)
// footer: optional footer content (empty string for no footer)
// width: available width for the navigation (0 = auto)
func TuiNavigation(items []TuiNavItem, cursorIndex int, activeIndex int, focused bool, header string, footer string, width int) string {
	if len(items) == 0 {
		emptyStyle := lipgloss.NewStyle().Foreground(ColorTextMuted).Italic(true)
		return emptyStyle.Render("No navigation items")
	}

	var sections []string

	// Header
	if header != "" {
		headerStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary).Bold(true)
		sections = append(sections, headerStyle.Render(header))
		sections = append(sections, TuiDivider(DividerSingle, width))
	}

	// Nav items
	var itemLines []string
	for i, item := range items {
		isCursor := i == cursorIndex
		isActive := i == activeIndex
		line := TuiNavItemRender(item, isCursor, isActive, focused, width)
		itemLines = append(itemLines, line)
	}
	sections = append(sections, strings.Join(itemLines, "\n"))

	// Footer
	if footer != "" {
		sections = append(sections, TuiDivider(DividerSingle, width))
		sections = append(sections, footer)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
