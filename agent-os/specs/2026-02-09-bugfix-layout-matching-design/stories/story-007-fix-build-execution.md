# Fix Build Execution View

> Story ID: LAYOUT-007
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
Feature: Korrektes Build-Execution-View-Layout
  Als RFZ-Entwickler
  moechte ich dass die Build-Execution-Ansicht Tree-Icons, Braille-Fortschrittsbalken und korrektes Spalten-Layout zeigt,
  damit die Build-Uebersicht dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Tree-Icons vor Komponenten

```gherkin
Scenario: Komponenten-Liste mit Tree-Icons
  Given ein Build laeuft mit 3 Komponenten
  When ich die Komponenten-Liste betrachte
  Then haben die ersten beiden Komponenten ein "├─" Tree-Icon
  And die letzte Komponente hat ein "└─" Tree-Icon
  And das Tree-Icon steht vor dem Status-Symbol
```

### Szenario 2: Braille-Fortschrittsbalken pro Komponente

```gherkin
Scenario: Fortschrittsbalken-Stil pro Komponente
  Given ein Build laeuft fuer "traktion"
  When ich die Progress-Spalte betrachte
  Then zeigt der Fortschrittsbalken Braille-Block-Zeichen (⣿)
  And laufende Builds zeigen blaue Bloecke
  And fehlerhafte Builds zeigen rote Bloecke
  And erfolgreiche Builds zeigen gruene Bloecke
```

### Szenario 3: Overall Progress mit Block-Stil

```gherkin
Scenario: Gesamtfortschritt Balken-Stil
  Given ein Build ist zu 67% abgeschlossen
  When ich den Overall Progress betrachte
  Then zeigt der gefuellte Bereich dunkle Block-Zeichen (█)
  And der leere Bereich zeigt helle Block-Zeichen (░)
  And die Prozentanzeige steht rechts vom Balken
```

### Szenario 4: Komponentenliste volle Breite

```gherkin
Scenario: Komponenten-Tabelle nutzt volle Breite
  Given ein Build laeuft in einem 120-Spalten-Terminal
  When ich die Komponentenliste betrachte
  Then nutzt die Tabelle die volle verfuegbare Breite des Content-Bereichs
  And "St" und "Component" sind links ausgerichtet
  And "Phase", "Progress" und "Time" sind rechts ausgerichtet
```

### Szenario 5: Pending-Badge verschwindet nach Abschluss

```gherkin
Scenario: Badge-Cleanup nach Build-Abschluss
  Given alle Builds sind abgeschlossen
  When ich die Progress-Box betrachte
  Then wird kein "Pending: 0" Badge mehr angezeigt
  And der "Running" Badge wird nicht mehr angezeigt
```

### Szenario 6: Kein Running-Badge in Progress-Box

```gherkin
Scenario: Keine Running-Badge-Anzeige
  Given Builds laufen
  When ich die Progress-Box betrachte
  Then wird der "Running: N" Badge nicht angezeigt
  And nur "Success", "Failed" und "Pending" Badges sind sichtbar (wenn > 0)
```

### Edge Cases

```gherkin
Scenario: Nur eine Komponente im Build
  Given ein Build laeuft mit nur 1 Komponente
  When ich die Komponenten-Liste betrachte
  Then hat die einzige Komponente ein "└─" Tree-Icon (letzte = einzige)
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/40-build-execution-starting.png`, `41-build-execution-running.png`, `44-build-execution-progress.png`, `49-build-execution-complete.png`
- Current: `references/current/current-build-running.png`, `references/current/current-build-runnning-2.png`, `references/current/current-build-finished.png`

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
| Frontend | `internal/ui/screens/build/execution.go` | Tree icons, full-width layout, column alignment, progress style, badge cleanup |
| Frontend | `internal/ui/components/progress.go` | Add braille block style mode |

---

### Technical Details

**WAS:**
- Replace `├── ` prefix with proper tree rendering: `├─` for all except last, `└─` for last item
- Use charm.land tree patterns where available (check `lipgloss` tree utilities)
- Change per-component progress bar to use braille blocks (⣿) with color coding
- Change overall progress bar to use ░ (empty) and █ (filled) pattern
- Remove "Running" badge from progress box status counters
- Hide "Pending: 0" when all builds finished
- Adjust column widths to use full available width

**WIE:**
- Tree icons: use `├─`/`└─` characters directly (standard Unicode box-drawing)
- For braille progress: create a simple function that maps percentage to braille block string with color
- For overall progress: use `strings.Repeat("█", filled) + strings.Repeat("░", empty)` with width calculation
- Badge visibility: conditional rendering based on count > 0
- Full-width: calculate available width from parent and distribute to columns proportionally

**WO:**
- `internal/ui/screens/build/execution.go`
- `internal/ui/components/progress.go` (add braille/block style)

**Abhaengigkeiten:** LAYOUT-001, LAYOUT-008

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Bubble Tea rendering patterns |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
