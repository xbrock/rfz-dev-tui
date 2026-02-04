# Technical Specification: Component Gallery Rendering Issues

> Spec ID: BUGFIX-GALLERY-RENDER
> Technical Analysis Date: 2026-02-03
> Analyzed By: tech-architecture Agent

---

## Root Cause Analysis

### Verified Root Cause: String Concatenation of Multi-line Components

The bug hypothesis has been **CONFIRMED** by code analysis. The issue is in `/Users/lix/xapps/rfz-tui/internal/ui/components/demo/gallery.go`.

**Current Implementation (Lines 106-114, 148-154):**
```go
// TuiBox section - string concatenation
sb.WriteString(components.TuiBox("Single Border", components.BoxSingle, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Double Border", components.BoxDouble, false))
// ... etc

// TuiButton section - string concatenation
sb.WriteString(components.TuiButton("Primary", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Secondary", components.ButtonSecondary, "", false))
// ... etc
```

**Why This Fails:**

1. **TuiBox renders as 3 lines** (confirmed in golden file `TestTuiBox_Single.golden`):
   ```
   Line 1: +---------+
   Line 2: | Content |
   Line 3: +---------+
   ```

2. **TuiButton Secondary renders as 3 lines** (confirmed in golden file `TestTuiButton_Secondary.golden`):
   ```
   Line 1: +------------+
   Line 2: |  Cancel    |
   Line 3: +------------+
   ```

3. **TuiButton Primary renders as 1 line** (confirmed in golden file `TestTuiButton_Primary.golden`):
   ```
     Build
   ```

4. **String concatenation with WriteString appends sequentially**, producing:
   ```
   +---------+
   | Content |
   +---------+  +----------+    <- "  " separator appears here
   | Content2 |
   +----------+                 <- Second box continues on new lines
   ```

   Instead of the expected side-by-side layout:
   ```
   +---------+  +----------+
   | Content |  | Content2 |
   +---------+  +----------+
   ```

### Evidence from Screenshot

The screenshot `/Users/lix/xapps/rfz-tui/Screenshot 2026-02-02 at 23.55.31.png` clearly shows:

1. **TuiBox Section**: Boxes appear stacked with disconnected right borders - the right border characters (like `+` and `|`) appear floating after a space gap
2. **TuiButton Section**: Buttons appear in a stair-step pattern - Primary on first line, Secondary (bordered) drops down, Destructive drops further

---

## Related Code Areas

### Primary File (Requires Modification)

| File | Lines | Function | Issue |
|------|-------|----------|-------|
| `internal/ui/components/demo/gallery.go` | 99-122 | `renderBoxSection()` | String concat for TuiBox |
| `internal/ui/components/demo/gallery.go` | 141-169 | `renderButtonSection()` | String concat for TuiButton |
| `internal/ui/components/demo/gallery.go` | 173-199 | `renderStatusSection()` | String concat for TuiStatus (potential, but single-line so OK) |

### Component Files (No Changes Required)

| File | Purpose | Status |
|------|---------|--------|
| `internal/ui/components/box.go` | TuiBox component | OK - returns multi-line via lipgloss |
| `internal/ui/components/button.go` | TuiButton component | OK - Secondary has border (3 lines) |
| `internal/ui/components/status.go` | TuiStatus component | OK - single line badges |

### Test Files (Require Update After Fix)

| File | Action |
|------|--------|
| `internal/ui/components/demo/gallery_test.go` | Verify tests pass |
| `internal/ui/components/demo/testdata/TestGallery_View_AfterResize.golden` | Regenerate with `-update` flag |

---

## Technical Approach

### Solution: Use `lipgloss.JoinHorizontal()`

The fix is straightforward - replace string concatenation with `lipgloss.JoinHorizontal()` which properly aligns multi-line strings side-by-side.

**Lipgloss is already imported** in gallery.go (line 10):
```go
import (
    "github.com/charmbracelet/lipgloss"
)
```

**JoinHorizontal is used extensively** in the boilerplate code (9 occurrences found), confirming this is the idiomatic approach:
- `boilerplate/internal/ui/modals/confirm/confirm.go`
- `boilerplate/internal/ui/screens/discover/discover.go`
- `boilerplate/internal/app/app.go`
- etc.

### Fix for renderBoxSection() (Lines 106-120)

**BEFORE (broken):**
```go
sb.WriteString("Border Variants:\n")
sb.WriteString(components.TuiBox("Single Border", components.BoxSingle, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Double Border", components.BoxDouble, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Rounded Border", components.BoxRounded, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Heavy Border", components.BoxHeavy, false))
sb.WriteString("\n\n")

sb.WriteString("Focus State:\n")
sb.WriteString(components.TuiBox("Normal", components.BoxSingle, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Focused", components.BoxSingle, true))
```

**AFTER (fixed):**
```go
sb.WriteString("Border Variants:\n")
boxRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiBox("Single Border", components.BoxSingle, false),
    "  ",
    components.TuiBox("Double Border", components.BoxDouble, false),
    "  ",
    components.TuiBox("Rounded Border", components.BoxRounded, false),
    "  ",
    components.TuiBox("Heavy Border", components.BoxHeavy, false),
)
sb.WriteString(boxRow)
sb.WriteString("\n\n")

sb.WriteString("Focus State:\n")
focusRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiBox("Normal", components.BoxSingle, false),
    "  ",
    components.TuiBox("Focused", components.BoxSingle, true),
)
sb.WriteString(focusRow)
```

### Fix for renderButtonSection() (Lines 148-168)

**BEFORE (broken):**
```go
sb.WriteString("Button Variants:\n")
sb.WriteString(components.TuiButton("Primary", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Secondary", components.ButtonSecondary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Destructive", components.ButtonDestructive, "", false))
sb.WriteString("\n\n")

sb.WriteString("With Shortcuts:\n")
sb.WriteString(components.TuiButton("Build", components.ButtonPrimary, "Enter", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Delete", components.ButtonDestructive, "D", false))
sb.WriteString("\n\n")

sb.WriteString("Focus State:\n")
sb.WriteString(components.TuiButton("Normal", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Focused", components.ButtonPrimary, "", true))
```

**AFTER (fixed):**
```go
sb.WriteString("Button Variants:\n")
variantRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Primary", components.ButtonPrimary, "", false),
    "  ",
    components.TuiButton("Secondary", components.ButtonSecondary, "", false),
    "  ",
    components.TuiButton("Destructive", components.ButtonDestructive, "", false),
)
sb.WriteString(variantRow)
sb.WriteString("\n\n")

sb.WriteString("With Shortcuts:\n")
shortcutRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Build", components.ButtonPrimary, "Enter", false),
    "  ",
    components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false),
    "  ",
    components.TuiButton("Delete", components.ButtonDestructive, "D", false),
)
sb.WriteString(shortcutRow)
sb.WriteString("\n\n")

sb.WriteString("Focus State:\n")
focusBtnRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Normal", components.ButtonPrimary, "", false),
    "  ",
    components.TuiButton("Focused", components.ButtonPrimary, "", true),
)
sb.WriteString(focusBtnRow)
```

### Alignment Choice: `lipgloss.Top`

Using `lipgloss.Top` alignment because:
1. Boxes should align at top when they have the same height
2. Single-line buttons (Primary/Destructive) should align with top of bordered buttons (Secondary)
3. This matches the boilerplate usage patterns

Alternative: `lipgloss.Center` could vertically center shorter elements, but `Top` is more predictable.

---

## Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Golden file mismatch after fix | Certain | Low | Regenerate with `go test -update` |
| Alignment issues in edge cases | Low | Low | Visual verification after fix |
| Breaking existing tests | Low | Medium | All tests use Contains checks, not exact matching |
| Performance impact | None | None | JoinHorizontal is O(n) string operation |

### Risk Level: LOW

This is a straightforward fix with minimal risk:
- Only modifies `gallery.go` (demo file, not production component)
- Does not change component APIs (`TuiBox`, `TuiButton` unchanged)
- Uses already-imported lipgloss functionality
- Follows established patterns from boilerplate

---

## Story Refinements

### FIX-001: Fix TuiBox Border Rendering

**Verified Complexity:** S (2 SP) - Confirmed

**WIE (Implementation Steps):**
1. Open `internal/ui/components/demo/gallery.go`
2. Locate `renderBoxSection()` function (line 98)
3. Replace lines 106-114 with `lipgloss.JoinHorizontal()` for Border Variants row
4. Replace lines 116-119 with `lipgloss.JoinHorizontal()` for Focus State row
5. Run `go build ./cmd/rfz/...` to verify compilation
6. Run `./rfz` to visually verify box borders are connected
7. Run `golangci-lint run ./internal/ui/components/demo/gallery.go`

**WO (Affected Files):**
| File | Lines | Change Type |
|------|-------|-------------|
| `internal/ui/components/demo/gallery.go` | 106-114 | MODIFY - Border Variants |
| `internal/ui/components/demo/gallery.go` | 116-119 | MODIFY - Focus State |

### FIX-002: Fix TuiButton Layout Alignment

**Verified Complexity:** S (2 SP) - Confirmed

**WIE (Implementation Steps):**
1. Open `internal/ui/components/demo/gallery.go`
2. Locate `renderButtonSection()` function (line 141)
3. Replace lines 148-154 with `lipgloss.JoinHorizontal()` for Button Variants row
4. Replace lines 156-162 with `lipgloss.JoinHorizontal()` for With Shortcuts row
5. Replace lines 164-167 with `lipgloss.JoinHorizontal()` for Focus State row
6. Run `go build ./cmd/rfz/...` to verify compilation
7. Run `./rfz` to visually verify buttons are horizontally aligned
8. Run `golangci-lint run ./internal/ui/components/demo/gallery.go`

**WO (Affected Files):**
| File | Lines | Change Type |
|------|-------|-------------|
| `internal/ui/components/demo/gallery.go` | 148-154 | MODIFY - Button Variants |
| `internal/ui/components/demo/gallery.go` | 156-162 | MODIFY - With Shortcuts |
| `internal/ui/components/demo/gallery.go` | 164-167 | MODIFY - Focus State |

### FIX-003: Add Regression Tests

**Verified Complexity:** XS (1 SP) - Confirmed

**WIE (Implementation Steps):**
1. After FIX-001 and FIX-002 are complete
2. Run `./rfz` and visually confirm all rendering is correct
3. Run `go test ./internal/ui/components/demo/... -update` to regenerate golden files
4. Run `go test ./internal/ui/components/... -v` to verify all tests pass
5. Run `golangci-lint run ./internal/ui/components/...`

**WO (Affected Files):**
| File | Action |
|------|--------|
| `internal/ui/components/demo/testdata/TestGallery_View_AfterResize.golden` | UPDATE (regenerate) |

---

## Estimation Summary

| Story | Original | Verified | Notes |
|-------|----------|----------|-------|
| FIX-001 | S (2 SP) | S (2 SP) | ~20 LOC changes, straightforward |
| FIX-002 | S (2 SP) | S (2 SP) | ~30 LOC changes, straightforward |
| FIX-003 | XS (1 SP) | XS (1 SP) | Just regenerate golden files |
| **Total** | **5 SP** | **5 SP** | Estimation accurate |

**Implementation Time Estimate:** 30-45 minutes for all three stories

---

## Verification Commands

```bash
# Build verification
go build ./cmd/rfz/...

# Visual verification (must be done manually)
./rfz
# Press q to quit after verifying:
# - TuiBox borders are fully connected
# - TuiButton rows are horizontally aligned

# Run tests (expect golden file mismatch before updating)
go test ./internal/ui/components/demo/... -v

# Update golden files (after visual verification)
go test ./internal/ui/components/demo/... -update

# Verify tests pass with new golden files
go test ./internal/ui/components/... -v

# Lint check
golangci-lint run ./internal/ui/components/...
```

---

## Additional Notes

### renderStatusSection() - No Fix Needed

The `renderStatusSection()` function (lines 173-199) also uses string concatenation, but **no fix is required** because:
- `TuiStatus()` returns single-line badges (no borders)
- `TuiStatusCompact()` returns single-line compact badges
- Single-line components concatenate correctly with `WriteString()`

The screenshot confirms status badges display correctly on a single horizontal line.

### Design Pattern for Future Components

For future gallery additions, use this pattern:
```go
row := lipgloss.JoinHorizontal(
    lipgloss.Top,
    component1,
    "  ",  // 2-space separator
    component2,
    "  ",
    component3,
)
sb.WriteString(row)
```

---

*Technical analysis complete. Ready for implementation.*
