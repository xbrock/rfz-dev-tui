package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// Static spinner tests (for golden file comparison)

func TestTuiSpinnerStatic_Braille(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerBraille, "", components.SpinnerColorCyan)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_Line(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerLine, "", components.SpinnerColorCyan)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_Circle(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerCircle, "", components.SpinnerColorCyan)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_Bounce(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerBounce, "", components.SpinnerColorCyan)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_WithLabel(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerBraille, "Compiling...", components.SpinnerColorCyan)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_Green(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerBraille, "Success", components.SpinnerColorGreen)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiSpinnerStatic_Yellow(t *testing.T) {
	output := components.TuiSpinnerStatic(components.SpinnerBraille, "Warning", components.SpinnerColorYellow)
	golden.RequireEqual(t, []byte(output))
}

// Model behavior tests

func TestTuiSpinnerModel_Init(t *testing.T) {
	m := components.NewTuiSpinner(components.SpinnerBraille, "Loading")
	cmd := m.Init()
	if cmd == nil {
		t.Error("Init should return a Tick command")
	}
}

func TestTuiSpinnerModel_Label(t *testing.T) {
	m := components.NewTuiSpinner(components.SpinnerBraille, "Loading")
	if m.Label() != "Loading" {
		t.Errorf("Expected label 'Loading', got '%s'", m.Label())
	}

	m.SetLabel("Processing")
	if m.Label() != "Processing" {
		t.Errorf("Expected label 'Processing', got '%s'", m.Label())
	}
}

func TestTuiSpinnerModel_View(t *testing.T) {
	m := components.NewTuiSpinner(components.SpinnerBraille, "Building...")
	output := m.View()
	// Should contain the label
	if len(output) == 0 {
		t.Error("View should return non-empty string")
	}
}

func TestTuiSpinnerModel_SetColor(t *testing.T) {
	m := components.NewTuiSpinner(components.SpinnerBraille, "Test")
	// Should not panic
	m.SetColor(components.SpinnerColorGreen)
	m.SetColor(components.SpinnerColorYellow)
	m.SetColor(components.SpinnerColorCyan)
}
