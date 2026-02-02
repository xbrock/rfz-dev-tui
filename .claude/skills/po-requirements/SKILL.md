---
description: PO Requirements - Guide for creating user stories and specifications
globs:
  - "agent-os/specs/**/*.md"
  - "agent-os/product/**/*.md"
alwaysApply: false
version: 1.0.0
---

# PO Requirements Skill

Guidance for creating and refining user stories, specifications, and requirements for the RFZ Developer CLI project.

## Quick Reference

### User Story Format

```markdown
## Story: [Feature Name]

**As a** RFZ developer,
**I want to** [capability],
**So that** [benefit].

### Acceptance Criteria
- [ ] AC1: [Specific, testable criterion]
- [ ] AC2: [Another criterion]

### Edge Cases
- EC1: [What happens when...]
- EC2: [What if...]
```

### Story Types for RFZ-CLI

| Type | Description | Example |
|------|-------------|---------|
| Screen | New TUI screen | "Build Configuration Modal" |
| Component | Reusable UI element | "TuiStatus Badge" |
| Service | Business logic | "Maven Execution Service" |
| Integration | External system | "Git Status Detection" |

## Project Context

### Product Brief Reference

- **Product**: RFZ Developer CLI
- **Users**: RFZ developers at Deutsche Bahn
- **Problem**: Complex Maven build workflows
- **Solution**: TUI for streamlined component building

### Key Documents

| Document | Purpose |
|----------|---------|
| `agent-os/product/product-brief.md` | Product vision and goals |
| `agent-os/product/tech-stack.md` | Technology decisions |
| `agent-os/product/design-system.md` | Visual design tokens |
| `agent-os/product/ux-patterns.md` | Interaction patterns |
| `agent-os/product/architecture-decision.md` | Architecture patterns |

## Story Quality Checklist

### Good Story Has

- [ ] Clear user benefit (not just technical task)
- [ ] Testable acceptance criteria
- [ ] Defined scope (not too big, not trivial)
- [ ] Edge cases identified
- [ ] Technical feasibility confirmed

### Acceptance Criteria Guidelines

**Good AC:**
- "User can navigate between 5 screens using number keys 1-5"
- "Build progress shows percentage with spinner"
- "Error state displays message in red with `x` symbol"

**Bad AC:**
- "Works correctly" (not testable)
- "Looks good" (not specific)
- "Fast" (not measurable)

## TUI-Specific Requirements

### Screen Stories Should Include

1. **Layout specification** (which panels, proportions)
2. **Keyboard shortcuts** (primary actions)
3. **Focus behavior** (Tab order, panel focus)
4. **Empty state** (what to show when no data)
5. **Error state** (how errors display)

### Component Stories Should Include

1. **Variants** (normal, focused, active, disabled)
2. **Lip Gloss styling** (colors, borders, padding)
3. **Accessibility** (symbols + colors, not color-only)

## Reference Screenshots

The `references/prototype-screenshots/` directory contains 77 UI screenshots:

| Range | Content |
|-------|---------|
| 01-09 | Welcome screen states |
| 10-19 | Build screen states |
| 20-29 | Configuration screen states |
| 30-39 | Log viewer states |
| 40-49 | Discover screen states |
| 50-59 | Modal states |
| 95-97 | Design system reference |

Use these to verify story acceptance against visual design.

## Story Dependencies

When writing stories, consider:

1. **Does this need a new port interface?** (for external systems)
2. **Does this need a new service?** (for business logic)
3. **Does this need a new domain type?** (for data modeling)
4. **Does this use existing Bubbles components?** (list, table, viewport)
5. **Does this need a custom component?** (TuiStatus, TuiNavItem)
