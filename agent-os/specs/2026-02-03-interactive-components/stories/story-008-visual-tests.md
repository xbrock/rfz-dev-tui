# Visual Regression Tests

> Story ID: INTER-008
> Spec: 2026-02-03-interactive-components
> Created: 2026-02-03
> Last Updated: 2026-02-03

**Priority**: High
**Type**: Test
**Estimated Effort**: S (2 SP)
**Dependencies**: INTER-007 (Extend Gallery)

---

## Feature

```gherkin
Feature: Visual Regression Test Coverage
  Als RFZ-Entwickler
  möchte ich Golden-File-Tests für alle interaktiven Komponenten,
  damit UI-Änderungen automatisch erkannt werden und Regressionen verhindert werden.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: TuiCheckbox Golden Files

```gherkin
Scenario: Checkbox hat Golden Files für alle Zustände
  Given die Visual Tests laufen
  When TuiCheckbox getestet wird
  Then existieren Golden Files für:
    | checkbox_unchecked.golden |
    | checkbox_checked.golden   |
    | checkbox_focused.golden   |
    | checkbox_disabled.golden  |
```

### Szenario 2: TuiRadio Golden Files

```gherkin
Scenario: Radio hat Golden Files für alle Varianten
  Given die Visual Tests laufen
  When TuiRadio getestet wird
  Then existieren Golden Files für:
    | radio_horizontal.golden        |
    | radio_vertical.golden          |
    | radio_selected.golden          |
    | radio_focused.golden           |
```

### Szenario 3: TuiList Golden Files

```gherkin
Scenario: List hat Golden Files für beide Modi
  Given die Visual Tests laufen
  When TuiList getestet wird
  Then existieren Golden Files für:
    | list_multiselect.golden   |
    | list_singleselect.golden  |
    | list_with_badges.golden   |
    | list_empty.golden         |
```

### Szenario 4: TuiSpinner Golden Files

```gherkin
Scenario: Spinner hat Golden Files für alle Varianten
  Given die Visual Tests laufen
  When TuiSpinner getestet wird
  Then existieren Golden Files für:
    | spinner_braille.golden |
    | spinner_line.golden    |
    | spinner_circle.golden  |
    | spinner_bounce.golden  |
```

### Szenario 5: TuiProgress Golden Files

```gherkin
Scenario: Progress hat Golden Files für alle Füllstände
  Given die Visual Tests laufen
  When TuiProgress getestet wird
  Then existieren Golden Files für:
    | progress_0_percent.golden   |
    | progress_50_percent.golden  |
    | progress_100_percent.golden |
```

### Szenario 6: Alle Tests bestehen

```gherkin
Scenario: Alle Visual Tests sind grün
  Given alle Golden Files existieren
  When ich "go test ./internal/ui/components/... -v" ausführe
  Then bestehen alle Tests
  And es gibt keine fehlenden Golden Files
```

### Szenario 7: Update-Modus funktioniert

```gherkin
Scenario: Golden Files können aktualisiert werden
  Given eine Komponente wurde geändert
  When ich "go test -update" ausführe
  Then werden die Golden Files aktualisiert
  And der nächste Test-Lauf besteht
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [x] FILE_EXISTS: internal/ui/components/checkbox_test.go
- [x] FILE_EXISTS: internal/ui/components/radio_test.go
- [x] FILE_EXISTS: internal/ui/components/list_test.go
- [x] FILE_EXISTS: internal/ui/components/textinput_test.go
- [x] FILE_EXISTS: internal/ui/components/spinner_test.go
- [x] FILE_EXISTS: internal/ui/components/progress_test.go

### Inhalt-Prüfungen

- [x] CONTAINS: checkbox_test.go enthält "golden" oder "Golden"
- [x] CONTAINS: radio_test.go enthält "golden" oder "Golden"

### Funktions-Prüfungen

- [x] TEST_PASS: go test ./internal/ui/components/... -v
- [x] TEST_PASS: go test ./internal/ui/components/demo/... -v

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Done - Implementation complete

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
- [x] Code implementiert und folgt Style Guide
- [x] Architektur-Vorgaben eingehalten
- [x] Security/Performance Anforderungen erfüllt

#### Qualitätssicherung
- [x] Alle Akzeptanzkriterien erfüllt
- [x] Unit Tests geschrieben und bestanden
- [x] Code Review durchgeführt und genehmigt

#### Dokumentation
- [x] Keine Linting Errors
- [x] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only (Test)

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Test | `internal/ui/components/checkbox_test.go` | Add golden file tests |
| Test | `internal/ui/components/radio_test.go` | Add golden file tests |
| Test | `internal/ui/components/list_test.go` | Add golden file tests |
| Test | `internal/ui/components/textinput_test.go` | Add golden file tests |
| Test | `internal/ui/components/spinner_test.go` | Add golden file tests |
| Test | `internal/ui/components/progress_test.go` | Add golden file tests |
| Test | `internal/ui/components/testdata/` | New directory for golden files |

---

### Technical Details

**WAS:**
- Add golden file test functions to each component's test file
- Create testdata directory for golden files
- Test all visual states for each component (unchecked/checked, focused/unfocused, etc.)
- Support -update flag to regenerate golden files
- Ensure 97+ visual states are covered across all components

**WIE (Architektur-Guidance ONLY):**
- Follow existing test patterns in component test files
- Use standard Go golden file pattern with testdata directory
- Each test renders component to string, compares against .golden file
- Use `go test -update` flag to regenerate golden files when intentional changes
- Test naming: TestCheckbox_Golden, TestRadio_Golden, etc.
- Golden file naming: `testdata/checkbox_unchecked.golden`, etc.
- Use fixed terminal width (120 cols) for consistent rendering

**WO:**
- Modify: `internal/ui/components/checkbox_test.go` (add ~40 LOC)
- Modify: `internal/ui/components/radio_test.go` (add ~50 LOC)
- Modify: `internal/ui/components/list_test.go` (add ~60 LOC)
- Modify: `internal/ui/components/textinput_test.go` (add ~40 LOC)
- Modify: `internal/ui/components/spinner_test.go` (add ~40 LOC)
- Modify: `internal/ui/components/progress_test.go` (add ~40 LOC)
- Create: `internal/ui/components/testdata/*.golden` (~20 files)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** INTER-007 (Extend Gallery) - all components must be complete

**Geschätzte Komplexität:** S (2 SP)

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No project-specific skills required |

---

### Completion Check

```bash
# Run all component tests
go test ./internal/ui/components/... -v

# Run demo/gallery tests
go test ./internal/ui/components/demo/... -v

# Verify golden files exist
ls internal/ui/components/testdata/*.golden | wc -l

# Lint check
golangci-lint run ./internal/ui/components/...
```
