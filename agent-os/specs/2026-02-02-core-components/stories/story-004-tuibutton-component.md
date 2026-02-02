# TuiButton Component

> Story ID: CORE-004
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
Feature: TuiButton Interactive Button Component
  Als TUI-Entwickler
  möchte ich Buttons mit verschiedenen Styles und Keyboard-Shortcuts anzeigen können,
  damit Benutzer Aktionen klar erkennen und auslösen können.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Primary Button

```gherkin
Scenario: Primary Button rendern
  Given ich möchte eine Hauptaktion darstellen
  When ich TuiButton mit ButtonPrimary aufrufe
  Then wird ein Button mit Cyan-Hintergrund gerendert
  And der Text ist fett und dunkel
```

### Szenario 2: Button mit Keyboard Shortcut

```gherkin
Scenario: Button zeigt Keyboard Shortcut an
  Given ich möchte den Button "Build" mit Shortcut "Enter" anzeigen
  When ich TuiButton mit label="Build" und shortcut="Enter" aufrufe
  Then wird "[Enter] Build" angezeigt
  And der Shortcut ist visuell abgesetzt
```

### Szenario 3: Destructive Button

```gherkin
Scenario: Destructive Button für gefährliche Aktionen
  Given ich möchte eine Löschaktion darstellen
  When ich TuiButton mit ButtonDestructive aufrufe
  Then wird ein Button mit rotem Hintergrund gerendert
  And der Text ist weiß und fett
```

### Szenario Outline: Alle Button Varianten

```gherkin
Scenario Outline: TuiButton unterstützt verschiedene Varianten
  Given ich habe Label "Action"
  When ich TuiButton mit <variant> aufrufe
  Then wird der Button mit <expected_color> Hintergrund gerendert

  Examples:
    | variant           | expected_color |
    | ButtonPrimary     | Cyan           |
    | ButtonSecondary   | Border/Outline |
    | ButtonDestructive | Red            |
```

### Szenario 4: Focused Button

```gherkin
Scenario: Button im Focus-Zustand
  Given ich habe einen Primary Button
  When ich focused=true setze
  Then wird der Button mit zusätzlichem visuellen Indikator gerendert
  And die Hervorhebung ist deutlich erkennbar
```

### Edge Case: Langer Label Text

```gherkin
Scenario: Button mit sehr langem Label
  Given ich habe ein Label mit 50 Zeichen
  When ich TuiButton aufrufe
  Then wird das Label auf eine sinnvolle Länge gekürzt
  And der Shortcut bleibt vollständig sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/button.go

### Inhalt-Prüfungen

- [ ] CONTAINS: button.go enthält "func TuiButton("
- [ ] CONTAINS: button.go enthält "ButtonPrimary"
- [ ] CONTAINS: button.go enthält "ButtonSecondary"
- [ ] CONTAINS: button.go enthält "ButtonDestructive"
- [ ] CONTAINS: button.go enthält "shortcut"
- [ ] CONTAINS: button.go enthält "focused"

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/button.go`

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
| Backend | internal/ui/components/button.go | CREATE - TuiButton component |

**Kritische Integration Points:**
- button.go → styles.go: Verwendet StyleButtonPrimary, ColorCyan, ColorDestructive

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- TuiButton Funktion mit 3 Varianten
- Keyboard Shortcut Anzeige
- Focus State

**WIE:**
- Stateless rendering function
- ButtonVariant als string const:
  ```go
  type ButtonVariant string
  const (
      ButtonPrimary     ButtonVariant = "primary"
      ButtonSecondary   ButtonVariant = "secondary"
      ButtonDestructive ButtonVariant = "destructive"
  )
  ```
- Main function:
  - `func TuiButton(label string, variant ButtonVariant, shortcut string, focused bool) string`
- Format: "[Shortcut] Label" wenn shortcut != "", otherwise just "Label"
- Variant styles (from design-system.md):
  - Primary: Background(ColorCyan), Foreground(ColorBackground), Bold(true), Padding(0, 2)
  - Secondary: Border(BorderSingle), BorderForeground(ColorBorder), Foreground(ColorTextPrimary), Padding(0, 2)
  - Destructive: Background(ColorDestructive), Foreground(ColorTextPrimary), Bold(true), Padding(0, 2)
- Focus state: Bold(true).Underline(true)
- Edge case: Long labels truncated to reasonable length (preserve shortcut visibility)

**WO:**
- `internal/ui/components/button.go` (NEW) - ~120 LOC

**Abhängigkeiten:** CORE-001 (styles.go)

**Geschätzte Komplexität:** S (2 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify file exists
test -f internal/ui/components/button.go && echo "button.go exists"

# Verify required functions and constants
grep -q "func TuiButton" internal/ui/components/button.go && echo "TuiButton found"
grep -q "ButtonPrimary" internal/ui/components/button.go && echo "ButtonPrimary found"
grep -q "ButtonDestructive" internal/ui/components/button.go && echo "ButtonDestructive found"

# Verify build
go build ./internal/ui/components/...

# Verify lint
golangci-lint run ./internal/ui/components/button.go
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Alle CONTAINS checks bestanden
3. BUILD_PASS und LINT_PASS erfolgreich
