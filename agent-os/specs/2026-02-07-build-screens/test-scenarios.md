# Test Scenarios - Build Screens (Sprint 2.2)

**Spec:** 2026-02-07-build-screens
**Generated:** 2026-02-07

---

## 1. Domain Model (BUILD-001)

### Happy Path

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 1.1 | Mock provider returns components | Init MockComponentProvider, call Components() | 13 components with name and type |
| 1.2 | BuildConfig generates command | Set goal=clean install, profile=generate_local_config_files, port=11090, skipTests=true | `mvn clean install -Pgenerate_local_config_files,use_traktion_11090 -DskipTests` |
| 1.3 | Build phases are complete | Check all BuildPhase constants | Pending, Compiling, Testing, Packaging, Installing, Done, Failed |

### Edge Cases

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 1.4 | Config without profiles | Set goal=install only | `mvn install` (no -P flag) |
| 1.5 | Config with port only | Set goal=compile, port=8080 | `mvn compile -Puse_traktion_8080` |
| 1.6 | Unknown component type | Create ComponentType(99) | String() returns "Unknown" |

---

## 2. Component Selection (BUILD-002)

### Happy Path

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 2.1 | All components displayed | Navigate to Build screen | 13 components visible with type badges, "0/13 selected" |
| 2.2 | Toggle single component | Press Space on "boss" | "boss" marked [x], counter "1/13 selected" |
| 2.3 | Select all | Press "a" | All 13 selected, counter "13/13 selected" |
| 2.4 | Deselect all | With 3 selected, press "n" | None selected, counter "0/13 selected" |
| 2.5 | Open config after selection | Select 3 components, press Enter | Config modal opens with "Building 3 components: ..." |
| 2.6 | Navigate with arrow keys | Press down arrow from "boss" | Cursor moves to "fistiv" |

### Edge Cases

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 2.7 | Enter without selection | No components selected, press Enter | Nothing happens, stays on selection |
| 2.8 | Wrap-around navigation | At last item, press down | Cursor wraps to first item |

---

## 3. Configuration Modal (BUILD-003)

### Happy Path

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 3.1 | Default values shown | Open config modal | "clean install" selected, "target_env_dev" active, port 11090, skip tests on |
| 3.2 | Change Maven goal | Select "package" in goal section | Command preview updates to "mvn package ..." |
| 3.3 | Toggle profile | Activate "generate_local_config_files" | Both profiles active, preview updated |
| 3.4 | Change port | Select port 11091 | Preview shows "use_traktion_11091" |
| 3.5 | Toggle skip tests | Deactivate skip tests | "-DskipTests" removed from preview |
| 3.6 | Tab between sections | Press Tab from Maven Goal | Focus moves to Maven Profiles section |
| 3.7 | Start build | Press Enter on Start Build button | Modal closes, execution view shown |

### Edge Cases

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 3.8 | Cancel returns to selection | Press Esc in modal | Modal closes, selection preserved |
| 3.9 | Shift+Tab backwards | Press Shift+Tab from Maven Profiles | Focus moves back to Maven Goal |

---

## 4. Build Execution (BUILD-004)

### Happy Path

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 4.1 | Build starts | Start build with 3 components | Maven command visible, 3 rows in table, counters show pending |
| 4.2 | Phase progression | Watch simulation | Components progress: Pending -> Compiling -> Testing -> Packaging -> Installing -> Done |
| 4.3 | Overall progress updates | 1 of 3 done | Progress bar ~33%, counters updated |
| 4.4 | All complete | All 3 finish | Progress 100%, "New Build" button shown |
| 4.5 | Navigate between rows | Press down arrow | Cursor moves to next component row |

### Edge Cases

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 4.6 | Component fails | Random failure in Testing phase | Status "Failed" in red, others continue |
| 4.7 | Cancel running build | Press Esc during execution | All running/pending marked Failed, phase=completed |
| 4.8 | New build after completion | Press "n" after build done | Returns to selection, previous selection cleared |
| 4.9 | Concurrent build limit | Start 5+ components | Max 3 build simultaneously, others wait as Pending |

---

## 5. App Integration (BUILD-005)

### Happy Path

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 5.1 | Navigate via shortcut | Press "1" from Welcome | Build screen shown, nav active, status "SELECT" |
| 5.2 | Status bar reflects state | Enter config modal | Status shows "CONFIG" mode |
| 5.3 | Tab focus toggle | Press Tab from content | Focus switches to navigation sidebar |
| 5.4 | Escape to welcome | Press Esc from build | Welcome screen shown |

### Edge Cases

| # | Scenario | Steps | Expected |
|---|----------|-------|----------|
| 5.5 | Quit modal from build | Press "q" in content focus | Quit confirmation modal shown |
| 5.6 | Golden file regression | Run all TestApp_Build* tests | All match golden files exactly |
| 5.7 | Window resize | Resize terminal during build | All views adapt, no clipping |

---

## End-to-End Scenarios

### E2E-1: Full Build Workflow

1. Start app -> Welcome screen
2. Press "1" -> Build Components screen
3. Press Tab -> Content focused
4. Press "a" -> All 13 selected
5. Press Enter -> Config modal opens
6. Verify defaults -> Press Enter to start build
7. Watch execution -> Components progress through phases
8. Build completes -> "New Build" button shown
9. Press "n" -> Back to selection, all deselected

### E2E-2: Build with Cancel

1. Start app -> Press "1" -> Tab -> Select 3 components
2. Enter -> Config modal -> Start Build
3. During execution, press Esc -> All stopped
4. Press "n" -> Back to selection
