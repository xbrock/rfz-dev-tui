# TuiStatusBar Component

> Story ID: LAYOUT-007
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Frontend
**Estimated Effort**: S (2 SP)
**Dependencies**: LAYOUT-003 (TuiKeyHints)

---

## Feature

```gherkin
Feature: TuiStatusBar Bottom Bar
  Als RFZ-Entwickler
  möchte ich eine Status-Leiste am unteren Bildschirmrand,
  damit ich den aktuellen Status und verfügbare Shortcuts sehe.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Full-Width Bar

```gherkin
Scenario: StatusBar erstreckt sich über die volle Breite
  Given ich habe ein 120-Zeichen breites Terminal
  When die StatusBar gerendert wird
  Then ist die StatusBar genau 120 Zeichen breit
  And sie hat einen farblichen Hintergrund
```

### Szenario 2: Left Section (Status)

```gherkin
Scenario: StatusBar zeigt Status links
  Given ein Build läuft gerade
  When die StatusBar gerendert wird
  Then sehe ich "Build: Running" im linken Bereich
  And der Status hat eine passende Farbe (gelb)
```

### Szenario 3: Center Section (Info)

```gherkin
Scenario: StatusBar zeigt Info mittig
  Given ich bin auf dem "Build" Screen
  When die StatusBar gerendert wird
  Then sehe ich den Screen-Namen oder eine Info in der Mitte
```

### Szenario 4: Right Section (KeyHints)

```gherkin
Scenario: StatusBar zeigt KeyHints rechts
  Given der aktuelle Screen hat Shortcuts
  When die StatusBar gerendert wird
  Then sehe ich die KeyHints rechtsbündig
  And sie zeigen kontextbezogene Shortcuts
```

### Szenario 5: Separator vom Content

```gherkin
Scenario: StatusBar ist vom Content getrennt
  Given ich habe Content über der StatusBar
  When der Screen gerendert wird
  Then ist eine visuelle Trennung erkennbar
  And die StatusBar ist optisch abgesetzt
```

### Szenario 6: Integration mit KeyHints

```gherkin
Scenario: StatusBar nutzt TuiKeyHints Komponente
  Given ich habe KeyHints ["Enter Build", "Esc Cancel"]
  When die StatusBar mit diesen Hints gerendert wird
  Then sehe ich die formatierten Hints rechts
  And das Format entspricht TuiKeyHints
```

### Edge Case: Leere StatusBar

```gherkin
Scenario: StatusBar ohne Content
  Given ich habe keine Status-Informationen
  When die StatusBar gerendert wird
  Then wird eine leere Bar mit Hintergrund angezeigt
  And die Höhe bleibt konstant
```

### Edge Case: Lange Status-Texte

```gherkin
Scenario: Lange Texte werden gekürzt
  Given ich habe einen langen Status-Text
  When der Platz begrenzt ist
  Then wird der Text mit "..." gekürzt
  And die KeyHints bleiben vollständig
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/statusbar.go
- [ ] FILE_EXISTS: internal/ui/components/statusbar_test.go

### Inhalt-Prüfungen

- [ ] CONTAINS: statusbar.go enthält "TuiStatusBar"
- [ ] CONTAINS: statusbar.go enthält "TuiKeyHints" oder importiert keyhints

### Funktions-Prüfungen

- [ ] BUILD_PASS: go build ./internal/ui/components/...
- [ ] TEST_PASS: go test ./internal/ui/components/... -run TestStatusBar -v
- [ ] LINT_PASS: golangci-lint run ./internal/ui/components/statusbar.go

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

#### Integration DoD
- [ ] **Integration hergestellt: TuiStatusBar -> TuiKeyHints**
  - [ ] Import/Aufruf existiert in Code
  - [ ] Verbindung ist funktional (nicht nur Stub)
  - [ ] Validierung: KeyHints werden im Right-Section gerendert

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

**Betroffene Komponenten:**

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Frontend | `internal/ui/components/statusbar.go` | New file: TuiStatusBar |
| Frontend | `internal/ui/components/statusbar_test.go` | New file: Unit tests with golden files |

**Integration:** TuiStatusBar enthält TuiKeyHints (Composition)

**Kritische Integration Points:**
- TuiStatusBar (Source) -> TuiKeyHints (Target): Right section rendering

---

### Technical Details

**WAS:**
- `TuiStatusBar`: Renders a full-width bottom status bar with three sections
- Sections: Left (status), Center (info), Right (TuiKeyHints)
- Full-width background with `ColorCard`
- Integration with TuiKeyHints for right section

**WIE (Architektur-Guidance ONLY):**
- Pattern: Pure Lipgloss (stateless render function with width parameter)
- Three-column layout using `lipgloss.JoinHorizontal` with `lipgloss.Left`, `lipgloss.Center`, `lipgloss.Right` alignment
- Use `StyleFooter` from `styles.go` as base
- Left section: Status text with color based on state (green=success, yellow=running, red=error)
- Center section: Screen name or info with `ColorTextSecondary`
- Right section: Call `TuiKeyHints()` function from `keyhints.go`
- Width parameter required to calculate spacing between sections
- Use `Truncate()` helper for long status text

**WO:**
- `internal/ui/components/statusbar.go` - Main component implementation
- `internal/ui/components/statusbar_test.go` - Unit tests with golden files
- `internal/ui/components/testdata/` - Golden files for visual regression

**WER:** dev-team__frontend-developer

**Abhängigkeiten:** LAYOUT-003 (TuiKeyHints must be implemented first)

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
go test ./internal/ui/components/... -run TestStatusBar -v

# Lint check
golangci-lint run ./internal/ui/components/statusbar.go

# Verify TuiStatusBar function exists
grep -q "TuiStatusBar\|NewTuiStatusBar" internal/ui/components/statusbar.go

# Verify KeyHints integration
grep -q "TuiKeyHints\|keyhints" internal/ui/components/statusbar.go
```
