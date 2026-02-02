# Architecture Structure: RFZ Developer CLI

> Last Updated: 2026-02-02
> Version: 1.0.0
> Source: Boilerplate Implementation

## Overview

This document describes the directory structure and file organization for the RFZ Developer CLI boilerplate. The structure follows the Layered Architecture with Ports & Adapters pattern defined in `architecture-decision.md`.

## Directory Structure

```
boilerplate/
├── cmd/
│   └── rfz-cli/
│       └── main.go              # Application entry point
│
├── internal/
│   ├── app/                     # Application core
│   │   ├── app.go               # Main Bubble Tea model
│   │   ├── messages.go          # Shared message types
│   │   └── keymap.go            # Global key bindings
│   │
│   ├── ui/                      # PRESENTATION LAYER
│   │   ├── screens/             # Full-screen views
│   │   │   ├── welcome/
│   │   │   │   └── welcome.go   # Welcome screen (init + state)
│   │   │   ├── build/
│   │   │   │   └── build.go     # Build screen (component selection)
│   │   │   ├── logs/
│   │   │   │   └── logs.go      # Log viewer screen
│   │   │   ├── discover/
│   │   │   │   └── discover.go  # Component discovery screen
│   │   │   └── config/
│   │   │       └── config.go    # Configuration screen
│   │   │
│   │   ├── modals/              # Overlay dialogs
│   │   │   ├── buildconfig/
│   │   │   │   └── buildconfig.go
│   │   │   └── confirm/
│   │   │       └── confirm.go
│   │   │
│   │   └── components/          # Reusable UI components
│   │       ├── styles.go        # Lip Gloss style definitions
│   │       ├── status.go        # TuiStatus badge component
│   │       ├── navitem.go       # TuiNavItem menu component
│   │       └── statusbar.go     # TuiStatusBar footer component
│   │
│   ├── service/                 # APPLICATION LAYER
│   │   ├── build.go             # Build orchestration service
│   │   ├── scan.go              # Component scanning service
│   │   └── config.go            # Configuration management service
│   │
│   ├── domain/                  # DOMAIN LAYER
│   │   ├── component.go         # Component entity
│   │   ├── buildconfig.go       # Build configuration value object
│   │   └── logentry.go          # Log entry entity + BuildResult
│   │
│   └── infra/                   # INFRASTRUCTURE LAYER
│       ├── ports/               # Interface definitions
│       │   ├── maven.go         # MavenExecutor interface
│       │   ├── git.go           # GitClient interface
│       │   └── filesystem.go    # FileSystem interface
│       │
│       └── adapters/            # Interface implementations
│           ├── adapters.go      # Type re-exports
│           ├── maven_real.go    # Real Maven CLI execution
│           ├── maven_mock.go    # Mock for testing
│           ├── git_real.go      # Real Git CLI execution
│           └── git_mock.go      # Mock for testing
│
├── configs/
│   └── components.yaml          # Component registry file
│
├── testdata/
│   └── golden/                  # Golden files for visual tests
│
├── go.mod                       # Go module definition
└── Makefile                     # Build commands
```

## Layer Responsibilities

### Presentation Layer (`internal/ui/`)

**Purpose:** Handle user interface state, rendering, and user input.

**Components:**
- **screens/**: Full-screen views (welcome, build, logs, discover, config)
- **modals/**: Overlay dialogs (build config, confirmation)
- **components/**: Reusable UI elements (status badges, nav items, status bar)

**Key Files:**
- `components/styles.go`: All Lip Gloss style definitions from design-system.md
- `components/status.go`: TuiStatus component for build status badges
- `components/navitem.go`: TuiNavItem component for navigation menu
- `components/statusbar.go`: TuiStatusBar for footer with keyboard hints

**Rules:**
- All screens implement `tea.Model` interface
- Use Bubbles components (list, table, viewport, textinput) where possible
- All styling via Lip Gloss (never raw ANSI codes)
- View functions are pure (no side effects)

### Application Layer (`internal/service/`)

**Purpose:** Orchestrate business logic and coordinate between UI and infrastructure.

**Services:**
- `build.go`: Build orchestration, component ordering, progress tracking
- `scan.go`: Component discovery, registry management
- `config.go`: Application configuration loading/saving

**Key Patterns:**
- Services receive port interfaces via constructor injection
- All async operations return `tea.Cmd` (integrates with Bubble Tea)
- Services never access external systems directly (only through ports)

### Domain Layer (`internal/domain/`)

**Purpose:** Define core business entities and value objects.

**Entities:**
- `Component`: RFZ component with ID, name, type, path, dependencies
- `LogEntry`: Build log line with level, timestamp, message
- `BuildResult`: Build outcome with success, duration, logs

**Value Objects:**
- `BuildConfig`: Immutable build configuration (profiles, skip tests, etc.)

**Rules:**
- No dependencies on other layers
- Pure Go types (no framework dependencies)
- Immutable where possible

### Infrastructure Layer (`internal/infra/`)

**Purpose:** Connect to external systems (Maven, Git, filesystem).

**Ports (Interfaces):**
- `MavenExecutor`: Execute builds, stream output, cancel builds
- `GitClient`: Get status, check repository, fetch/pull
- `FileSystem`: Read/write files, scan directories

**Adapters (Implementations):**
- `*_real.go`: Production implementations (actual CLI execution)
- `*_mock.go`: Test implementations (predictable fake data)

**Key Pattern:**
- Ports define contracts in `ports/` directory
- Adapters implement contracts in `adapters/` directory
- Mock adapters enable visual regression testing without real Maven/Git

## File Conventions

### Naming

| Element | Convention | Example |
|---------|------------|---------|
| Packages | lowercase, single word | `build`, `config`, `ports` |
| Files | lowercase with underscores | `maven_real.go`, `maven_mock.go` |
| Types | PascalCase | `BuildService`, `MavenExecutor` |
| Message types | `*Msg` suffix | `BuildCompleteMsg`, `ConfigLoadedMsg` |

### Screen Structure

Each screen package contains:
- Single `*.go` file with Model, Init, Update, View
- Messages defined at bottom of file
- Helper functions for rendering sections

For complex screens, split into:
- `model.go`: Model struct and constructor
- `update.go`: Update function and message handling
- `view.go`: View function and render helpers
- `messages.go`: Screen-specific message types

### Style Organization

All styles in `components/styles.go`:
- Color tokens (ColorBackground, ColorCyan, etc.)
- Typography styles (StyleH1, StyleBody, etc.)
- Component styles (StyleBoxDefault, StyleNavItem, etc.)
- Status styles (StyleLogInfo, StyleLogError, etc.)

## Dependency Flow

```
main.go
    │
    ├── Creates adapters (real or mock based on env)
    │   └── adapters.NewRealMavenExecutor()
    │   └── adapters.NewMockMavenExecutor()
    │
    ├── Creates services with injected adapters
    │   └── service.NewBuildService(maven, fileSystem)
    │   └── service.NewScanService(fileSystem, git)
    │
    └── Creates app model with services
        └── app.New(buildSvc, scanSvc, configSvc)
            │
            └── Creates screen models with services
                └── build.New(buildSvc)
                └── discover.New(scanSvc)
                └── config.New(configSvc)
```

## Testing Strategy

### Unit Tests
- Test service logic with mock adapters
- Test model state transitions
- Test message handling

### Visual Regression Tests
- Use `teatest` for golden file comparisons
- Standard terminal size: 120x40
- Mock adapters for predictable rendering

### Golden File Location
- `testdata/golden/`: Screen snapshots
- Naming: `{screen}-{state}.golden`

## Extension Points

### Adding a New Screen

1. Create package in `internal/ui/screens/{name}/`
2. Create `{name}.go` with Model struct implementing `tea.Model`
3. Add screen constant in `internal/app/app.go`
4. Add screen model as field in `app.Model`
5. Add case in `app.Update()` for message routing
6. Add case in `app.View()` for rendering

### Adding a New Port

1. Define interface in `internal/infra/ports/{name}.go`
2. Create real adapter in `internal/infra/adapters/{name}_real.go`
3. Create mock adapter in `internal/infra/adapters/{name}_mock.go`
4. Re-export in `internal/infra/adapters/adapters.go`

### Adding a New Service

1. Create service in `internal/service/{name}.go`
2. Define constructor with port dependencies
3. Define service methods returning `tea.Cmd`
4. Define related message types
5. Inject into app via `main.go`

## Key Files Quick Reference

| Purpose | File |
|---------|------|
| Entry point | `cmd/rfz-cli/main.go` |
| Main app model | `internal/app/app.go` |
| Shared messages | `internal/app/messages.go` |
| All styles | `internal/ui/components/styles.go` |
| Status badges | `internal/ui/components/status.go` |
| Build service | `internal/service/build.go` |
| Component entity | `internal/domain/component.go` |
| Maven interface | `internal/infra/ports/maven.go` |
| Maven mock | `internal/infra/adapters/maven_mock.go` |

---

This structure provides:
- Clear separation of concerns across layers
- Testability via swappable adapters
- Consistency through shared styles and patterns
- Extensibility through well-defined extension points
