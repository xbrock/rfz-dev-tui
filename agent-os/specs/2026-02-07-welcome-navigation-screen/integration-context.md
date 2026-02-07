# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| WELCOME-002 | App Shell with full layout | Root Model with header, nav sidebar, content, statusbar, clock tick |
| WELCOME-001 | Entry point & demo rename | New app entry point, demo moved to separate cmd |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/app/app.go` → `app.New()` - Creates the full App Shell Model with header, navigation sidebar, content area, and status bar

### Services
<!-- New service classes/modules -->
_None yet_

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
_None yet_

### Types / Interfaces
<!-- New type definitions -->
- `internal/app/app.go` → `app.Model` - Root Bubble Tea model with width, height, focus, cursorIndex, activeIndex, currentTime
- `internal/app/messages.go` → `app.TickMsg` - Sent every second for clock updates
- `internal/app/messages.go` → `app.NavigateMsg` - Navigation request with Screen index

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- `cmd/rfz/main.go` is the main application entry point, uses `app.New()` with `tea.WithAltScreen()`
- `cmd/rfz-components-demo/main.go` is the component gallery demo (uses `demo.New()`)
- `cmd/layout-demo/` was removed (merged into rfz-components-demo conceptually)
- `internal/app/app.go` is now the full App Shell: header (title+clock), nav sidebar (30 chars fixed), content area, statusbar
- Navigation has 5 items: Build Components, View Logs, Discover, Configuration, Exit (indices 0-4)
- Focus system: `focusNav` / `focusContent` toggled via Tab key
- `activeIndex = -1` means Welcome/Home screen (no nav item selected)
- Content area currently shows screen title only - WELCOME-003 will fill in the Welcome Screen content
- Clock ticks every second via `tea.Every` → `TickMsg`
- Terminal minimum size check: 80x24, shows warning if too small
- Module path: `rfz-cli`

---

## File Change Summary

| File | Action | Story |
|------|--------|-------|
| cmd/rfz/main.go | Modified | WELCOME-001 |
| cmd/rfz-components-demo/main.go | Created | WELCOME-001 |
| cmd/layout-demo/main.go | Deleted | WELCOME-001 |
| internal/app/app.go | Created | WELCOME-001 |
| internal/app/app.go | Modified | WELCOME-002 |
| internal/app/messages.go | Created | WELCOME-002 |
