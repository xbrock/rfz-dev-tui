# Bug Specification: Build Complete Screen - Broken Navigation, Missing Actions, and Design Mismatch

**Type**: Bug Fix
**Created**: 2026-02-07
**Severity**: Critical
**Priority**: Urgent
**Status**: Open

## Problem Statement

The build completion screen in the RFZ Developer CLI has four interrelated issues that prevent users from navigating away, accessing key post-build actions, and seeing a visually consistent interface. After a build completes, the Tab key is trapped (cannot switch focus back to the Navigation sidebar), the Actions box is missing "View Logs" and "Rebuild Failed" buttons, the "New Build" action label appears twice (as a button and as a hint), and the overall screen layout does not match the approved reference prototype screenshots.

## Environment
- Platform: macOS / Terminal
- Version: Current development (rfz-cli)
- Context: Development
- Terminal: 120x40 canonical size

## Reproduction Steps

1. Start the RFZ CLI application (`go run ./cmd/rfz/main.go`)
2. Navigate to Build Components (press `1` or select from nav)
3. Select one or more components (Space to toggle, Enter to confirm)
4. Configure build options and start the build
5. Wait for the build to complete (all components reach Done/Failed state)
6. Observe the Actions box at the bottom of the build execution view
7. Press `Tab` -- focus does NOT move to the Navigation sidebar
8. Note the absence of "View Logs" and "Rebuild Failed" buttons
9. Note the duplicate "New Build" label (button + hint)
10. Compare overall screen layout with reference screenshots `45-build-execution-actions-focus.png` and `49-build-execution-complete.png`

## Expected Behavior

### Tab Navigation (Issue 1)
- After build completes (phase `phaseCompleted`), pressing `Tab` switches focus between the Navigation sidebar and the Build content area, exactly as it works during the selection phase.
- The app-level `handleKey` method must handle `Tab` for the completed build state rather than delegating it to the build screen's `handleCompletedKey`.

### Actions Box Buttons (Issue 2)
- The Actions box on the build complete screen shows three buttons on the left side:
  - `[View Logs] (l)` -- placeholder button; pressing `l` is a no-op for now (actual log modal is a separate feature)
  - `[Rebuild Failed] (r)` -- immediately re-starts the build for only the components that have `PhaseFailed` status, using the same build configuration
  - `[Back] (Esc)` -- returns to the component selection phase
- The right side of the Actions box shows a single hint: `Tab Switch Focus`

### Duplicate Label Removal (Issue 3)
- The `[New Build] n` button and the `n New Build` hint must both be removed.
- They are replaced by the three buttons and single hint described above.

### Design Match (Issue 4)
- The build complete screen layout must match the reference prototype screenshots:
  - `references/prototype-screenshots/45-build-execution-actions-focus.png` (running state with Actions focused)
  - `references/prototype-screenshots/49-build-execution-complete.png` (completed state)
- The reference shows:
  - A "Build Execution" box containing the Maven command preview
  - A "Components" box with a table header (`St`, `Component`, `Phase`, `Progress`, `Time`) and per-component rows showing status icon, name, phase label, progress bar, and elapsed time
  - A "Progress" box with an overall progress bar labeled `Overall:` with percentage, plus status counter badges (`Running: N`, `Success: N`, `Failed: N`, `Pending: N`)
  - An "Actions" box with button-style actions and hint text
- Current implementation deviates in layout structure, spacing, component styling (status counters are outside the Progress box, progress label text differs, etc.)

## Actual Behavior

### Tab Navigation (Issue 1)
- Line 173 of `internal/app/app.go`: the condition `m.build.IsCompleted()` causes the app to delegate ALL key input (including Tab) to `build.handleCompletedKey`.
- `handleCompletedKey` (line 313 of `internal/ui/screens/build/update.go`) only handles `up/k`, `down/j`, and `n`. It does not handle `tab`, so the key press is silently consumed and ignored.
- Result: user is trapped in the build content area with no way to switch focus via Tab.

### Actions Box (Issue 2)
- `viewExecutionActions()` (line 213 of `internal/ui/screens/build/execution.go`) renders only a single `[New Build]` button when `buildDone()` or `buildCanceled` is true.
- No "View Logs", "Rebuild Failed", or "Back" buttons exist.

### Duplicate Label (Issue 3)
- The function creates a `[New Build]` button (line 218) AND adds `{Key: "n", Label: "New Build"}` to the hints array (line 221).
- Both render on the same line in the Actions box, showing the label twice.

### Design Mismatch (Issue 4)
- The Components section uses a custom table with direct `lipgloss.JoinHorizontal` instead of matching the reference's bordered box with proper header styling.
- Status counters (`viewStatusCounters`) render outside the Progress box rather than inside it as shown in the reference.
- The Progress section label reads "Overall Progress" instead of "Overall:" as in the reference.
- The Actions box border uses `ColorBorder` (dim) instead of `ColorCyan` (bright) when the build content is focused.

## Root Cause Analysis

### Issue 1: Tab Key Trapped
**File**: `internal/app/app.go`, line 173
**Cause**: The condition `m.build.IsCompleted()` is included in the early-return block that delegates all keys to the build screen. The build screen's `handleCompletedKey` does not handle `tab`.
**Fix**: Remove `m.build.IsCompleted()` from the early-return condition so that Tab is handled by the app-level key handler. Alternatively, have `handleCompletedKey` return a message that the app intercepts for Tab.

### Issue 2: Missing Buttons
**File**: `internal/ui/screens/build/execution.go`, lines 217-223
**Cause**: The completed-state branch of `viewExecutionActions` only creates a single "New Build" button. No "View Logs", "Rebuild Failed", or "Back" buttons are defined.
**Fix**: Add the three buttons with their respective key bindings. "Rebuild Failed" requires a new handler in `handleCompletedKey` that filters `buildStates` to failed-only components and re-triggers `StartBuildMsg`.

### Issue 3: Duplicate Label
**File**: `internal/ui/screens/build/execution.go`, lines 218-223
**Cause**: The "New Build" label appears both as a `TuiButton` and as a `KeyHint` in the same action row.
**Fix**: Remove the button and hint for "New Build" entirely; replace with the new button set and a single "Tab Switch Focus" hint on the right.

### Issue 4: Design Mismatch
**Files**: `internal/ui/screens/build/execution.go` (entire `viewExecution`, `viewComponentTable`, `viewOverallProgress`, `viewStatusCounters`, `viewExecutionActions`)
**Cause**: The rendering functions were built incrementally without strict adherence to the reference prototype screenshots. Layout structure, spacing, box nesting, and label text diverge from the approved design.
**Fix**: Restructure the view functions to match the reference screenshots exactly, including proper box nesting for Components, Progress (with counters inside), and Actions sections.

## Impact Assessment
- **Users Affected**: All users who complete or cancel a build
- **Functionality Affected**: Post-build navigation, post-build actions (view logs, rebuild failed), visual design consistency
- **Business Impact**: Users cannot navigate away from the build complete screen via Tab, cannot rebuild failed components, and see an inconsistent interface that does not match the approved design

## Affected Files

| File | Changes Required |
|------|-----------------|
| `internal/app/app.go` | Fix Tab delegation for `phaseCompleted` state |
| `internal/ui/screens/build/update.go` | Add key handlers for `l`, `r`, `Esc` in `handleCompletedKey`; implement rebuild-failed logic |
| `internal/ui/screens/build/execution.go` | Rewrite `viewExecutionActions` with correct buttons/hints; restructure `viewExecution`, `viewComponentTable`, `viewOverallProgress`, `viewStatusCounters` to match reference design |
| `internal/ui/screens/build/model.go` | Add any needed state fields (e.g., action button focus index for completed state) |

## Acceptance Criteria

- [ ] After build completes, pressing `Tab` switches focus between Navigation sidebar and Build content area
- [ ] Actions box shows `[View Logs] (l)`, `[Rebuild Failed] (r)`, `[Back] (Esc)` buttons on the left
- [ ] Actions box shows `Tab Switch Focus` hint on the right
- [ ] Pressing `l` on build complete screen is accepted as a key binding (no-op for now; log modal is a separate feature)
- [ ] Pressing `r` on build complete screen re-starts the build for only failed components with the same configuration
- [ ] Pressing `Esc` on build complete screen returns to the component selection phase
- [ ] No duplicate action labels appear in the Actions box
- [ ] The "New Build" button and hint are removed
- [ ] Components section renders inside a bordered box with proper header matching the reference screenshots
- [ ] Progress section renders inside a bordered box containing both the progress bar (labeled "Overall:") and the status counter badges
- [ ] Actions box border highlights with `ColorCyan` when build content is focused
- [ ] Overall screen layout matches `references/prototype-screenshots/49-build-execution-complete.png`
- [ ] Running state matches `references/prototype-screenshots/45-build-execution-actions-focus.png`
- [ ] All existing tests continue to pass
- [ ] New regression tests cover the build complete screen state

## Reference Screenshots

- **Running with Actions focused**: `references/prototype-screenshots/45-build-execution-actions-focus.png`
- **Build complete**: `references/prototype-screenshots/49-build-execution-complete.png`
- **Build starting**: `references/prototype-screenshots/40-build-execution-starting.png`
- **Build running**: `references/prototype-screenshots/41-build-execution-running.png`
- **Build progress**: `references/prototype-screenshots/44-build-execution-progress.png`
