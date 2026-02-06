package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestStatusBar_FullWidth(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Status:      "Build: Running",
		StatusColor: components.ColorYellow,
		Info:        "Build",
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
		},
		Width: 120,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_LeftStatus(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Status:      "Build: Running",
		StatusColor: components.ColorYellow,
		Width:       80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_CenterInfo(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Info:  "Build",
		Width: 80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_RightKeyHints(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
			{Key: "q", Label: "Quit"},
		},
		Width: 80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_Separator(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Status:      "Success",
		StatusColor: components.ColorGreen,
		Info:        "Config",
		Hints: []components.KeyHint{
			{Key: "q", Label: "Quit"},
		},
		Width: 60,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_KeyHintsIntegration(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
		},
		Width: 120,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_Empty(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Width: 80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_LongStatusTruncation(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Status:      "Building component-registry with Maven 3.9.5 - Phase compile running",
		StatusColor: components.ColorYellow,
		Info:        "Build Dashboard",
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
			{Key: "q", Label: "Quit"},
		},
		Width: 80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_ZeroWidth(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Status: "Test",
		Width:  0,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}
