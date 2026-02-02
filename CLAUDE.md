# CLAUDE.md - RFZ Developer CLI

> RFZ Developer CLI Development Guide
> Last Updated: 2026-02-02
> Type: Go TUI Application (charm.land stack)

## Purpose

Essential guidance for Claude Code development on the RFZ Developer CLI - a Go-based Terminal User Interface for RFZ developers at Deutsche Bahn that streamlines Maven build workflows.

---

## CRITICAL RULE: Charm.land First - Custom Last

**MANDATORY for ALL agents. This rule OVERRIDES all other styling/component decisions.**

Before implementing ANY visual/UI element, you MUST check if charm.land provides it.

### Priority Order (MUST follow)

1. **Bubbles Component** - Use directly: `list`, `table`, `viewport`, `progress`, `spinner`, `help`, `textinput`, `paginator`
2. **Lip Gloss Styling** - Use for ALL styling (borders, colors, layout, padding, text)
3. **charmbracelet/log** - Use for all logging output
4. **Custom implementation** - ONLY when charm.land has NO solution

### Forbidden Patterns

| DO NOT | USE INSTEAD |
|--------|-------------|
| Custom border drawing (`---`, box-drawing chars) | `lipgloss.NewStyle().Border(lipgloss.RoundedBorder())` |
| Custom color codes / ANSI escapes | `lipgloss.Color("#ff0000")` |
| Manual string padding | `lipgloss.NewStyle().Padding(1, 2)` |
| Custom progress bar strings | `bubbles/progress` component |
| Custom spinner frames | `bubbles/spinner` component |
| Custom list rendering | `bubbles/list` component |
| Custom table rendering | `bubbles/table` component |
| Custom scrolling logic | `bubbles/viewport` component |

**Even custom components MUST use Lip Gloss for ALL internal styling.**

---

## Tech Stack

| Technology | Purpose |
|------------|---------|
| **Go 1.21+** | Language |
| **Bubble Tea** | TUI Framework (Elm architecture) |
| **Lip Gloss** | ALL styling (borders, colors, layout) |
| **Bubbles** | Components (list, table, viewport, progress, spinner, help) |
| **charmbracelet/log** | Structured logging |
| **teatest** | Visual regression testing |

Full details: `agent-os/product/tech-stack.md`

---

## Document Locations (Load on Demand)

### Product Information
- **Product Vision**: `agent-os/product/product-brief.md`
- **Product Summary**: `agent-os/product/product-brief-lite.md`
- **Technical Stack**: `agent-os/product/tech-stack.md`
- **Design System**: `agent-os/product/design-system.md` (colors, styles, spacing)
- **Development Roadmap**: `agent-os/product/roadmap.md`

### Development Standards
- **Tech Stack Defaults**: `agent-os/standards/tech-stack.md`
- **Code Style**: `agent-os/standards/code-style.md`
- **Best Practices**: `agent-os/standards/best-practices.md`

---

## Critical Rules

- **FOLLOW ALL INSTRUCTIONS** - Mandatory, not optional
- **ASK FOR CLARIFICATION** - If uncertain about any requirement
- **MINIMIZE CHANGES** - Edit only what's necessary
- **CHARM.LAND FIRST** - Always use Bubbles/Lip Gloss before custom code
- **NO CO-AUTHORED COMMITS** - Never add "Co-Authored-By" lines to commit messages

---

## Components Strategy

### Use from Bubbles (DO NOT reimplement)

| Component | Usage in RFZ CLI |
|-----------|------------------|
| `list` | Component selection, navigation lists |
| `table` | Component registry, detected components |
| `viewport` | Log viewer content area |
| `progress` | Build progress indicators |
| `spinner` | Loading states during builds |
| `textinput` | Configuration inputs, search |
| `help` | Keyboard shortcut display |
| `paginator` | Log pagination |

### Custom Components to Build (using Lip Gloss internally)

| Component | Description |
|-----------|-------------|
| TuiBox | Container with focus state (Lip Gloss borders) |
| TuiNavigation | Sidebar navigation container |
| TuiNavItem | Menu item with keyboard shortcut |
| TuiStatus | Build status badges (pending/running/success/failed) |
| TuiButton | Action buttons |
| TuiModal | Overlay dialog with backdrop |
| TuiRadio | Radio button group |
| TuiCheckbox | Checkbox with label |
| TuiTabs | Tab navigation |
| TuiStatusBar | Bottom bar with hints |
| TuiTree | Tree view for hierarchies |

---

## Sub-Agents

### Utility & Support
- **context-fetcher** - Load documents on demand
- **date-checker** - Determine today's date
- **file-creator** - Create files and apply templates
- **git-workflow** - Git operations, commits, PRs

---

## File Organization Rules

**CRITICAL - No Files in Project Root:**
- Implementation reports: `agent-os/specs/[spec-name]/implementation-reports/`
- Architecture docs: `agent-os/product/`
- Team docs: `agent-os/team/`

---

## References Directory

The `references/` directory contains project reference materials:

| Path | Content |
|------|---------|
| `references/prototype-screenshots/` | 77 UI screenshots showing all screens (welcome, navigation, build, config, logs, discover views). Files numbered sequentially: `01-welcome-default.png`, `10-build-*.png`, `20-config-*.png`, etc. |
| `references/tui-web-prototype/` | Visual design prototype (built with v0.dev). Use as UI/UX design reference only - actual implementation is Go TUI. |
| `references/user-flow-diagrams.md` | User flow diagrams documentation |
| `references/_test-data/TREE.txt` | Directory tree structure |
| `references/_test-data/WORKFLOW.md` | Workflow documentation |
| `references/_test-data/demo-components/` | Demo components (core, simulator, standalone) |

**Usage:** Read screenshots with the Read tool to see UI designs. The web prototype is for visual reference only (created via v0.dev) - this project is a Go TUI application.

---

## Quality Requirements

**Mandatory Checks:**
- Run linting after ALL code changes (`golangci-lint`)
- ALL lint errors must be fixed before task completion
- Visual regression tests must pass (97 UI states)

**Canonical Terminal Size for Testing:**
- Width: 120 columns
- Height: 40 rows

---

## Agent OS Commands

```bash
# Product Planning
/plan-product            # Single-product planning

# Feature Development
/create-spec             # Create detailed specifications
/execute-tasks           # Execute planned tasks

# Bug Management
/create-bug              # Create bug specification
/add-bug                 # Add bug to existing spec

# Documentation
/retroactive-doc         # Document existing features
```

---

**Remember:** This is a Go TUI application. ALL visual elements must use the charm.land stack. Check Bubbles and Lip Gloss FIRST before writing any custom rendering code.
