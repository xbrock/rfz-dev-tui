# Fix General Border Overflow

> Story ID: LAYOUT-008
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: Critical
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: None

---

## Feature

```gherkin
Feature: Korrekte Border-Berechnung ohne Ueberlauf
  Als RFZ-Entwickler
  moechte ich dass alle umrandeten Boxen innerhalb der Terminal-Breite bleiben,
  damit keine visuellen Artefakte durch Ueberlauf auf der rechten Seite entstehen.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Keine Ueberlauf bei 120 Spalten

```gherkin
Scenario: Boxen bei Standard-Terminal-Breite
  Given die Anwendung laeuft in einem 120-Spalten-Terminal
  When ich eine beliebige Seite (Welcome, Build, etc.) betrachte
  Then enden alle Boxen und Rahmen innerhalb der sichtbaren Terminal-Breite
  And keine Rahmenlinien gehen ueber den rechten Rand hinaus
```

### Szenario 2: Keine Ueberlauf bei 80 Spalten

```gherkin
Scenario: Boxen bei minimaler Terminal-Breite
  Given die Anwendung laeuft in einem 80-Spalten-Terminal
  When ich eine beliebige Seite betrachte
  Then enden alle Boxen innerhalb der 80 Spalten
  And der Inhalt wird ggf. gekuerzt aber nicht die Rahmen
```

### Szenario 3: Keine Ueberlauf bei grossem Terminal

```gherkin
Scenario: Boxen bei breitem Terminal
  Given die Anwendung laeuft in einem 200-Spalten-Terminal
  When ich eine beliebige Seite betrachte
  Then enden alle Boxen innerhalb der verfuegbaren Breite
  And es gibt keine visuellen Artefakte
```

### Edge Cases

```gherkin
Scenario: Terminal Resize waehrend Nutzung
  Given die Anwendung laeuft
  When ich das Terminal von 120 auf 80 Spalten verkleinere
  Then passen sich alle Boxen an die neue Breite an
  And kein Rahmen ragt ueber den rechten Rand hinaus
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: Any screenshot from `references/prototype-screenshots/` (all show correct borders)
- Current: Run the app and observe border overflow on right side

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

- [ ] Code implementiert und folgt Style Guide
- [ ] Alle Akzeptanzkriterien erfuellt
- [ ] `go build ./...` erfolgreich
- [ ] `golangci-lint run ./...` ohne Fehler
- [ ] Completion Check Commands erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | `internal/app/app.go` | Fix width calculations in viewBody, viewContent, viewNavigation |
| Frontend | `internal/ui/components/box.go` | Ensure TuiBoxWithWidth accounts for border chars correctly |
| Frontend | `internal/ui/screens/build/selection.go` | Fix content box width calculation |
| Frontend | `internal/ui/screens/build/execution.go` | Fix content box width calculation |

---

### Technical Details

**WAS:**
- Audit all width calculations in app.go: header, nav, content, status bar
- Fix content width: currently `m.width - navWidth - 4` may not account for all border chars
- Ensure that when Lip Gloss adds a border (2 chars wide: left + right), the content width is reduced accordingly
- Fix any screen-level width calculations that don't respect the parent container width

**WIE:**
- Central fix point: `app.go` width calculations for nav and content boxes
- Pattern: `maxWidth = availableWidth - borderLeft(1) - borderRight(1) - paddingLeft(N) - paddingRight(N)`
- Use `lipgloss.Width()` on rendered border boxes to verify they don't exceed terminal width
- For nested boxes (e.g., Build Execution has boxes inside the content area): pass correct inner width to screens

**WO:**
- `internal/app/app.go` (viewBody, viewNavigation, viewContent, contentWidth helper)
- `internal/ui/components/box.go` (TuiBoxWithWidth if needed)
- `internal/ui/screens/build/selection.go` (if inner widths need fixing)
- `internal/ui/screens/build/execution.go` (if inner widths need fixing)

**Abhaengigkeiten:** None

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Lip Gloss width/layout patterns |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
