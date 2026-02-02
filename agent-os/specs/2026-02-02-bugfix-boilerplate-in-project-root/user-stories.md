# User Stories: Boilerplate Cleanup

> Spec: BUGFIX-BOILERPLATE
> Created: 2026-02-02
> Updated: 2026-02-02 (Technical Analysis Complete)

---

## Story Index

| ID | Title | Priority | SP | Status | DoR |
|----|-------|----------|------|--------|-----|
| BUGFIX-001 | Identify and Document All Boilerplate Files | Must Have | 0 | Complete | Yes |
| BUGFIX-002 | Remove Boilerplate Files from Project Root | Must Have | 1 | Done | Yes |
| BUGFIX-003 | Add Regression Safeguard | Should Have | 2 | Done | Yes |

**Total Story Points: 3**

---

## BUGFIX-001: Identify and Document All Boilerplate Files

### Status: COMPLETE

Technical analysis completed by Technical Architecture Agent. All files verified via `diff` commands against `agent-os/product/boilerplate/`.

### Analysis Results

**Files Verified as Pure Boilerplate (REMOVE) - 29 files:**

| Path | Verification Method | Result |
|------|---------------------|--------|
| `Makefile` | `diff Makefile agent-os/product/boilerplate/Makefile` | Identical |
| `cmd/rfz-cli/main.go` | `diff cmd/... agent-os/product/boilerplate/cmd/...` | Identical |
| `configs/components.yaml` | `diff configs/... agent-os/product/boilerplate/configs/...` | Identical |
| `internal/app/app.go` | diff comparison | Identical |
| `internal/app/keymap.go` | diff comparison | Identical |
| `internal/app/messages.go` | diff comparison | Identical |
| `internal/domain/buildconfig.go` | diff comparison | Identical |
| `internal/domain/component.go` | diff comparison | Identical |
| `internal/domain/logentry.go` | diff comparison | Identical |
| `internal/infra/adapters/adapters.go` | diff comparison | Identical |
| `internal/infra/adapters/git_mock.go` | diff comparison | Identical |
| `internal/infra/adapters/git_real.go` | diff comparison | Identical |
| `internal/infra/adapters/maven_mock.go` | diff comparison | Identical |
| `internal/infra/adapters/maven_real.go` | diff comparison | Identical |
| `internal/infra/ports/filesystem.go` | diff comparison | Identical |
| `internal/infra/ports/git.go` | diff comparison | Identical |
| `internal/infra/ports/maven.go` | diff comparison | Identical |
| `internal/service/build.go` | diff comparison | Identical |
| `internal/service/config.go` | diff comparison | Identical |
| `internal/service/scan.go` | diff comparison | Identical |
| `internal/ui/components/navitem.go` | diff comparison | Identical |
| `internal/ui/components/statusbar.go` | diff comparison | Identical |
| `internal/ui/modals/buildconfig/buildconfig.go` | diff comparison | Identical |
| `internal/ui/modals/confirm/confirm.go` | diff comparison | Identical |
| `internal/ui/screens/build/build.go` | diff comparison | Identical |
| `internal/ui/screens/config/config.go` | diff comparison | Identical |
| `internal/ui/screens/discover/discover.go` | diff comparison | Identical |
| `internal/ui/screens/logs/logs.go` | diff comparison | Identical |
| `internal/ui/screens/welcome/welcome.go` | diff comparison | Identical |

**Files Verified as Actual Implementation (KEEP):**

| Path | Verification | Source |
|------|--------------|--------|
| `go.mod` | Different from boilerplate (has golden, reflow deps) | Core Components |
| `go.sum` | Different from boilerplate | Core Components |
| `internal/ui/components/styles.go` | Identical but KEEP (used by impl) | CORE-001 base |
| `internal/ui/components/status.go` | Identical but KEEP (used by impl) | CORE-005 base |
| `internal/ui/components/helpers.go` | NOT in boilerplate | CORE-001 |
| `internal/ui/components/box.go` | NOT in boilerplate | CORE-002 |
| `internal/ui/components/divider.go` | NOT in boilerplate | CORE-003 |
| `internal/ui/components/button.go` | NOT in boilerplate | CORE-004 |
| `internal/ui/components/*_test.go` | NOT in boilerplate (4 files) | CORE-006 |
| `internal/ui/components/testdata/*` | NOT in boilerplate (36 golden files) | CORE-006 |
| `internal/ui/components/demo/*` | NOT in boilerplate | CORE-007 |

---

## BUGFIX-002: Remove Boilerplate Files from Project Root

### Description

Remove all 29 boilerplate files from the project root while preserving the actual implementation files from the Core Components spec.

### Story Points: 1

### Definition of Ready: YES

- [x] Files to remove verified via diff
- [x] Files to keep verified via diff
- [x] Risk assessment completed (LOW risk)
- [x] Exact commands documented
- [x] Verification criteria defined

### Acceptance Criteria

- [x] Remove all 29 files listed as "Pure Boilerplate"
- [x] Preserve all files listed as "Actual Implementation"
- [x] `go build ./...` succeeds (no compile errors)
- [x] `go test ./internal/ui/components/...` passes
- [x] `golangci-lint run ./internal/ui/components/...` passes
- [x] Changes committed with clear message

### WIE (How to Implement)

**Step 1: Remove entire boilerplate directories**

```bash
cd /Users/lix/xapps/rfz-tui

# Remove entire directories (all contents are boilerplate)
rm -rf cmd/
rm -rf configs/
rm -rf internal/app/
rm -rf internal/domain/
rm -rf internal/infra/
rm -rf internal/service/
rm -rf internal/ui/screens/
rm -rf internal/ui/modals/
```

**Step 2: Remove individual boilerplate files**

```bash
# Remove root-level Makefile
rm Makefile

# Remove boilerplate-only component files
rm internal/ui/components/navitem.go
rm internal/ui/components/statusbar.go
```

**Step 3: Verify build and tests**

```bash
# Build must succeed
go build ./...

# Tests must pass
go test ./internal/ui/components/...

# Lint must pass
golangci-lint run ./internal/ui/components/...
```

**Step 4: Commit changes**

```bash
git add -A
git commit -m "fix: Remove accidentally copied boilerplate files from project root

Boilerplate from agent-os/product/boilerplate/ was accidentally copied
to project root in commit 802ecce. This caused compile errors due to
type mismatches between boilerplate code and the evolved styles package.

Removed:
- Makefile (boilerplate)
- cmd/ directory (boilerplate entry point)
- configs/ directory (boilerplate config)
- internal/app/ (boilerplate app layer)
- internal/domain/ (boilerplate domain models)
- internal/infra/ (boilerplate infrastructure)
- internal/service/ (boilerplate services)
- internal/ui/screens/ (boilerplate screens)
- internal/ui/modals/ (boilerplate modals)
- internal/ui/components/navitem.go (boilerplate)
- internal/ui/components/statusbar.go (boilerplate)

Preserved:
- go.mod, go.sum (actual dependencies)
- internal/ui/components/ core files (CORE-001 to CORE-007)
- All test files and golden test data

Fixes compile errors:
- logs.go:193 - float64 passed to Render()
- buildconfig.go:233 - int passed to Render()"
```

### WO (Where - File Paths)

**Files to DELETE:**

```
/Users/lix/xapps/rfz-tui/Makefile
/Users/lix/xapps/rfz-tui/cmd/rfz-cli/main.go
/Users/lix/xapps/rfz-tui/configs/components.yaml
/Users/lix/xapps/rfz-tui/internal/app/app.go
/Users/lix/xapps/rfz-tui/internal/app/keymap.go
/Users/lix/xapps/rfz-tui/internal/app/messages.go
/Users/lix/xapps/rfz-tui/internal/domain/buildconfig.go
/Users/lix/xapps/rfz-tui/internal/domain/component.go
/Users/lix/xapps/rfz-tui/internal/domain/logentry.go
/Users/lix/xapps/rfz-tui/internal/infra/adapters/adapters.go
/Users/lix/xapps/rfz-tui/internal/infra/adapters/git_mock.go
/Users/lix/xapps/rfz-tui/internal/infra/adapters/git_real.go
/Users/lix/xapps/rfz-tui/internal/infra/adapters/maven_mock.go
/Users/lix/xapps/rfz-tui/internal/infra/adapters/maven_real.go
/Users/lix/xapps/rfz-tui/internal/infra/ports/filesystem.go
/Users/lix/xapps/rfz-tui/internal/infra/ports/git.go
/Users/lix/xapps/rfz-tui/internal/infra/ports/maven.go
/Users/lix/xapps/rfz-tui/internal/service/build.go
/Users/lix/xapps/rfz-tui/internal/service/config.go
/Users/lix/xapps/rfz-tui/internal/service/scan.go
/Users/lix/xapps/rfz-tui/internal/ui/components/navitem.go
/Users/lix/xapps/rfz-tui/internal/ui/components/statusbar.go
/Users/lix/xapps/rfz-tui/internal/ui/modals/buildconfig/buildconfig.go
/Users/lix/xapps/rfz-tui/internal/ui/modals/confirm/confirm.go
/Users/lix/xapps/rfz-tui/internal/ui/screens/build/build.go
/Users/lix/xapps/rfz-tui/internal/ui/screens/config/config.go
/Users/lix/xapps/rfz-tui/internal/ui/screens/discover/discover.go
/Users/lix/xapps/rfz-tui/internal/ui/screens/logs/logs.go
/Users/lix/xapps/rfz-tui/internal/ui/screens/welcome/welcome.go
```

**Files to KEEP:**

```
/Users/lix/xapps/rfz-tui/go.mod
/Users/lix/xapps/rfz-tui/go.sum
/Users/lix/xapps/rfz-tui/internal/ui/components/styles.go
/Users/lix/xapps/rfz-tui/internal/ui/components/status.go
/Users/lix/xapps/rfz-tui/internal/ui/components/helpers.go
/Users/lix/xapps/rfz-tui/internal/ui/components/box.go
/Users/lix/xapps/rfz-tui/internal/ui/components/divider.go
/Users/lix/xapps/rfz-tui/internal/ui/components/button.go
/Users/lix/xapps/rfz-tui/internal/ui/components/box_test.go
/Users/lix/xapps/rfz-tui/internal/ui/components/divider_test.go
/Users/lix/xapps/rfz-tui/internal/ui/components/button_test.go
/Users/lix/xapps/rfz-tui/internal/ui/components/status_test.go
/Users/lix/xapps/rfz-tui/internal/ui/components/testdata/* (36 files)
/Users/lix/xapps/rfz-tui/internal/ui/components/demo/gallery.go
/Users/lix/xapps/rfz-tui/internal/ui/components/demo/gallery_test.go
/Users/lix/xapps/rfz-tui/internal/ui/components/demo/testdata/*
```

### Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Delete implementation files | Very Low | High | Files verified via diff |
| Build still fails | Very Low | Medium | Test after each step |
| Tests fail | Very Low | Medium | Run test suite |

**Overall Risk: LOW** - All deletions verified against boilerplate source.

---

## BUGFIX-003: Add Regression Safeguard

### Description

Add safeguards to prevent future accidental copies of boilerplate code to the project root.

### Story Points: 2

### Definition of Ready: YES

- [x] Options analyzed
- [x] Recommended approach defined
- [x] File locations identified

### Acceptance Criteria

- [x] Document boilerplate directory purpose in CLAUDE.md
- [x] Add README.md to boilerplate directory with clear warnings
- [ ] (Optional) Add CI check to detect boilerplate in root

### WIE (How to Implement)

**Step 1: Update CLAUDE.md**

Add section under "References Directory":

```markdown
### Boilerplate Directory

The `agent-os/product/boilerplate/` directory contains starter code templates:

| Path | Content |
|------|---------|
| `agent-os/product/boilerplate/` | Go project starter code (cmd, internal, configs) |

**WARNING:** This is REFERENCE MATERIAL ONLY. Do NOT copy these files to the project root.
The boilerplate provides a starting point structure but should be adapted and written fresh,
not copied verbatim.
```

**Step 2: Create boilerplate README**

Create `/Users/lix/xapps/rfz-tui/agent-os/product/boilerplate/README.md`:

```markdown
# RFZ CLI Boilerplate

> WARNING: DO NOT COPY THESE FILES TO PROJECT ROOT

This directory contains starter code templates for the RFZ CLI project.

## Purpose

These files serve as REFERENCE MATERIAL showing:
- Recommended project structure
- Example implementations
- API patterns to follow

## Usage

1. READ the boilerplate to understand patterns
2. WRITE your own implementation inspired by these patterns
3. DO NOT copy files directly to project root

## Why Not Copy?

1. Boilerplate may contain placeholder bugs
2. Implementation should evolve independently
3. Copying causes duplicate code and merge conflicts
4. API changes in real code will break copied boilerplate

## Directory Structure

- `cmd/` - Example entry point
- `configs/` - Example configuration
- `internal/` - Example internal packages
```

**Step 3: (Optional) Add CI check**

Create `.github/workflows/check-boilerplate.yml` or add to existing CI:

```yaml
- name: Check for boilerplate in root
  run: |
    # Check if boilerplate directories exist in root
    if [ -d "cmd" ] && [ -d "internal/app" ]; then
      echo "ERROR: Boilerplate directories detected in project root"
      echo "Remove cmd/, internal/app/, internal/domain/, etc."
      exit 1
    fi
```

### WO (Where - File Paths)

**Files to MODIFY:**

```
/Users/lix/xapps/rfz-tui/CLAUDE.md
```

**Files to CREATE:**

```
/Users/lix/xapps/rfz-tui/agent-os/product/boilerplate/README.md
```

**Files to CREATE (Optional):**

```
/Users/lix/xapps/rfz-tui/.github/workflows/check-boilerplate.yml
```

---

## Dependencies

```
BUGFIX-001 (Identify) --> BUGFIX-002 (Remove) --> BUGFIX-003 (Safeguard)
     COMPLETE                 READY                    READY
```

---

## Definition of Done (Global)

- [x] All boilerplate files identified and verified
- [x] All 29 boilerplate files removed from project root
- [x] `go build ./...` succeeds
- [x] `go test ./internal/ui/components/...` passes
- [x] `golangci-lint run ./internal/ui/components/...` passes
- [ ] Changes committed to main branch
- [x] Boilerplate in `agent-os/product/boilerplate/` remains intact
- [x] Documentation updated with safeguard information

---

## Technical Reference

See: `sub-specs/technical-spec.md` for detailed root cause analysis and full verification results.
