# Bug Specification: Boilerplate Code Accidentally Copied to Project Root

> Bug ID: BUGFIX-BOILERPLATE
> Created: 2026-02-02
> Severity: Critical
> Priority: Urgent
> Status: Open

---

## Problem Statement

Boilerplate code from `agent-os/product/boilerplate/` was accidentally copied to the project root (`./`) during the implementation of specification `2026-02-02-core-components`. This results in compile errors due to type mismatches between outdated boilerplate code and the newly developed styles package.

The issue was previously fixed in a worktree but returned after merge, indicating the fix was not properly applied to the main branch.

---

## Reproduction Steps

1. Clone the repository
2. Run `go build ./...`
3. Observe compile errors:
   ```
   # rfz-cli/internal/ui/modals/buildconfig
   internal/ui/modals/buildconfig/buildconfig.go:233:49: cannot use m.config.Parallelism (variable of type int) as string value in argument to components.StyleKeyboard.Render

   # rfz-cli/internal/ui/screens/logs
   internal/ui/screens/logs/logs.go:193:3: cannot use m.viewport.ScrollPercent() (value of type float64) as string value in argument to components.StyleBodyMuted.Render
   ```

---

## Expected Behavior

- Project root should contain only the actual implementation code
- Boilerplate should remain only in `agent-os/product/boilerplate/` as reference material
- `go build ./...` should complete successfully
- `go test ./...` should pass

---

## Actual Behavior

- Boilerplate files duplicated in project root
- Outdated boilerplate code references styles that have since been updated
- Compile errors block all development work
- Type mismatches between boilerplate's expected API and actual styles package

---

## Root Cause Analysis

**Commit History Analysis:**

1. **Commit `802ecce`** (feat(spec): Add specification for core-components) - Introduced boilerplate files to root
2. **Commits `2fbd7ed` to `a69837c`** - Developed Core Components (CORE-001 to CORE-007) in `internal/ui/components/`
3. The styles package evolved during Core Components implementation, but boilerplate files were not updated

**The Problem:**
The boilerplate was designed as a starting point template, but was accidentally committed to the project root alongside the actual implementation. As the actual implementation evolved (particularly styles.go), the boilerplate files became incompatible.

---

## Impact Assessment

| Impact Area | Severity | Description |
|-------------|----------|-------------|
| Build System | Critical | Project cannot compile |
| Development | Blocked | No new features can be developed |
| CI/CD | Blocked | All builds fail |
| Testing | Blocked | Cannot run tests |

---

## Files to Remove (Boilerplate in Root)

The following files in project root are duplicates of boilerplate and should be removed:

### Root Level Files
- `Makefile` - Identical to boilerplate
- `configs/components.yaml` - Boilerplate config

### cmd/ Directory
- `cmd/rfz-cli/main.go` - Boilerplate entry point

### internal/app/ Directory
- `internal/app/app.go` - Boilerplate app
- `internal/app/keymap.go` - Boilerplate keymaps
- `internal/app/messages.go` - Boilerplate messages

### internal/domain/ Directory
- `internal/domain/buildconfig.go` - Boilerplate domain model
- `internal/domain/component.go` - Boilerplate domain model
- `internal/domain/logentry.go` - Boilerplate domain model

### internal/infra/ Directory
- `internal/infra/adapters/adapters.go`
- `internal/infra/adapters/git_mock.go`
- `internal/infra/adapters/git_real.go`
- `internal/infra/adapters/maven_mock.go`
- `internal/infra/adapters/maven_real.go`
- `internal/infra/ports/filesystem.go`
- `internal/infra/ports/git.go`
- `internal/infra/ports/maven.go`

### internal/service/ Directory
- `internal/service/build.go`
- `internal/service/config.go`
- `internal/service/scan.go`

### internal/ui/components/ Directory (Boilerplate Only)
- `internal/ui/components/navitem.go` - Boilerplate (not part of Core Components spec)
- `internal/ui/components/statusbar.go` - Boilerplate (not part of Core Components spec)

### internal/ui/modals/ Directory
- `internal/ui/modals/buildconfig/buildconfig.go` - Boilerplate (causes compile error)
- `internal/ui/modals/confirm/confirm.go` - Boilerplate

### internal/ui/screens/ Directory
- `internal/ui/screens/build/build.go`
- `internal/ui/screens/config/config.go`
- `internal/ui/screens/discover/discover.go`
- `internal/ui/screens/logs/logs.go` - Boilerplate (causes compile error)
- `internal/ui/screens/welcome/welcome.go`

---

## Files to KEEP (Actual Implementation from Core Components Spec)

The following files were developed as part of `2026-02-02-core-components` and must NOT be removed:

### Core Files
- `go.mod` - Updated with actual dependencies (different from boilerplate)
- `go.sum` - Dependency lockfile

### internal/ui/components/ (Actual Implementation)
- `internal/ui/components/styles.go` - CORE-001: Styles Package (modified from boilerplate)
- `internal/ui/components/helpers.go` - CORE-001: Truncate helper
- `internal/ui/components/box.go` - CORE-002: TuiBox Component
- `internal/ui/components/divider.go` - CORE-003: TuiDivider Component
- `internal/ui/components/button.go` - CORE-004: TuiButton Component
- `internal/ui/components/status.go` - CORE-005: TuiStatus Component (modified from boilerplate)

### Test Files
- `internal/ui/components/box_test.go` - CORE-006
- `internal/ui/components/divider_test.go` - CORE-006
- `internal/ui/components/button_test.go` - CORE-006
- `internal/ui/components/status_test.go` - CORE-006
- `internal/ui/components/testdata/*` - Golden files

### Demo
- `internal/ui/components/demo/gallery.go` - CORE-007
- `internal/ui/components/demo/gallery_test.go` - CORE-007
- `internal/ui/components/demo/testdata/*` - Gallery golden files

---

## Verification Criteria

After fix is applied:

1. **Build succeeds**: `go build ./...` completes with no errors
2. **Tests pass**: `go test ./internal/ui/components/...` passes
3. **Lint passes**: `golangci-lint run ./internal/ui/components/...` passes
4. **Boilerplate intact**: Files in `agent-os/product/boilerplate/` remain unchanged
5. **Directory structure preserved**: Empty directories can be kept for future use

---

## Recommended Fix Approach

1. **Remove boilerplate files from root** - Delete all files listed in "Files to Remove" section
2. **Keep actual implementation** - Preserve all files listed in "Files to KEEP" section
3. **Verify build** - Run `go build ./...` to confirm fix
4. **Run tests** - Run `go test ./internal/ui/components/...`
5. **Add safeguard** - Consider adding `.gitignore` patterns or pre-commit hook to prevent future accidental copies

---

## Related Specifications

- `agent-os/specs/2026-02-02-core-components/` - Original spec where boilerplate was accidentally copied
