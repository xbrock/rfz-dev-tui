// Package adapters provides implementations of the port interfaces.
//
// This file contains a mock Git client for testing.
// It returns predefined results without executing real Git commands.
package adapters

import "rfz-cli/internal/infra/ports"

// MockGitClient implements GitClient for testing.
//
// This adapter is used in tests and UI development to provide
// predictable Git status without requiring Git installation.
type MockGitClient struct {
	// Statuses maps paths to predefined status results.
	Statuses map[string]*ports.GitStatus

	// IsRepo maps paths to repository status.
	IsRepo map[string]bool

	// Branches maps paths to branch names.
	Branches map[string]string
}

// NewMockGitClient creates a new MockGitClient.
func NewMockGitClient() *MockGitClient {
	return &MockGitClient{
		Statuses: make(map[string]*ports.GitStatus),
		IsRepo:   make(map[string]bool),
		Branches: make(map[string]string),
	}
}

// SetStatus sets a predefined status for a path.
func (c *MockGitClient) SetStatus(path string, status *ports.GitStatus) {
	c.Statuses[path] = status
}

// SetIsRepository sets whether a path is a repository.
func (c *MockGitClient) SetIsRepository(path string, isRepo bool) {
	c.IsRepo[path] = isRepo
}

// SetBranch sets the branch name for a path.
func (c *MockGitClient) SetBranch(path string, branch string) {
	c.Branches[path] = branch
}

// Status returns the mock status for the given directory.
func (c *MockGitClient) Status(path string) (*ports.GitStatus, error) {
	if status, ok := c.Statuses[path]; ok {
		return status, nil
	}

	// Default: clean repository on main branch
	return &ports.GitStatus{
		Branch:     "main",
		HasChanges: false,
	}, nil
}

// IsRepository returns the mock repository status.
func (c *MockGitClient) IsRepository(path string) bool {
	if isRepo, ok := c.IsRepo[path]; ok {
		return isRepo
	}
	return true // Default: is a repository
}

// CurrentBranch returns the mock branch name.
func (c *MockGitClient) CurrentBranch(path string) (string, error) {
	if branch, ok := c.Branches[path]; ok {
		return branch, nil
	}
	return "main", nil
}

// Fetch is a no-op for the mock.
func (c *MockGitClient) Fetch(path string) error {
	return nil
}

// Pull is a no-op for the mock.
func (c *MockGitClient) Pull(path string) error {
	return nil
}

// GetRemoteURL returns a mock remote URL.
func (c *MockGitClient) GetRemoteURL(path string) (string, error) {
	return "https://github.com/mock/repo.git", nil
}
