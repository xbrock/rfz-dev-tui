# Integration Validation

> Story ID: LAYOUT-998
> Spec: 2026-02-09-bugfix-layout-matching-design
> Created: 2026-02-09
> Last Updated: 2026-02-09

**Priority**: -
**Type**: System/Integration
**Estimated Effort**: -
**Dependencies**: LAYOUT-997

---

## Purpose

Execute integration tests from spec.md to validate the complete feature works end-to-end. Replaces Phase 4.5 of execute-tasks.

## Integration Tests

```bash
# Build passes
cd /Users/lix/xapps/rfz-tui && go build ./...

# Lint passes
cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...

# All test files pass
cd /Users/lix/xapps/rfz-tui && go test ./...
```

## Validation Scenarios

1. App launches without visual errors at 120x40
2. Navigation active/select states render correctly
3. Status bar shows badges at bottom
4. Build flow renders matching prototype
5. No border overflow at any terminal width
