// Package components provides shared UI components and styles.
//
// This file contains helper/utility functions for TUI components.
package components

import (
	"github.com/muesli/reflow/truncate"
)

// Truncate shortens a string to maxWidth characters, adding ellipsis if needed.
// If the text is shorter than maxWidth, it is returned unchanged.
// The ellipsis "…" counts toward the maxWidth limit.
func Truncate(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}
	return truncate.StringWithTail(text, uint(maxWidth), "…")
}
