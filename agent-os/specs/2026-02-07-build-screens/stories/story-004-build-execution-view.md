# Build Execution View

> Story ID: BUILD-004
> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Frontend
**Estimated Effort**: M
**Dependencies**: BUILD-001, BUILD-003

---

## Feature

```gherkin
Feature: Build-Ausfuehrung ueberwachen
  Als RFZ-Entwickler
  moechte ich den Fortschritt meiner Builds in Echtzeit sehen,
  damit ich den Status jeder Komponente verfolgen und bei Problemen schnell reagieren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Build-Ausfuehrung startet

```gherkin
Scenario: Build-Ansicht zeigt gestarteten Build
  Given ich habe einen Build mit 3 Komponenten gestartet
  When die Build-Ausfuehrungsansicht angezeigt wird
  Then sehe ich den Maven-Befehl oben als Referenz
  And sehe ich eine Tabelle mit 3 Komponenten
  And der Gesamtfortschrittsbalken zeigt 0%
  And die Statusanzeige zeigt "Running: 0, Success: 0, Failed: 0, Pending: 3"
```

### Szenario 2: Komponente durchlaeuft Build-Phasen

```gherkin
Scenario: Komponente wechselt durch Build-Phasen
  Given ein Build laeuft fuer "audiocon"
  When die Simulation fortschreitet
  Then wechselt der Status von "Pending" zu "Compiling"
  And dann zu "Testing"
  And dann zu "Packaging"
  And dann zu "Installing"
  And schliesslich zu "Done"
  And die verstrichene Zeit wird als "MM:SS" angezeigt
```

### Szenario 3: Fortschrittsanzeige aktualisiert sich

```gherkin
Scenario: Gesamtfortschritt aktualisiert sich bei Abschluss
  Given ein Build mit 3 Komponenten laeuft
  And 1 Komponente ist erfolgreich abgeschlossen
  When ich die Fortschrittsanzeige betrachte
  Then zeigt der Gesamtfortschrittsbalken ~33%
  And die Statusanzeige zeigt "Running: 2, Success: 1, Failed: 0, Pending: 0"
```

### Szenario 4: Build komplett erfolgreich

```gherkin
Scenario: Alle Komponenten erfolgreich gebaut
  Given ein Build mit 3 Komponenten laeuft
  When alle 3 Komponenten erfolgreich abgeschlossen sind
  Then zeigt der Gesamtfortschrittsbalken 100%
  And die Statusanzeige zeigt "Running: 0, Success: 3, Failed: 0, Pending: 0"
  And die Aktionen zeigen "New Build" und "View Logs"
```

### Szenario 5: Navigation zwischen Komponenten

```gherkin
Scenario: Mit Pfeiltasten zwischen Komponenten navigieren
  Given die Build-Ausfuehrungsansicht ist aktiv
  And der Cursor steht auf der ersten Komponente
  When ich die Pfeil-Runter-Taste druecke
  Then wechselt der Cursor zur naechsten Komponente
  And die fokussierte Zeile ist visuell hervorgehoben
```

### Edge Cases & Fehlerszenarien

### Szenario 6: Komponente schlaegt fehl

```gherkin
Scenario: Eine Komponente schlaegt beim Bauen fehl
  Given ein Build laeuft fuer "energierechnung"
  When die Komponente in der Testing-Phase fehlschlaegt
  Then zeigt der Status "Failed" in Rot
  And die Phase zeigt "Testing" (wo der Fehler auftrat)
  And die anderen Komponenten bauen weiter
  And die Statusanzeige zeigt "Failed: 1"
```

### Szenario 7: Build abbrechen

```gherkin
Scenario: Laufenden Build abbrechen
  Given ein Build mit 3 Komponenten laeuft
  And 2 Komponenten sind noch aktiv
  When ich Ctrl+C druecke oder "Cancel Build" auswaehle
  Then stoppen alle laufenden Komponenten sofort
  And der Status aller laufenden Komponenten wechselt zu "Cancelled"
  And die Aktionen zeigen "New Build"
```

### Szenario 8: Zurueck zur Komponentenauswahl

```gherkin
Scenario: Nach Build-Abschluss neue Auswahl treffen
  Given der Build ist abgeschlossen (alle Komponenten fertig)
  When ich "New Build" auswaehle
  Then wechsle ich zurueck zur Komponentenauswahl
  And meine vorherige Auswahl ist zurueckgesetzt
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/screens/build/execution.go
- [ ] FILE_EXISTS: internal/ui/screens/build/simulator.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/ui/screens/build/execution.go enthaelt "TuiStatus"
- [ ] CONTAINS: internal/ui/screens/build/execution.go enthaelt "TuiProgress"
- [ ] CONTAINS: internal/ui/screens/build/simulator.go enthaelt "BuildPhase"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `cd /Users/lix/xapps/rfz-tui && go build ./internal/ui/screens/build/...`
- [ ] LINT_PASS: `cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/ui/screens/build/...`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> Dieser Abschnitt wird vom Architect ausgefuellt (Step 3)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und pruefbar
- [x] Business Value verstanden

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhaengigkeiten identifiziert
- [x] Betroffene Komponenten bekannt
- [x] Erforderliche MCP Tools dokumentiert (falls zutreffend)
- [x] Story ist angemessen geschaetzt (max 5 Dateien, 400 LOC)

#### Full-Stack Konsistenz
- [x] Alle betroffenen Layer identifiziert
- [x] Integration Type bestimmt
- [x] Kritische Integration Points dokumentiert (wenn Full-stack)
- [x] Handover-Dokumente definiert (bei Multi-Layer)

### DoD (Definition of Done) - Vom Architect

- [ ] Execution view renders component table with per-row status, spinner, progress, time
- [ ] Build simulator sends timed BuildTickMsg and BuildPhaseMsg
- [ ] Overall progress bar tracks completion percentage
- [ ] Status counters display running/success/failed/pending counts
- [ ] Per-component build phases transition correctly (Pending -> Compiling -> Testing -> Packaging -> Installing -> Done)
- [ ] Random failure simulation works for at least one component
- [ ] Keyboard handling for cancellation (Escape, Ctrl+C)
- [ ] Actions section shows appropriate buttons based on build state
- [ ] All acceptance criteria met
- [ ] Tests written and passing
- [ ] No linting errors
- [ ] Completion Check commands successful

### Betroffene Layer & Komponenten

| Layer | Integration Type | Affected Components/Files |
|-------|------------------|--------------------------|
| UI Screen | Extends BUILD-002 | `internal/ui/screens/build/execution.go`, `internal/ui/screens/build/simulator.go`, updates to `model.go` and `update.go` |
| Domain | Consumes | `internal/domain/buildconfig.go` (BuildPhase), `internal/domain/component.go` |
| App Messages | New | `internal/app/messages.go` (BuildTickMsg, BuildPhaseMsg, BuildCompleteMsg) |
| UI Components | Reuses | `TuiStatus`, `TuiProgress`, `TuiSpinner`, `TuiButton`, Lip Gloss layout |
| Integration | Frontend-only | TUI application, no backend/database |

### Technical Details

**WAS** (What to create):
- Execution view rendering with component table (Status, Component, Phase, Progress, Time columns)
- Per-component build state tracking (current phase, elapsed time, progress percentage)
- Build simulator with timed phase transitions
- BuildTickMsg, BuildPhaseMsg, BuildCompleteMsg message types
- Overall progress calculation across all components
- Status counter aggregation (running/success/failed/pending)
- Actions section with context-appropriate buttons (Cancel Build / New Build, View Logs)
- Custom table rendering using Lip Gloss layout (bubbles/table doesn't support inline widgets)

**WIE** (How to implement):
- Execution view uses custom table rendering (not bubbles/table) to embed TuiStatus, TuiSpinner, TuiProgress per row
- Build simulator goroutine sends BuildTickMsg every 100ms, BuildPhaseMsg on phase transitions
- Each component has componentBuildState struct tracking current phase, start time, progress
- Phase transition logic: random duration per phase, 20 percent chance of failure in Testing phase
- Overall progress bar calculates (completed components / total components) percentage
- Status counters aggregate across all componentBuildState entries
- All rendering via Lip Gloss Place/JoinHorizontal/JoinVertical
- Phase colors mapped via components.ColorCyan (running), ColorGreen (success), ColorDestructive (failed)
- Cancellation sets all running components to Done/Failed state, stops simulator

**WO** (Where to implement):
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/execution.go` - Execution view rendering, table layout, status counters
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/simulator.go` - Build simulator with timed phase transitions, goroutine management
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/model.go` - Add componentBuildState slice, simulator control fields
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/update.go` - Handle BuildTickMsg, BuildPhaseMsg, BuildCompleteMsg
- `/Users/lix/xapps/rfz-tui/internal/app/messages.go` - Add BuildTickMsg, BuildPhaseMsg, BuildCompleteMsg types

**WER** (Who implements):
- Generic fullstack developer (no specialized domain agents available)

**Abhaengigkeiten** (Dependencies):
- BUILD-001 (needs BuildPhase enum, Component type)
- BUILD-003 (needs BuildConfig to know what components to execute)

**Geschaetzte Komplexitaet** (Estimated Complexity):
- M (Medium) - 5 files modified/created, ~300 LOC, complex table rendering and simulation timing

**NOTE**: This story is sized M which is at the complexity limit. If implementation proves difficult, consider splitting simulator into separate story.

**Relevante Skills** (Required Skills):

| Skill | Reason |
|-------|--------|
| go-bubbletea | Message-based state updates, Cmd/Msg patterns |
| go-concurrency | Goroutine management for build simulator |
| charm-lipgloss | Custom table layout with embedded widgets |
| domain-rfz-cli | Understanding TuiStatus, TuiProgress, TuiSpinner components |

### Completion Check

```bash
# Build build screen package
cd /Users/lix/xapps/rfz-tui && go build ./internal/ui/screens/build/...

# Lint build screen package
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/ui/screens/build/...
```
