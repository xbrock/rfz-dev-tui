# Integration Context

> **Purpose:** Cross-story context preservation for multi-session execution.
> **Auto-updated** after each story completion.
> **READ THIS** before implementing the next story.

---

## Completed Stories

| Story | Summary | Key Changes |
|-------|---------|-------------|
| BUILD-001 | Domain model with components, build config, mock provider | `internal/domain/` package created |
| BUILD-002 | Build Component Selection Screen with TuiList multi-select | `internal/ui/screens/build/` package created, app.go integrated |
| BUILD-003 | Build Configuration Modal with TuiModal, TuiRadio, TuiCheckbox | `config.go` created, `model.go`/`update.go`/`view.go`/`app.go` modified |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/ui/screens/build/model.go` -> `build.Model` - Build screen with component selection (New/SetSize/SetFocused/Init/Update/View)
- `internal/ui/screens/build/model.go` -> `build.OpenConfigMsg{Selected []string}` - Message sent when user confirms selection
- `internal/ui/screens/build/model.go` -> `build.StartBuildMsg{Config, Selected}` - Message sent when user starts build from config
- `internal/ui/screens/build/config.go` -> Config modal view with 5 sections (Goal, Profiles, Port, Options, Buttons)

### Services
<!-- New service classes/modules -->
- `internal/domain/mock_provider.go` -> `MockComponentProvider{}` - Returns 13 hardcoded RFZ components

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
- `internal/ui/screens/build/model.go` -> `build.Model.CurrentItemLabel() string` - Returns label of cursor-focused component
- `internal/ui/screens/build/model.go` -> `build.Model.IsConfiguring() bool` - Returns true when in config phase
- `internal/ui/screens/build/model.go` -> `build.Model.OpenConfig(selected []string) Model` - Transitions to config phase
- `internal/ui/screens/build/model.go` -> `build.Model.SetTermSize(w, h int) Model` - Stores terminal dimensions for modal overlay

### Types / Interfaces
<!-- New type definitions -->
- `internal/domain/component.go` -> `Component{Name, Type}` - RFZ component with name and category
- `internal/domain/component.go` -> `ComponentType` (Core, Simulation, Standalone)
- `internal/domain/component.go` -> `ComponentProvider` interface - abstraction for component discovery
- `internal/domain/buildconfig.go` -> `BuildConfig{Goal, Profiles, Port, SkipTests}` - Maven build parameters
- `internal/domain/buildconfig.go` -> `BuildConfig.ToCommand() string` - Generates Maven command string
- `internal/domain/buildconfig.go` -> `MavenGoal` (clean, install, clean install, package, compile, test)
- `internal/domain/buildconfig.go` -> `BuildPhase` (Pending, Compiling, Testing, Packaging, Installing, Done, Failed)

---

## Integration Notes

<!-- Important integration information for subsequent stories -->
- Import build screen: `import "rfz-cli/internal/ui/screens/build"`
- Build screen delegates to parent via `build.OpenConfigMsg` when Enter is pressed with selections
- App handles `OpenConfigMsg` by calling `build.OpenConfig(selected)` to transition to config phase
- Build screen has internal `buildPhase` state machine: selecting -> configuring -> executing -> completed
- Config phase renders TuiModal overlay over entire screen with 5 sections (Goal, Profiles, Port, Options, Buttons)
- Tab cycles between config sections, Esc cancels back to selection, Enter on Start Build sends `StartBuildMsg`
- When `build.IsConfiguring()` is true, app delegates ALL key events to build (including Tab/Esc)
- App.go handles focus delegation: Tab switches focus to content, then key events go to build.Update()
- Status bar shows "SELECT" mode badge when selecting, "CONFIG" mode badge when configuring
- Default build config: clean install, target_env_dev profile, port 11090, skip tests enabled
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
| internal/ui/screens/build/model.go | Created | BUILD-002 |
| internal/ui/screens/build/update.go | Created | BUILD-002 |
| internal/ui/screens/build/view.go | Created | BUILD-002 |
| internal/ui/screens/build/selection.go | Created | BUILD-002 |
| internal/app/app.go | Modified | BUILD-002 |
| internal/app/app_test.go | Modified | BUILD-002 |
| internal/app/testdata/TestApp_BuildScreen.golden | Created | BUILD-002 |
| internal/ui/screens/build/config.go | Created | BUILD-003 |
| internal/ui/screens/build/model.go | Modified | BUILD-003 |
| internal/ui/screens/build/update.go | Modified | BUILD-003 |
| internal/ui/screens/build/view.go | Modified | BUILD-003 |
| internal/domain/buildconfig.go | Modified | BUILD-003 |
| internal/app/app.go | Modified | BUILD-003 |
