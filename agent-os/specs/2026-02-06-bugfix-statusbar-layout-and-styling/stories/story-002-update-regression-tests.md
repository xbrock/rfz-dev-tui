# Story 002: Update Regression Tests for StatusBar and KeyHints

**Status**: Open
**Type**: Test
**Points**: 3
**Dependencies**: Story 001

## Beschreibung
Als QA-Spezialist muss ich die bestehenden Golden-Tests aktualisieren und neue Regression-Tests erstellen, damit die neue Badge-basierte StatusBar automatisch getestet wird.

## Fachliche Beschreibung
Die bestehenden Golden-Tests für TuiStatusBar (8 Tests) und TuiKeyHints (6 Tests) müssen an das neue Badge-Layout angepasst werden. Zusätzlich werden neue Tests für Mode-Badges und Context-Badges benötigt.

## Technische Verfeinerung

### WAS (Anforderungen)
- Bestehende Golden-Test-Dateien aktualisieren (`.golden` files)
- Neue Tests für Mode-Badge Rendering (verschiedene Farben)
- Neue Tests für Context-Badge Rendering
- Tests für KeyHints ohne Punkt-Separatoren
- Tests für "q Quit" Separation
- Tests für FooterItemActive() Funktion

### WIE (Implementierung)
- Update all `.golden` files to match new badge-based rendering output
- Add test cases: `TestStatusBar_ModeBadge`, `TestStatusBar_ContextBadge`, `TestStatusBar_BadgeColors`
- Update `TestKeyHints_Multiple` to verify no dot separators
- Add `TestFooterItemActive` for the new function
- Run with `-update` flag to regenerate golden files after visual verification

### WO (Betroffene Dateien)
- `internal/ui/components/statusbar_test.go` — Update existing + add new test cases
- `internal/ui/components/keyhints_test.go` — Update separator tests
- `internal/ui/components/testdata/` — Updated golden files

### WER (Zuständigkeit)
- **Primary**: Frontend-developer
- **Review**: Architect

## Definition of Ready (DoR)
- [ ] Story 1 abgeschlossen
- [ ] Neues Badge-Layout implementiert und verifiziert
- [ ] Test-Strategie festgelegt (golden testing mit charmbracelet/x/exp/golden)

## Definition of Done (DoD)
- [ ] Alle bestehenden Tests aktualisiert
- [ ] Neue Badge-Tests implementiert
- [ ] Alle Tests bestehen mit Fix
- [ ] Golden files regeneriert und verifiziert
- [ ] `go test ./internal/ui/components/...` erfolgreich
