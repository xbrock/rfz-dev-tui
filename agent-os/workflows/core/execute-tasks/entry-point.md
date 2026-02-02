---
description: Entry point for task execution - routes to appropriate phase
globs:
alwaysApply: false
version: 3.6
encoding: UTF-8
---

# Task Execution Entry Point

## What's New in v3.6

**System Stories Support:**
- Detects System Stories (story-997, 998, 999) in specs
- Routes to Phase 3 when System Stories are pending (even if regular backlog is empty)
- Phase 4.5 and 5 become legacy checks (skip if System Stories exist and are Done)
- Backward compatible: Specs without System Stories work unchanged

**Requires create-spec v3.0:**
- System Stories (997, 998, 999) are automatically generated
- story-997: Code Review
- story-998: Integration Validation (replaces Phase 4.5)
- story-999: Finalize PR (replaces Phase 5)

## What's New in v3.5

**Integration Verification (Phase 3):**
- `verify_integration_requirements` step: Prüft VOR Implementierung welche Verbindungen nötig sind
- `self_review` erweitert: Verifiziert dass Verbindungen AKTIV hergestellt wurden (nicht nur Code existiert)
- FIX: "Komponenten gebaut aber nicht verbunden" - Erzwingt echte Integration per Story

**Requires create-spec v2.9:**
- Komponenten-Verbindungen Matrix im Implementation Plan
- Integration DoD items in Stories mit Verbindungs-Verantwortung

## What's New in v3.4

**Hybrid Template Lookup:**
- Templates searched in order: local (`agent-os/templates/`) → global (`~/.agent-os/templates/`)
- Fixes "template not found" for projects without local templates
- Applies to: kanban-board, test-scenarios, user-todos templates

**Handover Documentation (Spec only):**
- Phase 3: Collects user-todos during implementation (tasks requiring manual action)
- Phase 5: Generates test-scenarios.md with Happy-Path, Edge-Cases, Fehlerfälle
- Phase 5: Finalizes user-todos.md with summary and priority classification
- New templates: test-scenarios-template.md, user-todos-template.md

## What's New in v3.3

**External Worktree Support:**
- Worktrees are now located OUTSIDE the project: `../{project}-worktrees/{feature}`
- CWD check updated to support external worktree paths
- No symlinks needed - worktree contains full `.claude/` and `agent-os/` folders

## What's New in v3.2

**Worktree CWD Check:**
- Entry point now checks if agent is running in the correct working directory
- When Git Strategy is "worktree", validates CWD matches the Worktree Path
- Displays clear warning with copy-paste command to switch directories
- Automatically detects Claude mode (Max vs API) for correct command suggestion
- Backward compatible: Branch strategy and legacy specs work unchanged

## What's New in v3.1

**Kanban Auto-Sync:**
- New stories added after kanban creation are now automatically detected
- Entry point syncs new `user-story-*.md` and `bug-*.md` files to existing kanban
- No more "forgotten" tasks when using `/add-bug` or `/add-todo` mid-session

## What's New in v3.0

**Phase Files Updated:**
- `spec-phase-3.md` → Direct Execution (no sub-agents)
- `backlog-phase-2.md` → Direct Execution (no sub-agents)

**Key Changes:**
- Main agent implements stories directly
- Skills auto-load via glob patterns
- Self-review replaces separate review agents
- Self-learning mechanism added

---

# Task Execution Entry Point

## Overview

Lightweight router that detects current state and loads ONLY the relevant phase.
This reduces context usage by ~70-80% compared to loading the full workflow.

**Phase Files Location:** `agent-os/workflows/core/execute-tasks/`

---

## Execution Mode Detection

<mode_detection>
  WHEN /execute-tasks is invoked:

  1. CHECK: Was a parameter provided?

     IF parameter = "backlog":
       SET: EXECUTION_MODE = "backlog"
       GOTO: Backlog State Detection

     ELSE IF parameter = [spec-name]:
       SET: EXECUTION_MODE = "spec"
       SET: SELECTED_SPEC = [spec-name]
       GOTO: Spec State Detection

     ELSE (no parameter):
       GOTO: Auto-Detection

  <auto_detection>
    CHECK: Are there active kanban boards?

    ```bash
    # Check for active spec kanbans
    SPEC_KANBANS=$(ls agent-os/specs/*/kanban-board.md 2>/dev/null | head -5)

    # Check for active backlog kanban (today)
    TODAY=$(date +%Y-%m-%d)
    BACKLOG_KANBAN=$(ls agent-os/backlog/kanban-${TODAY}.md 2>/dev/null)

    # Check for pending backlog stories
    BACKLOG_STORIES=$(ls agent-os/backlog/user-story-*.md agent-os/backlog/bug-*.md 2>/dev/null | wc -l)
    ```

    IF active kanban exists (spec or backlog):
      DETECT: Which kanban is active
      RESUME: That execution automatically

    ELSE IF backlog has stories AND specs exist:
      ASK via AskUserQuestion:
      "What would you like to execute?
      1. Execute Backlog ([N] quick tasks)
      2. Execute Spec (select from available)
      3. View status only"

    ELSE IF only backlog has stories:
      SET: EXECUTION_MODE = "backlog"

    ELSE IF only specs exist:
      SET: EXECUTION_MODE = "spec"

    ELSE:
      ERROR: "No tasks to execute. Use /add-todo or /create-spec first."
  </auto_detection>
</mode_detection>

---

## Backlog State Detection

<backlog_routing>
  USE: date-checker to get current date (YYYY-MM-DD)

  CHECK: Does today's kanban exist?
  ```bash
  ls agent-os/backlog/kanban-${TODAY}.md 2>/dev/null
  ```

  IF NO kanban:
    LOAD: @agent-os/workflows/core/execute-tasks/backlog-phase-1.md
    STOP: After loading

  IF kanban exists:
    READ: agent-os/backlog/kanban-${TODAY}.md
    EXTRACT: "Current Phase" from Resume Context

    <kanban_sync>
      ## Auto-Sync: Check for New Stories (v3.1)

      BEFORE loading phase, sync any new stories added after kanban creation:

      1. LIST: All story files in backlog folder (excluding done/)
         ```bash
         ls agent-os/backlog/user-story-*.md agent-os/backlog/bug-*.md 2>/dev/null
         ```

      2. EXTRACT: Story IDs already in kanban
         - Parse "## Backlog", "## In Progress", "## Done" sections
         - Collect all Story IDs listed

      3. COMPARE: Find new stories
         FOR EACH file in backlog folder:
           EXTRACT: Story ID from filename (e.g., "US-001" from "user-story-US-001-title.md")
           IF Story ID NOT in kanban:
             ADD to NEW_STORIES list

      4. IF NEW_STORIES is not empty:
         READ: Each new story file
         EXTRACT: Title, Type, Priority, Points

         UPDATE: kanban-${TODAY}.md
           - ADD new stories to "## Backlog" table
           - UPDATE "Board Status" totals
           - ADD Change Log entry: "{TIMESTAMP} | Synced {N} new stories: {STORY_IDS}"

         INFORM user:
         "📥 **Kanban Sync:** Added {N} new stories to today's board: {STORY_IDS}"
    </kanban_sync>

    | Current Phase | Load Phase File |
    |---------------|-----------------|
    | 1-complete | backlog-phase-2.md |
    | story-complete | backlog-phase-2.md |
    | all-stories-done | backlog-phase-3.md |
    | complete | INFORM: "Backlog execution complete for today" |

    LOAD: Appropriate phase file
    STOP: After loading
</backlog_routing>

---

## Spec State Detection

<spec_routing>
  IF SELECTED_SPEC not set:
    LIST: Available specs
    ```bash
    ls -1 agent-os/specs/ | sort -r
    ```
    IF 1 spec: SET SELECTED_SPEC automatically
    IF multiple: ASK user via AskUserQuestion

  CHECK: Does kanban-board.md exist?
  ```bash
  ls agent-os/specs/${SELECTED_SPEC}/kanban-board.md 2>/dev/null
  ```

  IF NO kanban-board.md:
    LOAD: @agent-os/workflows/core/execute-tasks/spec-phase-1.md
    STOP: After loading

  IF kanban-board.md exists:
    READ: agent-os/specs/${SELECTED_SPEC}/kanban-board.md
    EXTRACT: "Current Phase" from Resume Context
    EXTRACT: "Git Strategy" from Resume Context (if present)
    EXTRACT: "Worktree Path" from Resume Context (if present)

    <cwd_check>
      ## Worktree CWD Check (v3.3)

      **Purpose:** Ensure agent is running in correct directory for worktree-based specs.

      1. GET: Git Strategy from Resume Context
         - "worktree" → Worktree strategy active
         - "branch" → Branch strategy (no CWD check needed)
         - Not set or "(none)" → Legacy spec (no CWD check needed)

      2. IF Git Strategy = "worktree":
         GET: Worktree Path from Resume Context (e.g., `../projekt-x-worktrees/my-feature`)
         GET: Current working directory

         ```bash
         # Get current working directory
         CWD=$(pwd)

         # Get worktree basename for comparison
         WORKTREE_BASENAME=$(basename "${WORKTREE_PATH}")
         CWD_BASENAME=$(basename "${CWD}")
         ```

         COMPARE: Check if CWD is the correct worktree
         - Compare directory basenames (feature name)
         - Verify parent directory ends with "-worktrees"

         ```bash
         # Check if we're in the right worktree
         CWD_PARENT=$(basename "$(dirname "${CWD}")")

         # Valid if: basename matches AND parent ends with "-worktrees"
         if [[ "${CWD_BASENAME}" == "${WORKTREE_BASENAME}" ]] && \
            [[ "${CWD_PARENT}" == *-worktrees ]]; then
           echo "In correct worktree"
         fi
         ```

         IF CWD is NOT the correct worktree:
           DETECT: Claude mode for correct command

           <mode_detection_logic>
             **Determine Claude startup command:**

             The agent should check its startup context:
             - If running with Claude Max account → use `claude`
             - If running with API token (GLM/Anthropic API) → use `claude --dangerously-skip-permissions`

             **Note:** In practice, check the environment or startup flags.
             For simplicity, assume API mode if `ANTHROPIC_API_KEY` is set or
             if started with `--dangerously-skip-permissions`.
           </mode_detection_logic>

           SET: CLAUDE_CMD based on detected mode
           - Claude Max: `claude`
           - API Mode: `claude --dangerously-skip-permissions`

           OUTPUT:
           ---
           ## ⚠️ Wrong Working Directory!

           **You are not in the correct worktree directory.**

           | Current | Expected |
           |---------|----------|
           | `{CWD}` | `{WORKTREE_PATH}` |

           **To continue, run this command:**
           ```bash
           cd {WORKTREE_PATH} && {CLAUDE_CMD}
           ```

           Then run `/execute-tasks` again.

           ---

           **STOP:** Execution cannot continue from wrong directory.
           ---

           STOP: Do not proceed - wrong working directory

         ELSE (CWD is correct worktree):
           CONTINUE: Proceed to phase loading

      3. IF Git Strategy = "branch" OR not set:
         CONTINUE: No CWD check needed, proceed normally
    </cwd_check>

    <system_stories_check>
      ## System Stories Detection (v3.6)

      **Before routing, check for System Stories:**

      1. CHECK: Do System Stories exist in this spec?
         ```bash
         ls agent-os/specs/${SELECTED_SPEC}/stories/story-997*.md \
            agent-os/specs/${SELECTED_SPEC}/stories/story-998*.md \
            agent-os/specs/${SELECTED_SPEC}/stories/story-999*.md 2>/dev/null
         ```

      2. IF System Stories exist:
         EXTRACT: Status of each System Story from story files

         SET: HAS_SYSTEM_STORIES = true
         SET: SYSTEM_STORIES_DONE = true if ALL (997, 998, 999) have Status: Done

      3. IF no System Stories:
         SET: HAS_SYSTEM_STORIES = false
         NOTE: Use legacy Phase 4.5 and 5 routing
    </system_stories_check>

    | Current Phase | Condition | Load Phase File |
    |---------------|-----------|-----------------|
    | 1-complete | - | spec-phase-2.md |
    | 2-complete | - | spec-phase-3.md |
    | story-complete | - | spec-phase-3.md |
    | all-stories-done | HAS_SYSTEM_STORIES = true AND NOT SYSTEM_STORIES_DONE | spec-phase-3.md (execute System Stories) |
    | all-stories-done | HAS_SYSTEM_STORIES = false | spec-phase-4-5.md (legacy) |
    | all-stories-done | SYSTEM_STORIES_DONE = true | spec-phase-5.md (legacy check only) |
    | 5-ready | - | spec-phase-5.md |
    | complete | - | INFORM: "Spec execution complete. PR created." |

    **Routing Logic (v3.6):**

    ```
    IF Current Phase = "all-stories-done":
      IF HAS_SYSTEM_STORIES AND story-997/998/999 NOT all Done:
        # System Stories pending - continue Phase 3
        LOAD: spec-phase-3.md

      ELSE IF NOT HAS_SYSTEM_STORIES:
        # Legacy spec - use old Phase 4.5
        LOAD: spec-phase-4-5.md

      ELSE IF SYSTEM_STORIES_DONE:
        # System Stories done - legacy check
        LOAD: spec-phase-5.md (will skip to completion)
    ```

    LOAD: Appropriate phase file
    STOP: After loading
</spec_routing>

---

## Phase File Reference

| Mode | Phase | File | Purpose |
|------|-------|------|---------|
| Spec | 1 | spec-phase-1.md | Initialize + Kanban |
| Spec | 2 | spec-phase-2.md | Git Worktree |
| Spec | 3 | spec-phase-3.md | Execute Story |
| Spec | 4.5 | spec-phase-4-5.md | Integration Validation |
| Spec | 5 | spec-phase-5.md | Finalize + PR |
| Backlog | 1 | backlog-phase-1.md | Daily Kanban |
| Backlog | 2 | backlog-phase-2.md | Execute Story |
| Backlog | 3 | backlog-phase-3.md | Daily Summary |

---

## Shared Resources

Common resources used across phases:

| Resource | Location |
|----------|----------|
| Resume Context Schema | shared/resume-context.md |
| Error Handling | shared/error-handling.md |
| Skill Extraction | shared/skill-extraction.md |
