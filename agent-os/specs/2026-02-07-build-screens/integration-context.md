# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| BUILD-001 | Domain model with components, build config, mock provider | `internal/domain/` package created |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
_None yet_

### Services
<!-- New service classes/modules -->
- `internal/domain/mock_provider.go` -> `MockComponentProvider{}` - Returns 13 hardcoded RFZ components

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
_None yet_

### Types / Interfaces
<!-- New type definitions -->
- `internal/domain/component.go` -> `Component{Name, Type}` - RFZ component with name and category
- `internal/domain/component.go` -> `ComponentType` (Core, Simulation, Standalone)
- `internal/domain/component.go` -> `ComponentProvider` interface - abstraction for component discovery
- `internal/domain/buildconfig.go` -> `BuildConfig{Goal, Profiles, Port, SkipTests}` - Maven build parameters
- `internal/domain/buildconfig.go` -> `BuildConfig.ToCommand() string` - Generates Maven command string
- `internal/domain/buildconfig.go` -> `MavenGoal` (install, clean install, package, compile, test)
- `internal/domain/buildconfig.go` -> `BuildPhase` (Pending, Compiling, Testing, Packaging, Installing, Done, Failed)

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- Import domain types: `import "rfz-cli/internal/domain"`
- Use `domain.MockComponentProvider{}` to get components: `provider.Components()` returns `[]domain.Component`
- Each `Component` has `.Name` (string) and `.Type` (ComponentType with `.String()` method)
- Use `domain.BuildConfig{}.ToCommand()` to generate Maven command strings
- `BuildPhase` enum has `.String()` method for display text
- Component order in mock matches prototype screenshot (boss first, konfiguration last)

---

## File Change Summary

| File | Action | Story |
|------|--------|-------|
| internal/domain/component.go | Created | BUILD-001 |
| internal/domain/buildconfig.go | Created | BUILD-001 |
| internal/domain/mock_provider.go | Created | BUILD-001 |
| internal/domain/domain_test.go | Created | BUILD-001 |
