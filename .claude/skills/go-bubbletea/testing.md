# Testing with teatest

Visual regression testing patterns for RFZ Developer CLI.

## Overview

- **Primary Tool**: teatest (charmbracelet/teatest)
- **Canonical Size**: 120 columns x 40 rows
- **Golden Files**: `testdata/golden/`
- **Coverage**: 97 UI states (from prototype screenshots)

## Setup

### Test File Structure

```go
// internal/ui/screens/build/build_test.go
package build_test

import (
    "testing"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/teatest"

    "rfz-tui/internal/ui/screens/build"
)
```

### Test Helper

```go
// internal/testutil/testutil.go
package testutil

import tea "github.com/charmbracelet/bubbletea"

const (
    CanonicalWidth  = 120
    CanonicalHeight = 40
)

func SetCanonicalSize(m tea.Model) (tea.Model, tea.Cmd) {
    return m.Update(tea.WindowSizeMsg{
        Width:  CanonicalWidth,
        Height: CanonicalHeight,
    })
}
```

## Golden File Tests

### Basic Golden Test

```go
func TestBuildScreen_Default(t *testing.T) {
    // Create model
    m := build.New()

    // Set canonical size
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})

    // Get view output
    got := m.View()

    // Compare with golden file
    // Creates testdata/golden/build-default.golden if it doesn't exist
    teatest.AssertGoldenString(t, got, "build-default.golden")
}
```

### With Mock Data

```go
func TestBuildScreen_WithComponents(t *testing.T) {
    // Create with mock service
    mockScanner := adapters.NewMockFileScanner()
    mockScanner.Components = []domain.Component{
        {ID: "boss", Name: "Boss", Type: domain.ComponentTypeCore},
        {ID: "fistiv", Name: "Fistiv", Type: domain.ComponentTypeSimulation},
    }

    svc := service.NewScanService(mockScanner)
    m := build.New(svc)

    // Load components
    m, _ = m.Update(build.ComponentsLoadedMsg{Components: mockScanner.Components})
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})

    got := m.View()
    teatest.AssertGoldenString(t, got, "build-with-components.golden")
}
```

### Testing States

```go
func TestBuildScreen_States(t *testing.T) {
    tests := []struct {
        name       string
        setup      func(m *build.Model)
        goldenFile string
    }{
        {
            name:       "default empty",
            setup:      func(m *build.Model) {},
            goldenFile: "build-empty.golden",
        },
        {
            name: "with selection",
            setup: func(m *build.Model) {
                m.Select(0)
                m.Select(2)
            },
            goldenFile: "build-selected.golden",
        },
        {
            name: "building in progress",
            setup: func(m *build.Model) {
                m.SetBuilding(true)
            },
            goldenFile: "build-running.golden",
        },
        {
            name: "build complete success",
            setup: func(m *build.Model) {
                m.SetBuildResult(domain.BuildResult{
                    Success: true,
                    Duration: 120 * time.Second,
                })
            },
            goldenFile: "build-success.golden",
        },
        {
            name: "build failed",
            setup: func(m *build.Model) {
                m.SetBuildResult(domain.BuildResult{
                    Success: false,
                    Error: errors.New("compilation error"),
                })
            },
            goldenFile: "build-failed.golden",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            m := build.New()
            tt.setup(&m)
            m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})

            got := m.View()
            teatest.AssertGoldenString(t, got, tt.goldenFile)
        })
    }
}
```

## Testing Update Logic

### Key Press Tests

```go
func TestBuildScreen_KeyNavigation(t *testing.T) {
    m := build.New()
    loadTestComponents(&m)
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})

    // Press 'j' to move down
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})

    if m.CursorIndex() != 1 {
        t.Errorf("expected cursor at 1, got %d", m.CursorIndex())
    }
}
```

### Selection Tests

```go
func TestBuildScreen_Selection(t *testing.T) {
    m := build.New()
    loadTestComponents(&m)

    // Press space to toggle selection
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})

    if !m.IsSelected(0) {
        t.Error("expected item 0 to be selected")
    }

    // Press space again to deselect
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})

    if m.IsSelected(0) {
        t.Error("expected item 0 to be deselected")
    }
}
```

### Select All/None

```go
func TestBuildScreen_SelectAll(t *testing.T) {
    m := build.New()
    loadTestComponents(&m) // 5 components

    // Press 'a' for select all
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})

    if m.SelectedCount() != 5 {
        t.Errorf("expected 5 selected, got %d", m.SelectedCount())
    }

    // Press 'n' for select none
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'n'}})

    if m.SelectedCount() != 0 {
        t.Errorf("expected 0 selected, got %d", m.SelectedCount())
    }
}
```

## Testing Commands

### Async Command Results

```go
func TestBuildScreen_BuildCommand(t *testing.T) {
    mockMaven := adapters.NewMockMavenExecutor()
    mockMaven.Results["boss"] = &domain.BuildResult{Success: true}

    svc := service.NewBuildService(mockMaven, nil)
    m := build.New(svc)
    loadTestComponents(&m)
    m.Select(0) // Select boss

    // Trigger build
    _, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})

    // Execute command
    msg := cmd()

    // Apply result message
    m, _ = m.Update(msg)

    if !m.BuildComplete() {
        t.Error("expected build to be complete")
    }
}
```

## Testing with Mocks

### Mock Adapters

```go
// internal/infra/adapters/maven_mock.go
type MockMavenExecutor struct {
    Results  map[string]*domain.BuildResult
    Errors   map[string]error
    Executed []string
}

func NewMockMavenExecutor() *MockMavenExecutor {
    return &MockMavenExecutor{
        Results: make(map[string]*domain.BuildResult),
        Errors:  make(map[string]error),
    }
}

func (m *MockMavenExecutor) Execute(cfg domain.BuildConfig) (*domain.BuildResult, error) {
    m.Executed = append(m.Executed, cfg.ComponentID)

    if err, ok := m.Errors[cfg.ComponentID]; ok {
        return nil, err
    }

    if result, ok := m.Results[cfg.ComponentID]; ok {
        return result, nil
    }

    return &domain.BuildResult{Success: true}, nil
}
```

### Using Test Data

```go
// testdata/components.go
package testdata

import "rfz-tui/internal/domain"

func DemoComponents() []domain.Component {
    return []domain.Component{
        {ID: "boss", Name: "Boss", Type: domain.ComponentTypeCore},
        {ID: "fistiv", Name: "Fistiv", Type: domain.ComponentTypeSimulation},
        {ID: "audiocon", Name: "Audiocon", Type: domain.ComponentTypeStandalone},
        {ID: "traktion", Name: "Traktion", Type: domain.ComponentTypeCore},
        {ID: "signalsteuerung", Name: "Signalsteuerung", Type: domain.ComponentTypeCore},
    }
}
```

## Updating Golden Files

### When Changes Are Intentional

```bash
# Update all golden files
go test ./... -update

# Update specific test
go test ./internal/ui/screens/build -update -run TestBuildScreen_Default
```

### Golden File Review Process

1. Make code changes
2. Run tests (expect failures if UI changed)
3. Review diffs in golden files
4. If changes are correct: `go test -update`
5. Commit golden files with code changes

## Test Organization

### File Naming

```
testdata/golden/
├── build-default.golden
├── build-selected.golden
├── build-running.golden
├── build-success.golden
├── build-failed.golden
├── logs-default.golden
├── logs-with-content.golden
├── config-scanpaths.golden
└── ...
```

### Coverage Goals

| Screen | States to Test |
|--------|----------------|
| Welcome | Default, with version |
| Build | Empty, with components, selected, building, success, failed |
| Logs | Empty, with logs, filtered, searching |
| Discover | Scanning, with results, empty |
| Config | Each tab, editing, saving |
| Modals | Build config, confirm |

## CI Integration

### GitHub Actions Workflow

```yaml
# .github/workflows/test.yml
name: Test

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run tests
        run: go test ./...

      - name: Check golden files
        run: |
          go test ./... -update
          git diff --exit-code testdata/
```
