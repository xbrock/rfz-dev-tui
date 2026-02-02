---
description: Go Bubble Tea TUI Development - Patterns and best practices
globs:
  - "**/*.go"
  - "internal/ui/**/*"
  - "internal/app/**/*"
alwaysApply: false
version: 1.0.0
---

# Go Bubble Tea TUI Skill

Development patterns for the RFZ Developer CLI using Go and the charm.land stack.

## Quick Reference

### Tech Stack

| Library | Version | Purpose |
|---------|---------|---------|
| Go | 1.21+ | Language |
| bubbletea | v0.25+ | TUI framework (Elm architecture) |
| bubbles | v0.18+ | Pre-built components |
| lipgloss | v0.9+ | Terminal styling |
| log | v0.3+ | Structured logging |
| teatest | v0.0.1+ | Visual testing |

### Charm.land First Rule

**MANDATORY**: Before implementing ANY UI element:

1. Check if Bubbles has a component
2. Check if Lip Gloss can style it
3. Only then consider custom implementation

### Forbidden Patterns

| DO NOT | USE INSTEAD |
|--------|-------------|
| `---` or `═══` borders | `lipgloss.Border()` |
| `\033[31m` ANSI codes | `lipgloss.Color("#ef4444")` |
| Manual `strings.Repeat(" ", n)` | `lipgloss.Padding()` |
| Custom progress bar strings | `bubbles/progress` |
| Custom spinner frames | `bubbles/spinner` |
| Custom list rendering | `bubbles/list` |

## Sub-Documents

| Document | Content |
|----------|---------|
| [components.md](components.md) | Bubbles component usage |
| [styling.md](styling.md) | Lip Gloss patterns |
| [testing.md](testing.md) | teatest visual regression |
| [dos-and-donts.md](dos-and-donts.md) | Project learnings |

## Bubble Tea Model Pattern

### Basic Model Structure

```go
package myscreen

import (
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Model struct {
    // Bubbles components
    list list.Model

    // Local state
    focused bool

    // Dimensions (from WindowSizeMsg)
    width  int
    height int
}

func New() Model {
    return Model{
        list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
    }
}

func (m Model) Init() tea.Cmd {
    return nil // or initial command
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.list.SetSize(msg.Width-4, msg.Height-6)
        return m, nil

    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        }
    }

    // Delegate to child components
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    return lipgloss.JoinVertical(
        lipgloss.Left,
        m.headerView(),
        m.list.View(),
        m.footerView(),
    )
}
```

### Message Types

```go
// Screen-specific messages
type itemSelectedMsg struct {
    index int
}

type dataLoadedMsg struct {
    items []domain.Component
}

type errorMsg struct {
    err error
}

// Commands that return messages
func loadDataCmd(svc *service.ScanService) tea.Cmd {
    return func() tea.Msg {
        items, err := svc.Scan()
        if err != nil {
            return errorMsg{err: err}
        }
        return dataLoadedMsg{items: items}
    }
}
```

## Lip Gloss Styling

### Color Tokens (from design-system.md)

```go
var (
    ColorBackground    = lipgloss.Color("#1e1e2e")
    ColorCard          = lipgloss.Color("#2a2a3e")
    ColorBorder        = lipgloss.Color("#4a4a5e")
    ColorCyan          = lipgloss.Color("#0891b2")
    ColorGreen         = lipgloss.Color("#22c55e")
    ColorYellow        = lipgloss.Color("#eab308")
    ColorDestructive   = lipgloss.Color("#ef4444")
    ColorBrand         = lipgloss.Color("#ec0016")  // DB Red
    ColorTextPrimary   = lipgloss.Color("#f4f4f5")
    ColorTextSecondary = lipgloss.Color("#a1a1aa")
    ColorTextMuted     = lipgloss.Color("#71717a")
)
```

### Border Styles

```go
var (
    StyleBoxDefault = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(ColorBorder).
        Padding(1, 2)

    StyleBoxFocused = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(ColorCyan).  // Cyan for focus
        Padding(1, 2)

    StyleBoxModal = lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).  // Double for modals
        BorderForeground(ColorTextPrimary).
        Padding(1, 2)
)
```

### Layout Helpers

```go
// Horizontal join (side by side)
lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

// Vertical join (stacked)
lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

// Center content
lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)

// Set fixed dimensions
style.Width(40).Height(10)
```

## Bubbles Components

### List Component

```go
import "github.com/charmbracelet/bubbles/list"

// Create with custom delegate
delegate := list.NewDefaultDelegate()
delegate.Styles.SelectedTitle = lipgloss.NewStyle().
    Foreground(ColorCyan).
    Bold(true)

l := list.New(items, delegate, width, height)
l.Title = "Build Components"
l.Styles.Title = lipgloss.NewStyle().Bold(true)
```

### Viewport Component

```go
import "github.com/charmbracelet/bubbles/viewport"

vp := viewport.New(width, height)
vp.SetContent(logContent)

// Handle scrolling in Update
vp, cmd = vp.Update(msg)
```

### Progress Component

```go
import "github.com/charmbracelet/bubbles/progress"

p := progress.New(progress.WithDefaultGradient())
p.SetPercent(0.45)
view := p.View()
```

### Spinner Component

```go
import "github.com/charmbracelet/bubbles/spinner"

s := spinner.New()
s.Spinner = spinner.Dot
s.Style = lipgloss.NewStyle().Foreground(ColorCyan)
```

## Testing with teatest

### Golden File Test

```go
func TestBuildScreen_Default(t *testing.T) {
    m := build.New()
    m, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})

    got := m.View()
    golden.Assert(t, got, "build-default.golden")
}
```

### Update test to verify state changes

```go
func TestBuildScreen_Selection(t *testing.T) {
    m := build.New()

    // Simulate key press
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})

    // Assert state changed
    if !m.IsSelected(0) {
        t.Error("expected item 0 to be selected")
    }
}
```

## Common Patterns

### Focus Management

```go
type FocusArea int

const (
    FocusNavigation FocusArea = iota
    FocusContent
    FocusActions
)

func (m *Model) NextFocus() {
    m.focused = (m.focused + 1) % 3
}
```

### Keyboard Shortcuts

```go
case tea.KeyMsg:
    switch msg.String() {
    case "1", "2", "3", "4":
        // Screen navigation
        return m, navigateCmd(msg.String())
    case "q":
        return m, tea.Quit
    case "tab":
        m.NextFocus()
        return m, nil
    case " ":
        m.ToggleSelection()
        return m, nil
    }
```

### Async Operations

```go
// Command that performs async work
func buildComponentsCmd(svc *service.BuildService, components []string) tea.Cmd {
    return func() tea.Msg {
        result, err := svc.Build(components)
        if err != nil {
            return buildErrorMsg{err: err}
        }
        return buildCompleteMsg{result: result}
    }
}

// Handle in Update
case buildCompleteMsg:
    m.building = false
    m.result = msg.result
    return m, nil
```
