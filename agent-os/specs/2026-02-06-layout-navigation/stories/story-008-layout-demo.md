# Layout Navigation Demo

> Story ID: LAYOUT-008
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: LAYOUT-001 to LAYOUT-007

---

## Feature

```gherkin
Feature: Layout Navigation Demo Program
  Als RFZ-Entwickler
  möchte ich alle Layout-Komponenten in einer Demo sehen,
  damit ich ihre Funktionsweise verstehe und testen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Demo startet

```gherkin
Scenario: Demo-Programm startet erfolgreich
  Given ich führe das Demo-Programm aus
  When das Programm startet
  Then sehe ich den Demo-Screen
  And alle Komponenten sind sichtbar
```

### Szenario 2: Alle 8 Komponenten angezeigt

```gherkin
Scenario: Demo zeigt alle Layout-Komponenten
  Given ich bin im Demo-Programm
  When ich durch die Demo scrolle
  Then sehe ich TuiNavigation mit Items
  And sehe ich TuiModal
  And sehe ich TuiKeyHints
  And sehe ich TuiTable
  And sehe ich TuiTree
  And sehe ich TuiTabs
  And sehe ich TuiStatusBar
```

### Szenario 3: Interaktive Navigation

```gherkin
Scenario: Navigation ist interaktiv
  Given ich sehe die TuiNavigation in der Demo
  When ich mit j/k navigiere
  Then bewegt sich der Cursor durch die Items
  And das fokussierte Item ist hervorgehoben
```

### Szenario 4: Modal öffnen/schließen

```gherkin
Scenario: Modal kann geöffnet werden
  Given ich bin in der Demo
  When ich "m" oder Enter drücke (je nach Kontext)
  Then öffnet sich ein Demo-Modal
  When ich Escape drücke
  Then schließt sich das Modal
```

### Szenario 5: Tree Expand/Collapse

```gherkin
Scenario: Tree Nodes können auf/zugeklappt werden
  Given ich sehe den TuiTree in der Demo
  When ich einen Node mit Kindern fokussiere
  And Enter drücke
  Then klappt der Node auf oder zu
```

### Szenario 6: Scroll durch alle Sektionen

```gherkin
Scenario: Demo ist scrollbar
  Given die Demo hat mehr Inhalt als das Terminal
  When ich mit j/k scrolle
  Then bewegt sich der Viewport
  And ich kann alle Komponenten erreichen
```

### Szenario 7: Quit mit q

```gherkin
Scenario: Demo kann beendet werden
  Given ich bin im Demo-Programm
  When ich "q" drücke
  Then wird das Programm beendet
  And ich bin zurück im Terminal
```

### Edge Case: Kleines Terminal

```gherkin
Scenario: Demo passt sich an Terminalgröße an
  Given ich habe ein 80x24 Terminal
  When die Demo gerendert wird
  Then passt sich der Inhalt an
  And bleibt lesbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/demo/layout_gallery.go
- [ ] FILE_EXISTS: cmd/layout-demo/main.go

### Inhalt-Prüfungen

- [ ] CONTAINS: layout_gallery.go enthält "TuiNavigation"
- [ ] CONTAINS: layout_gallery.go enthält "TuiModal"
- [ ] CONTAINS: layout_gallery.go enthält "TuiTree"
- [ ] CONTAINS: layout_gallery.go enthält "TuiTabs"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./cmd/layout-demo/...
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/demo/layout_gallery.go

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
- [x] Demo program runs without errors
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
| Frontend | `internal/ui/components/demo/layout_gallery.go` | New file: Layout Demo Bubble Tea Model |
| Frontend | `cmd/layout-demo/main.go` | New file: Demo entry point |

**Integration:** Verwendet alle 7 Layout-Komponenten (Navigation, Modal, KeyHints, Table, Tree, Tabs, StatusBar)

---

### Technical Details

**WAS:**
- `LayoutGalleryModel`: Bubble Tea Model showcasing all layout components
- Interactive demo with keyboard navigation
- Modal open/close, tree expand/collapse, tab switching
- Entry point in `cmd/layout-demo/main.go`

**WIE (Architektur-Guidance ONLY):**
- Pattern: Full Bubble Tea application with Model/Update/View
- Reference existing demo pattern in `internal/ui/components/demo/gallery_test.go`
- Create sections for each component type
- State management:
  - `navCursor` for TuiNavigation
  - `modalOpen` for TuiModal
  - `treeNodes` with expanded states for TuiTree
  - `activeTab` for TuiTabs
- Key bindings: j/k for navigation, Enter for select/toggle, Escape for close, q for quit
- Use `bubbles/viewport` for scrollable content if needed
- Layout: TuiNavigation on left, component showcase on right, TuiStatusBar at bottom

**WO:**
- `internal/ui/components/demo/layout_gallery.go` - Demo Model implementation
- `cmd/layout-demo/main.go` - Entry point with `tea.NewProgram()`

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** LAYOUT-001 to LAYOUT-007 (alle Komponenten muessen implementiert sein)

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
go build ./cmd/layout-demo/...

# Lint check
golangci-lint run ./internal/ui/components/demo/layout_gallery.go
golangci-lint run ./cmd/layout-demo/main.go

# Verify all components are used
grep -q "TuiNavigation" internal/ui/components/demo/layout_gallery.go
grep -q "TuiModal" internal/ui/components/demo/layout_gallery.go
grep -q "TuiTree" internal/ui/components/demo/layout_gallery.go
grep -q "TuiTabs" internal/ui/components/demo/layout_gallery.go
grep -q "TuiStatusBar" internal/ui/components/demo/layout_gallery.go
grep -q "TuiKeyHints" internal/ui/components/demo/layout_gallery.go
grep -q "TuiTable" internal/ui/components/demo/layout_gallery.go
```
