# Code Style Guide: Go + Bubble Tea

> RFZ Developer CLI Code Standards
> Last Updated: 2026-02-02
> Stack: Go 1.21+ | Bubble Tea | Lip Gloss | Bubbles

---

## CRITICAL RULE: Charm.land First - Custom Last

**This rule takes precedence over ALL other guidelines.**

Before implementing ANY visual/UI element:

1. **Check Bubbles first** - Use existing components (list, table, viewport, progress, spinner, help, textinput, paginator)
2. **Check Lip Gloss second** - Use for ALL styling (borders, colors, padding, layout)
3. **Check charmbracelet/log third** - Use for all logging output
4. **Custom implementation LAST** - Only when charm.land has NO solution

### Forbidden Patterns

```go
// NEVER DO THIS
fmt.Sprintf("%-20s", text)           // Manual padding
strings.Repeat("-", 40)               // Manual borders
"\033[31m" + text + "\033[0m"        // ANSI escape codes
fmt.Sprintf("| %s |", content)        // Manual box drawing

// ALWAYS DO THIS
lipgloss.NewStyle().Padding(0, 2).Render(text)
lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(content)
lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render(text)
```

---

## Go Formatting

### Tools (Mandatory)

```bash
# Run before every commit
gofmt -w .
goimports -w .
golangci-lint run
```

### Indentation

- Use tabs for indentation (Go standard)
- Use spaces for alignment within lines
- Let `gofmt` handle all formatting decisions

### Line Length

- Soft limit: 100 characters
- Hard limit: 120 characters
- Break long function calls at parameters

---

## Naming Conventions

### Go Standard Naming

| Element | Convention | Example |
|---------|------------|---------|
| Packages | lowercase, single word | `tui`, `build`, `config` |
| Exported types | PascalCase | `BuildModel`, `LogViewer` |
| Unexported types | camelCase | `buildState`, `logEntry` |
| Exported functions | PascalCase | `NewBuildModel`, `Update` |
| Unexported functions | camelCase | `parseOutput`, `formatTime` |
| Constants | PascalCase or camelCase | `MaxRetries`, `defaultPort` |
| Interfaces | PascalCase, -er suffix when appropriate | `Builder`, `Renderer` |

### Bubble Tea Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Model types | `[Component]Model` | `BuildModel`, `LogViewerModel` |
| Message types | `[Action]Msg` | `BuildStartMsg`, `LogReceivedMsg` |
| Command functions | `[action]Cmd` | `startBuildCmd`, `fetchLogsCmd` |
| View helper functions | `render[Element]` | `renderHeader`, `renderStatus` |
| Style variables | `[element]Style` | `headerStyle`, `activeItemStyle` |
| Key bindings struct | `keyMap` | `type keyMap struct` |

### Message Type Naming

```go
// Internal messages (within component)
type tickMsg time.Time
type errMsg error

// External messages (between components)
type BuildStartedMsg struct {
    ComponentID string
    StartTime   time.Time
}

type BuildCompletedMsg struct {
    ComponentID string
    Success     bool
    Duration    time.Duration
}
```

---

## File Organization

### Directory Structure

```
internal/
  tui/
    app.go              # Main application model
    keys.go             # Global key bindings
    styles.go           # Global Lip Gloss styles
    messages.go         # Shared message types

    components/         # Reusable TUI components
      tuibox/
        tuibox.go       # TuiBox component
        styles.go       # TuiBox-specific styles
      tuistatus/
        tuistatus.go
        styles.go
      ...

    screens/            # Full screen views
      build/
        model.go        # BuildModel struct and methods
        update.go       # Update function
        view.go         # View function
        keys.go         # Screen-specific key bindings
        messages.go     # Screen-specific messages
        styles.go       # Screen-specific styles
      logs/
        ...
      discover/
        ...
      config/
        ...

  domain/               # Business logic (non-TUI)
    maven/
      executor.go
      parser.go
    git/
      status.go
    component/
      scanner.go
      registry.go

  config/               # Application configuration
    config.go
    paths.go
```

### File Naming Rules

- One Model per screen/component directory
- Separate `model.go`, `update.go`, `view.go` for complex screens
- Keep simple components in single file
- Styles always in `styles.go`
- Messages always in `messages.go`

### File Size Guidelines

- Model files: < 200 lines
- Update functions: < 150 lines per file
- View functions: < 100 lines per file
- Break large files into focused sub-files

---

## Bubble Tea Model Structure

### Standard Model Template

```go
package build

import (
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

// Model represents the build screen state
type Model struct {
    // Bubbles components (use existing components)
    list     list.Model
    viewport viewport.Model

    // State
    selected    []string
    building    bool
    err         error

    // Dimensions
    width  int
    height int

    // Key bindings
    keys keyMap
}

// New creates a new Model with default state
func New() Model {
    return Model{
        list:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
        keys:     defaultKeyMap(),
        selected: make([]string, 0),
    }
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
    return nil
}

// Update handles messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle messages...
}

// View renders the model
func (m Model) View() string {
    // Render using Lip Gloss...
}
```

### Model Rules

1. **Immutability**: Never mutate state directly, return new Model
2. **No side effects in View**: View only renders, never modifies state
3. **Commands for async**: All I/O operations return tea.Cmd
4. **Type all messages**: Use typed message structs, not raw values

---

## Lip Gloss Style Definitions

### Style Organization

```go
// styles.go - Define all styles in dedicated file

package build

import "github.com/charmbracelet/lipgloss"

var (
    // Layout styles
    containerStyle = lipgloss.NewStyle().
        Padding(1, 2).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("62"))

    // Text styles
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("170"))

    // State-based styles
    successStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("42"))

    errorStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("196"))

    // Adaptive colors for light/dark terminals
    subtleColor = lipgloss.AdaptiveColor{Light: "250", Dark: "238"}
)
```

### Style Rules

1. **Define once, reuse everywhere**: Styles in `styles.go`
2. **Use adaptive colors**: Support both light and dark terminals
3. **Name by purpose**: `errorStyle` not `redStyle`
4. **Compose styles**: Use `.Inherit()` and `.Copy()` for variants

---

## Error Handling

### Error Patterns

```go
// Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to start build for %s: %w", componentID, err)
}

// Error messages (user-facing via TUI)
type errMsg struct {
    err error
}

func (e errMsg) Error() string {
    return e.err.Error()
}

// Command that can fail
func startBuildCmd(componentID string) tea.Cmd {
    return func() tea.Msg {
        err := maven.Build(componentID)
        if err != nil {
            return errMsg{err: err}
        }
        return buildCompletedMsg{componentID: componentID}
    }
}
```

### Error Display

```go
// In View function
if m.err != nil {
    errorBox := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("196")).
        Padding(1, 2).
        Render(m.err.Error())

    return lipgloss.JoinVertical(lipgloss.Left, content, errorBox)
}
```

---

## Comments

### When to Comment

```go
// Comment WHY, not WHAT
// Using viewport instead of custom scrolling because charm.land recommends it
// and it handles edge cases like terminal resize automatically.
viewport := viewport.New(width, height)

// Document non-obvious business logic
// Maven profiles must be applied in this order: base profile first,
// then component-specific overrides, as per RFZ build conventions.
profiles := append([]string{"base"}, componentProfiles...)

// Document complex calculations
// Terminal height minus header (3 lines), footer (2 lines), and borders (2 lines)
contentHeight := m.height - 7
```

### Comment Style

```go
// Single-line comments for brief explanations
// Start with capital letter, no period unless multiple sentences.

/*
Multi-line comments for longer explanations.
Use sparingly - prefer multiple single-line comments.
*/

// TODO(username): Description of what needs to be done
// FIXME(username): Description of bug to fix
```

---

## Imports

### Import Organization

```go
import (
    // Standard library
    "fmt"
    "strings"
    "time"

    // Third-party (charm.land first)
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/viewport"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/log"

    // Other third-party
    "github.com/spf13/cobra"

    // Internal packages
    "rfz-cli/internal/domain/maven"
    "rfz-cli/internal/tui/components/tuistatus"
)
```

### Import Rules

- Three groups: stdlib, third-party, internal
- Blank line between groups
- Alphabetical within groups
- Use aliases only when necessary (`tea "github.com/charmbracelet/bubbletea"`)

---

## Testing

### Test File Organization

```go
// model_test.go - Unit tests for model
package build

import (
    "testing"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/teatest"
)

func TestModel_Update_BuildStart(t *testing.T) {
    m := New()
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40}).(Model)

    // Test state transition
    m, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter}).(Model)

    if !m.building {
        t.Error("expected building to be true")
    }
    if cmd == nil {
        t.Error("expected command to be returned")
    }
}
```

### Golden File Tests

```go
func TestModel_View_GoldenFile(t *testing.T) {
    m := New()
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40}).(Model)

    tm := teatest.NewTestModel(t, m)
    teatest.WaitFor(t, tm, func(b []byte) bool {
        return strings.Contains(string(b), "Build Components")
    })

    out := tm.FinalModel().(Model).View()
    golden.Assert(t, out, "build-components-default.golden")
}
```

---

## Key Bindings

### Key Binding Structure

```go
// keys.go
package build

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
    Up       key.Binding
    Down     key.Binding
    Select   key.Binding
    Build    key.Binding
    Help     key.Binding
    Quit     key.Binding
}

func defaultKeyMap() keyMap {
    return keyMap{
        Up: key.NewBinding(
            key.WithKeys("up", "k"),
            key.WithHelp("up/k", "move up"),
        ),
        Down: key.NewBinding(
            key.WithKeys("down", "j"),
            key.WithHelp("down/j", "move down"),
        ),
        Select: key.NewBinding(
            key.WithKeys(" "),
            key.WithHelp("space", "toggle selection"),
        ),
        Build: key.NewBinding(
            key.WithKeys("b", "enter"),
            key.WithHelp("b/enter", "start build"),
        ),
        Help: key.NewBinding(
            key.WithKeys("?"),
            key.WithHelp("?", "toggle help"),
        ),
        Quit: key.NewBinding(
            key.WithKeys("q", "ctrl+c"),
            key.WithHelp("q", "quit"),
        ),
    }
}

// ShortHelp returns key bindings for short help
func (k keyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Up, k.Down, k.Select, k.Build, k.Help}
}

// FullHelp returns key bindings for full help
func (k keyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down, k.Select},
        {k.Build, k.Help, k.Quit},
    }
}
```

---

## Quick Reference

### Do

- Use `gofmt` and `goimports`
- Use Bubbles components before custom
- Use Lip Gloss for ALL styling
- Define styles in `styles.go`
- Use typed messages
- Return tea.Cmd for async operations
- Keep View functions pure

### Do Not

- Use manual string padding
- Use ANSI escape codes
- Use custom border drawing
- Mutate model state directly
- Perform I/O in View functions
- Use `fmt.Println` in TUI code

---

**Remember:** When in doubt, check charm.land documentation first. The answer is almost always already implemented in Bubbles or achievable with Lip Gloss.
