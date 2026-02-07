# Requirements Clarification - Build Screens (Sprint 2.2)

**Created:** 2026-02-07
**Status:** Pending User Approval

## Feature Overview

Implement the three Build sub-screens within the "Build Components" navigation tab: Component Selection (multi-select list with category badges), Build Configuration Modal (Maven goals, profiles, port, skip tests, command preview), and Build Execution View (real-time progress tracking with per-component status, phases, and timing). All using mocked domain data with an interface for Phase 3 swap.

## Target Users

RFZ developers at Deutsche Bahn who need to select, configure, and monitor Maven builds for multiple component types (Core, Simulator, Standalone).

## Business Value

Transforms the placeholder Build Components screen into a fully interactive build workflow UI. This is the core value proposition of the RFZ CLI - making Maven builds visual and keyboard-driven. Completing Sprint 2.2 enables the primary user journey: select components -> configure build -> monitor execution.

## Functional Requirements

### 1. Build Component Selection Screen
- Display scrollable list of RFZ components with multi-select (checkboxes)
- Each component shows a category badge: Core (cyan), Simulation (yellow), Standalone (green)
- Selection counter: "Select components to build (X/N selected)"
- Keyboard shortcuts: Space=toggle, a=Select All, n=Clear Selection
- Cursor indicator (>) on focused component
- Highlighted row for current cursor position
- Actions section with buttons: [Build Selected] (Enter), [Select All] (a), [Clear Selection] (n)
- Tab switches focus between component list and actions
- Legend at bottom: [x] Selected, [ ] Not selected, > Current
- Status bar shows: SELECT mode, current component name, keyboard hints

### 2. Build Configuration Modal
- Opens when user triggers "Build Selected" with components selected
- Shows "Building N components: comp1, comp2, comp3" header
- **Maven Goal** (radio): clean, install, package, clean install (default: clean install)
- **Maven Profiles** (multi-select checkboxes): target_env_dev (default: checked), generate_local_config_files
- **Traktion Port** (radio): Port 11090 (default), Port 11091 - appends use_traktion_* profile
- **Build Options** (checkbox): Skip Tests (default: checked) - adds -DskipTests
- **Command Preview**: Live-updating preview of the assembled Maven command (e.g., `$ mvn clean-install -Pgenerate_local_config_files,use_traktion_11090 -DskipTests`)
- Tab navigates between sections
- Arrow keys / h/l navigate within sections
- Space/Enter toggles options
- Cancel (Esc) / Start Build (Enter on Start button)
- Double border, centered overlay, dims background

### 3. Build Execution View
- Replaces the component selection view in-place (same Build tab)
- **Command display**: Shows the full Maven command being executed
- **Components table**: Columns: St (status icon), Component, Phase, Progress, Time
  - Per-component status: Pending (circle), Running (spinner), Success (checkmark), Failed (X)
  - Build phases: Pending -> Compiling -> Testing -> Packaging -> Installing -> Done/Failed
  - Per-component progress bar
  - Elapsed time per component (MM:SS)
- **Progress section**: Overall progress bar with percentage, status counters (Running: N, Success: N, Failed: N, Pending: N)
- **Actions section**: [View Logs] (l), [Cancel Build] (Ctrl+C)
- **Cursor navigation**: Navigate between components in the list
- **Status bar**: BUILD mode, current component name, RUNNING badge

### 4. Build Completion
- When all components finish: show final status per component (Success/Failed)
- Progress bar shows 100% (or partial if failures)
- Actions change to: [New Build] (to return to selection), [View Logs] (l)
- Status bar updates to show COMPLETE or FAILED

### 5. Build Cancellation
- Ctrl+C or Cancel Build button stops all builds immediately
- Running components switch to "Cancelled" status
- Show summary with cancelled state
- Allow "New Build" to start fresh

### 6. Mock Build Simulation
- Domain interface: `ComponentProvider` with mock implementation
- Hardcoded 13 components matching prototype: boss, fistiv, audiocon, traktion, signalsteuerung, weichensteuerung, simkern, fahrdynamik, energierechnung, zuglauf, stellwerk, diagnose, konfiguration
- Component types: Core, Simulation, Standalone (matching prototype badges)
- Timed simulation: each component progresses through phases over ~5-15 seconds
- ~20% random failure chance per component (fails at random phase)
- All components build in parallel (simulated)

## Affected Areas & Dependencies

- **internal/app/app.go** - Replace placeholder Build screen with real Build model, add modal handling for config
- **internal/ui/screens/build/** (NEW) - Build selection, config modal, execution views
- **internal/domain/** (NEW) - Component, BuildConfig, BuildStatus domain types with mock provider
- **internal/ui/components/** - Existing components used extensively (TuiList, TuiModal, TuiRadio, TuiCheckbox, TuiButton, TuiStatus, TuiProgress, TuiSpinner, TuiBox, TuiStatusBar, TuiKeyHints, TuiTable, TuiDivider)
- **internal/ui/components/styles.go** - May need additional style definitions for build-specific elements
- **internal/app/messages.go** - New message types for build state transitions

## Edge Cases & Error Scenarios

- **No components selected**: "Build Selected" button is disabled/no-op, show hint to select components
- **Build failure**: Component shows "Failed" status in red, phase shows where it failed, other components continue
- **Build cancellation**: All running components immediately stop, show "Cancelled" status
- **Terminal resize during build**: Layout reflows, progress bars adjust width
- **Empty component list**: Show informative empty state (shouldn't happen with mock data)
- **Modal escape**: Esc cancels configuration and returns to selection without starting build
- **All components fail**: Summary shows all failed, "New Build" option available

## Security & Permissions

Not applicable for Phase 2 (mocked UI). No real Maven execution or file system access.

## Performance Considerations

- Timer-based simulation should use Bubble Tea's tick mechanism (already exists as TickMsg)
- Spinner animations should use bubbles/spinner (already wrapped)
- Progress bar updates should be smooth (bubbles/progress handles this)
- Component list should handle scrolling efficiently for 13+ items

## Scope Boundaries

**IN SCOPE:**
- Build Component Selection screen with full keyboard navigation
- Build Configuration Modal with all 5 sections
- Build Execution View with timed simulation
- Mock domain model with interface for Phase 3
- Golden file visual tests for all prototype states (~25 screenshots: 10-17, 20-34, 40-49)
- Build cancellation and failure states
- Status bar updates per view state

**OUT OF SCOPE:**
- Real Maven execution (Phase 3)
- Real component scanning/discovery (Phase 3)
- Build history/statistics (Future)
- Log viewer integration (Sprint 2.3 - only navigation stub)
- Configuration persistence (Phase 3)
- Component filtering/search in the list
- Parallel build orchestration (Phase 3)

## Open Questions

None - all requirements clarified through dialog.

## Proposed User Stories (High Level)

1. **Domain Model & Mock Data** - Create domain types (Component, BuildConfig, BuildResult) with ComponentProvider interface and mock implementation
2. **Build Component Selection Screen** - Multi-select component list with badges, actions, keyboard navigation
3. **Build Configuration Modal** - Full modal with 5 sections (Goal, Profiles, Port, Options, Preview)
4. **Build Execution View** - Progress tracking with timed simulation, per-component status, phases
5. **App Integration & Screen Transitions** - Wire build screens into app model, replace placeholder, handle transitions between selection/config/execution views

---
*Review this document carefully. Once approved, detailed user stories will be generated.*
