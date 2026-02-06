package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// TuiNavItemRender tests

func TestNavItem_Normal(t *testing.T) {
	item := components.TuiNavItem{Label: "Build", Number: 1, Shortcut: "1"}
	output := components.TuiNavItemRender(item, false, false, false, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavItem_Cursor(t *testing.T) {
	item := components.TuiNavItem{Label: "Logs", Number: 3, Shortcut: "3"}
	output := components.TuiNavItemRender(item, true, false, false, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavItem_Cursor_Focused(t *testing.T) {
	item := components.TuiNavItem{Label: "Logs", Number: 3, Shortcut: "3"}
	output := components.TuiNavItemRender(item, true, false, true, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavItem_Active(t *testing.T) {
	item := components.TuiNavItem{Label: "Build", Number: 1, Shortcut: "1"}
	output := components.TuiNavItemRender(item, false, true, false, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavItem_NoShortcut(t *testing.T) {
	item := components.TuiNavItem{Label: "Settings", Number: 5, Shortcut: ""}
	output := components.TuiNavItemRender(item, false, false, false, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavItem_LongLabel(t *testing.T) {
	item := components.TuiNavItem{Label: "Build Components and Deploy to Production Server", Number: 1, Shortcut: "1"}
	output := components.TuiNavItemRender(item, false, false, false, 0)
	golden.RequireEqual(t, []byte(output))
}

// TuiNavigation tests

func TestNavigation_Empty(t *testing.T) {
	output := components.TuiNavigation(nil, -1, -1, false, "", "", 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavigation_Basic(t *testing.T) {
	items := []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "1"},
		{Label: "Build", Number: 2, Shortcut: "2"},
		{Label: "Logs", Number: 3, Shortcut: "3"},
		{Label: "Config", Number: 4, Shortcut: "4"},
	}
	output := components.TuiNavigation(items, -1, 0, false, "", "", 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavigation_WithCursor(t *testing.T) {
	items := []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "1"},
		{Label: "Build", Number: 2, Shortcut: "2"},
		{Label: "Logs", Number: 3, Shortcut: "3"},
		{Label: "Config", Number: 4, Shortcut: "4"},
	}
	output := components.TuiNavigation(items, 2, 0, true, "", "", 0)
	golden.RequireEqual(t, []byte(output))
}

func TestNavigation_WithHeader(t *testing.T) {
	items := []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "1"},
		{Label: "Build", Number: 2, Shortcut: "2"},
	}
	output := components.TuiNavigation(items, -1, 0, false, "Navigation", "", 20)
	golden.RequireEqual(t, []byte(output))
}

func TestNavigation_WithFooter(t *testing.T) {
	items := []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "1"},
		{Label: "Build", Number: 2, Shortcut: "2"},
	}
	output := components.TuiNavigation(items, -1, 0, false, "", "j/k navigate", 20)
	golden.RequireEqual(t, []byte(output))
}

func TestNavigation_WithHeaderAndFooter(t *testing.T) {
	items := []components.TuiNavItem{
		{Label: "Welcome", Number: 1, Shortcut: "1"},
		{Label: "Build", Number: 2, Shortcut: "2"},
		{Label: "Logs", Number: 3, Shortcut: "3"},
	}
	output := components.TuiNavigation(items, 1, 0, true, "Navigation", "j/k navigate", 25)
	golden.RequireEqual(t, []byte(output))
}
