# Screen Switching & Navigation

> Story ID: WELCOME-004
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Status**: Done
**Priority**: High
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: WELCOME-002

---

## Feature

```gherkin
Feature: Bildschirmwechsel und Navigation
  Als RFZ-Entwickler
  moechte ich mit Tastenkuerzeln zwischen verschiedenen Bildschirmen wechseln,
  damit ich schnell auf alle Funktionen der Anwendung zugreifen kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Navigation mit Pfeiltasten

```gherkin
Scenario: Cursorbewegung in der Navigation
  Given ich befinde mich auf dem Welcome Screen
  And der Cursor steht auf "1. Build Components"
  When ich die Pfeiltaste nach unten druecke
  Then steht der Cursor auf "2. View Logs"
  And der Statusbar zeigt "View Logs" als fokussierten Eintrag
```

### Szenario 2: Schnellnavigation mit Ziffern

```gherkin
Scenario: Direkter Bildschirmwechsel mit Zifferntaste
  Given ich befinde mich auf dem Welcome Screen
  When ich die Taste "3" druecke
  Then wechselt der Inhaltsbereich zum Discover-Platzhalter
  And ich sehe "Discover - Coming Soon"
  And der Statusbar zeigt "Discover" als aktiven Bildschirm
```

### Szenario 3: Platzhalter-Bildschirme

```gherkin
Scenario Outline: Platzhalter fuer noch nicht implementierte Bildschirme
  Given ich befinde mich auf dem Welcome Screen
  When ich zur Taste "<taste>" druecke
  Then sehe ich den Platzhalter "<titel> - Coming Soon"
  And den Hinweis "Press Esc to return to Welcome"

  Examples:
    | taste | titel |
    | 1 | Build Components |
    | 2 | View Logs |
    | 3 | Discover |
    | 4 | Configuration |
```

### Szenario 4: Fokus-Wechsel mit Tab

```gherkin
Scenario: Fokus zwischen Navigation und Inhalt wechseln
  Given der Fokus liegt auf der Navigationsleiste
  When ich die Tab-Taste druecke
  Then wechselt der Fokus zum Inhaltsbereich
  And die Navigationsleiste verliert den Fokus-Rahmen
```

### Szenario 5: Zurueck zum Welcome Screen

```gherkin
Scenario: Escape-Taste kehrt zum Welcome Screen zurueck
  Given ich befinde mich auf dem "Build Components" Platzhalter
  When ich die Escape-Taste druecke
  Then kehre ich zum Welcome Screen zurueck
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Navigation-Cursor umwickelt am Ende
  Given der Cursor steht auf "5. Exit" (letzter Eintrag)
  When ich die Pfeiltaste nach unten druecke
  Then springt der Cursor zu "1. Build Components" (erster Eintrag)
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [x] FILE_EXISTS: internal/ui/screens/placeholder/placeholder.go

### Inhalt-Pruefungen

- [x] CONTAINS: internal/ui/screens/placeholder/placeholder.go enthaelt "Coming Soon"
- [x] CONTAINS: internal/app/app.go enthaelt "activeScreen"

### Funktions-Pruefungen

- [x] BUILD_PASS: `go build ./internal/ui/screens/placeholder/...`
- [x] BUILD_PASS: `go build ./internal/app/...`
- [x] LINT_PASS: `golangci-lint run ./internal/ui/screens/...`

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
- [x] Code implementiert und folgt Style Guide
- [x] Keyboard Routing korrekt (Global vs Panel vs Screen)
- [x] Fokus-Management implementiert (Tab wechselt Panels)

#### Qualitaetssicherung
- [x] Alle Akzeptanzkriterien erfuellt
- [x] Build kompiliert ohne Fehler
- [x] Alle Platzhalter-Bildschirme korrekt angezeigt

#### Dokumentation
- [x] Keine Linting Errors
- [x] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | internal/ui/screens/placeholder/placeholder.go | Generischer Platzhalter-Screen |
| Frontend | internal/app/app.go | Screen-Routing, Focus-Management, Keyboard-Handling |

**Kritische Integration Points:**
- internal/app/app.go -> internal/ui/screens/placeholder/ (Platzhalter in Content Area rendern)
- Keyboard input -> Navigation state -> Screen switching -> View update

---

### Technical Details

**WAS:** Platzhalter-Screen erstellen, Screen-Switching und Focus-Management in app.go implementieren

**WIE:**
- Placeholder-Screen: Minimales tea.Model mit Title und width/height Feldern, View() rendert zentriert "{Title} - Coming Soon" und "Press Esc to return to Welcome" mit StyleH2 und StyleBodyMuted
- Screen-Konstanten als iota-basierte Enum definieren: ScreenWelcome, ScreenBuild, ScreenLogs, ScreenDiscover, ScreenConfig
- app.Model erhaelt activeScreen Feld und 4 Placeholder-Instanzen neben dem Welcome-Screen
- Globale Key-Behandlung in app.Update(): Tasten "1"-"4" setzen activeScreen auf entsprechenden Screen und aktualisieren cursorIndex/activeIndex
- Tab-Taste wechselt focusPanel zwischen FocusNav und FocusContent (enum)
- Wenn FocusNav: Pfeiltasten hoch/runter bewegen cursorIndex, Enter auf einem Menuepunkt aktiviert den Screen
- Wenn FocusContent: Keys werden an den aktiven Screen weitergeleitet
- Esc-Taste: Wenn nicht auf Welcome, zurueck zu ScreenWelcome setzen
- Cursor-Wrapping: cursorIndex = (cursorIndex + 1) % len(items) bzw. (cursorIndex - 1 + len(items)) % len(items)
- TuiStatusBar Config dynamisch aus activeScreen und cursorIndex zusammenbauen
- Focus-Zustand visuell: TuiBox um Navigation mit focused=true/false, TuiBox um Content mit focused=true/false
- Referenz-Pattern: internal/ui/components/demo/layout_gallery.go fuer Focus-Management Muster

**WO:**
- internal/ui/screens/placeholder/placeholder.go (neu - ~60 LOC)
- internal/app/app.go (erweitern - Screen-Routing, Focus-Management, ~100 LOC zusaetzlich)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** WELCOME-002 (App Shell Model mit Layout muss existieren, da Screen-Routing und Focus-Management in app.go eingebaut werden)

**Geschaetzte Komplexitaet:** S (1 neue Datei + 1 Erweiterung, ~160 LOC)

---

### Completion Check

```bash
go build ./internal/ui/screens/placeholder/...
go build ./internal/app/...
golangci-lint run ./internal/...
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS/CONTAINS checks bestanden
2. Alle BUILD_PASS commands exit 0
3. Zifferntasten 1-4 wechseln zu Platzhaltern, Esc kehrt zurueck
