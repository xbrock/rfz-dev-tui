# Test Scenarios: 2026-02-07-bugfix-build-complete-screen

**Generated:** 2026-02-07
**Spec:** agent-os/specs/2026-02-07-bugfix-build-complete-screen

---

## Story 1: Fix Tab Navigation and Key Bindings on Build Complete Screen

### Happy Path

1. **Tab switches focus in completed state**
   - Start a build, wait for completion
   - Press `Tab`
   - **Expected:** Focus moves from content area to navigation sidebar
   - Press `Tab` again
   - **Expected:** Focus returns to content area

2. **View Logs keybinding**
   - On Build Complete screen, press `l`
   - **Expected:** Key is accepted without error (no-op placeholder for future feature)

3. **Rebuild Failed components**
   - Complete a build where at least one component failed
   - Press `r`
   - **Expected:** Only failed components are rebuilt with same config; `StartBuildMsg` sent with filtered component list

4. **Back to selection**
   - On Build Complete screen, press `Esc`
   - **Expected:** Returns to component selection phase (not Welcome screen); build state is reset

5. **Actions box layout**
   - On Build Complete screen, verify Actions box
   - **Expected:** Left side shows `[View Logs] (l)`, `[Rebuild Failed] (r)`, `[Back] (Esc)`; right side shows only `Tab Switch Focus` hint

6. **StatusBar hints**
   - On Build Complete screen, check StatusBar
   - **Expected:** Shows `Tab Switch Focus`, `Navigate`, `Esc Back`

### Edge Cases

1. **Rebuild with no failed components**
   - Complete a build where all components succeed
   - Press `r`
   - **Expected:** No-op, no rebuild triggered

2. **Running state actions**
   - During build execution (not complete), verify Actions box
   - **Expected:** Shows `[View Logs] (l)` and `[Cancel Build] (Esc)` only

3. **Cursor navigation in completed state**
   - On Build Complete screen with multiple components
   - Press `j`/`down` and `k`/`up`
   - **Expected:** Cursor moves between component rows

### Error Scenarios

1. **No duplicate labels**
   - On any build state, check for "New Build" label
   - **Expected:** "New Build" button and hint are completely removed

---

## Story 2: Redesign Build Complete Screen to Match Reference Screenshots

### Happy Path

1. **Four separate boxes layout**
   - Start a build, observe the execution screen
   - **Expected:** Four separate bordered boxes stacked vertically: "Build Execution", "Components", "Progress", "Actions"

2. **Build Execution box**
   - **Expected:** Shows command preview with `$` prompt and green text, wrapped in rounded border

3. **Components box**
   - **Expected:** Table with headers `St`, `Component`, `Phase`, `Progress`, `Time`; wrapped in rounded border with cyan border color; uses `TuiStatusCompact` single characters for status

4. **Focused row highlight**
   - Navigate between component rows
   - **Expected:** Focused row has full-width cyan background highlight with `>` prefix; non-focused rows have `|--` prefix

5. **Progress box**
   - **Expected:** Label reads "Overall:" (not "Overall Progress"); progress bar with percentage; status counter badges (Running, Success, Failed, Pending) as colored pills INSIDE the Progress box

6. **Actions box border color**
   - When content is focused: **Expected:** Actions box border is cyan (`#0891b2`)
   - When content is not focused: **Expected:** Actions box border is default (`#4a4a5e`)

### Edge Cases

1. **Status counter badges styling**
   - During build with mixed statuses
   - **Expected:** Each counter rendered as colored pill badge with background color: Running=cyan, Success=green, Failed=red, Pending=secondary

2. **Consistent spacing**
   - **Expected:** Single blank line between each of the four boxes

### Visual Verification

Compare against reference screenshots:
- `references/prototype-screenshots/40-build-execution-starting.png` - Starting state
- `references/prototype-screenshots/45-build-execution-actions-focus.png` - Actions focused
- `references/prototype-screenshots/49-build-execution-complete.png` - Completed state

---

## Regressions-Checkliste

- [ ] Welcome screen renders correctly
- [ ] Navigation sidebar works on all screens
- [ ] Build selection screen unaffected
- [ ] Build configuration screen unaffected
- [ ] Build execution progress updates in real-time
- [ ] Build cancellation still works during execution
- [ ] Existing golden tests pass (`go test ./internal/app/...`)
- [ ] All lint checks pass (`golangci-lint run`)

---

## Automatisierungs-Hinweise

**Test Commands:**
```bash
# Run all tests
go test ./...

# Run build screen tests specifically
go test ./internal/ui/screens/build/...

# Run app integration tests
go test ./internal/app/...

# Regenerate golden files (if layout changed)
go test ./internal/app/... -update

# Lint check
golangci-lint run
```

**Key Files:**
- `internal/app/app.go` - App-level key handling and StatusBar
- `internal/ui/screens/build/execution.go` - Build screen rendering
- `internal/ui/screens/build/update.go` - Build screen key handlers
- `internal/ui/screens/build/model.go` - Build screen state model
