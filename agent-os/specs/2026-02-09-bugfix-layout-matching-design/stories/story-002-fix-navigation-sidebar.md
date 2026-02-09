# Fix Navigation Sidebar Styling

> Story ID: LAYOUT-002
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: High
**Type**: Frontend
**Status**: Done
**Estimated Effort**: S
**Dependencies**: LAYOUT-001

---

## Feature

```gherkin
Feature: Korrektes Navigation-Sidebar-Styling
  Als RFZ-Entwickler
  moechte ich dass die Navigation korrekte Active/Select-States, rechtsbuendig ausgerichtete Shortcuts und Tree-Style Hints anzeigt,
  damit die Navigation dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Active State ueberschreibt Select State

```gherkin
Scenario: Active und Selected Item gleichzeitig
  Given ich bin auf der "Build Components" Seite (aktiv)
  And ich navigiere mit j/k zum "Build Components" Eintrag
  When der Eintrag sowohl aktiv als auch ausgewaehlt ist
  Then hat der Eintrag den hellblauen Active-Hintergrund
  And ein Pfeil-Symbol ">" wird vor dem Eintrag angezeigt
```

### Szenario 2: Nur Active (ohne Select)

```gherkin
Scenario: Nur Active State ohne Cursor
  Given ich bin auf der "Build Components" Seite (aktiv)
  And der Cursor ist auf einem anderen Navigationseintrag
  When ich den aktiven Eintrag betrachte
  Then hat "Build Components" den hellblauen Active-Hintergrund
  And es wird KEIN Pfeil-Symbol angezeigt
```

### Szenario 3: Shortcuts sind rechtsbuendig

```gherkin
Scenario: Shortcut-Nummern rechts ausgerichtet
  Given die Navigation ist sichtbar
  When ich die Navigationspunkte betrachte
  Then sind die Shortcut-Nummern (1, 2, 3, 4, q) am rechten Rand der Nav-Box ausgerichtet
```

### Szenario 4: Shortcut-Hints als Tree-Liste

```gherkin
Scenario: Navigation Shortcut-Hints im Tree-Format
  Given die Navigation ist sichtbar
  When ich die Tastaturkuerzel unter den Navigationspunkten betrachte
  Then werden sie als Tree-Liste dargestellt mit Verbindungslinien
  And "1-5 Quick nav" ist das letzte Element mit Abschluss-Linie
```

### Szenario 5: Nav-Container schrumpft auf Inhaltshoehe

```gherkin
Scenario: Navigation Container Hoehe
  Given die Anwendung laeuft in einem 120x40 Terminal
  When ich die Navigation betrachte
  Then endet die Nav-Box nach dem letzten Hint-Eintrag
  And es gibt freien Hintergrund-Raum unterhalb der Nav-Box
```

### Edge Cases

```gherkin
Scenario: Navigation bei minimaler Terminal-Hoehe
  Given das Terminal hat nur 24 Zeilen Hoehe
  When die Anwendung geladen wird
  Then werden alle 5 Navigationspunkte und die Hints noch angezeigt
  And nichts wird abgeschnitten
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/01-welcome-default.png`, `02-nav-build-focused.png` through `06-nav-exit-focused.png`
- Current: `references/current/current-nav-container-build-active-exit-selected.png`

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
- [x] Code implementiert und folgt Style Guide
- [x] Architektur-Vorgaben eingehalten
- [x] Alle Akzeptanzkriterien erfuellt

#### Qualitaetssicherung
- [x] `go build ./...` erfolgreich
- [x] `golangci-lint run ./...` ohne Fehler
- [x] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/navigation.go` | Fix active/select priority, shortcut alignment, tree hints, content height |
| Frontend | `internal/ui/components/keyhints.go` | Add tree-style rendering mode |
| Frontend | `internal/app/app.go` | Update nav rendering to use content-based height |

---

### Technical Details

**WAS:**
- Update `TuiNavItemRender` state priority: active always shows bg color; cursor adds arrow on top
- Fix shortcut right-alignment: use lipgloss.Width() to calculate padding to right edge
- Add tree-style rendering to key hints (├── prefix for items, └── for last)
- Remove fixed height from nav container, let it shrink to content

**WIE (Architektur-Guidance):**
- Follow existing `TuiNavigation()` function signature - add parameters if needed
- Use `lipgloss.JoinVertical()` for tree hint rendering
- Use existing tree symbols from `styles.go` (SymbolTreeBranch, SymbolTreeLast, or define new ones)
- For content-based height: render nav content first, then wrap in border box without explicit height constraint

**WO:**
- `internal/ui/components/navigation.go`
- `internal/ui/components/keyhints.go`
- `internal/app/app.go` (viewNavigation function)

**Abhaengigkeiten:** LAYOUT-001

**Geschaetzte Komplexitaet:** S

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Bubble Tea TUI patterns |

---

### Completion Check

```bash
# Build passes
cd /Users/lix/xapps/rfz-tui && go build ./...

# Lint passes
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...

# Navigation file updated
grep -q "tree\|Tree\|TreeBranch\|treeBranch" /Users/lix/xapps/rfz-tui/internal/ui/components/navigation.go || grep -q "tree\|Tree" /Users/lix/xapps/rfz-tui/internal/ui/components/keyhints.go
```
