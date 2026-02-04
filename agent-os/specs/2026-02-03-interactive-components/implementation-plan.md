# Implementation Plan - Interactive Components (Sprint 1.2)

**Created:** 2026-02-03
**Status:** Pending Approval
**Spec:** 2026-02-03-interactive-components

---

## Executive Summary

Build 6 interactive TUI components (TuiList, TuiCheckbox, TuiRadio, TuiTextInput, TuiSpinner, TuiProgress) using charm-style Unicode symbols and Bubbles library wrappers, then extend the existing component gallery to showcase them with full visual regression test coverage.

---

## Architecture Decisions

### Component Architecture Strategy

Based on codebase exploration of existing components (`box.go`, `button.go`, `status.go`), we follow two patterns:

| Pattern | Components | Rationale |
|---------|------------|-----------|
| **Stateless Render Functions** | TuiCheckbox, TuiRadio | Pure functions returning strings, state managed by caller |
| **Bubbles Wrapper Models** | TuiSpinner, TuiProgress, TuiTextInput | Need animation/internal state, wrap Bubbles with styling |
| **Hybrid (Render + Model)** | TuiList | Render function for items + optional Model for full list with viewport |

### Charm-Style Design Tokens

Add to `styles.go`:

```go
// Charm-style symbols
const (
    SymbolCheckboxUnchecked = "☐"  // U+2610
    SymbolCheckboxChecked   = "☑"  // U+2611
    SymbolRadioUnselected   = "◯"  // U+25EF
    SymbolRadioSelected     = "◉"  // U+25C9
    SymbolCursor            = "›"  // U+203A (guillemet)
)
```

### Bubbles Integration

Existing dependency: `github.com/charmbracelet/bubbles v0.18.0`

| Bubbles Component | Our Wrapper | Customization |
|-------------------|-------------|---------------|
| `bubbles/spinner` | TuiSpinner | Custom frame sets (braille, line, circle, bounce) |
| `bubbles/progress` | TuiProgress | Custom colors (yellow→green gradient), block chars |
| `bubbles/textinput` | TuiTextInput | RFZ styles (border, colors, prompt) |

---

## Komponenten-Übersicht

### Neue Komponenten

| Komponente | Datei | Typ | Abhängigkeit |
|------------|-------|-----|--------------|
| TuiCheckbox | `checkbox.go` | Stateless Function | styles.go |
| TuiRadio | `radio.go` | Stateless Function | styles.go |
| TuiList | `list.go` | Stateless + Optional Model | checkbox.go, styles.go |
| TuiSpinner | `spinner.go` | Bubbles Wrapper Model | bubbles/spinner |
| TuiProgress | `progress.go` | Bubbles Wrapper Model | bubbles/progress |
| TuiTextInput | `textinput.go` | Bubbles Wrapper Model | bubbles/textinput |

### Änderungen an Bestehendem

| Datei | Änderung |
|-------|----------|
| `styles.go` | Add charm symbol constants + new style tokens |
| `demo/gallery.go` | Add 6 new render sections |

### Komponenten-Verbindungen

| Source | Target | Verbindungstyp | Zuständige Story |
|--------|--------|----------------|------------------|
| TuiCheckbox | TuiList | Used by (list items) | INTER-001 |
| TuiRadio | styles.go | Uses symbols from | INTER-003 |
| TuiSpinner | bubbles/spinner | Wraps | INTER-005 |
| TuiProgress | bubbles/progress | Wraps | INTER-006 |
| TuiTextInput | bubbles/textinput | Wraps | INTER-004 |
| All components | demo/gallery.go | Showcased in | INTER-007 |

---

## Umsetzungsphasen

### Phase 1: Foundation (Story 1-2)
**Goal:** Add charm symbols to styles.go, implement simplest stateless components

1. Add charm symbol constants to `styles.go`
2. Implement `TuiCheckbox` - simplest component, single render function
3. Implement unit tests for TuiCheckbox

### Phase 2: Form Controls (Story 3-4)
**Goal:** Complete stateless form controls

4. Implement `TuiRadio` - similar pattern to checkbox but with group logic
5. Implement `TuiTextInput` - wrap Bubbles textinput with RFZ styling
6. Add unit tests for both

### Phase 3: Feedback Components (Story 5-6)
**Goal:** Implement animated/stateful feedback components

7. Implement `TuiSpinner` - wrap Bubbles spinner with 4 custom frame sets
8. Implement `TuiProgress` - wrap Bubbles progress with 4 custom styles
9. Add unit tests for both

### Phase 4: List Component (Story 1 - complex)
**Goal:** Most complex component combining others

10. Implement `TuiList` render functions (item rendering with checkbox/cursor)
11. Implement optional `TuiListModel` for full interactive list with viewport
12. Add comprehensive tests

### Phase 5: Gallery & Visual Tests (Story 7-8)
**Goal:** Showcase and validate all components

13. Extend `demo/gallery.go` with 6 new sections
14. Create golden file tests for all component states
15. Verify all 97+ visual states pass

---

## Abhängigkeiten

```
styles.go (charm symbols)
    │
    ├──► TuiCheckbox
    │       │
    │       └──► TuiList (uses checkbox for multi-select)
    │
    ├──► TuiRadio
    │       │
    │       └──► TuiList (uses radio for single-select)
    │
    ├──► TuiSpinner (wraps bubbles/spinner)
    │
    ├──► TuiProgress (wraps bubbles/progress)
    │
    └──► TuiTextInput (wraps bubbles/textinput)

All ──► demo/gallery.go ──► Visual Tests
```

**Critical Path:** styles.go → TuiCheckbox → TuiRadio → TuiList → Gallery → Tests

---

## Risiken & Mitigationen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| Unicode symbol rendering inconsistency | Medium | Medium | Test on multiple terminals; provide ASCII fallback |
| Bubbles API changes | Low | High | Pin exact version in go.mod (already v0.18.0) |
| List performance with 100+ items | Low | Medium | Use viewport for virtualization |
| Golden test brittleness | Medium | Low | Use structural assertions where possible |

---

## Self-Review Ergebnisse

### 1. VOLLSTÄNDIGKEIT
- ✅ All 6 components from requirements covered
- ✅ All charm-style symbols specified
- ✅ Gallery extension included
- ✅ Visual tests included

### 2. KONSISTENZ
- ✅ Pattern aligns with existing components (box.go, button.go)
- ✅ Uses same color tokens and style system
- ✅ Follows Bubble Tea architecture

### 3. RISIKEN
- Identified: Unicode rendering, API stability, performance
- Mitigations: Fallbacks, version pinning, viewport usage

### 4. ALTERNATIVEN
- **Considered:** Building spinner/progress from scratch
- **Rejected:** Violates "Charm.land First" rule, Bubbles already provides excellent implementations
- **Decision:** Wrap Bubbles with custom styling

### 5. KOMPONENTEN-VERBINDUNGEN
- ✅ All new components have at least one connection
- ✅ All connections have assigned stories
- ✅ No orphan components
- ✅ Gallery integration ensures all components are exercised

---

## Minimalinvasiv-Optimierungen

### Wiederverwendung

| Bestehender Code | Wiederverwendung |
|------------------|------------------|
| `styles.go` color tokens | All new components use existing colors |
| `helpers.go` Truncate() | TuiList item label truncation |
| `demo/gallery.go` pattern | Follow existing section render pattern |
| Bubbles viewport | Already used in gallery, reuse for TuiList |
| Existing test patterns | Follow `*_test.go` patterns |

### Änderungsumfang

| Aktion | Dateien | LOC (geschätzt) |
|--------|---------|-----------------|
| New files | 6 component files + 6 test files | ~800 LOC |
| Modified files | styles.go (+20), gallery.go (+200) | ~220 LOC |
| **Total** | 14 files | ~1020 LOC |

### Feature-Preservation Checkliste

- [x] All requirements from Clarification are abgedeckt
- [x] Kein Feature wurde geopfert
- [x] Alle Akzeptanzkriterien bleiben erfüllbar
- [x] Minimalinvasiv: Using existing patterns, no architectural changes

---

## Story Ableitung

Based on this plan, stories will be derived as follows:

| Story ID | Title | Phase | Abhängig von |
|----------|-------|-------|--------------|
| INTER-001 | TuiList Component | 4 | INTER-002, INTER-003 |
| INTER-002 | TuiCheckbox Component | 1 | - |
| INTER-003 | TuiRadio Component | 2 | - |
| INTER-004 | TuiTextInput Component | 2 | - |
| INTER-005 | TuiSpinner Component | 3 | - |
| INTER-006 | TuiProgress Component | 3 | - |
| INTER-007 | Extend Component Gallery | 5 | All above |
| INTER-008 | Visual Regression Tests | 5 | INTER-007 |

**Parallel Execution Possible:** INTER-002, INTER-003, INTER-004, INTER-005, INTER-006 (after styles.go update)

---

## Validierungsbefehle

```bash
# Build validation
go build ./cmd/rfz/...

# Unit tests
go test ./internal/ui/components/... -v

# Visual regression tests
go test ./internal/ui/components/demo/... -v

# Lint check
golangci-lint run ./internal/ui/components/...
```

---

*Plan erstellt mit Agent OS /create-spec v2.8*
