# Bug Specification: Component Gallery Rendering Issues

> Spec ID: BUGFIX-GALLERY-RENDER
> Created: 2026-02-03
> Severity: Critical
> Priority: Urgent
> Status: Ready for Implementation

---

## Problem Statement

The Component Gallery demo screen exhibits two critical visual rendering defects that compromise the integrity of the TUI component library demonstration:

1. **TuiBox Border Disconnection**: Right-side borders appear as disconnected/floating characters instead of properly connected box corners when multiple boxes are rendered side-by-side
2. **TuiButton Vertical Misalignment**: Button variants display in a staggered/cascading pattern instead of horizontal alignment when placed in rows

These issues affect the visual credibility of the component library and indicate underlying rendering logic problems that could propagate to the main application.

---

## Environment

| Attribute | Value |
|-----------|-------|
| Platform | macOS (Darwin 25.2.0) |
| Terminal | Standard terminal (120x40 canonical size) |
| Go Version | 1.21+ |
| Framework | Bubble Tea + Lip Gloss (charm.land stack) |
| Component | Component Gallery Demo (`cmd/rfz/main.go`) |

---

## Bug Details

### Bug 1: TuiBox Border Rendering

**Location:** `internal/ui/components/box.go` / `internal/ui/components/demo/gallery.go`

**Symptom:**
When rendering multiple TuiBox components horizontally side-by-side using string concatenation, the right-side borders appear disconnected from their corners. The border characters float independently rather than forming a cohesive box shape.

**Affected Variants:**
- BoxDouble (Double Border)
- BoxRounded (Rounded Border)
- BoxHeavy (Heavy Border)
- Focus State boxes (Normal/Focused)

**Visual Evidence:**
Screenshot shows boxes like:
```
Double Border   ┐
               ─┤    <- disconnected right border
               ─┘
```

Instead of expected:
```
╔═══════════════╗
║ Double Border ║
╚═══════════════╝
```

**Root Cause Analysis:**
Multi-line rendered components (boxes with top/middle/bottom rows) cannot be horizontally concatenated with simple string operations. Each box renders as multiple lines. When using `sb.WriteString(box1); sb.WriteString("  "); sb.WriteString(box2)`, the lines are not properly aligned - only the first line of each box is concatenated, with subsequent lines appearing on new rows.

**Technical Investigation:**
- `TuiBox()` returns a multi-line string via `lipgloss.Style.Render()`
- `gallery.go` concatenates boxes with: `sb.WriteString(TuiBox(...)); sb.WriteString("  ")`
- This approach fails because multi-line strings need `lipgloss.JoinHorizontal()` for proper side-by-side layout

---

### Bug 2: TuiButton Layout Misalignment

**Location:** `internal/ui/components/button.go` / `internal/ui/components/demo/gallery.go`

**Symptom:**
Button variants appear vertically staggered when rendered in rows. Primary appears first, Secondary drops down, Destructive drops further. This creates a cascading stair-step effect instead of a clean horizontal row.

**Affected Areas:**
- "Button Variants:" row (Primary, Secondary, Destructive)
- "With Shortcuts:" row ([Enter] Build, [Esc] Cancel, [D] Delete)
- "Focus State:" row (Normal, Focused)

**Visual Evidence:**
Screenshot shows:
```
Primary
       Secondary
                  Destructive
```

Instead of expected:
```
Primary   Secondary   Destructive
```

**Root Cause Analysis:**
The `ButtonSecondary` variant includes a border (via `lipgloss.Border(BorderSingle)`), making it a multi-line component (3 lines: top border, content, bottom border). Primary and Destructive are single-line (no border). When concatenated with `WriteString()`, the differing heights cause vertical misalignment.

**Technical Investigation:**
- `TuiButton()` for Secondary returns: top border line + content line + bottom border line
- `TuiButton()` for Primary/Destructive returns: single content line
- String concatenation treats each component as sequential text, not aligned blocks
- Solution requires `lipgloss.JoinHorizontal()` with proper alignment

---

## Reproduction Steps

1. Navigate to project directory: `/Users/lix/xapps/rfz-tui`
2. Build the application: `go build -o rfz ./cmd/rfz`
3. Run the component gallery: `./rfz`
4. Observe the Component Gallery screen
5. Note the TuiBox section showing disconnected borders
6. Note the TuiButton section showing staggered buttons
7. Press `q` to quit

**Expected Result:**
- All boxes display with properly connected borders forming complete rectangles
- All button rows display horizontally aligned at the same baseline

**Actual Result:**
- Box borders appear disconnected with floating right-side characters
- Button rows display in a cascading stair-step pattern

---

## Impact Assessment

| Impact Area | Severity | Description |
|-------------|----------|-------------|
| Visual Quality | Critical | Component gallery fails to demonstrate component quality |
| Developer Trust | High | Broken demo undermines confidence in the component library |
| User Experience | High | Visual artifacts distract and confuse users |
| Component Reuse | Medium | Same rendering pattern may be used elsewhere, propagating bugs |
| Documentation | Medium | Screenshots/demos cannot accurately represent intended design |

---

## Acceptance Criteria

### AC1: TuiBox Border Integrity
- [ ] All TuiBox variants (Single, Double, Rounded, Heavy) render with fully connected borders
- [ ] Focused state boxes display with complete cyan borders
- [ ] Multiple boxes placed side-by-side align properly at the same vertical baseline
- [ ] No floating or disconnected border characters visible

### AC2: TuiButton Horizontal Alignment
- [ ] All button variants in a row align horizontally at the same baseline
- [ ] Secondary buttons (with borders) align with Primary/Destructive buttons (without borders)
- [ ] Buttons with shortcuts align properly with other buttons
- [ ] Focus state buttons align properly with normal buttons

### AC3: Visual Regression Prevention
- [ ] Golden file tests updated to reflect correct rendering
- [ ] Gallery test passes with new correct output
- [ ] All existing component unit tests still pass

### AC4: Code Quality
- [ ] Fix uses Lip Gloss layout functions (JoinHorizontal, Place, etc.) per charm.land first principle
- [ ] No custom string manipulation for layout alignment
- [ ] Code passes linting: `golangci-lint run ./internal/ui/components/...`

---

## Technical Approach

### Recommended Fix

**For gallery.go:**
Replace string concatenation with Lip Gloss horizontal joining:

```go
// BEFORE (broken)
sb.WriteString(components.TuiBox("Double Border", components.BoxDouble, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Rounded Border", components.BoxRounded, false))

// AFTER (fixed)
boxes := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiBox("Double Border", components.BoxDouble, false),
    "  ",
    components.TuiBox("Rounded Border", components.BoxRounded, false),
)
sb.WriteString(boxes)
```

**For button alignment:**
Use `lipgloss.JoinHorizontal()` with `lipgloss.Center` or `lipgloss.Top` alignment:

```go
// BEFORE (broken)
sb.WriteString(components.TuiButton("Primary", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Secondary", components.ButtonSecondary, "", false))

// AFTER (fixed)
buttons := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Primary", components.ButtonPrimary, "", false),
    "  ",
    components.TuiButton("Secondary", components.ButtonSecondary, "", false),
)
sb.WriteString(buttons)
```

---

## Files Affected

| File | Change Type | Description |
|------|-------------|-------------|
| `internal/ui/components/demo/gallery.go` | MODIFY | Replace string concat with lipgloss.JoinHorizontal() |
| `internal/ui/components/demo/gallery_test.go` | MODIFY | Update golden file expectations |
| `testdata/golden/components/gallery/*.golden` | UPDATE | Regenerate golden files with correct output |

---

## Verification Commands

```bash
# Build verification
go build ./cmd/rfz/...

# Run component gallery manually
./rfz

# Run tests
go test ./internal/ui/components/... -v

# Update golden files (after visual verification)
go test ./internal/ui/components/... -update

# Lint check
golangci-lint run ./internal/ui/components/...
```

---

## Related Documentation

- **Design System:** `agent-os/product/design-system.md`
- **Tech Stack:** `agent-os/product/tech-stack.md`
- **Original Spec:** `agent-os/specs/2026-02-02-core-components/spec.md`
- **Screenshot Evidence:** `Screenshot 2026-02-02 at 23.55.31.png`

---

## References

- [Lip Gloss Layout Functions](https://github.com/charmbracelet/lipgloss#joining-paragraphs)
- [Lip Gloss JoinHorizontal](https://pkg.go.dev/github.com/charmbracelet/lipgloss#JoinHorizontal)
- charm.land First Principle (CLAUDE.md)

---

*Detailed stories in: user-stories.md*
