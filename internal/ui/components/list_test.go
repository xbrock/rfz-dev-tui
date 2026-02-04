package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// TuiListItemRender tests

func TestTuiListItem_Unchecked(t *testing.T) {
	item := components.TuiListItem{Label: "boss", Badge: "", Selected: false}
	output := components.TuiListItemRender(item, false, components.ListMultiSelect, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_Checked(t *testing.T) {
	item := components.TuiListItem{Label: "boss", Badge: "", Selected: true}
	output := components.TuiListItemRender(item, false, components.ListMultiSelect, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_Cursor(t *testing.T) {
	item := components.TuiListItem{Label: "boss", Badge: "", Selected: false}
	output := components.TuiListItemRender(item, true, components.ListMultiSelect, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_Cursor_Focused(t *testing.T) {
	item := components.TuiListItem{Label: "boss", Badge: "", Selected: false}
	output := components.TuiListItemRender(item, true, components.ListMultiSelect, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_WithBadge(t *testing.T) {
	item := components.TuiListItem{Label: "boss", Badge: "Core", Selected: true}
	output := components.TuiListItemRender(item, true, components.ListMultiSelect, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_SingleSelect_Unselected(t *testing.T) {
	item := components.TuiListItem{Label: "clean install", Badge: "", Selected: false}
	output := components.TuiListItemRender(item, false, components.ListSingleSelect, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListItem_SingleSelect_Selected(t *testing.T) {
	item := components.TuiListItem{Label: "clean install", Badge: "", Selected: true}
	output := components.TuiListItemRender(item, false, components.ListSingleSelect, false)
	golden.RequireEqual(t, []byte(output))
}

// TuiList tests

func TestTuiList_Empty(t *testing.T) {
	output := components.TuiList(nil, -1, components.ListMultiSelect, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiList_MultiSelect(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Selected: true},
		{Label: "fistiv", Selected: true},
		{Label: "core-api", Selected: false},
		{Label: "simulator", Selected: false},
	}
	output := components.TuiList(items, 1, components.ListMultiSelect, true, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiList_MultiSelect_WithCounter(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Selected: true},
		{Label: "fistiv", Selected: true},
		{Label: "core-api", Selected: false},
		{Label: "simulator", Selected: true},
	}
	output := components.TuiList(items, 0, components.ListMultiSelect, true, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiList_SingleSelect(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "clean", Selected: false},
		{Label: "install", Selected: true},
		{Label: "package", Selected: false},
	}
	output := components.TuiList(items, 1, components.ListSingleSelect, true, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiList_WithBadges(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Badge: "Core", Selected: true},
		{Label: "simulator", Badge: "Plugin", Selected: false},
		{Label: "docs", Badge: "Docs", Selected: false},
	}
	output := components.TuiList(items, 0, components.ListMultiSelect, true, false)
	golden.RequireEqual(t, []byte(output))
}

// TuiListBox tests

func TestTuiListBox_Focused(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Selected: true},
		{Label: "fistiv", Selected: false},
	}
	output := components.TuiListBox(items, 0, components.ListMultiSelect, true, false, "Components")
	golden.RequireEqual(t, []byte(output))
}

func TestTuiListBox_Unfocused(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Selected: true},
		{Label: "fistiv", Selected: false},
	}
	output := components.TuiListBox(items, 0, components.ListMultiSelect, false, false, "Components")
	golden.RequireEqual(t, []byte(output))
}

// Helper function tests

func TestToggleSelection_MultiSelect(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "a", Selected: false},
		{Label: "b", Selected: false},
		{Label: "c", Selected: false},
	}

	// Toggle item 1
	result := components.ToggleSelection(items, 1, components.ListMultiSelect)
	if !result[1].Selected {
		t.Error("Item 1 should be selected after toggle")
	}
	if result[0].Selected || result[2].Selected {
		t.Error("Other items should remain unselected")
	}

	// Toggle again
	result = components.ToggleSelection(result, 1, components.ListMultiSelect)
	if result[1].Selected {
		t.Error("Item 1 should be deselected after second toggle")
	}
}

func TestToggleSelection_SingleSelect(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "a", Selected: true},
		{Label: "b", Selected: false},
		{Label: "c", Selected: false},
	}

	// Select item 2
	result := components.ToggleSelection(items, 2, components.ListSingleSelect)
	if !result[2].Selected {
		t.Error("Item 2 should be selected")
	}
	if result[0].Selected || result[1].Selected {
		t.Error("Other items should be deselected in single-select mode")
	}
}

func TestSelectAll(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "a", Selected: false},
		{Label: "b", Selected: false},
	}
	result := components.SelectAll(items)
	for i, item := range result {
		if !item.Selected {
			t.Errorf("Item %d should be selected", i)
		}
	}
}

func TestDeselectAll(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "a", Selected: true},
		{Label: "b", Selected: true},
	}
	result := components.DeselectAll(items)
	for i, item := range result {
		if item.Selected {
			t.Errorf("Item %d should be deselected", i)
		}
	}
}

func TestGetSelected(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "a", Selected: true},
		{Label: "b", Selected: false},
		{Label: "c", Selected: true},
	}
	selected := components.GetSelected(items)
	if len(selected) != 2 || selected[0] != 0 || selected[1] != 2 {
		t.Errorf("Expected [0, 2], got %v", selected)
	}
}

func TestGetSelectedLabels(t *testing.T) {
	items := []components.TuiListItem{
		{Label: "boss", Selected: true},
		{Label: "core", Selected: false},
		{Label: "sim", Selected: true},
	}
	labels := components.GetSelectedLabels(items)
	if len(labels) != 2 || labels[0] != "boss" || labels[1] != "sim" {
		t.Errorf("Expected [boss, sim], got %v", labels)
	}
}
