# Contributing to deck

deck is a terminal slide presenter built with [Bubbletea v2](https://charm.land/bubbletea). It presents markdown files as navigable slide decks with live hot reload, progressive reveal, column layouts, code execution, and vim-style navigation.

## Prerequisites

- [Go](https://go.dev/) (version from `go.mod`)
- [golangci-lint](https://golangci-lint.run/)
- [pre-commit](https://pre-commit.com/) (`pip install pre-commit` or `brew install pre-commit`)

The repository ships a devcontainer (`.devcontainer/`) with the full toolchain preinstalled — opening the repo in it is the quickest way to get a working environment.

## Build, test & lint

```bash
# Build
make build

# Run with the bundled example
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

## Before you open a PR

- Make sure all CI checks pass locally first — run the formatter, linter, and tests.
- Run `pre-commit run --all-files` (this repo uses pre-commit hooks).
- A local commit gate runs `go vet` and golangci-lint and will block until they pass — fix any findings before committing.

## Branching & commits

- Branch off `main`; never commit directly to `main`.
- Use [Conventional Commits](https://www.conventionalcommits.org/) prefixes (`feat:`, `fix:`, `docs:`, `chore:`, `refactor:`, `test:`, …).
- Sign your commits where possible (`git commit -S`).
- Keep each PR focused; delete dead code rather than commenting it out.

## Pull requests

- Open the PR against `main`.
- Every PR runs CI and an automated code review. Resolve **all** review threads before the PR is merged.
- A PR is merged once CI is green and the review is approved.

## Releases

Releases are opt-in. Before merging, add one of `semver:patch`, `semver:minor`, or `semver:major` to the PR to cut a release on merge; with no label, merging does not release. A release publishes a single immutable `vX.Y.Z` tag with AI-generated release notes.
