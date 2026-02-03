# Specification: Interactive Components (Sprint 1.2)

> Spec ID: 2026-02-03-interactive-components
> Created: 2026-02-03
> Status: Ready for Execution

---

## Overview

Build 6 interactive TUI components for the RFZ Developer CLI component library: TuiList, TuiCheckbox, TuiRadio, TuiTextInput, TuiSpinner, and TuiProgress. These components use charm-style Unicode symbols and wrap Bubbles library components where applicable, following the established patterns from Sprint 1.1.

---

## User Stories

| Story ID | Title | Type | Priority | Dependencies |
|----------|-------|------|----------|--------------|
| INTER-001 | TuiList Component | Frontend | High | INTER-002, INTER-003 |
| INTER-002 | TuiCheckbox Component | Frontend | High | None |
| INTER-003 | TuiRadio Component | Frontend | High | None |
| INTER-004 | TuiTextInput Component | Frontend | Medium | None |
| INTER-005 | TuiSpinner Component | Frontend | Medium | None |
| INTER-006 | TuiProgress Component | Frontend | Medium | None |
| INTER-007 | Extend Component Gallery | Frontend | High | INTER-001 through INTER-006 |
| INTER-008 | Visual Regression Tests | Test | High | INTER-007 |

---

## Spec Scope

### What's Included

- **TuiList**: Scrollable selection list with single/multi-select modes, `›` cursor
- **TuiCheckbox**: Toggle checkbox with `☐`/`☑` ballot box symbols
- **TuiRadio**: Radio button group with `◯`/`◉` circle symbols, horizontal/vertical layouts
- **TuiTextInput**: Bubbles textinput wrapper with RFZ styling
- **TuiSpinner**: Bubbles spinner wrapper with 4 variants (braille, line, circle, bounce)
- **TuiProgress**: Bubbles progress wrapper with 4 styles (block gradient, braille, ASCII, simple)
- **Gallery Extension**: Add 6 new sections to existing component demo
- **Visual Tests**: Golden file tests for all component states and variants

### What's NOT Included

- Screen-level implementations (Phase 2)
- Real data integration
- Theme customization
- Complex validation logic
- Accessibility features beyond keyboard navigation

---

## Expected Deliverables

### Files to Create

| File | Description |
|------|-------------|
| `internal/ui/components/checkbox.go` | TuiCheckbox component |
| `internal/ui/components/checkbox_test.go` | TuiCheckbox unit tests |
| `internal/ui/components/radio.go` | TuiRadio component |
| `internal/ui/components/radio_test.go` | TuiRadio unit tests |
| `internal/ui/components/list.go` | TuiList component |
| `internal/ui/components/list_test.go` | TuiList unit tests |
| `internal/ui/components/textinput.go` | TuiTextInput wrapper |
| `internal/ui/components/textinput_test.go` | TuiTextInput unit tests |
| `internal/ui/components/spinner.go` | TuiSpinner wrapper |
| `internal/ui/components/spinner_test.go` | TuiSpinner unit tests |
| `internal/ui/components/progress.go` | TuiProgress wrapper |
| `internal/ui/components/progress_test.go` | TuiProgress unit tests |

### Files to Modify

| File | Changes |
|------|---------|
| `internal/ui/components/styles.go` | Add charm symbol constants |
| `internal/ui/components/demo/gallery.go` | Add 6 new component sections |

### Testable Outcomes

- [ ] All 6 components compile without errors
- [ ] All unit tests pass
- [ ] All visual regression tests pass
- [ ] Component gallery displays all new components
- [ ] `golangci-lint` reports 0 issues

---

## Integration Requirements

**Integration Type:** Frontend-only (TUI component library)

### Integration Test Commands

```bash
# Build validation
go build ./cmd/rfz/...

# Unit tests for all components
go test ./internal/ui/components/... -v

# Visual regression tests
go test ./internal/ui/components/demo/... -v

# Lint check
golangci-lint run ./internal/ui/components/...
```

### End-to-End Scenarios

1. **Component Gallery Display**
   - Run `go run ./cmd/rfz/...`
   - All 10 component sections visible (4 existing + 6 new)
   - Scroll through gallery with j/k keys
   - Requires MCP: No

2. **Visual State Verification**
   - Run visual regression test suite
   - All golden files match expected output
   - Requires MCP: No

---

## Technical References

- **Design System:** `agent-os/product/design-system.md`
- **Tech Stack:** `agent-os/product/tech-stack.md`
- **Existing Components:** `internal/ui/components/` (box.go, button.go, status.go, divider.go)
- **Prototype Screenshots:** `references/prototype-screenshots/` (10-build-*, 20-config-*, 40-build-*)

---

## Notes

- All components follow "Charm.land First" rule from CLAUDE.md
- Use existing color tokens from styles.go
- Follow stateless render function pattern from existing components
- Wrap Bubbles components for spinner, progress, textinput
