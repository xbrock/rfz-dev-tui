# Story Index

> Spec: Core Components
> Created: 2026-02-02
> Last Updated: 2026-02-02 (System Stories Added)

## Overview

This document provides an overview of all user stories for the Core Components specification.

**Total Stories**: 10 (7 regular + 3 system)
**Estimated Effort**: 20 SP

---

## Story Summary

### Regular Stories

| Story ID | Title | Type | Priority | Dependencies | Status | Points |
|----------|-------|------|----------|--------------|--------|--------|
| CORE-001 | Styles Package | Backend | Critical | None | Ready | 3 |
| CORE-002 | TuiBox Component | Backend | Critical | CORE-001 | Ready | 2 |
| CORE-003 | TuiDivider Component | Backend | High | CORE-001 | Ready | 1 |
| CORE-004 | TuiButton Component | Backend | Critical | CORE-001 | Ready | 2 |
| CORE-005 | TuiStatus Component | Backend | Critical | CORE-001 | Ready | 2 |
| CORE-006 | teatest Infrastructure | Test | High | CORE-001 to 005 | Ready | 3 |
| CORE-007 | Component Gallery | Backend | Medium | CORE-006 | Ready | 3 |

### System Stories (execute after all regular stories)

| Story ID | Title | Type | Priority | Dependencies | Status | Points |
|----------|-------|------|----------|--------------|--------|--------|
| CORE-997 | Code Review | System | Critical | All regular stories | Ready | 2 |
| CORE-998 | Integration Validation | System | Critical | CORE-997 | Ready | 1 |
| CORE-999 | Finalize PR | System | Critical | CORE-998 | Ready | 1 |

---

## Dependency Graph

```
CORE-001 (Styles Package) [No dependencies]
    │
    ├───────────────────────────────────────┐
    │                                       │
    ▼                                       ▼
CORE-002 (TuiBox)                    CORE-003 (TuiDivider)
    │                                       │
    │                                       │
    ▼                                       ▼
CORE-004 (TuiButton)                 CORE-005 (TuiStatus)
    │                                       │
    └───────────────────┬───────────────────┘
                        │
                        ▼
              CORE-006 (teatest Infrastructure)
                        │
                        ▼
              CORE-007 (Component Gallery)
                        │
                        ▼
        ════════════════════════════════════
                  SYSTEM STORIES
        ════════════════════════════════════
                        │
                        ▼
              CORE-997 (Code Review)
                        │
                        ▼
              CORE-998 (Integration Validation)
                        │
                        ▼
              CORE-999 (Finalize PR)
```

---

## Execution Plan

### Phase A: Foundation (Sequential)
1. **CORE-001**: Styles Package (must complete first)

### Phase B: Core Components (Parallel)
After CORE-001 completes, these can run in parallel:
- **CORE-002**: TuiBox Component
- **CORE-003**: TuiDivider Component
- **CORE-004**: TuiButton Component
- **CORE-005**: TuiStatus Component

### Phase C: Testing (Sequential)
6. **CORE-006**: teatest Infrastructure (depends on all components)

### Phase D: Demo (Sequential)
7. **CORE-007**: Component Gallery (depends on testing)

### Phase E: System Stories (Sequential, after all regular stories)
8. **CORE-997**: Code Review (Opus reviews entire diff)
9. **CORE-998**: Integration Validation (runs integration tests)
10. **CORE-999**: Finalize PR (creates PR, cleanup)

---

## Story Files

Individual story files are located in the `stories/` subdirectory:

### Regular Stories
- `stories/story-001-styles-package.md`
- `stories/story-002-tuibox-component.md`
- `stories/story-003-tuidivider-component.md`
- `stories/story-004-tuibutton-component.md`
- `stories/story-005-tuistatus-component.md`
- `stories/story-006-teatest-infrastructure.md`
- `stories/story-007-component-gallery.md`

### System Stories
- `stories/story-997-code-review.md`
- `stories/story-998-integration-validation.md`
- `stories/story-999-finalize-pr.md`

---

## Blocked Stories

None - All stories have completed DoR (Architect refinement completed 2026-02-02).

---

## Effort Breakdown

| Phase | Stories | Total Points | Parallel Possible |
|-------|---------|--------------|-------------------|
| Foundation | CORE-001 | 3 SP | No |
| Components | CORE-002 to 005 | 7 SP | Yes (4 parallel) |
| Testing | CORE-006 | 3 SP | No |
| Demo | CORE-007 | 3 SP | No |
| System | CORE-997 to 999 | 4 SP | No |

**Total**: 20 SP
