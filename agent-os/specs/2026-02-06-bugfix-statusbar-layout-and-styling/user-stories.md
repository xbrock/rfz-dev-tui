# User Stories - StatusBar Layout and Styling

## Overview
Bug fix implementation stories for the TuiStatusBar component that needs to match the approved design mockups with badge-based layout, colored mode pills, context badges, and styled key hints.

---

## Story 1: Redesign TuiStatusBar with Badge-Based Layout

### Beschreibung
Als Entwickler muss ich die TuiStatusBar-Komponente umbauen, damit sie farbige Mode-Badges, Context-Badges und gestylte Key-Hints anzeigt, wie in den Design-Mockups definiert.

### Fachliche Beschreibung
Der Bug verursacht eine plain-text 3-Spalten-Darstellung anstatt der badge-basierten Darstellung mit farbigen Pills (Mode-Badge wie "LOGS", "SELECT"), Context-Badges (aktive Komponente) und individuell gestylten Key-Hints ohne Punkt-Separatoren.

### Technische Verfeinerung

#### WAS (Anforderungen)
- `FooterItemActive()` Funktion implementieren (definiert in design-system.md, fehlt im Code)
- `TuiStatusBarConfig` erweitern um Mode-Badge und Context-Badge Felder
- Badge-Rendering mit farbigen Hintergründen (Pill-Styling) implementieren
- `TuiKeyHints` Separator-Logik ändern: Spacing statt Punkt-Separatoren
- Layout umbauen: `[Mode Badge] [Context Badge] ... [Key Hints] ... [q Quit]`
- Sicherstellen, dass keine Side-Effects in anderen Komponenten entstehen

#### WIE (Implementierung)
- `styles.go`: Add `FooterItemActive()`, `FooterItem()` functions, add `StyleFooterBadge` styles per mode
- `statusbar.go`: Restructure `TuiStatusBarConfig` to include `ModeBadge`, `ModeBadgeColor`, `ContextBadge` fields; rewrite layout from 3-column to badge-based
- `keyhints.go`: Change separator from `" · "` to `"  "` (double space); separate quit hint to far right
- `layout_gallery.go`: Update `renderStatusBar()` to use new badge-based config

#### WO (Betroffene Dateien)
- `internal/ui/components/styles.go` — Add FooterItemActive, FooterItem, badge styles
- `internal/ui/components/statusbar.go` — Restructure config and rendering logic
- `internal/ui/components/keyhints.go` — Change separator style
- `internal/ui/components/demo/layout_gallery.go` — Update status bar usage

#### WER (Zuständigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

### Definition of Ready (DoR)
- [x] Bug-Beschreibung verstanden
- [x] Reproduction Steps verifiziert
- [x] Betroffene Komponenten identifiziert (statusbar.go, keyhints.go, styles.go)
- [x] Technical-Spec vom Architekten vorhanden

### Definition of Done (DoD)
- [ ] Root-Cause dokumentiert
- [ ] Fix implementiert (badge layout, FooterItemActive, styled hints)
- [ ] StatusBar zeigt Mode-Badge mit farbigem Hintergrund
- [ ] StatusBar zeigt Context-Badge mit Selektionsinfo
- [ ] KeyHints ohne Punkt-Separatoren, mit Spacing
- [ ] "q Quit" visuell getrennt am rechten Rand
- [ ] Visuelle Ausgabe stimmt mit should.png und should-2.png überein
- [ ] Keine Regression in bestehenden Tests
- [ ] Architect Review bestanden

### Story Points
5

### Dependencies
None

---

## Story 2: Update Regression Tests for StatusBar and KeyHints

### Beschreibung
Als QA-Spezialist muss ich die bestehenden Golden-Tests aktualisieren und neue Regression-Tests erstellen, damit die neue Badge-basierte StatusBar automatisch getestet wird.

### Fachliche Beschreibung
Die bestehenden Golden-Tests für TuiStatusBar (8 Tests) und TuiKeyHints (6 Tests) müssen an das neue Badge-Layout angepasst werden. Zusätzlich werden neue Tests für Mode-Badges und Context-Badges benötigt.

### Technische Verfeinerung

#### WAS (Anforderungen)
- Bestehende Golden-Test-Dateien aktualisieren (`.golden` files)
- Neue Tests für Mode-Badge Rendering (verschiedene Farben)
- Neue Tests für Context-Badge Rendering
- Tests für KeyHints ohne Punkt-Separatoren
- Tests für "q Quit" Separation
- Tests für FooterItemActive() Funktion

#### WIE (Implementierung)
- Update all `.golden` files to match new badge-based rendering output
- Add test cases: `TestStatusBar_ModeBadge`, `TestStatusBar_ContextBadge`, `TestStatusBar_BadgeColors`
- Update `TestKeyHints_Multiple` to verify no dot separators
- Add `TestFooterItemActive` for the new function
- Run with `-update` flag to regenerate golden files after visual verification

#### WO (Betroffene Dateien)
- `internal/ui/components/statusbar_test.go` — Update existing + add new test cases
- `internal/ui/components/keyhints_test.go` — Update separator tests
- `internal/ui/components/testdata/` — Updated golden files

#### WER (Zuständigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

### Definition of Ready (DoR)
- [ ] Story 1 abgeschlossen
- [ ] Neues Badge-Layout implementiert und verifiziert
- [ ] Test-Strategie festgelegt (golden testing mit charmbracelet/x/exp/golden)

### Definition of Done (DoD)
- [ ] Alle bestehenden Tests aktualisiert
- [ ] Neue Badge-Tests implementiert
- [ ] Alle Tests bestehen mit Fix
- [ ] Golden files regeneriert und verifiziert
- [ ] `go test ./internal/ui/components/...` erfolgreich

### Story Points
3

### Dependencies
- Story 1 (Bug Fix)
