@CONTRIBUTING.md

# deck

Terminal slide presenter built with [Bubbletea v2](https://charm.land/bubbletea). Present markdown files as navigable slide decks with live hot reload, progressive reveal, column layouts, code execution, and vim-style navigation.

## Quick Start

```bash
# Build
make build

# Run with example
make run

# Install globally
make install

# Test
go test ./... -count=1

# Lint
golangci-lint run ./...

# Format
gofmt -l -w .
```

## Architecture

`deck` is a single-binary TUI application with no server component.

Key packages under `internal/`:

| Package | Responsibility |
|---------|---------------|
| `app` | Root Bubbletea model, wires navigation, rendering, and key handling |
| `model` | Slide data types (Slide, Chunk, Frontmatter, ColumnLayout) |
| `parse` | Markdown parser — splits on `---` delimiters, extracts YAML frontmatter |
| `render` | Glamour-based markdown renderer with lipgloss layout |
| `nav` | Slide navigation state machine |
| `search` | Incremental search across slide content |
| `code` | Code block extraction and execution |
| `diff` | Unified diff rendering |
| `watch` | fsnotify-based file watcher for hot reload |
| `version` | Build-time version stamping via ldflags |

Data flow: `main.go` loads content from stdin or file path → `parse` splits slides → `app` drives the Bubbletea event loop → `render` produces terminal output.

## Build Variables

Version info (`Version`, `Commit`, `Date`) is injected into `internal/version` via `-ldflags` at build time. `make build` handles this automatically.

## Slide Format

See the README for slide format and frontmatter options. Parsing lives in `internal/parse`; frontmatter fields map to `internal/model.Frontmatter`.

## Hooks

Hooks are configured in `.claude/hooks/`.

- `session-start.sh` — Installs Go and golangci-lint in Claude Code Web ephemeral containers.
- `pre-tool-use.sh` — Logs tool calls for auditing (only when `$CLAUDE_PROJECT_DIR` is set).
- `post-tool-use.sh` — Runs `gofmt` on any `.go` files modified by the Write or Edit tools.
- `stop.sh` — Runs `go vet ./...` and `golangci-lint run ./...` after each Claude response; exits 2 on issues (blocking).
- `subagent-stop.sh` — Stub for subagent coordination.
- `session-end.sh` — Stub for session cleanup.
