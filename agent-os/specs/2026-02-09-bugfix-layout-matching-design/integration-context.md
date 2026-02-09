# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| LAYOUT-001 | Updated style tokens and header layout | ColorNavActiveBg added, StyleHeader uses BorderTop, viewHeader restructured |
| LAYOUT-008 | Fixed border overflow in all bordered boxes | Width() adjusted for border chars in app.go, selection.go, execution.go, box.go |
| LAYOUT-002 | Fixed nav sidebar styling | navigation.go rewritten, keyhints.go tree mode added, app.go nav height removed |
| LAYOUT-003 | Fixed status bar layout | keyhints.go pipe separator, statusbar.go 3rd badge, app.go status bar badges |
| LAYOUT-004 | Fixed welcome screen layout | welcome.go: white subtitle, braille divider, 3 badges, tree hints |
| LAYOUT-005 | Fixed build components screen | list.go: circle symbols, right-aligned badges, cursor row highlight; selection.go: updated legend |
| LAYOUT-006 | Fixed config modal section hint styling | config.go: sectionHint() helper, cyan keys + muted text in all section hints |
| LAYOUT-007 | Fixed build execution view layout | execution.go: tree icons, braille progress, block overall progress, full-width columns, badge cleanup |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
_None yet_

### Services
<!-- New service classes/modules -->
_None yet_

### Hooks / Utilities
- `build.Model.execBoxWidth()` - returns `m.width - 2` for bordered boxes in build execution view
- `components.TuiKeyHintsTree(hints)` - renders key hints as vertical tree-style list with `├──`/`└──` connectors
- `TuiStatusBarConfig.StateBadge` - optional 3rd badge (e.g., "RUNNING", "COMPLETE") rendered after context badge
- `TuiStatusBarConfig.StateBadgeColor` - color for state badge (defaults to ColorYellow)

### List Functions
- `components.TuiListWidth(items, cursor, mode, focused, counter, width)` - renders list with right-aligned badges within given width
- `components.TuiListItemRenderWidth(item, cursor, mode, focused, width)` - renders single item with right-aligned badge

### Color Tokens
- `components.ColorNavActiveBg` (#164e63) - light blue/teal background for active nav items

### Types / Interfaces
<!-- New type definitions -->
_None yet_

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- Key hints now use `" | "` pipe separator instead of double-space separator
- Status bar supports 3 badges: ModeBadge + ContextBadge + StateBadge (optional)
- Build execution shows: BUILD + component name + RUNNING; Build complete shows: DONE + component name + COMPLETE
- StyleHeader now uses `BorderTop(true)` instead of `BorderBottom(true)` - red accent line is on top
- StyleNavItemActive now uses `ColorNavActiveBg` instead of `ColorSecondary` - dark teal background
- viewHeader() layout changed: title + time/info on same line, subtitle on separate line below
- **CRITICAL pattern:** Lip Gloss `Width(N)` sets content+padding area. Border adds 2 more chars visually. For a box with desired visual width W, use `Width(W - 2)` when border is enabled
- Nav box: `Width(navWidth - 2)` → visual width = navWidth (30)
- Content box: `Width(m.width - navWidth - 2)` → visual width = m.width - navWidth (90)
- Build inner boxes: use `execBoxWidth()` = `m.width - 2` where m.width is the content inner width
- `TuiBoxWithWidth()` now correctly uses `Width(width - 2)` to produce exact visual width
- Nav items: active state always wins over cursor state; cursor adds `›` prefix on active items
- Nav shortcuts are right-aligned using gap calculation: `width - textWidth - shortcutWidth`
- Nav box no longer has fixed height - shrinks to content (no `Height()` on nav box style)
- `viewNavigation()` no longer takes a height parameter
- Welcome screen subtitle uses white (`ColorTextPrimary`) not cyan
- Welcome screen divider uses braille blocks (`⣿`) in muted color instead of single-line divider
- Welcome screen badges: v1.0.0 (brand red bg), Deutsche Bahn (gray bg), Internal Tool (teal bg #164e63)
- Welcome screen hints use `TuiKeyHintsTree()` component for tree-style rendering
- Multi-select lists now use ○/◉ circle symbols instead of ☐/☑ checkboxes, with ◉ in green
- `SymbolCircleUnselected` (○) and `SymbolCircleSelected` (◉) added to styles.go
- List items support right-aligned badges via `TuiListWidth()` / `TuiListItemRenderWidth()` with width parameter
- Cursor+focused rows get `ColorNavActiveBg` background highlight across full width
- Build selection legend uses ◉/○/› symbols instead of [x]/[ ]/>
- Config modal section hints use `sectionHint()` helper: keys in cyan bold, description text in muted gray
- Config modal bottom hints already use `TuiKeyHints()` with pipe separators (from LAYOUT-003)
- Build execution component rows use tree icons: `├─` for non-last, `└─` for last
- Per-component progress uses braille blocks `⣿` colored by status (cyan=running, green=done, red=failed, muted dots=pending)
- Overall progress uses block chars: `█` (filled, cyan) + `░` (empty, muted) with percentage
- Status counters: Running badge removed; Pending badge hidden when count is 0
- Component table columns dynamically sized: name fills remaining width; Phase/Progress/Time right-aligned

---

## File Change Summary

| File | Action | Story |
|------|--------|-------|
| internal/ui/components/styles.go | Modified | LAYOUT-001 |
| internal/app/app.go | Modified | LAYOUT-001, LAYOUT-008 |
| internal/ui/components/box.go | Modified | LAYOUT-008 |
| internal/ui/screens/build/selection.go | Modified | LAYOUT-008 |
| internal/ui/screens/build/execution.go | Modified | LAYOUT-008 |
| internal/ui/components/navigation.go | Modified | LAYOUT-002 |
| internal/ui/components/keyhints.go | Modified | LAYOUT-002 |
| internal/app/app.go | Modified | LAYOUT-002 |
| internal/ui/components/keyhints.go | Modified | LAYOUT-003 |
| internal/ui/components/statusbar.go | Modified | LAYOUT-003 |
| internal/app/app.go | Modified | LAYOUT-003 |
| internal/ui/screens/welcome/welcome.go | Modified | LAYOUT-004 |
| internal/ui/components/list.go | Modified | LAYOUT-005 |
| internal/ui/components/styles.go | Modified | LAYOUT-005 |
| internal/ui/screens/build/selection.go | Modified | LAYOUT-005 |
| internal/ui/screens/build/config.go | Modified | LAYOUT-006 |
| internal/ui/screens/build/execution.go | Modified | LAYOUT-007 |
