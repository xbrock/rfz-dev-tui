# Definition of Done (DoD)

> RFZ Developer CLI
> Last Updated: 2026-02-02

## Overview

A task/story is considered **done** when all of the following criteria are met.

---

## Code Quality

### Mandatory Checks

- [ ] **Lint passes**: `golangci-lint run ./...` reports no errors
- [ ] **Tests pass**: `go test ./...` all tests pass
- [ ] **Build succeeds**: `go build ./cmd/rfz-cli/` compiles without errors
- [ ] **No regressions**: Golden file tests match (or intentionally updated)

### Code Standards

- [ ] Follows Go idioms and conventions
- [ ] Uses Lip Gloss for ALL styling (no raw ANSI codes)
- [ ] Uses Bubbles components when available (no reinventing)
- [ ] Follows layered architecture (Presentation -> Application -> Domain -> Infrastructure)
- [ ] Proper error handling on all error returns

---

## Functionality

- [ ] Acceptance criteria from story are met
- [ ] Edge cases identified and handled
- [ ] Error states display user-friendly messages
- [ ] Empty states show helpful guidance

---

## Visual Quality (for UI changes)

- [ ] Matches design system colors and borders
- [ ] Focus indicators visible at all times
- [ ] Status uses symbols AND colors (accessibility)
- [ ] Tested at canonical size (120x40)
- [ ] Golden file updated if UI intentionally changed

---

## Architecture

- [ ] Code placed in correct layer (ui/service/domain/infra)
- [ ] External dependencies accessed through ports (interfaces)
- [ ] Services receive dependencies via injection
- [ ] No business logic in presentation layer

---

## Testing

- [ ] Unit tests for new functionality
- [ ] Golden file tests for new UI states
- [ ] Mock adapters used (no real Maven/Git in tests)
- [ ] Test coverage maintained

---

## Documentation (if applicable)

- [ ] Code is self-documenting (clear naming)
- [ ] Complex logic has comments explaining "why"
- [ ] Public APIs have doc comments
- [ ] Skills updated if new patterns discovered

---

## Review Checklist

Before marking complete:

1. Run `golangci-lint run ./...` - all clean?
2. Run `go test ./...` - all pass?
3. Run `go build ./cmd/rfz-cli/` - builds?
4. Review own code for obvious issues
5. Check golden file diffs are intentional
