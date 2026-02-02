# Domain: Developer Workflows

Common workflows for RFZ developers using the CLI.

## Daily Workflows

### Morning Start

**Goal:** Get workspace ready for development

1. Open RFZ-CLI
2. Navigate to Discover (press `3`)
3. Review Git status of components
4. Pull latest changes if needed (external)
5. Build all components to verify

**In CLI:**
```
Press 3 -> Review status -> Press 1 -> Select all (a) -> Build
```

### Quick Build Cycle

**Goal:** Build after code changes

1. Make code changes (external editor)
2. Open RFZ-CLI / Navigate to Build
3. Select changed component(s)
4. Build with skip tests
5. Review result

**In CLI:**
```
Press 1 -> Navigate to component -> Space to select -> Enter -> Skip tests -> Build
```

### Full Build Before Commit

**Goal:** Verify all tests pass before commit

1. Navigate to Build screen
2. Select all components (`a`)
3. Configure: DO NOT skip tests
4. Execute build
5. Review logs if failures
6. Fix issues, repeat

**In CLI:**
```
Press 1 -> a (select all) -> Enter -> Ensure tests enabled -> Build
```

### Investigating Build Failure

**Goal:** Find and fix build error

1. Note which component failed
2. Navigate to Logs (press `2`)
3. Select failed component
4. Enable error-only filter (`e`)
5. Find error message
6. Fix issue
7. Rebuild

**In CLI:**
```
Press 2 -> Select component -> e (error filter) -> Find error -> Fix -> Press 1 -> Rebuild
```

## Component Management

### Adding a New Component

1. Create component directory (external)
2. Create pom.xml (external)
3. Open RFZ-CLI
4. Navigate to Configuration (`4`)
5. Verify scan path includes new location
6. Navigate to Discover (`3`)
7. Trigger scan if needed
8. Verify component appears
9. Build to verify

### Checking Component Status

1. Navigate to Discover (`3`)
2. View component list
3. Check Git status (branch, dirty/clean)
4. Check last commit info

### Updating Scan Paths

1. Navigate to Configuration (`4`)
2. Select Scan Paths section
3. Add new path or edit existing
4. Save changes
5. Return to Discover to rescan

## Build Scenarios

### Build Single Component

```
1 -> Navigate to component -> Space -> Enter -> Build
```

### Build All Core Components

```
1 -> Filter/navigate to Core type -> Select all Core -> Enter -> Build
```

### Build Component and Dependencies

```
1 -> Select target component -> Auto-select deps (future) -> Enter -> Build
```

### Rebuild Failed Components Only

After a partial failure:
```
2 (review logs) -> Fix issue -> 1 -> Select only failed -> Enter -> Build
```

## Log Analysis

### Real-time Log Monitoring

During build:
```
L (while building) -> Watch logs -> f (toggle follow)
```

### Searching Logs

```
2 -> Select component -> / (search) -> Type query -> Navigate results
```

### Filter by Level

```
2 -> Select component -> e (errors only) or number keys (1-5 for levels)
```

## Keyboard Shortcuts Summary

### Global Navigation

| Key | Action |
|-----|--------|
| `1` | Go to Build screen |
| `2` | Go to Logs screen |
| `3` | Go to Discover screen |
| `4` | Go to Configuration |
| `q` | Quit (with confirm) |
| `?` | Help |

### List Navigation

| Key | Action |
|-----|--------|
| `j` / `Down` | Move down |
| `k` / `Up` | Move up |
| `Space` | Toggle selection |
| `Enter` | Confirm/Open |
| `a` | Select all |
| `n` | Clear selection |

### Build Actions

| Key | Action |
|-----|--------|
| `Enter` | Open build config (from selection) |
| `L` | View logs (during build) |
| `Ctrl+C` | Cancel build |

### Log Viewer

| Key | Action |
|-----|--------|
| `f` | Toggle follow mode |
| `e` | Toggle error-only |
| `/` | Search |

## Workflow Tips

### Efficiency

- Use number keys for quick screen navigation
- Learn `j/k` navigation (faster than arrows)
- Use `a` to select all, then deselect unwanted
- Skip tests during iteration, run before commit

### Troubleshooting

- Always check logs after failure
- Error filter (`e`) helps find issues quickly
- Check Git status if build suddenly fails
- Verify dependencies built successfully

### Best Practices

- Build dependencies before dependents
- Run full build with tests before commits
- Keep scan paths organized
- Review discover regularly for new components

## Notes

- Add team-specific workflows as needed
- Document project-specific shortcuts
- Update with efficiency discoveries
