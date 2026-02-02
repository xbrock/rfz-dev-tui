// Package service contains application services that orchestrate business logic.
//
// This file defines the BuildService which handles build orchestration,
// including component selection, Maven execution, and progress tracking.
package service

import (
	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/infra/adapters"
)

// BuildService orchestrates Maven build operations.
//
// The service:
// - Receives ports (interfaces) via constructor injection
// - Encapsulates business logic for building components
// - Returns tea.Cmd for async operations (integrates with Bubble Tea)
// - Never accesses external systems directly (always through ports)
type BuildService struct {
	maven   adapters.MavenExecutor
	scanner adapters.FileSystem
}

// NewBuildService creates a new BuildService with the given dependencies.
//
// In production, pass real implementations (RealMavenExecutor, RealFileSystem).
// In testing, pass mock implementations for predictable behavior.
func NewBuildService(maven adapters.MavenExecutor, scanner adapters.FileSystem) *BuildService {
	return &BuildService{
		maven:   maven,
		scanner: scanner,
	}
}

// Build starts a build for the given configuration.
//
// Returns a tea.Cmd that executes the build asynchronously.
// The command will send BuildProgressMsg and BuildCompleteMsg messages.
func (s *BuildService) Build(config domain.BuildConfig) tea.Cmd {
	return func() tea.Msg {
		if s.maven == nil {
			return BuildErrorMsg{
				Err:     errNoMavenExecutor,
				Context: "Build service not properly initialized",
			}
		}

		// TODO: Implement actual build execution
		// 1. Resolve component build order (dependencies first)
		// 2. For each component, call maven.Execute()
		// 3. Stream progress messages via channel
		// 4. Return final result

		// Placeholder: return success
		return BuildCompleteMsg{
			Result: domain.BuildResult{
				Success: true,
			},
		}
	}
}

// BuildWithProgress starts a build and returns a channel for progress updates.
//
// This allows the UI to receive real-time build progress.
func (s *BuildService) BuildWithProgress(config domain.BuildConfig) (<-chan BuildProgressMsg, tea.Cmd) {
	progress := make(chan BuildProgressMsg, 100)

	cmd := func() tea.Msg {
		defer close(progress)

		// TODO: Implement streaming build execution
		// Send progress updates via the channel
		// Example:
		// progress <- BuildProgressMsg{ComponentID: "comp1", Phase: "compile", Percent: 50}

		return BuildCompleteMsg{
			Result: domain.BuildResult{
				Success: true,
			},
		}
	}

	return progress, cmd
}

// CancelBuild attempts to cancel a running build.
func (s *BuildService) CancelBuild(buildID string) tea.Cmd {
	return func() tea.Msg {
		if s.maven == nil {
			return BuildErrorMsg{
				Err:     errNoMavenExecutor,
				Context: "Cannot cancel build",
			}
		}

		// TODO: Implement build cancellation
		// Call maven.Cancel() if supported

		return BuildCancelledMsg{BuildID: buildID}
	}
}

// ResolveBuildOrder determines the order in which components should be built.
//
// Components are ordered based on their dependencies (dependency graph traversal).
// Returns an error if there are circular dependencies.
func (s *BuildService) ResolveBuildOrder(components []domain.Component) ([]domain.Component, error) {
	// TODO: Implement topological sort for build order
	// 1. Build dependency graph
	// 2. Detect cycles (error if found)
	// 3. Return components in build order

	// Placeholder: return as-is
	return components, nil
}

// Build-related messages

// BuildProgressMsg reports build progress for a component.
type BuildProgressMsg struct {
	ComponentID string
	Phase       string // compile, test, package, install
	Percent     int    // 0-100
	Message     string
}

// BuildCompleteMsg is sent when a build finishes.
type BuildCompleteMsg struct {
	Result domain.BuildResult
}

// BuildErrorMsg is sent when a build encounters an error.
type BuildErrorMsg struct {
	Err     error
	Context string
}

func (e BuildErrorMsg) Error() string {
	if e.Context != "" {
		return e.Context + ": " + e.Err.Error()
	}
	return e.Err.Error()
}

// BuildCancelledMsg is sent when a build is cancelled.
type BuildCancelledMsg struct {
	BuildID string
}

// Sentinel errors
var (
	errNoMavenExecutor = &serviceError{"maven executor not configured"}
)

type serviceError struct {
	message string
}

func (e *serviceError) Error() string {
	return e.message
}
