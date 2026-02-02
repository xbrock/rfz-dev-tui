# Component Gallery

> Story ID: CORE-007
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: Medium
**Type**: Backend
**Estimated Effort**: S (3 SP)
**Dependencies**: CORE-006

---

## Feature

```gherkin
Feature: Component Gallery Demo Screen
  Als TUI-Entwickler
  möchte ich eine Demo-Ansicht aller Komponenten haben,
  damit ich alle Varianten visuell verifizieren und präsentieren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Gallery zeigt alle Komponenten

```gherkin
Scenario: Component Gallery zeigt alle verfügbaren Komponenten
  Given ich starte die Component Gallery
  When die Ansicht geladen wird
  Then sehe ich Sektionen für TuiBox, TuiDivider, TuiButton und TuiStatus
  And jede Sektion zeigt alle Varianten der Komponente
```

### Szenario 2: TuiBox Varianten sichtbar

```gherkin
Scenario: Alle TuiBox Varianten werden angezeigt
  Given ich bin in der Component Gallery
  When ich die TuiBox Sektion betrachte
  Then sehe ich Single, Double, Rounded, Heavy Border Varianten
  And ich sehe die Focused-State Variante
```

### Szenario 3: TuiButton Varianten sichtbar

```gherkin
Scenario: Alle TuiButton Varianten werden angezeigt
  Given ich bin in der Component Gallery
  When ich die TuiButton Sektion betrachte
  Then sehe ich Primary, Secondary, Destructive Varianten
  And ich sehe Buttons mit und ohne Keyboard Shortcuts
```

### Szenario 4: TuiStatus Varianten sichtbar

```gherkin
Scenario: Alle TuiStatus Varianten werden angezeigt
  Given ich bin in der Component Gallery
  When ich die TuiStatus Sektion betrachte
  Then sehe ich Pending, Running, Success, Failed, Error Badges
  And ich sehe die Compact-Varianten
```

### Szenario 5: Keyboard Navigation

```gherkin
Scenario: Navigation durch die Gallery mit Tastatur
  Given ich bin in der Component Gallery
  When ich die Pfeiltasten oder j/k benutze
  Then scrolle ich durch die verschiedenen Sektionen
  And die aktuelle Sektion ist visuell hervorgehoben
```

### Edge Case: Terminal Größe

```gherkin
Scenario: Gallery passt sich an Terminal-Größe an
  Given mein Terminal ist 80x24
  When die Gallery gerendert wird
  Then werden Komponenten angemessen angeordnet
  And nichts wird abgeschnitten
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/demo/gallery.go
- [ ] FILE_EXISTS: internal/ui/components/demo/gallery_test.go
- [ ] FILE_EXISTS: testdata/golden/gallery/

### Inhalt-Prüfungen

- [ ] CONTAINS: gallery.go enthält "type Gallery struct"
- [ ] CONTAINS: gallery.go enthält "func (g Gallery) Update("
- [ ] CONTAINS: gallery.go enthält "func (g Gallery) View("
- [ ] CONTAINS: gallery.go enthält "TuiBox"
- [ ] CONTAINS: gallery.go enthält "TuiButton"
- [ ] CONTAINS: gallery.go enthält "TuiStatus"
- [ ] CONTAINS: gallery.go enthält "TuiDivider"

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/demo/...`
- [ ] TEST_PASS: `go test ./internal/ui/components/demo/... -v`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/demo/`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

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
- [x] Handover-Dokumente definiert (bei Multi-Layer)

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
- [ ] Dokumentation aktualisiert
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Backend-only

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Backend | internal/ui/components/demo/gallery.go | CREATE - Gallery screen |
| Test | internal/ui/components/demo/gallery_test.go | CREATE - Gallery tests |
| Test | testdata/golden/gallery/ | CREATE - Gallery golden files |

**Kritische Integration Points:**
- gallery.go → box.go: Rendert TuiBox Beispiele
- gallery.go → button.go: Rendert TuiButton Beispiele
- gallery.go → status.go: Rendert TuiStatus Beispiele
- gallery.go → divider.go: Rendert TuiDivider Beispiele
- gallery.go → styles.go: Verwendet alle Styles für Layout

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- Gallery als Bubble Tea Model
- Zeigt alle Komponenten in Sektionen
- Keyboard Navigation zwischen Sektionen
- Golden File Tests für Gallery

**WIE:**
- Implementiert tea.Model Interface (Init, Update, View):
  ```go
  type Gallery struct {
      viewport viewport.Model
      sections []string
      focused  int
      width    int
      height   int
  }

  func (g Gallery) Init() tea.Cmd
  func (g Gallery) Update(msg tea.Msg) (tea.Model, tea.Cmd)
  func (g Gallery) View() string
  ```
- Verwendet bubbles/viewport für Scrolling (for content overflow)
- Sektionen mit TuiBox umrahmt:
  - Section 1: TuiBox variants (Single, Double, Rounded, Heavy, Focused)
  - Section 2: TuiDivider variants (Single, Double)
  - Section 3: TuiButton variants (Primary, Secondary, Destructive + shortcuts + focus)
  - Section 4: TuiStatus variants (all 5 states + compact)
- Keyboard bindings:
  - j/down: scroll down
  - k/up: scroll up
  - q/ctrl+c: quit
- Layout: Sections stacked vertically, separated by TuiDivider
- Each section title uses StyleH2 from styles.go

**WO:**
- `internal/ui/components/demo/gallery.go` (NEW) - ~200 LOC
- `internal/ui/components/demo/gallery_test.go` (NEW)
- `testdata/golden/gallery/` (NEW directory)

**Abhängigkeiten:** CORE-006 (requires working test infrastructure and all components)

**Geschätzte Komplexität:** S (3 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify files exist
test -f internal/ui/components/demo/gallery.go && echo "gallery.go exists"
test -f internal/ui/components/demo/gallery_test.go && echo "gallery_test.go exists"
test -d testdata/golden/gallery && echo "golden/gallery exists"

# Verify Gallery struct and methods
grep -q "type Gallery struct" internal/ui/components/demo/gallery.go && echo "Gallery struct found"
grep -q "func (g Gallery) View" internal/ui/components/demo/gallery.go && echo "View method found"

# Verify components are used
grep -q "TuiBox" internal/ui/components/demo/gallery.go && echo "TuiBox used"
grep -q "TuiButton" internal/ui/components/demo/gallery.go && echo "TuiButton used"
grep -q "TuiStatus" internal/ui/components/demo/gallery.go && echo "TuiStatus used"

# Build and test
go build ./internal/ui/components/demo/...
go test ./internal/ui/components/demo/... -v
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Gallery zeigt alle 4 Komponenten
3. BUILD_PASS und TEST_PASS erfolgreich
