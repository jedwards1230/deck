#!/bin/bash
# Hook: PreToolUse
# Fires before every tool call. Receives the tool input as JSON on stdin.
# IMPORTANT: This fires on EVERY tool call. Keep it fast and silent.
# Exit 0 = allow, exit 2 = block (stderr message shown to Claude).

set -euo pipefail

input=$(cat)

# If jq is not available, exit silently to keep the hook fast and quiet.
if ! command -v jq >/dev/null 2>&1; then
  exit 0
fi

tool_name=$(echo "$input" | jq -r '.tool_name // empty' 2>/dev/null || true)

# Log tool calls for auditing, if a project directory is available.
# Failures in logging must not block tool execution, so ignore errors.
if [ -n "${CLAUDE_PROJECT_DIR:-}" ]; then
    log_dir="$CLAUDE_PROJECT_DIR/.claude/logs"
    log_file="$log_dir/tool_calls.log"
    {
        mkdir -p "$log_dir"
        printf '%s\t%s\n' "$(date -Iseconds)" "${tool_name:-}" >>"$log_file"
    } || true
fi

exit 0
