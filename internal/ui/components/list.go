// Package components provides shared UI components and styles.
//
// This file contains the TuiList scrollable selection component.
package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ListSelectMode defines the selection behavior of a list.
type ListSelectMode string

const (
	// ListMultiSelect allows multiple items to be selected (checkboxes).
	ListMultiSelect ListSelectMode = "multi"

	// ListSingleSelect allows only one item to be selected (radio buttons).
	ListSingleSelect ListSelectMode = "single"
)

// maxListItemLabelWidth is the maximum width for list item labels before truncation.
const maxListItemLabelWidth = 50

// TuiListItem represents a single item in a list.
type TuiListItem struct {
	Label    string // Display text for the item
	Badge    string // Optional category badge (e.g., "Core", "Plugin")
	Selected bool   // Whether the item is selected
}

// TuiListItemRender renders a single list item.
// item: the item to render
// cursor: whether the cursor is on this item
// mode: selection mode (multi or single)
// focused: whether this item is focused (cursor + list focused)
func TuiListItemRender(item TuiListItem, cursor bool, mode ListSelectMode, focused bool) string {
	var parts []string

	// Add cursor indicator
	cursorStyle := lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	if cursor {
		parts = append(parts, cursorStyle.Render(SymbolListPointer))
	} else {
		parts = append(parts, " ")
	}

	// Add selection symbol based on mode
	switch mode {
	case ListMultiSelect:
		if item.Selected {
			parts = append(parts, SymbolCheckboxChecked)
		} else {
			parts = append(parts, SymbolCheckboxUnchecked)
		}
	case ListSingleSelect:
		if item.Selected {
			parts = append(parts, SymbolRadioSelected)
		} else {
			parts = append(parts, SymbolRadioUnselected)
		}
	}

	// Truncate label if needed
	label := Truncate(item.Label, maxListItemLabelWidth)

	// Apply style based on focus/selection state
	var labelStyle lipgloss.Style
	switch {
	case cursor && focused:
		labelStyle = lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	case item.Selected:
		labelStyle = lipgloss.NewStyle().Foreground(ColorTextPrimary)
	default:
		labelStyle = lipgloss.NewStyle().Foreground(ColorTextSecondary)
	}
	parts = append(parts, labelStyle.Render(label))

	// Add badge if present
	if item.Badge != "" {
		badgeStyle := lipgloss.NewStyle().
			Background(ColorSecondary).
			Foreground(ColorTextMuted).
			Padding(0, 1)
		parts = append(parts, badgeStyle.Render(item.Badge))
	}

	return strings.Join(parts, " ")
}

// TuiList renders a complete list with all items.
// items: slice of items to render
// cursorIndex: index of the current cursor position (-1 for no cursor)
// mode: selection mode (multi or single)
// focused: whether the list is focused
// showCounter: whether to show selection counter ("3/13 selected")
func TuiList(items []TuiListItem, cursorIndex int, mode ListSelectMode, focused bool, showCounter bool) string {
	if len(items) == 0 {
		emptyStyle := lipgloss.NewStyle().Foreground(ColorTextMuted).Italic(true)
		return emptyStyle.Render("No items")
	}

	var lines []string

	// Render each item
	for i, item := range items {
		isCursor := i == cursorIndex
		line := TuiListItemRender(item, isCursor, mode, focused)
		lines = append(lines, line)
	}

	// Add selection counter if requested
	if showCounter {
		selectedCount := 0
		for _, item := range items {
			if item.Selected {
				selectedCount++
			}
		}
		counterStyle := lipgloss.NewStyle().Foreground(ColorTextMuted)
		counter := counterStyle.Render(fmt.Sprintf("%d/%d selected", selectedCount, len(items)))
		lines = append(lines, "", counter)
	}

	return strings.Join(lines, "\n")
}

// TuiListBox renders a list inside a bordered box.
// items: slice of items to render
// cursorIndex: index of the current cursor position (-1 for no cursor)
// mode: selection mode (multi or single)
// focused: whether the list is focused (affects border color)
// showCounter: whether to show selection counter
// title: optional title for the box (empty string for no title)
func TuiListBox(items []TuiListItem, cursorIndex int, mode ListSelectMode, focused bool, showCounter bool, title string) string {
	content := TuiList(items, cursorIndex, mode, focused, showCounter)

	// Choose border style based on focus
	var boxStyle lipgloss.Style
	if focused {
		boxStyle = StyleBoxFocused
	} else {
		boxStyle = StyleBoxDefault
	}

	// Add title if provided
	if title != "" {
		titleStyle := lipgloss.NewStyle().Foreground(ColorTextSecondary).Bold(true)
		content = titleStyle.Render(title) + "\n\n" + content
	}

	return boxStyle.Render(content)
}

// TuiListHelpers provides utility functions for list state management.

// ToggleSelection toggles the selection state of an item at the given index.
// For multi-select, it toggles the item.
// For single-select, it deselects all others and selects the given item.
func ToggleSelection(items []TuiListItem, index int, mode ListSelectMode) []TuiListItem {
	if index < 0 || index >= len(items) {
		return items
	}

	result := make([]TuiListItem, len(items))
	copy(result, items)

	switch mode {
	case ListMultiSelect:
		result[index].Selected = !result[index].Selected
	case ListSingleSelect:
		for i := range result {
			result[i].Selected = i == index
		}
	}

	return result
}

// SelectAll selects all items (multi-select mode only).
func SelectAll(items []TuiListItem) []TuiListItem {
	result := make([]TuiListItem, len(items))
	for i, item := range items {
		item.Selected = true
		result[i] = item
	}
	return result
}

// DeselectAll deselects all items.
func DeselectAll(items []TuiListItem) []TuiListItem {
	result := make([]TuiListItem, len(items))
	for i, item := range items {
		item.Selected = false
		result[i] = item
	}
	return result
}

// GetSelected returns indices of all selected items.
func GetSelected(items []TuiListItem) []int {
	var selected []int
	for i, item := range items {
		if item.Selected {
			selected = append(selected, i)
		}
	}
	return selected
}

// GetSelectedLabels returns labels of all selected items.
func GetSelectedLabels(items []TuiListItem) []string {
	var labels []string
	for _, item := range items {
		if item.Selected {
			labels = append(labels, item.Label)
		}
	}
	return labels
}
