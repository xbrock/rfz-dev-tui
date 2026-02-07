package app

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"
)

// fixedTime returns a deterministic time for golden file tests.
func fixedTime() time.Time {
	return time.Date(2026, 1, 15, 14, 30, 0, 0, time.UTC)
}

// initModel creates a model with fixed time and the given terminal size,
// ready for rendering.
func initModel(width, height int) Model {
	m := New()
	m.currentTime = fixedTime()
	updated, _ := m.Update(tea.WindowSizeMsg{Width: width, Height: height})
	return updated.(Model)
}

// sendKey sends a key message to the model and returns the updated model.
func sendKey(m Model, key string) Model {
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)})
	return updated.(Model)
}

func TestApp_WelcomeDefault(t *testing.T) {
	m := initModel(120, 40)
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_NavBuildFocused(t *testing.T) {
	m := initModel(120, 40)
	// Cursor starts at index 0 (Build Components) by default
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_NavLogsFocused(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "j") // Move cursor down to Logs (index 1)
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_NavDiscoverFocused(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "j") // index 1
	m = sendKey(m, "j") // index 2 (Discover)
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_NavConfigFocused(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "j") // index 1
	m = sendKey(m, "j") // index 2
	m = sendKey(m, "j") // index 3 (Configuration)
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_NavExitFocused(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "j") // index 1
	m = sendKey(m, "j") // index 2
	m = sendKey(m, "j") // index 3
	m = sendKey(m, "j") // index 4 (Exit)
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_BuildScreen(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "1") // Navigate to Build Components
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_PlaceholderLogs(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "2") // Navigate to View Logs
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_PlaceholderDiscover(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "3") // Navigate to Discover
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_PlaceholderConfig(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "4") // Navigate to Configuration
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_ExitModal(t *testing.T) {
	m := initModel(120, 40)
	m = sendKey(m, "q") // Open quit confirmation modal
	golden.RequireEqual(t, []byte(m.View()))
}

func TestApp_TerminalTooSmall(t *testing.T) {
	m := initModel(60, 15)
	golden.RequireEqual(t, []byte(m.View()))
}
