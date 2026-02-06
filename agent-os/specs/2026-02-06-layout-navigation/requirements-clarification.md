# Requirements Clarification - Layout & Navigation Components

**Created:** 2026-02-06
**Status:** Pending User Approval
**Sprint:** 1.3 - Layout & Navigation (Week 3)

## Feature Overview

Implementation of 8 Layout & Navigation TUI components for the RFZ Developer CLI, completing the component library foundation. These components enable screen navigation, data display, overlay dialogs, and keyboard hint visualization.

## Target Users

RFZ developers at Deutsche Bahn who need intuitive terminal navigation, data tables for component listing, modal dialogs for configuration, and clear keyboard shortcuts display.

## Business Value

- **Completes Foundation Phase:** Sprint 1.3 finalizes the component library, enabling screen implementation in Phase 2
- **Navigation Ready:** App shell and screen switching infrastructure prepared
- **Data Visualization:** Tables and trees ready for component registry and dependency views
- **User Guidance:** Keyboard hints and status bar improve discoverability

## Functional Requirements

### 1. TuiNavigation (Sidebar Container)
- Vertical sidebar container for navigation items
- Fixed width (configurable)
- Optional header section
- Optional footer section (for keyboard hints)
- Border with focus state support
- Scroll support for many items

### 2. TuiNavItem (Navigation Item)
- Three states: Normal, Focused (cursor), Active (current screen)
- Keyboard shortcut display (e.g., "1", "2", "3")
- Icon support (optional, ASCII-based)
- Hover/cursor indicator ("> " prefix)

### 3. TuiModal (Overlay Dialog)
- Double border styling (prominent)
- Centered positioning
- Title bar
- Content area (scrollable if needed)
- Footer with action buttons
- **Focus Trapping:** Escape closes, Tab cycles through buttons
- Backdrop rendering (dim background)

### 4. TuiTable (Data Table)
- Wraps `bubbles/table` component
- RFZ styling applied (colors, borders)
- Selectable rows
- Column headers
- Zebra striping (optional)
- Pagination support (via bubbles)

### 5. TuiTabs (Tab Navigation)
- Horizontal tab bar
- **Numeric shortcuts (1-9)** for direct tab access
- Visual indicator for active tab
- Focus state for keyboard navigation
- Tab labels with optional counts/badges

### 6. TuiStatusBar (Bottom Bar)
- Full-width bottom bar
- Left/center/right sections
- Status indicators (e.g., "Build: Running")
- Context-sensitive keyboard hints
- Separator from main content

### 7. TuiTree (Tree View)
- **Simple expand/collapse functionality**
- Indentation for hierarchy levels
- Expand/collapse icons (▶/▼ or +/-)
- Keyboard navigation (j/k or arrows)
- Node labels with optional metadata
- For: Component dependencies, file structures

### 8. TuiKeyHints (Keyboard Shortcuts Display)
- Compact display of available shortcuts
- Key + Label format (e.g., "Enter Select")
- Separator between items
- Adaptive width (fits available space)
- Context-aware (shows current screen shortcuts)

## Affected Areas & Dependencies

### Existing Components Used
- `internal/ui/components/styles.go` - Color tokens, border styles
- `internal/ui/components/helpers.go` - Render utilities
- `internal/ui/components/box.go` - Container patterns

### External Dependencies
- `github.com/charmbracelet/bubbles/table` - For TuiTable
- `github.com/charmbracelet/bubbles/viewport` - For scrollable content
- `github.com/charmbracelet/lipgloss` - All styling

### New Files to Create
- `internal/ui/components/navigation.go` - TuiNavigation + TuiNavItem
- `internal/ui/components/modal.go` - TuiModal
- `internal/ui/components/table.go` - TuiTable wrapper
- `internal/ui/components/tabs.go` - TuiTabs
- `internal/ui/components/statusbar.go` - TuiStatusBar
- `internal/ui/components/tree.go` - TuiTree
- `internal/ui/components/keyhints.go` - TuiKeyHints

## Edge Cases & Error Scenarios

| Edge Case | Expected Behavior |
|-----------|-------------------|
| Modal without buttons | Escape still closes, no Tab cycling |
| Empty table | Show "No data" message |
| Empty tree | Show "No items" message |
| Tabs > 9 | Only first 9 have numeric shortcuts |
| Very long nav items | Truncate with ellipsis |
| Tree too deep | Limit visible depth, show "..." |
| Modal larger than terminal | Scroll content, modal respects bounds |

## Security & Permissions

Not applicable - UI components only, no external data access.

## Performance Considerations

- Table should handle 100+ rows without lag
- Tree should handle 50+ nodes with reasonable performance
- Modal backdrop rendering should be efficient
- No unnecessary re-renders on unchanged state

## Scope Boundaries

**IN SCOPE:**
- All 8 components from Sprint 1.3
- Unit tests (teatest) for all components
- Golden file tests for key variants
- **Separate Demo Program** for new components
- Integration with existing styles package
- Bubbles wrapper approach for table

**OUT OF SCOPE:**
- Screen implementation (Phase 2)
- Real data integration
- Extending existing Gallery (separate demo instead)
- Horizontal navigation variant
- Multi-select tree
- Drag-and-drop functionality

## Open Questions

None - all requirements clarified.

## Proposed User Stories (High Level)

1. **TuiNavigation + TuiNavItem** - Vertical navigation sidebar with items
2. **TuiModal** - Overlay dialog with focus trapping
3. **TuiTable** - Bubbles/table wrapper with RFZ styling
4. **TuiTabs** - Tab navigation with numeric shortcuts
5. **TuiStatusBar** - Bottom status bar component
6. **TuiTree** - Simple expand/collapse tree view
7. **TuiKeyHints** - Keyboard shortcuts display
8. **Layout Navigation Demo** - Separate demo program showcasing all new components
9. **Visual Regression Tests** - Golden file tests for all component states

---

*Review this document carefully. Once approved, the implementation plan and detailed user stories will be generated.*
