// Package main is the entry point for the RFZ Developer CLI.
//
// This file is responsible for:
// - Creating adapter implementations (real or mock based on environment)
// - Creating services with injected dependencies
// - Initializing the main application model
// - Running the Bubble Tea program
package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"rfz-cli/internal/app"
	"rfz-cli/internal/infra/adapters"
	"rfz-cli/internal/service"
)

func main() {
	// Determine if we should use mock adapters
	// Set RFZ_MOCK=true for UI development without real Maven/Git
	useMocks := os.Getenv("RFZ_MOCK") == "true"

	// Create infrastructure adapters
	// In production: real implementations that call Maven, Git, etc.
	// In development/testing: mock implementations for predictable UI testing
	var (
		mavenExec  adapters.MavenExecutor
		gitClient  adapters.GitClient
		fileSystem adapters.FileSystem
	)

	if useMocks {
		// TODO: Create mock adapters for development
		// mavenExec = adapters.NewMockMavenExecutor()
		// gitClient = adapters.NewMockGitClient()
		// fileSystem = adapters.NewMockFileSystem()
		log.Println("Running with mock adapters")
	} else {
		// TODO: Create real adapters for production
		// mavenExec = adapters.NewRealMavenExecutor("/usr/bin/mvn")
		// gitClient = adapters.NewRealGitClient()
		// fileSystem = adapters.NewRealFileSystem()
		log.Println("Running with real adapters")
	}

	// Create application services with injected dependencies
	// Services encapsulate business logic and use ports for external access
	buildSvc := service.NewBuildService(mavenExec, fileSystem)
	scanSvc := service.NewScanService(fileSystem, gitClient)
	configSvc := service.NewConfigService()

	// Create the main application model
	// The app model contains all screen models and routes messages between them
	model := app.New(buildSvc, scanSvc, configSvc)

	// Configure and run the Bubble Tea program
	// WithAltScreen: Uses alternate screen buffer (restores terminal on exit)
	// WithMouseAllMotion: Optional - enable if mouse support is needed
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
