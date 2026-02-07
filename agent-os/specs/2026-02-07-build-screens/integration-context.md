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
| BUILD-004 | Build Execution View with simulator, progress, status counters | `execution.go`/`simulator.go` created, `model.go`/`update.go`/`view.go`/`app.go` modified |

---

## New Exports & APIs

### Components
<!-- New UI components created -->
- `internal/ui/screens/build/model.go` -> `build.Model` - Build screen with component selection (New/SetSize/SetFocused/Init/Update/View)
- `internal/ui/screens/build/model.go` -> `build.OpenConfigMsg{Selected []string}` - Message sent when user confirms selection
- `internal/ui/screens/build/model.go` -> `build.StartBuildMsg{Config, Selected}` - Message sent when user starts build from config
- `internal/ui/screens/build/config.go` -> Config modal view with 5 sections (Goal, Profiles, Port, Options, Buttons)
- `internal/ui/screens/build/execution.go` -> Execution view with component table, progress bar, status counters, actions
- `internal/ui/screens/build/simulator.go` -> Build simulator with timed phase transitions via tea.Tick

### Services
<!-- New service classes/modules -->
- `internal/domain/mock_provider.go` -> `MockComponentProvider{}` - Returns 13 hardcoded RFZ components

### Hooks / Utilities
<!-- New hooks, helpers, utilities -->
- `internal/ui/screens/build/model.go` -> `build.Model.CurrentItemLabel() string` - Returns label of cursor-focused component
- `internal/ui/screens/build/model.go` -> `build.Model.IsConfiguring() bool` - Returns true when in config phase
- `internal/ui/screens/build/model.go` -> `build.Model.OpenConfig(selected []string) Model` - Transitions to config phase
- `internal/ui/screens/build/model.go` -> `build.Model.SetTermSize(w, h int) Model` - Stores terminal dimensions for modal overlay
- `internal/ui/screens/build/model.go` -> `build.Model.IsExecuting() bool` - Returns true when build is running
- `internal/ui/screens/build/model.go` -> `build.Model.IsCompleted() bool` - Returns true when build has finished
- `internal/ui/screens/build/model.go` -> `build.BuildTickMsg` - Sent every 100ms during build execution
- `internal/ui/screens/build/model.go` -> `build.BuildPhaseMsg{ComponentIndex, Phase}` - Phase transition message
- `internal/ui/screens/build/model.go` -> `build.BuildCompleteMsg{}` - All components finished

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
- App handles `StartBuildMsg`, `BuildTickMsg`, `BuildPhaseMsg`, `BuildCompleteMsg` by delegating to build.Update()
- When `build.IsExecuting()` or `build.IsCompleted()` is true, app delegates ALL key events to build
- Status bar shows "BUILD" mode badge when executing, "DONE" mode badge when completed
- Simulator uses `tea.Tick(100ms)` pattern, no goroutines - max 3 concurrent component builds
- Build simulator has 20% chance of failure during Testing phase
- Esc during execution cancels build (marks running/pending as failed), "n" after completion resets to selection
- Execution view uses custom Lip Gloss table layout with TuiStatus badges, TuiProgress bars per row

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
| internal/ui/screens/build/execution.go | Created | BUILD-004 |
| internal/ui/screens/build/simulator.go | Created | BUILD-004 |
| internal/ui/screens/build/model.go | Modified | BUILD-004 |
| internal/ui/screens/build/update.go | Modified | BUILD-004 |
| internal/ui/screens/build/view.go | Modified | BUILD-004 |
| internal/app/app.go | Modified | BUILD-004 |
