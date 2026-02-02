// Package ports defines interfaces (ports) for external dependencies.
//
// Ports are the boundaries between our application and external systems.
// They define contracts that adapters must implement.
//
// This file defines the MavenExecutor interface for Maven build operations.
package ports

import "rfz-cli/internal/domain"

// MavenExecutor defines the interface for executing Maven builds.
//
// This port abstracts the Maven CLI so that:
// - Production code can use real Maven execution
// - Tests can use mock implementations for predictable behavior
// - Different Maven versions can be supported via different adapters
type MavenExecutor interface {
	// Execute runs a Maven build with the given configuration.
	//
	// This is a blocking operation that runs until the build completes.
	// For streaming output, use ExecuteWithProgress.
	Execute(config domain.BuildConfig, componentPath string) (*domain.BuildResult, error)

	// ExecuteWithProgress runs a Maven build and streams progress.
	//
	// The returned channel receives log entries as they are produced.
	// The channel is closed when the build completes.
	ExecuteWithProgress(config domain.BuildConfig, componentPath string) (<-chan domain.LogEntry, error)

	// Cancel attempts to cancel a running build.
	//
	// Returns an error if the build cannot be cancelled or is not running.
	Cancel(buildID string) error

	// IsAvailable checks if Maven is available on the system.
	//
	// Returns true if Maven can be executed, false otherwise.
	IsAvailable() bool

	// Version returns the Maven version string.
	//
	// Returns an error if Maven is not available.
	Version() (string, error)
}

// MavenExecutorOption is a functional option for configuring MavenExecutor.
type MavenExecutorOption func(*MavenExecutorConfig)

// MavenExecutorConfig holds configuration for MavenExecutor.
type MavenExecutorConfig struct {
	// MavenPath is the path to the Maven executable.
	MavenPath string

	// JavaHome is the JAVA_HOME environment variable.
	JavaHome string

	// Timeout is the maximum build duration in seconds.
	Timeout int

	// Environment contains additional environment variables.
	Environment map[string]string
}

// DefaultMavenConfig returns the default MavenExecutor configuration.
func DefaultMavenConfig() MavenExecutorConfig {
	return MavenExecutorConfig{
		MavenPath:   "/usr/bin/mvn",
		Timeout:     3600, // 1 hour
		Environment: make(map[string]string),
	}
}

// WithMavenPath sets the Maven executable path.
func WithMavenPath(path string) MavenExecutorOption {
	return func(c *MavenExecutorConfig) {
		c.MavenPath = path
	}
}

// WithJavaHome sets the JAVA_HOME environment variable.
func WithJavaHome(path string) MavenExecutorOption {
	return func(c *MavenExecutorConfig) {
		c.JavaHome = path
	}
}

// WithTimeout sets the build timeout in seconds.
func WithTimeout(seconds int) MavenExecutorOption {
	return func(c *MavenExecutorConfig) {
		c.Timeout = seconds
	}
}
