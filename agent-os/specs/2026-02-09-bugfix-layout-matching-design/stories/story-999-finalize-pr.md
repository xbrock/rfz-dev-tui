# Finalize PR

> Story ID: LAYOUT-999
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: -
**Type**: System/Finalization
**Estimated Effort**: -
**Dependencies**: LAYOUT-998

---

## Purpose

Finalize the pull request: create test scenarios documentation, generate user TODOs, create PR, and clean up worktree. Replaces Phase 5 of execute-tasks.

## Tasks

1. **Test Scenarios** - Document manual visual test scenarios for reviewer
2. **User TODOs** - List any remaining items for the user (e.g., golden file updates)
3. **Create PR** - Push branch and create pull request with summary
4. **Cleanup** - Clean up any temporary files or worktree artifacts

## PR Template

```markdown
## Summary
Fix TUI layout and styling to match approved design prototype across all implemented screens.

## Changes
- Title bar: Red accent line on top, correct title hierarchy
- Navigation: Active/select states, shortcut alignment, tree hints, content-based height
- Status bar: 3-badge system, pipe-separated hints, gray bg
- Welcome screen: Logo colors, braille line, version badges
- Build components: Circle checkboxes, category alignment, action button styles
- Config modal: Shortcut label colors, hint separators
- Build execution: Tree icons, braille progress bars, badge cleanup
- General: Fixed border overflow on all screens
```
