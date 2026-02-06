package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTabs_Single(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
	}
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_MultipleTabs(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs"},
		{Label: "Discover"},
		{Label: "Config"},
	}
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_NumericShortcuts(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs"},
		{Label: "Discover"},
		{Label: "Config"},
	}
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_FocusedTab(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs"},
		{Label: "Discover"},
		{Label: "Config"},
	}
	// Focus on "Logs" (index 1)
	output := components.TuiTabs(tabs, 1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_WithBadge(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs", Badge: 5},
		{Label: "Config"},
	}
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_Empty(t *testing.T) {
	output := components.TuiTabs(nil, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_WidthTruncation(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs"},
		{Label: "Discover"},
		{Label: "Config"},
	}
	// Width 25 should cut off some tabs
	output := components.TuiTabs(tabs, -1, 25)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_ShortcutLimit(t *testing.T) {
	tabs := make([]components.TuiTab, 12)
	for i := range tabs {
		tabs[i] = components.TuiTab{Label: "Tab"}
	}
	tabs[0].Active = true
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_FocusAndActive(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build", Active: true},
		{Label: "Logs"},
		{Label: "Config"},
	}
	// Active is "Build" (index 0), focus on "Config" (index 2)
	output := components.TuiTabs(tabs, 2, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestTabs_MultipleBadges(t *testing.T) {
	tabs := []components.TuiTab{
		{Label: "Build"},
		{Label: "Logs", Badge: 12},
		{Label: "Errors", Badge: 3, Active: true},
		{Label: "Config"},
	}
	output := components.TuiTabs(tabs, -1, 0)
	golden.RequireEqual(t, []byte(output))
}
