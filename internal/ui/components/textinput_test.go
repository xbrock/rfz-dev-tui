package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiTextInput_Empty(t *testing.T) {
	m := components.NewTuiTextInput("Enter port number", "")
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_Empty_Focused(t *testing.T) {
	m := components.NewTuiTextInput("Enter port number", "")
	m.Focus()
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_WithValue(t *testing.T) {
	m := components.NewTuiTextInput("Enter port number", "")
	m.SetValue("11090")
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_WithValue_Focused(t *testing.T) {
	m := components.NewTuiTextInput("Enter port number", "")
	m.SetValue("11090")
	m.Focus()
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_WithPrompt(t *testing.T) {
	m := components.NewTuiTextInput("Enter command", "$")
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_WithPrompt_Focused(t *testing.T) {
	m := components.NewTuiTextInput("Enter command", "$")
	m.Focus()
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_WithPrompt_Value(t *testing.T) {
	m := components.NewTuiTextInput("Enter command", "$")
	m.SetValue("mvn clean install")
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_Disabled(t *testing.T) {
	m := components.NewTuiTextInput("Enter path", "")
	m.SetValue("/usr/local/bin")
	m.SetDisabled(true)
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_CharLimit(t *testing.T) {
	m := components.NewTuiTextInput("Max 10 chars", "")
	m.SetCharLimit(10)
	m.SetValue("1234567890")
	output := m.View()
	golden.RequireEqual(t, []byte(output))
}

func TestTuiTextInput_ValueMethods(t *testing.T) {
	m := components.NewTuiTextInput("Test", "")

	// Test SetValue and Value
	m.SetValue("hello")
	if m.Value() != "hello" {
		t.Errorf("Expected value 'hello', got '%s'", m.Value())
	}

	// Test CharLimit
	m.SetCharLimit(5)
	if m.CharLimit() != 5 {
		t.Errorf("Expected char limit 5, got %d", m.CharLimit())
	}

	// Test Disabled
	m.SetDisabled(true)
	if !m.Disabled() {
		t.Error("Expected disabled to be true")
	}

	// Test Focused after disable (should blur)
	if m.Focused() {
		t.Error("Expected not focused after disable")
	}
}
