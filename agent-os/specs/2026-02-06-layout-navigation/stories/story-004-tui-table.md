# TuiTable Component

> Story ID: LAYOUT-004
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiTable Data Table Component
  Als RFZ-Entwickler
  möchte ich Komponenten in einer Tabelle sehen,
  damit ich schnell Übersicht über alle Komponenten bekomme.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Tabelle mit Spaltenüberschriften

```gherkin
Scenario: Tabelle zeigt Column Headers
  Given ich habe eine Tabelle mit Spalten "Name", "Type", "Status"
  When die Tabelle gerendert wird
  Then sehe ich die Spaltenüberschriften in der ersten Zeile
  And die Headers sind fett/hervorgehoben dargestellt
```

### Szenario 2: Selektierbare Zeilen

```gherkin
Scenario: Zeile kann ausgewählt werden
  Given ich habe eine Tabelle mit 10 Zeilen
  When ich mit j/k zur 5. Zeile navigiere
  Then ist die 5. Zeile hervorgehoben
  And ich kann mit Enter die Zeile auswählen
```

### Szenario 3: Zebra Striping (optional)

```gherkin
Scenario: Tabelle hat abwechselnde Zeilenfarben
  Given ich habe Zebra Striping aktiviert
  When die Tabelle gerendert wird
  Then haben gerade und ungerade Zeilen unterschiedliche Hintergründe
  And die Lesbarkeit ist verbessert
```

### Szenario 4: Pagination bei vielen Zeilen

```gherkin
Scenario: Tabelle scrollt bei vielen Zeilen
  Given ich habe eine Tabelle mit 100 Zeilen
  When nur 20 Zeilen sichtbar sind
  Then kann ich mit j/k scrollen
  And die aktuelle Position wird angezeigt
```

### Szenario 5: RFZ-Styling angewandt

```gherkin
Scenario: Tabelle nutzt RFZ Design System
  Given ich rendere eine TuiTable
  When die Tabelle dargestellt wird
  Then verwendet sie ColorBorder für Rahmen
  And die Textfarben entsprechen dem Design System
```

### Edge Case: Leere Tabelle

```gherkin
Scenario: Tabelle ohne Daten
  Given ich habe eine Tabelle ohne Zeilen
  When die Tabelle gerendert wird
  Then sehe ich "No data" als Hinweis
  And die Headers bleiben sichtbar
```

### Edge Case: Lange Zellenwerte

```gherkin
Scenario: Lange Zellenwerte werden gekürzt
  Given ich habe eine Zelle mit sehr langem Text
  When die Spaltenbreite fixiert ist
  Then wird der Text mit "..." gekürzt
  And die Tabellenstruktur bleibt intakt
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/table.go
- [ ] FILE_EXISTS: internal/ui/components/table_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: table.go enthält "TuiTable"
- [ ] CONTAINS: table.go enthält "bubbles/table"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestTable -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/table.go

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Done

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

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/table.go` | New file: TuiTable wrapper around bubbles/table |
| Frontend | `internal/ui/components/table_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add table-specific styles if needed |

**Integration:** Wraps `github.com/charmbracelet/bubbles/table` with RFZ styling

---

### Technical Details

**WAS:**
- `NewTuiTable`: Factory function to create a styled bubbles/table.Model
- `TuiTableStyles`: Returns pre-configured table.Styles with RFZ design tokens
- Support for column headers, selectable rows, optional zebra striping
- "No data" message for empty tables

**WIE (Architektur-Guidance ONLY):**
- Pattern: Bubbles Wrapper (wraps `bubbles/table` with RFZ styling)
- Import `github.com/charmbracelet/bubbles/table`
- Create factory function that returns configured `table.Model`
- Apply RFZ colors via `table.Styles` configuration:
  - Header: `StyleH3` or bold with `ColorTextSecondary`
  - Selected row: `ColorSecondary` background
  - Cell: `ColorTextPrimary`
  - Border: `ColorBorder`
- Use `Truncate()` helper for long cell values
- Reference design-system.md for zebra striping colors

**WO:**
- `internal/ui/components/table.go` - Wrapper implementation
- `internal/ui/components/table_test.go` - Unit tests with golden files
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** None (bubbles/table is external dependency)

**Geschätzte Komplexität:** S

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No special skills required |

---

### Completion Check

```bash
# Build validation
go build ./internal/ui/components/...

# Unit tests
go test ./internal/ui/components/... -run TestTable -v

# Lint check
golangci-lint run ./internal/ui/components/table.go

# Verify TuiTable uses bubbles/table
grep -q "bubbles/table" internal/ui/components/table.go

# Verify TuiTable function exists
grep -q "TuiTable\|NewTuiTable" internal/ui/components/table.go
```
