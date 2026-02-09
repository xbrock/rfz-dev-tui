# Implementation Plan - Layout Matching Design

**Created:** 2026-02-09
**Status:** Draft
**Spec:** 2026-02-09-bugfix-layout-matching-design
**Input:** Approved requirements-clarification.md

---

## Executive Summary

Fix visual discrepancies across all implemented TUI screens to match the approved design prototype. The changes are purely cosmetic/layout fixes touching the component library (`internal/ui/components/`), the app shell (`internal/app/app.go`), and screen implementations (`internal/ui/screens/`). No functional/business logic changes needed.

## Architektur-Entscheidungen

1. **Component Library First** - Fix shared components (styles, navigation, statusbar, list, progress) BEFORE fixing screens, since screens depend on these
2. **Minimal Changes** - Only modify what's needed to match the prototype; preserve all existing functionality
3. **Charm.land Compliance** - Continue using Lip Gloss for ALL styling; no custom ANSI codes or manual border drawing
4. **Width Calculation Fix** - Centralize border overflow fix in the app shell rather than patching each screen individually
5. **Tree Component Reuse** - Use existing `TuiTree` component patterns for nav hints and build execution tree rendering where possible

## Komponenten-Uebersicht

### Bestehende Komponenten (zu aendern)

| Komponente | Datei | Aenderung |
|------------|-------|-----------|
| styles.go | `internal/ui/components/styles.go` | Add new color tokens, update nav item styles for light blue active bg, add status bar styles |
| navigation.go | `internal/ui/components/navigation.go` | Fix active/select state priority, shortcut right-alignment, tree-style hints, content-based height |
| statusbar.go | `internal/ui/components/statusbar.go` | 3-badge system, pipe-separated hints, gray bg full width, q Quit right |
| list.go | `internal/ui/components/list.go` | Change multi-select symbols from checkbox to circle (only for component list mode) |
| keyhints.go | `internal/ui/components/keyhints.go` | Add tree-style rendering mode |
| progress.go | `internal/ui/components/progress.go` | Add braille block style, add filled/empty block style for overall |
| App Shell | `internal/app/app.go` | Fix header layout (red line top, subtitle+time same line), fix width calculations to prevent border overflow, nav content-based height |
| Welcome Screen | `internal/ui/screens/welcome/welcome.go` | Logo colors, subtitle white, braille line, 3 badges, tree hints |
| Build Selection | `internal/ui/screens/build/selection.go` | Hint overflow fix, action button colors, legend icons |
| Build Config | `internal/ui/screens/build/config.go` | Shortcut label colors, hint separators |
| Build Execution | `internal/ui/screens/build/execution.go` | Full-width component list, tree icons, progress bar styles, badge cleanup |

### Neue Komponenten

Keine neuen Komponenten noetig - alle Aenderungen sind an bestehenden Komponenten.

## Komponenten-Verbindungen

| Source | Target | Verbindungsart | Zustaendige Story |
|--------|--------|----------------|-------------------|
| styles.go (new tokens) | navigation.go | Import: neue Farb-Tokens fuer active state | Story 1 + 2 |
| styles.go (new tokens) | statusbar.go | Import: neue Styles fuer 3-Badge System | Story 1 + 3 |
| navigation.go (tree hints) | keyhints.go | Nutzt tree-style Rendering | Story 2 |
| app.go (width fix) | all screens | Korrekte Width wird an alle Screens propagiert | Story 8 |
| list.go (circle symbols) | build/selection.go | Neuer Modus fuer circle checkboxes | Story 5 |
| progress.go (braille style) | build/execution.go | Neuer Progress-Stil | Story 7 |

## Umsetzungsphasen

### Phase 1: Foundation Fixes (Stories 1, 8)
- Fix styles/color tokens
- Fix general border overflow issue in app shell
- These are prerequisites for all other stories

### Phase 2: Shell Components (Stories 2, 3)
- Fix navigation sidebar (depends on styles)
- Fix status bar (depends on styles)

### Phase 3: Welcome Screen (Story 4)
- Fix welcome screen layout (depends on styles + status bar)

### Phase 4: Build Screen Fixes (Stories 5, 6, 7)
- Fix build components list
- Fix config modal
- Fix build execution view
- These can run in parallel

## Abhaengigkeiten

```
Story 1 (Styles) ──┬──> Story 2 (Navigation)
                    ├──> Story 3 (Status Bar)
Story 8 (Borders) ──┤
                    ├──> Story 4 (Welcome) [after 2, 3]
                    ├──> Story 5 (Build Components)
                    ├──> Story 6 (Config Modal)
                    └──> Story 7 (Build Execution)
```

## Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| Width calculations break at edge terminal sizes | Medium | Medium | Test at 80x24 (min), 120x40 (canonical), and 200x50 (large) |
| Tree hints rendering differs from prototype | Low | Low | Compare against prototype screenshots pixel-by-pixel |
| Braille characters not supported in all terminals | Low | Medium | Use standard block characters as fallback pattern |
| Progress bar styles conflict with existing TuiProgress | Low | Low | Add new style mode, don't change default behavior |

---

## Self-Review Ergebnisse

### 1. VOLLSTAENDIGKEIT
- Alle 8 Anforderungsbereiche aus der Clarification sind abgedeckt (Title Bar, Nav, Status Bar, Welcome, Build Components, Config Modal, Build Execution, Border Overflow)
- Jede Anforderung ist einer konkreten Story zugeordnet
- Pre-Implementation Screenshot Review ist als Pflicht in jeder Story vorgesehen

### 2. KONSISTENZ
- Alle Aenderungen nutzen Lip Gloss (kein custom ANSI)
- Farb-Tokens werden zentral in styles.go definiert und von allen Komponenten importiert
- Width-Fix ist zentral im App Shell, nicht pro-Screen dupliziert

### 3. RISIKEN
- Border overflow ist ein fundamentaler Layout-Bug der frueh gefixt werden muss (Story 8 in Phase 1)
- Keine Breaking Changes an bestehenden Component APIs - nur Erweiterungen

### 4. ALTERNATIVEN
- Statt circle-symbols im List-Component koennte man einen neuen Component erstellen -> Nicht noetig, List hat bereits `ListSelectMode` parameter
- Statt tree-style hints in navigation koennte man `TuiTree` component direkt nutzen -> Overkill, einfache String-Formatierung reicht fuer statische Hint-Listen

### 5. KOMPONENTEN-VERBINDUNGEN
- Alle Verbindungen validiert: styles.go -> navigation.go, statusbar.go, list.go, progress.go
- Keine verwaisten Komponenten
- Jede Verbindung hat eine zustaendige Story

---

## Minimalinvasiv-Optimierungen

### 1. WIEDERVERWENDUNG
- `TuiListItem` hat bereits `ListMultiSelect` und `ListSingleSelect` Modes -> Neuen Mode `ListCircleSelect` hinzufuegen oder das Symbol im Multi-Select-Modus konfigurierbar machen
- `TuiTree` Component existiert bereits fuer hierarchische Darstellung -> Patterns fuer Tree-Icons in Build Execution wiederverwendbar
- `TuiKeyHints` existiert bereits -> Tree-Style als zusaetzlichen Rendering-Modus hinzufuegen statt neues Component
- `TuiProgress` existiert als Wrapper um Bubbles Progress -> Neuen Style-Modus fuer Braille-Blocks hinzufuegen

### 2. AENDERUNGSUMFANG
- styles.go: ~20 neue Zeilen (Farb-Tokens, Style-Varianten)
- navigation.go: ~40 geaenderte Zeilen (State-Prioritaet, Shortcut-Alignment)
- statusbar.go: ~30 geaenderte Zeilen (3-Badge Layout, Pipe-Separatoren)
- app.go: ~20 geaenderte Zeilen (Header Layout, Width Fix)
- Screens: je ~20-40 geaenderte Zeilen
- **Geschaetzter Gesamtumfang: ~250-350 LOC geaendert**

### 3. FEATURE-PRESERVATION CHECKLISTE
- [x] Alle Requirements aus Clarification sind abgedeckt
- [x] Kein Feature wird geopfert
- [x] Alle Akzeptanzkriterien bleiben erfuellbar
- [x] Bestehende Funktionalitaet (Build-Flow, Navigation, Keyboard Shortcuts) bleibt erhalten
- [x] Bestehende Tests bleiben kompatibel (Golden Files muessen aktualisiert werden)

---

*Implementation Plan v1.0 - Ready for Review*
