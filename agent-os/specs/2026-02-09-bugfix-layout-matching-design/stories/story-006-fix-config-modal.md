# Fix Config Modal Styling

> Story ID: LAYOUT-006
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: XS
**Dependencies**: LAYOUT-001, LAYOUT-008

---

## Feature

```gherkin
Feature: Korrektes Config-Modal-Styling
  Als RFZ-Entwickler
  moechte ich dass das Build-Configuration-Modal korrekte Shortcut-Farben und Pipe-getrennte Hints zeigt,
  damit das Modal dem genehmigten Design-Prototyp entspricht.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Shortcut-Label Farben

```gherkin
Scenario: Shortcut-Key in Blau, Beschreibung in Grau
  Given das Build-Configuration-Modal ist geoeffnet
  When ich die Shortcut-Hints in den Sektionen betrachte (z.B. "←→ or h/l to select")
  Then sind die Tasten-Symbole "←→" und "h/l" in blauer Farbe
  And der beschreibende Text "or" und "to select" ist in grauer Farbe
```

### Szenario 2: Bottom Hints mit Pipe-Separator

```gherkin
Scenario: Navigation-Hints am unteren Rand des Modals
  Given das Build-Configuration-Modal ist geoeffnet
  When ich die Hinweise unter den Cancel/Start-Buttons betrachte
  Then sind die Hints durch "|" getrennt: "Tab Switch sections | Enter Confirm | Esc Cancel"
```

### Edge Cases

```gherkin
Scenario: Modal bei schmalem Terminal
  Given das Terminal hat 80 Spalten Breite
  When das Build-Configuration-Modal geoeffnet wird
  Then passen alle Sektionen in den verfuegbaren Platz
  And die Shortcut-Hints werden nicht abgeschnitten
```

---

## Pre-Implementation Requirement

**MANDATORY:** Before writing any code, READ and visually compare:
- Prototype: `references/prototype-screenshots/20-config-modal-goals-focus.png` through `34-config-modal-actions-start.png`
- Current: `references/current/currect-build-modal.png`

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

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | `internal/ui/screens/build/config.go` | Fix shortcut label colors (blue keys, grey text), add pipe separators to bottom hints |

---

### Technical Details

**WAS:**
- Update section hint text styling: key parts in `ColorCyan`, description parts in `ColorTextMuted`
- Add " | " separators between bottom modal hints ("Tab Switch sections | Enter Confirm | Esc Cancel")

**WIE:**
- Use `lipgloss.NewStyle().Foreground(ColorCyan)` for key parts and `ColorTextMuted` for description
- Follow existing hint rendering pattern but split key and label rendering

**WO:**
- `internal/ui/screens/build/config.go`

**Abhaengigkeiten:** LAYOUT-001, LAYOUT-008

**Geschaetzte Komplexitaet:** XS

**Relevante Skills:**

| Skill | Pfad | Grund |
|-------|------|-------|
| go-bubbletea | .claude/skills/go-bubbletea.md | Lip Gloss styling |

---

### Completion Check

```bash
cd /Users/lix/xapps/rfz-tui && go build ./...
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...
```
