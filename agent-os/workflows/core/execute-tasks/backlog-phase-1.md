---
description: Backlog Phase 1 - Initialize Daily Kanban
version: 3.2
---

# Backlog Phase 1: Initialize Daily Kanban

## Purpose
Create today's Kanban Board for backlog execution.

## Entry Condition
- EXECUTION_MODE = "backlog"
- No kanban-[TODAY].md exists

## Actions

<step name="get_today_date">
  USE: date-checker to get current date
  SET: TODAY = YYYY-MM-DD
</step>

<step name="collect_ready_stories">
  LIST: All story files in backlog (excluding done/ folder)
  ```bash
  ls agent-os/backlog/user-story-*.md agent-os/backlog/bug-*.md 2>/dev/null
  ```
  NOTE: Completed stories are in agent-os/backlog/done/ and won't be listed

  FOR EACH story file:
    READ: Story file
    CHECK: DoR status (all [x] checked?)
    IF DoR complete: ADD to ready_stories list
    ELSE: ADD to blocked_stories list
</step>

<step name="create_daily_kanban" subagent="file-creator">
  USE: file-creator subagent

  PROMPT: "Create daily kanban board for backlog execution.

  Output: agent-os/backlog/kanban-{TODAY}.md

  Content Structure:
  ```markdown
  # Backlog Kanban - {TODAY}

  > Daily task execution board
  > Created: {TODAY}

  ## Resume Context

  | Field | Value |
  |-------|-------|
  | **Execution Mode** | Backlog |
  | **Current Phase** | 1-complete |
  | **Next Phase** | 2 - Execute Story |
  | **Current Story** | None |
  | **Last Action** | Daily kanban created |
  | **Next Action** | Execute first story |

  ---

  ## Board Status

  - **Total Stories**: {TOTAL}
  - **Completed**: 0
  - **In Progress**: 0
  - **Backlog**: {READY_COUNT}
  - **Blocked**: {BLOCKED_COUNT}

  ---

  ## Backlog

  | Story ID | Title | Type | Priority | Points |
  |----------|-------|------|----------|--------|
  {READY_STORIES_TABLE}

  ---

  ## In Progress

  <!-- Currently executing story -->

  ---

  ## Done

  <!-- Completed stories today -->

  ---

  ## Blocked

  {BLOCKED_STORIES_TABLE}

  ---

  ## Change Log

  | Time | Action |
  |------|--------|
  | {TIMESTAMP} | Daily kanban created with {TOTAL} stories |
  ```

  Replace placeholders with actual values."

  WAIT: For file-creator completion
</step>

## Phase Completion

<phase_complete>
  OUTPUT to user:
  ---
  ## Backlog Phase 1 Complete: Daily Kanban Created

  **Date:** {TODAY}
  **Stories Ready:** {READY_COUNT}
  **Blocked:** {BLOCKED_COUNT}

  **Kanban:** agent-os/backlog/kanban-{TODAY}.md

  **Next Phase:** Execute First Story

  ---
  **To continue, run:**
  ```
  /clear
  /execute-tasks backlog
  ```
  ---

  STOP: Do not proceed to Backlog Phase 2
</phase_complete>
