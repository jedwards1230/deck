#!/bin/bash
# Devcontainer entrypoint - runs on every container start.
# Keep this idempotent (safe to run multiple times).

set -euo pipefail

echo "=== Deck Devcontainer Health Check ==="

# Tool version validation
check_tool() {
  local name="$1"
  local cmd="$2"
  if version=$($cmd 2>/dev/null); then
    printf "  %-18s %s\n" "$name" "$version"
  else
    printf "  %-18s %s\n" "$name" "NOT FOUND"
  fi
}

echo "Tools:"
check_tool "go" "go version"
check_tool "golangci-lint" "golangci-lint --version"
check_tool "yq" "yq --version"

# Install Go dependencies
if [ -f go.mod ]; then
  go mod download
fi

echo "=== Health Check Complete ==="
