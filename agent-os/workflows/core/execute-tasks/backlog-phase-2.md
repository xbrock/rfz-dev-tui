---
description: Backlog Phase 2 - Execute one backlog story (Direct Execution v3.0)
version: 3.0
---

# Backlog Phase 2: Execute Story (Direct Execution)

## What's New in v3.0

- **No Sub-Agent Delegation**: Main agent implements story directly
- **Skills Load Automatically**: Via glob patterns in .claude/skills/
- **Self-Review**: DoD checklist instead of separate review
- **Self-Learning**: Updates dos-and-donts.md when learning

## Purpose

Execute ONE backlog story. Simpler than spec execution (no git worktree, no integration phase).

## Entry Condition

- kanban-[TODAY].md exists
- Resume Context shows: Phase 1-complete OR story-complete
- Stories remain in Backlog

## Actions

<step name="load_state">
  READ: agent-os/backlog/kanban-{TODAY}.md
  EXTRACT: Resume Context
  IDENTIFY: Next story from Backlog
</step>

<step name="story_selection">
  SELECT: First story from Backlog section
  (Backlog stories have no dependencies - execute in order)
</step>

<step name="update_kanban_in_progress">
  UPDATE: kanban-{TODAY}.md
    - MOVE: Selected story from Backlog to "In Progress"
    - UPDATE Board Status
    - SET Resume Context: Current Story = [story-id]
    - ADD Change Log entry
</step>

<step name="load_story">
  ### Load Story Details

  READ: Story file from agent-os/backlog/

  EXTRACT:
  - Story ID and Title
  - Feature description
  - Acceptance Criteria
  - DoD Checklist
  - Domain reference (if specified)

  NOTE: Skills load automatically when you edit matching files.
</step>

<step name="implement">
  ### Direct Implementation (v3.0)

  **The main agent implements the story directly.**

  <implementation_process>
    1. UNDERSTAND: Story requirements

    2. IMPLEMENT: The task
       - Create/modify files as needed
       - Skills load automatically when editing matching files
       - Keep it focused (backlog tasks are smaller)

    3. RUN: Tests
       - Ensure tests pass

    4. VERIFY: Acceptance criteria satisfied

    **This is a quick task:**
    - No extensive refactoring
    - Keep changes minimal
    - Focus on the specific requirement
  </implementation_process>
</step>

<step name="self_review">
  ### Self-Review with DoD Checklist

  <review_process>
    1. READ: DoD checklist from story

    2. VERIFY each item:
       - [ ] Implementation complete
       - [ ] Tests passing
       - [ ] Linter passes
       - [ ] Acceptance criteria met

    3. RUN: Completion Check commands from story

    IF all checks pass:
      PROCEED to self_learning_check
    ELSE:
      FIX issues and re-verify
  </review_process>
</step>

<step name="self_learning_check">
  ### Self-Learning Check (v3.0)

  <learning_detection>
    REFLECT: Did you learn something during implementation?

    IF YES:
      1. IDENTIFY: The learning
      2. LOCATE: Target dos-and-donts.md file
         - Frontend: .claude/skills/frontend-[framework]/dos-and-donts.md
         - Backend: .claude/skills/backend-[framework]/dos-and-donts.md
         - DevOps: .claude/skills/devops-[stack]/dos-and-donts.md

      3. APPEND: Learning entry
         ```markdown
         ### [DATE] - [Short Title]
         **Context:** [What you were trying to do]
         **Issue:** [What didn't work]
         **Solution:** [What worked]
         ```

    IF NO learning:
      SKIP: No update needed
  </learning_detection>
</step>

<step name="move_story_to_done">
  MOVE: Story file to done/ folder
  ```bash
  mkdir -p agent-os/backlog/done
  mv agent-os/backlog/{STORY_FILE} agent-os/backlog/done/
  ```
  NOTE: This prevents the story from being picked up in future kanbans
</step>

<step name="story_commit" subagent="git-workflow">
  UPDATE: kanban-{TODAY}.md
    - MOVE: Story from "In Progress" to "Done"
    - UPDATE Board Status
    - ADD Change Log entry

  USE: git-workflow subagent
  "Commit backlog story {STORY_ID}:

  **WORKING_DIR:** {PROJECT_ROOT}

  - Message: fix/feat: {STORY_ID} [Story Title]
  - Stage all changes including:
    - Implementation files
    - Moved story file in done/
    - Any dos-and-donts.md updates
  - Push to current branch"
</step>

## Phase Completion

<phase_complete>
  CHECK: Remaining stories in Backlog

  IF backlog NOT empty:
    UPDATE: kanban-{TODAY}.md Resume Context
      - Current Phase: story-complete
      - Next Phase: 2 - Execute Story (next)
      - Current Story: None

    OUTPUT to user:
    ---
    ## Story Complete: {STORY_ID}

    **Progress:** {COMPLETED} of {TOTAL} stories today

    **Self-Learning:** [Updated/No updates]

    **Next:** Execute next story

    ---
    **To continue, run:**
    ```
    /clear
    /execute-tasks backlog
    ```
    ---

    STOP: Do not proceed to next story

  ELSE (backlog empty):
    UPDATE: kanban-{TODAY}.md Resume Context
      - Current Phase: all-stories-done
      - Next Phase: 3 - Daily Summary

    OUTPUT to user:
    ---
    ## All Backlog Stories Complete!

    **Today's Progress:** {TOTAL} stories completed

    **Next Phase:** Daily Summary

    ---
    **To continue, run:**
    ```
    /clear
    /execute-tasks backlog
    ```
    ---

    STOP: Do not proceed to Backlog Phase 3
</phase_complete>

---

## Quick Reference: v3.0 Changes

| v2.x (Sub-Agents) | v3.0 (Direct Execution) |
|-------------------|-------------------------|
| extract_skill_paths_backlog | Skills auto-load via globs |
| DELEGATE to dev-team__* | Main agent implements |
| quick_review (separate) | Self-review with DoD |
| - | self_learning_check (NEW) |

**Benefits:**
- Full context for each task
- Faster execution (no delegation overhead)
- Self-learning improves backlog workflow too
