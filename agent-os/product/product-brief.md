# Product Brief: RFZ Developer CLI

> Terminal Orchestration Tool for RFZ Developers
> Last Updated: 2026-02-02
> Version: 1.0.0
> Organization: Deutsche Bahn (Internal Tool)

---

## Executive Summary

The RFZ Developer CLI is a Go-based Terminal User Interface (TUI) application built with the charm.land stack that streamlines the Maven build workflow for RFZ developers. It provides an intuitive interface for component selection, build configuration, execution monitoring, and log viewing, reducing the cognitive load of managing complex Maven builds across a heterogeneous component ecosystem.

---

## 1. Product Mission

### Pitch

A Go-based Terminal User Interface (TUI) application built with the charm.land stack that streamlines the Maven build workflow for RFZ developers by providing an intuitive interface for component selection, build configuration, execution monitoring, and log viewing.

### Vision Statement

Empower RFZ developers at Deutsche Bahn with a purpose-built terminal tool that transforms the complex, error-prone process of managing Maven builds into an efficient, visually-rich, keyboard-driven workflow. By deeply understanding the RFZ component ecosystem, the tool eliminates friction in daily development tasks while providing confidence through comprehensive visual regression testing.

---

## 2. Target Users

### Primary User Segments

| Segment | Description | Key Needs |
|---------|-------------|-----------|
| **RFZ Developers** | Deutsche Bahn developers who frequently rebuild Maven components | Quick component selection, correct profile application, build status visibility |
| **Multi-Component Managers** | Developers managing builds across Core, Simulators, and Standalone components | Category-based organization, batch operations, consistent interface |
| **Terminal Power Users** | Technical users comfortable with CLI who want efficient keyboard-driven workflows | Fast navigation, keyboard shortcuts, minimal mouse usage |

### User Personas

**Persona 1: Daily Builder**
- Rebuilds 2-5 components multiple times per day
- Needs to quickly select components and kick off builds
- Values build status visibility and log access
- Pain: Typing long Maven commands with correct profiles

**Persona 2: Integration Developer**
- Works across multiple component types (Core + Simulators)
- Needs to understand Git status and component health
- Values discovery and scanning features
- Pain: Tracking which components are clean/dirty across repos

**Persona 3: New Team Member**
- Learning the RFZ component ecosystem
- Needs guidance on available profiles and options
- Values clear UI and discoverable features
- Pain: Not knowing which Maven profiles to use for which components

---

## 3. Problem Statement

### Core Problems

| # | Problem | Impact | Current Workaround |
|---|---------|--------|-------------------|
| 1 | **Complex Maven Commands** | Building with correct profiles requires remembering long commands with specific flags and profile combinations | Developers maintain shell aliases or scripts, which become outdated |
| 2 | **Inconsistent Components** | Components vary in structure and available profiles; no single source of truth | Manual documentation, tribal knowledge |
| 3 | **Config Generation Complexity** | Generating local config files requires specific profile combinations that vary by component | Copy-paste from team docs, frequent errors |
| 4 | **Multiple Repositories** | Some components are in the main repo, others are standalone; different Git workflows | Manual tracking, context switching |
| 5 | **Directory Structure Variations** | Developers organize their workspace differently; tools must adapt | Hardcoded paths, per-developer configuration |

### Problem Severity

```
Complex Maven Commands      [##########] Critical - Daily friction
Inconsistent Components     [########  ] High - Causes build failures
Config Generation           [#######   ] High - Time-consuming errors
Multiple Repositories       [######    ] Medium - Context switching overhead
Directory Variations        [#####     ] Medium - Setup friction for new devs
```

---

## 4. Solution Overview

### Core Features

#### 4.1 Build Components Screen
The primary interface for selecting which components to build.

**Capabilities:**
- Multi-select component list with category badges (Core, Simulator, Standalone)
- Visual distinction between component types
- Keyboard navigation with vim-style shortcuts
- Component status indicators (last build result, Git status)
- Filter and search functionality

**User Flow:**
1. User navigates to Build screen (keyboard shortcut or navigation)
2. User selects one or more components using checkboxes
3. User presses Enter or 'b' to open Build Configuration modal
4. Selected components are built with chosen configuration

#### 4.2 Build Configuration Modal
Modal dialog for configuring build parameters before execution.

**Capabilities:**
- Maven goals selection (clean, install, package, etc.)
- Profile selection (component-specific available profiles)
- Port selection for local config (11090/11091)
- Skip tests toggle
- Additional Maven flags input
- Configuration presets

**Configuration Options:**
| Option | Type | Default | Description |
|--------|------|---------|-------------|
| Goals | Multi-select | clean install | Maven goals to execute |
| Profiles | Multi-select | (component-specific) | Maven profiles to activate |
| Port | Radio | 11090 | Local config port |
| Skip Tests | Checkbox | false | Add -DskipTests flag |
| Skip Checkstyle | Checkbox | false | Add -Dcheckstyle.skip flag |
| Offline | Checkbox | false | Add -o flag |

#### 4.3 Build Execution View
Real-time monitoring of build progress and status.

**Capabilities:**
- Real-time progress indicators per component
- Status transitions: Pending -> Compiling -> Testing -> Packaging -> Installing -> Done/Failed
- Live Maven output streaming
- Abort/cancel build functionality
- Build time tracking
- Parallel build visualization

**Status States:**
| Status | Visual | Description |
|--------|--------|-------------|
| Pending | Gray circle | Queued, not yet started |
| Compiling | Yellow spinner | compile phase running |
| Testing | Blue spinner | test phase running |
| Packaging | Purple spinner | package phase running |
| Installing | Cyan spinner | install phase running |
| Done | Green checkmark | Build succeeded |
| Failed | Red X | Build failed |

#### 4.4 Log Viewer
Component-based log viewing with filtering and navigation.

**Capabilities:**
- Component-based log selection
- Log level filters (ALL, INFO, WARN, ERROR, DEBUG)
- Follow mode (auto-scroll to latest)
- Search within logs
- Copy log content
- Clear logs
- Export logs to file

**Filter Levels:**
| Filter | Shows | Keyboard |
|--------|-------|----------|
| ALL | Everything | 1 |
| INFO | INFO and above | 2 |
| WARN | WARN and above | 3 |
| ERROR | ERROR only | 4 |
| DEBUG | All including DEBUG | 5 |

#### 4.5 Discover Screen
Repository scanning and component detection.

**Capabilities:**
- Scan configured paths for components
- Display Git status (clean/dirty) per component
- Show current branch per repository
- Display last commit info
- Detect component type (Core, Simulator, Standalone)
- Manual refresh and auto-refresh options

**Discovery Information:**
| Field | Description |
|-------|-------------|
| Component Name | Detected from pom.xml artifactId |
| Type | Core, Simulator, or Standalone |
| Path | Absolute path to component root |
| Git Status | Clean, Dirty (modified files count) |
| Branch | Current Git branch name |
| Last Commit | Short hash + message + relative time |

#### 4.6 Configuration Screen
Application settings and component registry management.

**Capabilities:**
- Scan paths management (add/remove/edit paths to scan)
- Component registry view (all known components)
- Detected components status (found/missing)
- Maven settings (MAVEN_HOME, settings.xml path)
- Terminal settings (colors, size preferences)
- Export/import configuration

**Configuration Sections:**
| Section | Description |
|---------|-------------|
| Scan Paths | Directories to scan for components |
| Component Registry | Known components with expected locations |
| Maven Settings | Maven installation and configuration |
| Display Settings | Terminal and visual preferences |

---

## 5. Technical Architecture

### Technology Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| Language | Go 1.21+ | Performance, single binary deployment |
| TUI Framework | Bubble Tea (charm.land) | Elm-architecture TUI framework |
| Styling | Lip Gloss (charm.land) | Terminal styling and layouts |
| Components | Bubbles (charm.land) | Pre-built TUI components (list, table, viewport, etc.) |
| Logging | charmbracelet/log | Structured colorized logging with Lip Gloss integration |
| Testing | teatest + image comparison | Hybrid visual regression testing |
| CI/CD | GitHub Actions | Automated testing and releases |

### Canonical Terminal Size

- **Width:** 120 columns
- **Height:** 40 rows
- **Rationale:** Common developer terminal size, sufficient for complex UIs

### Component Library (Design System)

The application uses a hybrid approach: leveraging Bubbles built-in components where available, and building custom components only where needed. All components follow idiomatic Bubble Tea patterns (Model-Update-View).

#### Use from Bubbles (charm.land/bubbles)

These components are provided by the Bubbles library and should be used directly:

| Bubbles Component | Usage in RFZ CLI | Key Features |
|-------------------|------------------|--------------|
| **list** | Component selection, navigation lists | Filtering, pagination, help integration, spinner |
| **progress** | Build progress indicators | Animation via Harmonica, customizable |
| **spinner** | Loading states during builds | Multiple built-in styles |
| **textinput** | Configuration inputs, search | Unicode, pasting, in-place scrolling |
| **table** | Component registry, detected components | Scrolling, column customization |
| **viewport** | Log viewer content area | Scrollable content, essential for logs |
| **help** | Keyboard shortcut display | Standardized key binding format |
| **paginator** | Log pagination (optional) | Navigation controls |

#### Build Custom (not in Bubbles)

These components require custom implementation:

| Component | Variants | Description |
|-----------|----------|-------------|
| TuiBox | single, double, rounded, heavy, focused | Lip Gloss wrapper with focus state styling |
| TuiDivider | single, double | Simple horizontal/vertical separators |
| TuiNavigation | - | Sidebar navigation container |
| TuiNavItem | normal, focused, active | Navigation menu item with shortcut display |
| TuiStatus | pending, running, success, failed, error | Build status badges with icons |
| TuiButton | default, primary, danger | Action buttons with keyboard shortcuts |
| TuiModal | - | Overlay dialog with double border and backdrop |
| TuiRadio | - | Radio button group for single selection |
| TuiCheckbox | - | Checkbox with label for multi-selection |
| TuiTabs | - | Tab navigation for configuration sections |
| TuiStatusBar | - | Bottom bar with context info and hints |
| TuiTree | - | Tree view for file/component hierarchies |

#### Logging Strategy (charmbracelet/log)

The application uses `charmbracelet/log` for:
- **Application logging**: Structured logs with levels (Debug, Info, Warn, Error)
- **Colorized output**: Lip Gloss styling integration for consistent look
- **Log parsing**: Parse Maven output and display with level-based coloring in viewport
- **JSON/text formatters**: For log export and debugging

### Architecture Principles

1. **Bubble Tea Elm Architecture**: All state managed through Model-Update-View cycle
2. **Component Composition**: Complex UIs built from composable primitives
3. **Message Passing**: Components communicate via typed messages
4. **Immutable Updates**: State changes return new state, never mutate
5. **Pure View Functions**: View functions only render, no side effects

### ⚠️ CRITICAL: Charm.land First - Custom Last

**This is a MANDATORY rule for all development agents.**

Before implementing ANY visual/UI element, agents MUST check if charm.land already provides it:

#### Priority Order (MUST follow)

1. **Bubbles Component** → Use directly (list, table, viewport, progress, spinner, help, textinput, paginator)
2. **Lip Gloss Styling** → Use for ALL styling needs:
   - Borders: `lipgloss.Border()` with `NormalBorder()`, `RoundedBorder()`, `DoubleBorder()`, `ThickBorder()`, `HiddenBorder()`
   - Colors: `lipgloss.Color()` and `lipgloss.AdaptiveColor()`
   - Layout: `lipgloss.Place()`, `lipgloss.JoinHorizontal()`, `lipgloss.JoinVertical()`
   - Padding/Margin: `.Padding()`, `.Margin()`
   - Text styling: `.Bold()`, `.Italic()`, `.Underline()`, `.Faint()`
3. **charmbracelet/log** → Use for all logging output styling
4. **Custom implementation** → ONLY when charm.land has NO solution

#### Forbidden Patterns (DO NOT implement)

| ❌ DO NOT | ✅ USE INSTEAD |
|-----------|----------------|
| Custom border drawing with `─│┌┐└┘` | `lipgloss.NewStyle().Border(lipgloss.NormalBorder())` |
| Custom color codes/ANSI escapes | `lipgloss.Color("#ff0000")` or `lipgloss.AdaptiveColor{}` |
| Manual string padding | `lipgloss.NewStyle().Padding(1, 2)` |
| Custom progress bar strings | `bubbles/progress` component |
| Custom spinner frames | `bubbles/spinner` component |
| Custom list rendering | `bubbles/list` component |
| Custom table rendering | `bubbles/table` component |
| Custom scrolling logic | `bubbles/viewport` component |

#### When Custom IS Allowed

Custom implementation is permitted ONLY for:
- **Business logic components** (TuiStatus badges, TuiNavItem with shortcuts)
- **Composite components** (TuiModal = viewport + Lip Gloss border + overlay logic)
- **Domain-specific rendering** (Maven build phases, Git status indicators)

Even then, the custom component MUST use Lip Gloss for ALL styling internally.

---

## 6. Development Approach

### Phase 1: Component Library + Visual Testing Foundation
**Timeline:** Weeks 1-3
**Goal:** Build embedded TUI component library with structural match to web prototype

**Deliverables:**
- [ ] All design system components implemented (TuiBox, TuiList, etc.)
- [ ] teatest infrastructure setup
- [ ] Screenshot testing with hybrid approach (structural + visual)
- [ ] Mock domain logic for all external dependencies
- [ ] Demo components from references/_test-data integrated for testing
- [ ] Component showcase/gallery screen for visual verification

**Success Criteria:**
- All components render correctly at 120x40
- Golden file tests passing for all component variants
- No external dependencies (Maven, Git) required for tests

### Phase 2: Screen Implementation with Visual Tests
**Timeline:** Weeks 4-7
**Goal:** Implement all 5 main screens + 2 modals matching web prototype

**Deliverables:**
- [ ] Welcome/Home screen
- [ ] Build Components screen
- [ ] Build Configuration modal
- [ ] Build Execution view
- [ ] Log Viewer screen
- [ ] Discover screen
- [ ] Configuration screen
- [ ] Visual tests for all 97 screenshots from prototype

**Success Criteria:**
- Visual parity with web prototype screenshots
- All keyboard navigation working
- Golden file tests for all 97 UI states
- Still using mocked domain logic

### Phase 3: Real Functionality
**Timeline:** Weeks 8-12
**Goal:** Connect real Maven execution, Git operations, and component scanning

**Deliverables:**
- [ ] Maven execution (clean, install, package with profiles)
- [ ] Git operations (status, checkout, pull)
- [ ] Component scanning with registry matching
- [ ] Configuration persistence
- [ ] Full integration testing
- [ ] Error handling and edge cases

**Success Criteria:**
- Successful builds of real RFZ components
- Accurate Git status detection
- Reliable component discovery
- Integration tests with real (or realistically mocked) components

---

## 7. User Experience

### Navigation Model

```
+-------------------+
|    Navigation     |  [1] Build  [2] Logs  [3] Discover  [4] Config
+-------------------+
|                   |
|   Main Content    |  Screen-specific content
|      Area         |
|                   |
+-------------------+
|    Status Bar     |  Context info, keyboard hints
+-------------------+
```

### Keyboard Shortcuts

| Context | Key | Action |
|---------|-----|--------|
| Global | 1-4 | Navigate to screen |
| Global | ? | Show help |
| Global | q | Quit (with confirmation) |
| Global | Esc | Close modal / Cancel |
| Lists | j/k or arrows | Navigate up/down |
| Lists | Space | Toggle selection |
| Lists | Enter | Confirm / Open |
| Build | b | Start build |
| Build | c | Configure build |
| Logs | f | Toggle follow mode |
| Logs | / | Search |
| Logs | 1-5 | Set filter level |

### Visual Design Principles

1. **Information Density**: Pack useful info without clutter
2. **Visual Hierarchy**: Clear distinction between primary and secondary content
3. **Status Visibility**: Always show what's happening (build status, Git status)
4. **Keyboard Discoverability**: Shortcuts visible in UI
5. **Consistent Patterns**: Same interactions work the same everywhere

---

## 8. Success Metrics

### Primary Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| Build Initiation Time | < 10 seconds | Time from app launch to build start |
| UI Regression Rate | 0% | Regressions caught before release |
| Visual Test Coverage | 100% (97 states) | Screenshots with golden file tests |
| Team Adoption | 80% of RFZ team | Developers using tool weekly |

### Secondary Metrics

| Metric | Target | Description |
|--------|--------|-------------|
| Build Success Visibility | 100% | All build outcomes clearly displayed |
| Component Discovery Accuracy | 100% | All components correctly identified |
| Configuration Errors | 0 | No builds fail due to tool misconfiguration |

---

## 9. Competitive Analysis / Differentiation

### What Makes RFZ Developer CLI Different

| Differentiator | Description |
|----------------|-------------|
| **Purpose-Built for RFZ** | Deep understanding of component types, profiles, and workflows specific to RFZ |
| **Component Type Awareness** | Knows difference between Core, Simulator, and Standalone; applies correct defaults |
| **Visual Regression Testing** | Testing is a first-class feature, not afterthought; prevents AI-assisted development regressions |
| **Keyboard-First Design** | Optimized for developers who prefer terminal; minimal mouse required |
| **Build Progress Visibility** | Real-time phase tracking (Compiling -> Testing -> Packaging) not just pass/fail |

### Comparison to Alternatives

| Feature | RFZ CLI | Shell Scripts | IDE Maven Plugin |
|---------|---------|---------------|------------------|
| Component Selection | Visual multi-select | Manual typing | Project-based |
| Profile Awareness | Component-specific | Manual | Generic |
| Build Progress | Real-time phases | Output only | Basic |
| Git Integration | Built-in | Separate tool | Separate view |
| Keyboard Driven | Yes | Yes | No |
| Visual Consistency | Tested | N/A | IDE-dependent |

---

## 10. Risks and Mitigations

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Terminal compatibility issues | Medium | High | Test on target terminals early; use widely-supported ANSI codes |
| Component ecosystem changes | Low | Medium | Registry-based design allows updates without code changes |
| Visual test brittleness | Medium | Medium | Structural tests as primary, visual as secondary |
| Performance with many components | Low | Medium | Lazy loading, virtual scrolling for large lists |
| Maven version incompatibilities | Low | High | Abstract Maven execution; test with multiple versions |

---

## 11. Reference Materials

### Project Structure

```
references/
├── prototype-screenshots/       # 97 UI screenshots from web prototype
│   ├── 01-welcome-default.png
│   ├── 10-build-*.png
│   ├── 20-config-*.png
│   └── ...
├── tui-web-prototype/          # Visual design prototype (v0.dev)
├── user-flow-diagrams.md       # User flow documentation
└── _test-data/
    ├── TREE.txt                # Directory structure reference
    ├── WORKFLOW.md             # Workflow documentation
    └── demo-components/        # Test data for development
        ├── core/
        ├── simulator/
        └── standalone/
```

### Key Documents

- **Prototype Screenshots**: Visual reference for all UI states
- **Web Prototype**: Interactive design reference (UI/UX only, not code)
- **Demo Components**: Mock component data for testing
- **User Flow Diagrams**: Navigation and interaction flows

---

## 12. Glossary

| Term | Definition |
|------|------------|
| **RFZ** | The railway control system project at Deutsche Bahn |
| **Core Component** | Main application components in the primary repository |
| **Simulator** | Simulation components for testing railway scenarios |
| **Standalone** | Components maintained in separate repositories |
| **Profile** | Maven build profile (e.g., local-config, skip-tests) |
| **Golden File** | Reference screenshot for visual regression testing |
| **Bubble Tea** | Go framework for terminal user interfaces from charm.land |

---

## Appendix A: Screen Inventory

| # | Screen | Type | Description |
|---|--------|------|-------------|
| 1 | Welcome | Main | Landing screen with quick actions |
| 2 | Build Components | Main | Component selection for building |
| 3 | Build Configuration | Modal | Build options before execution |
| 4 | Build Execution | Main | Active build monitoring |
| 5 | Log Viewer | Main | Component log viewing |
| 6 | Discover | Main | Repository scanning and status |
| 7 | Configuration | Main | App settings management |

## Appendix B: Component Categories

| Category | Description | Examples |
|----------|-------------|----------|
| Core | Main application modules | rfz-core, rfz-api, rfz-persistence |
| Simulator | Railway simulation components | train-sim, signal-sim, track-sim |
| Standalone | Separate repository components | rfz-tools, rfz-config-generator |

---

**Note:** This product brief serves as the foundation for all development decisions. Refer to this document when planning features, making architectural choices, or prioritizing work. For implementation details, see the associated roadmap and technical specifications.
