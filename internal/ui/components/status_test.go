package components_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"rfz-cli/internal/ui/components"
)

func TestTuiStatus_Pending(t *testing.T) {
	output := components.TuiStatus(components.StatusPending)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatus_Running(t *testing.T) {
	output := components.TuiStatus(components.StatusRunning)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatus_Success(t *testing.T) {
	output := components.TuiStatus(components.StatusSuccess)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatus_Failed(t *testing.T) {
	output := components.TuiStatus(components.StatusFailed)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatus_Error(t *testing.T) {
	output := components.TuiStatus(components.StatusError)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatusCompact_Pending(t *testing.T) {
	output := components.TuiStatusCompact(components.StatusPending)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatusCompact_Running(t *testing.T) {
	output := components.TuiStatusCompact(components.StatusRunning)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatusCompact_Success(t *testing.T) {
	output := components.TuiStatusCompact(components.StatusSuccess)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatusCompact_Failed(t *testing.T) {
	output := components.TuiStatusCompact(components.StatusFailed)
	golden.RequireEqual(t, []byte(output))
}

func TestTuiStatusCompact_Error(t *testing.T) {
	output := components.TuiStatusCompact(components.StatusError)
	golden.RequireEqual(t, []byte(output))
}

func TestStatus_String(t *testing.T) {
	tests := []struct {
		status   components.Status
		expected string
	}{
		{components.StatusPending, "PENDING"},
		{components.StatusRunning, "RUNNING"},
		{components.StatusSuccess, "SUCCESS"},
		{components.StatusFailed, "FAILED"},
		{components.StatusError, "ERROR"},
		{components.Status(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if got := tt.status.String(); got != tt.expected {
			t.Errorf("Status(%d).String() = %q, want %q", tt.status, got, tt.expected)
		}
	}
}
