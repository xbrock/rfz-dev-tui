# Fix Welcome Screen Layout

> Story ID: LAYOUT-004
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: LAYOUT-001, LAYOUT-002, LAYOUT-003

---

## Feature

```gherkin
Feature: Korrektes Welcome-Screen-Layout
  Als RFZ-Entwickler
  moechte ich dass der Welcome-Screen das RFZ-CLI Logo in korrekten Farben, weissen Untertitel, Braille-Zierlinie und 3 Version-Badges zeigt,
  damit der Startbildschirm dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Logo in korrekten Farben

```gherkin
Scenario: RFZ-CLI ASCII-Logo Farben
  Given ich starte die RFZ-CLI Anwendung
  When der Welcome-Screen angezeigt wird
  Then sind die "RFZ" Buchstaben in DB-Rot (Brand-Farbe) dargestellt
  And die "CLI" Buchstaben sind in Cyan dargestellt
```

### Szenario 2: Untertitel in Weiss

```gherkin
Scenario: Terminal Orchestration Tool Farbe
  Given der Welcome-Screen ist sichtbar
  When ich den Text unter dem Logo betrachte
  Then ist "Terminal Orchestration Tool" in weisser Farbe dargestellt
```

### Szenario 3: Braille-Zierlinie statt einfacher Linie

```gherkin
Scenario: Dekorative Linie
  Given der Welcome-Screen ist sichtbar
  When ich die Linie unter dem Zitat betrachte
  Then ist sie als Braille-Block-Muster dargestellt (Punktmuster statt durchgezogene Linie)
```

### Szenario 4: Drei Version-Badges

```gherkin
Scenario: Version und Branding Badges
  Given der Welcome-Screen ist sichtbar
  When ich die Badge-Reihe unter der Zierlinie betrachte
  Then sehe ich 3 separate Badges nebeneinander:
  And "v1.0.0" hat einen roten Hintergrund
  And "Deutsche Bahn" hat einen dunkelgrauen Hintergrund mit hellem Text
  And "Internal Tool" hat einen hellblauen Hintergrund mit dunklem Text
```

### Szenario 5: Shortcut-Hints als Tree-View

```gherkin
Scenario: Welcome Screen Shortcut-Hints
  Given der Welcome-Screen ist sichtbar
  When ich die Tastaturkuerzel unten betrachte
  Then werden sie als Tree-Liste dargestellt
  And enthalten "↑↓/jk navigate", "Enter select", "q quit"
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/01-welcome-default.png`
- Current: Run the app and observe the welcome screen

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
- [x] Story ist angemessen geschaetzt (max 5 Dateien, 400 LOC)

#### Full-Stack Konsistenz
- [x] Alle betroffenen Layer identifiziert
- [x] Integration Type bestimmt
- [x] Kritische Integration Points dokumentiert
- [x] Handover-Dokumente definiert

---

### DoD (Definition of Done) - Vom Architect

- [x] Code implementiert und folgt Style Guide
- [x] Alle Akzeptanzkriterien erfuellt
- [x] `go build ./...` erfolgreich
- [x] `golangci-lint run ./...` ohne Fehler
- [x] Completion Check Commands erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | `internal/ui/screens/welcome/welcome.go` | Logo colors, subtitle white, braille line, 3 badges, tree hints |

---

### Technical Details

**WAS:**
- Fix ASCII logo colors: RFZ in `ColorBrand` (#ec0016), CLI in `ColorCyan` (#0891b2)
- Change "Terminal Orchestration Tool" to `ColorTextPrimary` (white)
- Replace simple divider line with braille block characters (e.g. `⣿⣿⣿⣿`)
- Create 3 separate badges: v1.0.0 (red bg), Deutsche Bahn (gray bg), Internal Tool (light blue bg)
- Use tree-style hints for keyboard shortcuts at bottom

**WIE:**
- Use existing `StyleASCIIArt` (brand red) and `StyleASCIIArtCyan` from design system
- For braille line: simple string of braille characters rendered with `ColorTextMuted`
- For badges: use `lipgloss.NewStyle().Background().Foreground().Padding(0,1)` pattern from existing badge code
- For tree hints: reuse the tree-style rendering from Story 002's keyhints update

**WO:**
- `internal/ui/screens/welcome/welcome.go`

**Abhaengigkeiten:** LAYOUT-001, LAYOUT-002 (tree hints), LAYOUT-003 (status bar)

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Lip Gloss styling patterns |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
