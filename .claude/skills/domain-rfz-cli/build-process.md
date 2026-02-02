# Domain: Build Process

Knowledge about Maven build workflows in the RFZ ecosystem.

## Maven Build Basics

### Standard Maven Goals

| Goal | Description | Usage |
|------|-------------|-------|
| `compile` | Compile source code | Fast check |
| `test` | Run unit tests | Verify tests pass |
| `package` | Create JAR/WAR | Create artifact |
| `install` | Install to local repo | Full build |
| `clean` | Delete target/ | Fresh start |

### Common Goal Combinations

```bash
# Full clean build
mvn clean install

# Skip tests for faster build
mvn clean install -DskipTests

# Compile only (fastest check)
mvn compile

# Run tests without packaging
mvn test
```

## Build Configuration

### Maven Profiles

Profiles customize the build for different environments.

**Common RFZ Profiles:**
- `dev` - Development settings
- `test` - Test environment
- `prod` - Production settings

**Usage:**
```bash
mvn install -P dev,test
```

### Skip Flags

| Flag | Effect |
|------|--------|
| `-DskipTests` | Skip test execution |
| `-Dmaven.test.skip=true` | Skip test compilation AND execution |
| `-DskipITs` | Skip integration tests |
| `-Dcheckstyle.skip` | Skip checkstyle |

## Build Output

### Log Levels

| Level | Shows | Maven Flag |
|-------|-------|------------|
| ERROR | Only errors | `-q` |
| WARN | Warnings + errors | default |
| INFO | Standard output | default |
| DEBUG | Verbose | `-X` |

### Progress Indicators

**Maven Lifecycle Phases Shown:**
1. `[INFO] Compiling...` - compile phase
2. `[INFO] Running tests...` - test phase
3. `[INFO] Building jar...` - package phase
4. `[INFO] Installing...` - install phase

### Error Patterns

| Pattern | Meaning |
|---------|---------|
| `[ERROR] COMPILATION ERROR` | Source code error |
| `[ERROR] Failed to execute goal` | Goal failed |
| `Tests run: X, Failures: Y` | Test failure |
| `Could not resolve dependencies` | Missing dependency |

## Multi-Module Builds

### Reactor Build Order

When building multiple modules, Maven uses "reactor" order:
1. Reads all pom.xml files
2. Calculates dependency graph
3. Builds in dependency order

### Parallel Builds (Maven)

```bash
# Build with 4 threads
mvn -T 4 install

# One thread per CPU core
mvn -T 1C install
```

## Build Configuration Modal

### Options in RFZ-CLI

| Option | Default | Description |
|--------|---------|-------------|
| Goal | `install` | Maven goal to execute |
| Clean | `true` | Run `clean` before goal |
| Skip Tests | `false` | Add `-DskipTests` |
| Profiles | [] | Active Maven profiles |
| Extra Args | "" | Additional CLI arguments |

### Build Config Storage

```go
type BuildConfig struct {
    ComponentID string
    Goal        string   // "compile", "test", "package", "install"
    Clean       bool     // prepend "clean"
    SkipTests   bool     // add -DskipTests
    Profiles    []string // -P profile1,profile2
    ExtraArgs   []string // any additional arguments
}
```

## Build Execution Flow

### In RFZ-CLI

1. **Select Components**: User picks components to build
2. **Configure Build**: Open build config modal
3. **Start Build**: Trigger Maven execution
4. **Stream Output**: Show logs in real-time
5. **Track Progress**: Update status per component
6. **Report Results**: Show success/failure summary

### Build States

```
Pending -> Compiling -> Testing -> Packaging -> Installing -> Success
                    \-> Failed (at any phase)
```

## Error Recovery

### Common Fixes

| Error | Solution |
|-------|----------|
| Dependency not found | Build dependencies first |
| Test failure | View logs, fix test |
| Out of memory | Increase Maven heap |
| Permission denied | Check file permissions |

### Retry Strategy

1. View error logs
2. Fix the issue
3. Rebuild failed components only
4. Or rebuild all with `clean`

## Performance Tips

- Use `-DskipTests` for quick iteration
- Build only changed components
- Use parallel builds when independent
- Keep local Maven repo clean

## Notes

- Document project-specific Maven quirks
- Add common error patterns as discovered
- Update with performance findings
