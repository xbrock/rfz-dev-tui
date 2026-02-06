// Package components provides shared UI components and styles.
//
// This file contains the TuiNavigation sidebar container and TuiNavItem menu item.
package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// maxNavItemLabelWidth is the maximum width for nav item labels before truncation.
const maxNavItemLabelWidth = 30

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
	// Build the left part: cursor + number + label
	var leftParts []string

	if cursor {
		cursorStyle := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
		leftParts = append(leftParts, cursorStyle.Render(SymbolListPointer))
	}

	// Number prefix
	numberStr := fmt.Sprintf("%d", item.Number)

	// Determine label width budget for truncation
	labelWidth := maxNavItemLabelWidth
	if width > 0 {
		// Account for: cursor(2) + number(len) + dot(1) + space(1) + shortcut(len) + spacing(2)
		overhead := 4 + len(numberStr)
		if item.Shortcut != "" {
			overhead += len(item.Shortcut) + 2
		}
		if !cursor {
			overhead += 2 // padding for non-cursor items
		}
		available := width - overhead
		if available > 0 && available < labelWidth {
			labelWidth = available
		}
	}

	label := Truncate(item.Label, labelWidth)

	// Build the content based on state
	switch {
	case cursor && focused:
		// Focused cursor: cyan, bold
		style := StyleNavItemFocused
		leftParts = append(leftParts, style.Render(fmt.Sprintf("%s. %s", numberStr, label)))
	case active:
		// Active screen: background highlight
		text := fmt.Sprintf("%s. %s", numberStr, label)
		if item.Shortcut != "" {
			shortcutStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
			// Pad between label and shortcut
			if width > 0 {
				padding := width - lipgloss.Width(text) - len(item.Shortcut) - 4
				if padding < 1 {
					padding = 1
				}
				text += strings.Repeat(" ", padding) + shortcutStyle.Render(item.Shortcut)
			} else {
				text += "  " + shortcutStyle.Render(item.Shortcut)
			}
		}
		return StyleNavItemActive.Render(text)
	default:
		// Normal item
		leftParts = append(leftParts, StyleNavItem.Render(fmt.Sprintf("%s. %s", numberStr, label)))
	}

	result := strings.Join(leftParts, " ")

	// Add shortcut (right side) for non-active items
	if item.Shortcut != "" && !active {
		shortcutStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		result += "  " + shortcutStyle.Render(item.Shortcut)
	}

	return result
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
