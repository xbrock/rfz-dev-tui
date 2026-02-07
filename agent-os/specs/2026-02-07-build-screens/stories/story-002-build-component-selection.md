# Build Component Selection Screen

> Story ID: BUILD-002
> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: BUILD-001

---

## Feature

```gherkin
Feature: Build-Komponenten auswaehlen
  Als RFZ-Entwickler
  moechte ich Komponenten aus einer Liste auswaehlen koennen,
  damit ich gezielt bestimmte RFZ-Komponenten bauen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Komponentenliste anzeigen

```gherkin
Scenario: Alle Komponenten werden in der Build-Ansicht angezeigt
  Given ich bin auf dem Build Components Screen
  When die Ansicht geladen wird
  Then sehe ich eine Liste mit 13 RFZ-Komponenten
  And jede Komponente zeigt ein Typ-Badge (Core, Simulation oder Standalone)
  And der Zaehler zeigt "0/13 selected"
```

### Szenario 2: Einzelne Komponente auswaehlen

```gherkin
Scenario: Komponente per Leertaste auswaehlen
  Given ich bin auf dem Build Components Screen
  And der Cursor steht auf "boss"
  When ich die Leertaste druecke
  Then ist "boss" als ausgewaehlt markiert [x]
  And der Zaehler zeigt "1/13 selected"
```

### Szenario 3: Alle Komponenten auswaehlen

```gherkin
Scenario: Alle Komponenten mit Tastenkuerzel auswaehlen
  Given ich bin auf dem Build Components Screen
  And keine Komponente ist ausgewaehlt
  When ich "a" druecke
  Then sind alle 13 Komponenten als ausgewaehlt markiert
  And der Zaehler zeigt "13/13 selected"
```

### Szenario 4: Auswahl komplett loeschen

```gherkin
Scenario: Alle Auswahlen mit Tastenkuerzel loeschen
  Given ich bin auf dem Build Components Screen
  And 3 Komponenten sind ausgewaehlt
  When ich "n" druecke
  Then ist keine Komponente ausgewaehlt
  And der Zaehler zeigt "0/13 selected"
```

### Szenario 5: Build starten mit Auswahl

```gherkin
Scenario: Build-Konfiguration oeffnen nach Komponentenauswahl
  Given ich bin auf dem Build Components Screen
  And ich habe "audiocon", "traktion" und "signalsteuerung" ausgewaehlt
  When ich Enter druecke (Build Selected)
  Then oeffnet sich das Build-Konfigurations-Modal
  And das Modal zeigt "Building 3 components: audiocon, traktion, signalsteuerung"
```

### Szenario 6: Navigation in der Komponentenliste

```gherkin
Scenario: Mit Pfeiltasten durch die Liste navigieren
  Given ich bin auf dem Build Components Screen
  And der Cursor steht auf dem ersten Eintrag "boss"
  When ich die Pfeil-Runter-Taste druecke
  Then steht der Cursor auf "fistiv"
  And der aktuelle Eintrag ist visuell hervorgehoben
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Build starten ohne Auswahl
  Given ich bin auf dem Build Components Screen
  And keine Komponente ist ausgewaehlt
  When ich Enter druecke (Build Selected)
  Then passiert nichts (Aktion wird ignoriert)
  And ich bleibe auf dem Build Components Screen
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/screens/build/model.go
- [ ] FILE_EXISTS: internal/ui/screens/build/selection.go
- [ ] FILE_EXISTS: internal/ui/screens/build/view.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/ui/screens/build/model.go enthaelt "type Model struct"
- [ ] CONTAINS: internal/ui/screens/build/selection.go enthaelt "TuiList"

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

- [ ] Build screen Model follows Bubble Tea sub-model pattern (New/Init/Update/View)
- [ ] Component selection uses existing TuiList with multi-select mode
- [ ] Keyboard navigation implemented (up/down, space, a, n, enter)
- [ ] Actions section renders with TuiButton components
- [ ] State machine tracks selection vs configuring vs executing phases
- [ ] All acceptance criteria met
- [ ] Tests written and passing
- [ ] No linting errors
- [ ] Completion Check commands successful

### Betroffene Layer & Komponenten

| Layer | Integration Type | Affected Components/Files |
|-------|------------------|--------------------------|
| UI Screen | New | `internal/ui/screens/build/model.go`, `internal/ui/screens/build/view.go`, `internal/ui/screens/build/selection.go` |
| Domain | Consumes | `internal/domain/component.go` (ComponentProvider) |
| UI Components | Reuses | `TuiList`, `TuiButton`, `TuiBox`, `TuiKeyHints` from `internal/ui/components/` |
| Integration | Frontend-only | TUI application, no backend/database |

### Technical Details

**WAS** (What to create):
- Build screen Model struct with state machine (buildPhase enum: selecting/configuring/executing/completed)
- Fields for component list (TuiListItem slice), cursor position, selected indices
- ComponentProvider instance for loading component data
- Selection view rendering function using TuiList with badges
- Actions section with buttons (Build Selected, Select All, Deselect All, Cancel)
- Keyboard handling for navigation (up/down), selection (space), batch operations (a/n), and build trigger (enter)

**WIE** (How to implement):
- Follow Bubble Tea sub-model pattern exactly like welcome.Model (New/SetSize/Init/Update/View)
- Build screen owns its internal state machine via buildPhase field
- Use existing TuiList component with ListMultiSelect mode
- Convert domain.Component slice to TuiListItem slice with type badges (Core=cyan, Simulation=yellow, Standalone=green)
- Selection view composes TuiList + actions section + key hints using Lip Gloss layout
- All styling via components.styles.go (no manual padding or borders)
- State transitions: selecting -> (enter with selection) -> configuring
- Parent app delegates Update/View calls to build.Model

**WO** (Where to implement):
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/model.go` - Model struct, New(), SetSize(), Init(), buildPhase enum
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/view.go` - View() dispatcher based on buildPhase
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/selection.go` - Selection view rendering, converts components to TuiListItems
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/update.go` - Update() with keyboard handling for selection phase

**WER** (Who implements):
- Generic fullstack developer (no specialized domain agents available)

**Abhaengigkeiten** (Dependencies):
- BUILD-001 (needs Component type and ComponentProvider)

**Geschaetzte Komplexitaet** (Estimated Complexity):
- S (Small) - 4 files, ~250 LOC, straightforward Bubble Tea sub-model

**Relevante Skills** (Required Skills):

| Skill | Reason |
|-------|--------|
| go-bubbletea | Bubble Tea Model/Update/View pattern |
| charm-lipgloss | Layout with JoinVertical/Place |
| domain-rfz-cli | Understanding TUI component library |

### Completion Check

```bash
# Build build screen package
cd /Users/lix/xapps/rfz-tui && go build ./internal/ui/screens/build/...

# Lint build screen package
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/ui/screens/build/...
```
