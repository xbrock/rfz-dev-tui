# Technical Specification: Boilerplate Cleanup

> Spec: BUGFIX-BOILERPLATE
> Created: 2026-02-02
> Author: Technical Architecture Agent

---

## Root Cause Analysis

### How Boilerplate Got Copied

**Timeline of Events:**

1. **Commit `d2729b3`** (Initial Commit) - Repository initialized
2. **Commit `802ecce`** (feat(spec): Add specification for core-components)
   - Added Core Components specification files
   - **ACCIDENTALLY** copied boilerplate files to project root
   - Commit message mentions: "Also includes: Initial Go project structure (cmd, internal, configs)"
   - This was the mistake - boilerplate was meant as reference only

**Evidence from Git Log:**
```
git show --stat 802ecce
```
Shows 55 files changed, including:
- `Makefile` (boilerplate)
- `cmd/rfz-cli/main.go` (boilerplate)
- `configs/components.yaml` (boilerplate)
- `internal/app/*` (boilerplate)
- `internal/domain/*` (boilerplate)
- `internal/infra/*` (boilerplate)
- `internal/service/*` (boilerplate)
- `internal/ui/components/styles.go` (boilerplate - later modified)
- `internal/ui/screens/*` (boilerplate)
- `internal/ui/modals/*` (boilerplate)

3. **Commits `2fbd7ed` to `a69837c`** - Core Components implementation
   - CORE-001: Added `helpers.go` to internal/ui/components/
   - CORE-002: Added `box.go` and `box_test.go`
   - CORE-003: Added `divider.go` and `divider_test.go`
   - CORE-004: Added `button.go` and `button_test.go`
   - CORE-005: Verified existing `status.go` (from boilerplate)
   - CORE-006: Added test infrastructure and golden files
   - CORE-007: Added `demo/gallery.go` and tests

### Why Compile Errors Occur

The boilerplate files contain bugs that were never fixed because they were just templates:

**Error 1:** `internal/ui/screens/logs/logs.go:193`
```go
scrollInfo := components.StyleBodyMuted.Render(
    m.viewport.ScrollPercent(),  // float64, but Render expects string
)
```

**Error 2:** `internal/ui/modals/buildconfig/buildconfig.go:233`
```go
components.StyleKeyboard.Render(m.config.Parallelism)  // int, but Render expects string
```

These are simply bugs in the boilerplate code - `lipgloss.Style.Render()` only accepts string arguments.

---

## File Verification Results

### Files Verified as Boilerplate (REMOVE)

All files below are **identical** to their counterparts in `agent-os/product/boilerplate/`:

| File Path | Verification | Commit Added |
|-----------|--------------|--------------|
| `Makefile` | diff = identical | 802ecce |
| `cmd/rfz-cli/main.go` | diff = identical | 802ecce |
| `configs/components.yaml` | diff = identical | 802ecce |
| `internal/app/app.go` | diff = identical | 802ecce |
| `internal/app/keymap.go` | diff = identical | 802ecce |
| `internal/app/messages.go` | diff = identical | 802ecce |
| `internal/domain/buildconfig.go` | diff = identical | 802ecce |
| `internal/domain/component.go` | diff = identical | 802ecce |
| `internal/domain/logentry.go` | diff = identical | 802ecce |
| `internal/infra/adapters/adapters.go` | diff = identical | 802ecce |
| `internal/infra/adapters/git_mock.go` | diff = identical | 802ecce |
| `internal/infra/adapters/git_real.go` | diff = identical | 802ecce |
| `internal/infra/adapters/maven_mock.go` | diff = identical | 802ecce |
| `internal/infra/adapters/maven_real.go` | diff = identical | 802ecce |
| `internal/infra/ports/filesystem.go` | diff = identical | 802ecce |
| `internal/infra/ports/git.go` | diff = identical | 802ecce |
| `internal/infra/ports/maven.go` | diff = identical | 802ecce |
| `internal/service/build.go` | diff = identical | 802ecce |
| `internal/service/config.go` | diff = identical | 802ecce |
| `internal/service/scan.go` | diff = identical | 802ecce |
| `internal/ui/components/navitem.go` | diff = identical | 802ecce |
| `internal/ui/components/statusbar.go` | diff = identical | 802ecce |
| `internal/ui/modals/buildconfig/buildconfig.go` | diff = identical | 802ecce |
| `internal/ui/modals/confirm/confirm.go` | diff = identical | 802ecce |
| `internal/ui/screens/build/build.go` | diff = identical | 802ecce |
| `internal/ui/screens/config/config.go` | diff = identical | 802ecce |
| `internal/ui/screens/discover/discover.go` | diff = identical | 802ecce |
| `internal/ui/screens/logs/logs.go` | diff = identical | 802ecce |
| `internal/ui/screens/welcome/welcome.go` | diff = identical | 802ecce |

**Total: 29 files to remove**

### Files Verified as Implementation (KEEP)

| File Path | Verification | Notes |
|-----------|--------------|-------|
| `go.mod` | DIFFERENT | Has additional deps (golden, reflow) |
| `go.sum` | DIFFERENT | Updated lock file |
| `internal/ui/components/styles.go` | IDENTICAL to boilerplate | Used by implementation, KEEP |
| `internal/ui/components/status.go` | IDENTICAL to boilerplate | Used by implementation, KEEP |
| `internal/ui/components/helpers.go` | NOT in boilerplate | CORE-001 implementation |
| `internal/ui/components/box.go` | NOT in boilerplate | CORE-002 implementation |
| `internal/ui/components/divider.go` | NOT in boilerplate | CORE-003 implementation |
| `internal/ui/components/button.go` | NOT in boilerplate | CORE-004 implementation |
| `internal/ui/components/box_test.go` | NOT in boilerplate | CORE-006 tests |
| `internal/ui/components/divider_test.go` | NOT in boilerplate | CORE-006 tests |
| `internal/ui/components/button_test.go` | NOT in boilerplate | CORE-006 tests |
| `internal/ui/components/status_test.go` | NOT in boilerplate | CORE-006 tests |
| `internal/ui/components/testdata/*` | NOT in boilerplate | 36 golden files |
| `internal/ui/components/demo/gallery.go` | NOT in boilerplate | CORE-007 |
| `internal/ui/components/demo/gallery_test.go` | NOT in boilerplate | CORE-007 |
| `internal/ui/components/demo/testdata/*` | NOT in boilerplate | Gallery golden files |

---

## Risk Assessment

### Low Risk (Safe to Delete)

| Item | Reason |
|------|--------|
| `cmd/` | Not used, no imports from implementation |
| `configs/` | Not used, no imports from implementation |
| `internal/app/` | Not used, no imports from implementation |
| `internal/domain/` | Not used, no imports from implementation |
| `internal/infra/` | Not used, no imports from implementation |
| `internal/service/` | Not used, no imports from implementation |
| `internal/ui/screens/` | Not used, causes compile errors |
| `internal/ui/modals/` | Not used, causes compile errors |
| `internal/ui/components/navitem.go` | Not imported anywhere |
| `internal/ui/components/statusbar.go` | Not imported anywhere |
| `Makefile` | Not used (no build automation configured) |

### Verification Commands

Run these to verify no implementation depends on boilerplate:

```bash
# Check for imports of deleted packages
grep -r "rfz-cli/internal/app" internal/ui/components/
grep -r "rfz-cli/internal/domain" internal/ui/components/
grep -r "rfz-cli/internal/service" internal/ui/components/
grep -r "rfz-cli/internal/infra" internal/ui/components/

# Verify navitem.go and statusbar.go are not imported
grep -r "navitem" internal/ui/components/*.go
grep -r "statusbar" internal/ui/components/*.go
```

Expected: All grep commands return empty (no matches).

---

## Recommended Removal Strategy

### Phase 1: Remove Entire Directories (Safest)

```bash
# These directories are entirely boilerplate with no implementation files
rm -rf cmd/
rm -rf configs/
rm -rf internal/app/
rm -rf internal/domain/
rm -rf internal/infra/
rm -rf internal/service/
rm -rf internal/ui/screens/
rm -rf internal/ui/modals/
```

### Phase 2: Remove Individual Boilerplate Files

```bash
# Root level
rm Makefile

# Components directory - remove only boilerplate files
rm internal/ui/components/navitem.go
rm internal/ui/components/statusbar.go
```

### Phase 3: Cleanup Empty Directories

```bash
# internal/ui/ will still have components/ so no cleanup needed
# internal/ will still have ui/ so no cleanup needed
```

### Phase 4: Verify

```bash
# Build should succeed
go build ./...

# Tests should pass
go test ./internal/ui/components/...

# Lint should pass
golangci-lint run ./internal/ui/components/...
```

---

## Story Point Estimates

| Story | Complexity | Estimate |
|-------|------------|----------|
| BUGFIX-001 (Identify) | Already Complete | 0 SP |
| BUGFIX-002 (Remove) | Simple file deletion | 1 SP |
| BUGFIX-003 (Safeguard) | Documentation + optional CI | 2 SP |

**Total: 3 Story Points**

---

## Technical Notes for Implementation

### Files That Will Remain After Cleanup

```
rfz-tui/
├── .claude/
├── .git/
├── .gitignore
├── agent-os/
├── CLAUDE.md
├── go.mod
├── go.sum
├── internal/
│   └── ui/
│       └── components/
│           ├── box.go
│           ├── box_test.go
│           ├── button.go
│           ├── button_test.go
│           ├── demo/
│           │   ├── gallery.go
│           │   ├── gallery_test.go
│           │   └── testdata/
│           ├── divider.go
│           ├── divider_test.go
│           ├── helpers.go
│           ├── status.go
│           ├── status_test.go
│           ├── styles.go
│           └── testdata/
└── references/
```

### Post-Cleanup Module Path

The `go.mod` declares module path as `rfz-cli`. After cleanup:
- Only `internal/ui/components/` package exists
- All imports of `rfz-cli/internal/ui/components` will work
- No other internal packages will exist (that's fine for now)

---

## Appendix: Command Reference

### Full Cleanup Script

```bash
#!/bin/bash
# Boilerplate cleanup script for rfz-tui

set -e

echo "=== Removing boilerplate directories ==="
rm -rf cmd/
rm -rf configs/
rm -rf internal/app/
rm -rf internal/domain/
rm -rf internal/infra/
rm -rf internal/service/
rm -rf internal/ui/screens/
rm -rf internal/ui/modals/

echo "=== Removing boilerplate files ==="
rm -f Makefile
rm -f internal/ui/components/navitem.go
rm -f internal/ui/components/statusbar.go

echo "=== Verifying build ==="
go build ./...

echo "=== Running tests ==="
go test ./internal/ui/components/...

echo "=== Cleanup complete ==="
```

### Rollback (if needed)

```bash
# Reset to pre-cleanup state
git checkout HEAD -- .
```
