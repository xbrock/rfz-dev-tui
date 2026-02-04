# Requirements Clarification - Interactive Components (Sprint 1.2)

**Created:** 2026-02-03
**Status:** Pending User Approval

## Feature Overview

Implement all 6 interactive TUI components for Sprint 1.2 of the RFZ Developer CLI: TuiList, TuiCheckbox, TuiRadio, TuiTextInput, TuiSpinner, and TuiProgress. These components will use charm-style Unicode symbols and wrap Bubbles library components where applicable.

## Target Users

RFZ developers at Deutsche Bahn who need intuitive, keyboard-driven form controls and feedback indicators for the build configuration interface.

## Business Value

These interactive components are foundational building blocks required for Phase 2 screens (Build Components, Build Configuration Modal, Build Execution). Without them, the core user flows cannot be implemented. They enable efficient component selection, build option configuration, and real-time build progress feedback.

## Functional Requirements

### TuiList - Scrollable Selection List
- Support both single-select (radio-like) and multi-select (checkbox-like) modes
- Keyboard navigation: ↑/↓ or j/k for movement, Space/Enter for selection
- Cursor indicator using `›` (guillemet) charm-style symbol
- Support for item badges (like "Core", "Simulation", "Standalone" in build screen)
- Scrollable viewport for lists longer than display area
- Focus state with cyan border highlight
- Selection counter display (e.g., "3/13 selected")

### TuiCheckbox - Toggle Checkbox
- Charm-style symbols: `☐` unchecked, `☑` checked
- Label text with optional description
- Keyboard toggle: Space or Enter
- States: default, focused, checked, disabled
- Support for indeterminate state (optional, for "select all" scenarios)
- Cyan highlight for focused state

### TuiRadio - Radio Button Group
- Charm-style symbols: `◯` unselected, `◉` selected
- Horizontal and vertical layout options
- Single selection within group (selecting one deselects others)
- Keyboard navigation: ←/→ or h/l for horizontal, ↑/↓ or j/k for vertical
- Focus indicator with cyan highlight/border
- Support for disabled options

### TuiTextInput - Text Input Field
- Wrap Bubbles `textinput` component with RFZ styling
- Placeholder text support (muted style)
- States: default, focused, disabled, error
- Cursor blinking animation
- Character limit display (optional)
- Prompt prefix support (e.g., `$` or `>`)

### TuiSpinner - Loading Indicator
- Wrap Bubbles `spinner` component with custom frame sets
- 4 variants as per roadmap:
  - **Braille dots**: `⠈⠙⠛⠻⠹⠸⠰⠠` (primary, smooth animation)
  - **Line**: `|/-\` (fallback for limited terminals)
  - **Circle quarters**: `◴◷◶◵` (geometric option)
  - **Bounce**: `⠁⠂⠄⠠⠐⠈` (vertical bounce)
- Configurable speed
- Optional label text beside spinner
- Color variants: cyan (default), green (success), yellow (warning)

### TuiProgress - Progress Bar
- Wrap Bubbles `progress` component with custom rendering
- 4 styles as per roadmap:
  - **Block gradient**: `█▓▒░` (primary, smooth shading)
  - **Braille**: Ultra-smooth braille-based fill
  - **ASCII**: `[=====>    ]` (fallback)
  - **Simple**: `████████░░░░` (no gradient)
- Percentage display option
- Color gradient from yellow (0%) to green (100%)
- Determinate and indeterminate modes
- Width configuration

## Affected Areas & Dependencies

| Component/Area | Impact Description |
|----------------|-------------------|
| `internal/ui/components/` | New component files added |
| `internal/ui/components/styles.go` | May need additional style tokens |
| `internal/ui/components/demo/gallery.go` | Extended to showcase new components |
| Bubbles library | Dependency for spinner, progress, textinput wrappers |
| Existing TuiBox, TuiButton, TuiStatus | Must integrate visually (same design language) |

## Edge Cases & Error Scenarios

| Edge Case | Expected Behavior |
|-----------|-------------------|
| Empty list | Display "No items" message, disable selection |
| List with 100+ items | Smooth scrolling, viewport pagination |
| Very long item labels | Truncate with `...` using existing `Truncate()` helper |
| Disabled checkbox/radio in focus | Skip disabled items during navigation |
| Progress at exactly 0% or 100% | Distinct visual states (empty bar / full bar) |
| Spinner in non-animated context | Show static first frame for tests |
| Text input at max length | Prevent further input, show limit indicator |

## Security & Permissions

Not applicable - these are UI components with no security implications.

## Performance Considerations

- Spinner animation should not cause excessive CPU usage (use Bubble Tea tick messages)
- List rendering for 100+ items should remain smooth
- Progress bar updates should be throttled if called rapidly
- All components should render in < 1ms for visual test stability

## Scope Boundaries

**IN SCOPE:**
- All 6 interactive components (TuiList, TuiCheckbox, TuiRadio, TuiTextInput, TuiSpinner, TuiProgress)
- Charm-style Unicode symbols (◯/◉, ☐/☑, ›, braille, block gradient)
- Wrapping Bubbles components (spinner, progress, textinput)
- Full state coverage (default, focused, selected, disabled, error)
- Golden file visual tests for all variants and states
- Extending component gallery demo
- Unit tests for all components

**OUT OF SCOPE:**
- Screen-level implementations (those are Phase 2)
- Real data integration (using mock data only)
- Accessibility features beyond standard keyboard navigation
- Theme customization (using fixed design system)
- Complex validation logic (just visual states)

## Open Questions

None - all requirements clarified through dialog.

## Proposed User Stories (High Level)

1. **TuiList Component** - Implement scrollable list with single/multi-select modes and charm-style cursor
2. **TuiCheckbox Component** - Implement checkbox with ballot box symbols and all states
3. **TuiRadio Component** - Implement radio button group with circle symbols and layout options
4. **TuiTextInput Component** - Wrap Bubbles textinput with RFZ styling
5. **TuiSpinner Component** - Wrap Bubbles spinner with 4 custom frame variants
6. **TuiProgress Component** - Wrap Bubbles progress with 4 custom styles
7. **Extend Component Gallery** - Add new interactive components to existing demo
8. **Visual Regression Tests** - Golden file tests for all components and states

---

*Review this document carefully. Once approved, detailed user stories will be generated.*
