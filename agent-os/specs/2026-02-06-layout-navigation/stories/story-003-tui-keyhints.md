# TuiKeyHints Component

> Story ID: LAYOUT-003
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Frontend
**Estimated Effort**: XS (1 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiKeyHints Keyboard Shortcuts Display
  Als RFZ-Entwickler
  möchte ich verfügbare Keyboard-Shortcuts sehen,
  damit ich die Anwendung effizient per Tastatur bedienen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Key + Label Format

```gherkin
Scenario: KeyHints zeigt Taste und Aktion
  Given ich habe KeyHints mit "Enter" -> "Select"
  When die KeyHints gerendert werden
  Then sehe ich "Enter Select"
  And "Enter" ist in Cyan hervorgehoben
```

### Szenario 2: Mehrere Hints mit Separator

```gherkin
Scenario: Mehrere KeyHints durch Separator getrennt
  Given ich habe KeyHints ["Enter Select", "Esc Cancel", "q Quit"]
  When die KeyHints gerendert werden
  Then sehe ich "Enter Select • Esc Cancel • q Quit"
  And der Separator "•" ist in gedämpfter Farbe
```

### Szenario 3: Adaptive Breite

```gherkin
Scenario: KeyHints passen sich der verfügbaren Breite an
  Given ich habe 5 KeyHints
  When die verfügbare Breite 80 Zeichen ist
  Then werden alle Hints die reinpassen angezeigt
  And überzählige Hints werden nicht angezeigt
```

### Szenario 4: Context-aware Hints

```gherkin
Scenario: KeyHints ändern sich je nach Screen
  Given ich bin auf dem "Build" Screen
  When die KeyHints aktualisiert werden
  Then sehe ich build-spezifische Shortcuts
  And allgemeine Shortcuts wie "q Quit" bleiben
```

### Edge Case: Leere KeyHints

```gherkin
Scenario: Keine KeyHints definiert
  Given ich habe keine KeyHints konfiguriert
  When die KeyHints gerendert werden
  Then wird ein leerer String zurückgegeben
  And kein Separator wird angezeigt
```

### Edge Case: Einzelner Hint

```gherkin
Scenario: Nur ein KeyHint
  Given ich habe nur "q Quit" als KeyHint
  When die KeyHints gerendert werden
  Then sehe ich "q Quit" ohne Separator
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/keyhints.go
- [ ] FILE_EXISTS: internal/ui/components/keyhints_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: keyhints.go enthält "TuiKeyHints"
- [ ] CONTAINS: keyhints.go enthält "•" oder "separator"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestKeyHints -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/keyhints.go

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
| Frontend | `internal/ui/components/keyhints.go` | New file: TuiKeyHints |
| Frontend | `internal/ui/components/keyhints_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add separator symbol constant |

---

### Technical Details

**WAS:**
- `KeyHint`: Struct with Key and Label fields (e.g., Key="Enter", Label="Select")
- `TuiKeyHints`: Renders a horizontal list of keyboard hints with separators
- Format: "Key Label" with Key in cyan, separated by muted bullet

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (simple stateless render function)
- Use `ColorCyan` for key styling, `ColorTextSecondary` for label
- Use `ColorTextMuted` for separator character
- Separator: middle dot "." (U+00B7) rendered between hints
- Use `lipgloss.JoinHorizontal` for combining hints
- Optional: width parameter to truncate hints that don't fit
- Reference `StyleKeyboard` from `styles.go` for key styling

**WO:**
- `internal/ui/components/keyhints.go` - Main component implementation
- `internal/ui/components/keyhints_test.go` - Unit tests with golden files
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** None

**Geschätzte Komplexität:** XS

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
go test ./internal/ui/components/... -run TestKeyHints -v

# Lint check
golangci-lint run ./internal/ui/components/keyhints.go

# Verify TuiKeyHints function exists
grep -q "func TuiKeyHints" internal/ui/components/keyhints.go
```
