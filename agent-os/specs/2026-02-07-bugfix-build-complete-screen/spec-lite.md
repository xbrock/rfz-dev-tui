# Bug: Build Complete Screen - Broken Navigation, Missing Actions, and Design Mismatch

**Severity**: Critical | **Priority**: Urgent | **Type**: Frontend (Go TUI / Bubble Tea)

## Problem
After a build completes, Tab key is trapped (cannot switch focus to Navigation), the Actions box only shows a duplicate "New Build" label, "View Logs"/"Rebuild Failed"/"Back" buttons are missing, and the overall screen layout does not match the reference prototype screenshots.

## Root Cause
`app.go` line 173 delegates all keys (including Tab) to the build screen during `phaseCompleted`. The build screen's `handleCompletedKey` does not handle Tab, `l`, `r`, or `Esc`. The `viewExecutionActions` function only renders a single "New Build" button with a duplicate hint. Layout does not match the approved design in `45-build-execution-actions-focus.png` and `49-build-execution-complete.png`.

## Fix Scope
- 3 Stories: Navigation/Keybinding Fix + Layout Redesign + Regression Tests
- Assigned Agent: Frontend-developer
- Files: `app.go`, `update.go`, `execution.go`, `model.go`, new test files

## Acceptance
- Tab switches focus between Navigation and Content in build-complete state
- Actions box shows `[View Logs] (l)`, `[Rebuild Failed] (r)`, `[Back] (Esc)` left, `Tab Switch Focus` right
- Pressing `r` rebuilds only failed components with same config
- No duplicate labels in Actions box
- Screen layout matches reference prototype screenshots
- Regression tests cover all key bindings and layout states
