# App Integration & Screen Transitions

> Story ID: BUILD-005
> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: BUILD-002, BUILD-003, BUILD-004

**Integration:** `app.Model` -> `build.Model`

---

## Feature

```gherkin
Feature: Build-Screens in die Anwendung integrieren
  Als RFZ-Entwickler
  moechte ich nahtlos zwischen Navigation und Build-Screens wechseln koennen,
  damit der Build-Workflow sich natuerlich in die Anwendung einfuegt.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Zum Build-Screen navigieren

```gherkin
Scenario: Build Components Screen ueber Navigation oeffnen
  Given ich bin auf dem Welcome-Screen
  When ich die Taste "1" druecke
  Then sehe ich den Build Components Screen mit der Komponentenliste
  And die Navigation zeigt "Build Components" als aktiv
  And die Statusleiste zeigt "SELECT" als Modus
```

### Szenario 2: Statusleiste aktualisiert sich pro Build-Phase

```gherkin
Scenario Outline: Statusleiste spiegelt den aktuellen Build-Zustand wider
  Given ich bin im Build-Workflow
  When ich mich in der Phase "<phase>" befinde
  Then zeigt die Statusleiste den Modus "<modus>"
  And zeigt den Kontext "<kontext>"

  Examples:
    | phase              | modus   | kontext          |
    | Komponentenauswahl | SELECT  | boss             |
    | Build laeuft       | BUILD   | RUNNING          |
    | Build fertig       | BUILD   | COMPLETE         |
```

### Szenario 3: Fokus-Wechsel zwischen Navigation und Build

```gherkin
Scenario: Tab wechselt den Fokus
  Given ich bin auf dem Build Components Screen
  And der Fokus liegt auf dem Inhalt (Komponentenliste)
  When ich Tab druecke
  Then wechselt der Fokus zur Navigation
  And die Navigationsbox hat einen hervorgehobenen Rahmen
```

### Szenario 4: Escape zurueck zum Welcome-Screen

```gherkin
Scenario: Escape bringt zurueck zum Welcome-Screen
  Given ich bin auf dem Build Components Screen
  And kein Modal ist geoeffnet
  When ich Escape druecke
  Then sehe ich den Welcome-Screen
  And keine Navigation ist als aktiv markiert
```

### Szenario 5: Golden-File-Tests fuer alle Zustaende

```gherkin
Scenario: Visuelle Regression fuer Build-Zustaende
  Given die Anwendung laeuft im Test-Modus (120x40 Terminal)
  When ich alle Build-Screen-Zustaende durchlaufe
  Then stimmt jeder Zustand mit dem gespeicherten Golden File ueberein
  And kein visueller Unterschied wird erkannt
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Terminal-Groessenaenderung waehrend Build
  Given ein Build laeuft
  When die Terminalgroesse sich aendert
  Then passen sich alle Build-Views an die neue Groesse an
  And der Fortschrittsbalken hat die korrekte Breite
  And kein UI-Element wird abgeschnitten
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/screens/build/update.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/app/app.go enthaelt "build.Model" (nicht mehr "phBuild placeholder.Model")
- [ ] NOT_CONTAINS: internal/app/app.go enthaelt NICHT "phBuild"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `cd /Users/lix/xapps/rfz-tui && go build ./...`
- [ ] TEST_PASS: `cd /Users/lix/xapps/rfz-tui && go test ./...`
- [ ] LINT_PASS: `cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...`

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

- [ ] app.Model replaces placeholder.Model with build.Model
- [ ] App delegates Update/View to build.Model
- [ ] Focus delegation working (Tab switches between nav and build content)
- [ ] Status bar updates correctly per build phase (SELECT/BUILD/COMPLETE modes)
- [ ] Navigation highlights Build Components when active
- [ ] Escape returns to Welcome screen
- [ ] Size changes propagate to build.Model.SetSize()
- [ ] Golden file tests written for all distinct build UI states
- [ ] All acceptance criteria met
- [ ] Tests written and passing
- [ ] No linting errors
- [ ] Completion Check commands successful

### Betroffene Layer & Komponenten

| Layer | Integration Type | Affected Components/Files |
|-------|------------------|--------------------------|
| App | Integration | `internal/app/app.go` - replace phBuild with build.Model |
| App | Integration | `internal/app/messages.go` - already updated in BUILD-004 |
| UI Components | Integration | `internal/ui/components/statusbar.go` - may need mode updates |
| Tests | New | Golden file tests for build screen states |
| Integration | Frontend-only | TUI application, no backend/database |

### Technical Details

**WAS** (What to integrate):
- Replace placeholder.Model field with build.Model in app.Model
- Add build screen to app.Update switch cases for message delegation
- Add build screen to app.View switch cases for rendering
- Update status bar context based on build phase
- Update navigation active state when on build screen
- Wire focus delegation (Tab toggles between nav and build)
- Create golden file tests for build screen UI states

**WIE** (How to integrate):
- Change app.Model.phBuild from placeholder.Model to build.Model
- Initialize build.Model in app.New() with component provider
- In app.Update: when screen==screenBuild, delegate to build.Update()
- In app.View: when screen==screenBuild, render build.View()
- Update contentWidth/contentHeight calculations to account for build content
- Status bar mode field changes based on build.buildPhase (selecting=SELECT, executing=BUILD, completed=COMPLETE)
- Focus delegation already exists in app, just ensure build.Update handles focus correctly
- Golden file tests use teatest with deterministic model states (no real timers)

**WO** (Where to integrate):
- `/Users/lix/xapps/rfz-tui/internal/app/app.go` - Replace phBuild field, add switch cases, update New()
- `/Users/lix/xapps/rfz-tui/internal/app/app_test.go` - Add golden file tests for build screens (if not exists, create)
- `/Users/lix/xapps/rfz-tui/internal/ui/components/statusbar.go` - May need to support additional mode values

**WER** (Who implements):
- Generic fullstack developer (no specialized domain agents available)

**Abhaengigkeiten** (Dependencies):
- BUILD-002 (needs build.Model)
- BUILD-003 (needs config modal complete)
- BUILD-004 (needs execution view complete)

**Geschaetzte Komplexitaet** (Estimated Complexity):
- S (Small) - 2-3 files modified, ~50 LOC changes, mostly wiring existing pieces

**Relevante Skills** (Required Skills):

| Skill | Reason |
|-------|--------|
| go-bubbletea | Understanding Model composition and delegation |
| go-testing | Writing golden file tests with teatest |
| domain-rfz-cli | Understanding app.Model structure and patterns |

### Completion Check

```bash
# Build entire project
cd /Users/lix/xapps/rfz-tui && go build ./...

# Run all tests
cd /Users/lix/xapps/rfz-tui && go test ./...

# Lint entire project
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
