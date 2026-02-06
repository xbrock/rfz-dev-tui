# Specification: Layout & Navigation Components

> Spec ID: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06
> Status: Ready for Execution

---

## Overview

Implementation of 8 Layout & Navigation TUI components for the RFZ Developer CLI, completing Sprint 1.3 of the component library foundation. These components enable screen navigation, data display, overlay dialogs, and keyboard hint visualization using the charm.land stack.

---

## User Stories

| ID | Title | Priority | Type | Est. Effort | Dependencies |
|----|-------|----------|------|-------------|--------------|
| LAYOUT-001 | TuiNavigation + TuiNavItem | High | Frontend | S | None |
| LAYOUT-002 | TuiModal | High | Frontend | M | None |
| LAYOUT-003 | TuiKeyHints | High | Frontend | XS | None |
| LAYOUT-004 | TuiTable | High | Frontend | S | None |
| LAYOUT-005 | TuiTree | Medium | Frontend | M | None |
| LAYOUT-006 | TuiTabs | High | Frontend | S | None |
| LAYOUT-007 | TuiStatusBar | High | Frontend | S | LAYOUT-003 |
| LAYOUT-008 | Layout Navigation Demo | Medium | Frontend | S | LAYOUT-001 to 007 |
| LAYOUT-009 | Visual Regression Tests | High | Test | S | LAYOUT-001 to 007 |
| LAYOUT-997 | Code Review by Opus | High | Review | S | LAYOUT-001 to 009 |
| LAYOUT-998 | Integration Validation | High | Validation | XS | LAYOUT-997 |
| LAYOUT-999 | Finalize and Create PR | High | Release | XS | LAYOUT-998 |

---

## Spec Scope

### In Scope

- **8 Layout Components:**
  - TuiNavigation (sidebar container)
  - TuiNavItem (navigation item)
  - TuiModal (overlay dialog with focus trapping)
  - TuiTable (bubbles/table wrapper)
  - TuiTabs (tab navigation with numeric shortcuts)
  - TuiStatusBar (bottom status bar)
  - TuiTree (simple expand/collapse)
  - TuiKeyHints (keyboard shortcuts display)

- **Testing:**
  - Unit tests for all components
  - Golden file tests for visual regression
  - Separate demo program for showcase

- **Patterns:**
  - Bubbles wrapper pattern (TuiTable)
  - Pure Lipgloss pattern (all others)
  - Composition pattern (TuiNavigation → TuiNavItem)

### Out of Scope

- Screen implementation (Phase 2)
- Real data integration
- Extending existing Component Gallery
- Horizontal navigation variant
- Multi-select tree
- Drag-and-drop functionality

---

## Expected Deliverables

### New Files

```
internal/ui/components/
├── navigation.go          # TuiNavigation + TuiNavItem
├── navigation_test.go     # Tests
├── modal.go               # TuiModal
├── modal_test.go          # Tests
├── table.go               # TuiTable (bubbles wrapper)
├── table_test.go          # Tests
├── tabs.go                # TuiTabs
├── tabs_test.go           # Tests
├── statusbar.go           # TuiStatusBar
├── statusbar_test.go      # Tests
├── tree.go                # TuiTree
├── tree_test.go           # Tests
├── keyhints.go            # TuiKeyHints
├── keyhints_test.go       # Tests
└── demo/
    └── layout_gallery.go  # Separate Demo

cmd/
└── layout-demo/
    └── main.go            # Demo entry point
```

### Modified Files

```
internal/ui/components/
└── styles.go              # New style definitions
```

---

## Integration Requirements

**Integration Type:** Frontend-only

**Integration Tests:**

```bash
# All components build
go build ./internal/ui/components/...

# All tests pass
go test ./internal/ui/components/... -v

# Demo builds and runs
go build ./cmd/layout-demo/...

# Lint passes
golangci-lint run ./internal/ui/components/...
```

**End-to-End Scenarios:**

1. **Navigation Demo:** Start layout-demo, navigate through sidebar items
2. **Modal Demo:** Open modal, tab through buttons, close with Escape
3. **Tree Demo:** Expand/collapse nodes with Enter
4. **Table Demo:** Scroll through rows, select with Enter

---

## Technical Approach

### Pattern 1: Bubbles Wrapper (TuiTable)
- Wraps `bubbles/table`
- Applies RFZ styling via Lipgloss
- Simplified API for common use cases

### Pattern 2: Pure Lipgloss (Navigation, Modal, Tabs, StatusBar, Tree, KeyHints)
- Stateless render functions
- State managed by parent Bubble Tea Model
- Consistent with existing components

### Pattern 3: Composite Component (TuiNavigation)
- Combines TuiNavItem elements
- Container with optional header/footer

---

## Dependencies

### External
- `github.com/charmbracelet/bubbles/table` - For TuiTable
- `github.com/charmbracelet/bubbles/viewport` - For scrollable content
- `github.com/charmbracelet/lipgloss` - All styling

### Internal
- `internal/ui/components/styles.go` - Color tokens, border styles
- `internal/ui/components/helpers.go` - Truncate utility
- `internal/ui/components/button.go` - For Modal footer

---

## Quality Gates

- [ ] All 12 stories completed (9 regular + 3 system)
- [ ] All unit tests passing
- [ ] All golden file tests passing
- [ ] Demo program functional
- [ ] Lint checks passing
- [ ] Code review approved (LAYOUT-997)
- [ ] Integration validation passed (LAYOUT-998)
- [ ] PR created (LAYOUT-999)

---

*Generated with Agent OS /create-spec v3.0*
