# TuiProgress Component

> Story ID: INTER-006
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
Feature: TuiProgress Progress Bar Component
  Als RFZ-Entwickler
  möchte ich einen Fortschrittsbalken mit verschiedenen Stilen,
  damit ich den Build-Fortschritt visuell verfolgen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Block-Gradient Style (Standard)

```gherkin
Scenario: Progress zeigt Block-Gradient Füllung
  Given ich habe eine TuiProgress mit Style "block"
  When der Fortschritt bei 50% ist
  Then sehe ich eine halbe Füllung mit █▓▒░ Gradient
  And der gefüllte Teil geht fließend in den leeren über
```

### Szenario 2: Prozentanzeige

```gherkin
Scenario: Progress zeigt Prozentwert
  Given ich habe eine TuiProgress mit Prozentanzeige aktiviert
  When der Fortschritt bei 75% ist
  Then sehe ich "75%" neben dem Fortschrittsbalken
```

### Szenario 3: Farbverlauf basierend auf Fortschritt

```gherkin
Scenario Outline: Progress Farbe ändert sich mit Fortschritt
  Given ich habe eine TuiProgress
  When der Fortschritt bei <percent>% ist
  Then ist die Balkenfarbe <color>

  Examples:
    | percent | color |
    | 0       | Gelb  |
    | 50      | Gelb-Grün |
    | 100     | Grün  |
```

### Szenario 4: Voller Fortschritt

```gherkin
Scenario: Progress bei 100% zeigt vollständigen Balken
  Given ich habe eine TuiProgress
  When der Fortschritt bei 100% ist
  Then ist der Balken vollständig gefüllt
  And die Farbe ist Grün für Erfolg
```

### Szenario 5: Leerer Fortschritt

```gherkin
Scenario: Progress bei 0% zeigt leeren Balken
  Given ich habe eine TuiProgress
  When der Fortschritt bei 0% ist
  Then ist der Balken vollständig leer
  And nur die Rahmenzeichen sind sichtbar
```

### Szenario 6: Konfigurierbare Breite

```gherkin
Scenario: Progress mit benutzerdefinierter Breite
  Given ich habe eine TuiProgress mit Breite 40 Zeichen
  When der Balken gerendert wird
  Then hat der Fortschrittsbereich genau 40 Zeichen Breite
```

### Edge Case: ASCII Fallback

```gherkin
Scenario: Progress im ASCII-Stil
  Given ich habe eine TuiProgress mit Style "ascii"
  When der Fortschritt bei 50% ist
  Then sehe ich einen Balken im Format "[=====>    ]"
  And der Stil ist auch in einfachen Terminals korrekt
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [x] FILE_EXISTS: internal/ui/components/progress.go
- [x] FILE_EXISTS: internal/ui/components/progress_test.go

### Inhalt-Prüfungen

- [x] CONTAINS: progress.go enthält "bubbles/progress"
- [x] CONTAINS: progress.go enthält "ColorGreen" oder "ColorYellow"

### Funktions-Prüfungen

- [x] BUILD_PASS: go build ./internal/ui/components/...
- [x] TEST_PASS: go test ./internal/ui/components/... -run TestProgress -v
- [x] LINT_PASS: golangci-lint run ./internal/ui/components/progress.go

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

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/progress.go` | New file: TuiProgress Bubbles wrapper Model |
| Frontend | `internal/ui/components/progress_test.go` | New file: Unit tests for TuiProgress |

---

### Technical Details

**WAS:**
- Create TuiProgressModel wrapping bubbles/progress with RFZ styling
- Support block-gradient style (default) with Unicode block chars
- Support percentage display option
- Support color gradient (yellow at 0% to green at 100%)
- Support configurable width
- Support ASCII fallback style

**WIE (Architektur-Guidance ONLY):**
- Wrap `github.com/charmbracelet/bubbles/progress` (already in go.mod v0.18.0)
- Create Bubble Tea Model struct with Init/Update/View methods
- Configure bubbles progress with custom colors using progress.WithGradient()
- Use existing color tokens: ColorYellow, ColorGreen for gradient
- Constructor: `NewTuiProgress(width int, showPercent bool) TuiProgressModel`
- Provide SetPercent method and static render: `TuiProgress(percent float64, width int, showPercent bool) string`

**WO:**
- Create: `internal/ui/components/progress.go` (~100 LOC)
- Create: `internal/ui/components/progress_test.go` (~70 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** None (bubbles/progress already available)

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
go test ./internal/ui/components/... -run TestProgress -v

# Lint check
golangci-lint run ./internal/ui/components/progress.go

# Verify bubbles import
grep -q "bubbles/progress" internal/ui/components/progress.go
```
