# Visual Regression Tests

> Story ID: WELCOME-006
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Test
**Estimated Effort**: S
**Dependencies**: WELCOME-002, WELCOME-003, WELCOME-004, WELCOME-005

---

## Feature

```gherkin
Feature: Visuelle Regressionstests
  Als Entwickler des RFZ-CLI
  moechte ich Golden-File-Tests fuer alle UI-Zustaende haben,
  damit UI-Aenderungen automatisch erkannt und geprueft werden.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Welcome Screen Default-Zustand

```gherkin
Scenario: Golden File fuer Welcome Screen
  Given die App wird mit Terminal-Groesse 120x40 initialisiert
  When der Welcome Screen gerendert wird
  Then stimmt die Ausgabe mit dem Golden File "app-welcome-default" ueberein
```

### Szenario 2: Navigations-Fokus-Zustaende

```gherkin
Scenario Outline: Golden Files fuer Navigation
  Given die App wird mit Terminal-Groesse 120x40 initialisiert
  And der Cursor steht auf Menuepunkt <index>
  When der Bildschirm gerendert wird
  Then stimmt die Ausgabe mit dem Golden File "<golden_file>" ueberein

  Examples:
    | index | golden_file |
    | 0 | app-nav-build-focused |
    | 1 | app-nav-logs-focused |
    | 2 | app-nav-discover-focused |
    | 3 | app-nav-config-focused |
    | 4 | app-nav-exit-focused |
```

### Szenario 3: Platzhalter-Bildschirme

```gherkin
Scenario Outline: Golden Files fuer Platzhalter-Screens
  Given die App wird mit Terminal-Groesse 120x40 initialisiert
  When ich zum "<screen>" Platzhalter wechsle
  Then stimmt die Ausgabe mit dem Golden File "<golden_file>" ueberein

  Examples:
    | screen | golden_file |
    | Build Components | app-placeholder-build |
    | View Logs | app-placeholder-logs |
    | Discover | app-placeholder-discover |
    | Configuration | app-placeholder-config |
```

### Szenario 4: Exit Modal

```gherkin
Scenario: Golden File fuer Bestaetigungsdialog
  Given die App wird mit Terminal-Groesse 120x40 initialisiert
  When der Exit-Bestaetigungsdialog sichtbar ist
  Then stimmt die Ausgabe mit dem Golden File "app-exit-modal" ueberein
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Golden File fuer zu kleines Terminal
  Given die App wird mit Terminal-Groesse 60x15 initialisiert
  When der Bildschirm gerendert wird
  Then stimmt die Ausgabe mit dem Golden File "app-terminal-too-small" ueberein
  And die Meldung "Terminal too small" ist sichtbar
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/app/app_test.go

### Funktions-Pruefungen

- [ ] TEST_PASS: `go test ./internal/app/... -count=1`
- [ ] LINT_PASS: `golangci-lint run ./internal/app/...`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

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
- [x] Handover-Dokumente definiert (bei Multi-Layer)

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Golden File Tests fuer alle UI-Zustaende erstellt
- [ ] Tests laufen bei 120x40 Terminal-Groesse

#### Qualitaetssicherung
- [ ] Alle Tests bestehen beim ersten Lauf
- [ ] Golden Files eingecheckt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Test | internal/app/app_test.go | Golden File Tests fuer App Shell |
| Test | internal/app/testdata/golden/ | Golden Files fuer alle UI-Zustaende |

---

### Technical Details

**WAS:** Golden File Tests fuer alle App Shell UI-Zustaende erstellen

**WIE:**
- Bestehendes Test-Pattern befolgen: siehe internal/ui/components/statusbar_test.go als Referenz fuer golden.RequireEqual Muster
- Import "github.com/charmbracelet/x/exp/golden" fuer Golden-File-Vergleich
- Jeder Test-Case: app.New() erstellen, dann m.Update(tea.WindowSizeMsg{Width: 120, Height: 40}) senden
- View() aufrufen und Ergebnis als []byte an golden.RequireEqual(t, output) uebergeben
- Fuer Screen-Wechsel: Nach WindowSizeMsg weitere tea.KeyMsg senden (z.B. tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}} fuer Build)
- Fuer Modal: tea.KeyMsg mit "q" senden um Exit-Modal zu oeffnen
- Fuer zu kleines Terminal: tea.WindowSizeMsg{Width: 60, Height: 15} senden
- Test-Zustaende (12 Golden Files): welcome-default, nav-build-focused, nav-logs-focused, nav-discover-focused, nav-config-focused, nav-exit-focused, placeholder-build, placeholder-logs, placeholder-discover, placeholder-config, exit-modal, terminal-too-small
- Golden Files werden automatisch in internal/app/testdata/ generiert bei erstem Lauf mit -update Flag
- Canonical Terminal Size: 120x40 (ausser terminal-too-small Test)

**WO:**
- internal/app/app_test.go (neu - ~200 LOC, 12 Test-Funktionen)
- internal/app/testdata/golden/*.golden (auto-generiert bei -update)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** WELCOME-002, WELCOME-003, WELCOME-004, WELCOME-005 (alle UI-Features muessen implementiert sein, da Tests die vollstaendige View rendern)

**Geschaetzte Komplexitaet:** S (1 neue Test-Datei + generierte Golden Files, ~200 LOC)

---

### Completion Check

```bash
go test ./internal/app/... -count=1
golangci-lint run ./internal/app/...
```

**Story ist DONE wenn:**
1. Alle TEST_PASS commands exit 0
2. Golden Files existieren fuer alle getesteten Zustaende
3. Tests sind reproduzierbar
