# TuiList Component

> Story ID: INTER-001
> Spec: 2026-02-03-interactive-components
> Created: 2026-02-03
> Last Updated: 2026-02-03

**Priority**: High
**Type**: Frontend
**Estimated Effort**: M (3 SP)
**Dependencies**: INTER-002 (TuiCheckbox), INTER-003 (TuiRadio)

---

## Feature

```gherkin
Feature: TuiList Scrollable Selection Component
  Als RFZ-Entwickler
  möchte ich eine scrollbare Liste mit Single- und Multi-Select-Modi,
  damit ich Komponenten für den Build auswählen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Liste mit Cursor-Navigation

```gherkin
Scenario: Liste zeigt Cursor an aktueller Position
  Given ich habe eine TuiList mit 5 Einträgen
  When der Cursor auf dem 3. Eintrag steht
  Then sehe ich "›" vor dem 3. Eintrag
  And die anderen Einträge haben kein Cursor-Symbol
```

### Szenario 2: Multi-Select Modus mit Checkboxen

```gherkin
Scenario: Multi-Select Liste zeigt Checkboxen
  Given ich habe eine TuiList im Multi-Select-Modus
  When ich "boss" und "fistiv" ausgewählt habe
  Then sehe ich "☑ boss" und "☑ fistiv"
  And die anderen Einträge zeigen "☐"
```

### Szenario 3: Single-Select Modus mit Radio-Buttons

```gherkin
Scenario: Single-Select Liste zeigt Radio-Buttons
  Given ich habe eine TuiList im Single-Select-Modus
  When ich "clean install" ausgewählt habe
  Then sehe ich "◉ clean install"
  And die anderen Einträge zeigen "◯"
```

### Szenario 4: Auswahl-Zähler

```gherkin
Scenario: Liste zeigt Anzahl ausgewählter Einträge
  Given ich habe eine TuiList mit 13 Einträgen
  When 3 Einträge ausgewählt sind
  Then sehe ich "3/13 selected" als Zähler
```

### Szenario 5: Liste mit Badges

```gherkin
Scenario: Listeneinträge zeigen Kategorie-Badges
  Given ich habe eine TuiList mit Komponenten
  When ein Eintrag vom Typ "Core" ist
  Then sehe ich den Badge "Core" rechts neben dem Eintrag
  And der Badge hat eine entsprechende Hintergrundfarbe
```

### Szenario 6: Scrollbare Liste

```gherkin
Scenario: Lange Liste ist scrollbar
  Given ich habe eine TuiList mit 20 Einträgen
  When nur 10 Einträge sichtbar sind
  Then kann ich mit j/k oder Pfeiltasten scrollen
  And der Viewport zeigt immer den aktuellen Cursor
```

### Szenario 7: Fokus-Hervorhebung

```gherkin
Scenario: Fokussierte Liste hat Cyan-Border
  Given ich habe eine TuiList
  When die Liste fokussiert ist
  Then wird der Rahmen in Cyan hervorgehoben
```

### Edge Case: Leere Liste

```gherkin
Scenario: Leere Liste zeigt Hinweistext
  Given ich habe eine TuiList ohne Einträge
  When die Liste gerendert wird
  Then sehe ich "No items" als Hinweistext
  And die Auswahl-Funktion ist deaktiviert
```

### Edge Case: Sehr lange Eintragsnamen

```gherkin
Scenario: Lange Eintragsnamen werden gekürzt
  Given ich habe einen Eintrag "generate_local_config_files_for_development"
  When der Eintrag länger als die verfügbare Breite ist
  Then wird der Text mit "..." gekürzt
  And das Checkbox/Radio-Symbol bleibt sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [x] FILE_EXISTS: internal/ui/components/list.go
- [x] FILE_EXISTS: internal/ui/components/list_test.go

### Inhalt-Prüfungen

- [x] CONTAINS: list.go enthält "SymbolCursor" (uses SymbolListPointer)
- [x] CONTAINS: list.go enthält "MultiSelect" oder "SingleSelect"

### Funktions-Prüfungen

- [x] BUILD_PASS: go build ./internal/ui/components/...
- [x] TEST_PASS: go test ./internal/ui/components/... -run TestList -v
- [x] LINT_PASS: golangci-lint run ./internal/ui/components/list.go

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> **Status:** Done - Implementation complete

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
| Frontend | `internal/ui/components/styles.go` | Add SymbolCursor constant |
| Frontend | `internal/ui/components/list.go` | New file: TuiList and TuiListModel |
| Frontend | `internal/ui/components/list_test.go` | New file: Unit tests for TuiList |

**Integration:** TuiCheckbox + TuiRadio symbols used by TuiList

---

### Technical Details

**WAS:**
- Add cursor symbol to styles.go (U+203A guillemet)
- Create TuiListItem struct for list entries (label, badge, selected)
- Create stateless TuiListItemRender function for single item
- Create TuiListModel wrapping bubbles/viewport for scrollable lists
- Support multi-select mode (checkboxes) and single-select mode (radio)
- Support optional badges per item
- Support selection counter display ("3/13 selected")
- Support focus highlight with cyan border
- Handle empty list with "No items" message

**WIE (Architektur-Guidance ONLY):**
- Hybrid approach: stateless render functions + optional Bubble Tea Model
- Use TuiCheckbox/TuiRadio symbols from styles.go (SymbolCheckboxChecked, SymbolRadioSelected, etc.)
- Use existing Truncate() helper for long labels
- TuiListItem struct: `type TuiListItem struct { Label, Badge string; Selected bool }`
- TuiListItemRender: `TuiListItemRender(item TuiListItem, cursor bool, multiSelect bool, focused bool) string`
- TuiListModel wraps bubbles/viewport for scrolling large lists
- Use existing StyleBoxFocused for focus border
- Apply Lip Gloss for all badge styling (use existing StyleBadgeInfo pattern)

**WO:**
- Modify: `internal/ui/components/styles.go` (add ~2 lines for SymbolCursor)
- Create: `internal/ui/components/list.go` (~180 LOC)
- Create: `internal/ui/components/list_test.go` (~150 LOC)

**WER:** tech-architect (component library, frontend focus)

**Abhängigkeiten:** INTER-002 (TuiCheckbox), INTER-003 (TuiRadio) - uses their symbols

**Geschätzte Komplexität:** M (3 SP)

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No project-specific skills required |

---

### Completion Check

```bash
# Build validation
go build ./internal/ui/components/...

# Unit tests
go test ./internal/ui/components/... -run TestList -v

# Lint check
golangci-lint run ./internal/ui/components/list.go

# Verify cursor symbol exists
grep -q "SymbolCursor" internal/ui/components/styles.go

# Verify list uses checkbox/radio symbols
grep -q "SymbolCheckbox\|SymbolRadio" internal/ui/components/list.go
```
