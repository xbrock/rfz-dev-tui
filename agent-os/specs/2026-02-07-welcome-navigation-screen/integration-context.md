# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| WELCOME-001 | Entry point & demo rename | New app entry point, demo moved to separate cmd |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/app/app.go` → `app.New()` - Creates the top-level application Model (stub - handles q/ctrl+c quit and window resize)

### Services
<!-- New service classes/modules -->
_None yet_

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
_None yet_

### Types / Interfaces
<!-- New type definitions -->
- `internal/app/app.go` → `app.Model` - Top-level Bubble Tea model with width/height fields

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- `cmd/rfz/main.go` is the main application entry point, uses `app.New()` with `tea.WithAltScreen()`
- `cmd/rfz-components-demo/main.go` is the component gallery demo (uses `demo.New()`)
- `cmd/layout-demo/` was removed (merged into rfz-components-demo conceptually)
- `internal/app/app.go` is a minimal stub - WELCOME-002 should extend this with the App Shell layout (header, sidebar, content area)
- Module path: `rfz-cli`

---

## File Change Summary

| File | Action | Story |
|------|--------|-------|
| cmd/rfz/main.go | Modified | WELCOME-001 |
| cmd/rfz-components-demo/main.go | Created | WELCOME-001 |
| cmd/layout-demo/main.go | Deleted | WELCOME-001 |
| internal/app/app.go | Created | WELCOME-001 |
