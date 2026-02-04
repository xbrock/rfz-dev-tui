package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiCheckbox_Unchecked(t *testing.T) {
	output := components.TuiCheckbox("Skip Tests", false, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_Checked(t *testing.T) {
	output := components.TuiCheckbox("Skip Tests", true, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_Unchecked_Focused(t *testing.T) {
	output := components.TuiCheckbox("Enable Debug", false, true, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_Checked_Focused(t *testing.T) {
	output := components.TuiCheckbox("Enable Debug", true, true, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_Disabled(t *testing.T) {
	output := components.TuiCheckbox("Premium Feature", false, false, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_Checked_Disabled(t *testing.T) {
	output := components.TuiCheckbox("Premium Feature", true, false, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_LongLabel(t *testing.T) {
	output := components.TuiCheckbox("Generate local configuration files for development environment", false, false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiCheckbox_LongLabel_Focused(t *testing.T) {
	output := components.TuiCheckbox("Generate local configuration files for development environment", false, true, false)
	golden.RequireEqual(t, []byte(output))
}
