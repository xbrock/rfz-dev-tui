# Product Roadmap: RFZ Developer CLI

> Last Updated: 2026-02-02
> Version: 1.0.0
> Status: Planning

---

## Overview

This roadmap organizes development into three distinct phases, each with clear goals and success criteria. The approach emphasizes building a solid foundation (component library + visual testing) before implementing screens, and only connecting real functionality in the final phase.

### Timeline Summary

| Phase | Name | Duration | Focus |
|-------|------|----------|-------|
| 1 | Foundation | 3 weeks | Component Library + Visual Testing Infrastructure |
| 2 | Screens | 4 weeks | All Screens + Visual Test Coverage |
| 3 | Integration | 4-5 weeks | Real Maven/Git Functionality |

---

## MoSCoW Prioritization

### Must Have (MVP - Phases 1-2)
- [ ] TUI Component Library (all design system components)
- [ ] Visual testing infrastructure with teatest
- [ ] Golden file testing for all component variants
- [ ] Build Components screen with multi-select
- [ ] Build Configuration modal
- [ ] Build Execution view with status tracking
- [ ] Log Viewer with filtering
- [ ] Discover screen with component listing
- [ ] Configuration screen for scan paths
- [ ] Keyboard navigation throughout
- [ ] Visual tests for all 97 UI states

### Should Have (Phase 3)
- [ ] Real Maven execution (clean, install, package)
- [ ] Profile-aware build configuration
- [ ] Git status detection (clean/dirty)
- [ ] Component scanning with registry matching
- [ ] Configuration persistence
- [ ] Error handling and recovery
- [ ] Build output streaming

### Could Have (Phase 3+)
- [ ] Git operations (checkout, pull)
- [ ] Build history and statistics
- [ ] Component dependency visualization
- [ ] Custom keyboard shortcut configuration
- [ ] Theme customization
- [ ] Export/import configuration

### Won't Have (Future)
- [ ] Remote build execution
- [ ] Multi-user collaboration features
- [ ] Web-based dashboard
- [ ] Plugin system
- [ ] Notification integrations (Slack, email)

---

## Phase 1: Foundation (Weeks 1-3)

**Goal:** Build embedded TUI component library following idiomatic Bubble Tea patterns with structural match to web prototype screenshots.

**Success Criteria:**
- All design system components implemented and tested
- teatest infrastructure operational
- Golden file tests passing for all component variants
- Mock domain logic in place for all external dependencies
- Demo components integrated for testing
- No external dependencies (Maven, Git) required for tests

### Sprint 1.1: Core Components (Week 1)

| Task | Priority | Component | Description |
|------|----------|-----------|-------------|
| T1.1.1 | Must | TuiBox | Container with border styles (single, double, rounded, heavy, focused) |
| T1.1.2 | Must | TuiDivider | Horizontal/vertical separators (single, double) |
| T1.1.3 | Must | TuiButton | Button variants (default, primary, danger) with states |
| T1.1.4 | Must | TuiStatus | Status badges (pending, running, success, failed, error) |
| T1.1.5 | Must | Test Infra | teatest setup with golden file management |

**Deliverables:**
- [ ] Core container components
- [ ] Button component with all variants
- [ ] Status badge component
- [ ] teatest harness with screenshot capture
- [ ] Golden file comparison tooling

### Sprint 1.2: Interactive Components (Week 2)

| Task | Priority | Component | Description |
|------|----------|-----------|-------------|
| T1.2.1 | Must | TuiList | Scrollable list with single/multi-select |
| T1.2.2 | Must | TuiCheckbox | Checkbox with label |
| T1.2.3 | Must | TuiRadio | Radio button group |
| T1.2.4 | Must | TuiTextInput | Text input field |
| T1.2.5 | Must | TuiSpinner | Loading indicators (braille-dots, line, circle-quarters, bounce) |
| T1.2.6 | Must | TuiProgress | Progress bars (block, gradient, braille, ASCII) |

**Deliverables:**
- [ ] List component with selection modes
- [ ] Form input components (checkbox, radio, text)
- [ ] Animation components (spinner, progress)
- [ ] Golden file tests for all variants

### Sprint 1.3: Layout & Navigation (Week 3)

| Task | Priority | Component | Description |
|------|----------|-----------|-------------|
| T1.3.1 | Must | TuiNavigation | Navigation bar container |
| T1.3.2 | Must | TuiNavItem | Nav item with focus/active states |
| T1.3.3 | Must | TuiModal | Overlay dialog with double border |
| T1.3.4 | Must | TuiTable | Data table with selectable rows |
| T1.3.5 | Must | TuiTabs | Tab navigation component |
| T1.3.6 | Must | TuiStatusBar | Bottom status bar |
| T1.3.7 | Must | TuiTree | Tree view for hierarchies |
| T1.3.8 | Must | TuiKeyHints | Keyboard shortcut hints display |

**Deliverables:**
- [ ] Navigation components
- [ ] Modal overlay system
- [ ] Data display components (table, tree)
- [ ] Utility components (status bar, key hints)
- [ ] Component showcase/gallery screen
- [ ] Complete golden file coverage for all components

---

## Phase 2: Screen Implementation (Weeks 4-7)

**Goal:** Implement all 5 main screens + 2 modals matching web prototype with visual tests for all 97 UI states.

**Success Criteria:**
- Visual parity with prototype screenshots
- All keyboard navigation working
- Golden file tests for all 97 UI states
- Screen transitions smooth and correct
- Still using mocked domain logic

### Sprint 2.1: Welcome & Navigation (Week 4)

| Task | Priority | Screen | Description |
|------|----------|--------|-------------|
| T2.1.1 | Must | App Shell | Main application frame with navigation |
| T2.1.2 | Must | Welcome | Landing screen with quick actions |
| T2.1.3 | Must | Navigation | Tab-based screen switching (1-4 keys) |
| T2.1.4 | Must | Status Bar | Global status bar implementation |

**Screenshot Coverage:** 01-welcome-*.png series

**Deliverables:**
- [ ] Application shell with navigation
- [ ] Welcome/home screen
- [ ] Global keyboard navigation
- [ ] Visual tests for welcome states

### Sprint 2.2: Build Screens (Week 5)

| Task | Priority | Screen | Description |
|------|----------|--------|-------------|
| T2.2.1 | Must | Build Components | Component list with multi-select and badges |
| T2.2.2 | Must | Build Config Modal | Configuration options dialog |
| T2.2.3 | Must | Build Execution | Progress view with status tracking |

**Screenshot Coverage:** 10-build-*.png series

**Deliverables:**
- [ ] Build component selection screen
- [ ] Build configuration modal
- [ ] Build execution progress view
- [ ] Visual tests for all build states

### Sprint 2.3: Logs & Discover (Week 6)

| Task | Priority | Screen | Description |
|------|----------|--------|-------------|
| T2.3.1 | Must | Log Viewer | Log display with filtering |
| T2.3.2 | Must | Log Filters | Filter level selection (ALL/INFO/WARN/ERROR/DEBUG) |
| T2.3.3 | Must | Discover | Repository/component scanning view |
| T2.3.4 | Must | Component Status | Git status display per component |

**Screenshot Coverage:** 30-logs-*.png, 40-discover-*.png series

**Deliverables:**
- [ ] Log viewer with filtering
- [ ] Follow mode for logs
- [ ] Discover screen with component list
- [ ] Git status indicators
- [ ] Visual tests for logs and discover states

### Sprint 2.4: Configuration & Polish (Week 7)

| Task | Priority | Screen | Description |
|------|----------|--------|-------------|
| T2.4.1 | Must | Configuration | Settings management screen |
| T2.4.2 | Must | Scan Paths | Path management UI |
| T2.4.3 | Must | Component Registry | Registry view |
| T2.4.4 | Must | Polish | Edge cases, error states, transitions |
| T2.4.5 | Must | Full Coverage | Remaining visual tests to reach 97 |

**Screenshot Coverage:** 20-config-*.png series + remaining states

**Deliverables:**
- [ ] Configuration screen
- [ ] Scan paths management
- [ ] Component registry view
- [ ] All 97 visual tests passing
- [ ] Complete UI polish pass

---

## Phase 3: Real Functionality (Weeks 8-12)

**Goal:** Connect real Maven execution, Git operations, and component scanning to transform the mocked UI into a fully functional tool.

**Success Criteria:**
- Successful builds of real RFZ components
- Accurate Git status detection
- Reliable component discovery
- Integration tests with real/realistic components
- Error handling for all failure modes

### Sprint 3.1: Maven Integration (Weeks 8-9)

| Task | Priority | Feature | Description |
|------|----------|---------|-------------|
| T3.1.1 | Should | Maven Executor | Execute Maven commands with profiles |
| T3.1.2 | Should | Output Streaming | Real-time Maven output capture |
| T3.1.3 | Should | Status Detection | Parse Maven output for phase detection |
| T3.1.4 | Should | Error Handling | Handle Maven failures gracefully |
| T3.1.5 | Should | Profile Mapping | Map component types to available profiles |

**Deliverables:**
- [ ] Maven command execution
- [ ] Real-time output streaming to log viewer
- [ ] Build phase detection from Maven output
- [ ] Profile-aware configuration
- [ ] Integration tests with demo components

### Sprint 3.2: Git & Discovery (Weeks 10-11)

| Task | Priority | Feature | Description |
|------|----------|---------|-------------|
| T3.2.1 | Should | Git Status | Detect clean/dirty state per repo |
| T3.2.2 | Should | Branch Info | Get current branch name |
| T3.2.3 | Should | Commit Info | Get last commit hash and message |
| T3.2.4 | Should | Component Scanner | Scan directories for pom.xml |
| T3.2.5 | Should | Registry Matching | Match scanned components to registry |
| T3.2.6 | Could | Git Operations | Checkout, pull operations |

**Deliverables:**
- [ ] Git status detection
- [ ] Branch and commit information
- [ ] Directory scanning for components
- [ ] Component type detection
- [ ] Registry-based component identification

### Sprint 3.3: Configuration & Polish (Week 12)

| Task | Priority | Feature | Description |
|------|----------|---------|-------------|
| T3.3.1 | Should | Config Persistence | Save/load configuration |
| T3.3.2 | Should | Path Management | Add/remove scan paths |
| T3.3.3 | Should | Error Recovery | Graceful handling of all error conditions |
| T3.3.4 | Should | Edge Cases | Handle missing components, network issues |
| T3.3.5 | Should | Documentation | User guide and help system |

**Deliverables:**
- [ ] Configuration file persistence
- [ ] Complete error handling
- [ ] Help documentation
- [ ] Final integration testing
- [ ] Release candidate

---

## Future Considerations (Beyond Phase 3)

### Potential Features for v2.0

| Feature | Value | Complexity |
|---------|-------|------------|
| Build History | Track past builds with statistics | Medium |
| Dependency Graph | Visualize component dependencies | High |
| Custom Shortcuts | User-configurable keyboard shortcuts | Low |
| Themes | Light/dark mode, custom color schemes | Medium |
| Build Profiles | Save/load build configurations | Low |
| Parallel Builds | Build multiple components simultaneously | High |

### Integration Opportunities

| Integration | Description |
|-------------|-------------|
| CI/CD | Trigger Jenkins/GitHub Actions builds |
| Notifications | Slack/Teams notifications on build complete |
| IDE | VS Code extension for quick access |
| Metrics | Build time analytics dashboard |

### Technical Debt to Address

| Item | Description |
|------|-------------|
| Test Coverage | Increase unit test coverage beyond visual tests |
| Performance | Optimize for large component counts |
| Accessibility | Screen reader support |
| Documentation | Comprehensive API docs for component library |

---

## Risk Register

| Risk | Likelihood | Impact | Mitigation | Owner |
|------|------------|--------|------------|-------|
| Terminal compatibility | Medium | High | Test early on target terminals | Phase 1 |
| Visual test brittleness | Medium | Medium | Structural tests primary, visual secondary | Phase 1 |
| Maven version issues | Low | High | Abstract execution, test multiple versions | Phase 3 |
| Scope creep | Medium | Medium | Strict MoSCoW adherence | All phases |
| Integration complexity | Medium | Medium | Incremental integration, good abstractions | Phase 3 |

---

## Milestones

| Milestone | Target Date | Criteria |
|-----------|-------------|----------|
| M1: Component Library Complete | Week 3 | All components with golden tests |
| M2: UI Feature Complete | Week 7 | All 97 visual tests passing |
| M3: MVP Release | Week 12 | Real functionality working |
| M4: Team Adoption | Week 16 | 80% team using weekly |

---

**Note:** This roadmap is a living document. Update priorities based on user feedback, technical discoveries, and stakeholder input. Create detailed specs for features when they're ready for development using the `/create-spec` workflow.
