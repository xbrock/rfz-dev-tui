---
description: Domain Knowledge - RFZ Developer CLI business context
globs:
  - "**/*.go"
  - "agent-os/product/**/*.md"
alwaysApply: false
version: 1.0.0
---

# Domain Knowledge Skill - RFZ Developer CLI

Business domain knowledge for the RFZ Developer CLI project at Deutsche Bahn.

## Quick Reference

### What is RFZ?

RFZ (Rangier-Funk-Zentrale) is a railway signaling and communication system used by Deutsche Bahn for train yard operations. RFZ developers work with a complex suite of components that must be built, tested, and deployed together.

### The Problem This CLI Solves

RFZ developers frequently need to:
1. Build multiple related Maven components
2. Track build progress across components
3. View logs to diagnose failures
4. Manage component discovery across repositories

The RFZ Developer CLI provides a streamlined terminal interface for these workflows.

## Domain Areas

| Domain Area | Document | Description |
|-------------|----------|-------------|
| Components | [components.md](components.md) | RFZ component types and dependencies |
| Build Process | [build-process.md](build-process.md) | Maven build workflow and phases |
| Developer Workflows | [workflows.md](workflows.md) | Common developer tasks |

## Key Domain Terms

| Term | Definition |
|------|------------|
| **RFZ** | Rangier-Funk-Zentrale - railway signaling system |
| **Component** | A Maven module in the RFZ ecosystem |
| **Core Component** | Essential component required by others |
| **Simulation Component** | Component for testing/simulation |
| **Standalone Component** | Independent component |
| **Build Config** | Maven profiles and options for build |
| **Registry** | List of known components (components.yaml) |

## User Personas

### Primary: RFZ Developer

- Works on Deutsche Bahn internal systems
- Comfortable with terminal interfaces
- Familiar with Maven builds
- Needs to build multiple components frequently
- Values efficiency and keyboard shortcuts

### Secondary: CI/CD System

- Automated builds in GitHub Actions
- Requires headless/non-interactive mode (future)
- Needs deterministic output for logs

## Business Rules

### Component Building

1. Components may have dependencies on other components
2. Dependencies should build before dependents
3. Failed component does not block independent components
4. Build configuration is remembered per-component (future)

### Component Discovery

1. Scan configured paths for pom.xml files
2. Detect component type from pom.xml content
3. Show Git status (branch, dirty/clean) for each
4. Registry file (components.yaml) stores known components

## Integration Points

### Maven

- Execute `mvn` commands with configurable goals
- Parse output for progress and errors
- Support profiles: `-P profile1,profile2`
- Support skip flags: `-DskipTests`

### Git

- Detect current branch
- Show dirty/clean status
- Show last commit info

### File System

- Scan directories for components
- Read/write configuration files
- Manage logs

## Notes

- Update this skill when business logic changes
- Add new domain terms as they emerge
- Document exceptions to rules
