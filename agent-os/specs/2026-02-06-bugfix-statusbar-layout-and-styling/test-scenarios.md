# Test Scenarios - StatusBar Layout and Styling

**Spec:** 2026-02-06-bugfix-statusbar-layout-and-styling
**Datum:** 2026-02-06

---

## Story-001: Redesign TuiStatusBar with Badge-Based Layout

### Happy Path

1. **Mode Badge Rendering**
   - Given a TuiStatusBarConfig with ModeBadge="BUILD" and ModeBadgeColor=ColorYellow
   - When TuiStatusBar is rendered at width 80
   - Then a colored pill badge with "BUILD" text appears on the left

2. **Context Badge Rendering**
   - Given a TuiStatusBarConfig with ContextBadge="rfz-dispatcher"
   - When TuiStatusBar is rendered
   - Then a secondary-colored badge appears after the mode badge

3. **Full Layout**
   - Given badges, hints, and quit hint configured
   - When TuiStatusBar is rendered at width 120
   - Then layout shows: [ModeBadge] [ContextBadge] ...gap... [Hints] [q Quit]

4. **KeyHints Without Dots**
   - Given multiple KeyHint items
   - When TuiKeyHints is rendered
   - Then items are separated by double-space, no middle-dot character present

5. **Quit Hint Separation**
   - Given QuitHint is set
   - When TuiStatusBar is rendered
   - Then "q Quit" appears separated from other hints on the far right

### Edge Cases

6. **Zero Width**
   - Given Width=0
   - When TuiStatusBar is rendered
   - Then empty string is returned

7. **Mode Badge Only (No Hints)**
   - Given only ModeBadge is set
   - When TuiStatusBar is rendered
   - Then badge appears with gap filling remaining space

8. **Hints Only (No Badges)**
   - Given only Hints are set
   - When TuiStatusBar is rendered
   - Then hints appear right-aligned with gap on left

9. **Quit Hint Only**
   - Given only QuitHint is set
   - When TuiStatusBar is rendered
   - Then quit hint renders correctly

10. **Default Badge Colors**
    - Given ModeBadge without ModeBadgeColor
    - When TuiStatusBar is rendered
    - Then ColorCyan is used as default

11. **Long Badge Truncation**
    - Given long badges + hints at narrow width (80)
    - When TuiStatusBar is rendered
    - Then content fits without overflow, gap is at least 1

### Fehlerfalle

12. **Empty Config**
    - Given only Width is set (no badges, no hints)
    - When TuiStatusBar is rendered
    - Then empty styled bar is returned (no panic)

---

## Story-002: Update Regression Tests

### Happy Path

13. **All Golden Tests Pass**
    - Given golden test files are regenerated
    - When `go test ./internal/ui/components/...` is run
    - Then all 148 tests pass

14. **Badge Color Tests**
    - Given various badge color configurations
    - When rendered and compared to golden files
    - Then output matches expected golden data

15. **FooterItem Unit Test**
    - Given `FooterItem("Enter", "Select")`
    - When rendered
    - Then output matches golden file with styled key + label

16. **FooterItemActive Unit Test**
    - Given `FooterItemActive("Build")`
    - When rendered
    - Then output matches golden file with pill-style badge

### Edge Cases

17. **NoDotSeparator Regression**
    - Given multiple hints rendered via TuiKeyHints
    - When output is inspected
    - Then no middle-dot character exists in the output

---

## Automated Test Coverage

| Test | File | Status |
|------|------|--------|
| TestStatusBar_FullWidth | statusbar_test.go | PASS |
| TestStatusBar_ModeBadgeOnly | statusbar_test.go | PASS |
| TestStatusBar_ContextBadgeOnly | statusbar_test.go | PASS |
| TestStatusBar_RightKeyHints | statusbar_test.go | PASS |
| TestStatusBar_BadgesAndHints | statusbar_test.go | PASS |
| TestStatusBar_KeyHintsIntegration | statusbar_test.go | PASS |
| TestStatusBar_Empty | statusbar_test.go | PASS |
| TestStatusBar_LongBadgeTruncation | statusbar_test.go | PASS |
| TestStatusBar_ZeroWidth | statusbar_test.go | PASS |
| TestStatusBar_BadgeColors | statusbar_test.go | PASS |
| TestStatusBar_ContextBadgeCustomColor | statusbar_test.go | PASS |
| TestStatusBar_QuitHintOnly | statusbar_test.go | PASS |
| TestStatusBar_DefaultModeBadgeColor | statusbar_test.go | PASS |
| TestFooterItem | statusbar_test.go | PASS |
| TestFooterItemActive | statusbar_test.go | PASS |
| TestKeyHints_Single | keyhints_test.go | PASS |
| TestKeyHints_Multiple | keyhints_test.go | PASS |
| TestKeyHints_Empty | keyhints_test.go | PASS |
| TestKeyHints_WidthTruncation | keyhints_test.go | PASS |
| TestKeyHints_ContextAware | keyhints_test.go | PASS |
| TestKeyHints_TwoItems | keyhints_test.go | PASS |
| TestKeyHints_NoDotSeparator | keyhints_test.go | PASS |
