// Package adapters provides implementations of the port interfaces.
//
// This file contains a mock Maven executor for testing.
// It returns predefined results without executing real Maven builds.
package adapters

import (
	"time"

	"rfz-cli/internal/domain"
)

// MockMavenExecutor implements MavenExecutor for testing.
//
// This adapter is used in tests and UI development to provide
// predictable build results without requiring Maven installation.
type MockMavenExecutor struct {
	// Results maps component paths to predefined results.
	Results map[string]*domain.BuildResult

	// SimulateDelay adds artificial delay to simulate build time.
	SimulateDelay time.Duration

	// SimulateProgress determines whether to send progress updates.
	SimulateProgress bool

	// Logs contains predefined log entries to return.
	Logs []domain.LogEntry

	// FailOnPath causes builds to fail for specific paths.
	FailOnPath map[string]error
}

// NewMockMavenExecutor creates a new MockMavenExecutor.
func NewMockMavenExecutor() *MockMavenExecutor {
	return &MockMavenExecutor{
		Results:       make(map[string]*domain.BuildResult),
		FailOnPath:   make(map[string]error),
		SimulateDelay: 100 * time.Millisecond,
	}
}

// SetResult sets a predefined result for a component path.
func (e *MockMavenExecutor) SetResult(componentPath string, result *domain.BuildResult) {
	e.Results[componentPath] = result
}

// SetFailure sets a path to fail with the given error.
func (e *MockMavenExecutor) SetFailure(componentPath string, err error) {
	e.FailOnPath[componentPath] = err
}

// Execute runs a mock build.
func (e *MockMavenExecutor) Execute(config domain.BuildConfig, componentPath string) (*domain.BuildResult, error) {
	// Simulate delay
	if e.SimulateDelay > 0 {
		time.Sleep(e.SimulateDelay)
	}

	// Check for simulated failure
	if err, ok := e.FailOnPath[componentPath]; ok {
		return &domain.BuildResult{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	// Return predefined result if available
	if result, ok := e.Results[componentPath]; ok {
		return result, nil
	}

	// Default: return success
	return &domain.BuildResult{
		Success:   true,
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}, nil
}

// ExecuteWithProgress runs a mock build with progress updates.
func (e *MockMavenExecutor) ExecuteWithProgress(config domain.BuildConfig, componentPath string) (<-chan domain.LogEntry, error) {
	ch := make(chan domain.LogEntry, 100)

	go func() {
		defer close(ch)

		// Send predefined logs
		for _, log := range e.Logs {
			if e.SimulateDelay > 0 {
				time.Sleep(e.SimulateDelay / 10)
			}
			ch <- log
		}

		// If no logs defined, send some default progress
		if len(e.Logs) == 0 {
			phases := []string{"validate", "compile", "test", "package", "install"}
			for _, phase := range phases {
				if e.SimulateDelay > 0 {
					time.Sleep(e.SimulateDelay / 5)
				}
				ch <- domain.LogEntry{
					Level:   domain.LogLevelInfo,
					Message: "Building phase: " + phase,
				}
			}
		}
	}()

	return ch, nil
}

// Cancel is a no-op for the mock.
func (e *MockMavenExecutor) Cancel(buildID string) error {
	return nil
}

// IsAvailable always returns true for the mock.
func (e *MockMavenExecutor) IsAvailable() bool {
	return true
}

// Version returns a mock version string.
func (e *MockMavenExecutor) Version() (string, error) {
	return "3.9.6-mock", nil
}
