# Bug Spec Lite: Boilerplate in Project Root

> ID: BUGFIX-BOILERPLATE
> Severity: Critical | Priority: Urgent
> Created: 2026-02-02

## Problem

Boilerplate code from `agent-os/product/boilerplate/` was accidentally copied to project root, causing compile errors.

## Symptoms

```
go build ./...
# rfz-cli/internal/ui/modals/buildconfig
buildconfig.go:233:49: cannot use m.config.Parallelism (variable of type int) as string value
# rfz-cli/internal/ui/screens/logs
logs.go:193:3: cannot use m.viewport.ScrollPercent() (value of type float64) as string value
```

## Fix Summary

**Remove 29 boilerplate files from root:**
- `Makefile`, `cmd/`, `configs/`
- `internal/app/`, `internal/domain/`, `internal/infra/`, `internal/service/`
- `internal/ui/components/navitem.go`, `internal/ui/components/statusbar.go`
- `internal/ui/modals/`, `internal/ui/screens/`

**Keep actual implementation:**
- `go.mod`, `go.sum`
- `internal/ui/components/styles.go`, `helpers.go`, `box.go`, `divider.go`, `button.go`, `status.go`
- All `*_test.go` and `testdata/` files
- `internal/ui/components/demo/`

## Stories

| ID | Title | Priority |
|----|-------|----------|
| BUGFIX-001 | Identify boilerplate files | Must Have |
| BUGFIX-002 | Remove boilerplate from root | Must Have |
| BUGFIX-003 | Add regression safeguard | Should Have |

## Success Criteria

- `go build ./...` succeeds
- `go test ./internal/ui/components/...` passes
- Boilerplate in `agent-os/product/boilerplate/` unchanged
