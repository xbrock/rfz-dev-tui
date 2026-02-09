# Fix Build Components Screen

> Story ID: LAYOUT-005
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: LAYOUT-001, LAYOUT-008

---

## Feature

```gherkin
Feature: Korrektes Build-Components-Screen-Layout
  Als RFZ-Entwickler
  moechte ich dass die Build-Komponentenliste korrekte Checkbox-Icons, rechtsbuendig ausgerichtete Kategorien und korrekt gestylte Action-Buttons zeigt,
  damit die Komponentenauswahl dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Neue Checkbox-Icons

```gherkin
Scenario: Checkbox-Icons in Komponentenliste
  Given ich bin auf der Build-Components-Seite
  When ich die Komponentenliste betrachte
  Then zeigen nicht-ausgewaehlte Komponenten ein leeres Kreis-Symbol "○"
  And ausgewaehlte Komponenten zeigen ein gefuelltes gruenes Kreis-Symbol "◉"
```

### Szenario 2: Kategorien rechtsbuendig

```gherkin
Scenario: Kategorie-Badges rechts ausgerichtet
  Given ich bin auf der Build-Components-Seite
  When ich die Komponentenliste betrachte
  Then sind die Kategorie-Badges ("Core", "Simulation", "Standalone") am rechten Rand jeder Zeile positioniert
```

### Szenario 3: Select-State mit blauem Hintergrund

```gherkin
Scenario: Cursor auf einer Komponente
  Given ich bin auf der Build-Components-Seite
  When ich mit j/k den Cursor auf "traktion" bewege
  Then hat die "traktion"-Zeile einen blauen Hintergrund
  And ein Pfeil ">" erscheint vor dem Kreis-Symbol
```

### Szenario 4: Action-Button Styling

```gherkin
Scenario: Build Selected Button Styling
  Given ich bin auf der Build-Components-Seite mit Actions fokussiert
  When ich den "[Build Selected]" Button betrachte
  Then ist der Button-Text in gruener Schrift
  And wenn ich ihn auswaehle wechselt er zu gruenem Hintergrund mit dunkler Schrift
```

### Szenario 5: Shortcut-Hints ohne Zeilenumbruch

```gherkin
Scenario: Shortcut-Zeile passt in eine Zeile
  Given ich bin auf der Build-Components-Seite in einem 120-Spalten-Terminal
  When ich die Shortcut-Hints in der oberen Reihe betrachte ("Space Toggle a All n None")
  Then passen alle Hints in eine einzelne Zeile ohne Umbruch
```

### Szenario 6: Aktualisierte Legende

```gherkin
Scenario: Legende zeigt neue Icons
  Given ich bin auf der Build-Components-Seite
  When ich die Legende unter den Actions betrachte
  Then zeigt sie "◉ Selected  ○ Not selected  > Current"
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/10-build-empty-selection.png`, `11-build-list-navigation.png`, `12-build-partial-selection.png`, `15-build-actions-focused.png`
- Current: `references/current/current-build-components-pre-build.png`

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
| Frontend | `internal/ui/components/list.go` | Change multi-select symbols to ○/◉ for component list mode |
| Frontend | `internal/ui/screens/build/selection.go` | Fix hint overflow, action button colors, legend icons, category alignment |
| Frontend | `internal/ui/components/styles.go` | Add green button active style if needed |

---

### Technical Details

**WAS:**
- Change multi-select checkbox symbols from ☐/☑ to ○/◉ (in component list mode only)
- Make ◉ render in green color
- Fix category badges to align at right edge of list row
- Fix shortcut hints to never wrap to second line
- Update action button styles: Build Selected = green font (active: green bg + dark font), others = light font (active: blue bg)
- Update legend text to show ○/◉ symbols

**WIE:**
- In `list.go`: Add a parameter or new list select mode (e.g. `ListCircleSelect`) that uses ○/◉ instead of ☐/☑
- Use `lipgloss.Width()` for proper right-alignment of category badges within available row width
- For hint overflow: ensure the hint string fits within `contentWidth` before rendering
- Follow existing `TuiButton` variant pattern for green active style

**WO:**
- `internal/ui/components/list.go`
- `internal/ui/screens/build/selection.go`
- `internal/ui/components/styles.go` (if new button style needed)

**Abhaengigkeiten:** LAYOUT-001, LAYOUT-008

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Lip Gloss component patterns |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
