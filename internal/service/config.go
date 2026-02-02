// Package service contains application services that orchestrate business logic.
//
// This file defines the ConfigService which handles application configuration,
// including loading, saving, and validating settings.
package service

import (
	tea "github.com/charmbracelet/bubbletea"
)

// AppConfig represents the application configuration.
type AppConfig struct {
	// MavenPath is the path to the Maven executable.
	MavenPath string `json:"maven_path" yaml:"maven_path"`

	// WorkspacePath is the root workspace directory.
	WorkspacePath string `json:"workspace_path" yaml:"workspace_path"`

	// DefaultProfiles is the list of default Maven profiles.
	DefaultProfiles []string `json:"default_profiles" yaml:"default_profiles"`

	// DefaultSkipTests indicates whether to skip tests by default.
	DefaultSkipTests bool `json:"default_skip_tests" yaml:"default_skip_tests"`

	// DefaultCleanBuild indicates whether to do clean builds by default.
	DefaultCleanBuild bool `json:"default_clean_build" yaml:"default_clean_build"`

	// Parallelism is the default parallelism level.
	Parallelism int `json:"parallelism" yaml:"parallelism"`

	// RegistryPath is the path to the component registry.
	RegistryPath string `json:"registry_path" yaml:"registry_path"`
}

// DefaultConfig returns the default application configuration.
func DefaultConfig() AppConfig {
	return AppConfig{
		MavenPath:         "/usr/bin/mvn",
		WorkspacePath:     "~/workspace",
		DefaultProfiles:   []string{"base"},
		DefaultSkipTests:  false,
		DefaultCleanBuild: true,
		Parallelism:       4,
		RegistryPath:      "~/.rfz/registry.yaml",
	}
}

// ConfigService manages application configuration.
//
// The service handles:
// - Loading configuration from file
// - Saving configuration to file
// - Validating configuration values
// - Providing default values
type ConfigService struct {
	config     AppConfig
	configPath string
	loaded     bool
}

// NewConfigService creates a new ConfigService.
func NewConfigService() *ConfigService {
	return &ConfigService{
		config:     DefaultConfig(),
		configPath: "~/.rfz/config.yaml",
	}
}

// Load loads the configuration from disk.
func (s *ConfigService) Load() tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement configuration loading
		// 1. Read config file
		// 2. Parse YAML/JSON
		// 3. Merge with defaults

		s.loaded = true
		return ConfigLoadedMsg{Config: s.config}
	}
}

// Save saves the configuration to disk.
func (s *ConfigService) Save() tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement configuration saving
		// 1. Serialize config to YAML/JSON
		// 2. Write to config file

		return ConfigSavedMsg{Success: true}
	}
}

// Get returns the current configuration.
func (s *ConfigService) Get() AppConfig {
	return s.config
}

// Set updates the configuration.
func (s *ConfigService) Set(config AppConfig) {
	s.config = config
}

// SetMavenPath updates the Maven path.
func (s *ConfigService) SetMavenPath(path string) {
	s.config.MavenPath = path
}

// SetWorkspacePath updates the workspace path.
func (s *ConfigService) SetWorkspacePath(path string) {
	s.config.WorkspacePath = path
}

// SetProfiles updates the default profiles.
func (s *ConfigService) SetProfiles(profiles []string) {
	s.config.DefaultProfiles = profiles
}

// SetSkipTests updates the default skip tests setting.
func (s *ConfigService) SetSkipTests(skip bool) {
	s.config.DefaultSkipTests = skip
}

// SetParallelism updates the parallelism level.
func (s *ConfigService) SetParallelism(p int) {
	if p < 1 {
		p = 1
	}
	if p > 8 {
		p = 8
	}
	s.config.Parallelism = p
}

// Validate validates the current configuration.
func (s *ConfigService) Validate() tea.Cmd {
	return func() tea.Msg {
		errors := make([]string, 0)

		// TODO: Implement validation
		// 1. Check Maven executable exists
		// 2. Check workspace path exists
		// 3. Check parallelism is in valid range

		if len(errors) > 0 {
			return ConfigValidationMsg{
				Valid:  false,
				Errors: errors,
			}
		}

		return ConfigValidationMsg{Valid: true}
	}
}

// IsLoaded returns whether the configuration has been loaded.
func (s *ConfigService) IsLoaded() bool {
	return s.loaded
}

// Config-related messages

// ConfigLoadedMsg is sent when configuration is loaded.
type ConfigLoadedMsg struct {
	Config AppConfig
}

// ConfigSavedMsg is sent when configuration is saved.
type ConfigSavedMsg struct {
	Success bool
}

// ConfigValidationMsg is sent with validation results.
type ConfigValidationMsg struct {
	Valid  bool
	Errors []string
}
