// Package ports defines interfaces (ports) for external dependencies.
//
// This file defines the GitClient interface for Git operations.
package ports

// GitStatus represents the Git status of a directory.
type GitStatus struct {
	// Branch is the current branch name.
	Branch string

	// HasChanges indicates if there are uncommitted changes.
	HasChanges bool

	// Ahead is the number of commits ahead of remote.
	Ahead int

	// Behind is the number of commits behind remote.
	Behind int

	// StagedFiles is the list of staged file paths.
	StagedFiles []string

	// ModifiedFiles is the list of modified (unstaged) file paths.
	ModifiedFiles []string

	// UntrackedFiles is the list of untracked file paths.
	UntrackedFiles []string
}

// GitClient defines the interface for Git operations.
//
// This port abstracts Git CLI operations so that:
// - Production code can use real Git execution
// - Tests can use mock implementations
// - Git status can be retrieved without side effects
type GitClient interface {
	// Status returns the Git status for the given directory.
	//
	// Returns an error if the directory is not a Git repository.
	Status(path string) (*GitStatus, error)

	// IsRepository checks if the path is a Git repository.
	IsRepository(path string) bool

	// CurrentBranch returns the current branch name.
	CurrentBranch(path string) (string, error)

	// Fetch fetches from the remote.
	Fetch(path string) error

	// Pull pulls from the remote.
	Pull(path string) error

	// GetRemoteURL returns the URL of the origin remote.
	GetRemoteURL(path string) (string, error)
}

// GitClientConfig holds configuration for GitClient.
type GitClientConfig struct {
	// GitPath is the path to the Git executable.
	GitPath string

	// Timeout is the maximum operation duration in seconds.
	Timeout int
}

// DefaultGitConfig returns the default GitClient configuration.
func DefaultGitConfig() GitClientConfig {
	return GitClientConfig{
		GitPath: "/usr/bin/git",
		Timeout: 30,
	}
}
