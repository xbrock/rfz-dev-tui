# Bug: StatusBar Missing Mode Badges, Context Badges, and Styled Key Hints

**Severity**: Critical | **Priority**: Urgent | **Type**: Frontend

## Problem
The TuiStatusBar renders plain text in a 3-column layout instead of the designed badge-based layout with colored mode pills, context badges, and individually styled key hints. The `FooterItemActive()` function from the design system is not implemented. Visual output does not match the approved design mockups.

## Fix Scope
- 2 Stories: Bug Fix + Regression Test
- Assigned Agent: Frontend-developer
- Files: `statusbar.go`, `keyhints.go`, `styles.go`, `layout_gallery.go`, tests

## Acceptance
- Status bar shows colored mode/context badges matching designs
- Key hints rendered as spaced items (no dot separators)
- "q Quit" separated on far right
- All tests updated and passing
