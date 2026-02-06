# Technical Specification - StatusBar Layout and Styling

## Root Cause Analysis

### Suspected Root Cause

The TuiStatusBar was implemented as a basic 3-column text layout (left status | center info | right hints) without the badge/pill styling defined in the design system. Specifically:

1. **`FooterItemActive()` never implemented** — The design-system.md defines this function with `Background(ColorCyan)` pill styling, but it was never added to `styles.go`. The status bar has no concept of badges.

2. **`TuiStatusBarConfig` is too simple** — Only has `Status`, `StatusColor`, `Info`, `Hints`, `Width`. Missing fields for mode badge, context badge, and badge colors per screen.

3. **`TuiKeyHints` uses dot separators** — Uses `" · "` (middle dot) between items, but designs show spaced items without separators. Also doesn't support separating "q Quit" to the far right.

4. **Layout is center-aligned** — Current layout centers the info text and splits space evenly. Designs show left-aligned badges with right-aligned hints.

### Related Code Areas
- `internal/ui/components/statusbar.go` (97 lines): Core rendering — 3-column layout needs full restructure
- `internal/ui/components/keyhints.go` (64 lines): Separator change, quit separation
- `internal/ui/components/styles.go` (lines 367-377): Missing FooterItemActive, FooterItem functions
- `internal/ui/components/demo/layout_gallery.go` (lines 481-497): Consumer of TuiStatusBar
- `agent-os/product/design-system.md` (lines 463-499): Source of truth for badge styling

### Similar Past Issues
- None found — this is the initial implementation diverging from design spec.

## Technical Approach

### Investigation Strategy
1. Compare design-system.md definitions with actual styles.go implementation
2. Map each visual element in `should.png` / `should-2.png` to required code changes
3. Identify all consumers of TuiStatusBar and TuiKeyHints

### Fix Strategy

**Phase 1: Add Missing Style Functions (`styles.go`)**

Add `FooterItemActive()` and `FooterItem()` as defined in design-system.md:

```go
// FooterItem renders a key+label footer item
func FooterItem(key string, label string) string {
    keyStyle := lipgloss.NewStyle().
        Foreground(ColorCyan).
        Bold(true)
    labelStyle := lipgloss.NewStyle().
        Foreground(ColorTextSecondary)
    return keyStyle.Render(key) + " " + labelStyle.Render(label)
}

// FooterItemActive renders an active/current footer badge (pill)
func FooterItemActive(label string) string {
    return lipgloss.NewStyle().
        Background(ColorCyan).
        Foreground(ColorBackground).
        Bold(true).
        Padding(0, 1).
        Render(label)
}
```

**Phase 2: Extend `TuiStatusBarConfig` (`statusbar.go`)**

```go
type TuiStatusBarConfig struct {
    // Badge section (left)
    ModeBadge      string         // Mode text (e.g., "LOGS", "SELECT", "BUILD")
    ModeBadgeColor lipgloss.Color // Background color for mode badge
    ContextBadge   string         // Context text (e.g., "rfz-dispatcher", "Boss")
    ContextExtra   string         // Extra context (e.g., "FOLLOW")
    ContextExtraColor lipgloss.Color // Color for extra context badge

    // Key hints section (right)
    Hints    []KeyHint // Main key hints
    QuitHint *KeyHint  // Separate quit hint (far right)

    // Layout
    Width int
}
```

**Phase 3: Restructure StatusBar Layout (`statusbar.go`)**

New layout: `[ModeBadge] [ContextBadge] [ContextExtra]  ...gap...  [Hints]  [QuitHint]`

- Left: Render ModeBadge as pill using `FooterItemActive`-style with custom color
- Left: Render ContextBadge as secondary badge
- Right: Render hints without dot separators, spaced
- Far right: Render quit hint separated

**Phase 4: Update `TuiKeyHints` (`keyhints.go`)**

- Change separator from `" · "` to `"  "` (double space)
- This keeps the width calculation simple while removing the visual dots

**Phase 5: Update Consumer (`layout_gallery.go`)**

Update `renderStatusBar()` to use new config fields.

### Risk Assessment
- **Complexity**: Medium
- **Side Effect Risk**: Low — StatusBar is a leaf component, only used by layout
- **Areas to Test**: StatusBar rendering, KeyHints rendering, layout_gallery demo

## Story Refinements

### Story 1: Bug Fix
**WIE (Implementation):**
1. Add `FooterItemActive()` and `FooterItem()` to `styles.go`
2. Extend `TuiStatusBarConfig` with badge fields
3. Rewrite `TuiStatusBar()` rendering logic for badge-based layout
4. Update `TuiKeyHints()` separator from `" · "` to `"  "`
5. Update `layout_gallery.go` to use new config

**WO (Affected Files):**
- `internal/ui/components/styles.go` — Add 2 functions (~15 lines)
- `internal/ui/components/statusbar.go` — Restructure config + rendering (~80 lines rewrite)
- `internal/ui/components/keyhints.go` — Change separator (~2 lines)
- `internal/ui/components/demo/layout_gallery.go` — Update renderStatusBar (~15 lines)

### Story 2: Regression Test
**WO (Test Files):**
- `internal/ui/components/statusbar_test.go` — Update 8 tests + add 3 new
- `internal/ui/components/keyhints_test.go` — Update separator tests
- `internal/ui/components/testdata/*.golden` — Regenerate all golden files

## Estimation
- Story 1 (Bug Fix): 5 Story Points
- Story 2 (Regression Test): 3 Story Points
- Total: 8 Story Points
