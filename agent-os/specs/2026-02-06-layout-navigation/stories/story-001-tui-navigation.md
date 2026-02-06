# TuiNavigation + TuiNavItem Components

> Story ID: LAYOUT-001
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiNavigation Sidebar Container
  Als RFZ-Entwickler
  möchte ich eine vertikale Navigation im Sidebar-Stil,
  damit ich zwischen verschiedenen Screens navigieren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Navigation zeigt Items

```gherkin
Scenario: Navigation zeigt alle Menu-Items
  Given ich habe eine TuiNavigation mit 4 Items
  When die Navigation gerendert wird
  Then sehe ich alle 4 Items untereinander
  And jedes Item zeigt Nummer und Label
```

### Szenario 2: Aktives Item hervorgehoben

```gherkin
Scenario: Aktives Item ist visuell hervorgehoben
  Given ich bin auf dem "Build" Screen
  When die Navigation gerendert wird
  Then ist "Build" mit Hintergrundfarbe hervorgehoben
  And andere Items haben normalen Stil
```

### Szenario 3: Fokussiertes Item mit Cursor

```gherkin
Scenario: Fokussiertes Item zeigt Cursor
  Given ich navigiere mit j/k durch die Navigation
  When der Cursor auf "Logs" steht
  Then sehe ich "> " vor dem "Logs" Item
  And das Item ist in Cyan hervorgehoben
```

### Szenario 4: Keyboard Shortcuts angezeigt

```gherkin
Scenario: Items zeigen Keyboard Shortcuts
  Given ich habe TuiNavItems mit Shortcuts "1", "2", "3"
  When die Navigation gerendert wird
  Then sehe ich die Shortcuts rechts neben den Labels
  And sie sind in gedämpfter Farbe dargestellt
```

### Szenario 5: Navigation mit Header

```gherkin
Scenario: Navigation zeigt optionalen Header
  Given ich habe eine TuiNavigation mit Header "Navigation"
  When die Navigation gerendert wird
  Then sehe ich "Navigation" über den Items
  And der Header ist vom Inhalt getrennt
```

### Szenario 6: Navigation mit Footer

```gherkin
Scenario: Navigation zeigt optionalen Footer
  Given ich habe eine TuiNavigation mit Footer für KeyHints
  When die Navigation gerendert wird
  Then sehe ich den Footer unter den Items
  And der Footer ist durch eine Linie getrennt
```

### Edge Case: Lange Labels

```gherkin
Scenario: Lange Labels werden gekürzt
  Given ich habe ein Item mit Label "Build Components and Deploy"
  When die Navigation eine feste Breite hat
  Then wird das Label mit "..." gekürzt
  And die Nummer und der Shortcut bleiben sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/navigation.go
- [ ] FILE_EXISTS: internal/ui/components/navigation_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: navigation.go enthält "TuiNavigation"
- [ ] CONTAINS: navigation.go enthält "TuiNavItem"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestNav -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/navigation.go

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
| Frontend | `internal/ui/components/navigation.go` | New file: TuiNavigation + TuiNavItem |
| Frontend | `internal/ui/components/navigation_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add navigation-specific symbols if needed |

**Integration:** TuiNavigation enthält N TuiNavItem-Elemente (Composition)

---

### Technical Details

**WAS:**
- `TuiNavItem`: Renders a single navigation menu item with number, label, and keyboard shortcut
- `TuiNavigation`: Renders a complete sidebar navigation with optional header/footer, containing multiple TuiNavItem elements
- States: normal, focused (cursor), active (current screen)
- Support for keyboard shortcuts display (muted color on right)

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (stateless render functions, state managed by parent)
- Reference existing patterns from `list.go` for item rendering with cursor/selection states
- Use existing styles: `StyleNavItem`, `StyleNavItemFocused`, `StyleNavItemActive` from `styles.go`
- Use `Truncate()` helper from `helpers.go` for long labels
- Use `lipgloss.JoinVertical` for stacking nav items
- Header/footer separated by `TuiDivider` (from `divider.go`)
- Shortcuts rendered with `StyleKeyboard` or `ColorTextMuted`

**WO:**
- `internal/ui/components/navigation.go` - Main component implementation
- `internal/ui/components/navigation_test.go` - Unit tests with golden files
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** None

**Geschätzte Komplexität:** S

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
go test ./internal/ui/components/... -run TestNav -v

# Lint check
golangci-lint run ./internal/ui/components/navigation.go

# Verify TuiNavigation function exists
grep -q "func TuiNavigation" internal/ui/components/navigation.go

# Verify TuiNavItem function exists
grep -q "func TuiNavItem" internal/ui/components/navigation.go
```
