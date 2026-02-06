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

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/ui/components/navigation.go` → `TuiNavItemRender(item, cursor, active, focused, width)` - Renders single nav item
- `internal/ui/components/navigation.go` → `TuiNavigation(items, cursorIndex, activeIndex, focused, header, footer, width)` - Renders full sidebar navigation
- `internal/ui/components/modal.go` → `TuiModal(config, termWidth, termHeight)` - Renders centered overlay dialog with double border and backdrop

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
