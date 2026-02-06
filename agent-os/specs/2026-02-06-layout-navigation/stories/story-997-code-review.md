# Code Review by Opus

> Story ID: LAYOUT-997
> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

**Priority**: High
**Type**: Review
**Estimated Effort**: S (2 SP)
**Dependencies**: LAYOUT-001 to LAYOUT-009

---

## Feature

```gherkin
Feature: Code Review by Opus
  Als Tech Lead
  moechte ich einen Code Review durch Claude Opus,
  damit die Code-Qualitaet vor dem Merge sichergestellt ist.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Alle Dateien reviewen

```gherkin
Scenario: Opus reviewt alle neuen/geaenderten Dateien
  Given alle Stories LAYOUT-001 bis LAYOUT-009 sind implementiert
  When Opus den Code Review durchfuehrt
  Then werden alle neuen Dateien geprueft
  And alle geaenderten Dateien werden geprueft
```

### Szenario 2: Style Guide Compliance

```gherkin
Scenario: Code folgt Style Guide
  Given Opus reviewt den Code
  When Style-Violations gefunden werden
  Then werden diese dokumentiert
  And Korrekturvorschlaege gemacht
```

### Szenario 3: Pattern Compliance

```gherkin
Scenario: Architektur-Patterns eingehalten
  Given Opus reviewt den Code
  When Pattern-Violations gefunden werden (z.B. kein Lipgloss verwendet)
  Then werden diese als kritisch markiert
  And muessen vor Merge behoben werden
```

### Szenario 4: Test Coverage

```gherkin
Scenario: Ausreichende Test-Abdeckung
  Given Opus reviewt den Code
  When Test Coverage geprueft wird
  Then haben alle Komponenten Unit Tests
  And alle visuellen States haben Golden Files
```

---

## Technische Verifikation (Automated Checks)

### Review-Scope

- [ ] `internal/ui/components/navigation.go` - TuiNavigation + TuiNavItem
- [ ] `internal/ui/components/modal.go` - TuiModal
- [ ] `internal/ui/components/keyhints.go` - TuiKeyHints
- [ ] `internal/ui/components/table.go` - TuiTable
- [ ] `internal/ui/components/tree.go` - TuiTree
- [ ] `internal/ui/components/tabs.go` - TuiTabs
- [ ] `internal/ui/components/statusbar.go` - TuiStatusBar
- [ ] `internal/ui/components/styles.go` - Style additions
- [ ] `internal/ui/components/demo/layout_gallery.go` - Demo
- [ ] `cmd/layout-demo/main.go` - Demo entry point
- [ ] All `*_test.go` files for new components

### Funktions-Pruefungen

- [ ] BUILD_PASS: go build ./...
- [ ] TEST_PASS: go test ./internal/ui/components/... -v
- [ ] LINT_PASS: golangci-lint run ./...

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
- [ ] Code Review abgeschlossen
- [ ] Alle kritischen Findings behoben
- [ ] Review-Bericht erstellt

#### Qualitaetssicherung
- [ ] Alle Akzeptanzkriterien erfuellt
- [ ] Style Guide Compliance bestaetigt
- [ ] Pattern Compliance bestaetigt

#### Dokumentation
- [ ] Review-Findings dokumentiert
- [ ] Keine offenen kritischen Issues

---

### Betroffene Layer & Komponenten

**Integration Type:** Review-only

**Betroffene Komponenten:**

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Review | All new files from LAYOUT-001 to LAYOUT-009 | Code Review |

---

### Technical Details

**WAS:**
- Vollstaendiger Code Review aller neuen/geaenderten Dateien
- Pruefung auf Style Guide Compliance
- Pruefung auf Architektur-Pattern Compliance
- Pruefung auf ausreichende Test Coverage

**WIE (Architektur-Guidance ONLY):**
- Review durchfuehren mit Fokus auf:
  - Lipgloss/Bubbles First Rule (keine manuellen ANSI codes)
  - Verwendung bestehender Styles aus styles.go
  - Konsistente Benennung und Dokumentation
  - Error Handling
  - Test Coverage fuer alle visuellen States
- Findings in Review-Bericht dokumentieren
- Kritische Issues muessen vor Merge behoben werden

**WO:**
- Review-Bericht: `agent-os/specs/2026-02-06-layout-navigation/implementation-reports/code-review-997.md`

**WER:** Claude Opus (via /execute-tasks or manual review)

**Abhaengigkeiten:** LAYOUT-001 to LAYOUT-009 (alle Stories muessen implementiert sein)

**Geschaetzte Komplexitaet:** S

---

### Relevante Skills

| Skill | Pfad | Grund |
|-------|------|-------|
| N/A | - | No special skills required |

---

### Completion Check

```bash
# All builds pass
go build ./...

# All tests pass
go test ./internal/ui/components/... -v

# Lint passes
golangci-lint run ./...

# Review report exists
test -f agent-os/specs/2026-02-06-layout-navigation/implementation-reports/code-review-997.md
```
