# Requirements Clarification - Layout Matching Design

**Created:** 2026-02-09
**Status:** Pending User Approval

## Feature Overview

Fix the RFZ-CLI TUI layout and styling across all implemented screens to match the design prototype. The current implementation has numerous visual discrepancies in the title bar, navigation, status bar, welcome screen, build components screen, config modal, and build execution view.

## Target Users

RFZ developers at Deutsche Bahn who use the CLI daily for Maven build workflows. Visual consistency with the approved design ensures a professional, intuitive tool.

## Business Value

The TUI must match the approved design prototype to maintain visual coherence, brand alignment (Deutsche Bahn), and usability standards. Inconsistent UI creates confusion and reduces developer trust in the tool.

## Functional Requirements

### 1. Title Bar (Header)
- Red accent line must be at the TOP of the app (above title content), not below it
- Line 1: "RFZ-CLI v1.0.0" (main title, bold)
- Line 2 left: "Terminal Orchestration Tool" (cyan); Line 2 right: "HH:MM:SS | Deutsche Bahn Internal"
- Both subtitle and time/branding on the SAME line

### 2. Navigation Sidebar
- Active nav item: light blue background color
- Nav item shortcuts: aligned to the RIGHT side of the nav box
- State priority: Active state OVERRIDES select state
  - Selected + Active = active bg + arrow prefix
  - Active only = active bg, no arrow
  - Selected only = arrow prefix, no bg highlight
- Shortcut hints displayed as tree list:
  ```
  ├── ↑/k Up
  ├── ↓/j Down
  ├── Enter Select
  └── 1-5 Quick nav
  ```
- Nav container: shrinks to content height (not full screen height)

### 3. Status Bar (Footer)
- Full-width gray background
- Left side: 3 badges
  - Badge 1: Area (e.g. "BUILD", "HOME", "LOGS") - colored badge
  - Badge 2: Current selection (e.g. "fistiv") - component name
  - Badge 3 (optional): State (e.g. "RUNNING", "COMPLETE")
- After badges, left side: Nav hints separated by `|`: "Tab Focus | ↑↓/jk Nav | Enter Select | Esc Back"
- Right side: "q Quit" only

### 4. Welcome/Home Screen
- RFZ-CLI ASCII art logo with correct colors (RFZ in red/brand, CLI in cyan)
- "Terminal Orchestration Tool" in white color
- Decorative line below quote: braille block pattern (⣿⣿⣿...) instead of simple line
- Version/branding as 3 separate badges:
  - "v1.0.0" - red bg (brand color)
  - "Deutsche Bahn" - dark grey bg, light grey font
  - "Internal Tool" - light blue bg, dark font
- Navigation shortcut hints in tree view format

### 5. Build Components Screen
- Shortcut hints in top row: must NOT overflow to second row
- Checkbox icons: ○ (unselected) and ◉ (green fill, selected) - component list only
- Component category badges: aligned RIGHT at end of each row
- Select state: blue bg with arrow prefix + dot; Active: filled dot only
- Actions styling:
  - [Build Selected]: green font by default; active = green bg + dark font (light for command badge)
  - Other actions: light font by default; active = blue bg + light font
- Legend below actions: updated to show new ○/◉ icons

### 6. Config Modal
- Shortcut labels: key in blue, description in grey (e.g. "←→" blue, "or h/l to select" grey)
- Bottom navigation hints separated by `|`

### 7. Build Execution View
- Component list: full width
- Column layout: St + Component name left-aligned, rest right-aligned
- Tree icons before status indicator:
  - ├─ for all items, └─ for last item (use charm.land tree patterns where possible)
  - Status icons: ○ pending, spinner running, ✓ success, ✗ failed
- Per-component progress: smaller width, braille blocks (⣿⣿⣿), blue=running, red=error, green=success
- Overall progress: more spacing, use ░░░░ (empty) and ████ (filled) pattern
- No "Running" badge in progress section
- "Pending" badge disappears when all builds finished

### 8. General Layout
- ALL bordered boxes must stay within screen width (currently overflow on right side)
- Proper width calculation accounting for border characters

## Affected Areas & Dependencies
- `internal/ui/components/` - TuiBox, TuiNavigation, TuiNavItem, TuiStatusBar, TuiKeyHints, TuiButton, TuiCheckbox, TuiProgress, TuiTree
- `internal/ui/screens/` - Welcome, Build, BuildExecution, ConfigModal screens
- `internal/ui/styles/` - Color tokens, component styles
- `internal/ui/layout/` - App shell, header, footer, main layout

## Edge Cases & Error Scenarios
- Terminal resize: layouts must recalculate within bounds
- Very long component names: should truncate, not break layout
- Many components (>13): list should scroll, badges stay aligned
- Empty states: should still render correct borders/padding

## Security & Permissions
N/A - Pure UI/visual fixes

## Performance Considerations
- No performance impact expected - these are rendering/styling changes only
- Must maintain smooth 60fps-equivalent rendering for spinners/progress

## Scope Boundaries

**IN SCOPE:**
- Title bar layout and styling
- Navigation sidebar: active/select states, shortcuts alignment, tree hints, content-based height
- Status bar: 3-badge system, nav hints, gray bg, q Quit right-aligned
- Welcome screen: logo colors, subtitle, braille line, badges, tree hints
- Build Components: checkbox icons (○/◉), category alignment, action button styles, legend
- Config Modal: shortcut label colors, hint separators
- Build Execution: tree icons, progress bar styles, component list layout, badge cleanup
- General: fix all border overflow issues
- Component library updates where needed to support these changes

**OUT OF SCOPE:**
- New screens (Logs, Discover, Configuration)
- Functional changes (build logic, Maven integration)
- Config modal checkbox style (keeps [x]/[ ] format)
- New keyboard shortcuts or navigation flows
- Adding new features not in the current implementation

## Mandatory Pre-Implementation Requirement

**CRITICAL FOR EVERY STORY:** Before writing any code, the implementing developer MUST:

1. **READ the prototype reference screenshots** from `references/prototype-screenshots/` relevant to the area being fixed (use the Read tool to view .png files)
2. **READ the current state screenshots** from `references/current/` to understand what currently looks wrong
3. **COMPARE** both visually to fully understand the delta before making changes
4. **Document** which specific screenshots were reviewed in the story's implementation notes

This ensures the developer has the exact visual context needed rather than working solely from text descriptions.

**Relevant prototype screenshots by area:**
- Title Bar / Header: `01-welcome-default.png` (shows correct header)
- Navigation: `01-welcome-default.png`, `02-nav-build-focused.png` through `06-nav-exit-focused.png`
- Status Bar: all `01-*` through `49-*` (footer visible in all)
- Welcome Screen: `01-welcome-default.png`
- Build Components: `10-build-empty-selection.png` through `17-build-action-clear.png`
- Config Modal: `20-config-modal-goals-focus.png` through `34-config-modal-actions-start.png`
- Build Execution: `40-build-execution-starting.png` through `49-build-execution-complete.png`

**Current state screenshots:** `references/current/` directory

## Open Questions
None - all clarifications resolved.

## Proposed User Stories (High Level)

1. **Fix Title Bar / Header Layout** - Red line on top, correct title hierarchy, subtitle + time on same line
2. **Fix Navigation Sidebar Styling** - Active/select states, shortcut alignment, tree hints, content height
3. **Fix Status Bar** - 3-badge system, nav hints with separators, gray bg, q Quit right
4. **Fix Welcome Screen** - Logo colors, subtitle, braille line, version badges, tree hints
5. **Fix Build Components Screen** - Checkbox icons, category alignment, action button styles, select states, legend
6. **Fix Config Modal Styling** - Shortcut label colors, hint separators
7. **Fix Build Execution View** - Tree icons, progress bars, component list layout, badge cleanup
8. **Fix General Border Overflow** - All boxes must stay within terminal width

---
*Review this document carefully. Once approved, detailed user stories will be generated.*
