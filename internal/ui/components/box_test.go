package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiBox_Single(t *testing.T) {
	output := components.TuiBox("Content", components.BoxSingle, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Single_Focused(t *testing.T) {
	output := components.TuiBox("Content", components.BoxSingle, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Double(t *testing.T) {
	output := components.TuiBox("Content", components.BoxDouble, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Double_Focused(t *testing.T) {
	output := components.TuiBox("Content", components.BoxDouble, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Rounded(t *testing.T) {
	output := components.TuiBox("Content", components.BoxRounded, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Rounded_Focused(t *testing.T) {
	output := components.TuiBox("Content", components.BoxRounded, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Heavy(t *testing.T) {
	output := components.TuiBox("Content", components.BoxHeavy, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBox_Heavy_Focused(t *testing.T) {
	output := components.TuiBox("Content", components.BoxHeavy, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBoxWithWidth_Truncation(t *testing.T) {
	longContent := "This is a very long content that should be truncated"
	output := components.TuiBoxWithWidth(longContent, components.BoxSingle, false, 30)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiBoxWithWidth_Empty(t *testing.T) {
	output := components.TuiBoxWithWidth("", components.BoxSingle, false, 20)
	golden.RequireEqual(t, []byte(output))
}
