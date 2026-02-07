# Exit Confirmation Modal

> Story ID: WELCOME-005
> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Last Updated: 2026-02-07

**Status**: Done
**Priority**: Medium
**Type**: Frontend
**Estimated Effort**: XS
**Dependencies**: WELCOME-002

---

## Feature

```gherkin
Feature: Beendigungs-Bestaetigung
  Als RFZ-Entwickler
  moechte ich vor dem Beenden der Anwendung eine Bestaetigung sehen,
  damit ich nicht versehentlich die Anwendung schliesse.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Bestaetigung erscheint bei q-Taste

```gherkin
Scenario: Bestaetigung beim Druecken von q
  Given ich befinde mich in der Anwendung
  When ich die Taste "q" druecke
  Then erscheint ein Bestaetigungsdialog "Are you sure you want to quit?"
  And der Dialog hat die Optionen "Yes" und "No"
```

### Szenario 2: Bestaetigung bei Exit-Menuepunkt

```gherkin
Scenario: Bestaetigung beim Auswaehlen von Exit
  Given der Cursor steht auf "5. Exit" in der Navigation
  When ich Enter druecke
  Then erscheint der Bestaetigungsdialog
```

### Szenario 3: Anwendung beenden mit Yes

```gherkin
Scenario: Beenden wird bestaetigt
  Given der Bestaetigungsdialog ist sichtbar
  When ich "Yes" auswaehle
  Then wird die Anwendung beendet
```

### Szenario 4: Abbrechen mit No

```gherkin
Scenario: Beenden wird abgebrochen
  Given der Bestaetigungsdialog ist sichtbar
  When ich "No" auswaehle
  Then schliesst sich der Dialog
  And ich befinde mich wieder auf dem vorherigen Bildschirm
```

### Edge Cases & Fehlerszenarien

```gherkin
Scenario: Escape schliesst den Dialog
  Given der Bestaetigungsdialog ist sichtbar
  When ich die Escape-Taste druecke
  Then schliesst sich der Dialog ohne die Anwendung zu beenden
```

```gherkin
Scenario: Modal ueberlagert gesamten Bildschirm
  Given ich befinde mich auf dem Build Components Platzhalter
  When der Bestaetigungsdialog erscheint
  Then ueberlagert der Dialog den gesamten Bildschirminhalt
  And keine Tasten ausser Dialog-Tasten sind aktiv
```

---

## Technische Verifikation (Automated Checks)

### Inhalt-Pruefungen

- [x] CONTAINS: internal/app/app.go enthaelt "showModal"
- [x] CONTAINS: internal/app/app.go enthaelt "TuiModal"
- [x] CONTAINS: internal/app/app.go enthaelt "tea.Quit"

### Funktions-Pruefungen

- [x] BUILD_PASS: `go build ./internal/app/...`
- [x] LINT_PASS: `golangci-lint run ./internal/app/...`

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
- [x] Modal ueberlagert korrekt (captures all input)
- [x] Yes = tea.Quit, No/Esc = dismiss modal

#### Qualitaetssicherung
- [x] Alle Akzeptanzkriterien erfuellt
- [x] Build kompiliert ohne Fehler

#### Dokumentation
- [x] Keine Linting Errors
- [x] Completion Check Commands alle erfolgreich (exit 0)

---

### Betroffene Layer & Komponenten

**Integration Type:** Frontend-only

| Layer | Komponenten | Aenderung |
|-------|-------------|----------|
| Frontend | internal/app/app.go | Modal state + rendering + input handling |

**Kritische Integration Points:**
- internal/app/app.go -> internal/ui/components/modal.go (TuiModal fuer Bestaetigungsdialog)

---

### Technical Details

**WAS:** Exit-Bestaetigungsdialog in app.go integrieren mit TuiModal

**WIE:**
- Neue Felder in app.Model: showModal bool, modalFocusIndex int (0=Yes, 1=No)
- Trigger: Taste "q" (global) oder Enter auf Nav-Item "5. Exit" setzt showModal = true, modalFocusIndex = 1 (No ist Standard-Fokus fuer Sicherheit)
- Update-Logik: Wenn showModal == true, werden ALLE Key-Events vom Modal-Handler abgefangen (kein Durchleiten an Navigation oder Screen)
- Modal-Handler: Links/Rechts-Pfeiltasten oder Tab wechseln modalFocusIndex zwischen 0 und 1, Enter fuehrt fokussierten Button aus, Esc schliesst Modal
- TuiModal() Aufruf mit TuiModalConfig: Title="Quit RFZ-CLI?", Content="Are you sure you want to quit?", Buttons=[{Label:"Yes", Variant:ButtonPrimary, Shortcut:"y"}, {Label:"No", Variant:ButtonSecondary, Shortcut:"n"}], FocusedIndex=modalFocusIndex
- TuiModal() erhaelt termWidth und termHeight fuer Zentrierung mit Backdrop
- View-Logik: Wenn showModal, TuiModal(config, width, height) rendern STATT des normalen Layouts (TuiModal rendert selbst den Backdrop)
- Shortcut-Tasten: "y" -> tea.Quit, "n" -> showModal = false

**WO:**
- internal/app/app.go (erweitern - ~60 LOC zusaetzlich fuer Modal-State und Handler)

**WER:** dev-team__frontend-developer

**Abhaengigkeiten:** WELCOME-002 (App Shell muss existieren mit grundlegendem Update/View-Flow)

**Geschaetzte Komplexitaet:** XS (0 neue Dateien, ~60 LOC Erweiterung)

---

### Completion Check

```bash
go build ./internal/app/...
golangci-lint run ./internal/app/...
```

**Story ist DONE wenn:**
1. Alle CONTAINS checks bestanden
2. Alle BUILD_PASS commands exit 0
3. q-Taste und Exit-Menuepunkt zeigen Bestaetigungsdialog
