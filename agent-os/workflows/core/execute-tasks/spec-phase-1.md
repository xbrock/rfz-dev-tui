---
description: Spec Phase 1 - Initialize and create Kanban Board
version: 3.2
---

# Spec Phase 1: Initialize

## What's New in v3.2

**Hybrid Template Lookup:**
- Templates are now searched in order: local → global
- Local: `agent-os/templates/docs/`
- Global: `~/.agent-os/templates/docs/`
- Fixes "template not found" issues for projects without local templates

## What's New in v3.1

**Integration Context:**
- Creates `integration-context.md` for cross-story context preservation
- Enables proper integration when stories execute in separate sessions

## Purpose
Select specification and create Kanban Board. One-time setup phase.

## Entry Condition
- No kanban-board.md exists for target spec

## Actions

<step name="spec_selection">
  CHECK: Did user provide spec name as parameter?

  IF parameter provided:
    VALIDATE: agent-os/specs/[spec-name]/ exists
    SET: SELECTED_SPEC = [spec-name]

  ELSE:
    LIST: Available specs
    ```bash
    ls -1 agent-os/specs/ | sort -r
    ```

    IF 1 spec: CONFIRM with user
    IF multiple: ASK user via AskUserQuestion
</step>

<step name="create_kanban_board" subagent="file-creator">
  USE: file-creator subagent

  PROMPT: "Create kanban board for spec: agent-os/specs/{SELECTED_SPEC}/

  Source: Parse story files in stories/ directory
  Output: agent-os/specs/{SELECTED_SPEC}/kanban-board.md

  **TEMPLATE LOOKUP (Hybrid - v3.2):**
  Search for template in this order:
  1. Local: agent-os/templates/docs/kanban-board-template.md
  2. Global: ~/.agent-os/templates/docs/kanban-board-template.md

  Use the FIRST one found.

  **CRITICAL: TEMPLATE COMPLIANCE**
  You MUST use the EXACT format from the template. Do NOT invent your own format.
  - READ the template file FIRST
  - COPY the structure EXACTLY
  - Only replace {{PLACEHOLDER}} variables with actual values

  STEPS:
  1. FIND template using hybrid lookup (local first, then global)
  2. READ the template file
  2. LIST all story files in agent-os/specs/{SELECTED_SPEC}/stories/
  3. FOR EACH story file:
     - READ the story file
     - VALIDATE DoR:
       * CHECK: All DoR checkboxes are marked [x]
       * IF any [ ] unchecked: STORY_STATUS = ⚠️ Blocked
       * IF all [x]: STORY_STATUS = ✅ Ready
     - EXTRACT: Story ID, Title, Type, Dependencies, Points

  4. CREATE kanban board by COPYING template structure and replacing:
     - {{SPEC_NAME}} → Spec folder name
     - {{TOTAL_STORIES}} → Count from story files
     - {{COMPLETED_COUNT}} → 0
     - {{IN_PROGRESS_COUNT}} → 0
     - {{IN_REVIEW_COUNT}} → 0
     - {{TESTING_COUNT}} → 0
     - {{BACKLOG_COUNT}} → Count of READY stories
     - {{BLOCKED_COUNT}} → Count of BLOCKED stories
     - {{CURRENT_PHASE}} → 1-complete
     - {{NEXT_PHASE}} → 2 - Git Worktree
     - {{SPEC_FOLDER}} → Spec folder name
     - {{WORKTREE_PATH}} → (pending)
     - {{GIT_BRANCH}} → (pending)
     - {{CURRENT_STORY}} → None
     - {{LAST_ACTION}} → Kanban board created
     - {{NEXT_ACTION}} → Setup git worktree
     - {{BACKLOG_STORIES}} → Story table rows
     - {{BLOCKED_STORIES}} → Blocked story table rows (if any)
     - Other sections → Empty or 'None'

  **IMPORTANT: Resume Context Format**
  The Resume Context MUST be a TABLE, not key-value pairs:
  ```markdown
  | Field | Value |
  |-------|-------|
  | **Current Phase** | 1-complete |
  | **Next Phase** | 2 - Git Worktree |
  ```
  NOT: **Current Phase:** 1-complete"

  WAIT: For file-creator completion
</step>

<step name="create_integration_context" subagent="file-creator">
  USE: file-creator subagent

  PROMPT: "Create integration context file for spec execution.

  Output: agent-os/specs/{SELECTED_SPEC}/integration-context.md

  Content:
  ```markdown
  # Integration Context

  > **Purpose:** Cross-story context preservation for multi-session execution.
  > **Auto-updated** after each story completion.
  > **READ THIS** before implementing the next story.

  ---

  ## Completed Stories

  | Story | Summary | Key Changes |
  |-------|---------|-------------|
  | - | No stories completed yet | - |

  ---

  ## New Exports & APIs

  ### Components
  <!-- New UI components created -->
  _None yet_

  ### Services
  <!-- New service classes/modules -->
  _None yet_

  ### Hooks / Utilities
  <!-- New hooks, helpers, utilities -->
  _None yet_

  ### Types / Interfaces
  <!-- New type definitions -->
  _None yet_

  ---

  ## Integration Notes

  <!-- Important integration information for subsequent stories -->
  _None yet_

  ---

  ## File Change Summary

  | File | Action | Story |
  |------|--------|-------|
  | - | No changes yet | - |
  ```
  "

  WAIT: For file-creator completion
</step>

## Phase Completion

<phase_complete>
  UPDATE: kanban-board.md Resume Context
    - Current Phase: 1-complete
    - Next Phase: 2 - Git Worktree

  OUTPUT to user:
  ---
  ## Phase 1 Complete: Initialization

  **Created:**
  - Kanban Board: agent-os/specs/{SELECTED_SPEC}/kanban-board.md
  - Integration Context: agent-os/specs/{SELECTED_SPEC}/integration-context.md
  - Stories loaded: [X] stories in Backlog

  **Next Phase:** Git Worktree Setup

  ---
  **To continue, run:**
  ```
  /clear
  /execute-tasks
  ```
  ---

  STOP: Do not proceed to Phase 2
</phase_complete>
