# Extend Component Gallery

> Story ID: INTER-007
> Spec: 2026-02-03-interactive-components
> Created: 2026-02-03
> Last Updated: 2026-02-03

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: INTER-001, INTER-002, INTER-003, INTER-004, INTER-005, INTER-006

---

## Feature

```gherkin
Feature: Extended Component Gallery Demo
  Als RFZ-Entwickler oder Designer
  möchte ich alle interaktiven Komponenten in der Gallery sehen,
  damit ich ihre Varianten und Zustände visuell prüfen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Gallery zeigt alle 10 Komponenten-Sektionen

```gherkin
Scenario: Gallery enthält alle Komponenten
  Given ich starte die Component Gallery
  When ich durch die Gallery scrolle
  Then sehe ich Sektionen für alle 10 Komponenten:
    | TuiBox       |
    | TuiDivider   |
    | TuiButton    |
    | TuiStatus    |
    | TuiCheckbox  |
    | TuiRadio     |
    | TuiList      |
    | TuiTextInput |
    | TuiSpinner   |
    | TuiProgress  |
```

### Szenario 2: TuiCheckbox Sektion

```gherkin
Scenario: Checkbox-Sektion zeigt alle Zustände
  Given ich bin in der TuiCheckbox-Sektion
  When ich die Sektion betrachte
  Then sehe ich Beispiele für:
    | unchecked |
    | checked   |
    | focused   |
    | disabled  |
```

### Szenario 3: TuiRadio Sektion

```gherkin
Scenario: Radio-Sektion zeigt alle Varianten
  Given ich bin in der TuiRadio-Sektion
  When ich die Sektion betrachte
  Then sehe ich Beispiele für:
    | horizontal layout |
    | vertical layout   |
    | focused state     |
```

### Szenario 4: TuiList Sektion

```gherkin
Scenario: List-Sektion zeigt beide Modi
  Given ich bin in der TuiList-Sektion
  When ich die Sektion betrachte
  Then sehe ich Beispiele für:
    | multi-select mode  |
    | single-select mode |
    | with badges        |
```

### Szenario 5: TuiSpinner Sektion

```gherkin
Scenario: Spinner-Sektion zeigt alle Varianten
  Given ich bin in der TuiSpinner-Sektion
  When ich die Sektion betrachte
  Then sehe ich statische Darstellungen für:
    | braille dots   |
    | line           |
    | circle quarters|
    | bounce         |
```

### Szenario 6: TuiProgress Sektion

```gherkin
Scenario: Progress-Sektion zeigt alle Stile
  Given ich bin in der TuiProgress-Sektion
  When ich die Sektion betrachte
  Then sehe ich Beispiele bei verschiedenen Füllständen:
    | 0%   - leer    |
    | 50%  - halb    |
    | 100% - voll    |
```

### Szenario 7: Gallery Navigation funktioniert

```gherkin
Scenario: Durch Gallery scrollen
  Given ich habe die Gallery gestartet
  When ich j/k oder Pfeiltasten drücke
  Then scrollt der Viewport durch alle Sektionen
  And ich kann alle Komponenten erreichen
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/demo/gallery.go (modified)

### Inhalt-Prüfungen

- [ ] CONTAINS: gallery.go enthält "renderCheckboxSection"
- [ ] CONTAINS: gallery.go enthält "renderRadioSection"
- [ ] CONTAINS: gallery.go enthält "renderListSection"
- [ ] CONTAINS: gallery.go enthält "renderTextInputSection"
- [ ] CONTAINS: gallery.go enthält "renderSpinnerSection"
- [ ] CONTAINS: gallery.go enthält "renderProgressSection"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./cmd/rfz/...
- [ ] TEST_PASS: go test ./internal/ui/components/demo/... -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/demo/...

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** READY - Alle DoR-Kriterien erfüllt

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
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/demo/gallery.go` | Add 6 new render sections for interactive components |

**Integration:** All 6 new components (TuiCheckbox, TuiRadio, TuiList, TuiTextInput, TuiSpinner, TuiProgress) showcased in gallery

---

### Technical Details

**WAS:**
- Add renderCheckboxSection() showing unchecked/checked/focused/disabled states
- Add renderRadioSection() showing horizontal/vertical layouts and selection states
- Add renderListSection() showing multi-select and single-select modes with badges
- Add renderTextInputSection() showing empty/focused/filled/disabled states
- Add renderSpinnerSection() showing all 4 variants (braille/line/circle/bounce) with static frames
- Add renderProgressSection() showing 0%/50%/100% fill states
- Integrate all new sections into main gallery render loop

**WIE (Architektur-Guidance ONLY):**
- Follow existing gallery section pattern (see renderBoxSection, renderButtonSection, etc.)
- Each section function returns string with header + examples
- Use existing StyleH2 for section headers
- Use lipgloss.JoinVertical for stacking examples
- For spinners, use static TuiSpinnerStatic() to show frame samples (no animation in gallery)
- For progress, render at fixed percentages (0, 25, 50, 75, 100)
- Maintain consistent spacing with existing sections

**WO:**
- Modify: `internal/ui/components/demo/gallery.go` (~200 LOC additions)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** INTER-001, INTER-002, INTER-003, INTER-004, INTER-005, INTER-006

**Geschätzte Komplexität:** S (2 SP)

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No project-specific skills required |

---

### Completion Check

```bash
# Build validation
go build ./cmd/rfz/...

# Run gallery demo tests
go test ./internal/ui/components/demo/... -v

# Lint check
golangci-lint run ./internal/ui/components/demo/...

# Verify all sections exist
grep -q "renderCheckboxSection" internal/ui/components/demo/gallery.go
grep -q "renderRadioSection" internal/ui/components/demo/gallery.go
grep -q "renderListSection" internal/ui/components/demo/gallery.go
grep -q "renderTextInputSection" internal/ui/components/demo/gallery.go
grep -q "renderSpinnerSection" internal/ui/components/demo/gallery.go
grep -q "renderProgressSection" internal/ui/components/demo/gallery.go
```
