# UX Patterns: RFZ Developer CLI

> Last Updated: 2026-02-02
> Version: 1.0.0
> Platform: Desktop TUI (Terminal User Interface)

## Overview

User experience patterns for the RFZ Developer CLI. This document defines navigation, interaction, feedback, and accessibility patterns for a consistent, keyboard-driven terminal user experience built on the charm.land stack (Bubble Tea, Lip Gloss, Bubbles).

**Target Environment:**
- Terminal size: 120 columns x 40 rows (canonical)
- Input: Keyboard-only (no mouse required)
- Users: RFZ developers comfortable with terminal interfaces
- Framework: Bubble Tea (Elm architecture)

---

## Navigation Patterns

### Primary Navigation

**Type:** Sidebar Navigation with Numbered Shortcuts

**Structure:**
```
+-- Navigation --+
| > 1. Build Components  1 |  <- Active item (> prefix, cyan highlight)
|   2. View Logs         2 |
|   3. Discover          3 |
|   4. Configuration     4 |
|   5. Exit              q |
+-------------------------+
|   shortcuts help        |
+-------------------------+
```

**Behavior:**
- Active state: `>` prefix + cyan text highlight + number badge
- Focused state: Cyan border around navigation panel
- Number shortcuts: Press `1-5` or `q` from anywhere for quick navigation
- Keyboard: `j/k` or arrow keys to navigate within panel
- Enter: Select focused navigation item

**Implementation (Lip Gloss):**
```go
// Active nav item style
activeStyle := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FFFF")).  // Cyan
    Bold(true)

// Focused panel style
focusedBorder := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("#00FFFF"))
```

### Secondary Navigation

**Type:** Tab Navigation (within Configuration screen)

**Usage:** Switch between configuration sections (Scan Paths, Registry, Detected)

**Behavior:**
- `Tab` key switches between sections
- Active tab: Underlined or highlighted
- Each section maintains its own scroll position

### Panel Focus Model

**Pattern:** Focus cycles between panels using Tab key

**Focus Order:**
1. Navigation Panel (sidebar)
2. Content Panel (main area)
3. Actions Panel (if present)

**Visual Indicators:**
- Focused panel: Cyan border
- Unfocused panel: Gray/dim border
- Current focus visible at all times

**Keyboard:**
- `Tab`: Move to next panel
- `Shift+Tab`: Move to previous panel (if supported)
- `Esc`: Return focus to navigation OR close modal

---

## User Flow Patterns

### Build Components Flow

**Description:** User selects components and configures a Maven build

**Steps:**
1. Navigate to Build screen (press `1` or select from nav)
2. Use `j/k` or arrows to navigate component list
3. Press `Space` to toggle component selection (`[x]` / `[ ]`)
4. Press `a` to select all, `n` to clear selection
5. Press `Enter` to open Build Configuration modal
6. Configure build options in modal
7. Press `Enter` on "Start Build" to begin execution
8. Monitor progress in Build Execution view

**Pattern:** Multi-select List -> Modal Configuration -> Execution View

**Component Selection:**
```
> [x] boss                                     Core
  [ ] fistiv                                   Core
  [x] audiocon                              Standalone
  [ ] traktion                                 Core
```
- `>` prefix: Current cursor position
- `[x]`: Selected
- `[ ]`: Not selected
- Right-aligned badges: Component category (Core, Simulation, Standalone)

### Log Viewing Flow

**Description:** User views build logs with filtering

**Steps:**
1. Navigate to Log Viewer (press `2`) OR press `L` during build execution
2. Select component from left panel
3. Logs load in viewport (right panel)
4. Use `j/k` or arrows to scroll
5. Press `f` to toggle follow mode (auto-scroll)
6. Press `e` to toggle error-only filter
7. Press `/` to search within logs

**Pattern:** List Selection -> Scrollable Viewport with Filters

### Discover Flow

**Description:** User scans repositories and views component status

**Steps:**
1. Navigate to Discover screen (press `3`)
2. View auto-scanned components with Git status
3. Use `j/k` to navigate component list
4. View details: branch, last commit, clean/dirty status
5. Press `Tab` to access Git actions panel

**Pattern:** Auto-populated List -> Detail View -> Action Buttons

### Configuration Flow

**Description:** User manages scan paths and views component registry

**Steps:**
1. Navigate to Configuration (press `4`)
2. Use `Tab` to switch between sections (Scan Paths, Registry, Detected)
3. In Scan Paths: Add/remove/edit paths
4. Registry and Detected: Read-only views
5. Press `Esc` to return to navigation

**Pattern:** Tabbed Sections with CRUD on editable section

---

## Interaction Patterns

### Buttons & Actions

**Primary Action (Cyan/Teal):**
- Style: `[Action Name]` with cyan text, keyboard shortcut in parentheses
- Placement: Right side of actions panel, or bottom of modal
- Examples: `[Build Selected] (Enter)`, `[Start Build] (Enter)`

**Secondary Action (Gray):**
- Style: `[Action Name]` with gray/dim text
- Placement: Left of primary action
- Examples: `[Select All] (a)`, `[Clear Selection] (n)`, `[Cancel] (Esc)`

**Destructive Action (Red):**
- Style: `[Action Name]` with red text
- Confirmation: Required for destructive actions
- Examples: `[Cancel Build] (Ctrl+C)`

**Action Display Pattern:**
```
+-- Actions -------------------------------------------+
| [Build Selected] (Enter)  [Select All] (a)  [Clear Selection] (n)  Tab Switch focus |
+-----------------------------------------------------+
```

**Implementation (Lip Gloss):**
```go
primaryBtn := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00CED1"))  // Cyan

secondaryBtn := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#808080"))  // Gray

destructiveBtn := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FF6B6B"))  // Red
```

### Selection Patterns

**Multi-Select (Checkboxes):**
```
[x] Selected item
[ ] Unselected item
```
- `Space`: Toggle selection on focused item
- `a`: Select all
- `n`: Clear all (select none)
- Show count: `(3/13 selected)`

**Single-Select (Radio Buttons):**
```
>(o) Port 11090   ( ) Port 11091
```
- `Space` or `Enter`: Select option
- `h/l` or arrow keys: Move between options
- Only one can be selected

**Focus Indicator:**
- `>` prefix on current item
- Cyan highlight on focused row
- Background color change (subtle)

### Forms & Input

**Text Input (using bubbles/textinput):**
- Pattern: Inline text input with cursor
- Placeholder: Dim text when empty
- Focus: Cursor visible, border highlight

**Radio Groups:**
- Horizontal layout for 2-3 options
- Vertical layout for 4+ options
- Selected: `(o)` filled circle
- Unselected: `( )` empty circle

**Checkboxes:**
- Pattern: `[x]` selected, `[ ]` unselected
- Toggle: `Space` or `Enter`

**Validation:**
- Pattern: Inline below field
- Real-time: Validate as user types (debounced)
- Error color: Red text

### Context Menus & Dropdowns

**Not used in TUI.** Actions are displayed in dedicated Actions panels with keyboard shortcuts.

### Modal Pattern

**Trigger:** `Enter` on selection (e.g., Build Configuration modal)

**Structure:**
```
+========== Modal Title =========================================+
||                                                              ||
||  Modal content with sections                                 ||
||                                                              ||
||  +-- Section 1 --+                                           ||
||  | Content       |                                           ||
||  +---------------+                                           ||
||                                                              ||
||  +-- Section 2 --+                                           ||
||  | Content       |                                           ||
||  +---------------+                                           ||
||                                                              ||
||                        [Cancel] (Esc)    > [Confirm] (Enter) ||
+===============================================================+
```

**Behavior:**
- Double border: Visual distinction from background
- Focus trap: Tab cycles within modal only
- Esc: Close modal, return focus to trigger
- Enter: Confirm action (on focused button)
- Background: Dimmed/overlay effect (if possible)

**Implementation (Lip Gloss):**
```go
modalStyle := lipgloss.NewStyle().
    Border(lipgloss.DoubleBorder()).
    BorderForeground(lipgloss.Color("#FFFFFF")).
    Padding(1, 2)
```

---

## Feedback Patterns

### Loading States

**Spinner (using bubbles/spinner):**
- Usage: During component scan, Git status check
- Style: Dot spinner or Line spinner
- Position: Inline with loading message
- Example: `o Scanning repositories...`

**Progress Bar (using bubbles/progress):**
- Usage: Build execution overall progress
- Style: Block-based progress bar with percentage
- Example: `[#########...........] 45%`

**Phase Indicators:**
- Build phases shown as text: `Compiling`, `Testing`, `Packaging`, `Installing`
- Color: Yellow/cyan for active phase

### Build Status Indicators

| Status | Symbol | Color | Description |
|--------|--------|-------|-------------|
| Pending | `o` | Gray | Queued, not yet started |
| Compiling | `.` (spinner) | Yellow | compile phase running |
| Testing | `.` (spinner) | Blue | test phase running |
| Packaging | `.` (spinner) | Purple | package phase running |
| Installing | `.` (spinner) | Cyan | install phase running |
| Success | `v` | Green | Build succeeded |
| Failed | `x` | Red | Build failed |

**Implementation:**
```go
statusStyles := map[string]lipgloss.Style{
    "pending":   lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")),
    "compiling": lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")),
    "testing":   lipgloss.NewStyle().Foreground(lipgloss.Color("#4169E1")),
    "packaging": lipgloss.NewStyle().Foreground(lipgloss.Color("#9370DB")),
    "installing": lipgloss.NewStyle().Foreground(lipgloss.Color("#00CED1")),
    "success":   lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32")),
    "failed":    lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6347")),
}
```

### Success Feedback

**Pattern:** Inline status change + summary in progress panel

**Build Success:**
- Component row: Green checkmark `v` + "Done"
- Summary badge: `Success: 3` (green)
- No modal/toast (inline feedback sufficient in TUI)

**Action Success:**
- Selection: Row highlight changes to indicate selection
- Configuration saved: Brief status message

### Error Handling

**Build Errors:**
- Component row: Red `x` + "Failed"
- Summary badge: `Failed: 1` (red)
- Press `L` to view detailed error logs

**Form Errors:**
- Pattern: Inline below field
- Color: Red text
- Clear: On valid input

**System Errors:**
- Pattern: Status bar message or dedicated error panel
- Message: User-friendly, actionable
- Example: "Maven not found. Check MAVEN_HOME configuration."

### Log Level Coloring

| Level | Color | Usage |
|-------|-------|-------|
| INFO | White/Default | Standard log messages |
| WARN | Yellow | Warning messages (e.g., deprecation) |
| ERROR | Red | Error messages |
| DEBUG | Gray/Dim | Debug information |

**Implementation (charmbracelet/log):**
```go
// Parse and colorize Maven output
infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))
errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6347"))
debugStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
```

---

## Empty States

### First-Run Empty State

**Component List (no components detected):**
```
+-- Build RFZ Components -------------------------------+
|                                                       |
|       No components detected.                         |
|                                                       |
|       Press [3] to go to Discover and scan paths,     |
|       or [4] to configure scan paths.                 |
|                                                       |
+-------------------------------------------------------+
```

**Pattern:** Helpful message + actionable guidance

### No Selection State

**Build Screen (nothing selected):**
```
Select components to build (0/13 selected)

[Build Selected] (Enter) is disabled until selection made
```

**Pattern:** Counter shows 0 selected, primary action disabled/dimmed

### No Results (Filter)

**Log Viewer (no logs match filter):**
```
No logs matching filter "ERROR"

Press [f] to change filter level or clear search.
```

### Error State

**Component Scan Failed:**
```
Failed to scan /path/to/components

Error: Directory not found

[Retry] (r)    [Edit Path] (e)
```

---

## Accessibility Patterns

### Keyboard Navigation

**Global Shortcuts:**
| Key | Action |
|-----|--------|
| `1-4` | Navigate to screen (Build, Logs, Discover, Config) |
| `q` | Quit application (with confirmation) |
| `?` | Show help |
| `Esc` | Close modal / Cancel / Back to navigation |
| `Tab` | Switch focus between panels |

**List Navigation:**
| Key | Action |
|-----|--------|
| `j` or `Down` | Move down |
| `k` or `Up` | Move up |
| `Space` | Toggle selection (multi-select) |
| `Enter` | Confirm / Open |
| `a` | Select all |
| `n` | Clear selection (none) |

**Build-Specific:**
| Key | Action |
|-----|--------|
| `b` | Start build (from component list) |
| `L` | View logs (during execution) |
| `Ctrl+C` | Cancel build |

**Log Viewer:**
| Key | Action |
|-----|--------|
| `f` | Toggle follow mode |
| `e` | Toggle error-only filter |
| `/` | Search |
| `1-5` | Set filter level (ALL, INFO, WARN, ERROR, DEBUG) |

### Focus Indicators

**Pattern:** Visible focus at all times

**Panel Focus:**
- Focused panel: Cyan border
- Unfocused panel: Gray/dim border

**Item Focus:**
- `>` prefix on current item
- Cyan text color OR
- Inverted background (dark on light)

**Button Focus:**
- `>` prefix before button
- Example: `> [Start Build] (Enter)`

**Implementation:**
```go
focusedItem := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FFFF")).
    Bold(true)

cursorPrefix := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FFFF")).
    Render("> ")
```

### Screen Reader Considerations

While TUI applications have limited screen reader support, follow these practices:

- Use semantic structure in output (clear sections, headers)
- Provide text alternatives for visual indicators
- Status changes announced in status bar
- Avoid relying solely on color (use symbols: `v` success, `x` failed)

### Color Independence

**All status indicators use both color AND symbol:**
- Success: Green `v` (not just green text)
- Failed: Red `x` (not just red text)
- Selected: `[x]` checkbox notation
- Unselected: `[ ]` checkbox notation
- Active radio: `(o)` filled
- Inactive radio: `( )` empty

---

## Status Bar Pattern

### Structure

```
+-- Status Bar (bottom of screen) ----------------------------------------+
| CONTEXT  item_name  | STATUS  | Tab Focus | Enter Select | Esc Back | q Quit |
+------------------------------------------------------------------------ +
```

**Components:**
1. **Context indicator**: Current screen/mode (BUILD, LOGS, SELECT, etc.)
2. **Item context**: Currently selected item name
3. **Status**: Running state if applicable (RUNNING, IDLE, etc.)
4. **Keyboard hints**: Available actions with shortcuts

**Implementation:**
```go
statusBar := lipgloss.NewStyle().
    Background(lipgloss.Color("#1E1E1E")).
    Foreground(lipgloss.Color("#FFFFFF")).
    Padding(0, 1)

contextBadge := lipgloss.NewStyle().
    Background(lipgloss.Color("#00CED1")).
    Foreground(lipgloss.Color("#000000")).
    Padding(0, 1).
    Bold(true)
```

---

## Scrolling & Viewport Patterns

### Long Lists (using bubbles/list)

**Behavior:**
- Virtual scrolling for large lists
- Scroll indicator on right side
- Focused item always visible
- Page up/down with `Ctrl+U/D` or `PgUp/PgDn`

### Log Viewport (using bubbles/viewport)

**Behavior:**
- Scrollable content area
- Line numbers displayed (optional)
- Scroll position indicator: `Line 1 of 53`
- Follow mode: Auto-scroll to new content
- Manual scroll: `j/k`, arrows, or mouse wheel

**Follow Mode:**
- Toggle with `f` key
- Indicator: `[Follow]` badge when active
- Pauses on manual scroll, resumes on toggle

---

## Error Recovery Patterns

### Build Failure Recovery

**Pattern:** Retry failed components only

```
Build complete with errors:
  v audiocon        Done (2m 15s)
  x traktion        Failed (1m 03s)
  v signalsteuerung Done (3m 22s)

[Rebuild Failed] (r)    [View Logs] (L)    [Back] (Esc)
```

**Actions:**
- `r`: Rebuild only failed components
- `L`: View logs to diagnose failure
- `Esc`: Return to component selection

### Cancel Confirmation

**Pattern:** Confirm before canceling long-running operation

```
Cancel build in progress?

3 components running, 2 completed.

[Continue Build]    [Cancel Build]
```

### Unsaved Configuration

**Pattern:** Warn before leaving with unsaved changes

```
Unsaved changes to scan paths.

[Save & Exit]    [Discard]    [Cancel]
```

---

## Charm.land Component Usage

### Required Bubbles Components

| Component | Usage | Configuration |
|-----------|-------|---------------|
| `list` | Component selection, navigation lists | Custom item renderer, filtering enabled |
| `viewport` | Log viewer content | Scrolling, line wrapping |
| `progress` | Build progress bars | Animated, percentage display |
| `spinner` | Loading indicators | Dot or Line style |
| `textinput` | Configuration inputs, search | Unicode support, placeholder |
| `table` | Component registry, detected components | Column sorting, selection |
| `help` | Keyboard shortcut display | Full/short modes |

### Lip Gloss Styling Requirements

**Borders (NEVER use manual box-drawing):**
```go
normalBorder := lipgloss.NormalBorder()   // Single line
doubleBorder := lipgloss.DoubleBorder()   // Modal borders
roundedBorder := lipgloss.RoundedBorder() // Friendly panels
```

**Colors (NEVER use raw ANSI codes):**
```go
cyan := lipgloss.Color("#00FFFF")
red := lipgloss.Color("#FF6347")
green := lipgloss.Color("#32CD32")
yellow := lipgloss.Color("#FFD700")
gray := lipgloss.Color("#808080")
```

**Layout (NEVER use manual string padding):**
```go
lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
```

---

## Implementation Guidelines

### Bubble Tea Model Structure

**Screen as Model:**
```go
type BuildScreen struct {
    list     list.Model      // Component list (bubbles)
    focused  FocusArea       // Current focus
    selected map[int]bool    // Selected components
    // ... other state
}
```

**Message Types:**
```go
type ComponentSelectedMsg struct{ Index int }
type BuildStartedMsg struct{ Components []Component }
type BuildProgressMsg struct{ Component string; Phase string }
type BuildCompleteMsg struct{ Results []BuildResult }
```

### Color Theme

| Element | Color | Hex |
|---------|-------|-----|
| Primary/Accent | Cyan | #00FFFF |
| Success | Green | #32CD32 |
| Error | Red | #FF6347 |
| Warning | Yellow | #FFD700 |
| Muted/Dim | Gray | #808080 |
| Background | Dark Gray | #1E1E1E |
| Text | White | #FFFFFF |

### Responsive Behavior

**Terminal Resize:**
- Recalculate layout on resize
- Maintain focus position
- Adjust panel proportions

**Minimum Size:**
- Width: 80 columns (graceful degradation)
- Height: 24 rows (minimum usable)
- Optimal: 120x40 (canonical for testing)

---

## UX Anti-Patterns to Avoid

**Navigation:**
- NO unclear focus state (always show focus)
- NO mouse-required interactions
- NO inconsistent shortcuts across screens

**Feedback:**
- NO blank screens during loading (show spinner/progress)
- NO generic errors ("Error 500")
- NO silent failures (always show status)

**TUI-Specific:**
- NO custom border characters (use Lip Gloss)
- NO raw ANSI escape codes (use Lip Gloss colors)
- NO manual string padding (use Lip Gloss layout)
- NO blocking operations without feedback

**Accessibility:**
- NO color-only status indicators (use symbols too)
- NO hidden keyboard shortcuts (display in UI)
- NO focus traps outside modals

---

## Notes & Assumptions

- Users are developers comfortable with terminal interfaces and vim-style navigation
- Keyboard-only operation is primary; mouse support is optional enhancement
- All styling MUST use Lip Gloss for consistency and testability
- Visual regression testing with teatest requires deterministic rendering
- Status bar always visible for context and available actions
- Modals use double border to distinguish from content panels
- Focus is always visible - no hidden focus states

---

## Testing Considerations

**Golden File Testing:**
- All UI states captured at 120x40 canonical size
- Deterministic rendering required (no animations in tests)
- 97 UI states to cover (from prototype screenshots)

**Visual Elements to Test:**
- Focus indicators visible and correct
- Status colors render correctly
- Modal overlay/borders distinct
- Empty states display helpful content
- Error states show actionable messages
