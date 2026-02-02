# Finalize PR

> Story ID: CORE-999
> Spec: Core Components
> Created: 2026-02-02
> Type: System/Finalization

**Priority**: Critical
**Type**: System
**Estimated Effort**: XS (1 SP)
**Dependencies**: CORE-998

---

## Purpose

Ersetzt Phase 5 - Test-Szenarien dokumentieren, User-Todos erstellen, PR erstellen, Worktree Cleanup.

---

## Feature

```gherkin
Feature: PR Finalization
  Als Entwickler
  möchte ich einen sauberen PR mit Dokumentation erstellen,
  damit das Feature bereit für Merge ist.
```

---

## Akzeptanzkriterien

### Szenario 1: PR erstellen

```gherkin
Scenario: PR wird mit vollständiger Dokumentation erstellt
  Given die Integration Validation (CORE-998) ist bestanden
  When der PR erstellt wird
  Then enthält der PR eine vollständige Beschreibung
  And alle Test-Szenarien sind dokumentiert
  And User-Todos sind aufgelistet (falls vorhanden)
```

---

## Checkliste

### PR Vorbereitung
- [ ] Alle Changes committed
- [ ] Branch ist aktuell mit main
- [ ] Commit Messages sind aussagekräftig

### PR Inhalt
- [ ] Titel beschreibt Feature kurz
- [ ] Beschreibung enthält:
  - Summary der Änderungen
  - Test Plan
  - Screenshots (falls UI relevant)
- [ ] Labels gesetzt (falls vorhanden)

### Post-PR
- [ ] Worktree aufgeräumt (falls verwendet)
- [ ] Temporäre Dateien entfernt

---

## Technisches Refinement

### DoR (Definition of Ready)
- [x] CORE-998 (Integration Validation) bestanden
- [x] Alle Tests grün
- [x] Keine Lint Errors

### DoD (Definition of Done)
- [ ] PR erstellt
- [ ] PR Beschreibung vollständig
- [ ] Ready for Review

### Completion Check

```bash
# Verify clean state before PR
go build ./internal/ui/components/... && \
go test ./internal/ui/components/... && \
golangci-lint run ./internal/ui/components/... && \
git status && \
echo "Ready for PR creation"
```

---

## PR Template

```markdown
## Summary
- Complete TUI component library (Phase 1, Week 1)
- Styles package with all design system tokens
- Core components: TuiBox, TuiDivider, TuiButton, TuiStatus
- teatest visual testing infrastructure
- Component gallery demo screen

## Test Plan
- [ ] `go test ./internal/ui/components/... -v` passes
- [ ] Golden files match expected output
- [ ] `golangci-lint run` passes
- [ ] Component gallery displays all variants correctly

## Files Changed
- internal/ui/components/styles.go (NEW)
- internal/ui/components/helpers.go (NEW)
- internal/ui/components/box.go (NEW)
- internal/ui/components/divider.go (NEW)
- internal/ui/components/button.go (NEW)
- internal/ui/components/status.go (NEW)
- internal/ui/components/*_test.go (NEW)
- internal/ui/components/demo/gallery.go (NEW)
- testdata/golden/components/* (NEW)
```

---

*System Story - Final step in execution*
