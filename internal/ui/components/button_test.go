package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiButton_Primary(t *testing.T) {
	output := components.TuiButton("Build", components.ButtonPrimary, "", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Primary_Focused(t *testing.T) {
	output := components.TuiButton("Build", components.ButtonPrimary, "", true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Primary_Shortcut(t *testing.T) {
	output := components.TuiButton("Build", components.ButtonPrimary, "Enter", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Primary_Shortcut_Focused(t *testing.T) {
	output := components.TuiButton("Build", components.ButtonPrimary, "Enter", true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Secondary(t *testing.T) {
	output := components.TuiButton("Cancel", components.ButtonSecondary, "", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Secondary_Focused(t *testing.T) {
	output := components.TuiButton("Cancel", components.ButtonSecondary, "", true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Secondary_Shortcut(t *testing.T) {
	output := components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Secondary_Shortcut_Focused(t *testing.T) {
	output := components.TuiButton("Cancel", components.ButtonSecondary, "Esc", true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Destructive(t *testing.T) {
	output := components.TuiButton("Delete", components.ButtonDestructive, "", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Destructive_Focused(t *testing.T) {
	output := components.TuiButton("Delete", components.ButtonDestructive, "", true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Destructive_Shortcut(t *testing.T) {
	output := components.TuiButton("Delete", components.ButtonDestructive, "D", false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiButton_Destructive_Shortcut_Focused(t *testing.T) {
	output := components.TuiButton("Delete", components.ButtonDestructive, "D", true)
	golden.RequireEqual(t, []byte(output))
}
