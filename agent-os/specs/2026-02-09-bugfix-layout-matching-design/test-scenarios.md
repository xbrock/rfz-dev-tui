# Test Scenarios - Layout Matching Design

**Spec:** 2026-02-09-bugfix-layout-matching-design
**Datum:** 2026-02-09
**Terminal:** 120x40

## Happy Path

### 1. Welcome Screen
- Launch app at 120x40
- Verify: Red accent line on top of header
- Verify: Title "RFZ-CLI v1.0.0" with subtitle "Terminal Orchestration Tool"
- Verify: ASCII logo with RFZ in brand red, CLI in cyan
- Verify: Braille divider line (not dashes)
- Verify: Three styled badges (version red, org gray, info teal)
- Verify: Tree-style keyboard hints at bottom

### 2. Navigation Sidebar
- Verify: Navigation box with header + divider
- Press j/k: Cursor moves with ">" prefix in cyan
- Verify: Active item has teal background (#164e63)
- Verify: Shortcuts (1, 2, 3, 4, q) right-aligned
- Verify: Tree-style hints in footer

### 3. Status Bar
- Verify: Gray background full-width bar
- Verify: Left side has colored badges (mode + context)
- Verify: Hints separated by pipes " | "
- Verify: "q Quit" hint far right

### 4. Build Components Screen
- Press 1 to navigate to Build
- Tab to focus content
- Verify: Circle symbols (not checkboxes): unfilled = "o", filled = "filled circle"
- Verify: Category badges right-aligned
- Verify: Cursor row has highlight background
- Verify: Legend shows correct symbols

### 5. Build Configuration Modal
- Select components and press Enter
- Verify: Modal overlay with section boxes
- Verify: Section hints use cyan for keys, muted for text
- Verify: Tab cycles between sections

### 6. Build Execution View
- Start build from config modal
- Verify: Tree prefixes on component rows (branch/last)
- Verify: Braille progress bars (not block chars for per-component)
- Verify: Block progress bar (filled/empty) for overall
- Verify: Status badges (Success green, Failed red, Pending gray)
- Verify: No "Running" badge in counters

## Edge Cases

### Terminal Resize
- Resize terminal during build execution
- Verify: Content reflows without border overflow
- Verify: All boxes stay within terminal bounds

### Minimum Terminal Size
- Resize to 80x24
- Verify: App renders correctly (minimum supported)
- Resize to 79x23
- Verify: "Terminal too small" message appears

### Empty Selection
- On build selection, deselect all (press n)
- Verify: Build button shows secondary (dimmed) variant
- Verify: Counter shows "0/X selected"

## Error Scenarios

### N/A
No error scenarios applicable for visual layout fixes. All changes are rendering-only.
