@CONTRIBUTING.md

# deck

Terminal slide presenter built with [Bubbletea v2](https://charm.land/bubbletea). Present markdown files as navigable slide decks with live hot reload, progressive reveal, column layouts, code execution, and vim-style navigation.

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

### Architecture decisions

- **No CGO**: the binary must be cross-compilable without a C toolchain. Avoid packages that require CGO.
- **No config files**: deck has zero configuration files. All behavior is driven by slide frontmatter and CLI flags.
- **Embedded tutorial**: `tutorial.md` is embedded at compile time via `//go:embed`. Keep it self-contained.

## Conventions

### Package organization

All business logic lives under `internal/`. The `main.go` entry point is intentionally thin — it loads content and starts the Bubbletea program.

### Bubbletea model conventions

- Each sub-model implements `tea.Model` (`Init`, `Update`, `View`).
- Messages are defined as unexported types in the package that produces them.
- Key bindings match on `tea.KeyPressMsg` via `msg.String()` (e.g. `"q"`, `"ctrl+c"`).

### Adding a new internal package

1. Create `internal/<name>/<name>.go`.
2. Export only the types and functions used by other packages.
3. Write a `<name>_test.go` alongside — table-driven tests preferred.
4. Register the package in `internal/app/app.go` if it needs to be wired into the model.

## Build Variables

Version info (`Version`, `Commit`, `Date`) is injected into `internal/version` via `-ldflags` at build time. `make build` handles this automatically. Do not hardcode version strings — use `internal/version.Short()` or `version.Info()`.

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
