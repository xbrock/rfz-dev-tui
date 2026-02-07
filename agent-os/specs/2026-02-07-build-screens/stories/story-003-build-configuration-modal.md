# Build Configuration Modal

> Story ID: BUILD-003
> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: BUILD-002

---

## Feature

```gherkin
Feature: Build-Konfiguration einstellen
  Als RFZ-Entwickler
  moechte ich vor dem Build Maven-Optionen konfigurieren koennen,
  damit ich den Build-Befehl an meine Beduerfnisse anpassen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Modal mit Standardwerten oeffnen

```gherkin
Scenario: Build-Konfigurationsmodal zeigt Standardwerte
  Given ich habe 3 Komponenten ausgewaehlt
  When das Build-Konfigurationsmodal sich oeffnet
  Then sehe ich "Building 3 components: audiocon, traktion, signalsteuerung"
  And "clean install" ist als Maven-Ziel vorausgewaehlt
  And "target_env_dev" ist als Profil aktiviert
  And Port 11090 ist vorausgewaehlt
  And "Skip Tests" ist aktiviert
  And die Befehlsvorschau zeigt "$ mvn clean-install -Ptarget_env_dev,use_traktion_11090 -DskipTests"
```

### Szenario 2: Maven-Ziel aendern

```gherkin
Scenario Outline: Maven-Ziel auswaehlen aktualisiert Befehlsvorschau
  Given das Build-Konfigurationsmodal ist geoeffnet
  When ich das Maven-Ziel "<ziel>" auswaehle
  Then zeigt die Befehlsvorschau "$ mvn <befehl>"

  Examples:
    | ziel           | befehl                                                        |
    | clean          | clean -Ptarget_env_dev,use_traktion_11090 -DskipTests         |
    | install        | install -Ptarget_env_dev,use_traktion_11090 -DskipTests       |
    | package        | package -Ptarget_env_dev,use_traktion_11090 -DskipTests       |
    | clean install  | clean install -Ptarget_env_dev,use_traktion_11090 -DskipTests |
```

### Szenario 3: Maven-Profile umschalten

```gherkin
Scenario: Zusaetzliches Profil aktivieren
  Given das Build-Konfigurationsmodal ist geoeffnet
  And nur "target_env_dev" ist aktiviert
  When ich "generate_local_config_files" per Leertaste aktiviere
  Then sind beide Profile aktiviert
  And die Befehlsvorschau zeigt "-Ptarget_env_dev,generate_local_config_files,use_traktion_11090"
```

### Szenario 4: Traktion-Port aendern

```gherkin
Scenario: Port von 11090 auf 11091 aendern
  Given das Build-Konfigurationsmodal ist geoeffnet
  And Port 11090 ist ausgewaehlt
  When ich Port 11091 auswaehle
  Then zeigt die Befehlsvorschau "use_traktion_11091" statt "use_traktion_11090"
```

### Szenario 5: Skip Tests deaktivieren

```gherkin
Scenario: Tests nicht ueberspringen
  Given das Build-Konfigurationsmodal ist geoeffnet
  And "Skip Tests" ist aktiviert
  When ich "Skip Tests" per Leertaste deaktiviere
  Then erscheint "-DskipTests" nicht mehr in der Befehlsvorschau
```

### Szenario 6: Tab-Navigation zwischen Sektionen

```gherkin
Scenario: Mit Tab zwischen Modal-Sektionen navigieren
  Given das Build-Konfigurationsmodal ist geoeffnet
  And der Fokus liegt auf "Maven Goal"
  When ich Tab druecke
  Then wechselt der Fokus zu "Maven Profiles"
  And die fokussierte Sektion ist visuell hervorgehoben
```

### Szenario 7: Build starten

```gherkin
Scenario: Build aus dem Konfigurationsmodal starten
  Given das Build-Konfigurationsmodal ist geoeffnet
  And ich habe meine Einstellungen vorgenommen
  When ich "Start Build" auswaehle (Enter)
  Then schliesst sich das Modal
  And die Build-Ausfuehrungsansicht wird angezeigt
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Modal abbrechen kehrt zur Auswahl zurueck
  Given das Build-Konfigurationsmodal ist geoeffnet
  When ich Escape druecke
  Then schliesst sich das Modal
  And ich bin zurueck auf dem Build Components Screen
  And meine Komponentenauswahl ist erhalten
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/screens/build/config.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/ui/screens/build/config.go enthaelt "TuiModal"
- [ ] CONTAINS: internal/ui/screens/build/config.go enthaelt "TuiRadio"
- [ ] CONTAINS: internal/ui/screens/build/config.go enthaelt "TuiCheckbox"

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

- [ ] Config modal renders inside existing TuiModal component
- [ ] Modal content composed from TuiRadio (Goal, Port) and TuiCheckbox (Profiles, SkipTests)
- [ ] Sectional focus navigation with Tab between 5 sections
- [ ] Command preview updates dynamically based on selections
- [ ] Modal state tracked in build.Model with configuring phase
- [ ] Keyboard handling for Escape (cancel), Enter (start build)
- [ ] All acceptance criteria met
- [ ] Tests written and passing
- [ ] No linting errors
- [ ] Completion Check commands successful

### Betroffene Layer & Komponenten

| Layer | Integration Type | Affected Components/Files |
|-------|------------------|--------------------------|
| UI Screen | Extends BUILD-002 | `internal/ui/screens/build/config.go`, updates to `model.go` and `update.go` |
| Domain | Consumes | `internal/domain/buildconfig.go` (BuildConfig, MavenGoal) |
| UI Components | Reuses | `TuiModal`, `TuiRadio`, `TuiCheckbox`, `TuiButton` from `internal/ui/components/` |
| Integration | Frontend-only | TUI application, no backend/database |

### Technical Details

**WAS** (What to create):
- Config modal rendering function with 5 sections (Header, Maven Goal, Maven Profiles, Traktion Port, Options)
- BuildConfig state in build.Model to track user selections
- Section focus state to enable Tab navigation between form sections
- Command preview that reactively shows generated Maven command
- Modal buttons (Start Build, Cancel)
- Keyboard handling for modal-specific keys (Tab, Escape, Enter)

**WIE** (How to implement):
- Build modal content as single string by composing TuiRadio/TuiCheckbox sections
- Pass composed content to TuiModal.Content field
- TuiModal handles double border, backdrop, centering automatically
- Track sectionFocus index in build.Model (0=Goal, 1=Profiles, 2=Port, 3=Options, 4=Buttons)
- Tab key cycles sectionFocus, dispatches to appropriate form component
- Each form section uses existing TuiRadio or TuiCheckbox renderers
- Command preview built via BuildConfig.ToCommand() called on every state change
- Modal opens when buildPhase transitions from selecting to configuring
- Modal closes (phase -> executing) on Enter, reverts (phase -> selecting) on Escape

**WO** (Where to implement):
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/config.go` - Modal rendering, section composition, command preview
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/model.go` - Add BuildConfig field, sectionFocus field
- `/Users/lix/xapps/rfz-tui/internal/ui/screens/build/update.go` - Add modal keyboard handlers (Tab, Escape, Enter when configuring)

**WER** (Who implements):
- Generic fullstack developer (no specialized domain agents available)

**Abhaengigkeiten** (Dependencies):
- BUILD-002 (needs build.Model with state machine)

**Geschaetzte Komplexitaet** (Estimated Complexity):
- S (Small) - 3 files modified/created, ~200 LOC, straightforward form composition

**Relevante Skills** (Required Skills):

| Skill | Reason |
|-------|--------|
| go-bubbletea | State management for modal focus |
| charm-lipgloss | Modal content composition with JoinVertical |
| domain-rfz-cli | Understanding existing TuiModal, TuiRadio, TuiCheckbox components |

### Completion Check

```bash
# Build build screen package
cd /Users/lix/xapps/rfz-tui && go build ./internal/ui/screens/build/...

# Lint build screen package
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/ui/screens/build/...
```
