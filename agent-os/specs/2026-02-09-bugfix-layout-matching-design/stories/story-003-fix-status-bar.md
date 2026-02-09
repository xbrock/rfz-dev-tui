# Fix Status Bar Layout

> Story ID: LAYOUT-003
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: LAYOUT-001

---

## Feature

```gherkin
Feature: Korrektes Status-Bar-Layout
  Als RFZ-Entwickler
  moechte ich dass die Status-Bar 3 Badges, Pipe-getrennte Nav-Hints und "q Quit" rechts anzeigt,
  damit die Footer-Leiste dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: 3 Badges werden angezeigt

```gherkin
Scenario: Status Bar Badges auf der Build-Seite
  Given ich bin auf der "Build Components" Seite
  And ich habe den Cursor auf "fistiv"
  When ich die Status-Bar am unteren Rand betrachte
  Then sehe ich 3 Badges links: "SELECT" (farbig), "fistiv" (Komponentenname), und optional einen Status-Badge
```

### Szenario 2: Nav-Hints mit Pipe-Separator

```gherkin
Scenario: Nav-Hints Format
  Given die Status-Bar ist sichtbar
  When ich die Tastaturkuerzel betrachte
  Then sind die Hints durch "|" getrennt: "Tab Focus | ↑↓ Nav | Enter Select | Esc Back"
```

### Szenario 3: q Quit rechts ausgerichtet

```gherkin
Scenario: Quit-Hint Positionierung
  Given die Status-Bar ist sichtbar
  When ich den rechten Rand der Status-Bar betrachte
  Then steht dort nur "q Quit"
  And es ist rechtsbuendig am Rand positioniert
```

### Szenario 4: Grauer Hintergrund ueber volle Breite

```gherkin
Scenario: Status Bar Hintergrund
  Given die Anwendung laeuft in einem 120-Spalten-Terminal
  When ich die Status-Bar betrachte
  Then hat die gesamte Bar einen grauen Hintergrund
  And der Hintergrund erstreckt sich ueber die volle Terminal-Breite
```

### Edge Cases

```gherkin
Scenario: Status Bar bei langen Komponentennamen
  Given der aktuelle Komponentenname ist "signalsteuerung" (15 Zeichen)
  When die Status-Bar gerendert wird
  Then wird der Badge korrekt angezeigt ohne Ueberlauf
  And die Nav-Hints bleiben lesbar
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/10-build-empty-selection.png` (status bar visible at bottom)
- Prototype: `references/prototype-screenshots/41-build-execution-running.png` (BUILD + component + RUNNING badges)
- Current: `references/current/current-statusbar-build-select-konfiguration-conponent.png`

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
| Frontend | `internal/ui/components/statusbar.go` | 3-badge system, pipe-separated hints, gray bg full width |
| Frontend | `internal/app/app.go` | Update `viewStatusBar()` to provide 3 badges and pipe hints |

---

### Technical Details

**WAS:**
- Update `TuiStatusBarConfig` to support 3 badges (area, selection, optional state)
- Change hint separator from space to " | "
- Ensure gray background covers full terminal width
- Move "q Quit" to far right, separated from other hints

**WIE:**
- Extend `TuiStatusBarConfig` struct with optional 3rd badge field
- Use `lipgloss.JoinHorizontal()` for badge layout
- Use full-width `lipgloss.NewStyle().Width(totalWidth).Background(grayColor)` for bar
- Pipe separator: render " | " between each hint in `TuiKeyHints` or in status bar layout

**WO:**
- `internal/ui/components/statusbar.go`
- `internal/app/app.go` (viewStatusBar function)

**Abhaengigkeiten:** LAYOUT-001

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Lip Gloss layout patterns |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
