# Welcome Screen

> Story ID: WELCOME-003
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Status**: Done
**Priority**: High
**Type**: Frontend
**Estimated Effort**: XS
**Dependencies**: WELCOME-002

---

## Feature

```gherkin
Feature: Welcome Screen
  Als RFZ-Entwickler
  moechte ich beim Start der Anwendung einen informativen Willkommensbildschirm sehen,
  damit ich sofort erkenne welches Tool ich nutze und wie ich starten kann.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: ASCII Art Logo wird angezeigt

```gherkin
Scenario: Willkommensbildschirm zeigt RFZ-CLI Logo
  Given ich starte die RFZ-CLI Anwendung
  When der Welcome Screen geladen ist
  Then sehe ich ein grosses ASCII-Art Logo "RFZ-CLI"
  And "RFZ" ist in Rot/Magenta dargestellt
  And "CLI" ist in Cyan dargestellt
```

### Szenario 2: Informationen und Badges

```gherkin
Scenario: Willkommensbildschirm zeigt Version und Infos
  Given ich befinde mich auf dem Welcome Screen
  When ich den mittleren Bereich betrachte
  Then sehe ich den Untertitel "Terminal Orchestration Tool"
  And ein Zitat in kursiver Schrift
  And ein rotes Version-Badge "v1.0.0"
  And den Text "Deutsche Bahn"
  And ein umrahmtes Badge "Internal Tool"
```

### Szenario 3: Bereitschaftsanzeige

```gherkin
Scenario: Willkommensbildschirm zeigt Bereitschaft
  Given ich befinde mich auf dem Welcome Screen
  When ich den unteren Bereich des Inhalts betrachte
  Then sehe ich "$ rfz-cli ready"
  And den Hinweis "Use navigation panel to get started"
  And Tastaturhinweise zum Navigieren
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Welcome Screen bei schmalem Terminal
  Given mein Terminal hat eine Breite von 80 Zeichen
  When der Welcome Screen dargestellt wird
  Then wird der Inhalt zentriert angezeigt
  And das ASCII-Art Logo passt in die verfuegbare Breite
```

---

## Technische Verifikation (Automated Checks)

### Datei-Pruefungen

- [ ] FILE_EXISTS: internal/ui/screens/welcome/welcome.go

### Inhalt-Pruefungen

- [ ] CONTAINS: internal/ui/screens/welcome/welcome.go enthaelt "RFZ"
- [ ] CONTAINS: internal/ui/screens/welcome/welcome.go enthaelt "Terminal Orchestration Tool"
- [ ] CONTAINS: internal/ui/screens/welcome/welcome.go enthaelt "v1.0.0"

### Funktions-Pruefungen

- [ ] BUILD_PASS: `go build ./internal/ui/screens/welcome/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/screens/welcome/...`

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
- [x] Charm.land First Regel eingehalten (alle Styling via Lip Gloss)
- [x] Design System Farben korrekt verwendet

#### Qualitaetssicherung
- [x] Alle Akzeptanzkriterien erfuellt
- [x] Build kompiliert ohne Fehler
- [x] Code Review durchgefuehrt und genehmigt

#### Dokumentation
- [x] Keine Linting Errors
- [x] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | internal/ui/screens/welcome/welcome.go | Neuer Welcome Screen mit ASCII Art und Badges |

**Kritische Integration Points:**
- internal/app/app.go -> internal/ui/screens/welcome/ (Welcome Screen in Content Area rendern)

---

### Technical Details

**WAS:** Welcome Screen Model und View erstellen mit ASCII Art Logo, Badges und Hilfehinweisen

**WIE:**
- Minimales tea.Model mit width/height Feldern (wird vom Parent via WindowSizeMsg gesetzt)
- Konstruktor New(width, height int) liefert initialisiertes Model
- View() rendert vertikal zentriert: ASCII Art Logo, Subtitle, Zitat, Divider, Badges, Status-Zeile, Key-Hints
- ASCII Art Logo: "RFZ" in StyleASCIIArt (ColorBrand/rot), "CLI" in StyleASCIIArtCyan (ColorCyan) - als mehrzeilige String-Konstante definieren
- Subtitle "Terminal Orchestration Tool" mit StyleHeaderSubtitle
- Zitat/Tagline in StyleTagline (italic, muted)
- Badge-Zeile: StyleBadgeVersion fuer "v1.0.0", StyleBody fuer "Deutsche Bahn", StyleBadgeInfo fuer "Internal Tool" - horizontal verbunden
- Status-Zeile: StylePrompt fuer "$" + StyleBody fuer "rfz-cli ready"
- Zentrierung: lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content) oder lipgloss.NewStyle().Width(width).Align(lipgloss.Center)
- SetSize(width, height) Methode fuer Resize vom Parent
- Referenz-Screenshot: references/prototype-screenshots/01-welcome-default.png

**WO:**
- internal/ui/screens/welcome/welcome.go (neu - ~120 LOC)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** WELCOME-002 (App Shell muss existieren, damit Welcome als Child-Screen eingebettet werden kann)

**Geschaetzte Komplexitaet:** XS (1 Datei, ~120 LOC)

---

### Completion Check

```bash
go build ./internal/ui/screens/welcome/...
golangci-lint run ./internal/ui/screens/welcome/...
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS/CONTAINS checks bestanden
2. Alle BUILD_PASS commands exit 0
3. Welcome Screen zeigt Logo, Badges und Hints korrekt
