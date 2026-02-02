# RFZ Developer CLI Mission (Lite)

> Last Updated: 2026-02-02
> Version: 1.0.0

## Elevator Pitch

A Go-based Terminal User Interface (TUI) application built with the charm.land stack (Bubble Tea, Lip Gloss, Bubbles) that streamlines the Maven build workflow for RFZ developers at Deutsche Bahn. It provides an intuitive interface for component selection, build configuration, execution monitoring, and log viewing - transforming complex Maven commands into efficient keyboard-driven workflows.

## Value Proposition

The RFZ Developer CLI reduces cognitive load by eliminating the need to memorize complex Maven command combinations across a heterogeneous component ecosystem (Core, Simulators, Standalone). It provides real-time visual feedback during builds, ensures consistency across component types, and delivers keyboard-driven efficiency for experienced developers. Deep visual regression testing (97 UI states) prevents regressions from AI-assisted development.

## Target Users

RFZ developers at Deutsche Bahn who frequently rebuild Maven components and need to manage builds across multiple component types, comfortable with terminal interfaces and preferring efficient keyboard-driven workflows.

## Core Features

- **Build Components Screen** - Multi-select component list with category badges (Core, Simulator, Standalone)
- **Build Configuration Modal** - Maven goals, profiles, port selection, skip tests toggle
- **Build Execution View** - Real-time progress with status phases (Pending -> Compiling -> Testing -> Packaging -> Installing -> Done/Failed)
- **Log Viewer** - Component-based log viewing with filters (ALL, INFO, WARN, ERROR, DEBUG) and follow mode
- **Discover Screen** - Repository scanning, Git status (clean/dirty), branch info, last commit
- **Configuration Screen** - Scan paths management, component registry, detected components status

## Problem Solved

Developers struggle with complex Maven commands requiring specific profiles per component, inconsistent component structures, config generation complexity, multiple repositories, and varying directory structures - leading to errors, wasted time, and reliance on tribal knowledge.

## Differentiator

Purpose-built for the RFZ component ecosystem with deep understanding of component types and their profiles. Visual regression testing is a core feature (not afterthought) with 97 UI state coverage, enabling confident AI-assisted development without UI regressions.

## 6-Month Success Goal

- 80% adoption within RFZ development team
- 100% visual test coverage (97 UI states)
- 0% UI regressions after visual test suite implementation
- Build initiation time reduced from typing commands to < 10 seconds

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go 1.21+ | Language |
| Bubble Tea | TUI Framework |
| Lip Gloss | Styling |
| Bubbles | Components (list, table, viewport, progress, spinner, help) |
| charmbracelet/log | Structured logging with Lip Gloss styling |
| teatest | Visual Testing |
| GitHub Actions | CI/CD |

## Component Strategy

**Use from Bubbles:** list, progress, spinner, textinput, table, viewport, help, paginator
**Build custom:** TuiBox, TuiNavigation, TuiStatus, TuiButton, TuiModal, TuiRadio, TuiCheckbox, TuiTabs, TuiStatusBar, TuiTree

## ⚠️ CRITICAL RULE: Charm.land First - Custom Last

**MANDATORY for all agents:** Before implementing ANY visual element, check if charm.land provides it.

**Priority:** 1) Bubbles component → 2) Lip Gloss styling → 3) charmbracelet/log → 4) Custom (LAST RESORT)

**FORBIDDEN:** Custom border drawing, custom ANSI colors, manual padding, custom progress/spinner/list/table/scrolling.
**USE INSTEAD:** `lipgloss.Border()`, `lipgloss.Color()`, `lipgloss.Padding()`, `bubbles/*` components.

Even custom components MUST use Lip Gloss for ALL internal styling.

---

**Note:** This is a condensed version of the full product brief. For complete details, see `/Users/lix/xapps/rfz-tui/agent-os/product/product-brief.md`
