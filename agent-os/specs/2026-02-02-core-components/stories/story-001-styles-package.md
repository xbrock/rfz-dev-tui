# Styles Package

> Story ID: CORE-001
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: Critical
**Type**: Backend
**Estimated Effort**: S (3 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: Design System Styles Package
  Als TUI-Entwickler
  möchte ich eine zentrale Styles-Bibliothek haben,
  damit alle Komponenten konsistente Farben, Typografie und Abstände verwenden.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Farb-Tokens verfügbar

```gherkin
Scenario: Farb-Tokens aus dem Design System sind verwendbar
  Given ich importiere das components Package
  When ich die Farb-Konstanten verwende
  Then sind alle 14 Farben aus dem Design System verfügbar
  And jede Farbe ist ein lipgloss.Color Typ
```

### Szenario 2: Border-Styles verfügbar

```gherkin
Scenario: Border-Styles sind konfiguriert
  Given ich importiere das components Package
  When ich BorderSingle, BorderDouble, BorderRounded oder BorderHeavy verwende
  Then erhalte ich den entsprechenden lipgloss.Border Typ
```

### Szenario 3: Spacing-Konstanten verfügbar

```gherkin
Scenario: Spacing-Konstanten sind definiert
  Given ich importiere das components Package
  When ich SpaceXS bis Space2XL verwende
  Then sind alle 7 Spacing-Werte als int-Konstanten verfügbar
  And die Werte entsprechen dem Design System (0, 1, 2, 3, 4, 6, 8)
```

### Szenario 4: Typography-Styles verfügbar

```gherkin
Scenario: Typography-Styles sind vorkonfiguriert
  Given ich importiere das components Package
  When ich StyleH1, StyleH2, StyleH3, StyleBody, StyleBodyMuted verwende
  Then sind alle Styles als lipgloss.Style verfügbar
  And verwenden die korrekten Farben und Formatierungen
```

### Edge Case: Adaptive Colors für Hell/Dunkel

```gherkin
Scenario: Adaptive Farben unterstützen beide Terminal-Modi
  Given ich importiere das components Package
  When ich ColorAdaptiveBackground oder ColorAdaptiveForeground verwende
  Then wird automatisch die richtige Farbe für Light/Dark Terminal gewählt
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/styles.go
- [ ] FILE_EXISTS: internal/ui/components/helpers.go

### Inhalt-Prüfungen

- [ ] CONTAINS: styles.go enthält "ColorBackground = lipgloss.Color"
- [ ] CONTAINS: styles.go enthält "ColorCyan = lipgloss.Color"
- [ ] CONTAINS: styles.go enthält "BorderSingle = lipgloss.NormalBorder"
- [ ] CONTAINS: styles.go enthält "SpaceXS = 1"
- [ ] CONTAINS: styles.go enthält "StyleH1 = lipgloss.NewStyle"
- [ ] CONTAINS: helpers.go enthält "func Truncate"

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/styles.go`

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
| Backend | internal/ui/components/styles.go | CREATE - Design tokens |
| Backend | internal/ui/components/helpers.go | CREATE - Utility functions |

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- styles.go mit allen Design-System-Tokens (Farben, Borders, Spacing, Typography)
- helpers.go mit Utility-Funktionen (Truncate)
- Keine externen Abhängigkeiten außer lipgloss

**WIE:**
- Alle Farben als package-level `var` mit `lipgloss.Color()`:
  - ColorBackground (#1e1e2e), ColorCard (#2a2a3e), ColorSecondary (#3a3a4e), ColorBorder (#4a4a5e)
  - ColorCyan (#0891b2), ColorGreen (#22c55e), ColorYellow (#eab308), ColorDestructive (#ef4444), ColorBrand (#ec0016)
  - ColorTextPrimary (#f4f4f5), ColorTextSecondary (#a1a1aa), ColorTextMuted (#71717a), ColorTextDisabled (#52525b)
- Alle Borders als `lipgloss.Border` Typ-Referenzen:
  - BorderSingle = lipgloss.NormalBorder()
  - BorderDouble = lipgloss.DoubleBorder()
  - BorderRounded = lipgloss.RoundedBorder()
  - BorderHeavy = lipgloss.ThickBorder()
- Alle Spacing-Werte als `const int`:
  - SpaceNone=0, SpaceXS=1, SpaceSM=2, SpaceMD=3, SpaceLG=4, SpaceXL=6, Space2XL=8
- Alle Typography Styles als `lipgloss.NewStyle()` Variablen:
  - StyleH1, StyleH2, StyleH3, StyleBody, StyleBodySecondary, StyleBodyMuted
- Adaptive Colors mit `lipgloss.AdaptiveColor{}` for light/dark terminal support
- Truncate(text string, maxWidth int) string function in helpers.go

**WO:**
- `internal/ui/components/styles.go` (NEW) - ~350 LOC
- `internal/ui/components/helpers.go` (NEW) - ~50 LOC

**Abhängigkeiten:** None (foundation package)

**Geschätzte Komplexität:** S (3 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify files exist
test -f internal/ui/components/styles.go && echo "styles.go exists"
test -f internal/ui/components/helpers.go && echo "helpers.go exists"

# Verify build
go build ./internal/ui/components/...

# Verify lint
golangci-lint run ./internal/ui/components/styles.go ./internal/ui/components/helpers.go
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Alle CONTAINS checks bestanden
3. BUILD_PASS und LINT_PASS erfolgreich
