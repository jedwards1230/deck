#!/bin/bash
# Hook: PostToolUse
# Fires after every tool call. Receives the tool output as JSON on stdin.
#
# Runs gofmt on any .go files modified by Write or Edit tool calls.

set -euo pipefail

output=$(cat)
tool_name=$(echo "$output" | jq -r '.tool_name // empty' 2>/dev/null || true)

echo "[hook:post-tool-use] tool=$tool_name"

# Run gofmt on .go files written or edited
if [[ "$tool_name" == "Write" || "$tool_name" == "Edit" ]]; then
  file_path=$(echo "$output" | jq -r '.tool_input.file_path // empty' 2>/dev/null || true)
  if [[ "$file_path" == *.go && -f "$file_path" ]]; then
    echo "[hook:post-tool-use] Running gofmt on $file_path"
    if command -v gofmt &>/dev/null; then
      gofmt -w "$file_path" && echo "[hook:post-tool-use] gofmt OK"
    else
      echo "[hook:post-tool-use] gofmt not found — skipping"
    fi
  fi
fi

exit 0
