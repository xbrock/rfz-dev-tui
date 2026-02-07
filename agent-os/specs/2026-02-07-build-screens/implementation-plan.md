# Implementation Plan - Build Screens (Sprint 2.2)

**Created:** 2026-02-07
**Status:** DRAFT
**Spec:** 2026-02-07-build-screens
**Input:** Approved requirements-clarification.md

---

## Executive Summary

Implement three interconnected Build sub-views (Component Selection, Configuration Modal, Build Execution) within the existing "Build Components" navigation tab. The implementation replaces the current placeholder screen, introduces a lightweight domain model with mock data provider, and adds timed build simulation with visual states matching the 16+ prototype screenshots.

## Architektur-Entscheidungen

### 1. Screen as Bubble Tea Sub-Model (not raw functions)

The Build screen is complex with multiple internal states (selection, modal, execution). Unlike the passive Welcome screen, it needs its own `Update()` method. Follow the sub-model pattern:

```
app.Model → build.Model (with internal state machine)
```

The build screen owns its internal focus, cursor, and view state. The parent app delegates key events and size updates to it.

### 2. State Machine for View Transitions

The Build screen has three distinct modes/phases:

```
selecting → configuring → executing → completed
    ↑_________|  (cancel)     ↑_________|  (new build)
```

Use a simple enum `buildPhase` to track which view to render. This avoids the complexity of multiple model instances.

### 3. Domain Model with Interface (Phase 3-Ready)

Create a thin domain layer with:
- `Component` type (simplified from boilerplate - only fields needed for UI)
- `BuildConfig` value object (goal, profiles, port, skipTests)
- `ComponentProvider` interface with `MockComponentProvider` implementation
- `BuildSimulator` for timed mock builds

The boilerplate domain models serve as reference but are too heavy for Phase 2. We'll create focused types.

### 4. Reuse Existing Components Extensively

All visual rendering uses existing components from `internal/ui/components/`:
- `TuiList` with `ListMultiSelect` for component selection (badges already supported)
- `TuiModal` for build configuration overlay
- `TuiRadio` for Maven Goal and Traktion Port
- `TuiCheckbox` for Maven Profiles and Skip Tests
- `TuiButton` for action buttons
- `TuiStatus` for per-component build status
- `TuiProgress` for overall progress bar
- `TuiSpinner` for running component animation
- `TuiBox` for section containers
- `TuiKeyHints` / `TuiStatusBar` for hints

### 5. Build Config Modal: Extend TuiModal

The existing `TuiModal` supports Title + Content + Buttons. The Build Config Modal needs multiple form sections inside Content. We'll compose the modal content from existing form components (TuiRadio, TuiCheckbox) rendered as a single string, passed to `TuiModal.Content`. The modal already handles double border, backdrop, and centering.

However, the config modal needs sectional focus (Tab navigates between sections), which requires a custom Build Config Modal model that manages internal focus state and delegates to the existing TuiModal for rendering.

### 6. Build Execution: Custom View with Existing Components

The execution table (St, Component, Phase, Progress, Time) needs per-row spinners and progress bars. This requires a custom rendering approach using Lip Gloss layout, composing `TuiStatus`, `TuiSpinner`, and `TuiProgress` per row. The existing `TuiTable` (bubbles/table wrapper) doesn't support inline widget rendering per cell.

### 7. Testing Strategy

- Golden file tests using teatest for all distinct UI states
- Static test helpers that set specific model states
- Match prototype screenshots series: 10-17, 20-34, 40-49
- Canonical terminal size: 120x40

## Komponenten-Uebersicht

### Neue Komponenten

| Komponente | Typ | Beschreibung |
|------------|-----|--------------|
| `internal/domain/component.go` | Domain Model | Component type, ComponentType enum, ComponentProvider interface |
| `internal/domain/buildconfig.go` | Domain Model | BuildConfig, MavenGoal, BuildPhase types |
| `internal/domain/mock_provider.go` | Mock Data | MockComponentProvider with 13 hardcoded components |
| `internal/ui/screens/build/model.go` | Screen Model | Build screen state machine (selecting/configuring/executing/completed) |
| `internal/ui/screens/build/update.go` | Screen Logic | Key handling for all build phases |
| `internal/ui/screens/build/view.go` | Screen View | Renders the appropriate view based on current phase |
| `internal/ui/screens/build/selection.go` | Sub-View | Component selection list + actions rendering |
| `internal/ui/screens/build/config.go` | Sub-View | Build configuration modal with 5 sections |
| `internal/ui/screens/build/execution.go` | Sub-View | Build execution progress view |
| `internal/ui/screens/build/simulator.go` | Simulation | Timed build simulation with random failures |

### Aenderungen an Bestehendem

| Komponente | Aenderung |
|------------|-----------|
| `internal/app/app.go` | Replace `phBuild placeholder.Model` with `build build.Model`, delegate update/view |
| `internal/app/messages.go` | Add BuildTickMsg, BuildPhaseMsg, BuildCompleteMsg types |
| `internal/ui/components/styles.go` | Add build-specific styles if needed (execution table header, phase colors) |
| `internal/ui/components/list.go` | Possibly extend badge styling for component type coloring (Core=cyan, Sim=yellow, Standalone=green) |

### Komponenten-Verbindungen

| Source | Target | Verbindung | Zustaendige Story |
|--------|--------|------------|-------------------|
| `app.Model` | `build.Model` | App creates, delegates Update/View, passes size | Story 5 (App Integration) |
| `build.Model` | `domain.ComponentProvider` | Build screen gets component list from provider | Story 1 (Domain) → Story 2 (Selection) |
| `build.Model` | `domain.BuildConfig` | Config modal produces BuildConfig, execution uses it | Story 3 (Config Modal) → Story 4 (Execution) |
| `build/selection` | `components.TuiList` | Renders component list using existing TuiList | Story 2 (Selection) |
| `build/config` | `components.TuiModal` | Renders config inside existing modal | Story 3 (Config Modal) |
| `build/execution` | `components.TuiStatus/Progress/Spinner` | Renders per-component status row | Story 4 (Execution) |
| `build/simulator` | `build.Model` | Simulator sends BuildTickMsg/BuildPhaseMsg to drive execution | Story 4 (Execution) |
| `app.messages` | `build.Model` | New message types flow from app.Update to build.Update | Story 5 (App Integration) |

## Umsetzungsphasen

### Phase A: Domain Foundation
1. Create domain types (Component, BuildConfig, BuildPhase, MavenGoal)
2. Create ComponentProvider interface + mock implementation
3. Unit tests for domain logic (BuildConfig.ToCommand(), etc.)

### Phase B: Build Selection Screen
1. Create build screen model with state machine
2. Implement selection view using TuiList with badges
3. Add keyboard handling (space, a, n, enter, tab)
4. Add actions section with buttons

### Phase C: Build Configuration Modal
1. Create config modal model with sectional focus
2. Implement 5 sections (Goal, Profiles, Port, Options, Preview)
3. Command preview updates dynamically based on selections
4. Wire modal open/close with selection screen

### Phase D: Build Execution View
1. Create execution view with component table
2. Implement build simulator with timed phase transitions
3. Add per-component status, spinner, progress, time
4. Add overall progress bar and status counters
5. Handle completion, failure, and cancellation

### Phase E: App Integration
1. Replace placeholder with build.Model in app.go
2. Wire message passing (BuildTickMsg, etc.)
3. Handle focus delegation (app → build)
4. Update status bar per build phase
5. Golden file tests for all states

## Abhaengigkeiten

```
Story 1 (Domain) ← Story 2 (Selection) ← Story 3 (Config Modal)
                                                    ↓
                                        Story 4 (Execution)
                                                    ↓
                                        Story 5 (App Integration)
```

- Story 1 has no dependencies (foundational)
- Story 2 depends on Story 1 (needs Component type)
- Story 3 depends on Story 2 (needs selection state to know which components to configure)
- Story 4 depends on Story 1 + Story 3 (needs BuildConfig + Component types)
- Story 5 depends on all others (integration)

## Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| TuiModal too rigid for config form | Mittel | Mittel | Compose content string from form components; if insufficient, render modal-like overlay manually with TuiBox + double border |
| Build simulation timing brittle in tests | Hoch | Niedrig | Use deterministic time steps in tests (not real timers); golden files test static states only |
| Execution table layout complex | Mittel | Mittel | Start simple (one row per component), iterate on column alignment; use Lip Gloss Place/Join |
| State machine complexity | Niedrig | Mittel | Keep phases as simple enum; avoid nested state machines |

---

## Self-Review Ergebnisse

### 1. Vollstaendigkeit
- All requirements from clarification are covered
- Component selection: full multi-select with badges, actions, keyboard shortcuts
- Config modal: all 5 sections (Goal, Profiles, Port, Options, Preview)
- Execution: per-component status, phases, progress, timing, cancellation
- Mock simulation with random failures
- Golden file tests for all prototype states

### 2. Konsistenz
- Architecture follows existing patterns (sub-model, function components, styles.go)
- Naming conventions match codebase (TuiXxx for components, lowercase packages)
- State management follows Elm architecture (Model/Update/View)
- No contradictions found

### 3. Risiken
- TuiModal flexibility is the main concern - mitigated by composition approach
- Build tick timing in tests - mitigated by deterministic test helpers
- No critical blockers identified

### 4. Alternativen
- **Alternative: Separate screen per build phase** - Rejected. The prototype clearly shows one tab with view transitions. Separate screens would break navigation model.
- **Alternative: Use bubbles/list instead of TuiList** - Rejected. TuiList already exists, matches the prototype styling, supports badges and multi-select. Switching would duplicate work.
- **Alternative: Full domain model now** - Rejected. Boilerplate domain model has fields not needed for UI (GroupID, ArtifactID, Version, Dependencies). Lighter types are more maintainable.

### 5. Komponenten-Verbindungen
- All new components have at least one connection (verified)
- All connections have assigned stories (verified)
- No orphaned components found
- Validation:
  - `domain/*` → connected to `build.Model` via Story 1/2
  - `build/model` → connected to `app.Model` via Story 5
  - `build/selection` → connected to `TuiList` via Story 2
  - `build/config` → connected to `TuiModal` via Story 3
  - `build/execution` → connected to `TuiStatus/Progress/Spinner` via Story 4
  - `build/simulator` → connected to `build.Model` via Story 4

---

## Minimalinvasiv-Optimierungen

### 1. Wiederverwendung

| Bestehender Code | Wiederverwendung |
|------------------|-----------------|
| `TuiList` + `TuiListItem` + helpers | Direct use for component selection - badges, multi-select, SelectAll/DeselectAll already implemented |
| `TuiModal` + `TuiModalConfig` | Container for config dialog - title, backdrop, centering already work |
| `TuiRadio` | Maven Goal and Traktion Port selection |
| `TuiCheckbox` | Skip Tests toggle |
| `TuiStatus` (all 6 variants) | Per-component build status display |
| `TuiProgress` (bubbles wrapper) | Overall and per-component progress |
| `TuiSpinner` (bubbles wrapper) | Running state animation |
| `TuiButton` (3 variants) | All action buttons |
| `TuiBox` | Section containers (Components, Progress, Actions) |
| `TuiDivider` | Section separators |
| `TuiKeyHints` | In-component keyboard hints |
| `TuiStatusBar` + config | Bottom bar (existing, needs mode/context updates) |
| `TickMsg` + `tickCmd()` | Timer infrastructure for build simulation |
| `focusArea` + Tab handling | Focus model between nav and content |
| Welcome screen pattern | Sub-model with New/SetSize/Update/View |
| Quit modal pattern | Modal overlay with key capture |

### 2. Aenderungsumfang

**Minimal changes needed:**
- `app.go`: ~20 lines changed (replace placeholder field + switch cases)
- `messages.go`: ~15 lines added (new message types)
- `styles.go`: ~10-20 lines added (build-specific colors/styles, if needed)
- `list.go`: 0 changes (badge coloring can be done at call site via TuiListItem.Badge)

**New code (estimated):**
- `domain/`: ~150 LOC (types + mock provider)
- `build/model.go`: ~100 LOC (state machine + fields)
- `build/update.go`: ~200 LOC (key handling for all phases)
- `build/view.go`: ~50 LOC (dispatcher to sub-views)
- `build/selection.go`: ~100 LOC (compose TuiList + actions)
- `build/config.go`: ~200 LOC (modal with 5 sections)
- `build/execution.go`: ~200 LOC (table + progress + actions)
- `build/simulator.go`: ~100 LOC (timed phase transitions)
- Tests: ~300 LOC

**Total estimated: ~1400 LOC new, ~50 LOC changed**

### 3. Feature-Preservation Checkliste

- [x] All requirements from Clarification are covered
- [x] No feature was sacrificed
- [x] All acceptance criteria remain achievable
- [x] Existing functionality (welcome, navigation, quit modal) unaffected
- [x] Existing golden file tests remain valid

---

*Created with Agent OS /create-spec v3.0 - Implementation Plan (Kollegen-Methode)*
