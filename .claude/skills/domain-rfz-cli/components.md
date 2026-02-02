# Domain: RFZ Components

Knowledge about RFZ component types and their relationships.

## Component Types

### Core Components

Core components are essential modules that other components depend on.

**Characteristics:**
- Required by multiple other components
- Must build successfully before dependents
- Typically include: base libraries, shared utilities, data models

**Examples from prototype:**
- boss
- traktion
- signalsteuerung
- fistiv (may also be simulation)

### Simulation Components

Components used for testing and simulation environments.

**Characteristics:**
- Provide test doubles or simulators
- May not be needed in production builds
- Used for integration testing

**Examples:**
- simulator-*
- mock-*

### Standalone Components

Independent components with no external dependencies.

**Characteristics:**
- Can build independently
- Do not depend on other RFZ components
- Often utilities or tools

**Examples from prototype:**
- audiocon

## Component Structure

### Standard Maven Layout

```
component-name/
├── pom.xml              # Maven project file
├── src/
│   ├── main/
│   │   ├── java/        # Java source
│   │   └── resources/   # Resources
│   └── test/
│       ├── java/        # Test source
│       └── resources/   # Test resources
└── target/              # Build output (generated)
```

### Key pom.xml Fields

```xml
<project>
    <groupId>de.deutschebahn.rfz</groupId>
    <artifactId>component-name</artifactId>
    <version>1.0.0-SNAPSHOT</version>

    <!-- Used to detect component type -->
    <packaging>jar</packaging>

    <dependencies>
        <!-- Dependencies on other RFZ components -->
    </dependencies>
</project>
```

## Component Detection

### How CLI Discovers Components

1. **Scan Path**: Look in configured directories
2. **Find pom.xml**: Identifies Maven project
3. **Parse Metadata**: Extract groupId, artifactId, version
4. **Classify Type**: Based on naming conventions or pom content
5. **Check Dependencies**: Identify inter-component dependencies

### Detection Heuristics

| Pattern | Classified As |
|---------|---------------|
| `*-core`, `*-common`, `*-base` | Core |
| `*-simulator`, `*-mock`, `*-test` | Simulation |
| No RFZ dependencies | Standalone |
| Has RFZ dependencies | Core (if depended on) or Standard |

## Component Registry

### components.yaml Format

```yaml
# configs/components.yaml
components:
  - id: boss
    name: Boss
    type: core
    path: /path/to/boss
    dependencies: []

  - id: audiocon
    name: Audiocon
    type: standalone
    path: /path/to/audiocon
    dependencies: []

  - id: traktion
    name: Traktion
    type: core
    path: /path/to/traktion
    dependencies:
      - boss
```

### Registry Management

- **Auto-discovery**: Scan paths populate registry
- **Manual entries**: Can add components manually
- **Persistence**: Saved between sessions
- **Refresh**: Re-scan updates registry

## Dependency Order

### Build Order Calculation

When building multiple components:

1. Build components with no dependencies first
2. Build components whose dependencies are satisfied
3. Repeat until all built or failure

**Example:**
- boss (no deps) -> build first
- traktion (depends on boss) -> build second
- audiocon (standalone) -> can build anytime

### Parallel Builds (Future)

- Independent components can build in parallel
- Dependent components wait for dependencies
- Optimize for fastest overall completion

## Component States

| State | Symbol | Meaning |
|-------|--------|---------|
| Pending | `o` | Not yet started |
| Building | `.` | Currently building (with phase) |
| Success | `v` | Build completed successfully |
| Failed | `x` | Build failed |

### Build Phases

During building, component passes through Maven phases:
1. compile
2. test
3. package
4. install

## Notes

- Update when new component types emerge
- Document component-specific build quirks
- Add dependency rules as discovered
