# teatest Infrastructure

> Story ID: CORE-006
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: High
**Type**: Test
**Estimated Effort**: S (3 SP)
**Dependencies**: CORE-001, CORE-002, CORE-003, CORE-004, CORE-005

---

## Feature

```gherkin
Feature: Visual Regression Testing mit teatest
  Als TUI-Entwickler
  möchte ich Visual Regression Tests mit Golden Files haben,
  damit UI-Änderungen automatisch erkannt und verifiziert werden können.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Golden File Test Setup

```gherkin
Scenario: Golden File Test-Infrastruktur ist eingerichtet
  Given die Test-Dateien existieren
  When ich `go test ./internal/ui/components/...` ausführe
  Then werden alle Komponenten gegen ihre Golden Files verglichen
  And Tests bestehen wenn die Ausgabe übereinstimmt
```

### Szenario 2: Golden Files generieren

```gherkin
Scenario: Golden Files können aktualisiert werden
  Given es gibt Änderungen an Komponenten
  When ich `go test ./internal/ui/components/... -update` ausführe
  Then werden neue Golden Files in testdata/golden/ geschrieben
```

### Szenario 3: Kanonische Terminal-Größe

```gherkin
Scenario: Tests verwenden konsistente Terminal-Größe
  Given ich einen Golden File Test ausführe
  When die Komponente gerendert wird
  Then wird eine 120x40 Terminal-Größe simuliert
```

### Szenario Outline: Alle Komponenten haben Tests

```gherkin
Scenario Outline: Jede Komponente hat Golden File Tests
  Given die Komponente <component> existiert
  When ich die Tests für <test_file> ausführe
  Then werden Golden Files in <golden_dir> verglichen

  Examples:
    | component   | test_file        | golden_dir              |
    | TuiBox      | box_test.go      | testdata/golden/box/    |
    | TuiDivider  | divider_test.go  | testdata/golden/divider/|
    | TuiButton   | button_test.go   | testdata/golden/button/ |
    | TuiStatus   | status_test.go   | testdata/golden/status/ |
```

### Edge Case: Fehlende Golden File

```gherkin
Scenario: Test schlägt fehl bei fehlender Golden File
  Given eine Golden File existiert nicht
  When ich den Test ohne -update Flag ausführe
  Then schlägt der Test fehl
  And eine hilfreiche Fehlermeldung wird angezeigt
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/box_test.go
- [ ] FILE_EXISTS: internal/ui/components/divider_test.go
- [ ] FILE_EXISTS: internal/ui/components/button_test.go
- [ ] FILE_EXISTS: internal/ui/components/status_test.go
- [ ] FILE_EXISTS: testdata/golden/box/
- [ ] FILE_EXISTS: testdata/golden/divider/
- [ ] FILE_EXISTS: testdata/golden/button/
- [ ] FILE_EXISTS: testdata/golden/status/

### Inhalt-Prüfungen

- [ ] CONTAINS: box_test.go enthält "func Test"
- [ ] CONTAINS: go.mod enthält "teatest"

### Funktions-Prüfungen

- [ ] TEST_PASS: `go test ./internal/ui/components/... -v`
- [ ] BUILD_PASS: `go build ./internal/ui/components/...`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

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
- [x] Handover-Dokumente definiert (bei Multi-Layer)

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
- [ ] Dokumentation aktualisiert
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Test

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Test | internal/ui/components/*_test.go | CREATE - Test files |
| Test | testdata/golden/*/  | CREATE - Golden file directories |
| Config | go.mod | UPDATE - Add teatest dependency |

**Kritische Integration Points:**
- Test files → All components: Import und Test aller Komponenten
- Tests → Golden files: Vergleich gegen gespeicherte Ausgaben

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- Test-Dateien für alle 4 Komponenten
- Golden File Verzeichnisstruktur
- teatest Dependency in go.mod

**WIE:**
- teatest für Golden File Vergleiche (github.com/charmbracelet/x/exp/teatest)
- Test pattern für jede Komponente:
  ```go
  func TestTuiBox_Single(t *testing.T) {
      output := components.TuiBox("Content", components.BoxSingle, false)
      golden.RequireEqual(t, []byte(output))
  }
  ```
- Golden files stored with .golden extension
- -update flag: `go test ./... -update` regenerates golden files
- Kanonische Terminal-Größe: 120 columns x 40 rows
- Test cases per component:
  - TuiBox: 5 variants (Single, Double, Rounded, Heavy) x 2 states (normal, focused) = 10 tests
  - TuiDivider: 2 variants (Single, Double) x 2 widths = 4 tests
  - TuiButton: 3 variants x 2 states (normal, focused) x 2 (with/without shortcut) = 12 tests
  - TuiStatus: 5 states x 2 formats (full, compact) = 10 tests
- Import: "github.com/charmbracelet/x/exp/teatest/golden"

**WO:**
- `internal/ui/components/box_test.go` (NEW)
- `internal/ui/components/divider_test.go` (NEW)
- `internal/ui/components/button_test.go` (NEW)
- `internal/ui/components/status_test.go` (NEW)
- `testdata/golden/box/` (NEW directory)
- `testdata/golden/divider/` (NEW directory)
- `testdata/golden/button/` (NEW directory)
- `testdata/golden/status/` (NEW directory)
- `go.mod` (UPDATE - add teatest dependency)

**Abhängigkeiten:** CORE-001 bis CORE-005 (all component implementations)

**Geschätzte Komplexität:** S (3 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify test files exist
test -f internal/ui/components/box_test.go && echo "box_test.go exists"
test -f internal/ui/components/divider_test.go && echo "divider_test.go exists"
test -f internal/ui/components/button_test.go && echo "button_test.go exists"
test -f internal/ui/components/status_test.go && echo "status_test.go exists"

# Verify golden directories exist
test -d testdata/golden/box && echo "golden/box exists"
test -d testdata/golden/divider && echo "golden/divider exists"
test -d testdata/golden/button && echo "golden/button exists"
test -d testdata/golden/status && echo "golden/status exists"

# Run tests
go test ./internal/ui/components/... -v
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. TEST_PASS erfolgreich
3. Golden Files für alle Komponenten existieren
