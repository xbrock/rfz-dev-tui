# Bug Specification: StatusBar Missing Mode Badges, Context Badges, and Styled Key Hints

**Type**: Bug Fix
**Created**: 2026-02-06
**Severity**: Critical
**Priority**: Urgent
**Status**: Open

## Problem Statement

The TuiStatusBar component renders a basic 3-column plain-text layout that fundamentally differs from the design specification. The "should" designs show a badge-based status bar with colored mode pills (e.g., "LOGS", "SELECT"), context badges (e.g., component name with status), and individually styled key hint items. The current implementation shows only plain text with dot-separated hints, missing all badge/pill styling entirely.

## Environment
- Platform: macOS / Terminal
- Version: Current development (Layout Demo)
- Context: Development

## Reproduction Steps
1. Run the application (`go run cmd/rfz/main.go`)
2. Observe the bottom status bar
3. Compare with design mockups (`should.png`, `should-2.png`)

## Expected Behavior

Based on design mockups (`should.png`, `should-2.png`):

1. **Left section** shows colored pill/badge for current mode:
   - "LOGS" badge with blue/cyan background when on Logs screen
   - "SELECT" badge with green background when on Build screen
   - Mode badge changes color per screen context

2. **Left section** also shows a context badge:
   - Active component name (e.g., "rfz-dispatcher") as a secondary badge
   - Optional status indicator (e.g., "FOLLOW" in green) next to context

3. **Right section** shows individually styled key hints:
   - Each hint rendered as a standalone item (e.g., "Tab Focus", "Enter Select")
   - No dot separators between items — spacing only
   - "q Quit" aligned to far right, visually separated

4. **Overall layout**: `[Mode Badge] [Context Badge]` ... `[Key Hints]` ... `[q Quit]`

## Actual Behavior

1. **Left section** shows plain bold text: "Layout Demo" (no badge/pill styling)
2. **Center section** shows plain text screen name: "Navigation"
3. **Right section** shows dot-separated hints: `Tab Section · m Modal · q Quit`
4. No colored badges or pills anywhere
5. No context/selection information shown
6. 3-column layout doesn't match the badge-based design

## Impact Assessment
- Users Affected: All users — status bar is visible on every screen
- Functionality Affected: Visual design fidelity, UX consistency
- Business Impact: Application looks unfinished, doesn't match approved designs

## Error Messages/Logs
No runtime errors — this is a visual/design mismatch.

## Root Cause Summary

1. `TuiStatusBar` uses a simple 3-column text layout instead of badge-based design
2. `FooterItemActive()` function is defined in design-system.md but NOT implemented in code
3. `TuiKeyHints` uses dot separators (`·`) instead of spaced individual items
4. `TuiStatusBarConfig` has no fields for mode badge, context badge, or badge colors
5. No concept of "mode" or "context" in the status bar data model

## Acceptance Criteria
- [ ] Status bar shows colored mode badge (pill) on the left matching current screen
- [ ] Status bar shows context badge with active selection/component info
- [ ] Key hints rendered as individually styled items without dot separators
- [ ] "q Quit" visually separated on the far right
- [ ] Badge colors change based on screen context (LOGS=cyan, SELECT=green, etc.)
- [ ] Visual output matches `should.png` and `should-2.png` designs
- [ ] All existing tests updated to match new rendering
- [ ] No regression in other components
