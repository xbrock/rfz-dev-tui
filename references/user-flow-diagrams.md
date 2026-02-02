# RFZ-CLI User Flow Diagrams

> Mermaid diagrams for pitch presentation
> Version: 1.0.0

---

## 1. Overview: Main Navigation Flow

```mermaid
flowchart TB
    subgraph entry["🚀 Entry"]
        START([Launch App]) --> WELCOME[Welcome Screen]
    end

    subgraph nav["📍 Main Navigation"]
        WELCOME --> NAV{Navigation Panel}
        NAV -->|"1"| BUILD[Build Components]
        NAV -->|"2"| LOGS[Log Viewer]
        NAV -->|"3"| DISCOVER[Discover]
        NAV -->|"4"| CONFIG[Configuration]
        NAV -->|"q"| EXIT([Exit App])
    end

    subgraph build_flow["🔨 Build Flow"]
        BUILD --> SELECT[Select Components]
        SELECT -->|"Enter"| MODAL[Build Config Modal]
        MODAL -->|"Start Build"| EXEC[Build Execution]
        EXEC -->|"Complete"| RESULT{Result}
        RESULT -->|"Success"| LOGS
        RESULT -->|"Failed"| RETRY[Rebuild Failed]
        RETRY --> EXEC
    end

    subgraph monitoring["📊 Monitoring"]
        LOGS --> FILTER[Filter & Search]
        LOGS --> FOLLOW[Follow Mode]
        EXEC -->|"L"| LOG_MODAL[Logs Modal]
        LOG_MODAL -->|"Esc"| EXEC
    end

    subgraph discovery["🔍 Discovery"]
        DISCOVER --> SCAN[Scan Repositories]
        SCAN --> STATUS[View Status]
        STATUS --> ACTIONS[Git Actions]
    end

    subgraph settings["⚙️ Settings"]
        CONFIG --> PATHS[Scan Paths]
        CONFIG --> REGISTRY[Component Registry]
        CONFIG --> DETECTED[Detected Components]
    end

    %% Cross-navigation
    BUILD -.->|"Esc"| NAV
    LOGS -.->|"Esc"| NAV
    DISCOVER -.->|"Esc"| NAV
    CONFIG -.->|"Esc"| NAV

    style WELCOME fill:#1e40af,color:#fff
    style BUILD fill:#059669,color:#fff
    style LOGS fill:#7c3aed,color:#fff
    style DISCOVER fill:#d97706,color:#fff
    style CONFIG fill:#dc2626,color:#fff
    style EXEC fill:#0891b2,color:#fff
```

---

## 2. Detailed: Build Workflow

```mermaid
flowchart TB
    subgraph selection["Component Selection"]
        B_START([Enter Build Screen]) --> B_LIST[Component List]
        B_LIST --> B_NAV{Navigate}
        B_NAV -->|"↑/k"| B_UP[Move Up]
        B_NAV -->|"↓/j"| B_DOWN[Move Down]
        B_UP --> B_LIST
        B_DOWN --> B_LIST

        B_LIST --> B_SELECT{Select Action}
        B_SELECT -->|"Space"| B_TOGGLE[Toggle Selection]
        B_SELECT -->|"a"| B_ALL[Select All]
        B_SELECT -->|"n"| B_NONE[Clear Selection]
        B_TOGGLE --> B_LIST
        B_ALL --> B_LIST
        B_NONE --> B_LIST
    end

    subgraph config["Build Configuration"]
        B_LIST -->|"Enter"| C_MODAL[Config Modal Opens]
        C_MODAL --> C_GOALS[Maven Goals]
        C_GOALS --> C_PROFILES[Profiles Selection]
        C_PROFILES --> C_PORT[Port Selection]
        C_PORT --> C_TESTS[Skip Tests Toggle]
        C_TESTS --> C_PREVIEW[Command Preview]
        C_PREVIEW --> C_ACTION{Action}
        C_ACTION -->|"Cancel"| B_LIST
        C_ACTION -->|"Start Build"| E_START
    end

    subgraph execution["Build Execution"]
        E_START([Build Starts]) --> E_QUEUE[Queue Components]
        E_QUEUE --> E_LOOP{For Each Component}

        E_LOOP --> E_PENDING[○ Pending]
        E_PENDING --> E_COMPILE[◐ Compiling]
        E_COMPILE --> E_TEST[◐ Testing]
        E_TEST --> E_PACKAGE[◐ Packaging]
        E_PACKAGE --> E_INSTALL[◐ Installing]
        E_INSTALL --> E_CHECK{Check Result}

        E_CHECK -->|"Success"| E_DONE[✓ Done]
        E_CHECK -->|"Failure"| E_FAILED[✗ Failed]

        E_DONE --> E_NEXT{More Components?}
        E_FAILED --> E_NEXT
        E_NEXT -->|"Yes"| E_LOOP
        E_NEXT -->|"No"| E_COMPLETE[Build Complete]
    end

    subgraph actions["Post-Build Actions"]
        E_COMPLETE --> A_MENU{Actions Menu}
        A_MENU -->|"L"| A_LOGS[View Build Logs]
        A_MENU -->|"Rebuild"| E_START
        A_MENU -->|"Back"| B_LIST
        A_LOGS -->|"Esc"| A_MENU
    end

    style B_START fill:#059669,color:#fff
    style E_START fill:#0891b2,color:#fff
    style E_DONE fill:#22c55e,color:#fff
    style E_FAILED fill:#ef4444,color:#fff
    style E_COMPLETE fill:#1e40af,color:#fff
```

---

## 3. Detailed: Log Viewer Flow

```mermaid
flowchart TB
    subgraph entry["Entry Points"]
        L_NAV([From Navigation]) --> L_MAIN
        L_BUILD([From Build Exec]) --> L_MODAL
    end

    subgraph main["Log Viewer Screen"]
        L_MAIN[Log Viewer] --> L_FOCUS{Focus Area}
        L_FOCUS -->|"Tab"| L_COMP[Components Panel]
        L_FOCUS -->|"Tab"| L_LOGS[Log Content Panel]
        L_FOCUS -->|"Tab"| L_FILTERS[Filter Panel]

        L_COMP --> L_SELECT[Select Component]
        L_SELECT --> L_LOAD[Load Logs]
        L_LOAD --> L_LOGS
    end

    subgraph filtering["Filtering & Search"]
        L_FILTERS --> F_LEVEL{Filter Level}
        F_LEVEL -->|"f"| F_ALL[ALL]
        F_LEVEL -->|"f"| F_INFO[INFO]
        F_LEVEL -->|"f"| F_WARN[WARN]
        F_LEVEL -->|"f"| F_ERROR[ERROR]
        F_LEVEL -->|"f"| F_DEBUG[DEBUG]

        L_FILTERS --> F_SEARCH[Search Query]
        F_SEARCH --> F_APPLY[Apply Filter]
        F_APPLY --> L_LOGS
    end

    subgraph viewing["Log Viewing"]
        L_LOGS --> V_SCROLL{Scroll}
        V_SCROLL -->|"↑/k"| V_UP[Scroll Up]
        V_SCROLL -->|"↓/j"| V_DOWN[Scroll Down]
        V_SCROLL -->|"Space"| V_FOLLOW[Toggle Follow Mode]

        V_FOLLOW --> V_AUTO{Auto-Scroll}
        V_AUTO -->|"On"| V_TAIL[Tail New Logs]
        V_AUTO -->|"Off"| V_MANUAL[Manual Navigation]
    end

    subgraph modal["Logs Modal (from Build)"]
        L_MODAL[Build Log Modal] --> M_VIEW[View Build Output]
        M_VIEW --> M_TOGGLE{Toggle}
        M_TOGGLE -->|"e"| M_ERRORS[Error-Only Mode]
        M_TOGGLE -->|"e"| M_ALL_LOGS[All Logs]
        M_VIEW -->|"Esc"| L_CLOSE([Return to Build])
    end

    style L_MAIN fill:#7c3aed,color:#fff
    style L_MODAL fill:#7c3aed,color:#fff
    style V_FOLLOW fill:#0891b2,color:#fff
```

---

## 4. Detailed: Discover Flow

```mermaid
flowchart TB
    subgraph scan["Repository Scanning"]
        D_START([Enter Discover]) --> D_SCAN[Scan Configured Paths]
        D_SCAN --> D_FIND[Find RFZ Components]
        D_FIND --> D_GIT[Check Git Status]
        D_GIT --> D_LIST[Component List]
    end

    subgraph display["Component Display"]
        D_LIST --> D_INFO{Component Info}
        D_INFO --> I_NAME[Component Name]
        D_INFO --> I_CAT[Category Badge]
        D_INFO --> I_BRANCH[Git Branch]
        D_INFO --> I_STATUS[Git Status]
        D_INFO --> I_COMMIT[Last Commit]

        I_STATUS --> S_CLEAN[🟢 Clean]
        I_STATUS --> S_DIRTY[🟡 Dirty]
        I_STATUS --> S_UNKNOWN[⚪ Unknown]
    end

    subgraph navigation["Navigation"]
        D_LIST --> D_NAV{Navigate}
        D_NAV -->|"↑/k"| D_PREV[Previous Component]
        D_NAV -->|"↓/j"| D_NEXT[Next Component]
        D_PREV --> D_LIST
        D_NEXT --> D_LIST
    end

    subgraph actions["Git Actions"]
        D_LIST -->|"Tab"| D_ACTIONS[Actions Panel]
        D_ACTIONS --> A_CLONE[Clone Repository]
        D_ACTIONS --> A_CHECKOUT[Checkout Branch]
        D_ACTIONS --> A_STATUS[Refresh Status]
        D_ACTIONS --> A_PULL[Pull Updates]
    end

    subgraph summary["Summary Stats"]
        D_LIST --> D_STATS[Statistics Bar]
        D_STATS --> ST_TOTAL[Total: N]
        D_STATS --> ST_CLEAN[Clean: N]
        D_STATS --> ST_DIRTY[Dirty: N]
        D_STATS --> ST_DEV[Not on develop: N]
    end

    style D_START fill:#d97706,color:#fff
    style S_CLEAN fill:#22c55e,color:#fff
    style S_DIRTY fill:#eab308,color:#000
    style S_UNKNOWN fill:#6b7280,color:#fff
```

---

## 5. Detailed: Configuration Flow

```mermaid
flowchart TB
    subgraph entry["Configuration Screen"]
        C_START([Enter Config]) --> C_SECTIONS{Section Tabs}
        C_SECTIONS -->|"Tab 1"| S_SCAN[Scan Configuration]
        C_SECTIONS -->|"Tab 2"| S_REG[Component Registry]
        C_SECTIONS -->|"Tab 3"| S_DET[Detected Components]
    end

    subgraph scan_config["Scan Configuration (Editable)"]
        S_SCAN --> SC_PATHS[Scan Paths List]
        SC_PATHS --> P_ADD[Add Path]
        SC_PATHS --> P_REMOVE[Remove Path]
        SC_PATHS --> P_TOGGLE[Toggle Enabled]
        SC_PATHS --> P_EDIT[Edit Path Value]

        S_SCAN --> SC_BEHAVIOR[Scan Behavior]
        SC_BEHAVIOR --> B_RECURSIVE[Recursive: On/Off]
        SC_BEHAVIOR --> B_DEPTH[Max Depth: 1-10]
        SC_BEHAVIOR --> B_EXCLUDE[Exclude Patterns]
    end

    subgraph registry["Component Registry (Read-Only)"]
        S_REG --> R_LIST[Registered Components]
        R_LIST --> R_DETAIL{Component Details}
        R_DETAIL --> RD_NAME[Name & Artifact ID]
        R_DETAIL --> RD_CAT[Category]
        R_DETAIL --> RD_PROF[Default Profiles]
        R_DETAIL --> RD_PORT[Supported Ports]
        R_DETAIL --> RD_ORDER[Build Order]
    end

    subgraph detected["Detected Components (Read-Only)"]
        S_DET --> D_LIST[Detection Results]
        D_LIST --> D_ITEM{Component Status}
        D_ITEM --> DI_FOUND[✓ Found + Path]
        D_ITEM --> DI_NOTFOUND[✗ Not Found]
        D_ITEM --> DI_CONFIG[Has Config Module]
    end

    subgraph navigation["Navigation"]
        C_SECTIONS --> N_TAB[Tab: Switch Sections]
        SC_PATHS --> N_ARROWS[↑↓: Navigate Items]
        R_LIST --> N_ARROWS
        D_LIST --> N_ARROWS
        N_ARROWS --> N_ENTER[Enter: Edit/View]
        N_ENTER --> N_ESC[Esc: Back]
    end

    style C_START fill:#dc2626,color:#fff
    style S_SCAN fill:#059669,color:#fff
    style S_REG fill:#6b7280,color:#fff
    style S_DET fill:#6b7280,color:#fff
```

---

## 6. State Machine: Application States

```mermaid
stateDiagram-v2
    [*] --> Welcome: Launch

    Welcome --> BuildScreen: Press 1
    Welcome --> LogViewer: Press 2
    Welcome --> Discover: Press 3
    Welcome --> Config: Press 4

    BuildScreen --> BuildConfig: Enter (with selection)
    BuildConfig --> BuildScreen: Cancel
    BuildConfig --> BuildExecution: Start Build

    BuildExecution --> LogsModal: Press L
    LogsModal --> BuildExecution: Esc
    BuildExecution --> BuildScreen: Back
    BuildExecution --> BuildExecution: Rebuild Failed

    BuildScreen --> Welcome: Esc
    LogViewer --> Welcome: Esc
    Discover --> Welcome: Esc
    Config --> Welcome: Esc

    Welcome --> [*]: Press q

    state BuildExecution {
        [*] --> Pending
        Pending --> Compiling
        Compiling --> Testing
        Testing --> Packaging
        Packaging --> Installing
        Installing --> Done
        Installing --> Failed
        Done --> [*]
        Failed --> [*]
    }
```

---

## 7. Component Architecture Overview

```mermaid
flowchart TB
    subgraph app["RFZ-CLI Application"]
        subgraph tui["TUI Layer"]
            BOX[TuiBox]
            LIST[TuiList]
            NAV[TuiNavigation]
            MODAL[TuiModal]
            BTN[TuiButton]
            PROG[TuiProgress]
            SPIN[TuiSpinner]
        end

        subgraph screens["Screen Components"]
            WELCOME_S[WelcomeScreen]
            BUILD_S[BuildScreen]
            EXEC_S[BuildExecutionView]
            LOGS_S[LogViewerScreen]
            DISC_S[DiscoverScreen]
            CONF_S[ConfigScreen]
        end

        subgraph modals["Modal Components"]
            CONFIG_M[BuildConfigModal]
            LOGS_M[LogsModal]
        end

        subgraph state["State Management"]
            APP_STATE[App State]
            BUILD_STATE[Build State]
            LOG_STATE[Log State]
            CONFIG_STATE[Config State]
        end

        subgraph services["Services"]
            MAVEN[Maven Executor]
            GIT[Git Operations]
            SCANNER[Component Scanner]
            LOGGER[Log Manager]
        end
    end

    subgraph external["External"]
        FS[File System]
        GIT_REPO[Git Repositories]
        MVN[Maven CLI]
    end

    %% Connections
    screens --> tui
    modals --> tui
    screens --> state
    modals --> state
    services --> external
    state --> services

    style app fill:#1e1e2e,color:#fff
    style tui fill:#313244,color:#fff
    style screens fill:#45475a,color:#fff
    style services fill:#585b70,color:#fff
```

---

## Usage Notes

### Rendering
These diagrams are written in Mermaid syntax and can be rendered in:
- GitHub README/Markdown files
- VS Code with Mermaid extension
- [Mermaid Live Editor](https://mermaid.live)
- Notion, Confluence, and other documentation tools
- Presentation tools that support Mermaid

### Color Legend
| Color | Meaning |
|-------|---------|
| 🟦 Blue (#1e40af) | Entry/Welcome |
| 🟩 Green (#059669) | Build/Success |
| 🟣 Purple (#7c3aed) | Logs/Monitoring |
| 🟧 Orange (#d97706) | Discover |
| 🟥 Red (#dc2626) | Config/Errors |
| 🔵 Cyan (#0891b2) | Execution/Active |

### Keyboard Shortcuts Reference
| Key | Action |
|-----|--------|
| `1-4` | Navigate to screen |
| `↑↓` / `jk` | Navigate list |
| `Enter` | Select/Confirm |
| `Space` | Toggle |
| `Tab` | Switch focus |
| `Esc` | Back/Cancel |
| `a` | Select all |
| `n` | Select none |
| `L` | View logs |
| `f` | Cycle filter |
| `q` | Quit |
