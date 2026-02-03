# TuiRadio Component

> Story ID: INTER-003
> Spec: 2026-02-03-interactive-components
> Created: 2026-02-03
> Last Updated: 2026-02-03

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiRadio Button Group Component
  Als RFZ-Entwickler
  möchte ich eine Radio-Button-Gruppe mit charm-style Kreis-Symbolen,
  damit ich eine Option aus mehreren Alternativen auswählen kann (z.B. Maven Goal).
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Radio-Button im unselected Zustand

```gherkin
Scenario: Radio-Button zeigt unselected Symbol
  Given ich habe eine Radio-Gruppe mit Optionen "clean", "install", "package"
  When keine Option ausgewählt ist
  Then sehe ich "◯ clean  ◯ install  ◯ package" mit leeren Kreis-Symbolen
```

### Szenario 2: Radio-Button im selected Zustand

```gherkin
Scenario: Radio-Button zeigt selected Symbol
  Given ich habe eine Radio-Gruppe mit Optionen "clean", "install", "package"
  When "install" ausgewählt ist
  Then sehe ich "◯ clean  ◉ install  ◯ package"
  And nur "install" hat das gefüllte Kreis-Symbol
```

### Szenario 3: Horizontales Layout

```gherkin
Scenario: Radio-Gruppe in horizontaler Anordnung
  Given ich habe eine Radio-Gruppe mit Layout "horizontal"
  When die Gruppe gerendert wird
  Then werden alle Optionen nebeneinander angezeigt
  And die Optionen sind durch Leerzeichen getrennt
```

### Szenario 4: Vertikales Layout

```gherkin
Scenario: Radio-Gruppe in vertikaler Anordnung
  Given ich habe eine Radio-Gruppe mit Layout "vertical"
  When die Gruppe gerendert wird
  Then wird jede Option in einer eigenen Zeile angezeigt
```

### Szenario 5: Fokussierte Option

```gherkin
Scenario: Fokussierte Radio-Option wird hervorgehoben
  Given ich habe eine Radio-Gruppe mit 4 Optionen
  When die zweite Option fokussiert ist
  Then wird die zweite Option in Cyan hervorgehoben
  And die anderen Optionen behalten ihre normale Farbe
```

### Edge Case: Nur eine Option

```gherkin
Scenario: Radio-Gruppe mit einzelner Option
  Given ich habe eine Radio-Gruppe mit nur einer Option "default"
  When die Gruppe gerendert wird
  Then wird die einzelne Option angezeigt
  And sie ist automatisch auswählbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/radio.go
- [ ] FILE_EXISTS: internal/ui/components/radio_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: radio.go enthält "SymbolRadioUnselected"
- [ ] CONTAINS: radio.go enthält "SymbolRadioSelected"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestRadio -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/radio.go

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** READY - Alle DoR-Kriterien erfüllt

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und prüfbar
- [x] Business Value verstanden

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhängigkeiten identifiziert
- [x] Betroffene Komponenten bekannt
- [x] Erforderliche MCP Tools dokumentiert (falls zutreffend)
- [x] Story ist angemessen geschätzt (max 5 Dateien, 400 LOC)

#### Full-Stack Konsistenz
- [x] Alle betroffenen Layer identifiziert
- [x] Integration Type bestimmt
- [x] Kritische Integration Points dokumentiert (wenn Full-stack)

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Code implementiert und folgt Style Guide
- [ ] Architektur-Vorgaben eingehalten
- [ ] Security/Performance Anforderungen erfüllt

#### Qualitätssicherung
- [ ] Alle Akzeptanzkriterien erfüllt
- [ ] Unit Tests geschrieben und bestanden
- [ ] Code Review durchgeführt und genehmigt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/styles.go` | Add SymbolRadioUnselected, SymbolRadioSelected constants |
| Frontend | `internal/ui/components/radio.go` | New file: TuiRadio and TuiRadioGroup render functions |
| Frontend | `internal/ui/components/radio_test.go` | New file: Unit tests for TuiRadio |

---

### Technical Details

**WAS:**
- Add charm-style radio symbols to styles.go (U+25EF, U+25C9)
- Create stateless TuiRadio render function for single radio button
- Create TuiRadioGroup function for rendering multiple options
- Support horizontal and vertical layouts
- Support selected/focused states

**WIE (Architektur-Guidance ONLY):**
- Follow existing stateless render pattern from `button.go`
- Use existing color tokens from `styles.go` (ColorCyan for focus)
- TuiRadio renders single item: `TuiRadio(label string, selected bool, focused bool) string`
- TuiRadioGroup renders group: `TuiRadioGroup(options []string, selectedIndex int, focusedIndex int, horizontal bool) string`
- Use lipgloss.JoinHorizontal/JoinVertical for layout
- Apply Lip Gloss styling for all visual rendering

**WO:**
- Modify: `internal/ui/components/styles.go` (add ~3 lines for symbols, if not already added by INTER-002)
- Create: `internal/ui/components/radio.go` (~80 LOC)
- Create: `internal/ui/components/radio_test.go` (~100 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** None (symbols may be added by INTER-002, check before adding)

**Geschätzte Komplexität:** XS (2 SP)

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No project-specific skills required |

---

### Completion Check

```bash
# Build validation
go build ./internal/ui/components/...

# Unit tests
go test ./internal/ui/components/... -run TestRadio -v

# Lint check
golangci-lint run ./internal/ui/components/radio.go

# Verify symbols exist in styles.go
grep -q "SymbolRadioUnselected" internal/ui/components/styles.go
grep -q "SymbolRadioSelected" internal/ui/components/styles.go
```
