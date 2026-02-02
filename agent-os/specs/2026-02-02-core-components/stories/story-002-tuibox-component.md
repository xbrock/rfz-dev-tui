# TuiBox Component

> Story ID: CORE-002
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: Critical
**Type**: Backend
**Estimated Effort**: S (2 SP)
**Dependencies**: CORE-001

---

## Feature

```gherkin
Feature: TuiBox Bordered Container Component
  Als TUI-Entwickler
  möchte ich einen TuiBox-Container mit verschiedenen Border-Styles,
  damit ich Inhalte visuell gruppieren und hervorheben kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Single Border Box

```gherkin
Scenario: TuiBox mit Single Border rendern
  Given ich habe Content "Hello World"
  When ich TuiBox mit BoxSingle Style aufrufe
  Then wird der Content in einem einfachen Rahmen (┌─┐) dargestellt
  And der Rahmen hat die Standard-Border-Farbe
```

### Szenario 2: Focused Box

```gherkin
Scenario: TuiBox mit Focus-Highlight
  Given ich habe Content "Focused Item"
  When ich TuiBox mit focused=true aufrufe
  Then wird der Rahmen in Cyan (#0891b2) dargestellt
  And der Content bleibt unverändert
```

### Szenario Outline: Alle Border Varianten

```gherkin
Scenario Outline: TuiBox unterstützt verschiedene Border-Styles
  Given ich habe Content "Test Content"
  When ich TuiBox mit <style> aufrufe
  Then wird der entsprechende Border-Typ verwendet

  Examples:
    | style      |
    | BoxSingle  |
    | BoxDouble  |
    | BoxRounded |
    | BoxHeavy   |
```

### Edge Case: Content Overflow

```gherkin
Scenario: Content länger als Box-Breite wird abgeschnitten
  Given ich habe Content mit 100 Zeichen
  And die Box hat eine Breite von 50 Zeichen
  When ich TuiBoxWithWidth aufrufe
  Then wird der Content auf 47 Zeichen gekürzt
  And "..." wird am Ende angehängt
```

### Edge Case: Leerer Content

```gherkin
Scenario: TuiBox mit leerem Content
  Given ich habe leeren Content ""
  When ich TuiBox aufrufe
  Then wird eine Box mit Minimum-Höhe gerendert
  And der Rahmen ist vollständig sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/box.go

### Inhalt-Prüfungen

- [ ] CONTAINS: box.go enthält "func TuiBox("
- [ ] CONTAINS: box.go enthält "func TuiBoxWithWidth("
- [ ] CONTAINS: box.go enthält "BoxSingle"
- [ ] CONTAINS: box.go enthält "BoxDouble"
- [ ] CONTAINS: box.go enthält "BoxRounded"
- [ ] CONTAINS: box.go enthält "BoxHeavy"
- [ ] CONTAINS: box.go enthält "focused"

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/box.go`

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

**Integration Type:** Backend-only

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Backend | internal/ui/components/box.go | CREATE - TuiBox component |

**Kritische Integration Points:**
- box.go → styles.go: Verwendet ColorBorder, ColorCyan, Border-Konstanten

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- TuiBox Funktion für bordered Container
- 5 Border-Varianten (Single, Double, Rounded, Heavy + Focused)
- Optional: fixe Breite mit Truncation

**WIE:**
- Stateless rendering function (kein Bubble Tea Model)
- BoxStyle als string const für Varianten:
  ```go
  type BoxStyle string
  const (
      BoxSingle  BoxStyle = "single"
      BoxDouble  BoxStyle = "double"
      BoxRounded BoxStyle = "rounded"
      BoxHeavy   BoxStyle = "heavy"
  )
  ```
- Main functions:
  - `func TuiBox(content string, style BoxStyle, focused bool) string`
  - `func TuiBoxWithWidth(content string, style BoxStyle, focused bool, width int) string`
- Verwendet styles.go für ColorBorder, ColorCyan, Border-Konstanten
- Focus state: BorderForeground(ColorCyan) when focused=true
- Truncate helper für Content Overflow (width - 3 for "..." and padding)
- All styling via lipgloss.NewStyle().Border().BorderForeground().Padding()

**WO:**
- `internal/ui/components/box.go` (NEW) - ~100 LOC

**Abhängigkeiten:** CORE-001 (styles.go, helpers.go)

**Geschätzte Komplexität:** S (2 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify file exists
test -f internal/ui/components/box.go && echo "box.go exists"

# Verify required functions
grep -q "func TuiBox" internal/ui/components/box.go && echo "TuiBox found"
grep -q "func TuiBoxWithWidth" internal/ui/components/box.go && echo "TuiBoxWithWidth found"

# Verify build
go build ./internal/ui/components/...

# Verify lint
golangci-lint run ./internal/ui/components/box.go
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Alle CONTAINS checks bestanden
3. BUILD_PASS und LINT_PASS erfolgreich
