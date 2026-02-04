# TuiCheckbox Component

> Story ID: INTER-002
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
Feature: TuiCheckbox Toggle Component
  Als RFZ-Entwickler
  möchte ich eine Checkbox-Komponente mit charm-style Symbolen,
  damit ich Optionen in Build-Konfigurationen ein- und ausschalten kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Checkbox im unchecked Zustand anzeigen

```gherkin
Scenario: Checkbox zeigt unchecked Symbol
  Given ich habe eine TuiCheckbox mit Label "Skip Tests"
  When die Checkbox nicht aktiviert ist
  Then sehe ich "☐ Skip Tests" mit dem leeren Ballot-Box-Symbol
```

### Szenario 2: Checkbox im checked Zustand anzeigen

```gherkin
Scenario: Checkbox zeigt checked Symbol
  Given ich habe eine TuiCheckbox mit Label "Skip Tests"
  When die Checkbox aktiviert ist
  Then sehe ich "☑ Skip Tests" mit dem angehakten Ballot-Box-Symbol
```

### Szenario 3: Checkbox mit Fokus-Hervorhebung

```gherkin
Scenario: Fokussierte Checkbox wird hervorgehoben
  Given ich habe eine TuiCheckbox mit Label "Enable Debug"
  When die Checkbox fokussiert ist
  Then wird die Checkbox in Cyan hervorgehoben
  And das Symbol und der Text sind visuell deutlich erkennbar
```

### Szenario 4: Deaktivierte Checkbox anzeigen

```gherkin
Scenario: Deaktivierte Checkbox ist ausgegraut
  Given ich habe eine TuiCheckbox mit Label "Premium Feature"
  When die Checkbox deaktiviert ist
  Then wird die Checkbox in gedämpften Farben dargestellt
  And die Interaktion ist visuell als nicht verfügbar erkennbar
```

### Edge Case: Sehr langer Label-Text

```gherkin
Scenario: Checkbox mit langem Label wird gekürzt
  Given ich habe eine TuiCheckbox mit Label "Generate local configuration files for development environment"
  When das Label länger als 40 Zeichen ist
  Then wird der Text mit "..." gekürzt
  And das Checkbox-Symbol bleibt vollständig sichtbar
```

---

## Technische Verifikation (Automated Checks)

> **Hinweis:** Wird vom Architect ausgefüllt

### Datei-Prüfungen

- [x] FILE_EXISTS: internal/ui/components/checkbox.go
- [x] FILE_EXISTS: internal/ui/components/checkbox_test.go

### Inhalt-Prüfungen

- [x] CONTAINS: checkbox.go enthält "SymbolCheckboxUnchecked"
- [x] CONTAINS: checkbox.go enthält "SymbolCheckboxChecked"

### Funktions-Prüfungen

- [x] BUILD_PASS: go build ./internal/ui/components/...
- [x] TEST_PASS: go test ./internal/ui/components/... -run TestCheckbox -v
- [x] LINT_PASS: golangci-lint run ./internal/ui/components/checkbox.go

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
- [x] Architektur-Vorgaben eingehalten (WIE section)
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
| Frontend | `internal/ui/components/styles.go` | Add SymbolCheckboxUnchecked, SymbolCheckboxChecked constants |
| Frontend | `internal/ui/components/checkbox.go` | New file: TuiCheckbox render function |
| Frontend | `internal/ui/components/checkbox_test.go` | New file: Unit tests for TuiCheckbox |

---

### Technical Details

**WAS:**
- Add charm-style checkbox symbols to styles.go (U+2610, U+2611)
- Create stateless TuiCheckbox render function returning styled string
- Support checked/unchecked/focused/disabled states
- Support label truncation using existing Truncate() helper

**WIE (Architektur-Guidance ONLY):**
- Follow existing stateless render pattern from `button.go` (function returns string, no Model)
- Use existing color tokens from `styles.go` (ColorCyan for focus, ColorTextMuted for disabled)
- Use existing `Truncate()` helper from `helpers.go` for long labels
- Apply Lip Gloss styling for all visual rendering (no manual ANSI codes)
- Export function signature: `TuiCheckbox(label string, checked bool, focused bool, disabled bool) string`

**WO:**
- Modify: `internal/ui/components/styles.go` (add ~5 lines for symbols)
- Create: `internal/ui/components/checkbox.go` (~60 LOC)
- Create: `internal/ui/components/checkbox_test.go` (~80 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** None

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
go test ./internal/ui/components/... -run TestCheckbox -v

# Lint check
golangci-lint run ./internal/ui/components/checkbox.go

# Verify symbols exist in styles.go
grep -q "SymbolCheckboxUnchecked" internal/ui/components/styles.go
grep -q "SymbolCheckboxChecked" internal/ui/components/styles.go
```
