# User Stories - Build Complete Screen Bug Fix

## Overview
Bug fix implementation stories for the build completion screen that has broken Tab navigation, missing action buttons (View Logs, Rebuild Failed, Back), a duplicate "New Build" label, and a design layout that does not match the approved reference prototype screenshots.

**Technical Spec**: `sub-specs/technical-spec.md`
**Total Story Points**: 13

---

## Story 1: Fix Tab Navigation and Key Bindings on Build Complete Screen

### Beschreibung
Als Entwickler muss ich die Tastaturnavigation auf dem Build-Complete-Screen reparieren, damit Tab wieder den Fokus zwischen Navigation und Content-Bereich wechselt, und die fehlenden Action-Buttons (View Logs, Rebuild Failed, Back) mit ihren Keybindings hinzufuegen.

### Fachliche Beschreibung
Nach Abschluss eines Builds ist die Tab-Taste gefangen, weil die App alle Tastatureingaben an den Build-Screen delegiert, der Tab nicht verarbeitet. Ausserdem fehlen die Buttons "View Logs" (l), "Rebuild Failed" (r) und "Back" (Esc) in der Actions-Box. Stattdessen wird nur ein "New Build"-Button angezeigt, dessen Label doppelt erscheint (als Button und als Hint). Die korrekte Actions-Box zeigt links die drei Buttons und rechts nur "Tab Switch Focus" als Hint.

### Technische Verfeinerung

#### WAS (Anforderungen)
- Tab-Taste im `phaseCompleted`-Zustand muss den Fokus zwischen Navigation und Content wechseln (wie im Selection-Zustand)
- `handleCompletedKey` muss Keybindings fuer `l` (View Logs), `r` (Rebuild Failed), `Esc` (Back) verarbeiten
- "View Logs" (`l`): Nur Keybinding registrieren, kein Modal oeffnen (separate Feature)
- "Rebuild Failed" (`r`): Nur fehlgeschlagene Komponenten mit gleicher Config neu bauen; `StartBuildMsg` mit gefilterter Komponentenliste senden
- "Back" (`Esc`): Zurueck zum Selection-Phase, Build-State zuruecksetzen
- `viewExecutionActions` muss drei Buttons links und einen "Tab Switch Focus"-Hint rechts rendern
- Doppeltes "New Build"-Label entfernen (Button und Hint)
- StatusBar-Hints fuer `phaseCompleted` aktualisieren (Tab Switch Focus statt n New Build)

#### WIE (Implementierung)

**Change 1: `internal/app/app.go` line 173 -- Remove `IsCompleted()` from early-return block**

Current code at line 173:
```go
if m.screen == screenBuild && (m.build.IsConfiguring() || m.build.IsExecuting() || m.build.IsCompleted()) {
```
Change to:
```go
if m.screen == screenBuild && (m.build.IsConfiguring() || m.build.IsExecuting()) {
```
This allows Tab (handled at lines 181-189) and other app-level keys to work in the completed state. The completed build screen receives keys via the content delegation block at lines 200-210 instead.

**Change 2: `internal/app/app.go` lines 191-197 -- Add completed-state Esc delegation**

The general Esc handler navigates to Welcome. For `phaseCompleted`, Esc must go back to selection instead. Modify the `esc` case:
```go
case "esc":
    // In build completed state, Esc returns to selection (handled by build screen)
    if m.screen == screenBuild && m.build.IsCompleted() {
        var cmd tea.Cmd
        m.build, cmd = m.build.Update(msg)
        return m, cmd
    }
    if m.screen != screenWelcome {
        m.navigateTo(screenWelcome)
        m.build = m.build.SetFocused(false)
    }
    return m, nil
```

**Change 3: `internal/app/app.go` lines 507-513 -- Update StatusBar hints**

Replace completed-state hints:
```go
hints = []components.KeyHint{
    {Key: "Tab", Label: "Switch Focus"},
    {Key: "\u2191\u2193", Label: "Navigate"},
    {Key: "Esc", Label: "Back"},
}
```

**Change 4: `internal/ui/screens/build/model.go` after line 192 -- Add `failedComponents()` helper**

```go
func (m Model) failedComponents() []string {
    var failed []string
    for _, s := range m.buildStates {
        if s.Phase == domain.PhaseFailed {
            failed = append(failed, s.Name)
        }
    }
    return failed
}
```

**Change 5: `internal/ui/screens/build/update.go` lines 312-334 -- Rewrite `handleCompletedKey`**

Replace entire function body. Remove `n` handler. Add:
- `l`: No-op placeholder for future log modal
- `r`: Call `m.failedComponents()`, if `len(failed) > 0` send `StartBuildMsg{Config: m.config, Selected: failed}`
- `esc`: Reset `m.phase = phaseSelecting`, clear `buildStates`, `simStates`, `buildCursor`, `buildCanceled`, deselect all items

**Change 6: `internal/ui/screens/build/update.go` lines 285-310 -- Add `l` no-op to `handleExecutionKey`**

Add `case "l":` with empty body (no-op) to match the `[View Logs]` button shown during execution.

**Change 7: `internal/ui/screens/build/execution.go` lines 212-257 -- Rewrite `viewExecutionActions`**

Two states:
- Running: `[View Logs] (l)` (secondary) + `[Cancel Build] (Esc)` (destructive). No hints on right.
- Completed/Canceled: `[View Logs] (l)` (secondary) + `[Rebuild Failed] (r)` (primary) + `[Back] (Esc)` (secondary). Right side: single hint `Tab Switch Focus`.
- Border color: `ColorCyan` when `m.focused` is true, `ColorBorder` otherwise.

#### WO (Betroffene Dateien)

| File | Lines | Change |
|------|-------|--------|
| `internal/app/app.go` | 173 | Remove `m.build.IsCompleted()` from early-return condition |
| `internal/app/app.go` | 191-197 | Add completed-state Esc delegation before general Esc handler |
| `internal/app/app.go` | 507-513 | Update StatusBar hints: `Tab Switch Focus`, `Navigate`, `Esc Back` |
| `internal/ui/screens/build/model.go` | After 192 | Add `failedComponents() []string` helper method |
| `internal/ui/screens/build/update.go` | 285-310 | Add `case "l":` no-op to `handleExecutionKey` |
| `internal/ui/screens/build/update.go` | 312-334 | Rewrite `handleCompletedKey`: remove `n`, add `l`, `r`, `esc` handlers |
| `internal/ui/screens/build/execution.go` | 212-257 | Rewrite `viewExecutionActions` with correct buttons per state |

#### WER (Zustaendigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

### Definition of Ready (DoR)
- [x] Bug-Beschreibung verstanden
- [x] Reproduction Steps verifiziert
- [x] Betroffene Komponenten identifiziert (`app.go`, `update.go`, `execution.go`, `model.go`)
- [x] Root-Cause-Analyse dokumentiert (Tab delegation, missing handlers, duplicate label)
- [x] Referenz-Screenshots vorhanden (`45-build-execution-actions-focus.png`, `49-build-execution-complete.png`)
- [x] Technische Analyse mit exakten Zeilennummern abgeschlossen (siehe `sub-specs/technical-spec.md`)
- [x] `failedComponents()` Helper-Methode spezifiziert
- [x] `StartBuildMsg` Wiederverwendung fuer Rebuild-Failed validiert (existierender Flow in `handleStartBuild` line 48)

### Definition of Done (DoD)
- [x] Tab-Taste wechselt Fokus zwischen Navigation und Content im Build-Complete-Zustand
- [x] Actions-Box zeigt `[View Logs] (l)`, `[Rebuild Failed] (r)`, `[Back] (Esc)` links
- [x] Actions-Box zeigt `Tab Switch Focus` als einzigen Hint rechts
- [x] Taste `l` wird als Keybinding akzeptiert (No-op, kein Fehler)
- [x] Taste `r` startet Build nur fuer fehlgeschlagene Komponenten mit gleicher Config
- [x] Taste `r` ist No-op wenn keine Komponenten fehlgeschlagen sind
- [x] Taste `Esc` kehrt zur Komponentenauswahl zurueck
- [x] Kein doppeltes "New Build"-Label sichtbar
- [x] "New Build" Button und Hint komplett entfernt
- [x] StatusBar zeigt `Tab Switch Focus | Navigate | Esc Back` fuer Build-Complete-Zustand
- [x] Running-State Actions zeigt `[View Logs] (l)` + `[Cancel Build] (Esc)`
- [x] Keine Regression in bestehenden Key-Handling-Tests
- [x] `go test ./internal/app/...` besteht (Golden-Files muessen ggf. regeneriert werden)
- [x] Architect Review bestanden

### Status: Done

### Story Points
5

### Dependencies
None

---

## Story 2: Redesign Build Complete Screen to Match Reference Screenshots

### Beschreibung
Als Entwickler muss ich das Layout und Styling des Build-Execution-Screens (sowohl waehrend der Ausfuehrung als auch nach Abschluss) ueberarbeiten, damit es den Referenz-Screenshots aus dem Prototyp entspricht.

### Fachliche Beschreibung
Der aktuelle Build-Execution-Screen weicht visuell vom genehmigten Prototyp ab. Die Referenz-Screenshots zeigen eine spezifische Box-Struktur: "Build Execution" mit Command-Preview, "Components" mit Tabelle (St, Component, Phase, Progress, Time), "Progress" mit Overall-Fortschrittsbalken und Status-Counter-Badges innerhalb derselben Box, und "Actions" mit Button-Aktionen. Die aktuelle Implementierung hat die Status-Counter ausserhalb der Progress-Box, verwendet "Overall Progress" statt "Overall:" als Label, und hat abweichendes Spacing und Styling.

### Technische Verfeinerung

#### WAS (Anforderungen)
- "Build Execution" Box: Command-Preview mit `$`-Prompt und gruener Schrift, umrahmt mit Rounded-Border
- "Components" Box: Tabelle mit Header-Zeile (`St`, `Component`, `Phase`, `Progress`, `Time`), Dashed-Divider, und Zeilen mit Status-Icon, Name, Phase-Label (farbig), Progress-Bar, Elapsed-Time. Cursor-Highlight fuer fokussierte Zeile (cyan Hintergrund). Umrahmt mit Rounded-Border.
- "Progress" Box: `Overall:` Label (nicht "Overall Progress"), Fortschrittsbalken mit Prozentanzeige, darunter Status-Counter-Badges (`Running: N`, `Success: N`, `Failed: N`, `Pending: N`) mit farbigen Badge-Styles. Alles innerhalb einer Rounded-Border-Box.
- "Actions" Box: Buttons und Hints wie in Story 1 definiert. Umrahmt mit Rounded-Border. Border-Farbe `ColorCyan` wenn Content fokussiert.
- Vertikales Layout: Build Execution -> Components -> Progress -> Actions, jeweils als separate Boxen mit konsistentem Spacing
- Referenz: `references/prototype-screenshots/45-build-execution-actions-focus.png` und `49-build-execution-complete.png`

#### WIE (Implementierung)

**Change 1: `internal/ui/screens/build/execution.go` lines 14-60 -- Rewrite `viewExecution()`**

Current structure wraps everything in one "Build Execution" box. New structure renders four separate bordered boxes:

```
StyleH2.Render("Build")  (cyan title, unchanged)

+-- Build Execution (BorderRounded, ColorBorder) --+
| $ mvn clean install -P... -DskipTests            |
+--------------------------------------------------+
                                                    (blank line)
+-- Components (BorderRounded, ColorCyan) ----------+
| St  Component          Phase    Progress    Time  |
| ------------------------------------------------- |
| > * audiocon           Testing  |||...      00:10 |
|   * traktion           Testing  |||...      00:08 |
+---------------------------------------------------+
                                                    (blank line)
+-- Progress (BorderRounded, ColorBorder) ----------+
| Overall:  [==========......] 67%                  |
| * Running: 3  v Success: 0  x Failed: 0  o P: 0  |
+---------------------------------------------------+
                                                    (blank line)
+-- Actions (BorderRounded, ColorCyan/ColorBorder) -+
| [View Logs] l   [Cancel Build] Esc               |
+---------------------------------------------------+
```

Remove the outer "Build Execution" wrapper box (lines 44-52). Instead, each section gets its own `lipgloss.NewStyle().Border(components.BorderRounded)...` box.

**Change 2: `internal/ui/screens/build/execution.go` lines 62-99 -- Rewrite `viewComponentTable()`**

- Change header column "Status" (line 82) to "St" and reduce `colStatus` from 11 to 4
- Wrap the entire table output (header + divider + rows) in a bordered box:
  ```go
  lipgloss.NewStyle().
      Border(components.BorderRounded).
      BorderForeground(components.ColorCyan).
      Padding(0, 1).
      Width(m.width).
      Render(components.StyleH3.Render("Components") + "\n" + tableContent)
  ```
- Use `components.DividerSingle` for the dashed divider (existing, line 87 -- keep this)

**Change 3: `internal/ui/screens/build/execution.go` lines 101-171 -- Rewrite `viewComponentRow()`**

Current: focused row only has bold cyan name (lines 118-119) and `> ` text prefix (line 161).
Reference: focused row has full-row cyan background highlight spanning all columns, with `>` cursor prefix.

Implementation:
- Use tree-branch prefix characters: `|-- ` for non-focused rows, `>   ` for focused row (matching reference screenshots 40, 44)
- For focused row, apply cyan background to the entire assembled row string:
  ```go
  rowStr := lipgloss.JoinHorizontal(lipgloss.Top, statusCell, nameCell, phaseCell, progressCell, timeCell)
  if isFocused {
      rowStr = lipgloss.NewStyle().Background(components.ColorCyan).Foreground(components.ColorBackground).Width(totalRowWidth).Render("> " + rowStr)
  } else {
      rowStr = "|-- " + rowStr
  }
  ```
- Use `TuiStatusCompact` (from `components/status.go` lines 98-129) instead of `TuiStatus` (full badge) for the `St` column to save space: compact renders single characters like `*` (running), `v` (success), `x` (failed), `o` (pending)

**Change 4: `internal/ui/screens/build/execution.go` lines 173-185 -- Rewrite `viewOverallProgress()`**

- Change label from `"Overall Progress"` (line 180) to `"Overall:"`
- Move status counters (currently rendered separately in `viewExecution` line 41) into this function
- Wrap everything in a bordered "Progress" box:
  ```go
  func (m Model) viewProgressBox() string {
      label := lipgloss.NewStyle().Foreground(components.ColorTextSecondary).Bold(true).Render("Overall:")
      bar := components.TuiProgress(m.overallProgress(), m.width-30, true)
      progressLine := label + "  " + bar
      counters := m.viewStatusCounters()
      content := progressLine + "\n" + counters
      return lipgloss.NewStyle().
          Border(components.BorderRounded).
          BorderForeground(components.ColorBorder).
          Padding(0, 1).
          Width(m.width).
          Render(components.StyleH3.Render("Progress") + "\n" + content)
  }
  ```

**Change 5: `internal/ui/screens/build/execution.go` lines 187-210 -- Rewrite `viewStatusCounters()`**

Current: plain inline text with colored numbers (`0 Running    2 Success    1 Failed    0 Pending`).
Reference: colored pill badges with icon prefix and background color.

Implementation using pill/badge style from reference screenshots:
```go
func (m Model) viewStatusCounters() string {
    running, success, failed, pending := m.statusCounts()

    // Each counter as a colored badge/pill
    runningBadge := lipgloss.NewStyle().
        Background(components.ColorCyan).
        Foreground(components.ColorBackground).
        Bold(true).
        Padding(0, 1).
        Render(fmt.Sprintf("* Running: %d", running))

    successBadge := lipgloss.NewStyle().
        Background(components.ColorGreen).
        Foreground(components.ColorBackground).
        Bold(true).
        Padding(0, 1).
        Render(fmt.Sprintf("v Success: %d", success))

    failedBadge := lipgloss.NewStyle().
        Background(components.ColorDestructive).
        Foreground(components.ColorTextPrimary).
        Bold(true).
        Padding(0, 1).
        Render(fmt.Sprintf("x Failed: %d", failed))

    pendingBadge := lipgloss.NewStyle().
        Background(components.ColorSecondary).
        Foreground(components.ColorTextPrimary).
        Padding(0, 1).
        Render(fmt.Sprintf("o Pending: %d", pending))

    return lipgloss.JoinHorizontal(lipgloss.Top,
        runningBadge, "  ", successBadge, "  ", failedBadge, "  ", pendingBadge)
}
```

**Change 6: Actions box border -- integrate with Story 1 viewExecutionActions rewrite**

Use `m.focused` field (from `model.go` line 100) to determine border color:
```go
borderColor := components.ColorBorder
if m.focused {
    borderColor = components.ColorCyan
}
```

#### WO (Betroffene Dateien)

| File | Lines | Change |
|------|-------|--------|
| `internal/ui/screens/build/execution.go` | 14-60 | Rewrite `viewExecution()`: four separate bordered boxes stacked vertically |
| `internal/ui/screens/build/execution.go` | 62-99 | Rewrite `viewComponentTable()`: wrap in "Components" bordered box, change "Status" to "St", reduce column width |
| `internal/ui/screens/build/execution.go` | 101-171 | Rewrite `viewComponentRow()`: full-row cyan background highlight, tree-branch prefix, `TuiStatusCompact` for St column |
| `internal/ui/screens/build/execution.go` | 173-185 | Rewrite as `viewProgressBox()`: label "Overall:", merge status counters inside, wrap in "Progress" bordered box |
| `internal/ui/screens/build/execution.go` | 187-210 | Rewrite `viewStatusCounters()`: colored pill badges with background color and icon prefix |
| `internal/ui/screens/build/execution.go` | 248-250 | Actions box border: use `m.focused` for `ColorCyan`/`ColorBorder` |
| `internal/ui/components/styles.go` | N/A | No changes needed -- existing `ColorCyan`, `ColorGreen`, `ColorDestructive`, `ColorSecondary`, `BorderRounded`, `StyleH3` cover all required styles |

#### WER (Zustaendigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

### Definition of Ready (DoR)
- [x] Referenz-Screenshots analysiert und dokumentiert
- [x] Abweichungen zwischen Ist- und Soll-Zustand identifiziert (see `sub-specs/technical-spec.md` Section 4)
- [x] Story 1 (Tab Navigation und Key Bindings) abgeschlossen oder parallel durchfuehrbar
- [x] Design-System-Styles verfuegbar (`components/styles.go` -- alle benoetigten Styles vorhanden)
- [x] `TuiStatusCompact` vorhanden in `components/status.go` lines 98-129 fuer compact St column
- [x] `TuiProgress` vorhanden in `components/progress.go` lines 98-129 fuer static progress bar
- [x] `TuiDivider` vorhanden in `components/divider.go` fuer dashed divider
- [x] `TuiButton` vorhanden in `components/button.go` fuer action buttons
- [x] Box-Rendering-Muster validiert (see existing `viewExecutionActions` line 248 pattern)

### Definition of Done (DoD)
- [ ] "Build Execution" Box mit Command-Preview in Rounded-Border (eigene Box, nicht Teil eines Wrappers)
- [ ] "Components" Box als eigene Rounded-Border-Box mit Titel "Components"
- [ ] Tabellen-Header: `St`, `Component`, `Phase`, `Progress`, `Time` (nicht "Status")
- [ ] `St` Spalte nutzt `TuiStatusCompact` (einzelne Zeichen: `*`, `v`, `x`, `o`)
- [ ] Fokussierte Zeile in Components-Tabelle mit cyan Hintergrund-Highlight (volle Zeile, nicht nur Name)
- [ ] Tree-Branch-Prefix: `|-- ` fuer normale Zeilen, `>   ` fuer fokussierte Zeile
- [ ] "Progress" Box als eigene Rounded-Border-Box mit Titel "Progress"
- [ ] Progress-Label: "Overall:" (nicht "Overall Progress")
- [ ] Status-Counter-Badges INNERHALB der Progress-Box (nicht darunter)
- [ ] Status-Counter als farbige Badges/Pills mit Hintergrundfarbe gerendert
- [ ] "Actions" Box als eigene Rounded-Border-Box mit Titel "Actions"
- [ ] Actions-Box-Border ist `ColorCyan` (`#0891b2`) wenn Content fokussiert, `ColorBorder` (`#4a4a5e`) sonst
- [ ] Vertikales Layout mit einzelner Leerzeile zwischen Boxen
- [ ] Visueller Vergleich mit `references/prototype-screenshots/45-build-execution-actions-focus.png` bestanden
- [ ] Visueller Vergleich mit `references/prototype-screenshots/49-build-execution-complete.png` bestanden
- [ ] Visueller Vergleich mit `references/prototype-screenshots/40-build-execution-starting.png` bestanden
- [ ] Keine Regression in bestehenden Rendering-Tests (Golden-Files regeneriert)
- [ ] Architect Review bestanden

### Story Points
5

### Dependencies
- Story 1 (for correct Actions box content, can be developed in parallel but must be integrated)

---

## Story 3: Add Regression Tests for Build Complete Screen

### Beschreibung
Als Entwickler muss ich Regressionstests fuer den Build-Complete-Screen erstellen, damit die korrigierten Keybindings, die neuen Action-Buttons und das ueberarbeitete Layout automatisch getestet werden.

### Fachliche Beschreibung
Es existieren derzeit keine Tests fuer den Build-Execution- und Build-Complete-Screen (`internal/ui/screens/build/` hat keine `_test.go`-Dateien). Tests muessen die Tab-Navigation, alle Keybindings (l, r, Esc), die Rebuild-Failed-Logik, und das visuelle Layout des Complete-Screens abdecken.

### Technische Verfeinerung

#### WAS (Anforderungen)
- Unit-Tests fuer `handleCompletedKey`:
  - Tab-Taste wird nicht vom Build-Screen konsumiert (verifizieren, dass kein Handler greift, da Tab vom App-Level verarbeitet wird)
  - `l`-Taste: Keybinding akzeptiert, kein Fehler, kein Zustandswechsel
  - `r`-Taste: Nur fehlgeschlagene Komponenten werden zum Rebuild ausgewaehlt; `StartBuildMsg` mit korrekter gefilterter Liste
  - `r`-Taste bei keinen Failed-Komponenten: No-op, kein Command
  - `Esc`-Taste: Phase wechselt zu `phaseSelecting`, Build-State wird zurueckgesetzt
  - `up/k`, `down/j`: Cursor-Navigation in der Tabelle
- Unit-Tests fuer `failedComponents()` Helper:
  - Keine fehlgeschlagenen: leere/nil Liste
  - Einige fehlgeschlagen: nur fehlgeschlagene Namen
  - Alle fehlgeschlagen: alle Namen
  - Leere buildStates: nil Liste
- Integration-Tests fuer Tab-Fokus-Wechsel im App-Level:
  - Build-Complete-Zustand + Tab = Fokus wechselt zu Navigation
  - Build-Complete-Zustand + Tab + Tab = Fokus zurueck zu Content
  - Build-Complete-Zustand + Esc = zurueck zu Selection (nicht Welcome)
- Golden-Tests fuer `viewExecution()` und `viewExecutionActions()`:
  - Build-Running-Zustand: Actions zeigen `[View Logs]` und `[Cancel Build]`
  - Build-Complete-Zustand: Actions zeigen `[View Logs]`, `[Rebuild Failed]`, `[Back]`
  - Build-Canceled-Zustand: Actions zeigen `[View Logs]`, `[Rebuild Failed]`, `[Back]`
- Canonical terminal size: 120x40

#### WIE (Implementierung)

**New file: `internal/ui/screens/build/update_test.go`**

Key handler unit tests. Use the `teststate.go` fixtures (`TestCompletedState`, `TestExecutingState`) to create models in deterministic states. Send `tea.KeyMsg` and assert:
- Phase transitions (e.g., `esc` -> `phaseSelecting`)
- Returned commands (e.g., `r` -> `StartBuildMsg` with filtered component names)
- Cursor position changes (e.g., `down` increments `buildCursor`)
- No-op behavior (e.g., `l` returns no command, no state change)

Test functions:
```go
func TestHandleCompletedKey_ViewLogs(t *testing.T)
func TestHandleCompletedKey_RebuildFailed(t *testing.T)
func TestHandleCompletedKey_RebuildFailed_NoFailures(t *testing.T)
func TestHandleCompletedKey_Back(t *testing.T)
func TestHandleCompletedKey_CursorUp(t *testing.T)
func TestHandleCompletedKey_CursorDown(t *testing.T)
func TestFailedComponents_Mixed(t *testing.T)
func TestFailedComponents_None(t *testing.T)
func TestFailedComponents_All(t *testing.T)
func TestFailedComponents_Empty(t *testing.T)
```

**New file: `internal/ui/screens/build/execution_test.go`**

Golden tests for view rendering. Use `charmbracelet/x/exp/golden` (already used in `app_test.go`).
```go
func TestViewExecution_Running(t *testing.T)
func TestViewExecution_Completed(t *testing.T)
func TestViewExecution_Canceled(t *testing.T)
func TestViewExecutionActions_Running(t *testing.T)
func TestViewExecutionActions_Completed(t *testing.T)
```

**Extend: `internal/app/app_test.go` after line 165**

Add integration tests using `build.TestCompletedState`:
```go
func TestApp_BuildCompleted_TabSwitchesFocus(t *testing.T)
func TestApp_BuildCompleted_TabBackToContent(t *testing.T)
func TestApp_BuildCompleted_EscReturnsToSelection(t *testing.T)
```

**New fixture: `internal/ui/screens/build/teststate.go` after line 53**

Add `TestCanceledState` for canceled build golden tests.

**Golden files directory: `internal/ui/screens/build/testdata/`**

Generated automatically by running `go test ./internal/ui/screens/build/... -update`.

#### WO (Betroffene Dateien)

| File | Type | Content |
|------|------|---------|
| `internal/ui/screens/build/update_test.go` | New | Key handler unit tests for `handleCompletedKey`, `failedComponents()` |
| `internal/ui/screens/build/execution_test.go` | New | Golden tests for `viewExecution()`, `viewExecutionActions()` |
| `internal/ui/screens/build/testdata/` | New dir | Golden files (auto-generated) |
| `internal/ui/screens/build/teststate.go` | Extend | Add `TestCanceledState` fixture after line 53 |
| `internal/app/app_test.go` | Extend | Add 3 integration tests after line 165: Tab focus, Tab back, Esc behavior |
| `internal/app/testdata/TestApp_BuildCompleted.golden` | Regenerate | Updated golden file for completed screen |
| `internal/app/testdata/TestApp_BuildExecuting.golden` | Regenerate | Updated golden file for executing screen |

#### WER (Zustaendigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

### Definition of Ready (DoR)
- [ ] Story 1 abgeschlossen (Keybindings implementiert)
- [ ] Story 2 abgeschlossen (Layout-Redesign implementiert)
- [x] Test-Strategie festgelegt (Golden-Testing mit `charmbracelet/x/exp/golden`, Unit-Tests, Integration-Tests)
- [x] `teststate.go` Fixtures vorhanden (`TestExecutingState` line 11, `TestCompletedState` line 34)
- [x] Existierende Test-Patterns analysiert (`internal/app/app_test.go` -- `initModel`, `sendKey`, `sendKeyWithCmd` helpers)
- [x] Golden-Testing-Framework verfuegbar (`charmbracelet/x/exp/golden` bereits in `go.mod`)

### Definition of Done (DoD)
- [ ] `handleCompletedKey` Tests fuer alle Keys (l, r, Esc, up/down) -- 6 Tests
- [ ] `failedComponents()` Tests fuer alle Edge-Cases -- 4 Tests
- [ ] Tab-Focus-Integration-Tests im App-Level bestehen -- 3 Tests
- [ ] Golden-Tests fuer Build-Running und Build-Complete Views -- 5 Tests
- [ ] Alle Tests bestehen mit `go test ./internal/ui/screens/build/... ./internal/app/...`
- [ ] Golden-Files generiert und visuell verifiziert bei 120x40
- [ ] Code-Coverage fuer `handleCompletedKey` >= 90%
- [ ] Code-Coverage fuer `failedComponents` = 100%
- [ ] Existierende Golden-Files (`TestApp_BuildExecuting`, `TestApp_BuildCompleted`) regeneriert und verifiziert
- [ ] `golangci-lint run` besteht
- [ ] Architect Review bestanden

### Story Points
3

### Dependencies
- Story 1 (Key Bindings Fix)
- Story 2 (Layout Redesign)
