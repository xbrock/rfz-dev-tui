# Implementation Plan: Core Components

**Created:** 2026-02-02
**Spec:** CORE - Core Components (Phase 1, Week 1)
**Status:** DRAFT

---

## Executive Summary

Build the foundational TUI component library from scratch: a comprehensive styles package implementing the design system, four core components (TuiBox, TuiDivider, TuiButton, TuiStatus), teatest visual testing infrastructure, and a component gallery demo. This replaces any existing boilerplate code with production-quality implementations.

---

## Architektur-Entscheidungen

### 1. Component Pattern: Stateless Render Functions
- Components are **stateless rendering functions** (not Bubble Tea models)
- Accept parameters, return `string` (rendered output)
- All styling via Lip Gloss (no ANSI codes, no manual padding)
- Reusable across screens without state management overhead
- Follow idiomatic Go patterns (functional options where appropriate)

### 2. Styles Architecture
- Single `styles.go` file for all design system tokens
- Exported package-level variables for colors, borders, spacing
- Pre-configured style templates for common patterns
- Helper functions for dynamic styling (e.g., focus states)

### 3. File Organization
```
internal/ui/components/
├── styles.go              # Design system tokens
├── helpers.go             # Utility functions (Truncate, etc.)
├── box.go                 # TuiBox component
├── divider.go             # TuiDivider component
├── button.go              # TuiButton component
├── status.go              # TuiStatus component
├── styles_test.go         # Style tests
├── box_test.go            # TuiBox tests + golden files
├── divider_test.go        # TuiDivider tests + golden files
├── button_test.go         # TuiButton tests + golden files
├── status_test.go         # TuiStatus tests + golden files
└── demo/
    ├── gallery.go         # Component gallery model
    └── gallery_test.go    # Gallery tests

testdata/golden/components/
├── box/                   # TuiBox golden files
├── divider/               # TuiDivider golden files
├── button/                # TuiButton golden files
├── status/                # TuiStatus golden files
└── gallery/               # Gallery golden files
```

### 4. Testing Strategy
- **teatest** for visual regression testing
- Golden files at canonical size 120x40
- One golden file per significant variant
- Tests run with `go test ./internal/ui/components/...`
- Update golden files with `-update` flag

### 5. Design System Compliance
- All colors from `design-system.md`
- All border styles (Single, Double, Rounded, Heavy)
- All spacing constants (XS through 2XL)
- Typography hierarchy (H1, H2, H3, Body)
- CRITICAL: Lip Gloss for ALL visual styling

---

## Komponenten-Übersicht

### Neue Komponenten

| Komponente | Datei | Beschreibung | Geschätzte LOC |
|------------|-------|--------------|----------------|
| Styles | `styles.go` | All design tokens | ~350 |
| Helpers | `helpers.go` | Truncate, width utils | ~50 |
| TuiBox | `box.go` | Bordered container | ~100 |
| TuiDivider | `divider.go` | Separators | ~60 |
| TuiButton | `button.go` | Buttons + shortcuts | ~120 |
| TuiStatus | `status.go` | Status badges | ~130 |
| ComponentGallery | `demo/gallery.go` | Demo screen | ~200 |
| Tests | `*_test.go` | All test files | ~400 |

**Total estimated:** ~1,410 LOC

### Komponenten-Verbindungen

| Source | Target | Verbindungstyp | Validierung | Zuständige Story |
|--------|--------|----------------|-------------|------------------|
| TuiBox | styles.go | Imports colors, borders | `grep -q "components.Color" box.go` | story-002 |
| TuiDivider | styles.go | Imports colors | `grep -q "components.Color" divider.go` | story-003 |
| TuiButton | styles.go | Imports button styles | `grep -q "components.Color" button.go` | story-004 |
| TuiStatus | styles.go | Imports colors | `grep -q "components.Color" status.go` | story-005 |
| box_test.go | TuiBox | Imports for testing | `go test ./internal/ui/components/` | story-006 |
| ComponentGallery | All components | Renders all | Gallery shows all variants | story-007 |

---

## Umsetzungsphasen

### Phase A: Foundation (Story 001)
**Styles Package**
- Implement all color tokens
- Implement border constants
- Implement spacing constants
- Implement typography styles
- Implement helper functions

**Why first:** All components depend on styles.

### Phase B: Core Components (Stories 002-005)
**Parallel implementation possible**
- TuiBox - depends only on styles
- TuiDivider - depends only on styles
- TuiButton - depends only on styles
- TuiStatus - depends only on styles

**Why parallel:** No dependencies between components.

### Phase C: Testing Infrastructure (Story 006)
**teatest Setup**
- Add teatest to go.mod
- Create test harness
- Generate golden files for all components
- Configure canonical size (120x40)

**Why after components:** Need components to test.

### Phase D: Demo (Story 007)
**Component Gallery**
- Create Bubble Tea model
- Render all component variants
- Add keyboard navigation
- Generate gallery golden files

**Why last:** Needs all components + testing.

---

## Abhängigkeiten

```
                    ┌─────────────┐
                    │  story-001  │
                    │   Styles    │
                    └──────┬──────┘
                           │
       ┌───────────────────┼───────────────────┐
       │                   │                   │
       ▼                   ▼                   ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  story-002  │     │  story-003  │     │  story-004  │
│   TuiBox    │     │ TuiDivider  │     │  TuiButton  │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       │            ┌──────┴──────┐            │
       │            │  story-005  │            │
       │            │  TuiStatus  │            │
       │            └──────┬──────┘            │
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  story-006   │
                    │   teatest    │
                    │ Infrastructure│
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  story-007   │
                    │  Component   │
                    │   Gallery    │
                    └──────────────┘
```

- Story 001: No dependencies (foundation)
- Stories 002-005: Depend on story-001 (can run parallel after 001)
- Story 006: Depends on 001-005
- Story 007: Depends on story-006

---

## Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| teatest API instability | Low | Medium | Pin version, check docs for breaking changes |
| Golden file terminal differences | Medium | Medium | Document exact terminal settings for tests |
| go mod dependency issues | Low | Low | Run `go mod tidy`, verify import paths |
| Design system color accuracy | Low | Medium | Verify hex codes against design-system.md |
| Lip Gloss version compatibility | Low | High | Pin to v0.9.1 as specified in tech stack |

---

## Self-Review Ergebnisse

### 1. VOLLSTÄNDIGKEIT
- [x] Styles Package - All design tokens covered
- [x] TuiBox - 5 variants + focused states
- [x] TuiDivider - Single/double variants
- [x] TuiButton - 3 variants + shortcut support
- [x] TuiStatus - 5 states + compact variant
- [x] teatest infrastructure - Golden file setup
- [x] Component gallery - Demo screen
- [x] Content overflow handling - Truncate helper

### 2. KONSISTENZ
- [x] All components use same pattern (stateless functions)
- [x] All components depend on styles.go
- [x] All components have corresponding tests
- [x] File naming is consistent (component.go, component_test.go)

### 3. RISIKEN
- [x] All risks identified with mitigations
- [x] No external service dependencies
- [x] No security concerns (UI-only)

### 4. ALTERNATIVEN
**Considered:** Splitting styles into multiple files (colors.go, typography.go)
**Decision:** Single styles.go - easier to import, matches design-system.md structure

**Considered:** Making components Bubble Tea models
**Decision:** Stateless functions - simpler, more reusable, consistent pattern

### 5. KOMPONENTEN-VERBINDUNGEN
- [x] JEDE neue Komponente ist mit styles.go verbunden
- [x] JEDE Verbindung ist einer Story zugeordnet
- [x] Keine verwaisten Komponenten
- [x] Gallery verbindet alle Komponenten

---

## Minimalinvasiv-Optimierungen

### 1. WIEDERVERWENDUNG
| Design System Element | Implementation |
|-----------------------|----------------|
| Color tokens (design-system.md) | Direct Lip Gloss Color definitions |
| Border styles | Lip Gloss built-in borders |
| Typography | Lip Gloss style methods |
| Spacing | Go constants |

**Insight:** The design-system.md already provides Lip Gloss code examples - these can guide implementation directly.

### 2. ÄNDERUNGSUMFANG
| Aktion | Dateien | Begründung |
|--------|---------|------------|
| CREATE | 8 source files | Core implementation |
| CREATE | 7 test files | Test coverage |
| CREATE | Golden file directories | Testing infrastructure |
| UPDATE | go.mod | Add teatest dependency |

**No modifications to existing code** - this is a fresh implementation.

### 3. FEATURE-PRESERVATION
- [x] Alle Requirements aus Clarification v2 sind abgedeckt
- [x] Kein Feature wurde geopfert
- [x] Alle Akzeptanzkriterien bleiben erfüllbar

---

## Technische Details

### Styles Package API
```go
// Colors
var ColorBackground = lipgloss.Color("#1e1e2e")
var ColorCyan = lipgloss.Color("#0891b2")
// ... all tokens from design-system.md

// Borders
var BorderSingle = lipgloss.NormalBorder()
var BorderDouble = lipgloss.DoubleBorder()
// ... all border types

// Spacing
const SpaceXS = 1
const SpaceSM = 2
// ... all spacing constants

// Pre-built styles
var StyleBoxDefault = lipgloss.NewStyle()...
var StyleButtonPrimary = lipgloss.NewStyle()...
```

### TuiBox API
```go
type BoxStyle string
const (
    BoxSingle  BoxStyle = "single"
    BoxDouble  BoxStyle = "double"
    BoxRounded BoxStyle = "rounded"
    BoxHeavy   BoxStyle = "heavy"
)

func TuiBox(content string, style BoxStyle, focused bool) string
func TuiBoxWithWidth(content string, style BoxStyle, focused bool, width int) string
```

### TuiDivider API
```go
type DividerStyle string
const (
    DividerSingle DividerStyle = "single"
    DividerDouble DividerStyle = "double"
)

func TuiDivider(style DividerStyle, width int) string
```

### TuiButton API
```go
type ButtonVariant string
const (
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonDestructive ButtonVariant = "destructive"
)

func TuiButton(label string, variant ButtonVariant, shortcut string, focused bool) string
```

### TuiStatus API
```go
type Status int
const (
    StatusPending Status = iota
    StatusRunning
    StatusSuccess
    StatusFailed
    StatusError
)

func TuiStatus(status Status) string
func TuiStatusCompact(status Status) string
```

### Helper Functions
```go
func Truncate(text string, maxWidth int) string
func GetTerminalWidth() int  // or use lipgloss.Width
```

---

## Validierung

### Integration Checks
```bash
# Build all components
go build ./internal/ui/components/...

# Run all tests
go test ./internal/ui/components/...

# Lint check
golangci-lint run ./internal/ui/components/...

# Verify no boilerplate code remains
# (manually check that implementations match this plan)
```

### Golden File Verification
```bash
# Generate/update golden files
go test ./internal/ui/components/... -update

# Verify golden files match
go test ./internal/ui/components/...
```

---

*Plan Status: DRAFT - Awaiting User Approval*
