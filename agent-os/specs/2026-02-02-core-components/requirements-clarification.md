# Requirements Clarification - Core Components

**Created:** 2026-02-02
**Updated:** 2026-02-02 (Revision 2)
**Status:** Pending User Approval
**Spec ID:** CORE
**Phase:** Phase 1 - Foundation (Week 1)

## Feature Overview

Build the foundational TUI component library **from scratch**: styles package (colors, typography, borders, spacing), core components (TuiBox, TuiDivider, TuiButton, TuiStatus), plus teatest visual testing infrastructure with golden file comparison.

**Important:** Existing boilerplate code in the codebase should be treated as inspiration/reference only - this spec implements the production-quality component library.

## Target Users

RFZ developers at Deutsche Bahn who need a consistent, visually tested component library as the foundation for the RFZ Developer CLI screens.

## Business Value

- **Foundation for all screens:** Every screen in Phase 2 depends on these core components
- **Visual regression safety:** teatest infrastructure enables confident AI-assisted development
- **Design consistency:** Centralized styles ensure uniform look and feel
- **Developer efficiency:** Reusable components reduce duplication and errors

## Functional Requirements

### FR1: Styles Package (NEW)
- All color tokens from design-system.md
- Typography styles (H1, H2, H3, Body, Muted, Code, Keyboard)
- Border style constants (Single, Double, Rounded, Heavy)
- Spacing constants (XS through 2XL)
- Box styles for different border types
- Button styles (Primary, Secondary, Destructive)
- Input styles (Normal, Focused, Error)
- Navigation styles
- Log/status styles
- Helper functions for common patterns

### FR2: TuiBox Component (NEW)
- Bordered container component with configurable border styles
- 5 border variants: Single, Double, Rounded, Heavy, Focused
- Each variant supports focused state (cyan border highlight)
- Configurable padding and width
- Content overflow handling with truncation/ellipsis

### FR3: TuiDivider Component (NEW)
- Horizontal separator lines
- 2 variants: Single line (─) and Double line (═)
- Configurable width (auto-fill or fixed)
- Uses design system border color

### FR4: TuiButton Component (NEW)
- 3 visual variants: Default (outlined), Primary (filled cyan), Danger (filled red)
- Support for optional keyboard shortcut display (e.g., "[Enter] Build")
- Focused state with visual indicator
- Disabled state with muted styling

### FR5: TuiStatus Component (NEW)
- Status badge for build/process states
- 5 states: Pending (gray), Running (cyan), Success (green), Failed (red), Error (red)
- Full badge format with colored background
- Compact format (single character) for list views
- Clear visual distinction between states

### FR6: teatest Infrastructure (NEW)
- Golden file test setup for visual regression
- Canonical terminal size: 120x40
- Golden file storage structure (`testdata/golden/components/`)
- Test harness for rendering components at canonical size
- Update mechanism for golden files
- Integration with `go test`

### FR7: Component Gallery/Demo (NEW)
- Runnable demo screen showing all component variants
- Visual verification tool during development
- Demonstrates all states and variants
- Keyboard navigation through sections

## Affected Areas & Dependencies

| Component | Impact |
|-----------|--------|
| `internal/ui/components/styles.go` | NEW - Complete styles package |
| `internal/ui/components/box.go` | NEW - TuiBox component |
| `internal/ui/components/divider.go` | NEW - TuiDivider component |
| `internal/ui/components/button.go` | NEW - TuiButton component |
| `internal/ui/components/status.go` | NEW - TuiStatus component |
| `internal/ui/components/demo/gallery.go` | NEW - Component gallery |
| `internal/ui/components/*_test.go` | NEW - Component tests |
| `go.mod` | UPDATE - Ensure charm.land deps + teatest |
| `testdata/golden/components/` | NEW - Golden file storage |

### External Dependencies
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/charmbracelet/bubbles` - Base components (key bindings)
- `github.com/charmbracelet/x/exp/teatest` - Visual testing

## Edge Cases & Error Scenarios

| Edge Case | Expected Behavior |
|-----------|-------------------|
| Content longer than box width | Truncate with ellipsis ("...") |
| Empty content in TuiBox | Render box with minimum height |
| Very long button labels | Truncate label, preserve shortcut |
| Zero-width divider | Default to terminal width |
| Invalid status state | Fallback to "Pending" style |

## Security & Permissions

Not applicable - this is a UI component library with no external data access.

## Performance Considerations

- Components should render in < 1ms
- No memory leaks from style objects
- Styles should be reusable (not recreated on each render)

## Scope Boundaries

**IN SCOPE:**
- Complete styles package with all design system tokens
- TuiBox with 5 border variants + focused states
- TuiDivider with single/double variants
- TuiButton with 3 variants + shortcut support
- TuiStatus with 5 status states + compact variant
- teatest golden file infrastructure
- Component gallery demo screen
- Content overflow handling (truncation)
- Unit tests with golden file comparison

**OUT OF SCOPE:**
- Interactive components (TuiList, TuiCheckbox, TuiRadio) - Phase 1 Week 2
- Navigation components (TuiNavigation, TuiNavItem) - Phase 1 Week 3
- Modal/overlay components - Phase 1 Week 3
- Animation/spinner components - Phase 1 Week 2
- Progress bar component - Phase 1 Week 2
- Screen implementations - Phase 2
- Existing boilerplate code (to be replaced by this implementation)

## Open Questions

None - requirements are complete based on dialog.

## Proposed User Stories (High Level)

1. **Styles Package** - Complete design system tokens and helper functions
2. **TuiBox Component** - Container with 5 border variants and focus states
3. **TuiDivider Component** - Horizontal separators (single/double)
4. **TuiButton Component** - Buttons with variants and shortcut support
5. **TuiStatus Component** - Status badges for 5 states + compact variant
6. **teatest Infrastructure** - Golden file testing setup and harness
7. **Component Gallery** - Demo screen showing all variants

---

*Review this document carefully. Once approved, I will create an Implementation Plan treating this as a fresh implementation (not building on boilerplate).*
