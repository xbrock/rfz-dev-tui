# User Stories: Component Gallery Rendering Issues

> Spec ID: BUGFIX-GALLERY-RENDER
> Created: 2026-02-03
> Total Stories: 3
> Estimated Effort: S (5 SP total)
> Technical Analysis: COMPLETE (2026-02-03)

---

## Story Index

| ID | Title | Type | Priority | Effort | Dependencies | DoR |
|----|-------|------|----------|--------|--------------|-----|
| FIX-001 | Fix TuiBox Border Rendering | Bugfix | Critical | S (2 SP) | None | READY |
| FIX-002 | Fix TuiButton Layout Alignment | Bugfix | Critical | S (2 SP) | None | READY |
| FIX-003 | Add Regression Tests | Test | High | XS (1 SP) | FIX-001, FIX-002 | READY |

---

# FIX-001: Fix TuiBox Border Rendering

> Story ID: FIX-001
> Spec: BUGFIX-GALLERY-RENDER
> Created: 2026-02-03
> Last Updated: 2026-02-03
> Technical Analysis: VERIFIED

**Priority**: Critical
**Type**: Bugfix
**Estimated Effort**: S (2 SP) - VERIFIED
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiBox Border-Rendering Korrektur
  Als TUI-Entwickler
  moechte ich dass TuiBox-Komponenten korrekt nebeneinander gerendert werden,
  damit die Borders vollstaendig verbunden sind und keine floating Characters erscheinen.
```

---

## Fachliche Beschreibung

Die TuiBox-Komponente rendert mehrzeilige Ausgaben (Top-Border, Content, Bottom-Border). Wenn mehrere Boxen horizontal nebeneinander platziert werden sollen, fuehrt einfache String-Konkatenation zu gebrochenen Borders. Die rechte Seite jeder Box erscheint als losgeloeste Zeichen anstatt als verbundene Box-Ecken.

**Problem:** String-Konkatenation von mehrzeiligen Komponenten funktioniert nicht
**Loesung:** Verwendung von `lipgloss.JoinHorizontal()` fuer korrektes horizontales Layout

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Einzelne Box Border Integritaet

```gherkin
Scenario: TuiBox zeigt vollstaendige Borders
  Given ich habe einen TuiBox mit BoxDouble Style
  When die Box gerendert wird
  Then sind alle vier Ecken korrekt verbunden
  And die rechte Border-Linie ist durchgehend
  And keine losgeloesten Zeichen sind sichtbar
```

### Szenario 2: Multiple Boxen Horizontal

```gherkin
Scenario: Mehrere TuiBoxen nebeneinander ausrichten
  Given ich habe drei TuiBoxen mit verschiedenen Styles
  When diese horizontal nebeneinander gerendert werden
  Then erscheinen alle Boxen auf der gleichen vertikalen Baseline
  And jede Box hat vollstaendige, verbundene Borders
  And der Abstand zwischen den Boxen ist gleichmaessig
```

### Szenario 3: Focus State Boxen

```gherkin
Scenario: Focused Boxen mit korrekten Borders
  Given ich habe eine normale und eine fokussierte TuiBox
  When diese nebeneinander gerendert werden
  Then zeigt die fokussierte Box cyan Borders
  And beide Boxen haben vollstaendige Border-Verbindungen
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/components/demo/gallery.go

### Inhalt-Pruefungen

- [ ] CONTAINS: gallery.go enthaelt "lipgloss.JoinHorizontal"
- [ ] NOT_CONTAINS: gallery.go enthaelt NICHT mehrfache `WriteString(TuiBox` ohne JoinHorizontal

### Funktions-Pruefungen

- [ ] BUILD_PASS: `go build ./cmd/rfz/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/demo/gallery.go`
- [ ] VISUAL_CHECK: Manueller Check dass Borders verbunden sind

---

## Technisches Refinement (vom Architect)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und pruefbar
- [x] Business Value verstanden (Component Gallery muss korrekt funktionieren)

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhaengigkeiten identifiziert (keine)
- [x] Betroffene Komponenten bekannt
- [x] Story ist angemessen geschaetzt (1 Datei, ~20 LOC Aenderungen)
- [x] Root cause verified by code analysis
- [x] lipgloss.JoinHorizontal already imported (line 10)

**DoR Status: READY**

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Code implementiert und folgt charm.land First Principle
- [ ] lipgloss.JoinHorizontal() verwendet statt String-Konkatenation
- [ ] Alle TuiBox-Varianten korrekt dargestellt

#### Qualitaetssicherung
- [ ] Alle Akzeptanzkriterien erfuellt
- [ ] Manueller Visual Check durchgefuehrt
- [ ] Code Review durchgefuehrt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Backend-only (Go TUI)

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Backend | internal/ui/components/demo/gallery.go | MODIFY - renderBoxSection() |

**Kritische Integration Points:**
- gallery.go -> lipgloss: Verwendung von JoinHorizontal() (already imported, line 10)
- gallery.go -> box.go: TuiBox bleibt unveraendert

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- `renderBoxSection()` Funktion korrigieren
- String-Konkatenation durch lipgloss.JoinHorizontal() ersetzen

**WIE (Verified Implementation Steps):**

1. Open `/Users/lix/xapps/rfz-tui/internal/ui/components/demo/gallery.go`
2. Locate `renderBoxSection()` function (line 98)
3. Replace lines 106-114 (Border Variants section):

```go
// VORHER (kaputt) - Lines 106-114:
sb.WriteString("Border Variants:\n")
sb.WriteString(components.TuiBox("Single Border", components.BoxSingle, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Double Border", components.BoxDouble, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Rounded Border", components.BoxRounded, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Heavy Border", components.BoxHeavy, false))
sb.WriteString("\n\n")

// NACHHER (korrigiert):
sb.WriteString("Border Variants:\n")
boxRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiBox("Single Border", components.BoxSingle, false),
    "  ",
    components.TuiBox("Double Border", components.BoxDouble, false),
    "  ",
    components.TuiBox("Rounded Border", components.BoxRounded, false),
    "  ",
    components.TuiBox("Heavy Border", components.BoxHeavy, false),
)
sb.WriteString(boxRow)
sb.WriteString("\n\n")
```

4. Replace lines 116-119 (Focus State section):

```go
// VORHER (kaputt) - Lines 116-119:
sb.WriteString("Focus State:\n")
sb.WriteString(components.TuiBox("Normal", components.BoxSingle, false))
sb.WriteString("  ")
sb.WriteString(components.TuiBox("Focused", components.BoxSingle, true))

// NACHHER (korrigiert):
sb.WriteString("Focus State:\n")
focusRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiBox("Normal", components.BoxSingle, false),
    "  ",
    components.TuiBox("Focused", components.BoxSingle, true),
)
sb.WriteString(focusRow)
```

5. Run `go build ./cmd/rfz/...` to verify compilation
6. Run `./rfz` to visually verify box borders are connected
7. Run `golangci-lint run ./internal/ui/components/demo/gallery.go`

**WO (Affected Files with Line Numbers):**

| File | Lines | Change Type | Description |
|------|-------|-------------|-------------|
| `internal/ui/components/demo/gallery.go` | 106-114 | MODIFY | Border Variants - replace 8 WriteString calls with JoinHorizontal |
| `internal/ui/components/demo/gallery.go` | 116-119 | MODIFY | Focus State - replace 3 WriteString calls with JoinHorizontal |

**Geschaetzte Komplexitaet:** S (2 SP) - VERIFIED

---

### Completion Check

```bash
# Verify build
go build ./cmd/rfz/...

# Verify lint
golangci-lint run ./internal/ui/components/demo/gallery.go

# Visual verification (manual)
./rfz
# Check TuiBox section - all borders should be complete
```

**Story ist DONE wenn:**
1. Alle TuiBox-Borders vollstaendig verbunden sind
2. Build und Lint erfolgreich
3. Manueller Visual Check bestanden

---

---

# FIX-002: Fix TuiButton Layout Alignment

> Story ID: FIX-002
> Spec: BUGFIX-GALLERY-RENDER
> Created: 2026-02-03
> Last Updated: 2026-02-03
> Technical Analysis: VERIFIED

**Priority**: Critical
**Type**: Bugfix
**Estimated Effort**: S (2 SP) - VERIFIED
**Dependencies**: None

---

## Feature

```gherkin
Feature: TuiButton Layout-Alignment Korrektur
  Als TUI-Entwickler
  moechte ich dass TuiButton-Komponenten horizontal korrekt ausgerichtet werden,
  damit alle Buttons in einer Reihe auf der gleichen Baseline erscheinen.
```

---

## Fachliche Beschreibung

Die TuiButton-Komponente hat unterschiedliche Hoehen je nach Variante: Secondary Buttons haben einen Border (3 Zeilen), waehrend Primary und Destructive Buttons keinen Border haben (1 Zeile). Einfache String-Konkatenation fuehrt zu einer Treppen-Stufung der Buttons anstatt horizontaler Ausrichtung.

**Problem:** Unterschiedliche Komponentenhoehen fuehren zu vertikaler Fehlausrichtung
**Loesung:** Verwendung von `lipgloss.JoinHorizontal()` mit Alignment-Parameter

**Root Cause Verified:**
- `TuiButton` with `ButtonSecondary` variant has `Border(BorderSingle)` (button.go line 56-60), producing 3 lines
- `TuiButton` with `ButtonPrimary`/`ButtonDestructive` variants have no border, producing 1 line
- Golden files confirm: `TestTuiButton_Secondary.golden` = 3 lines, `TestTuiButton_Primary.golden` = 1 line

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Button Varianten Ausrichtung

```gherkin
Scenario: Alle Button-Varianten horizontal ausgerichtet
  Given ich habe Primary, Secondary und Destructive Buttons
  When diese nebeneinander gerendert werden
  Then erscheinen alle drei Buttons auf der gleichen horizontalen Linie
  And der Secondary Button (mit Border) ist vertikal zentriert oder top-aligned
  And keine Treppen-Stufung ist sichtbar
```

### Szenario 2: Buttons mit Shortcuts

```gherkin
Scenario: Shortcut-Buttons horizontal ausgerichtet
  Given ich habe Buttons mit Shortcuts ([Enter] Build, [Esc] Cancel, [D] Delete)
  When diese nebeneinander gerendert werden
  Then erscheinen alle Buttons auf der gleichen Baseline
  And die Shortcuts sind korrekt formatiert
```

### Szenario 3: Focus State Buttons

```gherkin
Scenario: Focus State Buttons ausgerichtet
  Given ich habe einen normalen und einen fokussierten Button
  When diese nebeneinander gerendert werden
  Then sind beide Buttons horizontal ausgerichtet
  And der fokussierte Button zeigt Bold+Underline Styling
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/components/demo/gallery.go

### Inhalt-Pruefungen

- [ ] CONTAINS: gallery.go enthaelt "lipgloss.JoinHorizontal" in renderButtonSection
- [ ] NOT_CONTAINS: gallery.go enthaelt NICHT mehrfache `WriteString(TuiButton` ohne JoinHorizontal

### Funktions-Pruefungen

- [ ] BUILD_PASS: `go build ./cmd/rfz/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/demo/gallery.go`
- [ ] VISUAL_CHECK: Manueller Check dass Buttons ausgerichtet sind

---

## Technisches Refinement (vom Architect)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und pruefbar
- [x] Business Value verstanden

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhaengigkeiten identifiziert (keine)
- [x] Betroffene Komponenten bekannt
- [x] Story ist angemessen geschaetzt (1 Datei, ~30 LOC Aenderungen)
- [x] Root cause verified by code analysis (Secondary has border, Primary/Destructive do not)
- [x] lipgloss.JoinHorizontal already imported (line 10)

**DoR Status: READY**

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Code implementiert und folgt charm.land First Principle
- [ ] lipgloss.JoinHorizontal() verwendet statt String-Konkatenation
- [ ] Alle Button-Reihen korrekt ausgerichtet

#### Qualitaetssicherung
- [ ] Alle Akzeptanzkriterien erfuellt
- [ ] Manueller Visual Check durchgefuehrt
- [ ] Code Review durchgefuehrt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Backend-only (Go TUI)

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Backend | internal/ui/components/demo/gallery.go | MODIFY - renderButtonSection() |

**Kritische Integration Points:**
- gallery.go -> lipgloss: Verwendung von JoinHorizontal() (already imported, line 10)
- gallery.go -> button.go: TuiButton bleibt unveraendert

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- `renderButtonSection()` Funktion korrigieren
- String-Konkatenation durch lipgloss.JoinHorizontal() ersetzen

**WIE (Verified Implementation Steps):**

1. Open `/Users/lix/xapps/rfz-tui/internal/ui/components/demo/gallery.go`
2. Locate `renderButtonSection()` function (line 141)
3. Replace lines 148-154 (Button Variants section):

```go
// VORHER (kaputt) - Lines 148-154:
sb.WriteString("Button Variants:\n")
sb.WriteString(components.TuiButton("Primary", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Secondary", components.ButtonSecondary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Destructive", components.ButtonDestructive, "", false))
sb.WriteString("\n\n")

// NACHHER (korrigiert):
sb.WriteString("Button Variants:\n")
variantRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Primary", components.ButtonPrimary, "", false),
    "  ",
    components.TuiButton("Secondary", components.ButtonSecondary, "", false),
    "  ",
    components.TuiButton("Destructive", components.ButtonDestructive, "", false),
)
sb.WriteString(variantRow)
sb.WriteString("\n\n")
```

4. Replace lines 156-162 (With Shortcuts section):

```go
// VORHER (kaputt) - Lines 156-162:
sb.WriteString("With Shortcuts:\n")
sb.WriteString(components.TuiButton("Build", components.ButtonPrimary, "Enter", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Delete", components.ButtonDestructive, "D", false))
sb.WriteString("\n\n")

// NACHHER (korrigiert):
sb.WriteString("With Shortcuts:\n")
shortcutRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Build", components.ButtonPrimary, "Enter", false),
    "  ",
    components.TuiButton("Cancel", components.ButtonSecondary, "Esc", false),
    "  ",
    components.TuiButton("Delete", components.ButtonDestructive, "D", false),
)
sb.WriteString(shortcutRow)
sb.WriteString("\n\n")
```

5. Replace lines 164-167 (Focus State section):

```go
// VORHER (kaputt) - Lines 164-167:
sb.WriteString("Focus State:\n")
sb.WriteString(components.TuiButton("Normal", components.ButtonPrimary, "", false))
sb.WriteString("  ")
sb.WriteString(components.TuiButton("Focused", components.ButtonPrimary, "", true))

// NACHHER (korrigiert):
sb.WriteString("Focus State:\n")
focusBtnRow := lipgloss.JoinHorizontal(
    lipgloss.Top,
    components.TuiButton("Normal", components.ButtonPrimary, "", false),
    "  ",
    components.TuiButton("Focused", components.ButtonPrimary, "", true),
)
sb.WriteString(focusBtnRow)
```

6. Run `go build ./cmd/rfz/...` to verify compilation
7. Run `./rfz` to visually verify buttons are horizontally aligned
8. Run `golangci-lint run ./internal/ui/components/demo/gallery.go`

**WO (Affected Files with Line Numbers):**

| File | Lines | Change Type | Description |
|------|-------|-------------|-------------|
| `internal/ui/components/demo/gallery.go` | 148-154 | MODIFY | Button Variants - replace 6 WriteString calls with JoinHorizontal |
| `internal/ui/components/demo/gallery.go` | 156-162 | MODIFY | With Shortcuts - replace 6 WriteString calls with JoinHorizontal |
| `internal/ui/components/demo/gallery.go` | 164-167 | MODIFY | Focus State - replace 3 WriteString calls with JoinHorizontal |

**Geschaetzte Komplexitaet:** S (2 SP) - VERIFIED

---

### Completion Check

```bash
# Verify build
go build ./cmd/rfz/...

# Verify lint
golangci-lint run ./internal/ui/components/demo/gallery.go

# Visual verification (manual)
./rfz
# Check TuiButton section - all buttons should align horizontally
```

**Story ist DONE wenn:**
1. Alle Button-Reihen horizontal ausgerichtet sind
2. Build und Lint erfolgreich
3. Manueller Visual Check bestanden

---

---

# FIX-003: Add Regression Tests

> Story ID: FIX-003
> Spec: BUGFIX-GALLERY-RENDER
> Created: 2026-02-03
> Last Updated: 2026-02-03
> Technical Analysis: VERIFIED

**Priority**: High
**Type**: Test
**Estimated Effort**: XS (1 SP) - VERIFIED
**Dependencies**: FIX-001, FIX-002

---

## Feature

```gherkin
Feature: Regression Tests fuer Rendering-Fixes
  Als QA-Engineer
  moechte ich dass die Rendering-Fixes durch Golden-File Tests abgesichert sind,
  damit zukuenftige Aenderungen keine Regressionen einfuehren.
```

---

## Fachliche Beschreibung

Nach der Korrektur der TuiBox und TuiButton Rendering-Issues muessen die Golden-File Tests aktualisiert werden, um den korrekten Output zu reflektieren. Dies stellt sicher, dass zukuenftige Code-Aenderungen keine visuellen Regressionen einfuehren.

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Golden Files Aktualisiert

```gherkin
Scenario: Golden Files reflektieren korrektes Rendering
  Given die Rendering-Fixes FIX-001 und FIX-002 sind implementiert
  When ich die Golden Files regeneriere
  Then enthalten die neuen Golden Files korrekte Border-Darstellung
  And enthalten die neuen Golden Files ausgerichtete Button-Layouts
```

### Szenario 2: Tests Bestehen

```gherkin
Scenario: Alle Component Tests bestehen
  Given die Golden Files sind aktualisiert
  When ich `go test ./internal/ui/components/...` ausfuehre
  Then bestehen alle Tests ohne Fehler
  And keine Golden File Mismatches werden gemeldet
```

### Szenario 3: Regression Detection

```gherkin
Scenario: Zukuenftige Regressionen werden erkannt
  Given die Tests bestehen mit korrektem Output
  When jemand versehentlich String-Konkatenation wiedereinfuehrt
  Then schlagen die Golden File Tests fehl
  And der Fehler zeigt den Unterschied im gerenderten Output
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/components/demo/gallery_test.go

### Funktions-Pruefungen

- [ ] TEST_PASS: `go test ./internal/ui/components/... -v`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/...`

---

## Technisches Refinement (vom Architect)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und pruefbar
- [x] Business Value verstanden (Regression Prevention)

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert
- [x] Abhaengigkeiten identifiziert (FIX-001, FIX-002)
- [x] Betroffene Komponenten bekannt
- [x] Golden file location verified: `internal/ui/components/demo/testdata/TestGallery_View_AfterResize.golden`
- [x] Test file verified: `internal/ui/components/demo/gallery_test.go` uses `golden.RequireEqual`

**DoR Status: READY**

---

### DoD (Definition of Done) - Vom Architect

#### Implementierung
- [ ] Golden Files regeneriert mit `-update` Flag
- [ ] Alle Tests bestehen

#### Qualitaetssicherung
- [ ] Manueller Visual Check bestaetigt korrekten Output
- [ ] Test Coverage unveraendert oder verbessert

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Test

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Test | internal/ui/components/demo/gallery_test.go | VERIFY - Tests bestehen |
| Test | internal/ui/components/demo/testdata/TestGallery_View_AfterResize.golden | UPDATE - Regenerate |

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- Golden Files regenerieren nach Fixes
- Sicherstellen dass alle Tests bestehen

**WIE (Verified Implementation Steps):**

1. Ensure FIX-001 and FIX-002 are complete
2. Run `./rfz` and visually confirm all rendering is correct:
   - TuiBox borders are fully connected
   - TuiButton rows are horizontally aligned
3. Regenerate golden files:
   ```bash
   go test ./internal/ui/components/demo/... -update
   ```
4. Verify tests pass:
   ```bash
   go test ./internal/ui/components/... -v
   ```
5. Run lint check:
   ```bash
   golangci-lint run ./internal/ui/components/...
   ```

**WO (Affected Files):**

| File | Action | Description |
|------|--------|-------------|
| `internal/ui/components/demo/testdata/TestGallery_View_AfterResize.golden` | UPDATE | Regenerate with `-update` flag after visual verification |
| `internal/ui/components/demo/gallery_test.go` | VERIFY | No changes, just verify tests pass |

**Geschaetzte Komplexitaet:** XS (1 SP) - VERIFIED

---

### Completion Check

```bash
# Run all component tests
go test ./internal/ui/components/... -v

# Verify lint
golangci-lint run ./internal/ui/components/...
```

**Story ist DONE wenn:**
1. Alle Tests bestehen (`go test` exit code 0)
2. Golden Files sind aktualisiert
3. Lint erfolgreich

---

## Implementation Order

```
FIX-001 (TuiBox) ------+
                       +---> FIX-003 (Regression Tests)
FIX-002 (Button) ------+
```

FIX-001 und FIX-002 koennen parallel implementiert werden.
FIX-003 muss warten bis beide Fixes abgeschlossen sind.

---

## Technical Reference

For detailed technical analysis, see: `sub-specs/technical-spec.md`

---

*Parent Spec: spec.md*
