# TuiTabs Component

> Story ID: LAYOUT-006
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
Feature: TuiTabs Tab Navigation
  Als RFZ-Entwickler
  möchte ich zwischen verschiedenen Ansichten per Tab wechseln,
  damit ich schnell zwischen Build, Logs, Config navigieren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Horizontale Tab-Leiste

```gherkin
Scenario: Tabs werden horizontal angezeigt
  Given ich habe 4 Tabs ["Build", "Logs", "Discover", "Config"]
  When die Tabs gerendert werden
  Then sehe ich alle 4 Tabs nebeneinander
  And sie sind durch Trennzeichen separiert
```

### Szenario 2: Numerische Shortcuts

```gherkin
Scenario: Tabs haben numerische Shortcuts
  Given ich habe 4 Tabs
  When die Tabs gerendert werden
  Then sehe ich "1" beim ersten Tab
  And ich kann mit Taste "1" direkt zu diesem Tab wechseln
```

### Szenario 3: Aktiver Tab hervorgehoben

```gherkin
Scenario: Aktiver Tab ist visuell hervorgehoben
  Given "Build" ist der aktive Tab
  When die Tabs gerendert werden
  Then ist "Build" mit Hintergrund hervorgehoben
  And andere Tabs haben normalen Stil
```

### Szenario 4: Fokus-State beim Navigieren

```gherkin
Scenario: Fokussierter Tab zeigt Cursor
  Given ich navigiere mit Pfeiltasten durch die Tabs
  When der Fokus auf "Logs" liegt
  Then ist "Logs" mit Unterstreichung oder Rahmen markiert
  And unterscheidet sich vom "aktiven" Tab
```

### Szenario 5: Tab mit Badge/Count

```gherkin
Scenario: Tab zeigt Anzahl von Items
  Given "Logs" Tab hat 5 neue Einträge
  When die Tabs gerendert werden
  Then sehe ich "Logs (5)" oder "Logs •" als Indikator
```

### Szenario 6: Shortcut-Limit bei 9 Tabs

```gherkin
Scenario: Nur erste 9 Tabs haben Shortcuts
  Given ich habe 12 Tabs
  When die Tabs gerendert werden
  Then haben Tabs 1-9 numerische Shortcuts
  And Tabs 10-12 haben keine Shortcuts
```

### Edge Case: Einzelner Tab

```gherkin
Scenario: Nur ein Tab vorhanden
  Given ich habe nur 1 Tab
  When die Tabs gerendert werden
  Then wird der Tab ohne Separator angezeigt
  And er ist automatisch aktiv
```

### Edge Case: Lange Tab-Labels

```gherkin
Scenario: Lange Labels werden gekürzt
  Given ich habe einen Tab "Build Components"
  When wenig Platz verfügbar ist
  Then wird das Label mit "..." gekürzt
  And der Shortcut bleibt sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/tabs.go
- [ ] FILE_EXISTS: internal/ui/components/tabs_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: tabs.go enthält "TuiTabs"
- [ ] CONTAINS: tabs.go enthält "active" oder "Active"

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestTabs -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/tabs.go

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
| Frontend | `internal/ui/components/tabs.go` | New file: TuiTabs + TuiTab |
| Frontend | `internal/ui/components/tabs_test.go` | New file: Unit tests with golden files |
| Frontend | `internal/ui/components/styles.go` | Add StyleTabActive, StyleTabFocused, StyleTabNormal |

---

### Technical Details

**WAS:**
- `TuiTab`: Struct with Label, Badge (optional count/indicator), Active state
- `TuiTabs`: Renders horizontal tab bar with numeric shortcuts (1-9)
- States: normal, focused (keyboard navigation), active (current view)
- Numeric shortcuts: tabs 1-9 have shortcuts, tabs 10+ do not

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (stateless render function)
- Use `lipgloss.JoinHorizontal` to combine tabs
- Separator between tabs: pipe "|" or space with muted color
- Active tab: `ColorCyan` background or underline
- Focused tab: distinct from active (e.g., bold text or border)
- Normal tab: `ColorTextSecondary`
- Badge format: "Label (N)" or "Label ." with `ColorTextMuted`
- Use `Truncate()` helper for long labels
- Reference `design-system.md` Footer/StatusBar patterns

**WO:**
- `internal/ui/components/tabs.go` - Main component implementation
- `internal/ui/components/tabs_test.go` - Unit tests with golden files
- `internal/ui/components/styles.go` - Add tab-specific styles
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
go test ./internal/ui/components/... -run TestTabs -v

# Lint check
golangci-lint run ./internal/ui/components/tabs.go

# Verify TuiTabs function exists
grep -q "TuiTabs\|NewTuiTabs" internal/ui/components/tabs.go

# Verify active state handling exists
grep -q "active\|Active" internal/ui/components/tabs.go
```
