# TuiStatus Component

> Story ID: CORE-005
> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02

**Priority**: Critical
**Type**: Backend
**Estimated Effort**: S (2 SP)
**Dependencies**: CORE-001

---

## Feature

```gherkin
Feature: TuiStatus Build Status Badge Component
  Als TUI-Entwickler
  möchte ich Status-Badges für Build-Zustände anzeigen können,
  damit Benutzer den aktuellen Status auf einen Blick erkennen.
```

---

## Akzeptanzkriterien (Gherkin-Szenarien)

### Szenario 1: Success Status Badge

```gherkin
Scenario: Success Status wird grün angezeigt
  Given ein Build war erfolgreich
  When ich TuiStatus mit StatusSuccess aufrufe
  Then wird ein grüner Badge mit "SUCCESS" angezeigt
  And der Text ist fett und gut lesbar
```

### Szenario 2: Running Status Badge

```gherkin
Scenario: Running Status wird cyan angezeigt
  Given ein Build läuft gerade
  When ich TuiStatus mit StatusRunning aufrufe
  Then wird ein cyan Badge mit "RUNNING" angezeigt
```

### Szenario Outline: Alle Status-Varianten

```gherkin
Scenario Outline: TuiStatus unterstützt alle Build-Zustände
  Given ein Build hat Status <status>
  When ich TuiStatus aufrufe
  Then wird "<label>" mit <color> Hintergrund angezeigt

  Examples:
    | status        | label    | color |
    | StatusPending | PENDING  | Gray  |
    | StatusRunning | RUNNING  | Cyan  |
    | StatusSuccess | SUCCESS  | Green |
    | StatusFailed  | FAILED   | Red   |
    | StatusError   | ERROR    | Red   |
```

### Szenario 3: Compact Status Format

```gherkin
Scenario: Compact Status für Listen-Ansichten
  Given ich habe wenig Platz in einer Liste
  When ich TuiStatusCompact mit StatusSuccess aufrufe
  Then wird nur ein grünes "✓" Symbol angezeigt
  And kein Text
```

### Edge Case: Unknown Status

```gherkin
Scenario: Unbekannter Status wird behandelt
  Given ein ungültiger Status-Wert wird übergeben
  When ich TuiStatus aufrufe
  Then wird ein grauer "UNKNOWN" Badge angezeigt
  And es tritt kein Fehler auf
```

---

## Technische Verifikation (Automated Checks)

### Datei-Prüfungen

- [ ] FILE_EXISTS: internal/ui/components/status.go

### Inhalt-Prüfungen

- [ ] CONTAINS: status.go enthält "type Status int"
- [ ] CONTAINS: status.go enthält "StatusPending"
- [ ] CONTAINS: status.go enthält "StatusRunning"
- [ ] CONTAINS: status.go enthält "StatusSuccess"
- [ ] CONTAINS: status.go enthält "StatusFailed"
- [ ] CONTAINS: status.go enthält "StatusError"
- [ ] CONTAINS: status.go enthält "func TuiStatus("
- [ ] CONTAINS: status.go enthält "func TuiStatusCompact("

### Funktions-Prüfungen

- [ ] BUILD_PASS: `go build ./internal/ui/components/...`
- [ ] LINT_PASS: `golangci-lint run ./internal/ui/components/status.go`

---

## Required MCP Tools

None required.

---

## Technisches Refinement (vom Architect)

### DoR (Definition of Ready) - Vom Architect

#### Fachliche Anforderungen
- [x] Fachliche requirements klar definiert
- [x] Akzeptanzkriterien sind spezifisch und prüfbar
- [x] Business Value verstanden

#### Technische Vorbereitung
- [x] Technischer Ansatz definiert (WAS/WIE/WO)
- [x] Abhängigkeiten identifiziert
- [x] Betroffene Komponenten bekannt
- [x] Erforderliche MCP Tools dokumentiert (falls zutreffend)
- [x] Story ist angemessen geschätzt (max 5 Dateien, 400 LOC)

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
- [ ] Security/Performance Anforderungen erfüllt

#### Qualitätssicherung
- [ ] Alle Akzeptanzkriterien erfüllt
- [ ] Unit Tests geschrieben und bestanden
- [ ] Code Review durchgeführt und genehmigt

#### Dokumentation
- [ ] Dokumentation aktualisiert
- [ ] Keine Linting Errors
- [ ] Completion Check Commands alle erfolgreich

---

### Betroffene Layer & Komponenten

**Integration Type:** Backend-only

| Layer | Komponenten | Änderung |
|-------|-------------|----------|
| Backend | internal/ui/components/status.go | CREATE - TuiStatus component |

**Kritische Integration Points:**
- status.go → styles.go: Verwendet ColorGreen, ColorCyan, ColorDestructive, ColorSecondary

---

### Technical Details

**WER:** dev-team__go-developer

**WAS:**
- Status enum (iota) mit 5 Zuständen
- TuiStatus Funktion für volle Badge-Darstellung
- TuiStatusCompact für Symbol-Darstellung
- String() Methode für Status-Labels

**WIE:**
- Status als int type mit iota:
  ```go
  type Status int
  const (
      StatusPending Status = iota
      StatusRunning
      StatusSuccess
      StatusFailed
      StatusError
  )
  ```
- Main functions:
  - `func TuiStatus(status Status) string` - Full badge with label
  - `func TuiStatusCompact(status Status) string` - Icon only
  - `func (s Status) String() string` - Returns status label
- Status colors and labels:
  - Pending: ColorSecondary (gray), "PENDING", icon "○"
  - Running: ColorCyan, "RUNNING", icon "●"
  - Success: ColorGreen, "SUCCESS", icon "✓"
  - Failed: ColorDestructive, "FAILED", icon "✗"
  - Error: ColorDestructive, "ERROR", icon "!"
- Badge style: Background(color), Foreground(ColorBackground), Bold(true), Padding(0, 1)
- Edge case: Unknown status values return gray "UNKNOWN" badge

**WO:**
- `internal/ui/components/status.go` (NEW) - ~130 LOC

**Abhängigkeiten:** CORE-001 (styles.go)

**Geschätzte Komplexität:** S (2 SP)

**Relevante Skills:** N/A (no skill-index.md in project)

---

### Completion Check

```bash
# Verify file exists
test -f internal/ui/components/status.go && echo "status.go exists"

# Verify required types and functions
grep -q "type Status int" internal/ui/components/status.go && echo "Status type found"
grep -q "func TuiStatus" internal/ui/components/status.go && echo "TuiStatus found"
grep -q "func TuiStatusCompact" internal/ui/components/status.go && echo "TuiStatusCompact found"

# Verify build
go build ./internal/ui/components/...

# Verify lint
golangci-lint run ./internal/ui/components/status.go
```

**Story ist DONE wenn:**
1. Alle FILE_EXISTS checks bestanden
2. Alle CONTAINS checks bestanden
3. BUILD_PASS und LINT_PASS erfolgreich
