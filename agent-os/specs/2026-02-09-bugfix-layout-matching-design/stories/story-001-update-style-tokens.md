# Update Style Tokens and Shared Styles

> Story ID: LAYOUT-001
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: Critical
**Type**: Frontend
**Estimated Effort**: XS
**Dependencies**: None

---

## Feature

```gherkin
Feature: Aktualisierte Style-Tokens fuer Design-Konformitaet
  Als RFZ-Entwickler
  moechte ich dass die TUI-Anwendung konsistente Farb- und Style-Tokens verwendet,
  damit alle Komponenten einheitlich gestylt werden koennen.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Navigations-Active-State hat hellblaue Hintergrundfarbe

```gherkin
Scenario: Navigation Active State Farbe
  Given ich bin in der RFZ-CLI Anwendung
  When ich einen Navigationspunkt wie "Build Components" auswaehle
  Then hat der aktive Navigationspunkt einen hellblauen Hintergrund
  And der Text ist gut lesbar auf dem hellblauen Hintergrund
```

### Szenario 2: Header zeigt rote Linie oben

```gherkin
Scenario: Header rote Linie Position
  Given ich bin in der RFZ-CLI Anwendung
  When die Anwendung geladen ist
  Then befindet sich die rote DB-Akzentlinie am oberen Rand der Anwendung
  And "RFZ-CLI v1.0.0" steht als Haupttitel unter der roten Linie
  And "Terminal Orchestration Tool" und die Zeitanzeige stehen auf derselben Zeile
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Styles sind konsistent bei verschiedenen Terminal-Groessen
  Given das Terminal hat eine Groesse von 80x24
  When die Anwendung geladen wird
  Then werden alle Farben und Styles korrekt angezeigt
  And es gibt keine visuellen Artefakte
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/01-welcome-default.png` (header, overall layout)
- Current: `references/current/current-titlebar-header.png` (current header state)

---

## Technisches Refinement (vom Architect)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und pruefbar
- [x] Business Value verstanden

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhaengigkeiten identifiziert
- [x] Betroffene Komponenten bekannt
- [x] Erforderliche MCP Tools dokumentiert (falls zutreffend)
- [x] Story ist angemessen geschaetzt (max 5 Dateien, 400 LOC)

#### Full-Stack Konsistenz
- [x] Alle betroffenen Layer identifiziert
- [x] Integration Type bestimmt
- [x] Kritische Integration Points dokumentiert
- [x] Handover-Dokumente definiert

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Code implementiert und folgt Style Guide
- [ ] Architektur-Vorgaben eingehalten (WIE section)
- [ ] Alle Akzeptanzkriterien erfuellt

#### Qualitaetssicherung
- [ ] `go build ./...` erfolgreich
- [ ] `golangci-lint run ./...` ohne Fehler
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/styles.go` | Add light blue active bg color, update nav active style, update header styles for top-line layout |
| Frontend | `internal/app/app.go` | Update `viewHeader()` to place red line on top, subtitle+time on same line |

---

### Technical Details

**WAS:**
- Add new color token for navigation active background (light blue/cyan)
- Update `StyleNavItemActive` to use light blue background instead of gray
- Update `StyleHeader` to have red border on TOP (not bottom)
- Update header rendering: "RFZ-CLI v1.0.0" on first line, "Terminal Orchestration Tool" + time on second line

**WIE (Architektur-Guidance):**
- Follow existing color token pattern in `styles.go`
- Use `lipgloss.Border()` with `BorderTop(true)` instead of `BorderBottom(true)` for header
- Keep all styling in Lip Gloss - no custom ANSI

**WO:**
- `internal/ui/components/styles.go`
- `internal/app/app.go` (viewHeader function)

**Abhaengigkeiten:** None

**Geschaetzte Komplexitaet:** XS

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Bubble Tea patterns and Lip Gloss styling |

---

### Completion Check

```bash
# Build passes
cd /Users/lix/xapps/rfz-tui && go build ./...

# Lint passes
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...

# Styles file contains new active bg color
grep -q "ColorNavActiveBg\|NavActiveBg\|ActiveBg" /Users/lix/xapps/rfz-tui/internal/ui/components/styles.go

# Header has top border, not bottom
grep -q "BorderTop" /Users/lix/xapps/rfz-tui/internal/app/app.go
```

**Story ist DONE wenn:**
1. Alle Completion Check commands exit 0
2. Header rote Linie ist oben
3. Nav Active State hat hellblauen Hintergrund
