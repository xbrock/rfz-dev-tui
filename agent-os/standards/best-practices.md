# Best Practices: Go + Bubble Tea

> RFZ Developer CLI Development Guidelines
> Last Updated: 2026-02-02
> Stack: Go 1.21+ | Bubble Tea | Lip Gloss | Bubbles

---

## CRITICAL RULE: Charm.land First - Custom Last

**This rule is MANDATORY. Violation will cause code review rejection.**

### Priority Order (MUST Follow)

| Priority | Source | When to Use |
|----------|--------|-------------|
| 1st | **Bubbles** | list, table, viewport, progress, spinner, help, textinput, paginator |
| 2nd | **Lip Gloss** | ALL styling: borders, colors, padding, margins, layout, text formatting |
| 3rd | **charmbracelet/log** | ALL logging output |
| 4th | **Custom** | ONLY when charm.land has NO solution |

### Forbidden Implementations

| Category | NEVER DO | ALWAYS DO |
|----------|----------|-----------|
| Borders | Box-drawing chars (`---`, `\|`) | `lipgloss.Border(lipgloss.RoundedBorder())` |
| Colors | ANSI codes (`\033[31m`) | `lipgloss.Color("#ff0000")` |
| Padding | `fmt.Sprintf("%-20s")` | `lipgloss.Padding(0, 2)` |
| Progress | Custom `[=====>    ]` | `bubbles/progress` |
| Spinners | Custom frame arrays | `bubbles/spinner` |
| Lists | Manual item rendering | `bubbles/list` |
| Tables | Manual column alignment | `bubbles/table` |
| Scrolling | Custom scroll logic | `bubbles/viewport` |

### When Custom IS Allowed

Custom implementation is permitted ONLY for:

1. **Business logic components** - TuiStatus (build status badges), TuiNavItem (shortcuts)
2. **Composite components** - TuiModal (viewport + Lip Gloss border + overlay)
3. **Domain-specific rendering** - Maven phases, Git status indicators

**Even then, custom components MUST use Lip Gloss for ALL internal styling.**

---

## Bubble Tea Architecture Patterns

### The Elm Architecture (Model-Update-View)

```
User Input / System Event
         |
         v
    +----------+
    |  Update  |  Pure function: (Model, Msg) -> (Model, Cmd)
    +----------+
         |
         v
    +----------+
    |   Model  |  Single source of truth
    +----------+
         |
         v
    +----------+
    |   View   |  Pure function: Model -> String
    +----------+
         |
         v
    Terminal Output
```

### Golden Rules

1. **Model is immutable** - Update returns NEW model, never mutates
2. **View is pure** - No side effects, only renders state
3. **Commands for I/O** - All async/IO operations return tea.Cmd
4. **Messages are typed** - Use structs, not raw values

---

## State Management

### Single Source of Truth

```go
// GOOD: All state in Model
type Model struct {
    components   []Component
    selected     map[string]bool
    building     bool
    currentBuild *BuildState
    logs         []LogEntry
    err          error
}

// BAD: State scattered across globals
var globalSelectedComponents []string  // NEVER DO THIS
var isBuilding bool                    // NEVER DO THIS
```

### State Updates Pattern

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case buildStartMsg:
        // Return NEW model with updated state
        m.building = true
        m.currentBuild = &BuildState{
            ComponentID: msg.componentID,
            StartTime:   time.Now(),
        }
        return m, nil

    case buildCompleteMsg:
        m.building = false
        m.currentBuild = nil
        return m, nil
    }
    return m, nil
}
```

### Derived State

```go
// Compute derived values in View or helper methods
// Do NOT store computed values in Model

func (m Model) selectedCount() int {
    count := 0
    for _, selected := range m.selected {
        if selected {
            count++
        }
    }
    return count
}

func (m Model) View() string {
    // Use derived state in view
    header := fmt.Sprintf("Build (%d selected)", m.selectedCount())
    // ...
}
```

---

## Message Passing

### Message Design

```go
// Internal messages: lowercase, simple
type tickMsg time.Time
type errMsg struct{ error }

// External/shared messages: exported, descriptive
type BuildStartedMsg struct {
    ComponentID string
    Config      BuildConfig
    StartTime   time.Time
}

type BuildProgressMsg struct {
    ComponentID string
    Phase       string  // "compile", "test", "package", "install"
    Progress    float64 // 0.0 to 1.0
}

type BuildCompletedMsg struct {
    ComponentID string
    Success     bool
    Duration    time.Duration
    Output      string
}
```

### Command Patterns

```go
// Command that performs I/O
func startBuildCmd(cfg BuildConfig) tea.Cmd {
    return func() tea.Msg {
        result, err := maven.Execute(cfg)
        if err != nil {
            return errMsg{err}
        }
        return BuildCompletedMsg{
            ComponentID: cfg.ComponentID,
            Success:     result.Success,
            Duration:    result.Duration,
        }
    }
}

// Tick command for animations/polling
func tickCmd() tea.Cmd {
    return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

// Batch multiple commands
func (m Model) startMultipleBuildCmd() tea.Cmd {
    var cmds []tea.Cmd
    for id := range m.selected {
        cmds = append(cmds, startBuildCmd(BuildConfig{ComponentID: id}))
    }
    return tea.Batch(cmds...)
}
```

### Message Routing

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Handle global messages first
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }

    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }

    // Delegate to child components
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    cmds = append(cmds, cmd)

    m.viewport, cmd = m.viewport.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}
```

---

## Component Composition

### Embedding Bubbles Components

```go
type Model struct {
    // Embed Bubbles components
    list     list.Model
    viewport viewport.Model
    spinner  spinner.Model
    help     help.Model

    // Custom state
    items    []Item
    loading  bool
}

func New() Model {
    // Initialize Bubbles components with defaults
    l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
    l.Title = "Components"
    l.SetShowStatusBar(true)
    l.SetFilteringEnabled(true)

    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return Model{
        list:    l,
        spinner: s,
        help:    help.New(),
    }
}
```

### Custom Component Pattern

```go
// TuiStatus - A custom component for build status badges
package tuistatus

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Status int

const (
    Pending Status = iota
    Running
    Success
    Failed
)

type Model struct {
    Status Status
    Label  string
}

// Must use Lip Gloss for ALL styling
var (
    pendingStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("241")).
        Background(lipgloss.Color("236")).
        Padding(0, 1)

    runningStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("220")).
        Background(lipgloss.Color("236")).
        Padding(0, 1).
        Bold(true)

    successStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("42")).
        Background(lipgloss.Color("236")).
        Padding(0, 1)

    failedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("196")).
        Background(lipgloss.Color("236")).
        Padding(0, 1)
)

func (m Model) View() string {
    icon := m.icon()
    style := m.style()
    return style.Render(icon + " " + m.Label)
}

func (m Model) icon() string {
    switch m.Status {
    case Pending:
        return "o"
    case Running:
        return "*"
    case Success:
        return "+"
    case Failed:
        return "x"
    default:
        return "?"
    }
}

func (m Model) style() lipgloss.Style {
    switch m.Status {
    case Pending:
        return pendingStyle
    case Running:
        return runningStyle
    case Success:
        return successStyle
    case Failed:
        return failedStyle
    default:
        return pendingStyle
    }
}
```

---

## Layout with Lip Gloss

### Responsive Layout

```go
func (m Model) View() string {
    // Calculate available space
    headerHeight := 3
    footerHeight := 2
    contentHeight := m.height - headerHeight - footerHeight

    // Render sections
    header := m.renderHeader()
    content := m.renderContent(contentHeight)
    footer := m.renderFooter()

    // Join vertically
    return lipgloss.JoinVertical(
        lipgloss.Left,
        header,
        content,
        footer,
    )
}

func (m Model) renderContent(height int) string {
    // Split horizontal: sidebar + main
    sidebarWidth := 30
    mainWidth := m.width - sidebarWidth - 1 // -1 for divider

    sidebar := m.renderSidebar(sidebarWidth, height)
    main := m.renderMain(mainWidth, height)

    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        sidebar,
        main,
    )
}
```

### Centering and Placement

```go
// Center content in available space
func (m Model) renderModal() string {
    modalWidth := 60
    modalHeight := 20

    modal := lipgloss.NewStyle().
        Width(modalWidth).
        Height(modalHeight).
        Border(lipgloss.DoubleBorder()).
        BorderForeground(lipgloss.Color("62")).
        Padding(1, 2).
        Render(m.modalContent())

    // Center in terminal
    return lipgloss.Place(
        m.width,
        m.height,
        lipgloss.Center,
        lipgloss.Center,
        modal,
    )
}
```

---

## Testing Patterns

### Model Unit Tests

```go
func TestModel_SelectComponent(t *testing.T) {
    m := New()
    m.items = []Item{{ID: "comp-1"}, {ID: "comp-2"}}

    // Simulate selection
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{' '}}).(Model)

    if !m.selected["comp-1"] {
        t.Error("expected comp-1 to be selected")
    }
}
```

### Golden File Tests (Visual Regression)

```go
func TestBuildScreen_Default(t *testing.T) {
    m := build.New()

    // Set canonical size
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40}).(Model)

    // Load test data
    m.SetComponents(testdata.DemoComponents())

    // Compare against golden file
    got := m.View()
    golden.Assert(t, got, "build-screen-default.golden")
}

func TestBuildScreen_WithSelection(t *testing.T) {
    m := build.New()
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40}).(Model)
    m.SetComponents(testdata.DemoComponents())

    // Select first two components
    m.Select("core-1")
    m.Select("simulator-1")

    got := m.View()
    golden.Assert(t, got, "build-screen-selected.golden")
}
```

### Command Testing

```go
func TestStartBuildCmd_Success(t *testing.T) {
    // Mock Maven executor
    oldExecutor := maven.Executor
    maven.Executor = &mockExecutor{success: true}
    defer func() { maven.Executor = oldExecutor }()

    cmd := startBuildCmd(BuildConfig{ComponentID: "test-comp"})
    msg := cmd()

    result, ok := msg.(BuildCompletedMsg)
    if !ok {
        t.Fatal("expected BuildCompletedMsg")
    }
    if !result.Success {
        t.Error("expected success")
    }
}
```

---

## Performance Patterns

### Lazy Rendering

```go
// Only render visible items
func (m Model) renderList(height int) string {
    startIdx := m.scrollOffset
    endIdx := min(startIdx+height, len(m.items))

    var rows []string
    for i := startIdx; i < endIdx; i++ {
        rows = append(rows, m.renderItem(m.items[i]))
    }

    return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
```

### Efficient Updates

```go
// Avoid unnecessary allocations in hot paths
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        // Only update spinner if visible
        if m.loading {
            var cmd tea.Cmd
            m.spinner, cmd = m.spinner.Update(msg)
            return m, cmd
        }
        return m, nil
    }
    return m, nil
}
```

### Style Caching

```go
// Define styles as package-level vars (computed once)
var (
    headerStyle = lipgloss.NewStyle().Bold(true)
    itemStyle   = lipgloss.NewStyle().Padding(0, 1)
)

// NOT inside functions (computed every call)
func (m Model) renderItem(item Item) string {
    // BAD: Creates new style every call
    // style := lipgloss.NewStyle().Padding(0, 1)

    // GOOD: Use cached style
    return itemStyle.Render(item.Name)
}
```

---

## Error Handling

### User-Friendly Errors

```go
// Wrap errors with user context
func scanComponents(path string) ([]Component, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return nil, fmt.Errorf("cannot scan directory %q: %w", path, err)
    }
    // ...
}

// Display errors in UI
func (m Model) View() string {
    if m.err != nil {
        errorBox := lipgloss.NewStyle().
            Border(lipgloss.RoundedBorder()).
            BorderForeground(lipgloss.Color("196")).
            Foreground(lipgloss.Color("196")).
            Padding(1, 2).
            Width(m.width - 4).
            Render("Error: " + m.err.Error())

        return lipgloss.JoinVertical(lipgloss.Left, m.content, errorBox)
    }
    return m.content
}
```

### Recovery Pattern

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case errMsg:
        m.err = msg.error
        m.loading = false
        // Allow user to retry
        return m, nil

    case tea.KeyMsg:
        if msg.String() == "r" && m.err != nil {
            // Retry last operation
            m.err = nil
            return m, m.lastCmd
        }
    }
    return m, nil
}
```

---

## Logging (charmbracelet/log)

### Application Logging

```go
import "github.com/charmbracelet/log"

// Configure logger
func init() {
    log.SetLevel(log.DebugLevel)
    log.SetReportTimestamp(true)
}

// Structured logging
func startBuild(cfg BuildConfig) {
    log.Info("starting build",
        "component", cfg.ComponentID,
        "profiles", cfg.Profiles,
        "skipTests", cfg.SkipTests,
    )
}

func buildComplete(componentID string, success bool, duration time.Duration) {
    if success {
        log.Info("build completed",
            "component", componentID,
            "duration", duration,
        )
    } else {
        log.Error("build failed",
            "component", componentID,
            "duration", duration,
        )
    }
}
```

---

## Quick Reference Checklist

### Before Writing Code

- [ ] Is there a Bubbles component for this? (list, table, viewport, etc.)
- [ ] Can Lip Gloss handle the styling? (borders, colors, padding)
- [ ] Does charmbracelet/log cover logging needs?

### Model Checklist

- [ ] All state in Model struct
- [ ] No global variables
- [ ] Typed message structs
- [ ] Commands for async operations

### View Checklist

- [ ] Pure function (no side effects)
- [ ] Uses Lip Gloss for ALL styling
- [ ] Handles terminal resize
- [ ] Cached styles (not created in function)

### Update Checklist

- [ ] Returns new Model (no mutation)
- [ ] Delegates to child components
- [ ] Batches related commands
- [ ] Handles error messages

### Test Checklist

- [ ] Golden file for each UI state
- [ ] Tests at canonical size (120x40)
- [ ] Unit tests for state transitions
- [ ] Mocked external dependencies

---

**Remember:** The charm.land ecosystem solves 95% of TUI problems. Custom code should be the exception, not the rule. When you find yourself writing custom rendering logic, stop and check if charm.land already provides a solution.
