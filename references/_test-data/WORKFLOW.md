# RFZ Developer CLI - Build Workflow Documentation

## The Problem

RFZ developers need to rebuild components frequently during development. The Maven build process involves multiple profiles and different build targets depending on what the developer wants to achieve. Currently, developers must remember complex Maven commands with multiple profile activations.

### Pain Points

1. **Complex Maven Commands**: Building with correct profiles requires remembering long commands
2. **Inconsistent Components**: Components vary in structure and available profiles
3. **Config Generation**: Generating local config files requires specific profile combinations
4. **Multiple Repositories**: Some components are in the main repo, others are standalone
5. **Directory Structure Variations**: Developers organize their workspace differently

---

## Component Variations

### Structure Variations

| Aspect | Example: AudioCon | Example: SigSim | Example: MGW-Connector |
|--------|-------------------|-----------------|------------------------|
| Type | Core Component | Simulator | Standalone |
| Location | `rfz/rfz/audiocon` | `rfz/rfz/sigsim` | Standalone repo |
| Has Config Module | ✅ Yes | ✅ Yes | ❌ No |
| Port Profiles | 2 (11090, 11091) | 0 | 1 (11090 only) |
| Config Generation | ✅ Supported | ✅ Supported | ❌ N/A |

### Key Differences

#### Core Components (e.g., AudioCon, FISDV, DisCon)
- Located in main `rfz/rfz` repository
- **Have config modules** (e.g., `audiocon-config/`)
- Support **TWO port profiles**: `target_env_use_traktion_11090` and `target_env_use_traktion_11091`
- Support config generation with `generate_local_config_files` profile
- Used in production trains

#### Simulators (e.g., SigSim, RISSim)
- Located in main `rfz/rfz` repository
- **Have config modules** (e.g., `sigsim-config/`)
- **NO port profiles** (simulators don't need port differentiation)
- Support config generation with `generate_local_config_files` profile
- Used for development/testing only

#### Standalone Components (e.g., MGW-Connector)
- Located in **separate repositories** (not in `rfz/rfz`)
- **NO config modules** (configuration is handled differently)
- May have **ONE or NO port profiles** depending on component
- NO config generation support
- Independent release cycle

---

## Profile Availability Matrix

| Profile | Core Components | Simulators | Standalone | Description |
|---------|-----------------|------------|------------|-------------|
| `target_env_dev` | ✅ | ✅ | ✅ | Development environment |
| `target_env_test` | ✅ | ✅ | ✅ | Test/Staging environment |
| `target_env_prod` | ✅ | ❌ | ❌ | Production environment |
| `target_env_use_traktion_11090` | ✅ | ❌ | ✅/❌ | Port 11090 profile |
| `target_env_use_traktion_11091` | ✅ | ❌ | ❌ | Port 11091 profile (redundancy) |
| `generate_local_config_files` | ✅ | ✅ | ❌ | Triggers config file generation |

---

## Build Workflows

### Workflow 1: Standard Component Build

**Goal**: Rebuild one or more components with default settings (no config generation).

**What developers want**: Just recompile and test the component code.

**Maven Command**:
```bash
cd <component-root>
mvn clean install -P target_env_dev
```

**Example for AudioCon**:
```bash
cd D:/apps/rfz/rfz/audiocon
mvn clean install -P target_env_dev
```

**What it does**:
- Cleans previous build artifacts (`target/` folders)
- Compiles all modules (config + core)
- Runs unit tests
- Installs artifacts to local Maven repository (`~/.m2/repository`)
- Uses development environment settings
- **Does NOT generate config files**

**Build time**: ~1-3 minutes depending on component size

---

### Workflow 2: Config-Only Build

**Goal**: Regenerate local configuration files **without** full rebuild.

**What developers want**: Update config files after changing templates or properties, without waiting for full compilation.

**Maven Command** (navigate to config module):
```bash
cd <component-root>/<component>-config
mvn clean install -P target_env_dev,generate_local_config_files
```

**With port profile** (for components that support it):
```bash
cd <component-root>/<component>-config
mvn clean install -P target_env_dev,target_env_use_traktion_11090,generate_local_config_files
```

**Example for AudioCon with port 11090**:
```bash
cd D:/apps/rfz/rfz/audiocon/audiocon-config
mvn clean install -P target_env_dev,target_env_use_traktion_11090,generate_local_config_files
```

**Example for AudioCon with port 11091**:
```bash
cd D:/apps/rfz/rfz/audiocon/audiocon-config
mvn clean install -P target_env_dev,target_env_use_traktion_11091,generate_local_config_files
```

**Example for SigSim (no port profile)**:
```bash
cd D:/apps/rfz/rfz/sigsim/sigsim-config
mvn clean install -P target_env_dev,generate_local_config_files
```

**What it does**:
- Only builds the config module (much faster)
- Generates local configuration files based on templates
- Port profile determines which port configuration to generate (if applicable)
- Config files are typically generated to `target/config/` or a local config directory

**Build time**: ~10-30 seconds

**Important**: This workflow is **NOT available** for standalone components without config modules.

---

### Workflow 3: Full Build with Config Generation

**Goal**: Complete rebuild including config file generation (everything at once).

**What developers want**: Build everything and get fresh config files in one step.

**Maven Command** (from component root):
```bash
cd <component-root>
mvn clean install -P target_env_dev,generate_local_config_files
```

**With port profile**:
```bash
cd <component-root>
mvn clean install -P target_env_dev,target_env_use_traktion_11090,generate_local_config_files
```

**Example for AudioCon with port 11091**:
```bash
cd D:/apps/rfz/rfz/audiocon
mvn clean install -P target_env_dev,target_env_use_traktion_11091,generate_local_config_files
```

**Example for SigSim**:
```bash
cd D:/apps/rfz/rfz/sigsim
mvn clean install -P target_env_dev,generate_local_config_files
```

**What it does**:
- Full clean and install of all modules
- Also generates local configuration files during the install phase
- Combines standard build with config generation

**Build time**: ~1-3 minutes

---

## User Directory Structures

Developers organize their RFZ workspace in different ways. The CLI must be flexible enough to handle all of these:

### Structure 1: Standard GitHub Structure

This is the "official" structure as cloned from Git:

```
~/projects/
├── rfz/
│   └── rfz/                    # Main repository (double-nested!)
│       ├── audiocon/
│       ├── discon/
│       ├── fisdv/
│       ├── sigsim/
│       └── ... (30+ components)
├── mgw-connector/              # Standalone repo
├── monitoring/                 # Standalone repo
└── rfz-notification-service/   # Standalone repo
```

**Challenge**: The main repo has a `rfz/rfz/` double-nesting which can be confusing.

---

### Structure 2: Custom Flat Structure

Some developers flatten everything to avoid nesting:

```
~/rfz-projects/
├── audiocon/
├── discon/
├── fisdv/
├── sigsim/
├── mgw-connector/
└── monitoring/
```

**How they do it**: After cloning, they move component folders out of `rfz/rfz/` to a flat directory.

**Challenge**: Loses the distinction between main repo and standalone components.

---

### Structure 3: Custom Categorized Structure

Some developers organize by component type:

```
~/dev/
├── core/
│   ├── audiocon/
│   ├── discon/
│   └── fisdv/
├── simulators/
│   ├── sigsim/
│   └── rissim/
└── standalone/
    ├── mgw-connector/
    └── monitoring/
```

**Challenge**: CLI needs to search multiple directories to find components.

---

### Structure 4: Mixed Structure

Real-world scenario where some components are in the main repo, others are standalone:

```
D:/apps/
├── rfz/
│   └── rfz/
│       ├── audiocon/
│       ├── fisdv/
│       └── sigsim/
├── mgw-connector/              # Separate repo
├── monitoring/                 # Separate repo
└── rfz-developers-cli/         # The CLI itself
```

**This is the most common structure** in practice.

---

## Maven Profile Reference

### Environment Profiles

| Profile | Description | When to Use |
|---------|-------------|-------------|
| `target_env_dev` | Development environment | **Always** - default for local development |
| `target_env_test` | Test/Staging environment | When deploying to test train/system |
| `target_env_prod` | Production environment | When building for production trains |

### Port Profiles

RFZ trains have **two independent systems** (11090 and 11091) for redundancy.

| Profile | Description | Port | When to Use |
|---------|-------------|------|-------------|
| `target_env_use_traktion_11090` | Primary port configuration | 11090 | For system A (primary) |
| `target_env_use_traktion_11091` | Secondary port (redundancy) | 11091 | For system B (backup) |

**Note**: Not all components support both port profiles. Simulators typically have no port profiles.

### Special Profiles

| Profile | Description | When to Use |
|---------|-------------|-------------|
| `generate_local_config_files` | Triggers config file generation during install phase | When you need to regenerate config files |

---

## Component Detection Logic

The CLI uses a **static component registry** approach rather than dynamic discovery. This provides better control and consistency.

### How It Works

1. **Static Registry**: Component metadata is defined in a configuration file (`configs/components.yaml`)
2. **Directory Scanning**: CLI recursively scans configured directories to locate these known components
3. **Matching by ArtifactId**: Components are identified by matching their Maven `artifactId` in `pom.xml`
4. **User Configuration**: Scan paths and component list are user-configurable

### Component Registry Structure

The registry defines component metadata that doesn't change:

```yaml
components:
  - artifactId: audiocon
    name: AudioCon
    category: RFZ Core
    defaultProfiles: [target_env_dev]
    portProfiles:
      11090: []  # No profile needed for 11090
      11091: [target_env_use_traktion_11091]
    buildOrder: 30
    requiresMaven: true

  - artifactId: sigsim
    name: SigSim
    category: Simulators
    defaultProfiles: []  # Simulators don't use target_env
    portProfiles: {}     # No port profiles
    buildOrder: 100
    requiresMaven: true

  - artifactId: rfz-mgw-connector
    name: MGW-Connector
    category: Standalone
    defaultProfiles: []  # Standalone don't use target_env
    portProfiles:
      11090: [target_env_use_traktion_11090]
    buildOrder: 200
    requiresMaven: true
```

### Directory Scanning Configuration

Users can configure where to search for components:

```yaml
# User configuration (~/.rfz/config.yaml)
scan:
  paths:
    - D:/apps/rfz/rfz          # Main RFZ repository
    - D:/apps/rfz-components   # Alternative component location
    - D:/apps                  # Scan for standalone components
  recursive: true
  maxDepth: 5                  # Prevent infinite recursion
  exclude:
    - "*/target/*"
    - "*/.git/*"
```

**Default behavior**: If no paths configured, scan current directory `./` recursively.

### Scanning Algorithm

1. **Load registry**: Read component definitions from `configs/components.yaml`
2. **Scan directories**: Recursively search configured paths for `pom.xml` files
3. **Extract artifactId**: Parse each `pom.xml` to get `<artifactId>` element
4. **Match with registry**: Check if artifactId matches any registered component
5. **Detect config module**: Look for `<component>-config/` subdirectory
6. **Store location**: Record the absolute path where component was found

### Detected Component Structure

After scanning, the CLI knows both the **metadata** (from registry) and **location** (from scan):

```go
type DetectedComponent struct {
    // From registry
    ArtifactId       string
    Name             string
    Category         string
    DefaultProfiles  []string
    PortProfiles     map[int][]string
    BuildOrder       int
    
    // From scanning
    Path             string   // Absolute path where found
    HasConfigModule  bool     // Detected by checking for <component>-config/
    PomFile          string   // Path to pom.xml
    
    // Computed
    IsAvailable      bool     // Was it found during scan?
}
```

### Example: Detected Components

**AudioCon** (found in registry, located during scan):
```go
DetectedComponent{
    // From registry
    ArtifactId: "audiocon",
    Name: "AudioCon",
    Category: "RFZ Core",
    DefaultProfiles: ["target_env_dev"],
    PortProfiles: map[int][]string{
        11090: [],  // No profile for 11090
        11091: ["target_env_use_traktion_11091"],
    },
    BuildOrder: 30,
    
    // From scanning
    Path: "D:/apps/rfz-components/audiocon",
    HasConfigModule: true,  // Found audiocon-config/
    PomFile: "D:/apps/rfz-components/audiocon/pom.xml",
    
    // Computed
    IsAvailable: true,  // Found during scan
}
```

**FISDV** (in registry but NOT found during scan):
```go
DetectedComponent{
    // From registry
    ArtifactId: "fisdv",
    Name: "FISDV",
    Category: "RFZ Core",
    DefaultProfiles: ["target_env_dev"],
    PortProfiles: map[int][]string{
        11090: ["target_env_use_traktion_11090"],
        11091: ["target_env_use_traktion_11091"],
    },
    BuildOrder: 13,
    
    // From scanning
    Path: "",  // Not found
    HasConfigModule: false,
    PomFile: "",
    
    // Computed
    IsAvailable: false,  // User doesn't have this component
}
```

---

## Benefits of Static Registry Approach

1. **Consistent Metadata**: Component names, categories, and profiles are centrally defined
2. **Build Order Control**: Components can be built in a specific order (dependencies first)
3. **Profile Knowledge**: CLI knows which profiles are valid without parsing complex POMs
4. **Performance**: No need to parse every POM file for profile information
5. **Flexibility**: Users can still organize directories however they want
6. **Maintainability**: Adding new components means updating one registry file
7. **Validation**: CLI can warn about unknown components or incorrect profiles

---

## User Configuration

Users can customize where the CLI searches for components:

### Default Configuration
If user doesn't configure anything, the CLI scans the current directory:

```yaml
# ~/.rfz/config.yaml (auto-generated)
scan:
  paths:
    - ./  # Current directory
  recursive: true
  maxDepth: 5
```

### Custom Multi-Location Setup
For users with components in multiple locations:

```yaml
# ~/.rfz/config.yaml
scan:
  paths:
    - D:/apps/rfz/rfz              # Main repo
    - D:/apps/rfz-components        # Individual components
    - D:/apps                       # Standalone components
    - C:/projects/custom-rfz        # Custom location
  recursive: true
  maxDepth: 5
  exclude:
    - "*/node_modules/*"
    - "*/target/*"
    - "*/.git/*"
```

### Single Flat Directory
For users who organize all components in one folder:

```yaml
# ~/.rfz/config.yaml
scan:
  paths:
    - D:/rfz-workspace  # All components in one folder
  recursive: false      # No need to recurse
```

---

## Appendix: Real Maven Commands

For reference, here are the actual Maven commands developers type today:

### Standard Build
```bash
mvn clean install -P target_env_dev
```

### Build with Port 11090
```bash
mvn clean install -P target_env_dev,target_env_use_traktion_11090
```

### Config Only (Port 11090)
```bash
cd audiocon-config
mvn clean install -P target_env_dev,target_env_use_traktion_11090,generate_local_config_files
```

### Full Build with Config (Port 11091)
```bash
mvn clean install -P target_env_dev,target_env_use_traktion_11091,generate_local_config_files
```

### Parallel Build
```bash
mvn clean install -P target_env_dev -T 4
```

### Skip Tests
```bash
mvn clean install -P target_env_dev -DskipTests
```

### Offline Mode
```bash
mvn clean install -P target_env_dev -o
```

---

**This documentation describes the problem domain and build workflows. It does NOT describe the current CLI implementation.**
