# Domain Model & Mock Data Provider

> Story ID: BUILD-001
> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Backend
**Estimated Effort**: S
**Dependencies**: None

---

## Feature

```gherkin
Feature: Domain Model fuer Build-Workflow
  Als RFZ Developer CLI
  moechte ich ein Domain-Modell fuer Komponenten und Build-Konfigurationen haben,
  damit die Build-Screens typsichere Daten verwenden koennen.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Komponenten-Liste abrufen

```gherkin
Scenario: Mock-Provider liefert alle RFZ-Komponenten
  Given der MockComponentProvider ist initialisiert
  When ich die Komponentenliste abrufe
  Then erhalte ich 13 Komponenten
  And jede Komponente hat einen Namen und einen Typ (Core, Simulation oder Standalone)
```

### Szenario 2: Komponenten-Typen korrekt zugeordnet

```gherkin
Scenario Outline: Komponenten haben den korrekten Typ
  Given der MockComponentProvider ist initialisiert
  When ich die Komponente "<name>" abrufe
  Then hat sie den Typ "<typ>"

  Examples:
    | name              | typ         |
    | boss              | Core        |
    | fistiv            | Core        |
    | audiocon          | Standalone  |
    | simkern           | Simulation  |
    | fahrdynamik       | Simulation  |
    | diagnose          | Standalone  |
```

### Szenario 3: Build-Konfiguration zusammenstellen

```gherkin
Scenario: Standard-Build-Konfiguration erzeugen
  Given ich erstelle eine neue Build-Konfiguration
  When ich das Ziel "clean install" waehle
  And ich das Profil "generate_local_config_files" aktiviere
  And ich Port 11090 waehle
  And ich "Skip Tests" aktiviere
  Then wird der Maven-Befehl "mvn clean install -Pgenerate_local_config_files,use_traktion_11090 -DskipTests" generiert
```

### Szenario 4: Build-Phasen definiert

```gherkin
Scenario: Build-Phasen sind vollstaendig definiert
  Given das Domain-Modell ist geladen
  When ich die Build-Phasen abfrage
  Then existieren die Phasen: Pending, Compiling, Testing, Packaging, Installing, Done, Failed
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Build-Konfiguration ohne Profile
  Given ich erstelle eine neue Build-Konfiguration
  When ich keine Profile aktiviere
  And ich das Ziel "install" waehle
  Then wird der Maven-Befehl "mvn install" generiert (ohne -P Flag)
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/domain/component.go
- [ ] FILE_EXISTS: internal/domain/buildconfig.go
- [ ] FILE_EXISTS: internal/domain/mock_provider.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/domain/component.go enthaelt "ComponentProvider"
- [ ] CONTAINS: internal/domain/component.go enthaelt "ComponentTypeCore"
- [ ] CONTAINS: internal/domain/component.go enthaelt "ComponentTypeSimulation"
- [ ] CONTAINS: internal/domain/component.go enthaelt "ComponentTypeStandalone"
- [ ] CONTAINS: internal/domain/buildconfig.go enthaelt "ToCommand"
- [ ] CONTAINS: internal/domain/mock_provider.go enthaelt "MockComponentProvider"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `cd /Users/lix/xapps/rfz-tui && go build ./internal/domain/...`
- [ ] TEST_PASS: `cd /Users/lix/xapps/rfz-tui && go test ./internal/domain/...`
- [ ] LINT_PASS: `cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/domain/...`

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

- [ ] Code implemented following Bubble Tea patterns
- [ ] Domain types are minimal and focused (no boilerplate cruft)
- [ ] MockComponentProvider returns 13 hardcoded components
- [ ] BuildConfig.ToCommand() generates correct Maven commands
- [ ] All acceptance criteria met
- [ ] Unit tests written and passing for domain logic
- [ ] No linting errors
- [ ] Completion Check commands successful

### Betroffene Layer & Komponenten

| Layer | Integration Type | Affected Components/Files |
|-------|------------------|--------------------------|
| Domain | Internal | `internal/domain/component.go`, `internal/domain/buildconfig.go`, `internal/domain/mock_provider.go` |
| UI | None (consumed by BUILD-002) | N/A |
| Integration | Frontend-only | TUI application, no backend/database |

### Technical Details

**WAS** (What to create):
- Component type with Name, Type (Core/Simulation/Standalone)
- ComponentType enum with three variants
- ComponentProvider interface for abstraction
- MockComponentProvider implementation with 13 hardcoded RFZ components
- BuildConfig value object with Goal, Profiles, Port, SkipTests fields
- MavenGoal type for build goal selection
- BuildPhase enum for execution states (Pending, Compiling, Testing, Packaging, Installing, Done, Failed)
- ToCommand() method on BuildConfig that generates Maven command strings

**WIE** (How to implement):
- Create lightweight domain types focused only on UI needs
- Do NOT use heavyweight boilerplate domain model (no GroupID, ArtifactID, Version, Dependencies)
- ComponentProvider interface enables future replacement with real file-system scanner
- MockComponentProvider hardcodes 13 components matching references/_test-data/demo-components/
- BuildConfig uses Go structs with simple string/bool fields
- ToCommand() method builds Maven command string with conditional flags
- BuildPhase enum uses Go iota pattern
- Follow Go naming conventions (package domain, lowercase visibility by default)
- Pure data types with minimal methods

**WO** (Where to implement):
- `/Users/lix/xapps/rfz-tui/internal/domain/component.go` - Component, ComponentType, ComponentProvider interface
- `/Users/lix/xapps/rfz-tui/internal/domain/buildconfig.go` - BuildConfig, MavenGoal, BuildPhase types, ToCommand() method
- `/Users/lix/xapps/rfz-tui/internal/domain/mock_provider.go` - MockComponentProvider implementation

**WER** (Who implements):
- Generic fullstack developer (no specialized domain agents available)

**Abhaengigkeiten** (Dependencies):
- None (foundational story)

**Geschaetzte Komplexitaet** (Estimated Complexity):
- S (Small) - 3 files, ~150 LOC, simple data structures

**Relevante Skills** (Required Skills):

| Skill | Reason |
|-------|--------|
| go-fundamentals | Go struct definitions, interfaces, enums |
| domain-modeling | Lightweight domain types design |

### Completion Check

```bash
# Build domain package
cd /Users/lix/xapps/rfz-tui && go build ./internal/domain/...

# Run domain tests
cd /Users/lix/xapps/rfz-tui && go test ./internal/domain/...

# Lint domain package
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./internal/domain/...
```
