# Spec: Layout Matching Design

> Created: 2026-02-09
> Last Updated: 2026-02-09
> Status: Ready for Execution

## Overview

Fix the RFZ-CLI TUI layout and styling across all implemented screens (title bar, navigation, status bar, welcome, build components, config modal, build execution) to match the approved design prototype from `references/prototype-screenshots/`. Includes fixing border overflow issues affecting all screens.

## User Stories

1. LAYOUT-001: Update Style Tokens and Shared Styles
2. LAYOUT-002: Fix Navigation Sidebar Styling
3. LAYOUT-003: Fix Status Bar Layout
4. LAYOUT-004: Fix Welcome Screen Layout
5. LAYOUT-005: Fix Build Components Screen
6. LAYOUT-006: Fix Config Modal Styling
7. LAYOUT-007: Fix Build Execution View
8. LAYOUT-008: Fix General Border Overflow

## Spec Scope

- Visual/layout fixes only - no functional changes
- All 8 areas from the bug report (title bar, nav, status bar, welcome, build components, config modal, build execution, borders)
- Component library updates where needed
- Charm.land compliance maintained

## Out of Scope

- New screens (Logs, Discover, Configuration)
- Functional changes (build logic, Maven integration)
- New keyboard shortcuts or navigation flows
- Golden file test updates (will need updating but not part of this spec)

## Expected Deliverables

- All screens visually match prototype screenshots
- No border overflow at any terminal width
- Status bar shows 3-badge system
- Navigation uses correct active/select states
- Build execution uses tree icons and braille progress

## Integration Requirements

**Integration Type:** Frontend-only (Go TUI)

**Integration Test Commands:**
```bash
# Build passes
cd /Users/lix/xapps/rfz-tui && go build ./...

# No lint errors
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```

**End-to-End Scenarios:**
1. App launches without visual errors at 120x40
2. Navigation active/select states render correctly
3. Status bar shows badges at bottom
4. Build flow renders matching prototype
