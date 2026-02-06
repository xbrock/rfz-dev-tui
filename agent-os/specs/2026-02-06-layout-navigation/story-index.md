# Story Index: Layout & Navigation Components

> Spec: 2026-02-06-layout-navigation
> Created: 2026-02-06
> Last Updated: 2026-02-06

---

## Story Summary

| ID | Story | Status | Type | Est. | Dependencies | Blocked By |
|----|-------|--------|------|------|--------------|------------|
| LAYOUT-001 | TuiNavigation + TuiNavItem | Ready | Frontend | S | None | None |
| LAYOUT-002 | TuiModal | Ready | Frontend | M | None | None |
| LAYOUT-003 | TuiKeyHints | Ready | Frontend | XS | None | None |
| LAYOUT-004 | TuiTable | Ready | Frontend | S | None | None |
| LAYOUT-005 | TuiTree | Ready | Frontend | M | None | None |
| LAYOUT-006 | TuiTabs | Ready | Frontend | S | None | None |
| LAYOUT-007 | TuiStatusBar | Ready | Frontend | S | LAYOUT-003 | LAYOUT-003 |
| LAYOUT-008 | Layout Navigation Demo | Ready | Frontend | S | LAYOUT-001 to 007 | LAYOUT-001 to 007 |
| LAYOUT-009 | Visual Regression Tests | Ready | Test | S | LAYOUT-001 to 007 | LAYOUT-001 to 007 |
| LAYOUT-997 | Code Review by Opus | Ready | Review | S | LAYOUT-001 to 009 | LAYOUT-001 to 009 |
| LAYOUT-998 | Integration Validation | Ready | Validation | XS | LAYOUT-997 | LAYOUT-997 |
| LAYOUT-999 | Finalize and Create PR | Ready | Release | XS | LAYOUT-998 | LAYOUT-998 |

---

## Dependency Graph

```
LAYOUT-001 (Navigation) ──┐
LAYOUT-002 (Modal) ───────┤
LAYOUT-003 (KeyHints) ────┼──→ LAYOUT-007 (StatusBar) ──┐
LAYOUT-004 (Table) ───────┤                             ├──→ LAYOUT-008 (Demo) ──┐
LAYOUT-005 (Tree) ────────┤                             │        │               │
LAYOUT-006 (Tabs) ────────┘                             │        ▼               │
                                                        └──→ LAYOUT-009 (Tests) ─┤
                                                                                  │
                                                             LAYOUT-997 (Review) ◄┘
                                                                     │
                                                                     ▼
                                                        LAYOUT-998 (Validation)
                                                                     │
                                                                     ▼
                                                            LAYOUT-999 (PR)
```

---

## Execution Plan

### Phase 1: Basis-Komponenten (Parallel)
Stories: LAYOUT-001, LAYOUT-002, LAYOUT-003, LAYOUT-004, LAYOUT-005, LAYOUT-006

All can run in parallel - no dependencies between them.

### Phase 2: Integration
Stories: LAYOUT-007

Depends on LAYOUT-003 (TuiKeyHints) for integration.

### Phase 3: Validation
Stories: LAYOUT-008, LAYOUT-009

Both depend on all component stories being complete.

### Phase 4: Finalization (Sequential)
Stories: LAYOUT-997, LAYOUT-998, LAYOUT-999

System stories must run sequentially:
1. LAYOUT-997: Code Review by Opus
2. LAYOUT-998: Integration Validation
3. LAYOUT-999: Finalize and Create PR

---

## Story Files

| ID | File |
|----|------|
| LAYOUT-001 | [story-001-tui-navigation.md](stories/story-001-tui-navigation.md) |
| LAYOUT-002 | [story-002-tui-modal.md](stories/story-002-tui-modal.md) |
| LAYOUT-003 | [story-003-tui-keyhints.md](stories/story-003-tui-keyhints.md) |
| LAYOUT-004 | [story-004-tui-table.md](stories/story-004-tui-table.md) |
| LAYOUT-005 | [story-005-tui-tree.md](stories/story-005-tui-tree.md) |
| LAYOUT-006 | [story-006-tui-tabs.md](stories/story-006-tui-tabs.md) |
| LAYOUT-007 | [story-007-tui-statusbar.md](stories/story-007-tui-statusbar.md) |
| LAYOUT-008 | [story-008-layout-demo.md](stories/story-008-layout-demo.md) |
| LAYOUT-009 | [story-009-visual-tests.md](stories/story-009-visual-tests.md) |
| LAYOUT-997 | [story-997-code-review.md](stories/story-997-code-review.md) |
| LAYOUT-998 | [story-998-integration-validation.md](stories/story-998-integration-validation.md) |
| LAYOUT-999 | [story-999-finalize-pr.md](stories/story-999-finalize-pr.md) |

---

## Blocked Stories

Currently no stories are blocked by incomplete prerequisites.

---

## Total Estimated Effort

| Size | Count | Story Points |
|------|-------|--------------|
| XS | 3 | 3 |
| S | 7 | 14 |
| M | 2 | 6 |
| **Total** | **12** | **23 SP** |

---

*Generated with Agent OS /create-spec v3.0*
