package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// TuiModal tests

func TestModal_Basic(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Build Configuration",
		Content: "Select your build options below.",
		Width:   60,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_WithButtons(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Build Configuration",
		Content: "Ready to start the build?",
		Buttons: []components.TuiModalButton{
			{Label: "Cancel", Variant: components.ButtonSecondary, Shortcut: "Esc"},
			{Label: "Start Build", Variant: components.ButtonPrimary, Shortcut: "Enter"},
		},
		Width:        60,
		FocusedIndex: -1,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_ButtonFocused(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Build Configuration",
		Content: "Ready to start the build?",
		Buttons: []components.TuiModalButton{
			{Label: "Cancel", Variant: components.ButtonSecondary, Shortcut: "Esc"},
			{Label: "Start Build", Variant: components.ButtonPrimary, Shortcut: "Enter"},
		},
		Width:        60,
		FocusedIndex: 1,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_NoButtons(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Information",
		Content: "This is an informational modal with no actions.",
		Width:   50,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_NoTitle(t *testing.T) {
	config := components.TuiModalConfig{
		Content: "Content without a title.",
		Buttons: []components.TuiModalButton{
			{Label: "OK", Variant: components.ButtonPrimary, Shortcut: "Enter"},
		},
		Width:        40,
		FocusedIndex: 0,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_WithBackdrop(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Confirm",
		Content: "Are you sure?",
		Buttons: []components.TuiModalButton{
			{Label: "No", Variant: components.ButtonSecondary, Shortcut: "Esc"},
			{Label: "Yes", Variant: components.ButtonPrimary, Shortcut: "Enter"},
		},
		Width:        40,
		FocusedIndex: 1,
	}
	output := components.TuiModal(config, 120, 40)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_DefaultWidth(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Default Width Modal",
		Content: "This uses the default width of 60.",
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_WithHeight(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Fixed Height",
		Content: "Short content in a tall modal.",
		Width:   50,
		Height:  20,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_DestructiveButton(t *testing.T) {
	config := components.TuiModalConfig{
		Title:   "Delete Confirmation",
		Content: "This action cannot be undone.",
		Buttons: []components.TuiModalButton{
			{Label: "Cancel", Variant: components.ButtonSecondary, Shortcut: "Esc"},
			{Label: "Delete", Variant: components.ButtonDestructive, Shortcut: "Enter"},
		},
		Width:        50,
		FocusedIndex: 0,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestModal_EmptyContent(t *testing.T) {
	config := components.TuiModalConfig{
		Title: "Empty Modal",
		Buttons: []components.TuiModalButton{
			{Label: "Close", Variant: components.ButtonSecondary, Shortcut: "Esc"},
		},
		Width:        40,
		FocusedIndex: 0,
	}
	output := components.TuiModal(config, 0, 0)
	golden.RequireEqual(t, []byte(output))
}
