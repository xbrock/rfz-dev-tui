# TuiDivider Component

> Story ID: CORE-003
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: High
**Type**: Backend
**Estimated Effort**: XS (1 SP)
**Dependencies**: CORE-001

---

## Feature

```gherkin
Feature: TuiDivider Horizontal Separator Component
  Als TUI-Entwickler
  möchte ich horizontale Trennlinien erstellen können,
  damit ich Inhalte visuell voneinander abgrenzen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Single Line Divider

```gherkin
Scenario: TuiDivider mit einfacher Linie
  Given ich möchte Sektionen trennen
  When ich TuiDivider mit DividerSingle aufrufe
  Then wird eine horizontale Linie aus "─" Zeichen gerendert
  And die Linie hat die Standard-Border-Farbe
```

### Szenario 2: Double Line Divider

```gherkin
Scenario: TuiDivider mit doppelter Linie
  Given ich möchte eine wichtige Trennung anzeigen
  When ich TuiDivider mit DividerDouble aufrufe
  Then wird eine horizontale Linie aus "═" Zeichen gerendert
```

### Szenario 3: Fixe Breite

```gherkin
Scenario: TuiDivider mit fixer Breite
  Given ich möchte eine 40 Zeichen breite Trennlinie
  When ich TuiDivider mit width=40 aufrufe
  Then ist die Linie exakt 40 Zeichen breit
```

### Edge Case: Breite 0

```gherkin
Scenario: TuiDivider mit Breite 0
  Given ich rufe TuiDivider mit width=0 auf
  When die Komponente gerendert wird
  Then wird ein leerer String zurückgegeben
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/divider.go

### Inhalt-Prüfungen

- [ ] CONTAINS: divider.go enthält "func TuiDivider("
- [ ] CONTAINS: divider.go enthält "DividerSingle"
- [ ] CONTAINS: divider.go enthält "DividerDouble"
- [ ] CONTAINS: divider.go enthält "strings.Repeat"

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/divider.go`

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
| Backend | internal/ui/components/divider.go | CREATE - TuiDivider component |

**Kritische Integration Points:**
- divider.go → styles.go: Verwendet ColorBorder

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- TuiDivider Funktion für horizontale Trennlinien
- 2 Varianten: Single (─) und Double (═)
- Konfigurierbare Breite

**WIE:**
- Stateless rendering function
- DividerStyle als string const:
  ```go
  type DividerStyle string
  const (
      DividerSingle DividerStyle = "single"
      DividerDouble DividerStyle = "double"
  )
  ```
- Main function:
  - `func TuiDivider(style DividerStyle, width int) string`
- Verwendet strings.Repeat für Linienerzeugung:
  - Single: strings.Repeat("─", width)
  - Double: strings.Repeat("═", width)
- Lip Gloss für Farbgebung: Foreground(ColorBorder)
- Edge case: width <= 0 returns empty string

**WO:**
- `internal/ui/components/divider.go` (NEW) - ~60 LOC

**Abhängigkeiten:** CORE-001 (styles.go)

**Geschätzte Komplexität:** XS (1 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify file exists
test -f internal/ui/components/divider.go && echo "divider.go exists"

# Verify required functions
grep -q "func TuiDivider" internal/ui/components/divider.go && echo "TuiDivider found"

# Verify build
go build ./internal/ui/components/...

# Verify lint
golangci-lint run ./internal/ui/components/divider.go
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Alle CONTAINS checks bestanden
3. BUILD_PASS und LINT_PASS erfolgreich
