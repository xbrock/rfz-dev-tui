# Spec Requirements Document

> Spec: Build Screens (Sprint 2.2)
> Created: 2026-02-07
> Status: In Progress

## Overview

Implement the three Build sub-screens within the "Build Components" navigation tab of the RFZ Developer CLI: Component Selection (multi-select list with category badges), Build Configuration Modal (Maven goals, profiles, port, skip tests, command preview), and Build Execution View (real-time progress tracking with per-component status, phases, and timing). Uses mocked domain data with an interface designed for Phase 3 real-data swap.

This is the core user journey of the RFZ CLI: select components -> configure build -> monitor execution.

## User Stories

See individual story files in `stories/` directory:
- story-001: Domain Model & Mock Data Provider
- story-002: Build Component Selection Screen
- story-003: Build Configuration Modal
- story-004: Build Execution View
- story-005: App Integration & Screen Transitions

System stories (auto-generated):
- story-997: Code Review
- story-998: Integration Validation
- story-999: Finalize PR

## Spec Scope

- Build Component Selection screen with full keyboard navigation and multi-select
- Build Configuration Modal with all 5 sections (Maven Goal, Profiles, Traktion Port, Build Options, Command Preview)
- Build Execution View with timed mock simulation and per-component progress
- Domain model with ComponentProvider interface and mock implementation
- Build cancellation and failure states
- Golden file visual tests for all prototype states (~25 screenshots)
- Status bar updates per view state

## Out of Scope

- Real Maven execution (Phase 3 - Sprint 3.1)
- Real component scanning/discovery (Phase 3 - Sprint 3.2)
- Build history/statistics (Future)
- Log viewer integration (Sprint 2.3)
- Configuration persistence (Phase 3)
- Component filtering/search
- Parallel build orchestration (Phase 3)

## Expected Deliverable

A fully interactive build workflow within the RFZ CLI TUI:
- Component selection matching prototype screenshots 10-17
- Build configuration modal matching screenshots 20-34
- Build execution view matching screenshots 40-49
- All keyboard shortcuts functional
- Golden file tests for all distinct UI states
- Mock build simulation with timed progress and random failures
- Clean domain model interface ready for Phase 3 integration

## Integration Requirements

> These integration tests will be executed automatically after all stories complete.

**Integration Type:** Frontend-only (TUI screens with mock data)

- [ ] **Integration Test 1:** Application builds without errors
   - Command: `cd /Users/lix/xapps/rfz-tui && go build ./...`
   - Validates: All new code compiles
   - Requires MCP: no

- [ ] **Integration Test 2:** All tests pass
   - Command: `cd /Users/lix/xapps/rfz-tui && go test ./...`
   - Validates: Unit tests and golden file tests pass
   - Requires MCP: no

- [ ] **Integration Test 3:** Lint passes
   - Command: `cd /Users/lix/xapps/rfz-tui && golangci-lint run ./...`
   - Validates: Code quality standards met
   - Requires MCP: no

- [ ] **Integration Test 4:** Build screen navigation works
   - Command: `cd /Users/lix/xapps/rfz-tui && go test ./internal/app/ -run TestBuild -update`
   - Validates: Build screen integrates with app navigation
   - Requires MCP: no

**Integration Scenarios:**
- [ ] Scenario 1: User navigates to Build Components, sees component list, selects components, opens config modal, starts build, monitors execution to completion
- [ ] Scenario 2: User starts build, cancels mid-execution, returns to component selection

## Spec Documentation

- Requirements: agent-os/specs/2026-02-07-build-screens/requirements-clarification.md
- Implementation Plan: agent-os/specs/2026-02-07-build-screens/implementation-plan.md
- Story Index: agent-os/specs/2026-02-07-build-screens/story-index.md
