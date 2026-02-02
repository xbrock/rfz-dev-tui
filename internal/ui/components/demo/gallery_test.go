package demo_test

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components/demo"
)

func TestGallery_View_Initial(t *testing.T) {
	g := demo.New()
	// Before receiving WindowSizeMsg, should show initializing
	output := g.View()
	if !strings.Contains(output, "Initializing") {
		t.Errorf("Expected 'Initializing' message before window size, got %q", output)
	}
}

func TestGallery_View_AfterResize(t *testing.T) {
	g := demo.New()

	// Simulate window resize
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	g = model.(demo.Gallery)

	output := g.View()
	golden.RequireEqual(t, []byte(output))
}

func TestGallery_Contains_TuiBox(t *testing.T) {
	g := demo.New()
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	g = model.(demo.Gallery)

	output := g.View()
	if !strings.Contains(output, "TuiBox") {
		t.Error("Gallery should contain TuiBox section")
	}
}

func TestGallery_Contains_TuiDivider(t *testing.T) {
	g := demo.New()
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	g = model.(demo.Gallery)

	output := g.View()
	if !strings.Contains(output, "TuiDivider") {
		t.Error("Gallery should contain TuiDivider section")
	}
}

func TestGallery_Contains_TuiButton(t *testing.T) {
	g := demo.New()
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	g = model.(demo.Gallery)

	output := g.View()
	if !strings.Contains(output, "TuiButton") {
		t.Error("Gallery should contain TuiButton section")
	}
}

func TestGallery_Contains_TuiStatus(t *testing.T) {
	g := demo.New()
	// Use a larger height to ensure all content is visible
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 100})
	g = model.(demo.Gallery)

	output := g.View()
	if !strings.Contains(output, "TuiStatus") {
		t.Error("Gallery should contain TuiStatus section")
	}
}

func TestGallery_KeyNavigation_Quit(t *testing.T) {
	g := demo.New()
	model, _ := g.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	g = model.(demo.Gallery)

	// Test q key quits
	_, cmd := g.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if cmd == nil {
		t.Error("Expected quit command after pressing 'q'")
	}
}
