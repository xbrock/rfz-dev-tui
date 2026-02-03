# TuiSpinner Component

> Story ID: INTER-005
> Spec: 2026-02-03-interactive-components
> Created: 2026-02-03
> Last Updated: 2026-02-03

**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiSpinner Loading Indicator Component
  Als RFZ-Entwickler
  möchte ich einen animierten Lade-Indikator mit verschiedenen Stilen,
  damit ich während Build-Vorgängen visuelles Feedback erhalte.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Braille-Dots Spinner (Standard)

```gherkin
Scenario: Spinner zeigt Braille-Dots Animation
  Given ich habe einen TuiSpinner mit Variante "braille"
  When der Spinner aktiv ist
  Then sehe ich eine fließende Animation mit Braille-Zeichen
  And die Animation läuft gleichmäßig durch
```

### Szenario 2: Line Spinner (Fallback)

```gherkin
Scenario: Spinner zeigt Line Animation
  Given ich habe einen TuiSpinner mit Variante "line"
  When der Spinner aktiv ist
  Then sehe ich eine Animation mit |/-\ Zeichen
  And die Animation ist auch in einfachen Terminals sichtbar
```

### Szenario 3: Circle-Quarters Spinner

```gherkin
Scenario: Spinner zeigt Kreis-Viertel Animation
  Given ich habe einen TuiSpinner mit Variante "circle"
  When der Spinner aktiv ist
  Then sehe ich eine Animation mit ◴◷◶◵ Zeichen
  And die Animation vermittelt eine rotierende Bewegung
```

### Szenario 4: Bounce Spinner

```gherkin
Scenario: Spinner zeigt Bounce Animation
  Given ich habe einen TuiSpinner mit Variante "bounce"
  When der Spinner aktiv ist
  Then sehe ich eine vertikale Bounce-Animation
  And die Animation wirkt lebendig und dynamisch
```

### Szenario 5: Spinner mit Label

```gherkin
Scenario: Spinner zeigt begleitenden Text
  Given ich habe einen TuiSpinner mit Label "Compiling..."
  When der Spinner aktiv ist
  Then sehe ich das Spinner-Symbol gefolgt von "Compiling..."
  And beide Elemente sind in derselben Zeile
```

### Szenario 6: Farbvarianten

```gherkin
Scenario Outline: Spinner in verschiedenen Farben
  Given ich habe einen TuiSpinner mit Farbe "<color>"
  When der Spinner gerendert wird
  Then wird der Spinner in <color_display> dargestellt

  Examples:
    | color   | color_display |
    | cyan    | Cyan          |
    | green   | Grün          |
    | yellow  | Gelb          |
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/spinner.go
- [ ] FILE_EXISTS: internal/ui/components/spinner_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: spinner.go enthält "bubbles/spinner"
- [ ] CONTAINS: spinner.go enthält "braille" oder "Braille"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestSpinner -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/spinner.go

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
| Frontend | `internal/ui/components/spinner.go` | New file: TuiSpinner Bubbles wrapper Model |
| Frontend | `internal/ui/components/spinner_test.go` | New file: Unit tests for TuiSpinner |

---

### Technical Details

**WAS:**
- Create TuiSpinnerModel wrapping bubbles/spinner with RFZ styling
- Support 4 spinner variants: braille (default), line, circle, bounce
- Support optional label text
- Support color configuration (cyan, green, yellow)
- Provide static render function for gallery display

**WIE (Architektur-Guidance ONLY):**
- Wrap `github.com/charmbracelet/bubbles/spinner` (already in go.mod v0.18.0)
- Create Bubble Tea Model struct with Init/Update/View methods
- Define SpinnerVariant type with constants (SpinnerBraille, SpinnerLine, SpinnerCircle, SpinnerBounce)
- Map variants to bubbles spinner types (spinner.Dot, spinner.Line, etc.)
- Apply existing color tokens from `styles.go`
- Constructor: `NewTuiSpinner(variant SpinnerVariant, label string) TuiSpinnerModel`
- Add static helper: `TuiSpinnerStatic(variant SpinnerVariant, label string) string` for gallery

**WO:**
- Create: `internal/ui/components/spinner.go` (~100 LOC)
- Create: `internal/ui/components/spinner_test.go` (~70 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** None (bubbles/spinner already available)

**Geschätzte Komplexität:** S (2 SP)

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
go test ./internal/ui/components/... -run TestSpinner -v

# Lint check
golangci-lint run ./internal/ui/components/spinner.go

# Verify bubbles import
grep -q "bubbles/spinner" internal/ui/components/spinner.go
```
