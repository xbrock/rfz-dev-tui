# Test Scenarios - Welcome & Navigation Screen

**Spec:** 2026-02-07-welcome-navigation-screen
**Datum:** 2026-02-07

## Happy Path

### Scenario 1: Application Launch
1. Run `go run ./cmd/rfz`
2. Verify: Alt-screen opens with header, navigation sidebar, welcome content, status bar
3. Verify: Header shows "RFZ-CLI v1.0.0", "Terminal Orchestration Tool", live clock, "Deutsche Bahn Internal"
4. Verify: Navigation sidebar shows 5 items (Build Components, View Logs, Discover, Configuration, Exit)
5. Verify: Welcome screen shows ASCII art "RFZ CLI" logo, version badge, status line
6. Verify: Status bar shows "HOME" badge, key hints

### Scenario 2: Navigate Between Screens
1. Press `1` - Verify: "Build Components - Coming Soon" placeholder shown, nav highlights Build
2. Press `2` - Verify: "View Logs - Coming Soon" placeholder shown, nav highlights Logs
3. Press `3` - Verify: "Discover - Coming Soon" placeholder shown, nav highlights Discover
4. Press `4` - Verify: "Configuration - Coming Soon" placeholder shown, nav highlights Config
5. Press `Esc` - Verify: Returns to Welcome screen

### Scenario 3: Keyboard Navigation (Arrow Keys)
1. Press `j` (or Down) - Verify: Cursor moves to "View Logs" (index 1)
2. Press `j` again - Verify: Cursor moves to "Discover" (index 2)
3. Press `k` (or Up) - Verify: Cursor moves back to "View Logs" (index 1)
4. Press `Enter` - Verify: "View Logs" placeholder screen shown

### Scenario 4: Exit Confirmation
1. Press `q` - Verify: Modal dialog appears: "Quit RFZ-CLI?"
2. Verify: "No" button is focused by default (safety)
3. Press `n` - Verify: Modal closes, returns to previous state
4. Press `q` again - Verify: Modal reappears
5. Press `y` - Verify: Application exits cleanly

### Scenario 5: Focus Switching
1. Press `Tab` - Verify: Focus switches from navigation to content area
2. Verify: Content area border changes to cyan (focused)
3. Press `Tab` again - Verify: Focus returns to navigation sidebar

### Scenario 6: Wrap-around Navigation
1. Press `k` from top (index 0) - Verify: Cursor wraps to Exit (index 4)
2. Press `j` from Exit (index 4) - Verify: Cursor wraps to Build (index 0)

## Edge Cases

### Scenario 7: Terminal Too Small
1. Resize terminal to below 80x24
2. Verify: Shows "Terminal too small. Please resize to at least 80x24." message
3. Resize back to 80x24+ - Verify: Normal layout restored

### Scenario 8: Modal Key Isolation
1. Press `q` to open exit modal
2. Press `1`, `2`, `3`, `4` - Verify: No screen switching happens (modal captures input)
3. Press `j`, `k` - Verify: No navigation movement (modal captures input)
4. Press `Esc` - Verify: Modal closes

### Scenario 9: Modal Button Navigation
1. Press `q` to open exit modal
2. Press `Left` or `Right` - Verify: Focus toggles between Yes/No buttons
3. Press `Tab` - Verify: Focus toggles between Yes/No buttons
4. Focus on "No", press `Enter` - Verify: Modal closes
5. Press `q`, focus on "Yes", press `Enter` - Verify: Application exits

### Scenario 10: Ctrl+C Force Quit
1. At any point, press `Ctrl+C` - Verify: Application exits immediately (no modal)
2. With modal open, press `Ctrl+C` - Verify: Application exits immediately

### Scenario 11: Enter on Exit Nav Item
1. Navigate cursor to "Exit" (index 4) using `j` four times
2. Press `Enter` - Verify: Exit confirmation modal appears (not a screen switch)

## Fehlerszenarien

### Scenario 12: Component Demo Still Works
1. Run `go run ./cmd/rfz-components-demo`
2. Verify: Component gallery demo launches correctly
3. Verify: Old layout-demo entry point no longer exists

### Scenario 13: Rapid Key Input
1. Rapidly press `1 2 3 4 Esc 1 2 3 4` in quick succession
2. Verify: Application handles all inputs correctly without crash or visual glitch

## Automated Test Coverage

The following UI states are covered by golden file visual regression tests (`internal/app/app_test.go`):

| Test | State | Terminal Size |
|------|-------|---------------|
| TestApp_WelcomeDefault | Initial welcome screen | 120x40 |
| TestApp_NavBuildFocused | Build nav item cursor | 120x40 |
| TestApp_NavLogsFocused | Logs nav item cursor | 120x40 |
| TestApp_NavDiscoverFocused | Discover nav item cursor | 120x40 |
| TestApp_NavConfigFocused | Config nav item cursor | 120x40 |
| TestApp_NavExitFocused | Exit nav item cursor | 120x40 |
| TestApp_PlaceholderBuild | Build placeholder screen | 120x40 |
| TestApp_PlaceholderLogs | Logs placeholder screen | 120x40 |
| TestApp_PlaceholderDiscover | Discover placeholder screen | 120x40 |
| TestApp_PlaceholderConfig | Config placeholder screen | 120x40 |
| TestApp_ExitModal | Quit confirmation modal | 120x40 |
| TestApp_TerminalTooSmall | Small terminal warning | 60x15 |
