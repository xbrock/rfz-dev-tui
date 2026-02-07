# Spec Requirements Document

> Spec: Welcome & Navigation Screen
> Created: 2026-02-07
> Status: In Progress

## Overview

Implement the main App Shell (Header, Navigation sidebar, Content area, StatusBar) with the Welcome screen as default view and placeholder screens for Build/Logs/Discover/Configuration. This transforms the RFZ Developer CLI from component library demos into a real Bubble Tea application with full keyboard-driven navigation, screen switching, and terminal resize handling.

The App Shell establishes the foundational screen architecture (internal/app/ and internal/ui/screens/) that all subsequent Sprint 2.x screens will be built into, following the Layered Architecture with Ports & Adapters pattern defined in DEC-001 through DEC-006.

## User Stories

See: stories/ directory for individual story files

1. Story 001: Entry Point & Demo Rename
2. Story 002: App Shell Model with Layout
3. Story 003: Welcome Screen
4. Story 004: Screen Switching & Navigation
5. Story 005: Exit Confirmation Modal
6. Story 006: Visual Regression Tests

## Spec Scope

- App Shell layout (Header bar, Navigation sidebar, Content area, StatusBar)
- Welcome screen with ASCII art logo, badges, status, key hints
- Placeholder screens for Build Components, View Logs, Discover, Configuration
- Full keyboard navigation (global shortcuts, per-panel navigation, focus switching)
- Screen switching via navigation selection
- Exit confirmation modal dialog
- Terminal resize handling (dynamic layout)
- Real clock time in header (1-second tick)
- Entry point migration (cmd/rfz becomes real app)
- Demo rename (old gallery to cmd/rfz-components-demo)
- Visual regression tests for all key UI states

## Out of Scope

- Real Build Components screen (Sprint 2.2)
- Real Log Viewer screen (Sprint 2.3)
- Real Discover screen (Sprint 2.3)
- Real Configuration screen (Sprint 2.4)
- Maven execution, Git operations
- Configuration persistence
- Domain logic, service layer, infrastructure layer
- Theme customization

## Expected Deliverable

A fully functional App Shell that:
- Launches via `go run ./cmd/rfz` as a Bubble Tea alt-screen application
- Shows header with real clock time, navigation sidebar, welcome content, and status bar
- Allows keyboard navigation between menu items and screen switching
- Shows placeholder "Coming Soon" screens for Build/Logs/Discover/Config
- Shows exit confirmation modal when pressing q or selecting Exit
- Responds to terminal resize events
- Passes all visual regression golden file tests at 120x40
- Passes `golangci-lint` with no errors
- Compiles cleanly with `go build ./...`

## Integration Requirements

> These integration tests will be executed automatically after all stories complete.

**Integration Type:** Frontend-only

- [ ] **Integration Test 1:** Application builds without errors
  - Command: `cd /Users/lix/xapps/rfz-tui && go build ./cmd/rfz/...`
  - Validates: Main entry point compiles
  - Requires MCP: no

- [ ] **Integration Test 2:** Component demo still builds
  - Command: `cd /Users/lix/xapps/rfz-tui && go build ./cmd/rfz-components-demo/...`
  - Validates: Renamed demo still compiles
  - Requires MCP: no

- [ ] **Integration Test 3:** All tests pass
  - Command: `cd /Users/lix/xapps/rfz-tui && go test ./internal/... -count=1`
  - Validates: No regressions in existing tests + new tests pass
  - Requires MCP: no

- [ ] **Integration Test 4:** Lint passes
  - Command: `cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...`
  - Validates: Code quality
  - Requires MCP: no

**Integration Scenarios:**
- [ ] Scenario 1: Launch app, navigate between all menu items, see content switch for each, press q, confirm exit
- [ ] Scenario 2: Launch app, resize terminal, verify layout adjusts correctly

## Spec Documentation

- Requirements Clarification: agent-os/specs/2026-02-07-welcome-navigation-screen/requirements-clarification.md
- Implementation Plan: agent-os/specs/2026-02-07-welcome-navigation-screen/implementation-plan.md
- Story Index: agent-os/specs/2026-02-07-welcome-navigation-screen/story-index.md
- Stories: agent-os/specs/2026-02-07-welcome-navigation-screen/stories/
