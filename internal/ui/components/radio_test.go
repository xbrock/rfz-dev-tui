package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

// Single TuiRadio tests

func TestTuiRadio_Unselected(t *testing.T) {
	output := components.TuiRadio("clean", false, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadio_Selected(t *testing.T) {
	output := components.TuiRadio("clean", true, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadio_Unselected_Focused(t *testing.T) {
	output := components.TuiRadio("clean", false, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadio_Selected_Focused(t *testing.T) {
	output := components.TuiRadio("clean", true, true)
	golden.RequireEqual(t, []byte(output))
}

// TuiRadioGroup tests - Horizontal Layout

func TestTuiRadioGroup_Horizontal_NoneSelected(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, -1, -1, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Horizontal_SecondSelected(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, 1, -1, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Horizontal_SecondFocused(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, -1, 1, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Horizontal_SelectedAndFocused(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, 1, 1, true)
	golden.RequireEqual(t, []byte(output))
}

// TuiRadioGroup tests - Vertical Layout

func TestTuiRadioGroup_Vertical_NoneSelected(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, -1, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Vertical_SecondSelected(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, 1, -1, false)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Vertical_ThirdFocused(t *testing.T) {
	options := []string{"clean", "install", "package"}
	output := components.TuiRadioGroup(options, -1, 2, false)
	golden.RequireEqual(t, []byte(output))
}

// Edge cases

func TestTuiRadioGroup_SingleOption(t *testing.T) {
	options := []string{"default"}
	output := components.TuiRadioGroup(options, 0, 0, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_Empty(t *testing.T) {
	options := []string{}
	output := components.TuiRadioGroup(options, -1, -1, true)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiRadioGroup_FourOptions_Vertical(t *testing.T) {
	options := []string{"debug", "info", "warn", "error"}
	output := components.TuiRadioGroup(options, 1, 2, false)
	golden.RequireEqual(t, []byte(output))
}
