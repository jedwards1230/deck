#!/bin/bash
# Hook: Stop
# Fires when Claude finishes a response.
# Runs go vet and golangci-lint as quality gates — exit 2 blocks on issues.

set -euo pipefail

if [ -f go.mod ]; then
  if command -v go &>/dev/null; then
    go vet ./... 2>&1 || { echo "go vet failed" >&2; exit 2; }
  fi

  if command -v golangci-lint &>/dev/null; then
    golangci-lint run ./... 2>&1 || { echo "golangci-lint failed" >&2; exit 2; }
  fi
fi

exit 0
