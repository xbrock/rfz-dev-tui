// Package adapters provides implementations of the port interfaces.
//
// This file contains the real Git client implementation that
// executes actual Git commands via the command line.
package adapters

import "rfz-cli/internal/infra/ports"

// RealGitClient implements GitClient using actual Git CLI.
//
// This adapter is used in production for real Git operations.
type RealGitClient struct {
	config ports.GitClientConfig
}

// NewRealGitClient creates a new RealGitClient.
func NewRealGitClient() *RealGitClient {
	return &RealGitClient{
		config: ports.DefaultGitConfig(),
	}
}

// Status returns the Git status for the given directory.
func (c *RealGitClient) Status(path string) (*ports.GitStatus, error) {
	// TODO: Implement real Git status
	//
	// Implementation steps:
	// 1. Run "git status --porcelain" in the directory
	// 2. Parse output to get file statuses
	// 3. Run "git rev-list --count @{upstream}..HEAD" for ahead count
	// 4. Run "git rev-list --count HEAD..@{upstream}" for behind count
	// 5. Run "git branch --show-current" for branch name
	//
	// Example:
	// cmd := exec.Command(c.config.GitPath, "status", "--porcelain")
	// cmd.Dir = path
	// output, err := cmd.Output()
	// ...

	return &ports.GitStatus{
		Branch:     "main",
		HasChanges: false,
	}, nil
}

// IsRepository checks if the path is a Git repository.
func (c *RealGitClient) IsRepository(path string) bool {
	// TODO: Implement repository check
	//
	// Check if .git directory exists or run "git rev-parse --git-dir"
	return true // Placeholder
}

// CurrentBranch returns the current branch name.
func (c *RealGitClient) CurrentBranch(path string) (string, error) {
	// TODO: Implement branch retrieval
	//
	// Run "git branch --show-current"
	return "main", nil
}

// Fetch fetches from the remote.
func (c *RealGitClient) Fetch(path string) error {
	// TODO: Implement fetch
	//
	// Run "git fetch"
	return nil
}

// Pull pulls from the remote.
func (c *RealGitClient) Pull(path string) error {
	// TODO: Implement pull
	//
	// Run "git pull"
	return nil
}

// GetRemoteURL returns the URL of the origin remote.
func (c *RealGitClient) GetRemoteURL(path string) (string, error) {
	// TODO: Implement remote URL retrieval
	//
	// Run "git remote get-url origin"
	return "https://github.com/example/repo.git", nil
}
