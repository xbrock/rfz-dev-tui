# Definition of Ready (DoR)

> RFZ Developer CLI
> Last Updated: 2026-02-02

## Overview

A story/task is considered **ready** for implementation when all of the following criteria are met.

---

## Story Completeness

### Required Elements

- [ ] **User story format**: "As a [user], I want to [capability], so that [benefit]"
- [ ] **Acceptance criteria**: Clear, testable criteria listed
- [ ] **Edge cases identified**: What happens in unusual scenarios
- [ ] **Scope is clear**: Not too big, not too vague

### For UI Stories

- [ ] **Layout defined**: Which panels, proportions specified
- [ ] **Keyboard shortcuts**: Primary actions documented
- [ ] **Focus behavior**: Tab order, panel focus explained
- [ ] **Empty state**: What to show when no data
- [ ] **Error state**: How errors display
- [ ] **Reference screenshot**: If available from prototype

### For Service Stories

- [ ] **Inputs defined**: What data/parameters needed
- [ ] **Outputs defined**: What result returned
- [ ] **Error cases**: What can go wrong
- [ ] **Dependencies**: Which ports needed

---

## Technical Clarity

- [ ] **Feasibility confirmed**: Can be done with current architecture
- [ ] **Dependencies identified**: What must exist first
- [ ] **Affected layers clear**: Which layer(s) touched
- [ ] **Components identified**: Which Bubbles/custom components

---

## Context Available

- [ ] **Design reference accessible**: Screenshots, design-system.md
- [ ] **Architecture docs available**: architecture-decision.md
- [ ] **Related code identified**: Which files to read/modify

---

## Questions Resolved

- [ ] No blocking questions remaining
- [ ] Ambiguities clarified with PO
- [ ] Technical approach agreed

---

## Estimation

- [ ] Story is estimable (not "unknown effort")
- [ ] Story fits within a reasonable iteration

---

## Ready Checklist

Before starting implementation:

1. Read the full story including acceptance criteria
2. Review related prototype screenshots
3. Identify which files need modification
4. Understand which Bubbles components to use
5. Confirm no blocking questions
