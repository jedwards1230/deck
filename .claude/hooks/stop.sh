#!/bin/bash
# Hook: Stop
# Fires when Claude finishes generating a response.
#
# Runs go vet and golangci-lint as final quality checks after Claude makes changes.
# Only runs if .go files are present in the project directory.

set -euo pipefail

PROJECT_DIR="${CLAUDE_PROJECT_DIR:-.}"

# Only run if this is a Go project
if [ ! -f "${PROJECT_DIR}/go.mod" ]; then
  exit 0
fi

cd "${PROJECT_DIR}"

# Ensure Go is on PATH (needed in ephemeral Claude Code Web containers)
export PATH="/usr/local/go/bin:${GOPATH:-/root/go}/bin:$PATH"

echo "[hook:stop] Running go vet..."
if command -v go &>/dev/null; then
  go vet ./... 2>&1 && echo "[hook:stop] go vet: OK" || {
    echo "[hook:stop] go vet found issues"
    exit 2
  }
else
  echo "[hook:stop] go not found — skipping go vet"
fi

echo "[hook:stop] Running golangci-lint..."
if command -v golangci-lint &>/dev/null; then
  golangci-lint run ./... 2>&1 | tail -30 || {
    echo "[hook:stop] golangci-lint found issues"
    exit 2
  }
  echo "[hook:stop] golangci-lint: OK"
else
  echo "[hook:stop] golangci-lint not found — skipping"
fi

exit 0
