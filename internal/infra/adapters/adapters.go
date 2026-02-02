// Package adapters provides implementations of the port interfaces.
//
// Adapters are concrete implementations that connect to external systems:
// - Real adapters: Actual Maven, Git, and filesystem operations
// - Mock adapters: Fake implementations for testing
//
// This file re-exports the port interfaces for convenience and provides
// type aliases used by the main package.
package adapters

import "rfz-cli/internal/infra/ports"

// Re-export port interfaces as type aliases.
// This allows the main package to import only the adapters package.
type (
	// MavenExecutor is the interface for Maven build operations.
	MavenExecutor = ports.MavenExecutor

	// GitClient is the interface for Git operations.
	GitClient = ports.GitClient

	// FileSystem is the interface for filesystem operations.
	FileSystem = ports.FileSystem
)

// Re-export port types.
type (
	// GitStatus represents Git repository status.
	GitStatus = ports.GitStatus

	// FileInfo represents file metadata.
	FileInfo = ports.FileInfo

	// PomMetadata represents parsed pom.xml data.
	PomMetadata = ports.PomMetadata

	// MavenExecutorConfig holds Maven configuration.
	MavenExecutorConfig = ports.MavenExecutorConfig

	// GitClientConfig holds Git configuration.
	GitClientConfig = ports.GitClientConfig
)

// Re-export configuration functions.
var (
	DefaultMavenConfig = ports.DefaultMavenConfig
	DefaultGitConfig   = ports.DefaultGitConfig
	WithMavenPath      = ports.WithMavenPath
	WithJavaHome       = ports.WithJavaHome
	WithTimeout        = ports.WithTimeout
)
