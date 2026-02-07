# Implementation Plan: Welcome & Navigation Screen

**Created:** 2026-02-07
**Spec:** 2026-02-07-welcome-navigation-screen
**Status:** Draft
**Based on:** Approved requirements-clarification.md

---

## Executive Summary

Build the main App Shell (Header, Navigation sidebar, Content area, StatusBar) with Welcome screen as default view, placeholder screens for Build/Logs/Discover/Configuration, and full keyboard navigation. This transforms the project from component demos into a real application by establishing the `internal/app/` and `internal/ui/screens/` layers as defined in the architecture decisions (DEC-001 through DEC-006).

---

## Architektur-Entscheidungen

### Pattern: Hierarchical Model Composition (DEC-002)

The main `app.Model` in `internal/app/app.go` acts as root Bubble Tea model. It owns:
- Screen routing (which screen is active)
- Global state (terminal dimensions, focus panel, clock)
- Navigation state (cursor, active item)
- Modal overlay state

Child screen models (welcome, placeholder screens) handle their own state.

### Vereinfachung fuer Sprint 2.1

Since Sprint 2.1 uses **no domain logic, no services, no infrastructure**, we defer the full service layer:
- No `internal/service/`, `internal/domain/`, `internal/infra/` in this sprint
- Main entry point creates `app.Model` directly without dependency injection
- Placeholder screens are trivial (no state, static render)
- Welcome screen is stateless (renders from app-level state like version string)

This keeps the implementation minimal while establishing the correct architecture structure for future sprints.

### Message Flow

```
tea.Program
    |
    v
app.Model.Update()
    |
    +-- Global keys (q, 1-4, Tab) handled here
    +-- Modal overlay (if showing, captures all input)
    +-- Routes to active screen's Update()
    |
    v
screen.Model.Update()
    |
    +-- Screen-specific keys
    +-- Returns tea.Cmd if needed
```

### Layout Composition (View)

```
+--[ Header ]------------------------------------------+
| RFZ-CLI v1.0.0                  3:40 PM | DB Internal |
| Terminal Orchestration Tool                            |
+-------------------------------------------------------+
|  Navigation  |  Content Area (active screen)           |
|  (sidebar)   |                                         |
|  30 chars    |  fills remaining width                  |
|              |                                         |
+-------------------------------------------------------+
| [HOME] Build Components | Tab Focus | Enter | q Quit  |
+--[ StatusBar ]--------------------------------------- +
```

---

## Komponenten-Uebersicht

### Neue Komponenten/Dateien

| Komponente | Pfad | Beschreibung |
|------------|------|-------------|
| App Model | `internal/app/app.go` | Root Bubble Tea model with layout, routing, state |
| App Messages | `internal/app/messages.go` | Shared message types (NavigateMsg, TickMsg) |
| Welcome Screen | `internal/ui/screens/welcome/welcome.go` | Welcome screen with ASCII art, badges, hints |
| Placeholder Screen | `internal/ui/screens/placeholder/placeholder.go` | Generic "Coming Soon" placeholder for Build/Logs/Discover/Config |
| Entry Point | `cmd/rfz/main.go` | Replace gallery with real app |
| Demo Rename | `cmd/rfz-components-demo/main.go` | Renamed component gallery demo |

### Bestehende Komponenten (wiederverwendet, NICHT geaendert)

| Komponente | Pfad | Nutzung |
|------------|------|---------|
| TuiNavigation | `internal/ui/components/navigation.go` | Sidebar navigation rendering |
| TuiStatusBar | `internal/ui/components/statusbar.go` | Bottom status bar |
| TuiModal | `internal/ui/components/modal.go` | Exit confirmation dialog |
| TuiKeyHints | `internal/ui/components/keyhints.go` | Keyboard hint rendering |
| TuiBox | `internal/ui/components/box.go` | Bordered containers |
| TuiButton | `internal/ui/components/button.go` | Modal buttons |
| TuiDivider | `internal/ui/components/divider.go` | Visual separators |
| Design System | `internal/ui/components/styles.go` | All color tokens, styles |
| Component Gallery | `internal/ui/components/demo/gallery.go` | Existing demo (moved) |

### Potenzielle Aenderungen an bestehenden Dateien

| Datei | Aenderung | Grund |
|-------|-----------|-------|
| `internal/ui/components/styles.go` | Add header styles if missing | Header bar needs specific styles (title, subtitle, right-aligned info) |
| `go.mod` | No change expected | All dependencies already present |

---

## Komponenten-Verbindungen

| Source | Target | Verbindungsart | Zustaendige Story |
|--------|--------|---------------|-------------------|
| `cmd/rfz/main.go` | `internal/app/app.go` | Creates & runs App Model | Story 1 (Entry Point) |
| `internal/app/app.go` | `internal/ui/screens/welcome/` | Embeds welcome model, routes messages | Story 2 (App Shell) |
| `internal/app/app.go` | `internal/ui/screens/placeholder/` | Embeds placeholder models, routes messages | Story 3 (Screen Switching) |
| `internal/app/app.go` | `internal/ui/components/navigation.go` | Calls TuiNavigation() in View | Story 2 (App Shell) |
| `internal/app/app.go` | `internal/ui/components/statusbar.go` | Calls TuiStatusBar() in View | Story 2 (App Shell) |
| `internal/app/app.go` | `internal/ui/components/modal.go` | Calls TuiModal() for exit confirm | Story 4 (Exit Modal) |
| `internal/app/app.go` | `internal/app/messages.go` | Uses shared message types | Story 2 (App Shell) |

---

## Umsetzungsphasen

### Phase 1: Entry Point & Demo Rename
1. Rename `cmd/rfz/` to `cmd/rfz-components-demo/`
2. Create new `cmd/rfz/main.go` that creates the App Model
3. Ensure both entry points compile and run

### Phase 2: App Shell (Core)
1. Create `internal/app/messages.go` with shared types
2. Create `internal/app/app.go` with root Model:
   - Header rendering (title, subtitle, clock, right info)
   - Navigation sidebar using TuiNavigation
   - Content area routing
   - StatusBar using TuiStatusBar
   - Terminal resize handling
   - Tick command for clock
3. Layout composition in View()

### Phase 3: Welcome Screen
1. Create `internal/ui/screens/welcome/welcome.go`
   - ASCII art logo
   - Subtitle, quote, divider
   - Version badge, DB badge, Internal Tool badge
   - Status line ("$ rfz-cli ready")
   - Key hints

### Phase 4: Screen Switching & Placeholders
1. Create `internal/ui/screens/placeholder/placeholder.go`
   - Generic screen with title and "Coming Soon" message
2. Wire navigation selection to screen switching
3. Implement focus management (Tab between nav/content)
4. Keyboard routing (global keys, per-panel keys)

### Phase 5: Exit Confirmation Modal
1. Implement modal overlay in app.go using TuiModal
2. Show modal on q press or Exit nav selection
3. Yes = tea.Quit, No = dismiss modal
4. Modal captures all input while shown

### Phase 6: Visual Regression Tests
1. Create `internal/app/app_test.go` with golden file tests
2. Test states: welcome default, each nav focus, placeholder screens, exit modal, too-small terminal
3. Canonical size: 120x40

---

## Abhaengigkeiten

```
Phase 1 (Entry Point) ─── unabhaengig
Phase 2 (App Shell) ──── abhaengig von Phase 1
Phase 3 (Welcome) ────── abhaengig von Phase 2
Phase 4 (Navigation) ─── abhaengig von Phase 2
Phase 5 (Exit Modal) ─── abhaengig von Phase 2
Phase 6 (Tests) ──────── abhaengig von Phase 2-5
```

Phase 3, 4, 5 koennen parallel nach Phase 2.

---

## Risiken & Mitigationen

| Risiko | Likelihood | Impact | Mitigation |
|--------|-----------|--------|------------|
| Header styles not in styles.go | Medium | Low | styles.go already has StyleHeader, StyleHeaderTitle, StyleHeaderSubtitle, StyleFooter - verify before adding new ones |
| ASCII art rendering issues | Low | Low | Use simple block characters, test at 120x40 |
| Clock tick performance | Low | Low | 1-second tick is standard Bubble Tea pattern |
| Focus management complexity | Medium | Medium | Follow layout_gallery.go pattern which already demonstrates section-based focus |
| TuiNavigation API mismatch | Low | Medium | API already explored - function signature is known |

---

## Self-Review Ergebnisse

### 1. Vollstaendigkeit
- [x] App Shell (Header, Nav, Content, StatusBar) - covered in Phase 2
- [x] Welcome Screen content - covered in Phase 3
- [x] Screen switching with placeholders - covered in Phase 4
- [x] Exit confirmation modal - covered in Phase 5
- [x] Keyboard navigation (global + per-panel) - covered in Phase 2 + 4
- [x] Terminal resize handling - covered in Phase 2
- [x] Real clock time - covered in Phase 2
- [x] Entry point migration - covered in Phase 1
- [x] Demo rename - covered in Phase 1
- [x] Visual regression tests - covered in Phase 6
- [x] Minimum terminal size warning - covered in Phase 2

### 2. Konsistenz
- [x] Architecture follows DEC-001 through DEC-006
- [x] Uses existing components (TuiNavigation, TuiStatusBar, TuiModal)
- [x] All styling via Lip Gloss (Charm.land First rule)
- [x] Pattern consistent with existing layout_gallery.go demo

### 3. Risiken
- Header styles: Already exist in styles.go (StyleHeader, StyleHeaderTitle, StyleHeaderSubtitle)
- Focus management: Proven pattern from layout_gallery.go
- No service/domain layer needed for this sprint - keeps scope minimal

### 4. Alternativen
- **Considered:** Creating screens as interfaces vs concrete types
  - **Decision:** Concrete types - simpler, follows DEC-005, no need for abstraction yet
- **Considered:** Single file per screen vs model/view/update split
  - **Decision:** Single file for simple screens (welcome, placeholder), split when complex (future Build screen)

### 5. Komponenten-Verbindungen
- [x] JEDE neue Komponente hat mindestens eine Verbindung
- [x] JEDE Verbindung ist einer Story zugeordnet
- [x] Keine "verwaisten" Komponenten

---

## Minimalinvasiv-Optimierungen

### 1. Wiederverwendung

| Bestehend | Wiederverwendung |
|-----------|-----------------|
| `TuiNavigation()` | Direkt aufrufen mit nav items - kein Wrapper noetig |
| `TuiStatusBar()` | Direkt aufrufen mit config - kein Wrapper noetig |
| `TuiModal()` | Direkt aufrufen fuer Exit-Bestaetigungsdialog |
| `styles.go` tokens | Alle Farben, Borders, Typography bereits definiert |
| `layout_gallery.go` patterns | Focus management, section routing als Vorlage |

### 2. Aenderungsumfang

| Aenderung | Noetig? | Begruendung |
|-----------|---------|-------------|
| Neues `internal/app/` Package | JA | Architecture Decision DEC-002 erfordert root model |
| Neues `internal/ui/screens/` Package | JA | Architecture Decision DEC-005 erfordert screen models |
| Aendern von `styles.go` | WAHRSCHEINLICH NICHT | Header/footer styles bereits vorhanden (StyleHeader, StyleHeaderTitle, etc.) |
| Aendern bestehender Komponenten | NEIN | Alle APIs passen bereits |
| Neues `cmd/rfz-components-demo/` | JA | Demo muss umbenannt werden |

### 3. Feature-Preservation

- [x] Alle Requirements aus Clarification sind abgedeckt
- [x] Kein Feature wurde geopfert
- [x] Alle Akzeptanzkriterien bleiben erfuellbar
- [x] Bestehende Demo-Funktionalitaet bleibt erhalten (umbenannt)

---

## Story-Vorschlag (Ableitung aus Phasen)

| # | Story | Phase | Komplexitaet | Dateien |
|---|-------|-------|-------------|---------|
| 1 | Entry Point & Demo Rename | Phase 1 | XS | 2 files (main.go move + new main.go) |
| 2 | App Shell Model with Layout | Phase 2 | S | 2 files (app.go, messages.go) |
| 3 | Welcome Screen | Phase 3 | XS | 1 file (welcome.go) |
| 4 | Screen Switching & Navigation | Phase 4 | S | 1 file (placeholder.go) + updates to app.go |
| 5 | Exit Confirmation Modal | Phase 5 | XS | Update app.go |
| 6 | Visual Regression Tests | Phase 6 | S | 1 file (app_test.go) |

**Total:** 6 Stories, ~5-7 new files, 0 modified existing component files

---

*Implementation Plan ready for review.*
