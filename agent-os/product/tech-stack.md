# Technical Stack: RFZ Developer CLI

> Last Updated: 2026-02-02
> Version: 1.0.0
> Type: Desktop TUI Application (Terminal User Interface)

---

## Overview

The RFZ Developer CLI is a Go-based Terminal User Interface application built entirely on the charm.land ecosystem. This document captures all technology decisions with rationale for each choice.

---

## Core Language

**Language:** Go 1.21+

### Rationale

| Factor | Go Advantage |
|--------|--------------|
| **Single Binary** | Compiles to a single executable with no runtime dependencies; easy distribution to developers |
| **Cross-Platform** | Native compilation for macOS, Linux, Windows without code changes |
| **Performance** | Fast startup time critical for CLI tools; no JVM warmup |
| **Concurrency** | Goroutines handle concurrent Maven builds and real-time output streaming efficiently |
| **Static Typing** | Compile-time safety important for complex TUI state management |
| **charm.land Ecosystem** | Best-in-class TUI libraries are written in Go (Bubble Tea, Lip Gloss, Bubbles) |

### Version Requirement

- **Minimum:** Go 1.21
- **Reason:** Generics support (1.18+), improved error handling, and compatibility with latest charm.land libraries

---

## TUI Framework

**Framework:** Bubble Tea (github.com/charmbracelet/bubbletea)

### Rationale

| Factor | Bubble Tea Advantage |
|--------|---------------------|
| **Elm Architecture** | Model-Update-View pattern provides predictable state management; essential for complex multi-screen TUI |
| **Testability** | teatest library enables golden file testing for visual regression; critical success metric (97 UI states) |
| **Composability** | Components nest naturally; build complex UIs from simple parts |
| **Message Passing** | Typed messages enable loose coupling between components |
| **Active Development** | Maintained by Charm, well-documented, strong community |
| **Production Proven** | Used by major projects (gh CLI, soft-serve, glow) |

### Architecture Principles

1. **Immutable State:** Updates return new state, never mutate
2. **Pure View Functions:** View only renders, no side effects
3. **Typed Messages:** All component communication via explicit message types
4. **Single Source of Truth:** Application state in top-level Model

---

## Styling Library

**Library:** Lip Gloss (github.com/charmbracelet/lipgloss)

### Rationale

| Factor | Lip Gloss Advantage |
|--------|---------------------|
| **Declarative Styling** | CSS-like API for terminal styling; readable and maintainable |
| **Layout System** | Built-in joining, placing, and alignment functions |
| **Border Support** | Multiple border styles (Normal, Rounded, Double, Thick, Hidden) |
| **Color Handling** | Adaptive colors for light/dark terminals; true color support |
| **Padding/Margin** | First-class support for spacing without manual string manipulation |
| **Integration** | Native integration with Bubble Tea and Bubbles |

### Styling Standards

All visual elements MUST use Lip Gloss. See "Critical Rule" section below.

---

## Component Library

**Library:** Bubbles (github.com/charmbracelet/bubbles)

### Components Used

| Bubbles Component | RFZ CLI Usage | Key Features |
|-------------------|---------------|--------------|
| `list` | Component selection, navigation lists | Filtering, pagination, help integration, spinner |
| `table` | Component registry, detected components | Column customization, scrolling, selection |
| `viewport` | Log viewer content area | Scrollable content, essential for large logs |
| `progress` | Build progress indicators | Animation via Harmonica, customizable colors |
| `spinner` | Loading states during builds | Multiple built-in styles (Dot, Line, MiniDot, etc.) |
| `textinput` | Configuration inputs, search | Unicode, pasting, cursor management |
| `help` | Keyboard shortcut display | Standardized key binding format |
| `paginator` | Log pagination (optional) | Dot and Arabic numeral styles |

### Rationale

| Factor | Bubbles Advantage |
|--------|------------------|
| **Consistency** | Pre-built components ensure consistent UX patterns |
| **Accessibility** | Tested for keyboard navigation, screen reader hints |
| **Maintenance** | Upstream updates improve components without code changes |
| **Integration** | Native Bubble Tea integration; implements tea.Model |

---

## Logging

**Library:** charmbracelet/log (github.com/charmbracelet/log)

### Rationale

| Factor | charmbracelet/log Advantage |
|--------|----------------------------|
| **Lip Gloss Integration** | Styled output consistent with application theme |
| **Structured Logging** | Key-value pairs for context (component name, build phase) |
| **Level Support** | Debug, Info, Warn, Error levels match Log Viewer filters |
| **Formatters** | JSON and text output for log export functionality |
| **Performance** | Minimal allocation; suitable for high-frequency build output |

### Usage Patterns

- **Application Logging:** Internal debug and operational logs
- **Maven Output Parsing:** Parse and colorize build output by level
- **Log Export:** JSON formatter for file export feature

---

## Testing Strategy

**Approach:** Hybrid (Structural + Visual Regression)

### Primary: teatest (github.com/charmbracelet/teatest)

| Capability | Description |
|------------|-------------|
| **Golden Files** | Snapshot testing for terminal output |
| **Model Testing** | Test state transitions without rendering |
| **Message Testing** | Verify correct messages sent |

### Secondary: Image Comparison

| Capability | Description |
|------------|-------------|
| **Visual Regression** | Pixel-level comparison for critical UI states |
| **97 UI States** | Coverage for all prototype screenshots |
| **CI Integration** | Automated comparison in GitHub Actions |

### Rationale

| Factor | Hybrid Advantage |
|--------|-----------------|
| **Structural Tests** | Fast, stable, catch logic errors |
| **Visual Tests** | Catch styling regressions missed by structural tests |
| **AI Development Safety** | Visual regression testing prevents AI-assisted development from introducing UI bugs |

### Canonical Terminal Size

- **Width:** 120 columns
- **Height:** 40 rows
- **Rationale:** Common developer terminal size; sufficient for complex UIs; consistent test baseline

---

## CI/CD

**Platform:** GitHub Actions

### Pipeline Stages

| Stage | Actions |
|-------|---------|
| **Build** | `go build` for all target platforms |
| **Test** | `go test` with race detection |
| **Visual Test** | teatest golden file comparison |
| **Lint** | `golangci-lint` static analysis |
| **Release** | GoReleaser for multi-platform binaries |

### Rationale

| Factor | GitHub Actions Advantage |
|--------|-------------------------|
| **Integration** | Native GitHub repository integration |
| **Matrix Builds** | Test across Go versions and OS targets |
| **Free Tier** | Sufficient for internal tool development |
| **Actions Ecosystem** | Pre-built actions for Go, releases |

### Environments

- **Development:** Local `go run` with mock data
- **Testing:** CI with `references/_test-data/demo-components/`
- **Production:** Single binary distribution (no server)

---

## Distribution

**Method:** Single Binary

### Build Targets

| Platform | Architecture | Output |
|----------|--------------|--------|
| macOS | amd64, arm64 | `rfz-darwin-*` |
| Linux | amd64 | `rfz-linux-amd64` |
| Windows | amd64 | `rfz-windows-amd64.exe` |

### Rationale

| Factor | Single Binary Advantage |
|--------|------------------------|
| **No Dependencies** | Go compiles to static binary; no runtime needed |
| **Easy Distribution** | Download and run; no installation process |
| **Versioning** | Binary contains version; easy to verify |
| **Rollback** | Keep old binary for instant rollback |

---

## Development Tools

### Required

| Tool | Purpose | Version |
|------|---------|---------|
| Go | Compiler | 1.21+ |
| Git | Version control | Latest |
| golangci-lint | Static analysis | Latest |

### Recommended

| Tool | Purpose |
|------|---------|
| GoReleaser | Multi-platform builds |
| direnv | Environment management |
| delve | Debugger |

---

## Project Dependencies Summary

```go
// go.mod core dependencies
require (
    github.com/charmbracelet/bubbletea v0.25+
    github.com/charmbracelet/bubbles v0.18+
    github.com/charmbracelet/lipgloss v0.9+
    github.com/charmbracelet/log v0.3+
)

// Testing
require (
    github.com/charmbracelet/teatest v0.0.1+ // Note: Check for stable release
)
```

---

## Critical Rule: Charm.land First - Custom Last

**This is a MANDATORY rule for all development.**

Before implementing ANY visual/UI element, developers MUST check if charm.land already provides it.

### Priority Order (MUST follow)

1. **Bubbles Component** - Use directly (list, table, viewport, progress, spinner, help, textinput, paginator)
2. **Lip Gloss Styling** - Use for ALL styling needs:
   - Borders: `lipgloss.Border()` with `NormalBorder()`, `RoundedBorder()`, `DoubleBorder()`, `ThickBorder()`
   - Colors: `lipgloss.Color()` and `lipgloss.AdaptiveColor()`
   - Layout: `lipgloss.Place()`, `lipgloss.JoinHorizontal()`, `lipgloss.JoinVertical()`
   - Padding/Margin: `.Padding()`, `.Margin()`
   - Text styling: `.Bold()`, `.Italic()`, `.Underline()`, `.Faint()`
3. **charmbracelet/log** - Use for all logging output
4. **Custom implementation** - ONLY when charm.land has NO solution

### Forbidden Patterns

| DO NOT | USE INSTEAD |
|--------|-------------|
| Custom border drawing with `---` or box-drawing chars | `lipgloss.NewStyle().Border(lipgloss.NormalBorder())` |
| Custom color codes/ANSI escapes | `lipgloss.Color("#ff0000")` or `lipgloss.AdaptiveColor{}` |
| Manual string padding | `lipgloss.NewStyle().Padding(1, 2)` |
| Custom progress bar strings | `bubbles/progress` component |
| Custom spinner frames | `bubbles/spinner` component |
| Custom list rendering | `bubbles/list` component |
| Custom table rendering | `bubbles/table` component |
| Custom scrolling logic | `bubbles/viewport` component |

### When Custom IS Allowed

Custom implementation is permitted ONLY for:

- **Business logic components** (TuiStatus badges, TuiNavItem with shortcuts)
- **Composite components** (TuiModal = viewport + Lip Gloss border + overlay logic)
- **Domain-specific rendering** (Maven build phases, Git status indicators)

**Even then, the custom component MUST use Lip Gloss for ALL styling internally.**

---

## Custom Components to Build

These components require custom implementation because Bubbles does not provide them:

| Component | Variants | Description | Uses |
|-----------|----------|-------------|------|
| TuiBox | single, double, rounded, heavy, focused | Container with focus state | Lip Gloss borders |
| TuiDivider | single, double | Horizontal/vertical separators | Lip Gloss |
| TuiNavigation | - | Sidebar navigation container | Lip Gloss layout |
| TuiNavItem | normal, focused, active | Menu item with shortcut | Lip Gloss styling |
| TuiStatus | pending, running, success, failed, error | Build status badges | Lip Gloss colors |
| TuiButton | default, primary, danger | Action buttons | Lip Gloss styling |
| TuiModal | - | Overlay dialog with backdrop | Lip Gloss borders + viewport |
| TuiRadio | - | Radio button group | Lip Gloss styling |
| TuiCheckbox | - | Checkbox with label | Lip Gloss styling |
| TuiTabs | - | Tab navigation | Lip Gloss styling |
| TuiStatusBar | - | Bottom bar with hints | Lip Gloss layout |
| TuiTree | - | Tree view for hierarchies | Lip Gloss + custom logic |

---

## Technology Decision Log

| Decision | Chosen | Alternatives Considered | Rationale |
|----------|--------|------------------------|-----------|
| Language | Go | Rust, Python | charm.land ecosystem is Go; single binary; team familiarity |
| TUI Framework | Bubble Tea | tview, termui, tcell | Elm architecture; teatest support; active maintenance |
| Styling | Lip Gloss | tcell styles, ANSI codes | Declarative; integrates with Bubble Tea |
| Testing | teatest + visual | go test only | Visual regression critical for 97 UI states |
| CI/CD | GitHub Actions | GitLab CI, Jenkins | Native integration; free tier sufficient |

---

## Not Applicable to This Project

The following web application concerns are not applicable to this desktop TUI:

- **Database:** No persistent database (config file only)
- **Frontend Framework:** N/A (TUI, not web)
- **CSS Framework:** N/A (Lip Gloss instead)
- **Hosting:** N/A (distributed binary)
- **CDN:** N/A (no web assets)
- **Asset Storage:** N/A (no uploaded files)

---

**Note:** This technical stack is optimized for the RFZ Developer CLI's specific requirements: a purpose-built terminal tool for RFZ developers with strong visual regression testing. All technology choices prioritize the charm.land ecosystem for consistency, testability, and development velocity.
