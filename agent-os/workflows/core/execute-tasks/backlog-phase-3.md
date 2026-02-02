---
description: Backlog Phase 3 - Daily Summary
version: 3.3
---

# Backlog Phase 3: Daily Summary

## Purpose
Summarize today's work and update story-index.

## Entry Condition
- kanban-{TODAY}.md shows: all-stories-done
- All stories in Done column

## Actions

<step name="update_story_index">
  READ: agent-os/backlog/story-index.md

  FOR EACH completed story (from kanban Done column):
    UPDATE: Status = "Done"
    UPDATE: Completed date
    NOTE: Story files are already in agent-os/backlog/done/ (moved in Phase 2)

  UPDATE: Totals
    - Completed: +N
    - Ready for Execution: -N

  ADD: Today's kanban to Execution History table

  WRITE: Updated story-index.md
</step>

## Phase Completion

<phase_complete>
  UPDATE: kanban-{TODAY}.md Resume Context
    - Current Phase: complete
    - Next Phase: None

  OUTPUT to user:
  ---
  ## Daily Backlog Execution Complete!

  ### Today's Summary ({TODAY})

  **Completed Stories:**
  {LIST_OF_COMPLETED_STORIES}

  **Kanban:** agent-os/backlog/kanban-{TODAY}.md

  ### What's Next?
  1. Add more tasks: `/add-todo "[description]"`
  2. Create spec for larger features: `/create-spec`
  3. Tomorrow: `/execute-tasks backlog` for new daily kanban

  ---
  **Backlog execution finished for today.**
  ---
</phase_complete>
