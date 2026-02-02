// Package service contains application services that orchestrate business logic.
//
// This file defines the ScanService which handles component discovery,
// including filesystem scanning and registry management.
package service

import (
	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/domain"
	"rfz-cli/internal/infra/adapters"
)

// ScanService handles component discovery and registry management.
//
// The service:
// - Scans the filesystem for RFZ components
// - Maintains the component registry
// - Resolves component dependencies
type ScanService struct {
	fileSystem adapters.FileSystem
	git        adapters.GitClient
}

// NewScanService creates a new ScanService with the given dependencies.
func NewScanService(fileSystem adapters.FileSystem, git adapters.GitClient) *ScanService {
	return &ScanService{
		fileSystem: fileSystem,
		git:        git,
	}
}

// Scan discovers components in the workspace.
//
// Returns a tea.Cmd that scans asynchronously.
// The command will send ScanCompleteMsg or ScanErrorMsg.
func (s *ScanService) Scan(workspacePath string) tea.Cmd {
	return func() tea.Msg {
		if s.fileSystem == nil {
			return ScanErrorMsg{
				Err:     errNoFileSystem,
				Context: "Scan service not properly initialized",
			}
		}

		// TODO: Implement component scanning
		// 1. Walk the workspace directory
		// 2. Find pom.xml files
		// 3. Parse component metadata
		// 4. Detect component type
		// 5. Build dependency graph

		// Placeholder: return empty list
		return ScanCompleteMsg{
			Components: []domain.Component{},
		}
	}
}

// ScanWithFilter discovers components matching the given filter.
//
// Filter examples: "core/*", "simulator/*", etc.
func (s *ScanService) ScanWithFilter(workspacePath, filter string) tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement filtered scanning
		return ScanCompleteMsg{
			Components: []domain.Component{},
		}
	}
}

// GetGitStatus returns the Git status for a component.
//
// This is useful for showing which components have uncommitted changes.
func (s *ScanService) GetGitStatus(component domain.Component) tea.Cmd {
	return func() tea.Msg {
		if s.git == nil {
			return GitStatusMsg{
				ComponentID: component.ID,
				HasChanges:  false,
			}
		}

		// TODO: Implement Git status check
		// Call git.Status() for the component path

		return GitStatusMsg{
			ComponentID: component.ID,
			HasChanges:  false,
		}
	}
}

// LoadRegistry loads the component registry from disk.
func (s *ScanService) LoadRegistry(registryPath string) tea.Cmd {
	return func() tea.Msg {
		if s.fileSystem == nil {
			return RegistryLoadedMsg{
				Components: []domain.Component{},
			}
		}

		// TODO: Implement registry loading
		// 1. Read registry file (YAML/JSON)
		// 2. Parse component entries
		// 3. Validate component paths exist

		return RegistryLoadedMsg{
			Components: []domain.Component{},
		}
	}
}

// SaveRegistry saves the component registry to disk.
func (s *ScanService) SaveRegistry(registryPath string, components []domain.Component) tea.Cmd {
	return func() tea.Msg {
		if s.fileSystem == nil {
			return RegistrySavedMsg{Success: false}
		}

		// TODO: Implement registry saving
		// 1. Serialize components to YAML/JSON
		// 2. Write to registry file

		return RegistrySavedMsg{Success: true}
	}
}

// AddToRegistry adds a component to the registry.
func (s *ScanService) AddToRegistry(component domain.Component) tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement adding to registry
		return ComponentAddedMsg{Component: component}
	}
}

// RemoveFromRegistry removes a component from the registry.
func (s *ScanService) RemoveFromRegistry(componentID string) tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement removing from registry
		return ComponentRemovedMsg{ComponentID: componentID}
	}
}

// Scan-related messages

// ScanCompleteMsg is sent when component scanning completes.
type ScanCompleteMsg struct {
	Components []domain.Component
}

// ScanErrorMsg is sent when scanning encounters an error.
type ScanErrorMsg struct {
	Err     error
	Context string
}

func (e ScanErrorMsg) Error() string {
	if e.Context != "" {
		return e.Context + ": " + e.Err.Error()
	}
	return e.Err.Error()
}

// GitStatusMsg reports Git status for a component.
type GitStatusMsg struct {
	ComponentID  string
	HasChanges   bool
	ChangedFiles []string
	Branch       string
}

// RegistryLoadedMsg is sent when the registry is loaded.
type RegistryLoadedMsg struct {
	Components []domain.Component
}

// RegistrySavedMsg is sent when the registry is saved.
type RegistrySavedMsg struct {
	Success bool
}

// ComponentAddedMsg is sent when a component is added to the registry.
type ComponentAddedMsg struct {
	Component domain.Component
}

// ComponentRemovedMsg is sent when a component is removed from the registry.
type ComponentRemovedMsg struct {
	ComponentID string
}

// Sentinel errors
var (
	errNoFileSystem = &serviceError{"file system not configured"}
)
