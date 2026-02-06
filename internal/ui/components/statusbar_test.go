package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestStatusBar_FullWidth(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ModeBadge:      "BUILD",
		ModeBadgeColor: components.ColorYellow,
		ContextBadge:   "Build",
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
		},
		QuitHint: &components.KeyHint{Key: "q", Label: "Quit"},
		Width:    120,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_ModeBadgeOnly(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ModeBadge:      "BUILD",
		ModeBadgeColor: components.ColorYellow,
		Width:          80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_ContextBadgeOnly(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ContextBadge: "Build",
		Width:        80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_RightKeyHints(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
		},
		QuitHint: &components.KeyHint{Key: "q", Label: "Quit"},
		Width:    80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_BadgesAndHints(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ModeBadge:      "LOGS",
		ModeBadgeColor: components.ColorGreen,
		ContextBadge:   "Config",
		Hints: []components.KeyHint{
			{Key: "Tab", Label: "Focus"},
		},
		QuitHint: &components.KeyHint{Key: "q", Label: "Quit"},
		Width:    60,
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

func TestStatusBar_LongBadgeTruncation(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ModeBadge:      "SELECT",
		ModeBadgeColor: components.ColorCyan,
		ContextBadge:   "rfz-dispatcher",
		Hints: []components.KeyHint{
			{Key: "Enter", Label: "Build"},
			{Key: "Esc", Label: "Cancel"},
		},
		QuitHint: &components.KeyHint{Key: "q", Label: "Quit"},
		Width:    80,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}

func TestStatusBar_ZeroWidth(t *testing.T) {
	cfg := components.TuiStatusBarConfig{
		ModeBadge: "TEST",
		Width:     0,
	}
	output := components.TuiStatusBar(cfg)
	golden.RequireEqual(t, []byte(output))
}
