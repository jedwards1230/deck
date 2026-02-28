#!/bin/bash
# Hook: PostToolUse
# Fires after every tool call. Receives the tool output as JSON on stdin.
# IMPORTANT: This fires on EVERY tool call. Keep it fast and silent.

set -euo pipefail

output=$(cat)
tool_name=$(echo "$output" | jq -r '.tool_name // empty' 2>/dev/null || true)

# Auto-format Go files on write
if [[ "$tool_name" == "Write" || "$tool_name" == "Edit" ]]; then
  file_path=$(echo "$output" | jq -r '.tool_input.file_path // .tool_input.path // empty' 2>/dev/null || true)
  if [[ "$file_path" == *.go ]]; then
    gofmt -w "$file_path" 2>/dev/null || true
  fi
fi

exit 0
