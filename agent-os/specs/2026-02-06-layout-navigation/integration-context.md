# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| LAYOUT-001 | TuiNavigation sidebar + TuiNavItem menu items | `navigation.go`, `navigation_test.go`, 12 golden files |
| LAYOUT-002 | TuiModal overlay dialog with double border | `modal.go`, `modal_test.go`, 10 golden files |
| LAYOUT-003 | TuiKeyHints keyboard shortcut display | `keyhints.go`, `keyhints_test.go`, 6 golden files |
| LAYOUT-004 | TuiTable wrapper around bubbles/table | `table.go`, `table_test.go`, 7 golden files |
| LAYOUT-005 | TuiTree hierarchical view with expand/collapse | `tree.go`, `tree_test.go`, `styles.go` updated, 11 golden files |
| LAYOUT-006 | TuiTabs horizontal tab bar with shortcuts/badges | `tabs.go`, `tabs_test.go`, `styles.go` updated, 10 golden files |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/ui/components/navigation.go` → `TuiNavItemRender(item, cursor, active, focused, width)` - Renders single nav item
- `internal/ui/components/navigation.go` → `TuiNavigation(items, cursorIndex, activeIndex, focused, header, footer, width)` - Renders full sidebar navigation
- `internal/ui/components/modal.go` → `TuiModal(config, termWidth, termHeight)` - Renders centered overlay dialog with double border and backdrop
- `internal/ui/components/keyhints.go` → `TuiKeyHints(hints, width)` - Renders horizontal keyboard hints with middle-dot separators
- `internal/ui/components/table.go` → `NewTuiTable(cfg TuiTableConfig) table.Model` - Creates RFZ-styled bubbles/table
- `internal/ui/components/table.go` → `TuiTableStyles() table.Styles` - Returns pre-configured RFZ table styles
- `internal/ui/components/table.go` → `TuiTableEmpty(columns, width)` - Renders empty table with "No data" message
- `internal/ui/components/tree.go` → `TuiTree(nodes, cursorIndex, focused)` - Renders hierarchical tree view
- `internal/ui/components/tree.go` → `TuiTreeItem(node, depth, cursor, focused)` - Renders single tree node line
- `internal/ui/components/tree.go` → `VisibleNodeCount(nodes)` - Returns count of visible nodes for cursor bounds
- `internal/ui/components/tabs.go` → `TuiTabs(tabs, focusedIndex, width)` - Renders horizontal tab bar with numeric shortcuts

### Services
<!-- New service classes/modules -->
_None yet_

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
_None yet_

### Types / Interfaces
<!-- New type definitions -->
- `internal/ui/components/navigation.go` → `TuiNavItem{Label, Number, Shortcut}` - Navigation menu item struct
- `internal/ui/components/modal.go` → `TuiModalConfig{Title, Content, Buttons, Width, Height, FocusedIndex}` - Modal configuration
- `internal/ui/components/modal.go` → `TuiModalButton{Label, Variant, Shortcut}` - Modal footer button
- `internal/ui/components/keyhints.go` → `KeyHint{Key, Label}` - Single keyboard hint struct
- `internal/ui/components/keyhints.go` → `SymbolKeySeparator` - Middle dot separator constant
- `internal/ui/components/table.go` → `TuiTableConfig{Columns, Rows, Width, Height, Focused, ZebraStripe}` - Table configuration struct
- `internal/ui/components/tree.go` → `TuiTreeNode{Label, Metadata, Children, Expanded}` - Tree node struct
- `internal/ui/components/styles.go` → `SymbolExpanded`, `SymbolCollapsed` - Tree expand/collapse symbols
- `internal/ui/components/tabs.go` → `TuiTab{Label, Badge, Active}` - Tab item struct
- `internal/ui/components/tabs.go` → `SymbolTabSeparator` - Pipe separator constant
- `internal/ui/components/styles.go` → `StyleTabActive`, `StyleTabFocused`, `StyleTabNormal` - Tab state styles

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- Navigation uses pure Lipgloss stateless render functions (same pattern as `list.go`)
- Uses existing styles: `StyleNavItem`, `StyleNavItemFocused`, `StyleNavItemActive` from `styles.go`
- Uses `TuiDivider` for header/footer separators
- Uses `Truncate()` helper for long labels (max 30 chars)
- `SymbolListPointer` (`›`) used as cursor indicator
- Modal uses pure Lipgloss stateless render function with config struct pattern
- Modal composes `TuiButton` for footer actions and `TuiDivider` for section separators
- `BorderDouble` + `ColorCyan` border for modal frame
- `lipgloss.Place()` with `ColorBackground` whitespace for backdrop centering
- `FocusedIndex` in config controls which button appears focused (for Tab cycling)
- KeyHints uses pure Lipgloss stateless render function (same pattern as navigation/modal)
- `SymbolKeySeparator` ("·") used between hints with `ColorTextMuted` styling
- Keys rendered in `ColorCyan` + Bold, labels in `ColorTextSecondary`
- Width parameter enables adaptive truncation (omits hints that don't fit)
- TuiTable wraps `bubbles/table` with RFZ styling (different pattern: returns `table.Model` for Bubble Tea lifecycle)
- `NewTuiTable` is a factory function returning `table.Model`, not a stateless render function
- `TuiTableEmpty` is a stateless render for the empty-table case with headers + "No data" message
- Header: bold `ColorTextSecondary` with bottom border, Cell: `ColorTextPrimary`, Selected: bold `ColorSecondary` bg
- TuiTree uses pure Lipgloss stateless render functions (same pattern as navigation/list)
- Flatten-based rendering: tree nodes are flattened to visible list for cursor indexing
- `SymbolExpanded` ("▼") and `SymbolCollapsed` ("▶") for branch nodes, leaf nodes get spacing
- `maxTreeDepth=15` limits nesting; deeper nodes show "..." truncation indicator
- `VisibleNodeCount()` utility for cursor bounds checking in Bubble Tea models
- TuiTabs uses pure Lipgloss stateless render function (same pattern as keyhints)
- `SymbolTabSeparator` ("|") between tabs with `ColorTextMuted` styling
- Active tab: `StyleTabActive` with `ColorSecondary` background + bold
- Focused tab: `StyleTabFocused` with `ColorCyan` + bold + underline (distinct from active)
- Normal tab: shortcut in `ColorTextMuted`, label in `ColorTextSecondary`
- Numeric shortcuts `1:` through `9:` for first 9 tabs; tabs 10+ have no prefix
- Badge format: `Label (N)` in `ColorTextMuted`
- Width parameter omits tabs that don't fit (same adaptive pattern as keyhints)

---

## File Change Summary

| File | Action | Story |
|------|--------|-------|
| `internal/ui/components/navigation.go` | Created | LAYOUT-001 |
| `internal/ui/components/navigation_test.go` | Created | LAYOUT-001 |
| `internal/ui/components/testdata/TestNavItem_*.golden` (6 files) | Created | LAYOUT-001 |
| `internal/ui/components/testdata/TestNavigation_*.golden` (6 files) | Created | LAYOUT-001 |
| `internal/ui/components/modal.go` | Created | LAYOUT-002 |
| `internal/ui/components/modal_test.go` | Created | LAYOUT-002 |
| `internal/ui/components/testdata/TestModal_*.golden` (10 files) | Created | LAYOUT-002 |
| `internal/ui/components/keyhints.go` | Created | LAYOUT-003 |
| `internal/ui/components/keyhints_test.go` | Created | LAYOUT-003 |
| `internal/ui/components/testdata/TestKeyHints_*.golden` (6 files) | Created | LAYOUT-003 |
| `internal/ui/components/table.go` | Created | LAYOUT-004 |
| `internal/ui/components/table_test.go` | Created | LAYOUT-004 |
| `internal/ui/components/testdata/TestTable_*.golden` (7 files) | Created | LAYOUT-004 |
| `internal/ui/components/tree.go` | Created | LAYOUT-005 |
| `internal/ui/components/tree_test.go` | Created | LAYOUT-005 |
| `internal/ui/components/styles.go` | Modified | LAYOUT-005 |
| `internal/ui/components/testdata/TestTree*.golden` (11 files) | Created | LAYOUT-005 |
| `internal/ui/components/tabs.go` | Created | LAYOUT-006 |
| `internal/ui/components/tabs_test.go` | Created | LAYOUT-006 |
| `internal/ui/components/styles.go` | Modified | LAYOUT-006 |
| `internal/ui/components/testdata/TestTabs_*.golden` (10 files) | Created | LAYOUT-006 |
