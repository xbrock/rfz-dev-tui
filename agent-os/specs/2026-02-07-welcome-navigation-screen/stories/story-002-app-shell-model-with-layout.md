# App Shell Model with Layout

> Story ID: WELCOME-002
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Priority**: Critical
**Type**: Frontend
**Estimated Effort**: S
**Dependencies**: WELCOME-001

---

## Feature

```gherkin
Feature: App Shell Layout
  Als RFZ-Entwickler
  moechte ich eine Anwendung mit Header, Seitennavigation, Inhaltsbereich und Statusleiste sehen,
  damit ich mich in der Anwendung orientieren und effizient navigieren kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Header-Leiste wird korrekt angezeigt

```gherkin
Scenario: Header zeigt Titel und aktuelle Uhrzeit
  Given ich starte die RFZ-CLI Anwendung
  When die Anwendung vollstaendig geladen ist
  Then sehe ich "RFZ-CLI v1.0.0" oben links
  And darunter "Terminal Orchestration Tool" in Cyan
  And rechts die aktuelle Uhrzeit im Format "HH:MM:SS PM"
  And rechts daneben "Deutsche Bahn Internal"
```

### Szenario 2: Navigation Sidebar wird angezeigt

```gherkin
Scenario: Navigation mit 5 Menuepunkten
  Given ich starte die RFZ-CLI Anwendung
  When die Anwendung vollstaendig geladen ist
  Then sehe ich links eine Navigationsleiste mit dem Titel "Navigation"
  And die Menuepunkte "1. Build Components", "2. View Logs", "3. Discover", "4. Configuration", "5. Exit"
  And unterhalb der Menuepunkte Tastaturhinweise
```

### Szenario 3: Status Bar am unteren Rand

```gherkin
Scenario: Statusleiste zeigt aktuellen Kontext
  Given ich befinde mich auf dem Welcome Screen
  When ich den unteren Bildschirmrand betrachte
  Then sehe ich ein "HOME" Badge
  And den Namen des fokussierten Menuepunkts
  And Tastaturhinweise fuer die Navigation
  And rechts "q Quit"
```

### Szenario 4: Uhrzeit aktualisiert sich

```gherkin
Scenario: Header-Uhrzeit wird jede Sekunde aktualisiert
  Given die Anwendung zeigt die aktuelle Uhrzeit "3:40:33 PM"
  When 1 Sekunde vergeht
  Then zeigt die Uhrzeit "3:40:34 PM"
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Terminal zu klein
  Given mein Terminal ist kleiner als 80x24 Zeichen
  When ich die Anwendung starte
  Then sehe ich eine Meldung "Terminal too small. Please resize to at least 80x24."
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/app/app.go
- [ ] FILE_EXISTS: internal/app/messages.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/app/app.go enthaelt "type Model struct"
- [ ] CONTAINS: internal/app/app.go enthaelt "func (m Model) Init()"
- [ ] CONTAINS: internal/app/app.go enthaelt "func (m Model) Update("
- [ ] CONTAINS: internal/app/app.go enthaelt "func (m Model) View()"
- [ ] CONTAINS: internal/app/app.go enthaelt "tea.WindowSizeMsg"
- [ ] CONTAINS: internal/app/messages.go enthaelt "TickMsg"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `go build ./internal/app/...`
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
- [ ] Code implementiert und folgt Style Guide
- [ ] Architektur-Vorgaben eingehalten (DEC-002 Hierarchical Model)
- [ ] Charm.land First Regel eingehalten

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
| Frontend | internal/app/app.go | Root Bubble Tea Model mit Layout-Komposition |
| Frontend | internal/app/messages.go | Shared message types (TickMsg, NavigateMsg) |

**Kritische Integration Points:**
- internal/app/app.go -> internal/ui/components/navigation.go (TuiNavigation aufrufen)
- internal/app/app.go -> internal/ui/components/statusbar.go (TuiStatusBar aufrufen)

---

### Technical Details

**WAS:** Root Bubble Tea Model erstellen mit Header, Navigation Sidebar, Content Area und StatusBar Layout

**WIE:**
- Follow DEC-002: Hierarchical Model Composition - app.Model ist der Root-Knoten, der alle Screen-Models als Felder enthaelt
- Layout-Komposition in View() via lipgloss.JoinVertical (Header, Body, StatusBar) und lipgloss.JoinHorizontal (Navigation, Content) innerhalb des Body
- TuiNavigation() aus internal/ui/components/navigation.go direkt aufrufen mit 5 TuiNavItem Eintraegen (Build, Logs, Discover, Config, Exit), cursorIndex, activeIndex, focused-Flag
- TuiStatusBar() aus internal/ui/components/statusbar.go direkt aufrufen mit TuiStatusBarConfig (ModeBadge, ContextBadge, Hints, QuitHint, Width)
- Header-Rendering mit bestehenden StyleHeader, StyleHeaderTitle, StyleHeaderSubtitle aus styles.go
- tea.Model Interface implementieren: Init() liefert TickCmd, Update() routet Messages, View() komponiert Layout
- tea.Tick fuer 1-Sekunden Uhrzeitaktualisierung im Header (time.Now() formatiert als "3:04:05 PM")
- tea.WindowSizeMsg Handler speichert width/height und berechnet Panel-Dimensionen
- Navigation-Panel feste Breite ~30 Zeichen, Content Area fuellt den Rest (width - navWidth - Borders)
- Minimum-Terminal-Groesse pruefen (80x24), bei Unterschreitung Warnmeldung rendern statt Layout
- Focus-State als enum/int (FocusNav, FocusContent) fuer spaetere Tab-Navigation vorbereiten

**WO:**
- internal/app/app.go (neu - Root Model mit Layout, ~200-300 LOC)
- internal/app/messages.go (neu - TickMsg, NavigateMsg, ~30 LOC)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** WELCOME-001 (Entry Point muss existieren, damit cmd/rfz/main.go das app Package importieren kann)

**Geschaetzte Komplexitaet:** S (2 Dateien, ~250 LOC)

---

### Completion Check

```bash
go build ./internal/app/...
go build ./cmd/rfz/...
golangci-lint run ./internal/app/...
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS/CONTAINS checks bestanden
2. Alle BUILD_PASS commands exit 0
3. App startet und zeigt Header, Navigation, Content Area, StatusBar
