# Documentation

When making changes to this codebase, keep documentation current.

## What to Document Here

Add rule files to this directory for domain-specific guidance. Examples:

- `RENDERING.md` — lipgloss layout conventions, color themes
- `TESTING.md` — test patterns for Bubbletea models
- `RELEASE.md` — release process and versioning

## Patterns and Conventions

### Package organization

All business logic lives under `internal/`. The `main.go` entry point is intentionally thin — it loads content and starts the Bubbletea program.

### Bubbletea model conventions

- Each sub-model implements `tea.Model` (`Init`, `Update`, `View`)
- Messages are defined as unexported types in the package that produces them
- Key bindings use `charm.land/bubbletea/v2` key handling

### Adding a new internal package

1. Create `internal/<name>/<name>.go`
2. Export only the types and functions used by other packages
3. Write a `<name>_test.go` alongside — table-driven tests preferred
4. Register the package in `internal/app/app.go` if it needs to be wired into the model

### Version stamping

Do not hardcode version strings. Use `internal/version.Short()` or `version.Info()`. Build-time values are injected via `make build`.

## Architecture decisions

- **No CGO**: The binary must be cross-compilable without C toolchain. Avoid packages that require CGO.
- **No config files**: deck has zero configuration files. All behavior is driven by slide frontmatter and CLI flags.
- **Embedded tutorial**: `tutorial.md` is embedded at compile time via `//go:embed`. Keep it self-contained.
