# Finalize and Create PR

> Story ID: LAYOUT-999
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Release
**Estimated Effort**: XS (1 SP)
**Dependencies**: LAYOUT-998

---

## Feature

```gherkin
Feature: Finalize and Create PR
  Als Tech Lead
  moechte ich einen PR mit allen Aenderungen erstellen,
  damit die Layout-Komponenten in den main Branch gemerged werden koennen.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Test Scenarios dokumentiert

```gherkin
Scenario: Manuelle Test-Szenarien dokumentiert
  Given alle Stories sind implementiert
  When ich den PR erstelle
  Then enthaelt der PR Test-Szenarien fuer:
    | Scenario |
    | Navigation Demo: Navigate with j/k |
    | Modal Demo: Open/close with Enter/Escape |
    | Tree Demo: Expand/collapse nodes |
    | Table Demo: Scroll and select rows |
```

### Szenario 2: User TODOs dokumentiert

```gherkin
Scenario: Manuelle Pruefschritte fuer User
  Given der PR ist erstellt
  When der User den PR reviewt
  Then gibt es klare TODOs fuer manuelle Pruefung:
    | TODO |
    | Run layout-demo and verify all components visible |
    | Test keyboard navigation |
    | Verify visual styling matches design system |
```

### Szenario 3: PR erstellt

```gherkin
Scenario: PR wird mit gh CLI erstellt
  Given alle Quality Gates sind erfuellt
  When ich "gh pr create" ausfuehre
  Then wird ein PR mit aussagekraeftigem Titel erstellt
  And die Description enthaelt Summary und Test Plan
```

---

## Technische Verifikation (Automated Checks)

### Pre-PR Checks

- [ ] All integration tests from LAYOUT-998 passing
- [ ] Code review from LAYOUT-997 approved
- [ ] No uncommitted changes

### PR Content

- [ ] PR title follows convention: "feat(components): Add layout and navigation components"
- [ ] PR description includes Summary section
- [ ] PR description includes Test Plan section
- [ ] PR description includes link to spec

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Ready

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
- [x] Kritische Integration Points dokumentiert (wenn Full-stack)

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] PR erstellt mit gh CLI
- [ ] PR title und description korrekt

#### Qualitaetssicherung
- [ ] Test scenarios dokumentiert
- [ ] User TODOs dokumentiert

#### Dokumentation
- [ ] PR URL dokumentiert
- [ ] Spec als "Merged" markiert (nach Merge)

---

### Betroffene Layer & Komponenten

**Integration Type:** Release

**Betroffene Komponenten:**

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Release | PR creation | Create and document PR |

---

### Technical Details

**WAS:**
- Erstellen eines Pull Requests mit allen Aenderungen
- Dokumentation von Test-Szenarien und User TODOs
- PR Description mit Summary und Test Plan

**WIE (Architektur-Guidance ONLY):**
- Branch erstellen falls noch nicht vorhanden: `feature/layout-navigation`
- Alle Aenderungen committen
- PR erstellen mit `gh pr create`
- PR Template:
  ```
  ## Summary
  - Add 7 new layout/navigation components (TuiNavigation, TuiModal, TuiKeyHints, TuiTable, TuiTree, TuiTabs, TuiStatusBar)
  - Add layout-demo program for component showcase
  - Add visual regression tests with golden files

  ## Test Plan

  ### Automated Tests
  - [ ] `go test ./internal/ui/components/... -v` - All tests pass
  - [ ] `golangci-lint run ./...` - No lint errors

  ### Manual Tests
  - [ ] Run `go run ./cmd/layout-demo/` and verify:
    - [ ] Navigation: j/k moves cursor, Enter selects
    - [ ] Modal: Opens with shortcut, closes with Escape
    - [ ] Tree: Enter toggles expand/collapse
    - [ ] Tabs: 1-9 switches tabs, arrows navigate
    - [ ] StatusBar: Shows hints and status
  - [ ] Visual styling matches design-system.md

  Spec: agent-os/specs/2026-02-06-layout-navigation/spec.md
  ```

**WO:**
- PR auf GitHub

**WER:** dev-team__frontend-developer (oder wer auch immer den letzten Commit macht)

**Abhaengigkeiten:** LAYOUT-998 (Integration Validation muss erfolgreich sein)

**Geschaetzte Komplexitaet:** XS

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No special skills required |

---

### Completion Check

```bash
# Verify clean git status
git status

# Verify all tests pass one more time
go test ./internal/ui/components/... -v

# PR exists (after creation)
gh pr list --head feature/layout-navigation

# Mark spec as merged (after PR is merged)
# Update spec.md: Status: Merged
```

---

## Manual Test Scenarios

### 1. Navigation Demo
```
1. Run: go run ./cmd/layout-demo/
2. Verify navigation sidebar is visible on left
3. Press j/k to move cursor through items
4. Verify cursor indicator (>) moves
5. Verify focused item is cyan
6. Press number keys 1-4 to jump to items
```

### 2. Modal Demo
```
1. In layout-demo, press 'm' or designated key to open modal
2. Verify modal appears centered with double border
3. Verify backdrop dims the background
4. Press Tab to cycle through buttons
5. Press Escape to close modal
```

### 3. Tree Demo
```
1. Navigate to tree section in demo
2. Verify tree shows hierarchy with indentation
3. Verify collapsed nodes show >
4. Press Enter on a node with children
5. Verify it expands and shows <
6. Verify children become visible
```

### 4. Table Demo
```
1. Navigate to table section in demo
2. Verify column headers are visible
3. Press j/k to move through rows
4. Verify selected row is highlighted
5. Press Enter to select a row
```

---

## User TODOs After PR Merge

- [ ] Update roadmap.md to mark Sprint 1.3 as complete
- [ ] Update spec.md status to "Merged"
- [ ] Consider adding layout components to existing Component Gallery
- [ ] Plan integration with actual screens (Phase 2)
