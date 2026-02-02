package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiDivider_Single_Short(t *testing.T) {
	output := components.TuiDivider(components.DividerSingle, 20)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiDivider_Single_Long(t *testing.T) {
	output := components.TuiDivider(components.DividerSingle, 80)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiDivider_Double_Short(t *testing.T) {
	output := components.TuiDivider(components.DividerDouble, 20)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiDivider_Double_Long(t *testing.T) {
	output := components.TuiDivider(components.DividerDouble, 80)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiDivider_Zero_Width(t *testing.T) {
	output := components.TuiDivider(components.DividerSingle, 0)
	if output != "" {
		t.Errorf("Expected empty string for width=0, got %q", output)
	}
}
