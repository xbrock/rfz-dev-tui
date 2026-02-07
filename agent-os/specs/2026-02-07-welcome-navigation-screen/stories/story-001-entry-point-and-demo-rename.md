# Entry Point & Demo Rename

> Story ID: WELCOME-001
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: High
**Type**: Frontend
**Estimated Effort**: XS
**Dependencies**: None

---

## Feature

```gherkin
Feature: Application Entry Point
  Als RFZ-Entwickler
  moechte ich die RFZ-CLI ueber den Befehl "go run ./cmd/rfz" starten,
  damit ich die Anwendung wie gewohnt ausfuehren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Anwendung startet erfolgreich

```gherkin
Scenario: RFZ-CLI startet als TUI-Anwendung
  Given ich befinde mich im Projektverzeichnis rfz-tui
  When ich "go run ./cmd/rfz" ausfuehre
  Then startet eine Bubble Tea Anwendung im Alt-Screen-Modus
  And ich sehe die App Shell mit Header und Navigation
```

### Szenario 2: Component Gallery Demo bleibt verfuegbar

```gherkin
Scenario: Component Gallery Demo unter neuem Namen
  Given ich befinde mich im Projektverzeichnis rfz-tui
  When ich "go run ./cmd/rfz-components-demo" ausfuehre
  Then startet die bisherige Component Gallery Demo
  And alle Komponenten-Demos werden korrekt angezeigt
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Beide Programme kompilieren ohne Fehler
  Given das Projekt hat keine Compile-Fehler
  When ich "go build ./cmd/..." ausfuehre
  Then kompilieren beide Programme (rfz und rfz-components-demo) erfolgreich
  And es gibt keine Lint-Warnungen
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: cmd/rfz/main.go
- [ ] FILE_EXISTS: cmd/rfz-components-demo/main.go
- [ ] FILE_NOT_EXISTS: cmd/layout-demo/main.go

### Inhalt-Pruefungen

- [ ] CONTAINS: cmd/rfz/main.go enthaelt "internal/app"
- [ ] CONTAINS: cmd/rfz-components-demo/main.go enthaelt "demo.New()"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `go build ./cmd/rfz/...`
- [ ] BUILD_PASS: `go build ./cmd/rfz-components-demo/...`
- [ ] LINT_PASS: `golangci-lint run ./cmd/...`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

> Wird in Step 3 vom Architect ausgefuellt

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
- [ ] Code implementiert und folgt Style Guide
- [ ] Architektur-Vorgaben eingehalten
- [ ] Security/Performance Anforderungen erfuellt

#### Qualitaetssicherung
- [ ] Alle Akzeptanzkriterien erfuellt
- [ ] Build kompiliert ohne Fehler
- [ ] Code Review durchgefuehrt und genehmigt

#### Dokumentation
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | cmd/rfz/main.go | Neuer Entry Point fuer App Shell |
| Frontend | cmd/rfz-components-demo/main.go | Umbenannte Gallery Demo |
| Frontend | cmd/layout-demo/main.go | Zusammengelegt in rfz-components-demo |

---

### Technical Details

**WAS:** Neuen Entry Point fuer die echte Anwendung erstellen, bestehende Gallery-Demo umbenennen

**WIE:**
- cmd/rfz/main.go: Aktuellen demo.New() Aufruf ersetzen durch app.New() aus dem neuen internal/app Package, tea.NewProgram mit WithAltScreen beibehalten
- cmd/rfz-components-demo/main.go: Bisherigen Gallery-Code aus cmd/rfz/main.go hierhin verschieben (demo.New() + tea.NewProgram Muster bleibt identisch)
- cmd/layout-demo/: In rfz-components-demo integrieren oder separat behalten
- Beide Entry Points muessen unabhaengig kompilierbar und ausfuehrbar sein
- Kein neues Package anlegen - nur Dateien verschieben und den Import-Pfad in cmd/rfz/main.go auf internal/app aendern

**WO:**
- cmd/rfz/main.go (neu schreiben - importiert internal/app statt demo)
- cmd/rfz-components-demo/main.go (umbenannt von cmd/rfz/ - importiert weiterhin demo)
- cmd/layout-demo/main.go (optional umbenennen)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** Keine - dies ist die erste Story ohne Vorbedingungen. internal/app/app.go muss als Stub existieren (leeres Model mit Init/Update/View), damit cmd/rfz kompiliert.

**Geschaetzte Komplexitaet:** XS (2 Dateien, ~30 LOC)

---

### Completion Check

```bash
# Auto-Verify Commands
go build ./cmd/rfz/...
go build ./cmd/rfz-components-demo/...
golangci-lint run ./cmd/...
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS/CONTAINS checks bestanden
2. Alle BUILD_PASS commands exit 0
3. Git diff zeigt nur erwartete Aenderungen
