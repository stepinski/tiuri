# AGENTS.md — tiuri

Terminal RPG in Go. Inspired by "Letter to the King" (Tonke Dragt).
Tiuri is a young knight on a long journey — castles, forests, guards, traps.
ASCII map, fluid movement, no over-engineering.

## Stack
- Go, tcell/v2 for rendering and input
- No BubbleTea — direct terminal control
- No external game engines

## Architecture
- `[][]rune` map
- `x, y` position for Tiuri
- Game loop: read input → update state → render
- Entities: Tiuri, walls, paths, guards, items

## Hard stops
- NEVER create files without approval
- One file per turn maximum
- Explain approach first, I write the code
- No over-engineering — simplest thing that works first

## Verification
```bash
go build ./... && go vet ./...
```

## Current milestone
Map renders. Tiuri moves. One obstacle to avoid.
