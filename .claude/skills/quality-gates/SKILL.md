---
description: Quality Gates - Mandatory checks before completing any task
globs:
  - "**/*.go"
  - "**/*.md"
alwaysApply: true
version: 1.0.0
---

# Quality Gates Skill

Mandatory quality checks that apply to ALL code changes in the RFZ Developer CLI project.

## Quick Reference

### Before Completing Any Task

| Check | Command | Required |
|-------|---------|----------|
| Lint | `golangci-lint run ./...` | YES |
| Tests | `go test ./...` | YES |
| Build | `go build ./cmd/rfz-cli/` | YES |

### Quality Criteria

1. **No lint errors** - All golangci-lint issues must be resolved
2. **All tests pass** - Unit tests and golden file tests
3. **Clean build** - Code compiles without errors
4. **No regressions** - Visual regression tests at 120x40 canonical size

## Mandatory Workflow

### After ANY Code Change

```bash
# 1. Run linter (MUST pass)
golangci-lint run ./...

# 2. Run tests (MUST pass)
go test ./...

# 3. Verify build (MUST succeed)
go build ./cmd/rfz-cli/
```

### If Lint Fails

1. Fix ALL reported issues before continuing
2. Common fixes:
   - Unused variables: Remove or use `_`
   - Import order: Let `goimports` fix
   - Style issues: Follow Go idioms

### If Tests Fail

1. Analyze failure message
2. Fix the issue
3. Re-run tests
4. Update golden files ONLY if change is intentional

## Code Quality Rules

### Go Patterns

- Follow standard Go idioms
- Use `gofmt` formatting (automatic with golangci-lint)
- No unused code (imports, variables, functions)
- Error handling required for all error returns
- Context propagation for long-running operations

### Charm.land Rules

- **ALWAYS** use Lip Gloss for styling (never raw ANSI)
- **ALWAYS** use Bubbles components when available
- **NEVER** draw borders manually (use `lipgloss.Border`)
- **NEVER** use manual string padding (use `lipgloss.Padding`)

### Architecture Rules

- Presentation layer (`internal/ui/`) handles rendering only
- Services (`internal/service/`) handle business logic
- Domain (`internal/domain/`) contains pure types
- Ports (`internal/infra/ports/`) define interfaces
- Adapters (`internal/infra/adapters/`) implement ports

## Definition of Done

A task is complete when:

- [ ] All lint errors resolved
- [ ] All tests pass
- [ ] Build succeeds
- [ ] No visual regressions (teatest golden files match)
- [ ] Code follows project patterns (Charm.land first)
- [ ] No unnecessary complexity added

## Common Quality Issues

| Issue | Fix |
|-------|-----|
| Unused import | Remove the import |
| Unused variable | Use `_` or remove |
| Missing error check | Handle or explicitly ignore with `_ = err` |
| Lint timeout | Run on specific package: `golangci-lint run ./internal/ui/...` |
| Golden file mismatch | Review change, update if intentional |
