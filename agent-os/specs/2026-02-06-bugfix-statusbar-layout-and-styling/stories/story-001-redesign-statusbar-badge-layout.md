# Story 001: Redesign TuiStatusBar with Badge-Based Layout

**Status**: Done
**Type**: Frontend
**Points**: 5
**Dependencies**: None

## Beschreibung
Als Entwickler muss ich die TuiStatusBar-Komponente umbauen, damit sie farbige Mode-Badges, Context-Badges und gestylte Key-Hints anzeigt, wie in den Design-Mockups definiert.

## Fachliche Beschreibung
Der Bug verursacht eine plain-text 3-Spalten-Darstellung anstatt der badge-basierten Darstellung mit farbigen Pills (Mode-Badge wie "LOGS", "SELECT"), Context-Badges (aktive Komponente) und individuell gestylten Key-Hints ohne Punkt-Separatoren.

## Technische Verfeinerung

### WAS (Anforderungen)
- `FooterItemActive()` Funktion implementieren (definiert in design-system.md, fehlt im Code)
- `TuiStatusBarConfig` erweitern um Mode-Badge und Context-Badge Felder
- Badge-Rendering mit farbigen Hintergründen (Pill-Styling) implementieren
- `TuiKeyHints` Separator-Logik ändern: Spacing statt Punkt-Separatoren
- Layout umbauen: `[Mode Badge] [Context Badge] ... [Key Hints] ... [q Quit]`
- Sicherstellen, dass keine Side-Effects in anderen Komponenten entstehen

### WIE (Implementierung)
1. Add `FooterItemActive()` and `FooterItem()` to `styles.go`
2. Extend `TuiStatusBarConfig` with badge fields
3. Rewrite `TuiStatusBar()` rendering logic for badge-based layout
4. Update `TuiKeyHints()` separator from `" · "` to `"  "`
5. Update `layout_gallery.go` to use new config

### WO (Betroffene Dateien)
- `internal/ui/components/styles.go` — Add FooterItemActive, FooterItem, badge styles
- `internal/ui/components/statusbar.go` — Restructure config and rendering logic
- `internal/ui/components/keyhints.go` — Change separator style
- `internal/ui/components/demo/layout_gallery.go` — Update status bar usage

### WER (Zuständigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

## Definition of Ready (DoR)
- [x] Bug-Beschreibung verstanden
- [x] Reproduction Steps verifiziert
- [x] Betroffene Komponenten identifiziert (statusbar.go, keyhints.go, styles.go)
- [x] Technical-Spec vom Architekten vorhanden

## Definition of Done (DoD)
- [x] Root-Cause dokumentiert
- [x] Fix implementiert (badge layout, FooterItemActive, styled hints)
- [x] StatusBar zeigt Mode-Badge mit farbigem Hintergrund
- [x] StatusBar zeigt Context-Badge mit Selektionsinfo
- [x] KeyHints ohne Punkt-Separatoren, mit Spacing
- [x] "q Quit" visuell getrennt am rechten Rand
- [x] Visuelle Ausgabe stimmt mit should.png und should-2.png überein
- [x] Keine Regression in bestehenden Tests
- [ ] Architect Review bestanden
