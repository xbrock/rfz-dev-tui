# TuiModal Component

> Story ID: LAYOUT-002
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Frontend
**Estimated Effort**: M (3 SP)
**Dependencies**: None (nutzt TuiButton aus Sprint 1.1)

---

## Feature

```gherkin
Feature: TuiModal Overlay Dialog
  Als RFZ-Entwickler
  möchte ich ein Modal-Dialog für Build-Konfiguration,
  damit ich Einstellungen in einem fokussierten Overlay vornehmen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Modal mit Double Border

```gherkin
Scenario: Modal zeigt prominenten Double Border
  Given ich öffne ein Modal
  When das Modal gerendert wird
  Then sehe ich einen doppelten Rahmen (╔═══╗)
  And der Rahmen hebt sich vom Hintergrund ab
```

### Szenario 2: Modal zentriert

```gherkin
Scenario: Modal ist horizontal und vertikal zentriert
  Given ich habe ein Modal mit 40x15 Größe
  When das Modal im 120x40 Terminal gerendert wird
  Then ist das Modal mittig positioniert
  And der Abstand zu allen Seiten ist gleich
```

### Szenario 3: Title Bar

```gherkin
Scenario: Modal zeigt Titel
  Given ich habe ein Modal mit Titel "Build Configuration"
  When das Modal gerendert wird
  Then sehe ich "Build Configuration" im Titel-Bereich
  And der Titel ist vom Content getrennt
```

### Szenario 4: Footer mit Buttons

```gherkin
Scenario: Modal zeigt Action-Buttons im Footer
  Given ich habe ein Modal mit "Start" und "Cancel" Buttons
  When das Modal gerendert wird
  Then sehe ich beide Buttons im Footer-Bereich
  And "Start" ist als Primary Button hervorgehoben
```

### Szenario 5: Escape schließt Modal

```gherkin
Scenario: Modal schließt bei Escape
  Given ich habe ein offenes Modal
  When ich Escape drücke
  Then wird das Modal geschlossen
  And der Fokus kehrt zum vorherigen Element zurück
```

### Szenario 6: Tab cycled durch Buttons

```gherkin
Scenario: Tab wechselt zwischen Buttons
  Given ich habe ein Modal mit 2 Buttons
  When ich Tab drücke
  Then wechselt der Fokus zum nächsten Button
  And der fokussierte Button ist hervorgehoben
```

### Szenario 7: Backdrop dimmt Hintergrund

```gherkin
Scenario: Modal hat gedimmten Backdrop
  Given ich öffne ein Modal
  When das Modal gerendert wird
  Then ist der Hintergrund außerhalb des Modals gedimmt
  And nur das Modal ist klar sichtbar
```

### Edge Case: Modal ohne Buttons

```gherkin
Scenario: Modal ohne Buttons zeigt nur Content
  Given ich habe ein Modal ohne Buttons definiert
  When das Modal gerendert wird
  Then wird kein Footer angezeigt
  And Escape schließt trotzdem das Modal
```

### Edge Case: Scrollbarer Content

```gherkin
Scenario: Langer Content ist scrollbar
  Given ich habe ein Modal mit viel Content
  When der Content die Höhe überschreitet
  Then kann ich mit j/k durch den Content scrollen
  And der Footer bleibt fixiert
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [x] FILE_EXISTS: internal/ui/components/modal.go
- [x] FILE_EXISTS: internal/ui/components/modal_test.go

### Inhalt-Prüfungen

- [x] CONTAINS: modal.go enthält "TuiModal"
- [x] CONTAINS: modal.go enthält "BorderDouble"

### Funktions-Prüfungen

- [x] BUILD_PASS: go build ./internal/ui/components/...
- [x] TEST_PASS: go test ./internal/ui/components/... -run TestModal -v
- [x] LINT_PASS: golangci-lint run ./internal/ui/components/

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Done

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
- [x] Code implementiert und folgt Style Guide
- [x] Architektur-Vorgaben eingehalten
- [x] Security/Performance Anforderungen erfüllt

#### Qualitätssicherung
- [x] Alle Akzeptanzkriterien erfüllt
- [x] Unit Tests geschrieben und bestanden
- [x] Code Review durchgeführt und genehmigt

#### Dokumentation
- [x] Keine Linting Errors
- [x] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/modal.go` | New file: TuiModal component |
| Frontend | `internal/ui/components/modal_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add StyleModalTitle, StyleModalBackdrop if needed |

**Integration:** TuiModal verwendet TuiButton (from `button.go`) for Footer Actions

---

### Technical Details

**WAS:**
- `TuiModal`: Renders a centered overlay dialog with double border
- `TuiModalConfig`: Configuration struct for title, content, buttons, dimensions
- Support for title bar, content area, footer with buttons
- Backdrop rendering for dimmed background effect
- Button focus state tracking for Tab navigation

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (stateless render function with config struct)
- Use `BorderDouble` from `styles.go` for prominent modal border
- Use `lipgloss.Place()` for centering modal in terminal dimensions
- Reference `box.go` pattern for bordered container rendering
- Compose with `TuiButton` from `button.go` for footer actions
- Backdrop: Fill terminal with `ColorBackground` at lower opacity or solid color
- Use `lipgloss.JoinVertical` for stacking title/content/footer
- Divider between sections using `TuiDividerSingle` from `divider.go`

**WO:**
- `internal/ui/components/modal.go` - Main component implementation
- `internal/ui/components/modal_test.go` - Unit tests with golden files
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** None (TuiButton bereits vorhanden in `button.go`)

**Geschätzte Komplexität:** M

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No special skills required |

---

### Completion Check

```bash
# Build validation
go build ./internal/ui/components/...

# Unit tests
go test ./internal/ui/components/... -run TestModal -v

# Lint check
golangci-lint run ./internal/ui/components/modal.go

# Verify TuiModal function exists
grep -q "func TuiModal\|type TuiModal" internal/ui/components/modal.go

# Verify double border is used
grep -q "BorderDouble" internal/ui/components/modal.go
```
