# Core Components Specification

> Spec ID: CORE
> Created: 2026-02-02
> Phase: Phase 1 - Foundation (Week 1)
> Status: Ready for Implementation

---

## Overview

Build the foundational TUI component library from scratch: a comprehensive styles package implementing the design system, four core components (TuiBox, TuiDivider, TuiButton, TuiStatus), teatest visual testing infrastructure, and a component gallery demo.

## User Stories

| ID | Title | Type | Dependencies |
|----|-------|------|--------------|
| CORE-001 | Styles Package | Backend | None |
| CORE-002 | TuiBox Component | Backend | CORE-001 |
| CORE-003 | TuiDivider Component | Backend | CORE-001 |
| CORE-004 | TuiButton Component | Backend | CORE-001 |
| CORE-005 | TuiStatus Component | Backend | CORE-001 |
| CORE-006 | teatest Infrastructure | Test | CORE-001 to CORE-005 |
| CORE-007 | Component Gallery | Backend | CORE-006 |

## Scope

### Included

- Complete styles package with all design system tokens
- TuiBox with 5 border variants (Single, Double, Rounded, Heavy, Focused)
- TuiDivider with 2 variants (Single, Double)
- TuiButton with 3 variants (Primary, Secondary, Destructive) + shortcut support
- TuiStatus with 5 states (Pending, Running, Success, Failed, Error) + compact variant
- teatest golden file testing infrastructure
- Component gallery demo screen
- Content overflow handling (truncation)
- Unit tests with golden file comparison

### Excluded

- Interactive components (TuiList, TuiCheckbox, TuiRadio) - Phase 1 Week 2
- Navigation components (TuiNavigation, TuiNavItem) - Phase 1 Week 3
- Modal/overlay components - Phase 1 Week 3
- Animation/spinner components - Phase 1 Week 2
- Progress bar component - Phase 1 Week 2
- Screen implementations - Phase 2

## Expected Deliverables

### Files Created

```
internal/ui/components/
├── styles.go          # Design system tokens (~350 LOC)
├── helpers.go         # Utility functions (~50 LOC)
├── box.go             # TuiBox component (~100 LOC)
├── divider.go         # TuiDivider component (~60 LOC)
├── button.go          # TuiButton component (~120 LOC)
├── status.go          # TuiStatus component (~130 LOC)
├── styles_test.go     # Style tests
├── box_test.go        # TuiBox tests
├── divider_test.go    # TuiDivider tests
├── button_test.go     # TuiButton tests
├── status_test.go     # TuiStatus tests
└── demo/
    ├── gallery.go     # Component gallery (~200 LOC)
    └── gallery_test.go

testdata/golden/components/
├── box/               # TuiBox golden files
├── divider/           # TuiDivider golden files
├── button/            # TuiButton golden files
├── status/            # TuiStatus golden files
└── gallery/           # Gallery golden files
```

### Testable Outcomes

1. All components compile: `go build ./internal/ui/components/...`
2. All tests pass: `go test ./internal/ui/components/...`
3. No lint errors: `golangci-lint run ./internal/ui/components/...`
4. Golden files generated and matching

## Integration Requirements

**Integration Type:** Backend-only (Go TUI library)

### Integration Test Commands

```bash
# Build verification
go build ./internal/ui/components/...

# Unit tests with golden files
go test ./internal/ui/components/... -v

# Lint check
golangci-lint run ./internal/ui/components/...
```

### End-to-End Scenarios

1. **Component rendering** - Each component renders correctly at 120x40
2. **Golden file matching** - All variants match their golden files
3. **Gallery demo** - Gallery screen shows all component variants

---

*Detailed stories in: stories/*
