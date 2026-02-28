#!/bin/bash
# Hook: Stop
# Fires when Claude finishes a response.

set -euo pipefail

# Run go vet silently as a quick sanity check (non-blocking)
if command -v go &>/dev/null && [ -f go.mod ]; then
  go vet ./... >/dev/null 2>&1 || true
fi

exit 0
