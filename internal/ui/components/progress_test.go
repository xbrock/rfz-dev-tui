package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// Static progress bar tests

func TestTuiProgress_Zero(t *testing.T) {
	output := components.TuiProgress(0, 30, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_Zero_Percent(t *testing.T) {
	output := components.TuiProgress(0, 30, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_Half(t *testing.T) {
	output := components.TuiProgress(0.5, 30, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_Half_Percent(t *testing.T) {
	output := components.TuiProgress(0.5, 30, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_ThreeQuarter(t *testing.T) {
	output := components.TuiProgress(0.75, 30, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_Full(t *testing.T) {
	output := components.TuiProgress(1.0, 30, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_Full_Percent(t *testing.T) {
	output := components.TuiProgress(1.0, 30, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiProgress_CustomWidth(t *testing.T) {
	output := components.TuiProgress(0.5, 40, true)
	golden.RequireEqual(t, []byte(output))
}

// Model tests

func TestTuiProgressModel_Init(t *testing.T) {
	m := components.NewTuiProgress(30, true)
	cmd := m.Init()
	// Progress doesn't need to tick like spinner, so no initial command
	if cmd != nil {
		t.Error("Init should return nil for progress bar")
	}
}

func TestTuiProgressModel_SetPercent(t *testing.T) {
	m := components.NewTuiProgress(30, true)

	m.SetPercent(0.5)
	if m.Percent() != 0.5 {
		t.Errorf("Expected percent 0.5, got %f", m.Percent())
	}

	// Test clamping
	m.SetPercent(-0.5)
	if m.Percent() != 0 {
		t.Errorf("Expected percent to be clamped to 0, got %f", m.Percent())
	}

	m.SetPercent(1.5)
	if m.Percent() != 1 {
		t.Errorf("Expected percent to be clamped to 1, got %f", m.Percent())
	}
}

func TestTuiProgressModel_View(t *testing.T) {
	m := components.NewTuiProgress(30, true)
	m.SetPercent(0.5)
	output := m.View()
	if len(output) == 0 {
		t.Error("View should return non-empty string")
	}
}

func TestTuiProgressModel_SetShowPercent(t *testing.T) {
	m := components.NewTuiProgress(30, false)
	m.SetPercent(0.5)

	viewWithout := m.View()
	m.SetShowPercent(true)
	viewWith := m.View()

	// View with percent should be longer
	if len(viewWith) <= len(viewWithout) {
		t.Error("View with percent should be longer than without")
	}
}
