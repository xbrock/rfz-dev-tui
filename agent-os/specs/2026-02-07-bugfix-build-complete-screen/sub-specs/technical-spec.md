# Technical Specification: Build Complete Screen Bug Fix

**Spec**: 2026-02-07-bugfix-build-complete-screen
**Created**: 2026-02-07
**Type**: Technical Analysis
**Status**: Ready for Implementation

---

## 1. Root Cause Analysis

### Issue 1: Tab Key Trapped in Build Complete State

**Root Cause**: In `internal/app/app.go` line 173, the condition:

```go
if m.screen == screenBuild && (m.build.IsConfiguring() || m.build.IsExecuting() || m.build.IsCompleted()) {
```

delegates ALL key input (including Tab) to the build screen when `IsCompleted()` returns true. The build screen's `handleCompletedKey` (in `internal/ui/screens/build/update.go` lines 312-334) only handles `up/k`, `down/j`, and `n`. Tab is not handled, so it is silently consumed and discarded.

This early-return block is correct for `IsConfiguring()` and `IsExecuting()` because those phases require full key capture (config form navigation, execution-time Esc to cancel). However, `IsCompleted()` does not need full capture -- the user should be able to Tab between nav and content just like in `phaseSelecting`.

**Fix Location**: `internal/app/app.go` line 173

**Fix Strategy**: Remove `m.build.IsCompleted()` from the early-return condition. Instead, handle the completed state similarly to the selection phase: let Tab be processed by the app-level handler (lines 181-189), and delegate only content-focused keys to the build screen via the existing content delegation block (lines 200-210). Additionally, `Esc` in completed state should be handled by the build screen (not the app-level Esc handler which navigates to Welcome), so add a specific completed-state block before the general Esc handler.

---

### Issue 2: Missing Action Buttons (View Logs, Rebuild Failed, Back)

**Root Cause**: In `internal/ui/screens/build/execution.go` lines 217-223, the completed-state branch of `viewExecutionActions` only renders a single "New Build" button:

```go
if m.buildDone() || m.buildCanceled {
    newBuildBtn := components.TuiButton("New Build", components.ButtonPrimary, "n", false)
    buttons = newBuildBtn
    hints = []components.KeyHint{
        {Key: "n", Label: "New Build"},
        {Key: "\u2191\u2193", Label: "Navigate"},
    }
}
```

The reference screenshots (`45-build-execution-actions-focus.png`, `49-build-execution-complete.png`) show the running state should have `[View Logs] (l)` and `[Cancel Build] (Ctrl+C)` buttons, and the completed state should have `[View Logs] (l)`, `[Rebuild Failed] (r)`, and `[Back] (Esc)` buttons with a single `Tab Switch Focus` hint on the right.

No key handlers exist for `l`, `r`, or `Esc` in `handleCompletedKey` (`internal/ui/screens/build/update.go` lines 312-334).

**Fix Locations**:
- `internal/ui/screens/build/execution.go` lines 212-257 (`viewExecutionActions`)
- `internal/ui/screens/build/update.go` lines 312-334 (`handleCompletedKey`)
- `internal/ui/screens/build/update.go` lines 285-310 (`handleExecutionKey`) -- update running-state actions too
- `internal/ui/screens/build/model.go` -- add `failedComponents()` helper

**Fix Strategy**:

1. **viewExecutionActions**: Rewrite the function to render different button sets based on state:
   - **Running state** (`!m.buildDone() && !m.buildCanceled`): `[View Logs] (l)` + `[Cancel Build] (Esc)` (destructive variant), no hints on right
   - **Completed/Canceled state** (`m.buildDone() || m.buildCanceled`): `[View Logs] (l)` + `[Rebuild Failed] (r)` + `[Back] (Esc)`, with `Tab Switch Focus` hint on right

2. **handleCompletedKey**: Replace the `n` handler with:
   - `l`: No-op (placeholder for future log modal)
   - `r`: Filter `buildStates` for `PhaseFailed` components, collect their names, send `StartBuildMsg` with same config and filtered list. If no failed components, no-op.
   - `esc`: Reset to `phaseSelecting`, clear build state (same as current `n` handler logic)
   - Remove the `n` handler entirely

3. **handleExecutionKey**: The running state currently only has `Esc` to cancel. Add `l` as a no-op for consistency with the button shown.

4. **model.go**: Add `failedComponents() []string` helper that iterates `buildStates` and returns names where `Phase == domain.PhaseFailed`.

---

### Issue 3: Duplicate "New Build" Label

**Root Cause**: In `internal/ui/screens/build/execution.go` lines 218-223, the completed branch creates both a `TuiButton("New Build", ...)` on line 218 and a `KeyHint{Key: "n", Label: "New Build"}` on line 221. Both render on the same line, showing "New Build" twice -- once as a bracketed button `[New Build] n` and once as a hint `n New Build`.

**Fix**: This is resolved as part of Issue 2 -- the entire completed-state branch of `viewExecutionActions` is rewritten with the correct buttons and hints. The "New Build" button and hint are both removed.

---

### Issue 4: Design Layout Mismatch

**Root Cause**: The current `viewExecution()` (lines 14-60 of `execution.go`) was built incrementally and does not match the reference prototype. Key deviations identified by comparing the current screenshot (`references/current/Screenshot 2026-02-07 at 20.43.04.png`) against the reference screenshots:

| Aspect | Reference (Prototype) | Current Implementation |
|--------|----------------------|----------------------|
| **Layout Structure** | Four distinct bordered boxes: "Build Execution", "Components", "Progress", "Actions" -- each with its own `BorderRounded` and title | Single outer box "Build Execution" wrapping everything; Components is a raw table without its own border; status counters are outside Progress box |
| **Components Box** | Separate bordered box titled "Components" with table header `St Component ... Phase Progress Time`, dashed divider, component rows with tree-branch characters (`\|--`, `>`) for focused row, full-row cyan background highlight for focused row | No "Components" border box; table columns are `Status Component Phase Progress Time`; focused row only has bold cyan name, not full-row highlight; uses `> ` prefix instead of tree-branch characters |
| **Progress Box** | Bordered box titled "Progress" containing `Overall:` label + progress bar on first line, status counter badges (colored pills like `Running: 3`, `Success: 0`, etc.) on second line -- all inside the box | Label reads "Overall Progress" (not "Overall:"); status counters render below the progress box as separate text, not inside it; counters are plain text with colored numbers, not pill/badge styled |
| **Status Counter Badges** | Rendered as colored pills with icon prefix: `* Running: 3` (orange bg), `v Success: 0` (green bg), `x Failed: 0` (red bg), `o Pending: 0` (grey bg) | Rendered as inline text: `0 Running    2 Success    1 Failed    0 Pending` with colored numbers but no background pills |
| **Actions Box Border** | Border is `ColorCyan` when Actions has focus (screenshot 45) or when content area is focused; `ColorBorder` otherwise | Border always uses `ColorBorder` (line 250 of execution.go) regardless of focus state |
| **Component Table Header** | Columns: `St`, `Component`, `Phase`, `Progress`, `Time` (note: `St` not `Status`) | Columns: `Status`, `Component`, `Phase`, `Progress`, `Time` |
| **Focused Row Highlight** | Full-row cyan background highlight spanning all columns; `>` cursor and tree-branch `\|--` prefix characters | Only the component name gets bold cyan styling; `> ` text prefix for cursor |

**Fix Locations**:
- `internal/ui/screens/build/execution.go` -- all view functions

**Fix Strategy**: Restructure `viewExecution()` to produce four vertically stacked bordered boxes:

1. **"Build Execution" box**: `BorderRounded` with `ColorBorder` foreground. Title via `StyleH3.Render("Build Execution")`. Content: `$ mvn ...` command in green.

2. **"Components" box**: `BorderRounded` with `ColorCyan` foreground (since this is the main interactive area). Title via `StyleH3.Render("Components")`. Content: table with header `St | Component | Phase | Progress | Time`, dashed divider, component rows with tree-branch prefixes and full-row cyan background for focused row.

3. **"Progress" box**: `BorderRounded` with `ColorBorder` foreground. Title via `StyleH3.Render("Progress")`. Content: first line has `Overall:` label + progress bar + percentage; second line has status counter badges rendered as colored pills.

4. **"Actions" box**: `BorderRounded`. Border color is `ColorCyan` when content is focused, `ColorBorder` otherwise. Title via `StyleH3.Render("Actions")`. Content: buttons on left, hint on right.

The overall `viewExecution()` stacks: title ("Build") + all four boxes with single blank lines between them.

---

## 2. Detailed Code Changes

### File 1: `internal/app/app.go`

**Change 1a: Fix Tab delegation (line 173)**

Current code (line 173):
```go
if m.screen == screenBuild && (m.build.IsConfiguring() || m.build.IsExecuting() || m.build.IsCompleted()) {
    var cmd tea.Cmd
    m.build, cmd = m.build.Update(msg)
    return m, cmd
}
```

Change to:
```go
if m.screen == screenBuild && (m.build.IsConfiguring() || m.build.IsExecuting()) {
    var cmd tea.Cmd
    m.build, cmd = m.build.Update(msg)
    return m, cmd
}
```

This removes `m.build.IsCompleted()` so that Tab, number keys, and other app-level keys work normally. The completed state will fall through to the standard Tab handler (line 181) and then to the content delegation block (line 200).

**Change 1b: Handle completed-state Esc before the general Esc handler (after line 189)**

The general `Esc` handler at line 191 navigates to the Welcome screen. For `phaseCompleted`, Esc should go back to component selection instead. Add a check before the general Esc handler:

```go
case "esc":
    // In build completed state, Esc returns to selection (handled by build screen)
    if m.screen == screenBuild && m.build.IsCompleted() {
        var cmd tea.Cmd
        m.build, cmd = m.build.Update(msg)
        return m, cmd
    }
    if m.screen != screenWelcome {
        m.navigateTo(screenWelcome)
        m.build = m.build.SetFocused(false)
    }
    return m, nil
```

**Change 1c: Update StatusBar hints for completed state (lines 507-513)**

Current code:
```go
} else if m.screen == screenBuild && m.build.IsCompleted() {
    modeBadge = "DONE"
    contextBadge = "Build Complete"
    hints = []components.KeyHint{
        {Key: "n", Label: "New Build"},
        {Key: "\u2191\u2193", Label: "Navigate"},
    }
```

Change to:
```go
} else if m.screen == screenBuild && m.build.IsCompleted() {
    modeBadge = "DONE"
    contextBadge = "Build Complete"
    hints = []components.KeyHint{
        {Key: "Tab", Label: "Switch Focus"},
        {Key: "\u2191\u2193", Label: "Navigate"},
        {Key: "Esc", Label: "Back"},
    }
```

---

### File 2: `internal/ui/screens/build/update.go`

**Change 2a: Extend `handleCompletedKey` (lines 312-334)**

Replace the entire function body:

```go
func (m Model) handleCompletedKey(msg tea.KeyMsg) (Model, tea.Cmd) {
    switch msg.String() {
    case "up", "k":
        if m.buildCursor > 0 {
            m.buildCursor--
        }
    case "down", "j":
        if m.buildCursor < len(m.buildStates)-1 {
            m.buildCursor++
        }
    case "l":
        // Placeholder for View Logs modal (no-op)
    case "r":
        // Rebuild failed components only
        failed := m.failedComponents()
        if len(failed) > 0 {
            cfg := m.config
            return m, func() tea.Msg {
                return StartBuildMsg{Config: cfg, Selected: failed}
            }
        }
    case "esc":
        // Back to selection phase
        m.phase = phaseSelecting
        m.buildStates = nil
        m.simStates = nil
        m.buildCursor = 0
        m.buildCanceled = false
        m.items = components.DeselectAll(m.items)
    }
    return m, nil
}
```

**Change 2b: Add `l` no-op to `handleExecutionKey` (lines 285-310)**

Add a case for `l` in the switch:
```go
case "l":
    // Placeholder for View Logs modal (no-op during execution)
```

---

### File 3: `internal/ui/screens/build/model.go`

**Change 3a: Add `failedComponents()` helper**

Add after `buildDone()` method (after line 192):

```go
// failedComponents returns the names of components that have PhaseFailed status.
func (m Model) failedComponents() []string {
    var failed []string
    for _, s := range m.buildStates {
        if s.Phase == domain.PhaseFailed {
            failed = append(failed, s.Name)
        }
    }
    return failed
}
```

---

### File 4: `internal/ui/screens/build/execution.go`

This file requires the most extensive changes. Below is a function-by-function breakdown.

**Change 4a: Rewrite `viewExecution()` (lines 14-60)**

Restructure to render four separate bordered boxes stacked vertically:

```
Build (H2 title)
+-- Build Execution (rounded border) --+
| $ mvn clean install ...              |
+--------------------------------------+

+-- Components (rounded border, cyan) --+
| St  Component  Phase  Progress  Time  |
| ------------------------------------- |
| > * audiocon   Testing |||...  00:10  |
|   * traktion   Testing |||...  00:08  |
+---------------------------------------+

+-- Progress (rounded border) ----------+
| Overall:  [==========......] 67%      |
| * Running: 1  v Success: 2  ...       |
+---------------------------------------+

+-- Actions (rounded border) -----------+
| [View Logs] (l)  [Cancel Build] (Esc) |
+---------------------------------------+
```

Key changes:
- Remove the single outer "Build Execution" box that wraps everything
- Each section (Build Execution, Components, Progress, Actions) gets its own `BorderRounded` box
- Status counters move inside the Progress box
- Overall progress label changes from "Overall Progress" to "Overall:"

**Change 4b: Rewrite `viewComponentTable()` (lines 62-99)**

- Wrap in a bordered box with title "Components"
- Change header column from "Status" to "St"
- Use dashed divider
- Apply `BorderRounded` with `ColorCyan` foreground (the components table is the main interactive area)

**Change 4c: Rewrite `viewComponentRow()` (lines 101-171)**

- Use tree-branch prefix characters (`|--` for non-focused, `>` for focused) matching reference
- Apply full-row cyan background highlight for the focused row (not just the name)
- Ensure all cells in the focused row get the cyan background treatment

**Change 4d: Rewrite `viewOverallProgress()` (lines 173-185)**

- Change label from "Overall Progress" to "Overall:"
- Combine with status counters into a single bordered "Progress" box
- Return the entire box (not just the label + bar line)

**Change 4e: Rewrite `viewStatusCounters()` (lines 187-210)**

- Render each counter as a colored pill/badge with background:
  - Running: orange/cyan dot icon + text on background
  - Success: green check icon + text on background
  - Failed: red X icon + text on background
  - Pending: grey circle icon + text on background
- Format: icon + `Label: N` inside a styled badge with background color

**Change 4f: Rewrite `viewExecutionActions()` (lines 212-257)**

Two states:
- **Running**: `[View Logs] (l)` (secondary) + `[Cancel Build] (Esc)` (destructive)
- **Completed/Canceled**: `[View Logs] (l)` (secondary) + `[Rebuild Failed] (r)` (primary) + `[Back] (Esc)` (secondary). Right side: `Tab Switch Focus` hint
- Border color: `ColorCyan` when content area is focused (use `m.focused` field from model), `ColorBorder` otherwise

---

### File 5: `internal/ui/screens/build/teststate.go`

**Change 5a: Add test state for canceled build**

Add a `TestCanceledState` function for golden tests:

```go
func TestCanceledState(width, height, termW, termH int) Model {
    m := New(width, height)
    m.termW = termW
    m.termH = termH
    m.phase = phaseCompleted
    m.buildCanceled = true
    m.selectedComponents = []string{"boss", "konfiguration", "strecke"}
    m.config = domain.BuildConfig{
        Goal:      domain.GoalCleanInstall,
        Profiles:  []string{"target_env_dev"},
        Port:      11090,
        SkipTests: true,
    }
    m.buildStates = []componentBuildState{
        {Name: "boss", Phase: domain.PhaseDone, ...},
        {Name: "konfiguration", Phase: domain.PhaseFailed, ...},
        {Name: "strecke", Phase: domain.PhaseFailed, ...},
    }
    m.buildCursor = 0
    return m
}
```

---

## 3. Risk Assessment

### Low Risk
- **Issue 1 (Tab fix)**: Removing `m.build.IsCompleted()` from the early-return condition is a targeted, low-risk change. The completed state will use the same delegation path as `phaseSelecting`, which is well-tested.
- **Issue 3 (Duplicate label)**: Resolved automatically by the `viewExecutionActions` rewrite.

### Medium Risk
- **Issue 2 (Missing buttons + Rebuild Failed logic)**: The `r` (rebuild failed) handler sends a new `StartBuildMsg` with a filtered component list. This reuses the existing `handleStartBuild` flow (line 48), which reinitializes `buildStates` and starts the simulator. Risk: if the filtered list is empty (no failed components), a `StartBuildMsg` with empty `Selected` could cause issues. **Mitigation**: The `r` handler checks `len(failed) > 0` before sending the message.
- **Esc in completed state**: The app-level Esc handler must be modified to delegate to the build screen when completed. If this logic is wrong, Esc could stop working for other screens. **Mitigation**: The condition explicitly checks `m.screen == screenBuild && m.build.IsCompleted()` before delegating.

### Medium-High Risk
- **Issue 4 (Design mismatch)**: Restructuring the entire `viewExecution()` function and its sub-functions is the largest change. This affects the visual output for both the executing and completed states. **Mitigation**: Existing golden tests (`TestApp_BuildExecuting`, `TestApp_BuildCompleted` in `internal/app/app_test.go` lines 149-165) will need their golden files regenerated. New golden tests should be added for the Actions box in both states.

### Dependencies Between Issues
- Issues 1, 2, and 3 are interdependent -- they all touch the same key handling and action rendering code
- Issue 4 is semi-independent (visual restructuring) but must incorporate the button changes from Issue 2
- All four issues can be implemented in a single branch since they share affected files

---

## 4. Implementation Strategy

### Recommended Approach: Single Branch, Two Stories

**Story 1** (5 SP) covers Issues 1, 2, and 3 together since they are tightly coupled:
1. Fix `app.go` Tab delegation (Change 1a)
2. Add completed-state Esc delegation in `app.go` (Change 1b)
3. Update StatusBar hints (Change 1c)
4. Add `failedComponents()` helper to `model.go` (Change 3a)
5. Rewrite `handleCompletedKey` in `update.go` (Change 2a)
6. Add `l` no-op to `handleExecutionKey` (Change 2b)
7. Rewrite `viewExecutionActions` in `execution.go` (Change 4f) -- buttons only, not full layout

**Story 2** (5 SP) covers Issue 4:
1. Restructure `viewExecution()` -- four separate bordered boxes (Change 4a)
2. Rewrite `viewComponentTable()` with border and header changes (Change 4b)
3. Rewrite `viewComponentRow()` with full-row highlight (Change 4c)
4. Rewrite `viewOverallProgress()` with "Overall:" label (Change 4d)
5. Rewrite `viewStatusCounters()` with pill/badge styling (Change 4e)
6. Integrate button changes from Story 1 into the new layout

**Story 3** (3 SP) covers regression tests:
1. Unit tests for `handleCompletedKey` (all key bindings)
2. Unit tests for `failedComponents()` helper
3. Integration tests for Tab focus switching in completed state
4. Golden tests for executing and completed view states
5. Regenerate existing golden files

### Implementation Order
1. Story 1 first (functional correctness)
2. Story 2 second (visual match -- can run in parallel on same branch)
3. Story 3 last (tests require final rendered output)

### Verification Checklist
- [ ] `go build ./...` compiles without errors
- [ ] `go test ./internal/app/... -update` regenerates golden files
- [ ] `go test ./internal/ui/screens/build/... -update` generates new golden files
- [ ] `go test ./...` passes after golden file verification
- [ ] `golangci-lint run` passes
- [ ] Manual test: start app, run build, verify Tab works after completion
- [ ] Manual test: press `r` after build with failed components, verify rebuild starts
- [ ] Manual test: press `Esc` after build completion, verify return to selection
- [ ] Visual comparison against reference screenshots at 120x40

---

## 5. Affected Files Summary

| File | Lines Changed | Type of Change |
|------|--------------|----------------|
| `internal/app/app.go` | ~173, ~191, ~507-513 | Remove IsCompleted from early-return; add completed Esc delegation; update StatusBar hints |
| `internal/ui/screens/build/update.go` | ~285-310, ~312-334 | Add `l` to execution handler; rewrite completed handler with `l`, `r`, `esc` |
| `internal/ui/screens/build/model.go` | After ~192 | Add `failedComponents()` helper method |
| `internal/ui/screens/build/execution.go` | ~14-257 (entire file) | Restructure all view functions for design match |
| `internal/ui/screens/build/teststate.go` | After ~53 | Add `TestCanceledState` for golden tests |
| `internal/app/app_test.go` | After ~165 | Add Tab-focus integration tests for completed state |
| `internal/ui/screens/build/update_test.go` | New file | Key handler unit tests |
| `internal/ui/screens/build/execution_test.go` | New file | Golden tests for view rendering |
| `internal/app/testdata/` | Updated files | Regenerated golden files for BuildExecuting, BuildCompleted |

---

## 6. Complexity Assessment

| Story | Points | Rationale |
|-------|--------|-----------|
| Story 1: Tab Nav + Key Bindings + Buttons | 5 | Touches 4 files, requires careful key delegation logic, rebuild-failed needs filtering + message sending |
| Story 2: Design Layout Redesign | 5 | Full rewrite of execution.go view functions, pixel-perfect matching to reference, badge/pill styling |
| Story 3: Regression Tests | 3 | New test files, golden file generation, integration tests at app level |
| **Total** | **13** | |
