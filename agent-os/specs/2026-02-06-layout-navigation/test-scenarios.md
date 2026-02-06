# Test Scenarios: Layout Navigation Components

> Spec: 2026-02-06-layout-navigation
> Generated: 2026-02-06
> Components: TuiNavigation, TuiModal, TuiKeyHints, TuiTable, TuiTree, TuiTabs, TuiStatusBar

---

## Automated Tests

### Pre-requisites

```bash
# All components build
go build ./internal/ui/components/...

# All tests pass
go test ./internal/ui/components/... -v

# Demo builds
go build ./cmd/layout-demo/...

# Lint passes
golangci-lint run ./internal/ui/components/...
```

---

## Manual Test Scenarios

### 1. TuiNavigation (LAYOUT-001)

#### Happy Path
1. Run `go run ./cmd/layout-demo/`
2. Verify navigation sidebar is visible on the left
3. Verify all menu items show number + label (e.g., "1 Build")
4. Press `j` / `k` to move cursor up/down
5. Verify cursor indicator (`>`) follows the focused item
6. Verify focused item text is cyan
7. Press number keys `1`-`4` to jump directly to items
8. Press `Enter` to select an item
9. Verify active item has background color highlight

#### Edge Cases
- Navigation with only 1 item renders correctly
- Navigation with 0 items renders empty container
- Pressing `j` at last item does not scroll past end
- Pressing `k` at first item does not scroll past top

#### Error Cases
- Invalid item index does not cause panic
- Rapid key presses are handled correctly

---

### 2. TuiModal (LAYOUT-002)

#### Happy Path
1. Press designated key to open a modal
2. Verify modal appears centered on screen
3. Verify double border (`═══`) is visible
4. Verify title bar shows modal title
5. Verify backdrop dims the background content
6. Press `Tab` to cycle through modal buttons
7. Press `Enter` to activate focused button
8. Press `Escape` to close modal

#### Edge Cases
- Modal with long title truncates gracefully
- Modal with no buttons renders correctly
- Modal wider than terminal clips properly
- Multiple rapid open/close cycles work

#### Error Cases
- Closing an already-closed modal is a no-op
- Modal content overflow is handled

---

### 3. TuiKeyHints (LAYOUT-003)

#### Happy Path
1. Verify key hints display at bottom of component
2. Verify format: key in cyan + label in normal text
3. Verify multiple hints separated by `*` separator
4. Example: `Enter Select * Esc Cancel * q Quit`

#### Edge Cases
- Single hint renders without separator
- Empty hints list renders nothing
- Very long hint list wraps or truncates to fit width
- Wide terminal shows all hints on one line

#### Error Cases
- Hints with empty key or label are skipped

---

### 4. TuiTable (LAYOUT-004)

#### Happy Path
1. Verify table shows column headers (bold/highlighted)
2. Verify data rows display correctly under headers
3. Press `j`/`k` to navigate between rows
4. Verify selected row is highlighted
5. Press `Enter` to select a row
6. Verify zebra striping on alternating rows (if enabled)

#### Edge Cases
- Table with 0 rows shows headers only
- Table with 1 row renders correctly
- Very long cell content truncates with ellipsis
- Narrow terminal adjusts column widths

#### Error Cases
- Missing column data does not panic
- Navigating in empty table is a no-op

---

### 5. TuiTree (LAYOUT-005)

#### Happy Path
1. Verify tree shows hierarchical structure with indentation
2. Verify collapsed nodes show `>` icon
3. Press `Enter` on a collapsed parent node
4. Verify node expands and shows `v` icon
5. Verify children become visible with increased indent
6. Press `Enter` again to collapse
7. Navigate with `j`/`k` through visible nodes

#### Edge Cases
- Leaf nodes (no children) show no expand icon
- Deeply nested tree (5+ levels) indents correctly
- Tree with single root node renders correctly
- Collapsing parent hides all nested descendants

#### Error Cases
- Toggling a leaf node is a no-op
- Rapid expand/collapse does not corrupt state

---

### 6. TuiTabs (LAYOUT-006)

#### Happy Path
1. Verify tabs display horizontally
2. Verify each tab shows number + label (e.g., `1 Build`)
3. Verify active tab is highlighted with underline/color
4. Press number keys `1`-`4` to switch tabs
5. Press `Left`/`Right` arrows to navigate between tabs
6. Verify content area updates when tab changes

#### Edge Cases
- Single tab renders correctly
- Tab with very long label truncates
- Switching to already-active tab is a no-op
- More tabs than fit screen width handled gracefully

#### Error Cases
- Invalid tab number key press is ignored
- Switching tabs with no content area does not panic

---

### 7. TuiStatusBar (LAYOUT-007)

#### Happy Path
1. Verify status bar spans full terminal width
2. Verify left section shows status text
3. Verify center section shows info text
4. Verify right section shows key hints (via TuiKeyHints)
5. Verify background color fills entire bar

#### Edge Cases
- Status bar in narrow terminal truncates center section first
- Empty sections still maintain layout
- Very long status text truncates with ellipsis
- Width exactly matches terminal width (no overflow/underflow)

#### Error Cases
- Bar with no content renders as colored empty bar
- Resizing terminal updates bar width

---

### 8. Layout Demo (LAYOUT-008)

#### Happy Path
1. Run `go run ./cmd/layout-demo/`
2. Verify all 7 components are visible in the gallery
3. Verify each component section has a label/title
4. Verify interactive components respond to keyboard input
5. Press `q` to quit the demo cleanly

#### Edge Cases
- Demo runs in 80x24 terminal (minimum)
- Demo runs in 120x40 terminal (canonical)
- Demo runs in very wide terminal (200+ columns)

#### Error Cases
- Demo exits cleanly on `Ctrl+C`
- Demo handles terminal resize during runtime

---

## Visual Regression Tests (LAYOUT-009)

All visual states are covered by golden file tests:

```bash
go test ./internal/ui/components/... -v -run TestGolden
```

**Coverage:** 162 golden files covering all component states across focused, unfocused, active, and various size configurations.

---

## Integration Validation (LAYOUT-998)

Previously validated:

| Check | Result |
|-------|--------|
| `go build ./internal/ui/components/...` | PASSED |
| `go test ./internal/ui/components/... -v` | PASSED (189 tests) |
| `go build ./cmd/layout-demo/...` | PASSED |
| `golangci-lint run ./internal/ui/components/...` | PASSED (0 issues) |
