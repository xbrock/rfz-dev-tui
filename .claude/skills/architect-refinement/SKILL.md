---
description: Architect Refinement - Technical refinement guidance for stories
globs:
  - "agent-os/specs/**/*.md"
  - "internal/**/*.go"
alwaysApply: false
version: 1.0.0
---

# Architect Refinement Skill

Technical refinement guidance for the RFZ Developer CLI project architecture.

## Quick Reference

### Architecture Layers

```
Presentation (internal/ui/)     - Bubble Tea models, views, components
      |
Application (internal/service/) - Build, Scan, Config services
      |
Domain (internal/domain/)       - Component, BuildConfig, LogEntry
      |
Infrastructure (internal/infra/) - Ports (interfaces) + Adapters (implementations)
```

### Key Architecture Decisions

| ID | Decision |
|----|----------|
| DEC-001 | Layered Architecture with Ports & Adapters |
| DEC-002 | Hierarchical Model Composition |
| DEC-003 | Service Layer Pattern |
| DEC-004 | Ports & Adapters for External Dependencies |
| DEC-005 | Screen Models Own Their State |
| DEC-006 | Message Types Per Screen |

## Directory Structure

```
rfz-tui/
├── cmd/rfz-cli/main.go          # Entry point, dependency wiring
├── internal/
│   ├── app/                      # Main Bubble Tea model
│   │   ├── app.go                # Root model, message routing
│   │   ├── messages.go           # Shared message types
│   │   └── keymap.go             # Global key bindings
│   ├── ui/
│   │   ├── screens/              # Full-screen views
│   │   │   ├── welcome/
│   │   │   ├── build/
│   │   │   ├── logs/
│   │   │   ├── discover/
│   │   │   └── config/
│   │   ├── modals/               # Overlay dialogs
│   │   │   ├── buildconfig/
│   │   │   └── confirm/
│   │   └── components/           # Reusable UI elements
│   │       ├── styles.go         # Lip Gloss definitions
│   │       ├── status.go         # TuiStatus
│   │       └── navitem.go        # TuiNavItem
│   ├── service/                  # Application layer
│   │   ├── build.go
│   │   ├── scan.go
│   │   └── config.go
│   ├── domain/                   # Domain layer
│   │   ├── component.go
│   │   ├── buildconfig.go
│   │   └── logentry.go
│   └── infra/                    # Infrastructure layer
│       ├── ports/                # Interfaces
│       └── adapters/             # Implementations (real + mock)
├── configs/                      # Configuration files
├── testdata/golden/              # Golden files for visual tests
└── references/                   # Design reference materials
```

## Technical Refinement Checklist

### For Screen Stories

- [ ] Which layer owns the state? (screen model)
- [ ] What messages does it handle?
- [ ] What tea.Cmd does it return for async?
- [ ] What services does it need?
- [ ] Which Bubbles components can be used?

### For Service Stories

- [ ] What ports (interfaces) does it need?
- [ ] What domain types does it use?
- [ ] What errors can occur?
- [ ] How does it integrate with Bubble Tea? (tea.Cmd)

### For External Integration Stories

- [ ] Define port interface in `internal/infra/ports/`
- [ ] Create real adapter in `internal/infra/adapters/*_real.go`
- [ ] Create mock adapter in `internal/infra/adapters/*_mock.go`
- [ ] Mock provides deterministic data for testing

## Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Package | lowercase, singular | `build`, `config` |
| File | lowercase with underscores | `maven_real.go` |
| Type | PascalCase | `BuildService` |
| Message | PascalCase + `Msg` | `BuildCompleteMsg` |
| Interface | PascalCase (no I prefix) | `MavenExecutor` |
| Method | PascalCase | `Execute`, `Update` |

## Pattern: Bubble Tea Model

```go
// Screen model structure
type Model struct {
    // UI state (Bubbles components)
    list     list.Model
    viewport viewport.Model

    // Local state
    focused  FocusArea
    selected map[string]bool

    // Dimensions
    width  int
    height int

    // Services (injected)
    buildSvc *service.BuildService
}

// Implement tea.Model interface
func (m Model) Init() tea.Cmd
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd)
func (m Model) View() string
```

## Pattern: Service Layer

```go
// Service with injected ports
type BuildService struct {
    maven   ports.MavenExecutor
    scanner ports.FileScanner
}

// Constructor with dependency injection
func NewBuildService(maven ports.MavenExecutor, scanner ports.FileScanner) *BuildService

// Method returns tea.Cmd for async operation
func (s *BuildService) Build(cfg domain.BuildConfig) tea.Cmd {
    return func() tea.Msg {
        result, err := s.maven.Execute(cfg)
        if err != nil {
            return BuildErrorMsg{Err: err}
        }
        return BuildCompleteMsg{Result: result}
    }
}
```

## Pattern: Port Interface

```go
// internal/infra/ports/maven.go
type MavenExecutor interface {
    Execute(cfg domain.BuildConfig) (*domain.BuildResult, error)
    Cancel(buildID string) error
    StreamOutput(buildID string) <-chan string
}
```

## Testing Strategy

### Visual Regression (Primary)

- **Tool**: teatest
- **Size**: 120x40 canonical terminal
- **Location**: `testdata/golden/`
- **Naming**: `{screen}-{state}.golden`

### Unit Tests

- Test state transitions (model.Update)
- Test view output (model.View)
- Use mock adapters for deterministic results

### Integration Tests

- Test service coordination
- Test message flow through app
- Use mock adapters (no real Maven/Git)

## Technical Constraints

### Must Follow

- **Charm.land First**: Use Bubbles/Lip Gloss before custom code
- **Elm Architecture**: MVU pattern is mandatory (Bubble Tea)
- **Dependency Injection**: Services receive ports, not implementations
- **Immutable Updates**: Model.Update returns new model, never mutates

### Must Avoid

- Direct external calls from UI layer
- Mutable state in models
- Raw ANSI codes (use Lip Gloss)
- Manual border drawing (use Lip Gloss borders)
