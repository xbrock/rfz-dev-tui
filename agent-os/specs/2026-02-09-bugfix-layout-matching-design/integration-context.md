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

### Color Tokens
- `components.ColorNavActiveBg` (#164e63) - light blue/teal background for active nav items

### Types / Interfaces
<!-- New type definitions -->
_None yet_

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
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
