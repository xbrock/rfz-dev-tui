# Requirements Clarification - Welcome & Navigation Screen

**Created:** 2026-02-07
**Status:** Pending User Approval
**Roadmap Reference:** Sprint 2.1 (Phase 2: Screen Implementation)

## Feature Overview

Implement the main App Shell (Header, Navigation sidebar, Content area, StatusBar) with the Welcome screen as default view, placeholder screens for Build/Logs/Discover/Configuration, and full keyboard-driven navigation - transitioning the project from component library demos to a real application.

## Target Users

RFZ developers at Deutsche Bahn who use the terminal daily for Maven builds. They expect keyboard-driven workflows with vim-style navigation (j/k) and quick numeric shortcuts.

## Business Value

This is the foundational screen architecture that all subsequent screens will be built into. Without the App Shell, there is no application - only standalone component demos. This sprint transforms the component library into a usable application frame.

## Functional Requirements

### App Shell (Main Frame)
1. **Header Bar** (top of screen)
   - Left: "RFZ-CLI v1.0.0" title
   - Left below title: "Terminal Orchestration Tool" subtitle in cyan
   - Right: Real clock time (HH:MM:SS format) + "Deutsche Bahn Internal"
   - Bottom border in DB Red color
   - Header updates time every second via tick message

2. **Navigation Panel** (left sidebar, ~30 chars wide)
   - Bordered box with "Navigation" title
   - 5 menu items:
     - 1. Build Components (shortcut: 1)
     - 2. View Logs (shortcut: 2)
     - 3. Discover (shortcut: 3)
     - 4. Configuration (shortcut: 4)
     - 5. Exit (shortcut: q)
   - Focused item shows ">" prefix, cyan bold text, highlighted background
   - Below items: divider line, then key hints (Up/Down, Enter Select, 1-5 Quick nav)
   - Navigation wraps (bottom to top, top to bottom)

3. **Content Area** (right of navigation, fills remaining space)
   - Displays the currently active screen
   - Default: Welcome screen
   - Switches based on navigation selection
   - Bordered box with screen title in header position

4. **Status Bar** (bottom of screen, full width)
   - Left: Current mode badge (e.g., "HOME" in cyan background)
   - Left: Current focused item name (e.g., "Build Components")
   - Center: Context keyboard hints (Tab Focus, arrows Nav, Enter Select, Esc Back)
   - Right: "q Quit"
   - Uses existing TuiStatusBar component

### Welcome Screen (Content)
1. ASCII art "RFZ-CLI" logo
   - "RFZ" portion in DB Red/magenta
   - "CLI" portion in cyan
2. "Terminal Orchestration Tool" subtitle centered
3. Inspirational quote in italic muted text: *"First, solve the problem. Then, write the code."*
4. Divider line
5. Badge row: v1.0.0 badge (DB Red bg), "Deutsche Bahn" text, "Internal Tool" bordered badge
6. Status line: "$ rfz-cli ready" with blinking cursor
7. Help text: "Use navigation panel to get started"
8. Key hints: arrows navigate, Enter select, q quit

### Screen Switching
1. Selecting a nav item switches the content area to the corresponding screen
2. Placeholder screens for Build Components, View Logs, Discover, Configuration:
   - Show screen title
   - Show centered message: "[Screen Name] - Coming Soon"
   - Show hint: "Press Esc to return to Welcome"
3. Screens maintain no state between switches (stateless placeholders)

### Exit Behavior
1. Selecting "Exit" nav item or pressing q shows a confirmation modal dialog
2. Modal: "Are you sure you want to quit?" with Yes/No buttons
3. Uses existing TuiModal component
4. Yes = quit app, No = return to previous state

### Keyboard Navigation
1. **Global shortcuts:**
   - `q` - Show exit confirmation (from any screen)
   - `1-4` - Quick navigate to screens (always active)
   - `Tab` - Toggle focus between Navigation panel and Content area
2. **Navigation panel (when focused):**
   - `Up/k` - Move cursor up
   - `Down/j` - Move cursor down
   - `Enter` - Select focused item
3. **Content area (when focused):**
   - Content-specific keys (placeholder screens have no special keys)
   - `Esc` - Return focus to navigation / go back to Welcome

### Dynamic Layout
1. App responds to terminal resize events (tea.WindowSizeMsg)
2. Navigation panel: fixed width ~30 characters
3. Content area: fills remaining horizontal space
4. Height: full terminal height minus header and status bar
5. Minimum terminal size: 80x24 (show warning if smaller)

## Affected Areas & Dependencies

- **cmd/rfz/** - Replace existing component gallery with real application entry point
- **cmd/rfz-components-demo/** - Rename existing cmd/rfz/ gallery demo to this new location
- **internal/ui/app/** - NEW: Main application model (Bubble Tea model)
- **internal/ui/screens/** - NEW: Screen models (Welcome, placeholder screens)
- **internal/ui/components/** - EXISTING: Use TuiNavigation, TuiNavItem, TuiStatusBar, TuiModal, TuiKeyHints, TuiBox, TuiDivider, TuiStatus, TuiButton
- **internal/ui/components/styles.go** - May need new header/app-shell styles

## Edge Cases & Error Scenarios

- **Terminal too small** - If width < 80 or height < 24, show a centered message: "Terminal too small. Please resize to at least 80x24."
- **Rapid key presses** - Debounce not needed (Bubble Tea handles this)
- **Resize during modal** - Modal should reposition/resize correctly
- **Unknown key presses** - Ignored silently (no error messages)

## Security & Permissions

None - this is a local TUI application with no external connections in this sprint.

## Performance Considerations

- Clock tick updates every 1 second (not more frequent to avoid CPU usage)
- Render only when state changes (Bubble Tea's standard behavior)
- No blocking operations in Update() function

## Scope Boundaries

**IN SCOPE:**
- App Shell layout (Header, Navigation, Content, StatusBar)
- Welcome screen with full visual content
- Placeholder screens for Build/Logs/Discover/Configuration
- Screen switching via navigation
- Keyboard navigation (global + per-panel)
- Exit confirmation modal
- Terminal resize handling
- Real clock time in header
- Renaming existing cmd/rfz to cmd/rfz-components-demo
- Visual regression tests for key states

**OUT OF SCOPE:**
- Real Build Components screen (Sprint 2.2)
- Real Log Viewer screen (Sprint 2.3)
- Real Discover screen (Sprint 2.3)
- Real Configuration screen (Sprint 2.4)
- Maven execution, Git operations
- Configuration persistence
- Any real domain logic / data
- Theme customization

## Open Questions

None - all requirements clarified.

## Proposed User Stories (High Level)

1. **App Shell & Bubble Tea Model** - Create main application model with header, navigation, content area, and status bar layout
2. **Welcome Screen** - Implement the welcome/home screen content (ASCII art, badges, status, hints)
3. **Navigation & Screen Switching** - Keyboard navigation, focus management, and screen switching with placeholder screens
4. **Exit Confirmation Modal** - Quit confirmation dialog using TuiModal
5. **Entry Point & Demo Rename** - Replace cmd/rfz with real app, rename old demo to cmd/rfz-components-demo
6. **Visual Regression Tests** - Golden file tests for all key UI states (welcome, nav focus, placeholders, modal, resize)

---

*Review this document carefully. Once approved, an Implementation Plan will be created, followed by detailed user stories.*
