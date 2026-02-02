# RFZ CLI Boilerplate

> WARNING: DO NOT COPY THESE FILES TO PROJECT ROOT

This directory contains starter code templates for the RFZ CLI project.

## Purpose

These files serve as REFERENCE MATERIAL showing:
- Recommended project structure
- Example implementations
- API patterns to follow

## Usage

1. READ the boilerplate to understand patterns
2. WRITE your own implementation inspired by these patterns
3. DO NOT copy files directly to project root

## Why Not Copy?

1. Boilerplate may contain placeholder code that needs adaptation
2. Implementation should evolve independently
3. Copying causes duplicate code and merge conflicts
4. API changes in real code will break copied boilerplate (type mismatches)

## Directory Structure

```
boilerplate/
├── cmd/           - Example entry point (main.go)
├── configs/       - Example configuration (components.yaml)
├── internal/      - Example internal packages
│   ├── app/       - Application layer
│   ├── domain/    - Domain models
│   ├── infra/     - Infrastructure adapters
│   ├── service/   - Business services
│   └── ui/        - UI components, screens, modals
├── testdata/      - Test data fixtures
├── go.mod         - Module definition
└── Makefile       - Build commands
```

## Historical Context

In February 2026, these boilerplate files were accidentally copied to the project root,
causing compile errors due to type mismatches between the boilerplate code and the
evolved implementation. This was fixed in bugfix spec `2026-02-02-bugfix-boilerplate-in-project-root`.

The safeguards in this README and in CLAUDE.md were added to prevent this from happening again.
