package components_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestKeyHints_Single(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "q", Label: "Quit"},
	}
	output := components.TuiKeyHints(hints, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_Multiple(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "Enter", Label: "Select"},
		{Key: "Esc", Label: "Cancel"},
		{Key: "q", Label: "Quit"},
	}
	output := components.TuiKeyHints(hints, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_Empty(t *testing.T) {
	output := components.TuiKeyHints(nil, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_WidthTruncation(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "Enter", Label: "Select"},
		{Key: "Esc", Label: "Cancel"},
		{Key: "q", Label: "Quit"},
		{Key: "?", Label: "Help"},
		{Key: "Tab", Label: "Next"},
	}
	// Width 30 should only fit first two hints
	output := components.TuiKeyHints(hints, 30)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_ContextAware(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "Enter", Label: "Build"},
		{Key: "Space", Label: "Toggle"},
		{Key: "a", Label: "Select All"},
		{Key: "q", Label: "Quit"},
	}
	output := components.TuiKeyHints(hints, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_TwoItems(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "j/k", Label: "Navigate"},
		{Key: "Enter", Label: "Select"},
	}
	output := components.TuiKeyHints(hints, 0)
	golden.RequireEqual(t, []byte(output))
}

func TestKeyHints_NoDotSeparator(t *testing.T) {
	hints := []components.KeyHint{
		{Key: "Enter", Label: "Select"},
		{Key: "Esc", Label: "Cancel"},
		{Key: "q", Label: "Quit"},
	}
	output := components.TuiKeyHints(hints, 0)
	if strings.Contains(output, "·") {
		t.Error("expected no middle-dot separator, but found '·' in output")
	}
	if strings.Contains(output, " · ") {
		t.Error("expected no dot separator, but found ' · ' in output")
	}
	golden.RequireEqual(t, []byte(output))
}
