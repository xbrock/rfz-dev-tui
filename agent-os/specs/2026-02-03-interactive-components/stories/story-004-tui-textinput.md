# TuiTextInput Component

> Story ID: INTER-004
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
Feature: TuiTextInput Text Entry Component
  Als RFZ-Entwickler
  möchte ich ein Texteingabefeld mit RFZ-Styling,
  damit ich Konfigurationswerte eingeben kann (z.B. Port-Nummern, Pfade).
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Leeres Textfeld mit Placeholder

```gherkin
Scenario: Textfeld zeigt Placeholder-Text
  Given ich habe ein TuiTextInput mit Placeholder "Enter port number"
  When das Feld leer ist
  Then sehe ich den Placeholder-Text in gedämpfter Farbe
  And das Feld hat eine einfache Border
```

### Szenario 2: Fokussiertes Textfeld

```gherkin
Scenario: Fokussiertes Textfeld wird hervorgehoben
  Given ich habe ein TuiTextInput
  When das Feld fokussiert wird
  Then wird die Border in Cyan hervorgehoben
  And der Cursor blinkt im Eingabebereich
```

### Szenario 3: Textfeld mit Eingabe

```gherkin
Scenario: Textfeld zeigt eingegebenen Text
  Given ich habe ein TuiTextInput
  When ich "11090" eingebe
  Then sehe ich "11090" im Eingabefeld
  And der Cursor steht nach dem letzten Zeichen
```

### Szenario 4: Textfeld mit Prompt-Prefix

```gherkin
Scenario: Textfeld mit Prompt-Symbol
  Given ich habe ein TuiTextInput mit Prompt "$"
  When das Feld gerendert wird
  Then sehe ich "$ " vor dem Eingabebereich
  And das Prompt-Symbol ist in Gelb dargestellt
```

### Szenario 5: Deaktiviertes Textfeld

```gherkin
Scenario: Deaktiviertes Textfeld ist nicht editierbar
  Given ich habe ein deaktiviertes TuiTextInput
  When das Feld gerendert wird
  Then ist das Feld ausgegraut
  And der Inhalt ist sichtbar aber nicht editierbar
```

### Edge Case: Zeichenlimit erreicht

```gherkin
Scenario: Textfeld mit Zeichenlimit
  Given ich habe ein TuiTextInput mit maximal 10 Zeichen
  When ich 10 Zeichen eingegeben habe
  Then werden weitere Eingaben ignoriert
  And ein visueller Hinweis zeigt das Limit an
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/textinput.go
- [ ] FILE_EXISTS: internal/ui/components/textinput_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: textinput.go enthält "bubbles/textinput"
- [ ] CONTAINS: textinput.go enthält "ColorCyan"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestTextInput -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/textinput.go

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
| Frontend | `internal/ui/components/textinput.go` | New file: TuiTextInput Bubbles wrapper Model |
| Frontend | `internal/ui/components/textinput_test.go` | New file: Unit tests for TuiTextInput |

---

### Technical Details

**WAS:**
- Create TuiTextInputModel wrapping bubbles/textinput with RFZ styling
- Support placeholder text display
- Support focus state with cyan border highlight
- Support prompt prefix (e.g., "$") in yellow
- Support disabled state
- Support character limit

**WIE (Architektur-Guidance ONLY):**
- Wrap `github.com/charmbracelet/bubbles/textinput` (already in go.mod v0.18.0)
- Create Bubble Tea Model struct with Init/Update/View methods
- Apply existing input styles from `styles.go` (StyleInputNormal, StyleInputFocused)
- Use existing color tokens: ColorCyan for focus, ColorYellow for prompt, ColorTextMuted for placeholder
- Constructor: `NewTuiTextInput(placeholder string, prompt string) TuiTextInputModel`
- Provide helper methods: SetValue, Value, SetCharLimit, SetDisabled

**WO:**
- Create: `internal/ui/components/textinput.go` (~120 LOC)
- Create: `internal/ui/components/textinput_test.go` (~80 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** None (bubbles/textinput already available)

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
go test ./internal/ui/components/... -run TestTextInput -v

# Lint check
golangci-lint run ./internal/ui/components/textinput.go

# Verify bubbles import
grep -q "bubbles/textinput" internal/ui/components/textinput.go
```
