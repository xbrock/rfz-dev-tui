# Visual Regression Tests

> Story ID: LAYOUT-009
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Test
**Estimated Effort**: S (2 SP)
**Dependencies**: LAYOUT-001 to LAYOUT-007

---

## Feature

```gherkin
Feature: Visual Regression Tests für Layout-Komponenten
  Als RFZ-Entwickler
  möchte ich automatisierte Golden-File-Tests,
  damit UI-Änderungen sofort erkannt werden.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: TuiNavigation Golden Files

```gherkin
Scenario: Navigation-Varianten sind getestet
  Given ich habe TuiNavigation-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Normal (kein Fokus) |
    | Fokussiert |
    | Mit aktivem Item |
    | Mit Header |
    | Mit Footer |
```

### Szenario 2: TuiModal Golden Files

```gherkin
Scenario: Modal-Varianten sind getestet
  Given ich habe TuiModal-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Normal |
    | Mit Buttons |
    | Button fokussiert |
    | Backdrop |
```

### Szenario 3: TuiTable Golden Files

```gherkin
Scenario: Table-Varianten sind getestet
  Given ich habe TuiTable-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Normale Tabelle |
    | Zeile ausgewählt |
    | Zebra Striping |
    | Leere Tabelle |
```

### Szenario 4: TuiTree Golden Files

```gherkin
Scenario: Tree-Varianten sind getestet
  Given ich habe TuiTree-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Alle zugeklappt |
    | Alle aufgeklappt |
    | Gemischt |
    | Leerer Baum |
```

### Szenario 5: TuiTabs Golden Files

```gherkin
Scenario: Tabs-Varianten sind getestet
  Given ich habe TuiTabs-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Normal |
    | Tab aktiv |
    | Tab fokussiert |
    | Mit Badges |
```

### Szenario 6: TuiStatusBar Golden Files

```gherkin
Scenario: StatusBar-Varianten sind getestet
  Given ich habe TuiStatusBar-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Mit allen Sections |
    | Nur links |
    | Mit KeyHints |
```

### Szenario 7: TuiKeyHints Golden Files

```gherkin
Scenario: KeyHints-Varianten sind getestet
  Given ich habe TuiKeyHints-Tests
  When ich die Tests ausführe
  Then werden Golden Files für folgende States geprüft:
    | State |
    | Mehrere Hints |
    | Einzelner Hint |
    | Leere Hints |
```

### Szenario 8: Golden Files bei Änderungen aktualisieren

```gherkin
Scenario: Golden Files können aktualisiert werden
  Given eine Komponente hat sich geändert
  When ich "go test -update" ausführe
  Then werden die Golden Files aktualisiert
  And die neuen Versionen sind im Repository
```

### Edge Case: Terminal-Größe

```gherkin
Scenario: Tests mit kanonischer Terminal-Größe
  Given alle Tests laufen mit 120x40 Terminal
  When eine Komponente gerendert wird
  Then entspricht die Größe der kanonischen Größe
  And die Golden Files sind vergleichbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/navigation_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/modal_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/table_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/tree_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/tabs_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/statusbar_test.go (Golden Tests)
- [ ] FILE_EXISTS: internal/ui/components/keyhints_test.go (Golden Tests)

### Funktions-Prüfungen

- [ ] TEST_PASS: go test ./internal/ui/components/... -v

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Ready

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
- [ ] All golden file tests passing
- [ ] Code Review durchgeführt und genehmigt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Test-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Test | `internal/ui/components/navigation_test.go` | Add golden file tests for all nav states |
| Test | `internal/ui/components/modal_test.go` | Add golden file tests for modal states |
| Test | `internal/ui/components/keyhints_test.go` | Add golden file tests for keyhints |
| Test | `internal/ui/components/table_test.go` | Add golden file tests for table states |
| Test | `internal/ui/components/tree_test.go` | Add golden file tests for tree states |
| Test | `internal/ui/components/tabs_test.go` | Add golden file tests for tab states |
| Test | `internal/ui/components/statusbar_test.go` | Add golden file tests for statusbar |
| Test | `internal/ui/components/testdata/*.golden` | Golden files for all component states |

---

### Technical Details

**WAS:**
- Golden file tests for each layout component covering all visual states
- Tests use `github.com/charmbracelet/x/exp/golden` package
- Update capability with `-update` flag
- Canonical terminal size: 120x40

**WIE (Architektur-Guidance ONLY):**
- Pattern: Follow existing test patterns in `box_test.go`, `list_test.go`
- Use `golden.RequireEqual(t, []byte(output))` for visual assertions
- Test naming: `TestComponentName_State` (e.g., `TestTuiNavigation_Focused`)
- Each component should have tests for:
  - Normal/default state
  - Focused state (if applicable)
  - Active/selected state (if applicable)
  - Empty/edge case state
- Golden files auto-created in `testdata/` directory
- Update golden files: `go test ./internal/ui/components/... -update`

**WO:**
- `internal/ui/components/navigation_test.go` - Navigation golden tests
- `internal/ui/components/modal_test.go` - Modal golden tests
- `internal/ui/components/keyhints_test.go` - KeyHints golden tests
- `internal/ui/components/table_test.go` - Table golden tests
- `internal/ui/components/tree_test.go` - Tree golden tests
- `internal/ui/components/tabs_test.go` - Tabs golden tests
- `internal/ui/components/statusbar_test.go` - StatusBar golden tests
- `internal/ui/components/testdata/` - Golden files directory

**WER:** dev-team__qa-specialist

**Abhängigkeiten:** LAYOUT-001 to LAYOUT-007 (alle Komponenten muessen existieren)

**Geschätzte Komplexität:** S

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No special skills required |

---

### Completion Check

```bash
# All tests must pass
go test ./internal/ui/components/... -v

# Verify golden files exist
ls internal/ui/components/testdata/*.golden 2>/dev/null || echo "No golden files yet"

# Verify test functions exist for each component
grep -q "func TestNav" internal/ui/components/navigation_test.go
grep -q "func TestModal" internal/ui/components/modal_test.go
grep -q "func TestTable" internal/ui/components/table_test.go
grep -q "func TestTree" internal/ui/components/tree_test.go
grep -q "func TestTabs" internal/ui/components/tabs_test.go
grep -q "func TestStatusBar" internal/ui/components/statusbar_test.go
grep -q "func TestKeyHints" internal/ui/components/keyhints_test.go
```
