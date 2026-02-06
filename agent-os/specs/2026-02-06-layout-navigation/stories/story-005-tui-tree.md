# TuiTree Component

> Story ID: LAYOUT-005
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: M (3 SP)
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiTree Hierarchical View
  Als RFZ-Entwickler
  möchte ich Komponenten-Abhängigkeiten als Baumstruktur sehen,
  damit ich die Beziehungen zwischen Komponenten verstehe.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Baumstruktur mit Einrückung

```gherkin
Scenario: Tree zeigt Hierarchie mit Indentation
  Given ich habe einen Baum mit 3 Ebenen
  When der Baum gerendert wird
  Then sind Kinder-Nodes eingerückt
  And jede Ebene hat mehr Einrückung als die vorherige
```

### Szenario 2: Expand/Collapse Icons

```gherkin
Scenario: Nodes zeigen Expand/Collapse Icons
  Given ich habe einen Node mit Kindern
  When der Node zugeklappt ist
  Then sehe ich "▶" vor dem Node-Label
  When der Node aufgeklappt ist
  Then sehe ich "▼" vor dem Node-Label
```

### Szenario 3: Keyboard Navigation

```gherkin
Scenario: Navigation mit j/k durch den Baum
  Given ich habe einen Baum mit 10 Nodes
  When ich j drücke
  Then bewegt sich der Cursor nach unten
  When ich k drücke
  Then bewegt sich der Cursor nach oben
```

### Szenario 4: Toggle Expand mit Enter

```gherkin
Scenario: Enter klappt Node auf/zu
  Given ich bin auf einem Node mit Kindern
  When ich Enter drücke
  Then wechselt der Node zwischen expand/collapse
  And die Kinder werden sichtbar/unsichtbar
```

### Szenario 5: Blatt-Nodes ohne Icon

```gherkin
Scenario: Nodes ohne Kinder zeigen kein Expand-Icon
  Given ich habe einen Leaf-Node ohne Kinder
  When der Node gerendert wird
  Then wird kein ▶/▼ Icon angezeigt
  And der Node ist entsprechend eingerückt
```

### Szenario 6: Node-Labels mit Metadata

```gherkin
Scenario: Nodes zeigen zusätzliche Informationen
  Given ich habe einen Node "boss" mit Status "clean"
  When der Node gerendert wird
  Then sehe ich "boss" als Label
  And "clean" als Metadata rechts daneben
```

### Edge Case: Leerer Baum

```gherkin
Scenario: Baum ohne Nodes
  Given ich habe einen leeren Baum
  When der Baum gerendert wird
  Then sehe ich "No items" als Hinweis
```

### Edge Case: Tiefe Hierarchie

```gherkin
Scenario: Baum mit vielen Ebenen
  Given ich habe einen Baum mit 15 Ebenen
  When die Tiefe das Limit überschreitet
  Then werden tiefere Ebenen mit "..." angedeutet
  And die UI bleibt performant
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/tree.go
- [ ] FILE_EXISTS: internal/ui/components/tree_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: tree.go enthält "TuiTree"
- [ ] CONTAINS: tree.go enthält "▶" oder "▼" oder "expand"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestTree -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/tree.go

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Ready

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
| Frontend | `internal/ui/components/tree.go` | New file: TuiTree + TuiTreeNode |
| Frontend | `internal/ui/components/tree_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add SymbolExpanded, SymbolCollapsed constants |

---

### Technical Details

**WAS:**
- `TuiTreeNode`: Struct representing a tree node with Label, Metadata, Children, Expanded state
- `TuiTree`: Renders a hierarchical tree view with expand/collapse icons
- `TuiTreeItem`: Renders a single node with proper indentation
- Support for cursor navigation, expand/collapse toggle
- Leaf nodes (no children) show no expand icon

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (stateless render functions, tree state managed externally)
- Add symbols to `styles.go`: `SymbolExpanded = ""`  `SymbolCollapsed = ""`
- Indentation: multiply depth by constant (e.g., 2 spaces per level)
- Reference `list.go` for cursor/focus state rendering pattern
- Recursive rendering or flatten tree to visible nodes list
- Use `ColorCyan` for focused node, `ColorTextPrimary` for normal
- Metadata rendered with `ColorTextMuted` right of label
- Handle deep hierarchies: limit depth or show "..." indicator

**WO:**
- `internal/ui/components/tree.go` - Main component implementation
- `internal/ui/components/tree_test.go` - Unit tests with golden files
- `internal/ui/components/styles.go` - Add tree symbols
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** None

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
go test ./internal/ui/components/... -run TestTree -v

# Lint check
golangci-lint run ./internal/ui/components/tree.go

# Verify TuiTree function exists
grep -q "TuiTree\|NewTuiTree" internal/ui/components/tree.go

# Verify expand/collapse symbols exist
grep -q "▶\|▼\|Expand\|Collapse" internal/ui/components/tree.go
```
